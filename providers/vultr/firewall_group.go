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

package vultr

import (
	"context"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/vultr/govultr"
)

type FirewallGroupGenerator struct {
	VultrService
}

func (g *FirewallGroupGenerator) loadFirewallGroups(client *govultr.Client) ([]govultr.FirewallGroup, error) {
	firewallGroups, err := client.FirewallGroup.List(context.Background())
	if err != nil {
		return nil, err
	}
	for _, firewallGroup := range firewallGroups {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			firewallGroup.FirewallGroupID,
			firewallGroup.FirewallGroupID,
			"vultr_firewall_group",
			"vultr",
			[]string{}))
	}
	return firewallGroups, nil
}

func (g *FirewallGroupGenerator) loadFirewallRulesByIPType(client *govultr.Client, firewallGroupID string, ipType string) error {
	firewallRules, err := client.FirewallRule.ListByIPType(context.Background(), firewallGroupID, ipType)
	if err != nil {
		return err
	}
	for _, firewallRule := range firewallRules {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			strconv.Itoa(firewallRule.RuleNumber),
			strconv.Itoa(firewallRule.RuleNumber),
			"vultr_firewall_rule",
			"vultr",
			map[string]string{
				"firewall_group_id": firewallGroupID,
				"ip_type":           ipType,
			},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}

func (g *FirewallGroupGenerator) InitResources() error {
	client := g.generateClient()
	firewallGroups, err := g.loadFirewallGroups(client)
	if err != nil {
		return err
	}
	for _, firewallGroup := range firewallGroups {
		err := g.loadFirewallRulesByIPType(client, firewallGroup.FirewallGroupID, "v4")
		if err != nil {
			return err
		}
		err = g.loadFirewallRulesByIPType(client, firewallGroup.FirewallGroupID, "v6")
		if err != nil {
			return err
		}
	}
	return nil
}
