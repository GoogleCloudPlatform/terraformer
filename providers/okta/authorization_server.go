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
)

type AuthorizationServerGenerator struct {
	OktaService
}

func (g AuthorizationServerGenerator) createResources(authorizationServerList []*okta.AuthorizationServer) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, authorizationServer := range authorizationServerList {
		resourceType := "okta_auth_server"
		if authorizationServer.Name == "default" {
			resourceType = "okta_auth_server_default"
		}

		resources = append(resources, terraformutils.NewSimpleResource(
			authorizationServer.Id,
			"auth_server_"+authorizationServer.Name,
			resourceType,
			"okta",
			[]string{}))
	}
	return resources
}

func (g *AuthorizationServerGenerator) InitResources() error {
	ctx, client, e := g.Client()
	if e != nil {
		return e
	}

	output, err := getAuthorizationServers(ctx, client)
	if err != nil {
		return e
	}

	g.Resources = g.createResources(output)
	return nil
}

func getAuthorizationServers(ctx context.Context, client *okta.Client) ([]*okta.AuthorizationServer, error) {
	output, resp, err := client.AuthorizationServer.ListAuthorizationServers(ctx, nil)
	if err != nil {
		return nil, err
	}

	for resp.HasNextPage() {
		var nextAuthorizationServerSet []*okta.AuthorizationServer
		resp, _ = resp.Next(ctx, &nextAuthorizationServerSet)
		output = append(output, nextAuthorizationServerSet...)
	}

	return output, nil
}
