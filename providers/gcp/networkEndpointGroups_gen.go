package gcp

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/iterator"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
)

var networkEndpointGroupsAllowEmptyValues = []string{""}

var networkEndpointGroupsAdditionalFields = map[string]interface{}{}

type NetworkEndpointGroupsGenerator struct {
	GCPService
}

func (g *NetworkEndpointGroupsGenerator) InitResources() error {
	ctx := context.Background()
	computeService, err := compute.NewNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		return err
	}
	defer computeService.Close()

	req := &computepb.AggregatedListNetworkEndpointGroupsRequest{Project: g.GetArgs()["project"].(string)}

	it := computeService.AggregatedList(ctx, req)

	for {
		pair, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				return nil
			}
			return err
		}

		groups := pair.Value.GetNetworkEndpointGroups()

		for i := 0; i < len(groups); i++ {
			group := groups[i]
			zone := group.GetZone()
			res := terraformutils.NewResource(
				zone+"/"+group.GetName(),
				zone+"/"+group.GetName(),
				"google_compute_network_endpoint_group",
				g.ProviderName,
				map[string]string{
					"name":    group.GetName(),
					"project": g.GetArgs()["project"].(string),
					"region":  group.GetRegion(),
					"zone":    zone,
				},
				networkEndpointGroupsAllowEmptyValues,
				networkEndpointGroupsAdditionalFields,
			)
			g.Resources = append(g.Resources, res)
		}
	}
}
