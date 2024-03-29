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

package okta

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/okta/terraform-provider-okta/sdk"
)

type AppSignOnPolicyRuleGenerator struct {
	OktaService
}

func (g AppSignOnPolicyRuleGenerator) createResources(signOnPolicyRuleList []sdk.PolicyRule, policyID string, policyName string) []terraformutils.Resource {
	var resources []terraformutils.Resource

	for _, policyRule := range signOnPolicyRuleList {
		resources = append(resources, terraformutils.NewResource(
			policyRule.Id,
			"app_policyrule_signon_"+normalizeResourceName(policyName+"_"+policyRule.Name),
			"okta_app_signon_policy_rule",
			"okta",
			map[string]string{
				"policy_id": policyID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}

func (g *AppSignOnPolicyRuleGenerator) InitResources() error {
	var resources []terraformutils.Resource

	ctx, client, e := g.Client()
	if e != nil {
		return e
	}

	appSignOnPolicies, err := getAppSignOnPolicies(ctx, client)
	if err != nil {
		return err
	}

	for _, policy := range appSignOnPolicies {
		output, err := getAppSignOnPolicyRules(g, policy.Id)
		if err != nil {
			return err
		}

		resources = append(resources, g.createResources(output, policy.Id, policy.Name)...)
	}

	g.Resources = resources
	return nil
}

func getAppSignOnPolicyRules(g *AppSignOnPolicyRuleGenerator, policyID string) ([]sdk.PolicyRule, error) {
	ctx, client, e := g.APISupplementClient()
	if e != nil {
		return nil, e
	}

	output, resp, err := client.ListPolicyRules(ctx, policyID)
	if err != nil {
		return nil, e
	}

	for resp.HasNextPage() {
		var nextPolicySet []sdk.PolicyRule
		resp, _ = resp.Next(ctx, &nextPolicySet)
		output = append(output, nextPolicySet...)
	}

	return output, nil
}
