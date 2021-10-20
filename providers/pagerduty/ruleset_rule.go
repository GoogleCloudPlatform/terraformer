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

package pagerduty

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	pagerduty "github.com/heimweh/go-pagerduty/pagerduty"

	"fmt"
	"strconv"
	"strings"
)

type RulesetRuleGenerator struct {
	PagerDutyService
}

func (g *RulesetRuleGenerator) createRulesetRuleResources(client *pagerduty.Client) error {
	respRulesets, _, err := client.Rulesets.List()
	if err != nil {
		return err
	}
	for _, ruleset := range respRulesets.Rulesets {
		resp, _, err := client.Rulesets.ListRules(ruleset.ID)
		if err != nil {
			return err
		}
		for _, rulesetRule := range resp.Rules {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				rulesetRule.ID,
				fmt.Sprintf("%s_%s", strings.Replace(ruleset.Name, " ", "_", -1), strconv.Itoa(*rulesetRule.Position)),
				"pagerduty_ruleset_rule",
				g.ProviderName,
				map[string]string{
					"ruleset": ruleset.ID,
				},
				[]string{},
				map[string]interface{}{},
			))
		}
	}

	return nil
}

func (g *RulesetRuleGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createRulesetRuleResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
