// Copyright 2019 The Terraformer Authors.
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

package cloudflare

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	cf "github.com/cloudflare/cloudflare-go"
)

type AccessGenerator struct {
	CloudflareService
}

func (g *AccessGenerator) createAccessApplications(api *cf.API, zoneID string) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}
	accessApplications, _, err := api.AccessApplications(zoneID, cf.PaginationOptions{})
	if err != nil {
		return []terraformutils.Resource{}, err
	}

	for _, app := range accessApplications {
		resources = append(resources, terraformutils.NewResource(
			app.ID,
			fmt.Sprintf("%s_%s", app.Name, app.ID),
			"cloudflare_access_application",
			"cloudflare",
			map[string]string{
				"zone_id": zoneID,
				"name":    app.Name,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources, nil
}

func (g *AccessGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	zones, err := api.ListZones()
	if err != nil {
		return err
	}

	for _, zone := range zones {
		tmpRes, err := g.createAccessApplications(api, zone.ID)
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, tmpRes...)
	}

	return nil
}
