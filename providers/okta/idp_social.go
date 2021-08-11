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

type IdpSocialGenerator struct {
	OktaService
}

func (g IdpSocialGenerator) createResources(idpSocialList []*okta.IdentityProvider) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, idp := range idpSocialList {
		resources = append(resources, terraformutils.NewSimpleResource(
			idp.Id,
			"idp_"+normalizeResourceName(idp.Type+"_"+idp.Name),
			"okta_idp_social",
			"okta",
			[]string{}))

	}
	return resources
}

// Generate Terraform Resources from Okta API,
func (g *IdpSocialGenerator) InitResources() error {
	ctx, client, err := g.Client()
	if err != nil {
		return err
	}

	identityProviders, err := getIdpSocials(ctx, client)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(identityProviders)
	return nil
}

func getIdpSocials(ctx context.Context, client *okta.Client) ([]*okta.IdentityProvider, error) {
	idpSocialTypes := []string{"APPLE", "FACEBOOK", "GOOGLE", "LINKEDIN", "MICROSOFT"}
	var allIDPSocials []*okta.IdentityProvider

	for _, idpSocialType := range idpSocialTypes {
		qp := &query.Params{Type: idpSocialType, Limit: 1}
		output, resp, err := client.IdentityProvider.ListIdentityProviders(ctx, qp)
		if err != nil {
			return nil, err
		}

		for resp.HasNextPage() {
			var nextIdpSocialSet []*okta.IdentityProvider
			resp, _ = resp.Next(ctx, &nextIdpSocialSet)
			output = append(output, nextIdpSocialSet...)
		}

		allIDPSocials = append(allIDPSocials, output...)
	}

	return allIDPSocials, nil
}
