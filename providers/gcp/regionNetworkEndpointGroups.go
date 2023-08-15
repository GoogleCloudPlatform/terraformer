package gcp

import (
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/iterator"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"

	comp "google.golang.org/api/compute/v1"
)

var regionNetworkEndpointGroupsAllowEmptyValues = []string{""}

var regionNetworkEndpointGroupsAdditionalFields = map[string]interface{}{}

type RegionNetworkEndpointGroupsGenerator struct {
	GCPService
}

func (g *RegionNetworkEndpointGroupsGenerator) InitResources() error {
	ctx := context.Background()
	computeService, err := compute.NewRegionNetworkEndpointGroupsRESTClient(ctx)
	if err != nil {
		return err
	}
	defer computeService.Close()

	req := &computepb.ListRegionNetworkEndpointGroupsRequest{Project: g.GetArgs()["project"].(string), Region: g.GetArgs()["region"].(comp.Region).Name}

	it := computeService.List(ctx, req)

	for {
		group, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				return nil
			}
			return err
		}

		zoneparts := strings.Split(group.GetZone(), "/")
		zone := zoneparts[len(zoneparts)-1]

		regionparts := strings.Split(group.GetRegion(), "/")
		region := regionparts[len(regionparts)-1]

		res := terraformutils.NewResource(
			group.GetName(),
			group.GetName(),
			"google_compute_region_network_endpoint_group",
			g.ProviderName,
			map[string]string{
				"name":    group.GetName(),
				"project": g.GetArgs()["project"].(string),
				"region":  region,
				"zone":    zone,
			},
			regionNetworkEndpointGroupsAllowEmptyValues,
			regionNetworkEndpointGroupsAdditionalFields,
		)
		g.Resources = append(g.Resources, res)
	}
}
