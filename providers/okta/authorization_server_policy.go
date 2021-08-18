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
	"github.com/okta/okta-sdk-golang/v2/okta"
)

type AuthorizationServerPolicyGenerator struct {
	OktaService
}

func (g AuthorizationServerPolicyGenerator) createResources(authorizationServerPolicyList []*okta.AuthorizationServerPolicy, authorizationServerID string, authorizationServerName string) []terraformutils.Resource {
	var resources []terraformutils.Resource

	for _, authorizationServerPolicy := range authorizationServerPolicyList {
		resources = append(resources, terraformutils.NewResource(
			authorizationServerPolicy.Id,
			normalizeResourceName("auth_server_"+authorizationServerName+"_policy_"+authorizationServerPolicy.Name),
			"okta_auth_server_policy",
			"okta",
			map[string]string{
				"auth_server_id": authorizationServerID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}

func (g *AuthorizationServerPolicyGenerator) InitResources() error {
	var resources []terraformutils.Resource
	ctx, client, e := g.Client()
	if e != nil {
		return e
	}

	authorizationServers, err := getAuthorizationServers(ctx, client)
	if err != nil {
		return err
	}

	for _, authorizationServer := range authorizationServers {
		output, _, err := client.AuthorizationServer.ListAuthorizationServerPolicies(ctx, authorizationServer.Id)
		if err != nil {
			return err
		}

		resources = append(resources, g.createResources(output, authorizationServer.Id, authorizationServer.Name)...)
	}

	g.Resources = resources
	return nil
}
