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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"google.golang.org/api/compute/v1"
)

var networkEndpointGroupsAllowEmptyValues = []string{""}

var networkEndpointGroupsAdditionalFields = map[string]string{}

type NetworkEndpointGroupsGenerator struct {
	GCPService
}

// Run on networkEndpointGroupsList and create for each TerraformResource
func (g NetworkEndpointGroupsGenerator) createResources(ctx context.Context, networkEndpointGroupsList *compute.NetworkEndpointGroupsListCall) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := networkEndpointGroupsList.Pages(ctx, func(page *compute.NetworkEndpointGroupList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewResource(
				g.GetArgs()["zone"]+"/"+obj.Name,
				obj.Name,
				"google_compute_network_endpoint_group",
				"google",
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"],
					"region":  g.GetArgs()["region"],
					"zone":    g.GetArgs()["zone"],
				},
				networkEndpointGroupsAllowEmptyValues,
				networkEndpointGroupsAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each networkEndpointGroups create 1 TerraformResource
// Need networkEndpointGroups name as ID for terraform resource
func (g *NetworkEndpointGroupsGenerator) InitResources() error {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	networkEndpointGroupsList := computeService.NetworkEndpointGroups.List(g.GetArgs()["project"], g.GetArgs()["zone"])

	g.Resources = g.createResources(ctx, networkEndpointGroupsList)
	g.PopulateIgnoreKeys()
	return nil

}
