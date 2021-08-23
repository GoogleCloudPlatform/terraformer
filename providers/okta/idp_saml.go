// Copyright 2021 The Terraformer Authors.
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

type IdpSAMLGenerator struct {
	OktaService
}

func (g IdpSAMLGenerator) createResources(idpSAMLList []*okta.IdentityProvider) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, idp := range idpSAMLList {
		resources = append(resources, terraformutils.NewSimpleResource(
			idp.Id,
			"idp_"+normalizeResourceName(idp.Type+"_"+idp.Name),
			"okta_idp_saml",
			"okta",
			[]string{}))

	}
	return resources
}

func (g *IdpSAMLGenerator) InitResources() error {
	ctx, client, err := g.Client()
	if err != nil {
		return err
	}

	identityProviders, err := getIdpSAML(ctx, client)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(identityProviders)
	return nil
}

func getIdpSAML(ctx context.Context, client *okta.Client) ([]*okta.IdentityProvider, error) {
	qp := &query.Params{Type: "SAML2", Limit: 1}
	output, resp, err := client.IdentityProvider.ListIdentityProviders(ctx, qp)
	if err != nil {
		return nil, err
	}

	for resp.HasNextPage() {
		var nextIdpSAMLSet []*okta.IdentityProvider
		resp, err = resp.Next(ctx, &nextIdpSAMLSet)
		if err != nil {
			return nil, err
		}
		output = append(output, nextIdpSAMLSet...)
	}

	return output, nil
}
