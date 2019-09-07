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

var instanceTemplatesAllowEmptyValues = []string{""}

var instanceTemplatesAdditionalFields = map[string]string{}

type InstanceTemplatesGenerator struct {
	GCPService
}

// Run on instanceTemplatesList and create for each TerraformResource
func (g InstanceTemplatesGenerator) createResources(ctx context.Context, instanceTemplatesList *compute.InstanceTemplatesListCall) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := instanceTemplatesList.Pages(ctx, func(page *compute.InstanceTemplateList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewResource(
				obj.Name,
				obj.Name,
				"google_compute_instance_template",
				"google",
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"].(string),
					"region":  g.GetArgs()["region"].(compute.Region).Name,
				},
				instanceTemplatesAllowEmptyValues,
				instanceTemplatesAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each instanceTemplates create 1 TerraformResource
// Need instanceTemplates name as ID for terraform resource
func (g *InstanceTemplatesGenerator) InitResources() error {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	instanceTemplatesList := computeService.InstanceTemplates.List(g.GetArgs()["project"].(string))
	g.Resources = g.createResources(ctx, instanceTemplatesList)

	g.PopulateIgnoreKeys()
	return nil

}
