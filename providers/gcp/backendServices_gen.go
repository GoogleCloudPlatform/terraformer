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

var backendServicesAllowEmptyValues = []string{""}

var backendServicesAdditionalFields = map[string]string{}

type BackendServicesGenerator struct {
	GCPService
}

// Run on backendServicesList and create for each TerraformResource
func (g BackendServicesGenerator) createResources(ctx context.Context, backendServicesList *compute.BackendServicesListCall) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := backendServicesList.Pages(ctx, func(page *compute.BackendServiceList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewResource(
				obj.Name,
				obj.Name,
				"google_compute_backend_service",
				"google",
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"].(string),
				},
				backendServicesAllowEmptyValues,
				backendServicesAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each backendServices create 1 TerraformResource
// Need backendServices name as ID for terraform resource
func (g *BackendServicesGenerator) InitResources() error {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	backendServicesList := computeService.BackendServices.List(g.GetArgs()["project"].(string))
	g.Resources = g.createResources(ctx, backendServicesList)

	g.PopulateIgnoreKeys()
	return nil

}
