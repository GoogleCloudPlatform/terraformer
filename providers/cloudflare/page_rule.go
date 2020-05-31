// Copyright 2020 The Terraformer Authors.
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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	cf "github.com/cloudflare/cloudflare-go"
)

type PageRulesGenerator struct {
	CloudflareService
}

func (g *PageRulesGenerator) createPageRules(api *cf.API, zoneID string) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	pageRules, err := api.ListPageRules(zoneID)
	if err != nil {
		return resources, err
	}

	for _, pageRule := range pageRules {
		resources = append(resources, terraformutils.NewResource(
			pageRule.ID,
			pageRule.ID,
			"cloudflare_page_rule",
			"cloudflare",
			map[string]string{
				"zone_id": zoneID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources, nil
}

func (g *PageRulesGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	zones, err := api.ListZones()
	if err != nil {
		return err
	}

	for _, zone := range zones {
		resources, err := g.createPageRules(api, zone.ID)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, resources...)
	}

	return nil
}
