package plans

import (
	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/states"
	"github.com/zclconf/go-cty/cty"
)

// Changes describes various actions that Terraform will attempt to take if
// the corresponding plan is applied.
//
// A Changes object can be rendered into a visual diff (by the caller, using
// code in another package) for display to the user.
type Changes struct {
	// Resources tracks planned changes to resource instance objects.
	Resources []*ResourceInstanceChangeSrc

	// Outputs tracks planned changes output values.
	//
	// Note that although an in-memory plan contains planned changes for
	// outputs throughout the configuration, a plan serialized
	// to disk retains only the root outputs because they are
	// externally-visible, while other outputs are implementation details and
	// can be easily re-calculated during the apply phase. Therefore only root
	// module outputs will survive a round-trip through a plan file.
	Outputs []*OutputChangeSrc
}

// NewChanges returns a valid Changes object that describes no changes.
func NewChanges() *Changes {
	return &Changes{}
}

func (c *Changes) Empty() bool {
	for _, res := range c.Resources {
		if res.Action != NoOp {
			return false
		}
	}
	return true
}

// ResourceInstance returns the planned change for the current object of the
// resource instance of the given address, if any. Returns nil if no change is
// planned.
func (c *Changes) ResourceInstance(addr addrs.AbsResourceInstance) *ResourceInstanceChangeSrc {
	addrStr := addr.String()
	for _, rc := range c.Resources {
		if rc.Addr.String() == addrStr && rc.DeposedKey == states.NotDeposed {
			return rc
		}
	}

	return nil
}

// ResourceInstanceDeposed returns the plan change of a deposed object of
// the resource instance of the given address, if any. Returns nil if no change
// is planned.
func (c *Changes) ResourceInstanceDeposed(addr addrs.AbsResourceInstance, key states.DeposedKey) *ResourceInstanceChangeSrc {
	addrStr := addr.String()
	for _, rc := range c.Resources {
		if rc.Addr.String() == addrStr && rc.DeposedKey == key {
			return rc
		}
	}

	return nil
}

// OutputValue returns the planned change for the output value with the
//  given address, if any. Returns nil if no change is planned.
func (c *Changes) OutputValue(addr addrs.AbsOutputValue) *OutputChangeSrc {
	addrStr := addr.String()
	for _, oc := range c.Outputs {
		if oc.Addr.String() == addrStr {
			return oc
		}
	}

	return nil
}

// SyncWrapper returns a wrapper object around the receiver that can be used
// to make certain changes to the receiver in a concurrency-safe way, as long
// as all callers share the same wrapper object.
func (c *Changes) SyncWrapper() *ChangesSync {
	return &ChangesSync{
		changes: c,
	}
}

// ResourceInstanceChange describes a change to a particular resource instance
// object.
type ResourceInstanceChange struct {
	// Addr is the absolute address of the resource instance that the change
	// will apply to.
	Addr addrs.AbsResourceInstance

	// DeposedKey is the identifier for a deposed object associated with the
	// given instance, or states.NotDeposed if this change applies to the
	// current object.
	//
	// A Replace change for a resource with create_before_destroy set will
	// create a new DeposedKey temporarily during replacement. In that case,
	// DeposedKey in the plan is always states.NotDeposed, representing that
	// the current object is being replaced with the deposed.
	DeposedKey states.DeposedKey

	// Provider is the address of the provider configuration that was used
	// to plan this change, and thus the configuration that must also be
	// used to apply it.
	ProviderAddr addrs.AbsProviderConfig

	// Change is an embedded description of the change.
	Change

	// RequiredReplace is a set of paths that caused the change action to be
	// Replace rather than Update. Always nil if the change action is not
	// Replace.
	//
	// This is retained only for UI-plan-rendering purposes and so it does not
	// currently survive a round-trip through a saved plan file.
	RequiredReplace cty.PathSet

	// Private allows a provider to stash any extra data that is opaque to
	// Terraform that relates to this change. Terraform will save this
	// byte-for-byte and return it to the provider in the apply call.
	Private []byte
}

// Encode produces a variant of the reciever that has its change values
// serialized so it can be written to a plan file. Pass the implied type of the
// corresponding resource type schema for correct operation.
func (rc *ResourceInstanceChange) Encode(ty cty.Type) (*ResourceInstanceChangeSrc, error) {
	cs, err := rc.Change.Encode(ty)
	if err != nil {
		return nil, err
	}
	return &ResourceInstanceChangeSrc{
		Addr:            rc.Addr,
		DeposedKey:      rc.DeposedKey,
		ProviderAddr:    rc.ProviderAddr,
		ChangeSrc:       *cs,
		RequiredReplace: rc.RequiredReplace,
		Private:         rc.Private,
	}, err
}

// Simplify will, where possible, produce a change with a simpler action than
// the receiever given a flag indicating whether the caller is dealing with
// a normal apply or a destroy. This flag deals with the fact that Terraform
// Core uses a specialized graph node type for destroying; only that
// specialized node should set "destroying" to true.
//
// The following table shows the simplification behavior:
//
//     Action    Destroying?   New Action
//     --------+-------------+-----------
//     Create    true          NoOp
//     Delete    false         NoOp
//     Replace   true          Delete
//     Replace   false         Create
//
// For any combination not in the above table, the Simplify just returns the
// receiver as-is.
func (rc *ResourceInstanceChange) Simplify(destroying bool) *ResourceInstanceChange {
	if destroying {
		switch rc.Action {
		case Delete:
			// We'll fall out and just return rc verbatim, then.
		case CreateThenDelete, DeleteThenCreate:
			return &ResourceInstanceChange{
				Addr:         rc.Addr,
				DeposedKey:   rc.DeposedKey,
				Private:      rc.Private,
				ProviderAddr: rc.ProviderAddr,
				Change: Change{
					Action: Delete,
					Before: rc.Before,
					After:  cty.NullVal(rc.Before.Type()),
				},
			}
		default:
			return &ResourceInstanceChange{
				Addr:         rc.Addr,
				DeposedKey:   rc.DeposedKey,
				Private:      rc.Private,
				ProviderAddr: rc.ProviderAddr,
				Change: Change{
					Action: NoOp,
					Before: rc.Before,
					After:  rc.Before,
				},
			}
		}
	} else {
		switch rc.Action {
		case Delete:
			return &ResourceInstanceChange{
				Addr:         rc.Addr,
				DeposedKey:   rc.DeposedKey,
				Private:      rc.Private,
				ProviderAddr: rc.ProviderAddr,
				Change: Change{
					Action: NoOp,
					Before: rc.Before,
					After:  rc.Before,
				},
			}
		case CreateThenDelete, DeleteThenCreate:
			return &ResourceInstanceChange{
				Addr:         rc.Addr,
				DeposedKey:   rc.DeposedKey,
				Private:      rc.Private,
				ProviderAddr: rc.ProviderAddr,
				Change: Change{
					Action: Create,
					Before: cty.NullVal(rc.After.Type()),
					After:  rc.After,
				},
			}
		}
	}

	// If we fall out here then our change is already simple enough.
	return rc
}

// OutputChange describes a change to an output value.
type OutputChange struct {
	// Addr is the absolute address of the output value that the change
	// will apply to.
	Addr addrs.AbsOutputValue

	// Change is an embedded description of the change.
	//
	// For output value changes, the type constraint for the DynamicValue
	// instances is always cty.DynamicPseudoType.
	Change

	// Sensitive, if true, indicates that either the old or new value in the
	// change is sensitive and so a rendered version of the plan in the UI
	// should elide the actual values while still indicating the action of the
	// change.
	Sensitive bool
}

// Encode produces a variant of the reciever that has its change values
// serialized so it can be written to a plan file.
func (oc *OutputChange) Encode() (*OutputChangeSrc, error) {
	cs, err := oc.Change.Encode(cty.DynamicPseudoType)
	if err != nil {
		return nil, err
	}
	return &OutputChangeSrc{
		Addr:      oc.Addr,
		ChangeSrc: *cs,
		Sensitive: oc.Sensitive,
	}, err
}

// Change describes a single change with a given action.
type Change struct {
	// Action defines what kind of change is being made.
	Action Action

	// Interpretation of Before and After depend on Action:
	//
	//     NoOp     Before and After are the same, unchanged value
	//     Create   Before is nil, and After is the expected value after create.
	//     Read     Before is any prior value (nil if no prior), and After is the
	//              value that was or will be read.
	//     Update   Before is the value prior to update, and After is the expected
	//              value after update.
	//     Replace  As with Update.
	//     Delete   Before is the value prior to delete, and After is always nil.
	//
	// Unknown values may appear anywhere within the Before and After values,
	// either as the values themselves or as nested elements within known
	// collections/structures.
	Before, After cty.Value
}

// Encode produces a variant of the reciever that has its change values
// serialized so it can be written to a plan file. Pass the type constraint
// that the values are expected to conform to; to properly decode the values
// later an identical type constraint must be provided at that time.
//
// Where a Change is embedded in some other struct, it's generally better
// to call the corresponding Encode method of that struct rather than working
// directly with its embedded Change.
func (c *Change) Encode(ty cty.Type) (*ChangeSrc, error) {
	beforeDV, err := NewDynamicValue(c.Before, ty)
	if err != nil {
		return nil, err
	}
	afterDV, err := NewDynamicValue(c.After, ty)
	if err != nil {
		return nil, err
	}

	return &ChangeSrc{
		Action: c.Action,
		Before: beforeDV,
		After:  afterDV,
	}, nil
}
