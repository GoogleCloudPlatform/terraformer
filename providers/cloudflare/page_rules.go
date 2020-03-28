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

//cloudflare_access_rule
//cloudflare_rate_limit

package cloudflare

import (
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	cf "github.com/cloudflare/cloudflare-go"
)

type PageRulesGenerator struct {
	CloudflareService
}

func (*PageRulesGenerator) createPageRuleResources(api *cf.API, zoneID, zoneName string) ([]terraform_utils.Resource, error) {
	resources := []terraform_utils.Resource{}

	pageRules, err := api.ListPageRules(zoneID)
	if err != nil {
		log.Println(err)
		return resources, err
	}
	for _, pageRule := range pageRules {
		resources = append(resources, terraform_utils.NewResource(
			pageRule.ID,
			fmt.Sprintf("%s_%s", zoneName, pageRule.ID),
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
		panic(err)
		return err
	}

	zones, err := api.ListZones()
	if err != nil {
		panic(err)
		return err
	}

	funcs := []func(*cf.API, string, string) ([]terraform_utils.Resource, error){
		g.createPageRuleResources,
	}

	for _, zone := range zones {
		for _, f := range funcs {
			// Getting all firewall filters
			tmpRes, err := f(api, zone.ID, zone.Name)
			if err != nil {
				panic(err)
				return err
			}
			g.Resources = append(g.Resources, tmpRes...)
		}
	}

	return nil
}

func (g *PageRulesGenerator) PostConvertHook() error {
	for i, resourceRecord := range g.Resources {
		// If zone ID exists, delete zone name
		if _, zoneIDExist := resourceRecord.Item["zone_id"]; zoneIDExist {
			delete(g.Resources[i].Item, "zone")
		}

		if resourceRecord.InstanceInfo.Type == "cloudflare_page_rule" {
			if resourceRecord.Item["priority"].(string) == "0" {
				delete(g.Resources[i].Item, "priority")
			}
		}
	}

	return nil
}
