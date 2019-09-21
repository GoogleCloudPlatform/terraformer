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

type FirewallGenerator struct {
	CloudflareService
}

func (*FirewallGenerator) createZoneLockdownsResources(api *cf.API, zoneID, zoneName string) ([]terraform_utils.Resource, error) {
	resources := []terraform_utils.Resource{}
	page := 1

	for {
		zonelockdowns, err := api.ListZoneLockdowns(zoneID, page)
		if err != nil {
			log.Println(err)
			return resources, err
		}
		for _, zonelockdown := range zonelockdowns.Result {
			resources = append(resources, terraform_utils.NewResource(
				zonelockdown.ID,
				fmt.Sprintf("%s_%s", zoneName, zonelockdown.ID),
				"cloudflare_zone_lockdown",
				"cloudflare",
				map[string]string{
					"zone_id": zoneID,
					"zone":    zoneName,
				},
				[]string{},
				map[string]interface{}{},
			))
		}

		if zonelockdowns.TotalPages > page {
			page++
		} else {
			break
		}
	}

	return resources, nil
}

func (*FirewallGenerator) createAccessRuleResources(api *cf.API, zoneID, zoneName string) ([]terraform_utils.Resource, error) {
	resources := []terraform_utils.Resource{}
	accessRules, err := api.ListZoneAccessRules(zoneID, cf.AccessRule{}, 1)
	if err != nil {
		log.Println(err)
		return resources, err
	}

	for _, r := range accessRules.Result {
		resources = append(resources, terraform_utils.NewResource(
			r.ID,
			fmt.Sprintf("%s_%s", zoneName, r.ID),
			"cloudflare_access_rule",
			"cloudflare",
			map[string]string{
				"zone_id": zoneID,
				"zone":    zoneName,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources, nil
}

func (*FirewallGenerator) createFilterResources(api *cf.API, zoneID, zoneName string) ([]terraform_utils.Resource, error) {
	resources := []terraform_utils.Resource{}
	filters, err := api.Filters(zoneID, cf.PaginationOptions{})
	if err != nil {
		log.Println(err)
		return resources, err
	}

	for _, filter := range filters {
		resources = append(resources, terraform_utils.NewResource(
			filter.ID,
			fmt.Sprintf("%s_%s", zoneName, filter.ID),
			"cloudflare_filter",
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

func (*FirewallGenerator) createFirewallRuleResources(api *cf.API, zoneID, zoneName string) ([]terraform_utils.Resource, error) {
	resources := []terraform_utils.Resource{}

	fwrules, err := api.FirewallRules(zoneID, cf.PaginationOptions{})
	if err != nil {
		log.Println(err)
		return resources, err
	}
	for _, rule := range fwrules {
		resources = append(resources, terraform_utils.NewResource(
			rule.ID,
			fmt.Sprintf("%s_%s", zoneName, rule.ID),
			"cloudflare_firewall_rule",
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

func (g *FirewallGenerator) InitResources() error {
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
		g.createFirewallRuleResources,
		g.createFilterResources,
		g.createAccessRuleResources,
		g.createZoneLockdownsResources,
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

	g.PopulateIgnoreKeys()
	return nil
}

func (g *FirewallGenerator) PostConvertHook() error {
	for i, resourceRecord := range g.Resources {
		// If Zone Name exists, delete ZoneID
		if _, zoneIDExist := resourceRecord.Item["zone_id"]; zoneIDExist {
			delete(g.Resources[i].Item, "zone")
		}

		if resourceRecord.InstanceInfo.Type == "cloudflare_firewall_rule" {
			if resourceRecord.Item["priority"].(string) == "0" {
				delete(g.Resources[i].Item, "priority")
			}
		}

		// Reference to 'cloudflare_filter' resource in 'cloudflare_firewall_rule'
		if resourceRecord.InstanceInfo.Type == "cloudflare_filter" {
			continue
		}
		filterID := resourceRecord.Item["filter_id"]
		for _, filterResource := range g.Resources {
			if filterResource.InstanceInfo.Type != "cloudflare_filter" {
				continue
			}
			if filterID == filterResource.InstanceState.ID {
				g.Resources[i].Item["filter_id"] = "${cloudflare_filter." + filterResource.ResourceName + ".id}"
			}
		}
	}

	return nil
}
