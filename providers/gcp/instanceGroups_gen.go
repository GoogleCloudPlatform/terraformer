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

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

var instanceGroupsAllowEmptyValues = []string{""}

var instanceGroupsAdditionalFields = map[string]string{}

type InstanceGroupsGenerator struct {
	GCPService
}

// Run on instanceGroupsList and create for each TerraformResource
func (g InstanceGroupsGenerator) createResources(ctx context.Context, instanceGroupsList *compute.InstanceGroupsListCall) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := instanceGroupsList.Pages(ctx, func(page *compute.InstanceGroupList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewResource(
				g.GetArgs()["zone"]+"/"+obj.Name,
				obj.Name,
				"google_compute_instance_group",
				"google",
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"],
					"region":  g.GetArgs()["region"],
					"zone":    g.GetArgs()["zone"],
				},
				instanceGroupsAllowEmptyValues,
				instanceGroupsAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each instanceGroups create 1 TerraformResource
// Need instanceGroups name as ID for terraform resource
func (g *InstanceGroupsGenerator) InitResources() error {
	ctx := context.Background()
	c, err := google.DefaultClient(ctx, compute.ComputeReadonlyScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	instanceGroupsList := computeService.InstanceGroups.List(g.GetArgs()["project"], g.GetArgs()["zone"])

	g.Resources = g.createResources(ctx, instanceGroupsList)
	g.PopulateIgnoreKeys()
	return nil

}
