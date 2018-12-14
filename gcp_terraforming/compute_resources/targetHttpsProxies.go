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

var targetHttpsProxiesAllowEmptyValues = map[string]bool{}

var targetHttpsProxiesAdditionalFields = map[string]string{
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type TargetHttpsProxiesGenerator struct {
	gcp_generator.BasicGenerator
}

// Run on targetHttpsProxiesList and create for each TerraformResource
func (TargetHttpsProxiesGenerator) createResources(targetHttpsProxiesList *compute.TargetHttpsProxiesListCall, ctx context.Context, region, zone string) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	if err := targetHttpsProxiesList.Pages(ctx, func(page *compute.TargetHttpsProxyList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewTerraformResource(
				obj.Name,
				obj.Name,
				"google_compute_target_https_proxy",
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
// from each targetHttpsProxies create 1 TerraformResource
// Need targetHttpsProxies name as ID for terraform resource
func (g TargetHttpsProxiesGenerator) Generate(zone string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
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

	targetHttpsProxiesList := computeService.TargetHttpsProxies.List(project)

	resources := g.createResources(targetHttpsProxiesList, ctx, region, zone)
	metadata := terraform_utils.NewResourcesMetaData(resources, g.IgnoreKeys(resources), targetHttpsProxiesAllowEmptyValues, targetHttpsProxiesAdditionalFields)
	return resources, metadata, nil

}
