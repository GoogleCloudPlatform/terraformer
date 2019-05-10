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

var targetPoolsAllowEmptyValues = []string{""}

var targetPoolsAdditionalFields = map[string]string{}

type TargetPoolsGenerator struct {
	GCPService
}

// Run on targetPoolsList and create for each TerraformResource
func (g TargetPoolsGenerator) createResources(ctx context.Context, targetPoolsList *compute.TargetPoolsListCall) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := targetPoolsList.Pages(ctx, func(page *compute.TargetPoolList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewResource(
				obj.Name,
				obj.Name,
				"google_compute_target_pool",
				"google",
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"],
					"region":  g.GetArgs()["region"],
				},
				targetPoolsAllowEmptyValues,
				targetPoolsAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each targetPools create 1 TerraformResource
// Need targetPools name as ID for terraform resource
func (g *TargetPoolsGenerator) InitResources() error {
	ctx := context.Background()
	c, err := google.DefaultClient(ctx, compute.ComputeReadonlyScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	targetPoolsList := computeService.TargetPools.List(g.GetArgs()["project"], g.GetArgs()["region"])

	g.Resources = g.createResources(ctx, targetPoolsList)
	g.PopulateIgnoreKeys()
	return nil

}
