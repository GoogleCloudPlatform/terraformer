package local

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/backend"
	"github.com/hashicorp/terraform/command/clistate"
	"github.com/hashicorp/terraform/configs/configload"
	"github.com/hashicorp/terraform/plans/planfile"
	"github.com/hashicorp/terraform/states/statemgr"
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/terraform/tfdiags"
	"github.com/zclconf/go-cty/cty"
)

// backend.Local implementation.
func (b *Local) Context(op *backend.Operation) (*terraform.Context, statemgr.Full, tfdiags.Diagnostics) {
	// Make sure the type is invalid. We use this as a way to know not
	// to ask for input/validate.
	op.Type = backend.OperationTypeInvalid

	if op.LockState {
		op.StateLocker = clistate.NewLocker(context.Background(), op.StateLockTimeout, b.CLI, b.Colorize())
	} else {
		op.StateLocker = clistate.NewNoopLocker()
	}

	ctx, _, stateMgr, diags := b.context(op)
	return ctx, stateMgr, diags
}

func (b *Local) context(op *backend.Operation) (*terraform.Context, *configload.Snapshot, statemgr.Full, tfdiags.Diagnostics) {
	var diags tfdiags.Diagnostics

	// Get the latest state.
	log.Printf("[TRACE] backend/local: requesting state manager for workspace %q", op.Workspace)
	s, err := b.StateMgr(op.Workspace)
	if err != nil {
		diags = diags.Append(errwrap.Wrapf("Error loading state: {{err}}", err))
		return nil, nil, nil, diags
	}
	log.Printf("[TRACE] backend/local: requesting state lock for workspace %q", op.Workspace)
	if err := op.StateLocker.Lock(s, op.Type.String()); err != nil {
		diags = diags.Append(errwrap.Wrapf("Error locking state: {{err}}", err))
		return nil, nil, nil, diags
	}
	log.Printf("[TRACE] backend/local: reading remote state for workspace %q", op.Workspace)
	if err := s.RefreshState(); err != nil {
		diags = diags.Append(errwrap.Wrapf("Error loading state: {{err}}", err))
		return nil, nil, nil, diags
	}

	// Initialize our context options
	var opts terraform.ContextOpts
	if v := b.ContextOpts; v != nil {
		opts = *v
	}

	// Copy set options from the operation
	opts.Destroy = op.Destroy
	opts.Targets = op.Targets
	opts.UIInput = op.UIIn

	// Load the latest state. If we enter contextFromPlanFile below then the
	// state snapshot in the plan file must match this, or else it'll return
	// error diagnostics.
	log.Printf("[TRACE] backend/local: retrieving local state snapshot for workspace %q", op.Workspace)
	opts.State = s.State()

	var tfCtx *terraform.Context
	var ctxDiags tfdiags.Diagnostics
	var configSnap *configload.Snapshot
	if op.PlanFile != nil {
		var stateMeta *statemgr.SnapshotMeta
		// If the statemgr implements our optional PersistentMeta interface then we'll
		// additionally verify that the state snapshot in the plan file has
		// consistent metadata, as an additional safety check.
		if sm, ok := s.(statemgr.PersistentMeta); ok {
			m := sm.StateSnapshotMeta()
			stateMeta = &m
		}
		log.Printf("[TRACE] backend/local: building context from plan file")
		tfCtx, configSnap, ctxDiags = b.contextFromPlanFile(op.PlanFile, opts, stateMeta)
		// Write sources into the cache of the main loader so that they are
		// available if we need to generate diagnostic message snippets.
		op.ConfigLoader.ImportSourcesFromSnapshot(configSnap)
	} else {
		log.Printf("[TRACE] backend/local: building context for current working directory")
		tfCtx, configSnap, ctxDiags = b.contextDirect(op, opts)
	}
	diags = diags.Append(ctxDiags)
	if diags.HasErrors() {
		return nil, nil, nil, diags
	}
	log.Printf("[TRACE] backend/local: finished building terraform.Context")

	// If we have an operation, then we automatically do the input/validate
	// here since every option requires this.
	if op.Type != backend.OperationTypeInvalid {
		// If input asking is enabled, then do that
		if op.PlanFile == nil && b.OpInput {
			mode := terraform.InputModeProvider
			mode |= terraform.InputModeVar
			mode |= terraform.InputModeVarUnset

			log.Printf("[TRACE] backend/local: requesting interactive input, if necessary")
			inputDiags := tfCtx.Input(mode)
			diags = diags.Append(inputDiags)
			if inputDiags.HasErrors() {
				return nil, nil, nil, diags
			}
		}

		// If validation is enabled, validate
		if b.OpValidation {
			log.Printf("[TRACE] backend/local: running validation operation")
			validateDiags := tfCtx.Validate()
			diags = diags.Append(validateDiags)
		}
	}

	return tfCtx, configSnap, s, diags
}

func (b *Local) contextDirect(op *backend.Operation, opts terraform.ContextOpts) (*terraform.Context, *configload.Snapshot, tfdiags.Diagnostics) {
	var diags tfdiags.Diagnostics

	// Load the configuration using the caller-provided configuration loader.
	config, configSnap, configDiags := op.ConfigLoader.LoadConfigWithSnapshot(op.ConfigDir)
	diags = diags.Append(configDiags)
	if configDiags.HasErrors() {
		return nil, nil, diags
	}
	opts.Config = config

	variables, varDiags := backend.ParseVariableValues(op.Variables, config.Module.Variables)
	diags = diags.Append(varDiags)
	if diags.HasErrors() {
		return nil, nil, diags
	}
	if op.Variables != nil {
		opts.Variables = variables
	}

	tfCtx, ctxDiags := terraform.NewContext(&opts)
	diags = diags.Append(ctxDiags)
	return tfCtx, configSnap, diags
}

func (b *Local) contextFromPlanFile(pf *planfile.Reader, opts terraform.ContextOpts, currentStateMeta *statemgr.SnapshotMeta) (*terraform.Context, *configload.Snapshot, tfdiags.Diagnostics) {
	var diags tfdiags.Diagnostics

	const errSummary = "Invalid plan file"

	// A plan file has a snapshot of configuration embedded inside it, which
	// is used instead of whatever configuration might be already present
	// in the filesystem.
	snap, err := pf.ReadConfigSnapshot()
	if err != nil {
		diags = diags.Append(tfdiags.Sourceless(
			tfdiags.Error,
			errSummary,
			fmt.Sprintf("Failed to read configuration snapshot from plan file: %s.", err),
		))
		return nil, snap, diags
	}
	loader := configload.NewLoaderFromSnapshot(snap)
	config, configDiags := loader.LoadConfig(snap.Modules[""].Dir)
	diags = diags.Append(configDiags)
	if configDiags.HasErrors() {
		return nil, snap, diags
	}
	opts.Config = config

	// A plan file also contains a snapshot of the prior state the changes
	// are intended to apply to.
	priorStateFile, err := pf.ReadStateFile()
	if err != nil {
		diags = diags.Append(tfdiags.Sourceless(
			tfdiags.Error,
			errSummary,
			fmt.Sprintf("Failed to read prior state snapshot from plan file: %s.", err),
		))
		return nil, snap, diags
	}
	if currentStateMeta != nil {
		// If the caller sets this, we require that the stored prior state
		// has the same metadata, which is an extra safety check that nothing
		// has changed since the plan was created. (All of the "real-world"
		// state manager implementstions support this, but simpler test backends
		// may not.)
		if currentStateMeta.Lineage != "" && priorStateFile.Lineage != "" {
			if priorStateFile.Serial != currentStateMeta.Serial || priorStateFile.Lineage != currentStateMeta.Lineage {
				diags = diags.Append(tfdiags.Sourceless(
					tfdiags.Error,
					"Saved plan is stale",
					"The given plan file can no longer be applied because the state was changed by another operation after the plan was created.",
				))
			}
		}
	}
	// The caller already wrote the "current state" here, but we're overriding
	// it here with the prior state. These two should actually be identical in
	// normal use, particularly if we validated the state meta above, but
	// we do this here anyway to ensure consistent behavior.
	opts.State = priorStateFile.State

	plan, err := pf.ReadPlan()
	if err != nil {
		diags = diags.Append(tfdiags.Sourceless(
			tfdiags.Error,
			errSummary,
			fmt.Sprintf("Failed to read plan from plan file: %s.", err),
		))
		return nil, snap, diags
	}

	variables := terraform.InputValues{}
	for name, dyVal := range plan.VariableValues {
		val, err := dyVal.Decode(cty.DynamicPseudoType)
		if err != nil {
			diags = diags.Append(tfdiags.Sourceless(
				tfdiags.Error,
				errSummary,
				fmt.Sprintf("Invalid value for variable %q recorded in plan file: %s.", name, err),
			))
			continue
		}

		variables[name] = &terraform.InputValue{
			Value:      val,
			SourceType: terraform.ValueFromPlan,
		}
	}
	opts.Variables = variables
	opts.Changes = plan.Changes
	opts.Targets = plan.TargetAddrs
	opts.ProviderSHA256s = plan.ProviderSHA256s

	tfCtx, ctxDiags := terraform.NewContext(&opts)
	diags = diags.Append(ctxDiags)
	return tfCtx, snap, diags
}

const validateWarnHeader = `
There are warnings related to your configuration. If no errors occurred,
Terraform will continue despite these warnings. It is a good idea to resolve
these warnings in the near future.

Warnings:
`
