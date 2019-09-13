package local

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/backend"
	"github.com/hashicorp/terraform/states"
	"github.com/hashicorp/terraform/states/statefile"
	"github.com/hashicorp/terraform/states/statemgr"
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/terraform/tfdiags"
)

func (b *Local) opApply(
	stopCtx context.Context,
	cancelCtx context.Context,
	op *backend.Operation,
	runningOp *backend.RunningOperation) {
	log.Printf("[INFO] backend/local: starting Apply operation")

	var diags tfdiags.Diagnostics

	// If we have a nil module at this point, then set it to an empty tree
	// to avoid any potential crashes.
	if op.PlanFile == nil && !op.Destroy && !op.HasConfig() {
		diags = diags.Append(tfdiags.Sourceless(
			tfdiags.Error,
			"No configuration files",
			"Apply requires configuration to be present. Applying without a configuration "+
				"would mark everything for destruction, which is normally not what is desired. "+
				"If you would like to destroy everything, run 'terraform destroy' instead.",
		))
		b.ReportResult(runningOp, diags)
		return
	}

	// Setup our count hook that keeps track of resource changes
	countHook := new(CountHook)
	stateHook := new(StateHook)
	if b.ContextOpts == nil {
		b.ContextOpts = new(terraform.ContextOpts)
	}
	old := b.ContextOpts.Hooks
	defer func() { b.ContextOpts.Hooks = old }()
	b.ContextOpts.Hooks = append(b.ContextOpts.Hooks, countHook, stateHook)

	// Get our context
	tfCtx, _, opState, contextDiags := b.context(op)
	diags = diags.Append(contextDiags)
	if contextDiags.HasErrors() {
		b.ReportResult(runningOp, diags)
		return
	}

	// Setup the state
	runningOp.State = tfCtx.State()

	// If we weren't given a plan, then we refresh/plan
	if op.PlanFile == nil {
		// If we're refreshing before apply, perform that
		if op.PlanRefresh {
			log.Printf("[INFO] backend/local: apply calling Refresh")
			_, err := tfCtx.Refresh()
			if err != nil {
				diags = diags.Append(err)
				runningOp.Result = backend.OperationFailure
				b.ShowDiagnostics(diags)
				return
			}
		}

		// Perform the plan
		log.Printf("[INFO] backend/local: apply calling Plan")
		plan, err := tfCtx.Plan()
		if err != nil {
			diags = diags.Append(err)
			b.ReportResult(runningOp, diags)
			return
		}

		trivialPlan := plan.Changes.Empty()
		hasUI := op.UIOut != nil && op.UIIn != nil
		mustConfirm := hasUI && ((op.Destroy && (!op.DestroyForce && !op.AutoApprove)) || (!op.Destroy && !op.AutoApprove && !trivialPlan))
		if mustConfirm {
			var desc, query string
			if op.Destroy {
				if op.Workspace != "default" {
					query = "Do you really want to destroy all resources in workspace \"" + op.Workspace + "\"?"
				} else {
					query = "Do you really want to destroy all resources?"
				}
				desc = "Terraform will destroy all your managed infrastructure, as shown above.\n" +
					"There is no undo. Only 'yes' will be accepted to confirm."
			} else {
				if op.Workspace != "default" {
					query = "Do you want to perform these actions in workspace \"" + op.Workspace + "\"?"
				} else {
					query = "Do you want to perform these actions?"
				}
				desc = "Terraform will perform the actions described above.\n" +
					"Only 'yes' will be accepted to approve."
			}

			if !trivialPlan {
				// Display the plan of what we are going to apply/destroy.
				b.renderPlan(plan, runningOp.State, tfCtx.Schemas())
				b.CLI.Output("")
			}

			v, err := op.UIIn.Input(stopCtx, &terraform.InputOpts{
				Id:          "approve",
				Query:       query,
				Description: desc,
			})
			if err != nil {
				diags = diags.Append(errwrap.Wrapf("Error asking for approval: {{err}}", err))
				b.ReportResult(runningOp, diags)
				return
			}
			if v != "yes" {
				if op.Destroy {
					b.CLI.Info("Destroy cancelled.")
				} else {
					b.CLI.Info("Apply cancelled.")
				}
				runningOp.Result = backend.OperationFailure
				return
			}
		}
	}

	// Setup our hook for continuous state updates
	stateHook.StateMgr = opState

	// Start the apply in a goroutine so that we can be interrupted.
	var applyState *states.State
	var applyDiags tfdiags.Diagnostics
	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		log.Printf("[INFO] backend/local: apply calling Apply")
		_, applyDiags = tfCtx.Apply()
		// we always want the state, even if apply failed
		applyState = tfCtx.State()
	}()

	if b.opWait(doneCh, stopCtx, cancelCtx, tfCtx, opState) {
		return
	}

	// Store the final state
	runningOp.State = applyState
	err := statemgr.WriteAndPersist(opState, applyState)
	if err != nil {
		// Export the state file from the state manager and assign the new
		// state. This is needed to preserve the existing serial and lineage.
		stateFile := statemgr.Export(opState)
		if stateFile == nil {
			stateFile = &statefile.File{}
		}
		stateFile.State = applyState

		diags = diags.Append(b.backupStateForError(stateFile, err))
		b.ReportResult(runningOp, diags)
		return
	}

	diags = diags.Append(applyDiags)
	if applyDiags.HasErrors() {
		b.ReportResult(runningOp, diags)
		return
	}

	// If we've accumulated any warnings along the way then we'll show them
	// here just before we show the summary and next steps. If we encountered
	// errors then we would've returned early at some other point above.
	b.ShowDiagnostics(diags)

	// If we have a UI, output the results
	if b.CLI != nil {
		if op.Destroy {
			b.CLI.Output(b.Colorize().Color(fmt.Sprintf(
				"[reset][bold][green]\n"+
					"Destroy complete! Resources: %d destroyed.",
				countHook.Removed)))
		} else {
			b.CLI.Output(b.Colorize().Color(fmt.Sprintf(
				"[reset][bold][green]\n"+
					"Apply complete! Resources: %d added, %d changed, %d destroyed.",
				countHook.Added,
				countHook.Changed,
				countHook.Removed)))
		}

		// only show the state file help message if the state is local.
		if (countHook.Added > 0 || countHook.Changed > 0) && b.StateOutPath != "" {
			b.CLI.Output(b.Colorize().Color(fmt.Sprintf(
				"[reset]\n"+
					"The state of your infrastructure has been saved to the path\n"+
					"below. This state is required to modify and destroy your\n"+
					"infrastructure, so keep it safe. To inspect the complete state\n"+
					"use the `terraform show` command.\n\n"+
					"State path: %s",
				b.StateOutPath)))
		}
	}
}

// backupStateForError is called in a scenario where we're unable to persist the
// state for some reason, and will attempt to save a backup copy of the state
// to local disk to help the user recover. This is a "last ditch effort" sort
// of thing, so we really don't want to end up in this codepath; we should do
// everything we possibly can to get the state saved _somewhere_.
func (b *Local) backupStateForError(stateFile *statefile.File, err error) error {
	b.CLI.Error(fmt.Sprintf("Failed to save state: %s\n", err))

	local := statemgr.NewFilesystem("errored.tfstate")
	writeErr := local.WriteStateForMigration(stateFile, true)
	if writeErr != nil {
		b.CLI.Error(fmt.Sprintf(
			"Also failed to create local state file for recovery: %s\n\n", writeErr,
		))
		// To avoid leaving the user with no state at all, our last resort
		// is to print the JSON state out onto the terminal. This is an awful
		// UX, so we should definitely avoid doing this if at all possible,
		// but at least the user has _some_ path to recover if we end up
		// here for some reason.
		stateBuf := new(bytes.Buffer)
		jsonErr := statefile.Write(stateFile, stateBuf)
		if jsonErr != nil {
			b.CLI.Error(fmt.Sprintf(
				"Also failed to JSON-serialize the state to print it: %s\n\n", jsonErr,
			))
			return errors.New(stateWriteFatalError)
		}

		b.CLI.Output(stateBuf.String())

		return errors.New(stateWriteConsoleFallbackError)
	}

	return errors.New(stateWriteBackedUpError)
}

const stateWriteBackedUpError = `Failed to persist state to backend.

The error shown above has prevented Terraform from writing the updated state
to the configured backend. To allow for recovery, the state has been written
to the file "errored.tfstate" in the current working directory.

Running "terraform apply" again at this point will create a forked state,
making it harder to recover.

To retry writing this state, use the following command:
    terraform state push errored.tfstate
`

const stateWriteConsoleFallbackError = `Failed to persist state to backend.

The errors shown above prevented Terraform from writing the updated state to
the configured backend and from creating a local backup file. As a fallback,
the raw state data is printed above as a JSON object.

To retry writing this state, copy the state data (from the first { to the
last } inclusive) and save it into a local file called errored.tfstate, then
run the following command:
    terraform state push errored.tfstate
`

const stateWriteFatalError = `Failed to save state after apply.

A catastrophic error has prevented Terraform from persisting the state file
or creating a backup. Unfortunately this means that the record of any resources
created during this apply has been lost, and such resources may exist outside
of Terraform's management.

For resources that support import, it is possible to recover by manually
importing each resource using its id from the target system.

This is a serious bug in Terraform and should be reported.
`

const earlyStateWriteErrorFmt = `Error saving current state: %s

Terraform encountered an error attempting to save the state before cancelling
the current operation. Once the operation is complete another attempt will be
made to save the final state.
`
