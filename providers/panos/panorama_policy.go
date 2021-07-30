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

type PanoramaPolicyGenerator struct {
	PanosService
}

func (g *PanoramaPolicyGenerator) createResourcesFromList(o getGeneric, terraformResourceName string) (resources []terraformutils.Resource) {
	l, err := o.i.(getListWithTwoArgs).GetList(o.params[0], o.params[1])
	if err != nil || len(l) == 0 {
		return []terraformutils.Resource{}
	}

	var positionReference string
	id := o.params[0] + ":" + o.params[1] + ":" + strconv.Itoa(util.MoveTop) + "::"

	for k, r := range l {
		if k > 0 {
			id = o.params[0] + ":" + o.params[1] + ":" + strconv.Itoa(util.MoveAfter) + ":" + positionReference + ":"
		}

		id += base64.StdEncoding.EncodeToString([]byte(r))
		positionReference = r

		resources = append(resources, terraformutils.NewSimpleResource(
			id,
			normalizeResourceName(o.params[0]+":"+o.params[1]+":"+r),
			terraformResourceName,
			"panos",
			[]string{},
		))
	}

	return resources
}

func (g *PanoramaPolicyGenerator) createNATRuleGroupResources(dg string) (resources []terraformutils.Resource) {
	resources = append(resources, g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Policies.Nat, []string{dg, util.PreRulebase}},
		"panos_panorama_nat_rule_group")...,
	)

	resources = append(resources, g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Policies.Nat, []string{dg, util.Rulebase}},
		"panos_panorama_nat_rule_group")...,
	)

	resources = append(resources, g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Policies.Nat, []string{dg, util.PostRulebase}},
		"panos_panorama_nat_rule_group")...,
	)

	return resources
}

func (g *PanoramaPolicyGenerator) createPBFRuleGroupResources(dg string) (resources []terraformutils.Resource) {
	resources = append(resources, g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Policies.PolicyBasedForwarding, []string{dg, util.PreRulebase}},
		"panos_panorama_pbf_rule_group")...,
	)

	resources = append(resources, g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Policies.PolicyBasedForwarding, []string{dg, util.Rulebase}},
		"panos_panorama_pbf_rule_group")...,
	)

	resources = append(resources, g.createResourcesFromList(
		getGeneric{g.client.(*pango.Panorama).Policies.PolicyBasedForwarding, []string{dg, util.PostRulebase}},
		"panos_panorama_pbf_rule_group")...,
	)

	return resources
}

func (g *PanoramaPolicyGenerator) createSecurityRuleGroupRulebaseResources(dg, rulebase string) (resources []terraformutils.Resource) {
	l, err := g.client.(*pango.Panorama).Policies.Security.GetList(dg, rulebase)
	if err != nil || len(l) == 0 {
		return []terraformutils.Resource{}
	}

	var positionReference string
	id := dg + ":" + rulebase + ":" + strconv.Itoa(util.MoveTop) + "::"

	for k, r := range l {
		if k > 0 {
			id = dg + ":" + rulebase + ":" + strconv.Itoa(util.MoveAfter) + ":" + positionReference + ":"
		}

		id += base64.StdEncoding.EncodeToString([]byte(r))
		positionReference = r

		resources = append(resources, terraformutils.NewResource(
			id,
			normalizeResourceName(dg+":"+rulebase+":"+r),
			"panos_panorama_security_rule_group",
			"panos",
			map[string]string{
				"device_group":    dg,
				"rulebase":        rulebase,
				"rule.#":          "1", // Add just enough attributes to make the refresh work...
				"rule.0.name":     r,   // Add just enough attributes to make the refresh work...
				"rule.0.target.#": "0", // Add just enough attributes to make the refresh work...
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}

func (g *PanoramaPolicyGenerator) createSecurityRuleGroupResources(dg string) (resources []terraformutils.Resource) {
	resources = append(resources, g.createSecurityRuleGroupRulebaseResources(dg, util.PreRulebase)...)
	resources = append(resources, g.createSecurityRuleGroupRulebaseResources(dg, util.Rulebase)...)
	resources = append(resources, g.createSecurityRuleGroupRulebaseResources(dg, util.PostRulebase)...)

	return resources
}

func (g *PanoramaPolicyGenerator) InitResources() error {
	if err := g.Initialize(); err != nil {
		return err
	}

	dg, err := g.client.(*pango.Panorama).Panorama.DeviceGroup.GetList()
	if err != nil {
		return err
	}

	for _, v := range dg {
		g.Resources = append(g.Resources, g.createNATRuleGroupResources(v)...)
		g.Resources = append(g.Resources, g.createPBFRuleGroupResources(v)...)
		g.Resources = append(g.Resources, g.createSecurityRuleGroupResources(v)...)
	}

	return nil
}

func (g *PanoramaPolicyGenerator) PostConvertHook() error {
	for _, res := range g.Resources {
		if res.InstanceInfo.Type == "panos_panorama_nat_rule_group" {
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

		if res.InstanceInfo.Type == "panos_panorama_security_rule_group" {
			for _, rule := range res.Item["rule"].([]interface{}) {
				if _, ok := rule.(map[string]interface{})["hip_profiles"]; !ok {
					rule.(map[string]interface{})["hip_profiles"] = []string{"any"}
				}
			}
		}
	}

	return nil
}
