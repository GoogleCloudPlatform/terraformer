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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	cf "github.com/cloudflare/cloudflare-go"
)

type FirewallGenerator struct {
	CloudflareService
}

func (*FirewallGenerator) createZoneLockdownsResources(api *cf.API, zoneID, zoneName string) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}
	page := 1

	for {
		zonelockdowns, err := api.ListZoneLockdowns(zoneID, page)
		if err != nil {
			return resources, err
		}
		for _, zonelockdown := range zonelockdowns.Result {
			resources = append(resources, terraformutils.NewResource(
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

func (g *FirewallGenerator) createAccountAccessRuleResources(api *cf.API) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}
	rules, err := api.ListAccountAccessRules(api.AccountID, cf.AccessRule{}, 1)
	if err != nil {
		return resources, err
	}

	totalPages := rules.TotalPages
	for _, rule := range rules.Result {
		resources = append(resources, terraformutils.NewSimpleResource(
			rule.ID,
			rule.ID,
			"cloudflare_access_rule",
			"cloudflare",
			[]string{},
		))
	}

	for page := 2; page <= totalPages; page++ {
		rules, err := api.ListAccountAccessRules(api.AccountID, cf.AccessRule{}, page)
		if err != nil {
			return resources, err
		}
		for _, rule := range rules.Result {
			resources = append(resources, terraformutils.NewSimpleResource(
				rule.ID,
				rule.ID,
				"cloudflare_access_rule",
				"cloudflare",
				[]string{},
			))
		}
	}

	return resources, nil
}

func (*FirewallGenerator) createZoneAccessRuleResources(api *cf.API, zoneID, zoneName string) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}
	rules, err := api.ListZoneAccessRules(zoneID, cf.AccessRule{}, 1)
	if err != nil {
		return resources, err
	}

	totalPages := rules.TotalPages
	for _, r := range rules.Result {
		if strings.Compare(r.Scope.Type, "organization") != 0 {
			resources = append(resources, terraformutils.NewResource(
				r.ID,
				fmt.Sprintf("%s_%s", zoneName, r.ID),
				"cloudflare_access_rule",
				"cloudflare",
				map[string]string{
					"zone_id": zoneID,
				},
				[]string{},
				map[string]interface{}{},
			))
		}
	}

	for page := 2; page <= totalPages; page++ {
		rules, err := api.ListZoneAccessRules(zoneID, cf.AccessRule{}, page)
		if err != nil {
			return resources, err
		}
		for _, r := range rules.Result {
			if strings.Compare(r.Scope.Type, "organization") != 0 {
				resources = append(resources, terraformutils.NewResource(
					r.ID,
					fmt.Sprintf("%s_%s", zoneName, r.ID),
					"cloudflare_access_rule",
					"cloudflare",
					map[string]string{
						"zone_id": zoneID,
					},
					[]string{},
					map[string]interface{}{},
				))
			}
		}
	}

	return resources, nil
}

func (*FirewallGenerator) createFilterResources(api *cf.API, zoneID, zoneName string) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}
	filters, err := api.Filters(zoneID, cf.PaginationOptions{})
	if err != nil {
		return resources, err
	}

	for _, filter := range filters {
		resources = append(resources, terraformutils.NewResource(
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

func (*FirewallGenerator) createFirewallRuleResources(api *cf.API, zoneID, zoneName string) ([]terraformutils.Resource, error) {
	resources := []terraformutils.Resource{}

	fwrules, err := api.FirewallRules(zoneID, cf.PaginationOptions{})
	if err != nil {
		return resources, err
	}
	for _, rule := range fwrules {
		resources = append(resources, terraformutils.NewResource(
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

func (g *FirewallGenerator) createRateLimitResources(api *cf.API, zoneID, zoneName string) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource

	rateLimits, err := api.ListAllRateLimits(zoneID)
	if err != nil {
		return resources, err
	}
	for _, rateLimit := range rateLimits {
		resources = append(resources, terraformutils.NewSimpleResource(
			rateLimit.ID,
			fmt.Sprintf("%s_%s", zoneID, rateLimit.ID),
			"cloudflare_rate_limit",
			"cloudflare",
			[]string{}))
	}

	return resources, nil
}

func (g *FirewallGenerator) InitResources() error {
	api, err := g.initializeAPI()
	if err != nil {
		return err
	}

	if len(api.AccountID) > 0 {
		resources, err := g.createAccountAccessRuleResources(api)
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, resources...)

	}

	zones, err := api.ListZones()
	if err != nil {
		return err
	}

	funcs := []func(*cf.API, string, string) ([]terraformutils.Resource, error){
		g.createFirewallRuleResources,
		g.createFilterResources,
		g.createZoneAccessRuleResources,
		g.createZoneLockdownsResources,
		g.createRateLimitResources,
	}

	for _, zone := range zones {
		for _, f := range funcs {
			// Getting all firewall filters
			tmpRes, err := f(api, zone.ID, zone.Name)
			if err != nil {
				return err
			}
			g.Resources = append(g.Resources, tmpRes...)
		}
	}

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
