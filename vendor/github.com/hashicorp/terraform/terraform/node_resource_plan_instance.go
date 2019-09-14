package terraform

import (
	"fmt"

	"github.com/hashicorp/terraform/plans"
	"github.com/hashicorp/terraform/providers"
	"github.com/hashicorp/terraform/states"

	"github.com/hashicorp/terraform/addrs"
	"github.com/zclconf/go-cty/cty"
)

// NodePlannableResourceInstance represents a _single_ resource
// instance that is plannable. This means this represents a single
// count index, for example.
type NodePlannableResourceInstance struct {
	*NodeAbstractResourceInstance
	ForceCreateBeforeDestroy bool
}

var (
	_ GraphNodeSubPath              = (*NodePlannableResourceInstance)(nil)
	_ GraphNodeReferenceable        = (*NodePlannableResourceInstance)(nil)
	_ GraphNodeReferencer           = (*NodePlannableResourceInstance)(nil)
	_ GraphNodeResource             = (*NodePlannableResourceInstance)(nil)
	_ GraphNodeResourceInstance     = (*NodePlannableResourceInstance)(nil)
	_ GraphNodeAttachResourceConfig = (*NodePlannableResourceInstance)(nil)
	_ GraphNodeAttachResourceState  = (*NodePlannableResourceInstance)(nil)
	_ GraphNodeEvalable             = (*NodePlannableResourceInstance)(nil)
)

// GraphNodeEvalable
func (n *NodePlannableResourceInstance) EvalTree() EvalNode {
	addr := n.ResourceInstanceAddr()

	// Eval info is different depending on what kind of resource this is
	switch addr.Resource.Resource.Mode {
	case addrs.ManagedResourceMode:
		return n.evalTreeManagedResource(addr)
	case addrs.DataResourceMode:
		return n.evalTreeDataResource(addr)
	default:
		panic(fmt.Errorf("unsupported resource mode %s", n.Config.Mode))
	}
}

func (n *NodePlannableResourceInstance) evalTreeDataResource(addr addrs.AbsResourceInstance) EvalNode {
	config := n.Config
	var provider providers.Interface
	var providerSchema *ProviderSchema
	var change *plans.ResourceInstanceChange
	var state *states.ResourceInstanceObject
	var configVal cty.Value

	return &EvalSequence{
		Nodes: []EvalNode{
			&EvalGetProvider{
				Addr:   n.ResolvedProvider,
				Output: &provider,
				Schema: &providerSchema,
			},

			&EvalReadState{
				Addr:           addr.Resource,
				Provider:       &provider,
				ProviderSchema: &providerSchema,

				Output: &state,
			},

			// If we already have a non-planned state then we already dealt
			// with this during the refresh walk and so we have nothing to do
			// here.
			&EvalIf{
				If: func(ctx EvalContext) (bool, error) {
					depChanges := false

					// Check and see if any of our dependencies have changes.
					changes := ctx.Changes()
					for _, d := range n.StateReferences() {
						ri, ok := d.(addrs.ResourceInstance)
						if !ok {
							continue
						}
						change := changes.GetResourceInstanceChange(ri.Absolute(ctx.Path()), states.CurrentGen)
						if change != nil && change.Action != plans.NoOp {
							depChanges = true
							break
						}
					}

					refreshed := state != nil && state.Status != states.ObjectPlanned

					// If there are no dependency changes, and it's not a forced
					// read because we there was no Refresh, then we don't need
					// to re-read. If any dependencies have changes, it means
					// our config may also have changes and we need to Read the
					// data source again.
					if !depChanges && refreshed {
						return false, EvalEarlyExitError{}
					}
					return true, nil
				},
				Then: EvalNoop{},
			},

			&EvalValidateSelfRef{
				Addr:           addr.Resource,
				Config:         config.Config,
				ProviderSchema: &providerSchema,
			},

			&EvalReadData{
				Addr:           addr.Resource,
				Config:         n.Config,
				Dependencies:   n.StateReferences(),
				Provider:       &provider,
				ProviderAddr:   n.ResolvedProvider,
				ProviderSchema: &providerSchema,
				ForcePlanRead:  true, // _always_ produce a Read change, even if the config seems ready
				OutputChange:   &change,
				OutputValue:    &configVal,
				OutputState:    &state,
			},

			&EvalWriteState{
				Addr:           addr.Resource,
				ProviderAddr:   n.ResolvedProvider,
				ProviderSchema: &providerSchema,
				State:          &state,
			},

			&EvalWriteDiff{
				Addr:           addr.Resource,
				ProviderSchema: &providerSchema,
				Change:         &change,
			},
		},
	}
}

func (n *NodePlannableResourceInstance) evalTreeManagedResource(addr addrs.AbsResourceInstance) EvalNode {
	config := n.Config
	var provider providers.Interface
	var providerSchema *ProviderSchema
	var change *plans.ResourceInstanceChange
	var state *states.ResourceInstanceObject

	return &EvalSequence{
		Nodes: []EvalNode{
			&EvalGetProvider{
				Addr:   n.ResolvedProvider,
				Output: &provider,
				Schema: &providerSchema,
			},

			&EvalReadState{
				Addr:           addr.Resource,
				Provider:       &provider,
				ProviderSchema: &providerSchema,

				Output: &state,
			},

			&EvalValidateSelfRef{
				Addr:           addr.Resource,
				Config:         config.Config,
				ProviderSchema: &providerSchema,
			},

			&EvalDiff{
				Addr:                addr.Resource,
				Config:              n.Config,
				CreateBeforeDestroy: n.ForceCreateBeforeDestroy,
				Provider:            &provider,
				ProviderAddr:        n.ResolvedProvider,
				ProviderSchema:      &providerSchema,
				State:               &state,
				OutputChange:        &change,
				OutputState:         &state,
			},
			&EvalCheckPreventDestroy{
				Addr:   addr.Resource,
				Config: n.Config,
				Change: &change,
			},
			&EvalWriteState{
				Addr:           addr.Resource,
				ProviderAddr:   n.ResolvedProvider,
				State:          &state,
				ProviderSchema: &providerSchema,
			},
			&EvalWriteDiff{
				Addr:           addr.Resource,
				ProviderSchema: &providerSchema,
				Change:         &change,
			},
		},
	}
}
