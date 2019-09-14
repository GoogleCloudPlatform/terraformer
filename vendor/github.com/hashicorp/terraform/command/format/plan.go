package format

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/mitchellh/colorstring"

	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/plans"
	"github.com/hashicorp/terraform/states"
	"github.com/hashicorp/terraform/terraform"
)

// Plan is a representation of a plan optimized for display to
// an end-user, as opposed to terraform.Plan which is for internal use.
//
// DisplayPlan excludes implementation details that may otherwise appear
// in the main plan, such as destroy actions on data sources (which are
// there only to clean up the state).
type Plan struct {
	Resources []*InstanceDiff
}

// InstanceDiff is a representation of an instance diff optimized
// for display, in conjunction with DisplayPlan.
type InstanceDiff struct {
	Addr   *terraform.ResourceAddress
	Action plans.Action

	// Attributes describes changes to the attributes of the instance.
	//
	// For destroy diffs this is always nil.
	Attributes []*AttributeDiff

	Tainted bool
	Deposed bool
}

// AttributeDiff is a representation of an attribute diff optimized
// for display, in conjunction with DisplayInstanceDiff.
type AttributeDiff struct {
	// Path is a dot-delimited traversal through possibly many levels of list and map structure,
	// intended for display purposes only.
	Path string

	Action plans.Action

	OldValue string
	NewValue string

	NewComputed bool
	Sensitive   bool
	ForcesNew   bool
}

// PlanStats gives summary counts for a Plan.
type PlanStats struct {
	ToAdd, ToChange, ToDestroy int
}

// NewPlan produces a display-oriented Plan from a terraform.Plan.
func NewPlan(changes *plans.Changes) *Plan {
	log.Printf("[TRACE] NewPlan for %#v", changes)
	ret := &Plan{}
	if changes == nil {
		// Nothing to do!
		return ret
	}

	for _, rc := range changes.Resources {
		addr := rc.Addr
		log.Printf("[TRACE] NewPlan found %s (%s)", addr, rc.Action)
		dataSource := addr.Resource.Resource.Mode == addrs.DataResourceMode

		// We create "delete" actions for data resources so we can clean
		// up their entries in state, but this is an implementation detail
		// that users shouldn't see.
		if dataSource && rc.Action == plans.Delete {
			continue
		}

		if rc.Action == plans.NoOp {
			continue
		}

		// For now we'll shim this to work with our old types.
		// TODO: Update for the new plan types, ideally also switching over to
		// a structural diff renderer instead of a flat renderer.
		did := &InstanceDiff{
			Addr:   terraform.NewLegacyResourceInstanceAddress(addr),
			Action: rc.Action,
		}

		if rc.DeposedKey != states.NotDeposed {
			did.Deposed = true
		}

		// Since this is just a temporary stub implementation on the way
		// to us replacing this with the structural diff renderer, we currently
		// don't include any attributes here.
		// FIXME: Implement the structural diff renderer to replace this
		// codepath altogether.

		ret.Resources = append(ret.Resources, did)
	}

	// Sort the instance diffs by their addresses for display.
	sort.Slice(ret.Resources, func(i, j int) bool {
		iAddr := ret.Resources[i].Addr
		jAddr := ret.Resources[j].Addr
		return iAddr.Less(jAddr)
	})

	return ret
}

// Format produces and returns a text representation of the receiving plan
// intended for display in a terminal.
//
// If color is not nil, it is used to colorize the output.
func (p *Plan) Format(color *colorstring.Colorize) string {
	if p.Empty() {
		return "This plan does nothing."
	}

	if color == nil {
		color = &colorstring.Colorize{
			Colors: colorstring.DefaultColors,
			Reset:  false,
		}
	}

	// Find the longest path length of all the paths that are changing,
	// so we can align them all.
	keyLen := 0
	for _, r := range p.Resources {
		for _, attr := range r.Attributes {
			key := attr.Path

			if len(key) > keyLen {
				keyLen = len(key)
			}
		}
	}

	buf := new(bytes.Buffer)
	for _, r := range p.Resources {
		formatPlanInstanceDiff(buf, r, keyLen, color)
	}

	return strings.TrimSpace(buf.String())
}

// Stats returns statistics about the plan
func (p *Plan) Stats() PlanStats {
	var ret PlanStats
	for _, r := range p.Resources {
		switch r.Action {
		case plans.Create:
			ret.ToAdd++
		case plans.Update:
			ret.ToChange++
		case plans.DeleteThenCreate, plans.CreateThenDelete:
			ret.ToAdd++
			ret.ToDestroy++
		case plans.Delete:
			ret.ToDestroy++
		}
	}
	return ret
}

// ActionCounts returns the number of diffs for each action type
func (p *Plan) ActionCounts() map[plans.Action]int {
	ret := map[plans.Action]int{}
	for _, r := range p.Resources {
		ret[r.Action]++
	}
	return ret
}

// Empty returns true if there is at least one resource diff in the receiving plan.
func (p *Plan) Empty() bool {
	return len(p.Resources) == 0
}

// DiffActionSymbol returns a string that, once passed through a
// colorstring.Colorize, will produce a result that can be written
// to a terminal to produce a symbol made of three printable
// characters, possibly interspersed with VT100 color codes.
func DiffActionSymbol(action plans.Action) string {
	switch action {
	case plans.DeleteThenCreate:
		return "[red]-[reset]/[green]+[reset]"
	case plans.CreateThenDelete:
		return "[green]+[reset]/[red]-[reset]"
	case plans.Create:
		return "  [green]+[reset]"
	case plans.Delete:
		return "  [red]-[reset]"
	case plans.Read:
		return " [cyan]<=[reset]"
	case plans.Update:
		return "  [yellow]~[reset]"
	default:
		return "  ?"
	}
}

// formatPlanInstanceDiff writes the text representation of the given instance diff
// to the given buffer, using the given colorizer.
func formatPlanInstanceDiff(buf *bytes.Buffer, r *InstanceDiff, keyLen int, colorizer *colorstring.Colorize) {
	addrStr := r.Addr.String()

	// Determine the color for the text (green for adding, yellow
	// for change, red for delete), and symbol, and output the
	// resource header.
	color := "yellow"
	symbol := DiffActionSymbol(r.Action)
	oldValues := true
	switch r.Action {
	case plans.DeleteThenCreate, plans.CreateThenDelete:
		color = "yellow"
	case plans.Create:
		color = "green"
		oldValues = false
	case plans.Delete:
		color = "red"
	case plans.Read:
		color = "cyan"
		oldValues = false
	}

	var extraStr string
	if r.Tainted {
		extraStr = extraStr + " (tainted)"
	}
	if r.Deposed {
		extraStr = extraStr + " (deposed)"
	}
	if r.Action.IsReplace() {
		extraStr = extraStr + colorizer.Color(" [red][bold](new resource required)")
	}

	buf.WriteString(
		colorizer.Color(fmt.Sprintf(
			"[%s]%s [%s]%s%s\n",
			color, symbol, color, addrStr, extraStr,
		)),
	)

	for _, attr := range r.Attributes {

		v := attr.NewValue
		var dispV string
		switch {
		case v == "" && attr.NewComputed:
			dispV = "<computed>"
		case attr.Sensitive:
			dispV = "<sensitive>"
		default:
			dispV = fmt.Sprintf("%q", v)
		}

		updateMsg := ""
		switch {
		case attr.ForcesNew && r.Action.IsReplace():
			updateMsg = colorizer.Color(" [red](forces new resource)")
		case attr.Sensitive && oldValues:
			updateMsg = colorizer.Color(" [yellow](attribute changed)")
		}

		if oldValues {
			u := attr.OldValue
			var dispU string
			switch {
			case attr.Sensitive:
				dispU = "<sensitive>"
			default:
				dispU = fmt.Sprintf("%q", u)
			}
			buf.WriteString(fmt.Sprintf(
				"      %s:%s %s => %s%s\n",
				attr.Path,
				strings.Repeat(" ", keyLen-len(attr.Path)),
				dispU, dispV,
				updateMsg,
			))
		} else {
			buf.WriteString(fmt.Sprintf(
				"      %s:%s %s%s\n",
				attr.Path,
				strings.Repeat(" ", keyLen-len(attr.Path)),
				dispV,
				updateMsg,
			))
		}
	}

	// Write the reset color so we don't bleed color into later text
	buf.WriteString(colorizer.Color("[reset]\n"))
}
