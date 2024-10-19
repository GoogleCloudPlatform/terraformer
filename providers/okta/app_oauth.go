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
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/okta/okta-sdk-golang/v5/okta"
)

type AppOAuthGenerator struct {
	OktaService
}

func (g *AppOAuthGenerator) createResources(appList []okta.ListApplications200ResponseInner) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, app := range appList {
		if app.OpenIdConnectApplication != nil {
			resources = append(resources, terraformutils.NewSimpleResource(
				*app.OpenIdConnectApplication.Id,
				normalizeResourceName(*app.OpenIdConnectApplication.Id+"_"+app.OpenIdConnectApplication.Label),
				"okta_app_oauth",
				"okta",
				[]string{},
			))
		}
	}
	return resources
}

// Generate Terraform Resources from Okta API,
func (g *AppOAuthGenerator) InitResources() error {
	ctx, client, e := g.ClientV5()
	if e != nil {
		return e
	}

	appList, _, err := client.ApplicationAPI.ListApplications(ctx).Filter("name eq \"oidc_client\"").Execute()
	if err != nil {
		return fmt.Errorf("error listing applications: %w", err)
	}

	g.Resources = g.createResources(appList)
	return nil
}
