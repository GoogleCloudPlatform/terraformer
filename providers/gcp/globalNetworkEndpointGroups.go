// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.
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

// Generate TerraformResources from GCP API,
// from each networkEndpointGroups create 1 TerraformResource
// Need networkEndpointGroups name as ID for terraform resource
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
