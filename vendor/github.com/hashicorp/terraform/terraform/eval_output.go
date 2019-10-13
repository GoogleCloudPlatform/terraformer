package terraform

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/plans"
	"github.com/hashicorp/terraform/states"
)

// EvalDeleteOutput is an EvalNode implementation that deletes an output
// from the state.
type EvalDeleteOutput struct {
	Addr addrs.OutputValue
}

// TODO: test
func (n *EvalDeleteOutput) Eval(ctx EvalContext) (interface{}, error) {
	state := ctx.State()
	if state == nil {
		return nil, nil
	}

	state.RemoveOutputValue(n.Addr.Absolute(ctx.Path()))
	return nil, nil
}

// EvalWriteOutput is an EvalNode implementation that writes the output
// for the given name to the current state.
type EvalWriteOutput struct {
	Addr      addrs.OutputValue
	Sensitive bool
	Expr      hcl.Expression
	// ContinueOnErr allows interpolation to fail during Input
	ContinueOnErr bool
}

// TODO: test
func (n *EvalWriteOutput) Eval(ctx EvalContext) (interface{}, error) {
	addr := n.Addr.Absolute(ctx.Path())

	// This has to run before we have a state lock, since evaluation also
	// reads the state
	val, diags := ctx.EvaluateExpr(n.Expr, cty.DynamicPseudoType, nil)
	// We'll handle errors below, after we have loaded the module.

	state := ctx.State()
	if state == nil {
		return nil, nil
	}

	changes := ctx.Changes() // may be nil, if we're not working on a changeset

	// handling the interpolation error
	if diags.HasErrors() {
		if n.ContinueOnErr || flagWarnOutputErrors {
			log.Printf("[ERROR] Output interpolation %q failed: %s", n.Addr.Name, diags.Err())
			// if we're continuing, make sure the output is included, and
			// marked as unknown. If the evaluator was able to find a type
			// for the value in spite of the error then we'll use it.
			n.setValue(addr, state, changes, cty.UnknownVal(val.Type()))
			return nil, EvalEarlyExitError{}
		}
		return nil, diags.Err()
	}

	n.setValue(addr, state, changes, val)

	return nil, nil
}

func (n *EvalWriteOutput) setValue(addr addrs.AbsOutputValue, state *states.SyncState, changes *plans.ChangesSync, val cty.Value) {
	if val.IsKnown() && !val.IsNull() {
		// The state itself doesn't represent unknown values, so we null them
		// out here and then we'll save the real unknown value in the planned
		// changeset below, if we have one on this graph walk.
		log.Printf("[TRACE] EvalWriteOutput: Saving value for %s in state", addr)
		stateVal := cty.UnknownAsNull(val)
		state.SetOutputValue(addr, stateVal, n.Sensitive)
	} else {
		log.Printf("[TRACE] EvalWriteOutput: Removing %s from state (it is now null)", addr)
		state.RemoveOutputValue(addr)
	}

	// If we also have an active changeset then we'll replicate the value in
	// there. This is used in preference to the state where present, since it
	// *is* able to represent unknowns, while the state cannot.
	if changes != nil {
		// For the moment we are not properly tracking changes to output
		// values, and just marking them always as "Create" or "Destroy"
		// actions. A future release will rework the output lifecycle so we
		// can track their changes properly, in a similar way to how we work
		// with resource instances.

		var change *plans.OutputChange
		if !val.IsNull() {
			change = &plans.OutputChange{
				Addr:      addr,
				Sensitive: n.Sensitive,
				Change: plans.Change{
					Action: plans.Create,
					Before: cty.NullVal(cty.DynamicPseudoType),
					After:  val,
				},
			}
		} else {
			change = &plans.OutputChange{
				Addr:      addr,
				Sensitive: n.Sensitive,
				Change: plans.Change{
					// This is just a weird placeholder delete action since
					// we don't have an actual prior value to indicate.
					// FIXME: Generate real planned changes for output values
					// that include the old values.
					Action: plans.Delete,
					Before: cty.NullVal(cty.DynamicPseudoType),
					After:  cty.NullVal(cty.DynamicPseudoType),
				},
			}
		}

		cs, err := change.Encode()
		if err != nil {
			// Should never happen, since we just constructed this right above
			panic(fmt.Sprintf("planned change for %s could not be encoded: %s", addr, err))
		}
		log.Printf("[TRACE] EvalWriteOutput: Saving %s change for %s in changeset", change.Action, addr)
		changes.RemoveOutputChange(addr) // remove any existing planned change, if present
		changes.AppendOutputChange(cs)   // add the new planned change
	}
}
