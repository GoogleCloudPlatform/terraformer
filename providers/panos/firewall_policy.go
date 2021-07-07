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

package panos

import (
	"encoding/base64"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/PaloAltoNetworks/pango"
	"github.com/PaloAltoNetworks/pango/util"
)

type FirewallPolicyGenerator struct {
	PanosService
}

func (g *FirewallPolicyGenerator) createResourcesFromList(o getGeneric, terraformResourceName string) (resources []terraformutils.Resource) {
	l, err := o.i.(getListWithOneArg).GetList(o.params[0])
	if err != nil || len(l) == 0 {
		return []terraformutils.Resource{}
	}

	var positionReference string
	id := g.vsys + ":" + strconv.Itoa(util.MoveTop) + "::"

	for k, r := range l {
		if k > 0 {
			id = g.vsys + ":" + strconv.Itoa(util.MoveAfter) + ":" + positionReference + ":"
		}

		id += base64.StdEncoding.EncodeToString([]byte(r))
		positionReference = r

		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(r),
			terraformResourceName,
			"panos",
			[]string{},
		))
	}

	return resources
}

func (g *FirewallPolicyGenerator) createNATRuleGroupResources() []terraformutils.Resource {
	return g.createResourcesFromList(getGeneric{g.client.(*pango.Firewall).Policies.Nat, []string{g.vsys}}, "panos_nat_rule_group")
}

func (g *FirewallPolicyGenerator) createPBFRuleGroupResources() []terraformutils.Resource {
	return g.createResourcesFromList(getGeneric{g.client.(*pango.Firewall).Policies.PolicyBasedForwarding, []string{g.vsys}}, "panos_pbf_rule_group")
}

func (g *FirewallPolicyGenerator) createSecurityRuleGroupResources() []terraformutils.Resource {
	return g.createResourcesFromList(getGeneric{g.client.(*pango.Firewall).Policies.Security, []string{g.vsys}}, "panos_security_rule_group")
}

func (g *FirewallPolicyGenerator) InitResources() error {
	if err := g.Initialize(); err != nil {
		return err
	}

	g.Resources = append(g.Resources, g.createNATRuleGroupResources()...)
	g.Resources = append(g.Resources, g.createPBFRuleGroupResources()...)
	g.Resources = append(g.Resources, g.createSecurityRuleGroupResources()...)

	return nil
}

func (g *FirewallPolicyGenerator) PostConvertHook() error {
	for _, res := range g.Resources {
		if res.InstanceInfo.Type == "panos_nat_rule_group" {
			for _, rule := range res.Item["rule"].([]interface{}) {
				if _, ok := rule.(map[string]interface{})["translated_packet"]; ok {
					a := rule.(map[string]interface{})["translated_packet"].([]interface{})
					for _, b := range a {
						if _, okb := b.(map[string]interface{})["source"]; !okb {
							b.(map[string]interface{})["source"] = make(map[string]interface{})
						}
					}

					for _, b := range a {
						if _, okb := b.(map[string]interface{})["destination"]; !okb {
							b.(map[string]interface{})["destination"] = make(map[string]interface{})
						}
					}
				}
			}
		}

		if res.InstanceInfo.Type == "panos_security_rule_group" {
			for _, rule := range res.Item["rule"].([]interface{}) {
				if _, ok := rule.(map[string]interface{})["hip_profiles"]; !ok {
					rule.(map[string]interface{})["hip_profiles"] = []string{"any"}
				}
			}
		}
	}

	return nil
}
