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
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

type PasswordPolicyGenerator struct {
	OktaService
}

func (g PasswordPolicyGenerator) createResources(passwordPolicyList []*okta.Policy) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, passwordPolicy := range passwordPolicyList {
		resourceName := normalizeResourceName(passwordPolicy.Name)
		resourceType := "okta_policy_password"
		if passwordPolicy.Name == "Default Policy" {
			resourceType = "okta_policy_password_default"
		}
		resources = append(resources, terraformutils.NewSimpleResource(
			passwordPolicy.Id,
			"policy_password_"+resourceName,
			resourceType,
			"okta",
			[]string{}))
	}
	return resources
}

func (g *PasswordPolicyGenerator) InitResources() error {
	var output []*okta.Policy
	ctx, client, e := g.Client()
	if e != nil {
		return e
	}

	output, _ = getPasswordPolicies(ctx, client)
	g.Resources = g.createResources(output)
	return nil
}

func getPasswordPolicies(ctx context.Context, client *okta.Client) ([]*okta.Policy, error) {
	qp := query.NewQueryParams(query.WithType("PASSWORD"))
	output, resp, err := client.Policy.ListPolicies(ctx, qp)
	if err != nil {
		return nil, err
	}

	for resp.HasNextPage() {
		var nextPolicySet []*okta.Policy
		resp, _ = resp.Next(ctx, &nextPolicySet)
		output = append(output, nextPolicySet...)
	}

	return output, nil
}
