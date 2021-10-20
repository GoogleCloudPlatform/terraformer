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
	"strings"
)

type RulesetGenerator struct {
	PagerDutyService
}

func (g *RulesetGenerator) createRulesetResources(client *pagerduty.Client) error {
	resp, _, err := client.Rulesets.List()
	if err != nil {
		return err
	}

	for _, ruleset := range resp.Rulesets {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			ruleset.ID,
			strings.Replace(ruleset.Name, " ", "_", -1),
			"pagerduty_ruleset",
			g.ProviderName,
			[]string{},
		))
	}

	return nil
}

// golangci-lint says this function isn't used anywhere. Do we need it? Commenting it out to make the linter happy
/*
func (g *RulesetGenerator) createRulesetRuleResources(client *pagerduty.Client) error {
	resp, _, err := client.Rulesets.List()
	if err != nil {
		return err
	}

	for _, ruleset := range resp.Rulesets {
		rules, _, err := client.Rulesets.ListRules(ruleset.ID)
		if err != nil {
			return err
		}

		for _, rule := range rules.Rules {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				rule.ID,
				rule.ID,
				"pagerduty_ruleset_rule",
				g.ProviderName,
				[]string{},
			))
		}
	}

	return nil
}
*/
func (g *RulesetGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createRulesetResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
