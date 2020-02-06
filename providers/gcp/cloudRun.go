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

package gcp

import (
	"context"
	"fmt"
	compute "google.golang.org/api/compute/v1"
	"google.golang.org/api/run/v1"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

type CloudRunGenerator struct {
	GCPService
}

// Run on CloudRunServicesList and create for each TerraformResource
func (g CloudRunGenerator) createResources(servicesList *run.ListServicesResponse) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, item := range servicesList.Items {
		resources = append(resources, terraform_utils.NewResource(
			fmt.Sprintf(
				"locations/%s/namespaces/%s/services/%s",
				g.GetArgs()["region"].(compute.Region).Name,
				g.GetArgs()["project"].(string),
				item.Metadata.Name),
			g.GetArgs()["region"].(compute.Region).Name+"_"+item.Metadata.Name,
			"google_cloud_run_service",
			"google",
			map[string]string{
				"name": item.Metadata.Name,
				"project": g.GetArgs()["project"].(string),
				"region": g.GetArgs()["region"].(compute.Region).Name,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each CloudRun service create 1 TerraformResource
func (g *CloudRunGenerator) InitResources() error {
	ctx := context.Background()
	service, err := run.NewService(ctx)
	if err != nil {
		return err
	}

	servicesList, err := service.Projects.Locations.Services.List("projects/" + g.GetArgs()["project"].(string) + "/locations/" + g.GetArgs()["region"].(compute.Region).Name).Do()
	if err != nil {
		return err
	}

	g.Resources = g.createResources(servicesList)
	return nil

}
