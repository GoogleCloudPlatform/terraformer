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

type AppSWAGenerator struct {
	OktaService
}

func (g *AppSWAGenerator) createResources(appList []okta.ListApplications200ResponseInner) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, app := range appList {
		if app.BrowserPluginApplication != nil && app.BrowserPluginApplication.Id != nil && app.BrowserPluginApplication.Label != "" {
			resources = append(resources, terraformutils.NewSimpleResource(
				*app.BrowserPluginApplication.Id,
				normalizeResourceName(*app.BrowserPluginApplication.Id+"_"+app.BrowserPluginApplication.Label),
				"okta_app_swa",
				"okta",
				[]string{},
			))
		}
	}
	return resources
}

func (g *AppSWAGenerator) InitResources() error {
	ctx, client, err := g.ClientV5()
	if err != nil {
		return err
	}

	appList, resp, err := client.ApplicationAPI.ListApplications(ctx).Execute()
	if err != nil {
		return fmt.Errorf("error listing applications: %w", err)
	}

	allApplications := appList

	for resp.HasNextPage() {
		var nextAppList []okta.ListApplications200ResponseInner
		resp, err = resp.Next(&nextAppList)
		if err != nil {
			return fmt.Errorf("error fetching next page: %w", err)
		}

		allApplications = append(allApplications, nextAppList...)
	}

	g.Resources = g.createResources(allApplications)
	return nil
}

func (g *AppSWAGenerator) PostConvertHook() error {
	for i := range g.Resources {
		g.Resources[i].Item = escapeDollar(g.Resources[i].Item)
	}
	return nil
}
