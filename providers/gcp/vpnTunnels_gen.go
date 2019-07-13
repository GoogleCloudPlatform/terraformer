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

var vpnTunnelsAllowEmptyValues = []string{""}

var vpnTunnelsAdditionalFields = map[string]string{}

type VpnTunnelsGenerator struct {
	GCPService
}

// Run on vpnTunnelsList and create for each TerraformResource
func (g VpnTunnelsGenerator) createResources(ctx context.Context, vpnTunnelsList *compute.VpnTunnelsListCall) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := vpnTunnelsList.Pages(ctx, func(page *compute.VpnTunnelList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewResource(
				obj.Name,
				obj.Name,
				"google_compute_vpn_tunnel",
				"google",
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"].(string),
					"region":  g.GetArgs()["region"].(compute.Region).Name,
				},
				vpnTunnelsAllowEmptyValues,
				vpnTunnelsAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each vpnTunnels create 1 TerraformResource
// Need vpnTunnels name as ID for terraform resource
func (g *VpnTunnelsGenerator) InitResources() error {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	vpnTunnelsList := computeService.VpnTunnels.List(g.GetArgs()["project"].(string), g.GetArgs()["region"].(compute.Region).Name)
	g.Resources = g.createResources(ctx, vpnTunnelsList)

	g.PopulateIgnoreKeys()
	return nil

}
