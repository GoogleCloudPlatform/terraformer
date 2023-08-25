// Copyright 2022 The Terraformer Authors.
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

package sumologic

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/iancoleman/strcase"
	sumologic "github.com/sumovishal/sumologic-go-sdk/api"
)

type DashboardGenerator struct {
	SumoLogicService
}

func (g *DashboardGenerator) createResources(dashboards []sumologic.Dashboard) []terraformutils.Resource {
	resources := make([]terraformutils.Resource, len(dashboards))

	for i, dashboard := range dashboards {
		title := strcase.ToSnake(replaceSpaceAndDash(dashboard.Title))

		resource := terraformutils.NewSimpleResource(
			*dashboard.Id,
			fmt.Sprintf("%s-%s", title, *dashboard.Id),
			"sumologic_dashboard",
			g.ProviderName,
			[]string{})
		resources[i] = resource
	}

	return resources
}

func (g *DashboardGenerator) InitResources() error {
	client := g.Client()

	var resources []terraformutils.Resource
	var dashboards []sumologic.Dashboard

	req := client.DashboardManagementApi.ListDashboards(g.AuthCtx())
	req = req.Limit(100)

	respBody, _, err := client.DashboardManagementApi.ListDashboardsExecute(req)
	if err != nil {
		return err
	}
	dashboards = respBody.Dashboards
	for respBody.Next != nil {
		req = req.Token(respBody.GetNext())
		respBody, _, err = client.DashboardManagementApi.ListDashboardsExecute(req)
		if err != nil {
			return err
		}
		dashboards = append(dashboards, respBody.Dashboards...)
	}

	resources = g.createResources(dashboards)
	g.Resources = resources
	return nil
}
