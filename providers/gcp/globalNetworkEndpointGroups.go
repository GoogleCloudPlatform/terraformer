package gcp

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/iterator"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
)

var globalNetworkEndpointGroupsAllowEmptyValues = []string{""}

var globalNetworkEndpointGroupsAdditionalFields = map[string]interface{}{}

type GlobalNetworkEndpointGroupsGenerator struct {
	GCPService
}

func (g *GlobalNetworkEndpointGroupsGenerator) InitResources() error {
	ctx := context.Background()
	computeService, err := compute.NewGlobalNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		return err
	}
	defer computeService.Close()

	req := &computepb.ListGlobalNetworkEndpointGroupsRequest{Project: g.GetArgs()["project"].(string)}

	it := computeService.List(ctx, req)

	for {
		group, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				return nil
			}
			return err
		}

		res := terraformutils.NewResource(
			group.GetName(),
			group.GetName(),
			"google_compute_global_network_endpoint_group",
			g.ProviderName,
			map[string]string{
				"name":    group.GetName(),
				"project": g.GetArgs()["project"].(string),
				"region":  group.GetRegion(),
				"zone":    group.GetZone(),
			},
			globalNetworkEndpointGroupsAllowEmptyValues,
			globalNetworkEndpointGroupsAdditionalFields,
		)
		g.Resources = append(g.Resources, res)
	}
}
