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
package computeTerrforming

import (
	"context"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"waze/terraformer/gcp_terraforming/gcp_generator"
	"waze/terraformer/terraform_utils"
)

var regionBackendServicesAllowEmptyValues = map[string]bool{}

var regionBackendServicesAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type RegionBackendServicesGenerator struct {
	gcp_generator.BasicGenerator
}

// Run on regionBackendServicesList and create for each TerraformResource
func (RegionBackendServicesGenerator) createResources(regionBackendServicesList *compute.RegionBackendServicesListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := regionBackendServicesList.Pages(ctx, func(page *compute.BackendServiceList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_region_backend_service",
				"google",
				nil,
				map[string]string{
					"name":    obj.Name,
					"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
					"region":  region,
				},
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each regionBackendServices create 1 TerraformResource
// Need regionBackendServices name as ID for terraform resource
func (g RegionBackendServicesGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	region := strings.Join(strings.Split(zone, "-")[:len(strings.Split(zone, "-"))-1], "-")
	project := os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	regionBackendServicesList := computeService.RegionBackendServices.List(project, region)

	resources := g.createResources(regionBackendServicesList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, g.IgnoreKeys(resources), regionBackendServicesAllowEmptyValues, regionBackendServicesAdditionalFields)
	return resources, metadata, nil

}
