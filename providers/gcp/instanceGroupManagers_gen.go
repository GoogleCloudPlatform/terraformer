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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"google.golang.org/api/compute/v1"
)

var instanceGroupManagersAllowEmptyValues = []string{"^version.[0-9].name", "^auto_healing_policies.[0-9].health_check"}

var instanceGroupManagersAdditionalFields = map[string]interface{}{}

type InstanceGroupManagersGenerator struct {
	GCPService
}

// Run on instanceGroupManagersList and create for each TerraformResource
func (g InstanceGroupManagersGenerator) createResources(ctx context.Context, instanceGroupManagersList *compute.InstanceGroupManagersListCall, zone string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	if err := instanceGroupManagersList.Pages(ctx, func(page *compute.InstanceGroupManagerList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraformutils.NewResource(
				obj.Name,
				obj.Name,
				"google_compute_instance_group_manager",
				g.ProviderName,
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"].(string),

					"zone": zone,
				},
				instanceGroupManagersAllowEmptyValues,
				instanceGroupManagersAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Println(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each instanceGroupManagers create 1 TerraformResource
// Need instanceGroupManagers name as ID for terraform resource
func (g *InstanceGroupManagersGenerator) InitResources() error {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		return err
	}

	for _, zoneLink := range g.GetArgs()["region"].(compute.Region).Zones {
		t := strings.Split(zoneLink, "/")
		zone := t[len(t)-1]
		instanceGroupManagersList := computeService.InstanceGroupManagers.List(g.GetArgs()["project"].(string), zone)
		g.Resources = append(g.Resources, g.createResources(ctx, instanceGroupManagersList, zone)...)
	}

	return nil

}
