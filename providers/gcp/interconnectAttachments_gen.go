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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"google.golang.org/api/compute/v1"
)

var interconnectAttachmentsAllowEmptyValues = []string{""}

var interconnectAttachmentsAdditionalFields = map[string]interface{}{}

type InterconnectAttachmentsGenerator struct {
	GCPService
}

// Run on interconnectAttachmentsList and create for each TerraformResource
func (g InterconnectAttachmentsGenerator) createResources(ctx context.Context, interconnectAttachmentsList *compute.InterconnectAttachmentsListCall) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	if err := interconnectAttachmentsList.Pages(ctx, func(page *compute.InterconnectAttachmentList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraformutils.NewResource(
				obj.Name,
				obj.Name,
				"google_compute_interconnect_attachment",
				g.ProviderName,
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"].(string),
					"region":  g.GetArgs()["region"].(compute.Region).Name,
				},
				interconnectAttachmentsAllowEmptyValues,
				interconnectAttachmentsAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Println(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each interconnectAttachments create 1 TerraformResource
// Need interconnectAttachments name as ID for terraform resource
func (g *InterconnectAttachmentsGenerator) InitResources() error {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		return err
	}

	interconnectAttachmentsList := computeService.InterconnectAttachments.List(g.GetArgs()["project"].(string), g.GetArgs()["region"].(compute.Region).Name)
	g.Resources = g.createResources(ctx, interconnectAttachmentsList)

	return nil

}
