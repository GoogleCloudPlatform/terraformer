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

type AppSignOnPolicyGenerator struct {
	OktaService
}

func (g AppSignOnPolicyGenerator) createResources(appSignOnPolicyList []*okta.Policy) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, appSignOnPolicy := range appSignOnPolicyList {
		resourceName := normalizeResourceName(appSignOnPolicy.Name)
		resourceType := "okta_app_signon_policy"

		resources = append(resources, terraformutils.NewSimpleResource(
			appSignOnPolicy.Id,
			"app_signon_policy"+resourceName,
			resourceType,
			"okta",
			[]string{"description"}))
	}
	return resources
}

func (g *AppSignOnPolicyGenerator) InitResources() error {
	var output []*okta.Policy
	ctx, client, e := g.Client()
	if e != nil {
		return e
	}

	output, _ = getAppSignOnPolicies(ctx, client)
	g.Resources = g.createResources(output)
	return nil
}

func getAppSignOnPolicies(ctx context.Context, client *okta.Client) ([]*okta.Policy, error) {
	qp := query.NewQueryParams(query.WithType("ACCESS_POLICY"))
	var policies []*okta.Policy
	data, resp, err := client.Policy.ListPolicies(ctx, qp)
	if err != nil {
		return nil, err
	}

	for resp.HasNextPage() {
		var nextPolicies []*okta.Policy
		resp, _ = resp.Next(ctx, &nextPolicies)
		policies = append(policies, nextPolicies...)
	}
	for _, p := range data {
		policies = append(policies, p.(*okta.Policy))
	}
	return policies, nil
}
