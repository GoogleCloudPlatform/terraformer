package terraform

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/configs"
	"github.com/hashicorp/terraform/dag"
	"github.com/hashicorp/terraform/states"
)

// GraphNodeDestroyerCBD must be implemented by nodes that might be
// create-before-destroy destroyers, or might plan a create-before-destroy
// action.
type GraphNodeDestroyerCBD interface {
	// CreateBeforeDestroy returns true if this node represents a node
	// that is doing a CBD.
	CreateBeforeDestroy() bool

	// ModifyCreateBeforeDestroy is called when the CBD state of a node
	// is changed dynamically. This can return an error if this isn't
	// allowed.
	ModifyCreateBeforeDestroy(bool) error
}

// GraphNodeAttachDestroyer is implemented by applyable nodes that have a
// companion destroy node. This allows the creation node to look up the status
// of the destroy node and determine if it needs to depose the existing state,
// or replace it.
// If a node is not marked as create-before-destroy in the configuration, but a
// dependency forces that status, only the destroy node will be aware of that
// status.
type GraphNodeAttachDestroyer interface {
	// AttachDestroyNode takes a destroy node and saves a reference to that
	// node in the receiver, so it can later check the status of
	// CreateBeforeDestroy().
	AttachDestroyNode(n GraphNodeDestroyerCBD)
}

// ForcedCBDTransformer detects when a particular CBD-able graph node has
// dependencies with another that has create_before_destroy set that require
// it to be forced on, and forces it on.
//
// This must be used in the plan graph builder to ensure that
// create_before_destroy settings are properly propagated before constructing
// the planned changes. This requires that the plannable resource nodes
// implement GraphNodeDestroyerCBD.
type ForcedCBDTransformer struct {
}

func (t *ForcedCBDTransformer) Transform(g *Graph) error {
	for _, v := range g.Vertices() {
		dn, ok := v.(GraphNodeDestroyerCBD)
		if !ok {
			continue
		}

		if !dn.CreateBeforeDestroy() {
			// If there are no CBD decendent (dependent nodes), then we
			// do nothing here.
			if !t.hasCBDDescendent(g, v) {
				log.Printf("[TRACE] ForcedCBDTransformer: %q (%T) has no CBD descendent, so skipping", dag.VertexName(v), v)
				continue
			}

			// If this isn't naturally a CBD node, this means that an descendent is
			// and we need to auto-upgrade this node to CBD. We do this because
			// a CBD node depending on non-CBD will result in cycles. To avoid this,
			// we always attempt to upgrade it.
			log.Printf("[TRACE] ForcedCBDTransformer: forcing create_before_destroy on for %q (%T)", dag.VertexName(v), v)
			if err := dn.ModifyCreateBeforeDestroy(true); err != nil {
				return fmt.Errorf(
					"%s: must have create before destroy enabled because "+
						"a dependent resource has CBD enabled. However, when "+
						"attempting to automatically do this, an error occurred: %s",
					dag.VertexName(v), err)
			}
		} else {
			log.Printf("[TRACE] ForcedCBDTransformer: %q (%T) already has create_before_destroy set", dag.VertexName(v), v)
		}
	}
	return nil
}

// hasCBDDescendent returns true if any descendent (node that depends on this)
// has CBD set.
func (t *ForcedCBDTransformer) hasCBDDescendent(g *Graph, v dag.Vertex) bool {
	s, _ := g.Descendents(v)
	if s == nil {
		return true
	}

	for _, ov := range s.List() {
		dn, ok := ov.(GraphNodeDestroyerCBD)
		if !ok {
			continue
		}

		if dn.CreateBeforeDestroy() {
			// some descendent is CreateBeforeDestroy, so we need to follow suit
			log.Printf("[TRACE] ForcedCBDTransformer: %q has CBD descendent %q", dag.VertexName(v), dag.VertexName(ov))
			return true
		}
	}

	return false
}

// CBDEdgeTransformer modifies the edges of CBD nodes that went through
// the DestroyEdgeTransformer to have the right dependencies. There are
// two real tasks here:
//
//   1. With CBD, the destroy edge is inverted: the destroy depends on
//      the creation.
//
//   2. A_d must depend on resources that depend on A. This is to enable
//      the destroy to only happen once nodes that depend on A successfully
//      update to A. Example: adding a web server updates the load balancer
//      before deleting the old web server.
//
// This transformer requires that a previous transformer has already forced
// create_before_destroy on for nodes that are depended on by explicit CBD
// nodes. This is the logic in ForcedCBDTransformer, though in practice we
// will get here by recording the CBD-ness of each change in the plan during
// the plan walk and then forcing the nodes into the appropriate setting during
// DiffTransformer when building the apply graph.
type CBDEdgeTransformer struct {
	// Module and State are only needed to look up dependencies in
	// any way possible. Either can be nil if not availabile.
	Config *configs.Config
	State  *states.State

	// If configuration is present then Schemas is required in order to
	// obtain schema information from providers and provisioners so we can
	// properly resolve implicit dependencies.
	Schemas *Schemas

	// If the operation is a simple destroy, no transformation is done.
	Destroy bool
}

func (t *CBDEdgeTransformer) Transform(g *Graph) error {
	if t.Destroy {
		return nil
	}

	// Go through and reverse any destroy edges
	destroyMap := make(map[string][]dag.Vertex)
	for _, v := range g.Vertices() {
		dn, ok := v.(GraphNodeDestroyerCBD)
		if !ok {
			continue
		}
		dern, ok := v.(GraphNodeDestroyer)
		if !ok {
			continue
		}

		if !dn.CreateBeforeDestroy() {
			continue
		}

		// Find the resource edges
		for _, e := range g.EdgesTo(v) {
			switch de := e.(type) {
			case *DestroyEdge:
				// we need to invert the destroy edge from the create node
				log.Printf("[TRACE] CBDEdgeTransformer: inverting edge: %s => %s",
					dag.VertexName(de.Source()), dag.VertexName(de.Target()))

				// Found it! Invert.
				g.RemoveEdge(de)
				applyNode := de.Source()
				destroyNode := de.Target()
				g.Connect(&DestroyEdge{S: destroyNode, T: applyNode})
			default:
				// We cannot have any direct dependencies from creators when
				// the node is CBD without inducing a cycle.
				if _, ok := e.Source().(GraphNodeCreator); ok {
					log.Printf("[TRACE] CBDEdgeTransformer: removing non DestroyEdge to CBD destroy node: %s => %s", dag.VertexName(e.Source()), dag.VertexName(e.Target()))
					g.RemoveEdge(e)
				}
			}
		}

		// If the address has an index, we strip that. Our depMap creation
		// graph doesn't expand counts so we don't currently get _exact_
		// dependencies. One day when we limit dependencies more exactly
		// this will have to change. We have a test case covering this
		// (depNonCBDCountBoth) so it'll be caught.
		addr := dern.DestroyAddr()
		key := addr.ContainingResource().String()

		// Add this to the list of nodes that we need to fix up
		// the edges for (step 2 above in the docs).
		destroyMap[key] = append(destroyMap[key], v)
	}

	// If we have no CBD nodes, then our work here is done
	if len(destroyMap) == 0 {
		return nil
	}

	// We have CBD nodes. We now have to move on to the much more difficult
	// task of connecting dependencies of the creation side of the destroy
	// to the destruction node. The easiest way to explain this is an example:
	//
	// Given a pre-destroy dependence of: A => B
	//   And A has CBD set.
	//
	// The resulting graph should be: A => B => A_d
	//
	// They key here is that B happens before A is destroyed. This is to
	// facilitate the primary purpose for CBD: making sure that downstreams
	// are properly updated to avoid downtime before the resource is destroyed.
	depMap, err := t.depMap(g, destroyMap)
	if err != nil {
		return err
	}

	// We now have the mapping of resource addresses to the destroy
	// nodes they need to depend on. We now go through our own vertices to
	// find any matching these addresses and make the connection.
	for _, v := range g.Vertices() {
		// We're looking for creators
		rn, ok := v.(GraphNodeCreator)
		if !ok {
			continue
		}

		// Get the address
		addr := rn.CreateAddr()

		// If the address has an index, we strip that. Our depMap creation
		// graph doesn't expand counts so we don't currently get _exact_
		// dependencies. One day when we limit dependencies more exactly
		// this will have to change. We have a test case covering this
		// (depNonCBDCount) so it'll be caught.
		key := addr.ContainingResource().String()

		// If there is nothing this resource should depend on, ignore it
		dns, ok := depMap[key]
		if !ok {
			continue
		}

		// We have nodes! Make the connection
		for _, dn := range dns {
			log.Printf("[TRACE] CBDEdgeTransformer: destroy depends on dependence: %s => %s",
				dag.VertexName(dn), dag.VertexName(v))
			g.Connect(dag.BasicEdge(dn, v))
		}
	}

	return nil
}

func (t *CBDEdgeTransformer) depMap(g *Graph, destroyMap map[string][]dag.Vertex) (map[string][]dag.Vertex, error) {
	// Build the list of destroy nodes that each resource address should depend
	// on. For example, when we find B, we map the address of B to A_d in the
	// "depMap" variable below.

	// Use a nested map to remove duplicate edges.
	depMap := make(map[string]map[dag.Vertex]struct{})
	for _, v := range g.Vertices() {
		// We're looking for resources.
		rn, ok := v.(GraphNodeResource)
		if !ok {
			continue
		}

		// Get the address
		addr := rn.ResourceAddr()
		key := addr.String()

		// Get the destroy nodes that are destroying this resource.
		// If there aren't any, then we don't need to worry about
		// any connections.
		dns, ok := destroyMap[key]
		if !ok {
			continue
		}

		// Get the nodes that depend on this on. In the example above:
		// finding B in A => B. Since dependencies can span modules, walk all
		// descendents of the resource.
		des, _ := g.Descendents(v)
		for _, v := range des.List() {
			// We're looking for resources.
			rn, ok := v.(GraphNodeResource)
			if !ok {
				continue
			}

			// Keep track of the destroy nodes that this address
			// needs to depend on.
			key := rn.ResourceAddr().String()

			deps, ok := depMap[key]
			if !ok {
				deps = make(map[dag.Vertex]struct{})
			}

			for _, d := range dns {
				deps[d] = struct{}{}
			}
			depMap[key] = deps
		}
	}

	result := map[string][]dag.Vertex{}
	for k, m := range depMap {
		for v := range m {
			result[k] = append(result[k], v)
		}
	}

	return result, nil
}
