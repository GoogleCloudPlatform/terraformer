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

type AppSignOnPolicyGenerator struct {
	OktaService
}

func (g AppSignOnPolicyGenerator) createResources(policies []okta.ListPolicies200ResponseInner) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, policy := range policies {
		if policy.AccessPolicy == nil {
			continue
		}

		resourceName := normalizeResourceNameWithRandom(policy.AccessPolicy.GetName(), true)
		resourceID := policy.AccessPolicy.GetId()

		resources = append(resources, terraformutils.NewSimpleResource(
			resourceID,
			resourceName,
			"okta_app_signon_policy",
			"okta",
			[]string{}))
	}
	return resources
}

func (g *AppSignOnPolicyGenerator) InitResources() error {
	ctx, client, err := g.ClientV5()
	if err != nil {
		return err
	}

	policies, err := getAppSignOnPolicies(ctx, client)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(policies)
	return nil
}

func getAppSignOnPolicies(ctx context.Context, client *okta.APIClient) ([]okta.ListPolicies200ResponseInner, error) {
	policies, _, err := client.PolicyAPI.ListPolicies(ctx).Type_("ACCESS_POLICY").Execute()
	if err != nil {
		return nil, err
	}
	return policies, nil
}
