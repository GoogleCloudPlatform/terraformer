package terraform

import (
	"log"

	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/states"

	"github.com/hashicorp/terraform/configs"
	"github.com/hashicorp/terraform/dag"
)

// GraphNodeDestroyer must be implemented by nodes that destroy resources.
type GraphNodeDestroyer interface {
	dag.Vertex

	// DestroyAddr is the address of the resource that is being
	// destroyed by this node. If this returns nil, then this node
	// is not destroying anything.
	DestroyAddr() *addrs.AbsResourceInstance
}

// GraphNodeCreator must be implemented by nodes that create OR update resources.
type GraphNodeCreator interface {
	// CreateAddr is the address of the resource being created or updated
	CreateAddr() *addrs.AbsResourceInstance
}

// DestroyEdgeTransformer is a GraphTransformer that creates the proper
// references for destroy resources. Destroy resources are more complex
// in that they must be depend on the destruction of resources that
// in turn depend on the CREATION of the node being destroy.
//
// That is complicated. Visually:
//
//   B_d -> A_d -> A -> B
//
// Notice that A destroy depends on B destroy, while B create depends on
// A create. They're inverted. This must be done for example because often
// dependent resources will block parent resources from deleting. Concrete
// example: VPC with subnets, the VPC can't be deleted while there are
// still subnets.
type DestroyEdgeTransformer struct {
	// These are needed to properly build the graph of dependencies
	// to determine what a destroy node depends on. Any of these can be nil.
	Config *configs.Config
	State  *states.State

	// If configuration is present then Schemas is required in order to
	// obtain schema information from providers and provisioners in order
	// to properly resolve implicit dependencies.
	Schemas *Schemas
}

func (t *DestroyEdgeTransformer) Transform(g *Graph) error {
	// Build a map of what is being destroyed (by address string) to
	// the list of destroyers.
	destroyers := make(map[string][]GraphNodeDestroyer)
	destroyerAddrs := make(map[string]addrs.AbsResourceInstance)

	// Record the creators, which will need to depend on the destroyers if they
	// are only being updated.
	creators := make(map[string]GraphNodeCreator)

	// destroyersByResource records each destroyer by the AbsResourceAddress.
	// We use this because dependencies are only referenced as resources, but we
	// will want to connect all the individual instances for correct ordering.
	destroyersByResource := make(map[string][]GraphNodeDestroyer)
	for _, v := range g.Vertices() {
		switch n := v.(type) {
		case GraphNodeDestroyer:
			addrP := n.DestroyAddr()
			if addrP == nil {
				log.Printf("[WARN] DestroyEdgeTransformer: %q (%T) has no destroy address", dag.VertexName(n), v)
				continue
			}
			addr := *addrP

			key := addr.String()
			log.Printf("[TRACE] DestroyEdgeTransformer: %q (%T) destroys %s", dag.VertexName(n), v, key)
			destroyers[key] = append(destroyers[key], n)
			destroyerAddrs[key] = addr

			resAddr := addr.Resource.Resource.Absolute(addr.Module).String()
			destroyersByResource[resAddr] = append(destroyersByResource[resAddr], n)
		case GraphNodeCreator:
			addr := n.CreateAddr()
			creators[addr.String()] = n
		}
	}

	// If we aren't destroying anything, there will be no edges to make
	// so just exit early and avoid future work.
	if len(destroyers) == 0 {
		return nil
	}

	// Connect destroy despendencies as stored in the state
	for _, ds := range destroyers {
		for _, des := range ds {
			ri, ok := des.(GraphNodeResourceInstance)
			if !ok {
				continue
			}

			for _, resAddr := range ri.StateDependencies() {
				for _, desDep := range destroyersByResource[resAddr.String()] {
					log.Printf("[TRACE] DestroyEdgeTransformer: %s has stored dependency of %s\n", dag.VertexName(desDep), dag.VertexName(des))
					g.Connect(dag.BasicEdge(desDep, des))

				}
			}
		}
	}

	// connect creators to any destroyers on which they may depend
	for _, c := range creators {
		ri, ok := c.(GraphNodeResourceInstance)
		if !ok {
			continue
		}

		for _, resAddr := range ri.StateDependencies() {
			for _, desDep := range destroyersByResource[resAddr.String()] {
				log.Printf("[TRACE] DestroyEdgeTransformer: %s has stored dependency of %s\n", dag.VertexName(c), dag.VertexName(desDep))
				g.Connect(dag.BasicEdge(c, desDep))

			}
		}
	}

	// Go through and connect creators to destroyers. Going along with
	// our example, this makes: A_d => A
	for _, v := range g.Vertices() {
		cn, ok := v.(GraphNodeCreator)
		if !ok {
			continue
		}

		addr := cn.CreateAddr()
		if addr == nil {
			continue
		}

		for _, d := range destroyers[addr.String()] {
			// For illustrating our example
			a_d := d.(dag.Vertex)
			a := v

			log.Printf(
				"[TRACE] DestroyEdgeTransformer: connecting creator %q with destroyer %q",
				dag.VertexName(a), dag.VertexName(a_d))

			g.Connect(&DestroyEdge{S: a, T: a_d})

			// Attach the destroy node to the creator
			// There really shouldn't be more than one destroyer, but even if
			// there are, any of them will represent the correct
			// CreateBeforeDestroy status.
			if n, ok := cn.(GraphNodeAttachDestroyer); ok {
				if d, ok := d.(GraphNodeDestroyerCBD); ok {
					n.AttachDestroyNode(d)
				}
			}
		}
	}

	// This is strange but is the easiest way to get the dependencies
	// of a node that is being destroyed. We use another graph to make sure
	// the resource is in the graph and ask for references. We have to do this
	// because the node that is being destroyed may NOT be in the graph.
	//
	// Example: resource A is force new, then destroy A AND create A are
	// in the graph. BUT if resource A is just pure destroy, then only
	// destroy A is in the graph, and create A is not.
	providerFn := func(a *NodeAbstractProvider) dag.Vertex {
		return &NodeApplyableProvider{NodeAbstractProvider: a}
	}
	steps := []GraphTransformer{
		// Add the local values
		&LocalTransformer{Config: t.Config},

		// Add outputs and metadata
		&OutputTransformer{Config: t.Config},
		&AttachResourceConfigTransformer{Config: t.Config},
		&AttachStateTransformer{State: t.State},

		// Add all the variables. We can depend on resources through
		// variables due to module parameters, and we need to properly
		// determine that.
		&RootVariableTransformer{Config: t.Config},
		&ModuleVariableTransformer{Config: t.Config},

		TransformProviders(nil, providerFn, t.Config),

		// Must attach schemas before ReferenceTransformer so that we can
		// analyze the configuration to find references.
		&AttachSchemaTransformer{Schemas: t.Schemas},

		&ReferenceTransformer{},
	}

	// Go through all the nodes being destroyed and create a graph.
	// The resulting graph is only of things being CREATED. For example,
	// following our example, the resulting graph would be:
	//
	//   A, B (with no edges)
	//
	var tempG Graph
	var tempDestroyed []dag.Vertex
	for d := range destroyers {
		// d is the string key for the resource being destroyed. We actually
		// want the address value, which we stashed earlier.
		addr := destroyerAddrs[d]

		// This part is a little bit weird but is the best way to
		// find the dependencies we need to: build a graph and use the
		// attach config and state transformers then ask for references.
		abstract := NewNodeAbstractResourceInstance(addr)
		tempG.Add(abstract)
		tempDestroyed = append(tempDestroyed, abstract)

		// We also add the destroy version here since the destroy can
		// depend on things that the creation doesn't (destroy provisioners).
		destroy := &NodeDestroyResourceInstance{NodeAbstractResourceInstance: abstract}
		tempG.Add(destroy)
		tempDestroyed = append(tempDestroyed, destroy)
	}

	// Run the graph transforms so we have the information we need to
	// build references.
	log.Printf("[TRACE] DestroyEdgeTransformer: constructing temporary graph for analysis of references, starting from:\n%s", tempG.StringWithNodeTypes())
	for _, s := range steps {
		log.Printf("[TRACE] DestroyEdgeTransformer: running %T on temporary graph", s)
		if err := s.Transform(&tempG); err != nil {
			log.Printf("[TRACE] DestroyEdgeTransformer: %T failed: %s", s, err)
			return err
		}
	}
	log.Printf("[TRACE] DestroyEdgeTransformer: temporary reference graph:\n%s", tempG.String())

	// Go through all the nodes in the graph and determine what they
	// depend on.
	for _, v := range tempDestroyed {
		// Find all ancestors of this to determine the edges we'll depend on
		vs, err := tempG.Ancestors(v)
		if err != nil {
			return err
		}

		refs := make([]dag.Vertex, 0, vs.Len())
		for _, raw := range vs.List() {
			refs = append(refs, raw.(dag.Vertex))
		}

		refNames := make([]string, len(refs))
		for i, ref := range refs {
			refNames[i] = dag.VertexName(ref)
		}
		log.Printf(
			"[TRACE] DestroyEdgeTransformer: creation node %q references %s",
			dag.VertexName(v), refNames)

		// If we have no references, then we won't need to do anything
		if len(refs) == 0 {
			continue
		}

		// Get the destroy node for this. In the example of our struct,
		// we are currently at B and we're looking for B_d.
		rn, ok := v.(GraphNodeResourceInstance)
		if !ok {
			log.Printf("[TRACE] DestroyEdgeTransformer: skipping %s, since it's not a resource", dag.VertexName(v))
			continue
		}

		addr := rn.ResourceInstanceAddr()
		dns := destroyers[addr.String()]

		// We have dependencies, check if any are being destroyed
		// to build the list of things that we must depend on!
		//
		// In the example of the struct, if we have:
		//
		//   B_d => A_d => A => B
		//
		// Then at this point in the algorithm we started with B_d,
		// we built B (to get dependencies), and we found A. We're now looking
		// to see if A_d exists.
		var depDestroyers []dag.Vertex
		for _, v := range refs {
			rn, ok := v.(GraphNodeResourceInstance)
			if !ok {
				continue
			}

			addr := rn.ResourceInstanceAddr()
			key := addr.String()
			if ds, ok := destroyers[key]; ok {
				for _, d := range ds {
					depDestroyers = append(depDestroyers, d.(dag.Vertex))
					log.Printf(
						"[TRACE] DestroyEdgeTransformer: destruction of %q depends on %s",
						key, dag.VertexName(d))
				}
			}
		}

		// Go through and make the connections. Use the variable
		// names "a_d" and "b_d" to reference our example.
		for _, a_d := range dns {
			for _, b_d := range depDestroyers {
				if b_d != a_d {
					log.Printf("[TRACE] DestroyEdgeTransformer: %q depends on %q", dag.VertexName(b_d), dag.VertexName(a_d))
					g.Connect(dag.BasicEdge(b_d, a_d))
				}
			}
		}
	}

	return t.pruneResources(g)
}

// If there are only destroy instances for a particular resource, there's no
// reason for the resource node to prepare the state. Remove Resource nodes so
// that they don't fail by trying to evaluate a resource that is only being
// destroyed along with its dependencies.
func (t *DestroyEdgeTransformer) pruneResources(g *Graph) error {
	for _, v := range g.Vertices() {
		n, ok := v.(*NodeApplyableResource)
		if !ok {
			continue
		}

		// if there are only destroy dependencies, we don't need this node
		des, err := g.Descendents(n)
		if err != nil {
			return err
		}

		descendents := des.List()
		nonDestroyInstanceFound := false
		for _, v := range descendents {
			if _, ok := v.(*NodeApplyableResourceInstance); ok {
				nonDestroyInstanceFound = true
				break
			}
		}

		if nonDestroyInstanceFound {
			continue
		}

		// connect all the through-edges, then delete the node
		for _, d := range g.DownEdges(n).List() {
			for _, u := range g.UpEdges(n).List() {
				g.Connect(dag.BasicEdge(u, d))
			}
		}
		log.Printf("DestroyEdgeTransformer: pruning unused resource node %s", dag.VertexName(n))
		g.Remove(n)
	}
	return nil
}
