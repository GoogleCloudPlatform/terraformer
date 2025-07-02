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
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/okta/okta-sdk-golang/v5/okta"
)

type AppSignOnPolicyRuleGenerator struct {
	OktaService
}

func (g AppSignOnPolicyRuleGenerator) createResources(signOnPolicyRuleList []okta.ListPolicyRules200ResponseInner, policyID string) []terraformutils.Resource {
	var resources []terraformutils.Resource

	for _, policyRule := range signOnPolicyRuleList {
		if policyRule.AccessPolicyRule == nil {
			continue
		}

		resourceName := normalizeResourceNameWithRandom(policyRule.AccessPolicyRule.GetName(), true)

		resources = append(resources, terraformutils.NewResource(
			policyRule.AccessPolicyRule.GetId(),
			resourceName,
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
	ctx, client, err := g.ClientV5()
	if err != nil {
		return err
	}

	policies, err := getAppSignOnPolicies(ctx, client)
	if err != nil {
		return err
	}

	var allResources []terraformutils.Resource

	for _, policy := range policies {
		if policy.AccessPolicy == nil {
			continue
		}

		policyID := policy.AccessPolicy.GetId()

		policyRules, err := getAppSignOnPolicyRules(ctx, client, policyID)
		if err != nil {
			return err
		}

		resources := g.createResources(policyRules, policyID)

		allResources = append(allResources, resources...)
	}

	g.Resources = allResources

	return nil
}

func getAppSignOnPolicyRules(ctx context.Context, client *okta.APIClient, policyID string) ([]okta.ListPolicyRules200ResponseInner, error) {
	policyRules, _, err := client.PolicyAPI.ListPolicyRules(ctx, policyID).Execute()
	if err != nil {
		return nil, err
	}
	return policyRules, nil
}
