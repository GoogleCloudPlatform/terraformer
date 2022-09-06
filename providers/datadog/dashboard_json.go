// Copyright 2018 The Terraformer Authors.
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

package datadog

import (
	"context"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// DashboardJSONAllowEmptyValues ...
	DashboardJSONAllowEmptyValues = []string{"tags."}
)

// DashboardJSONGenerator ...
type DashboardJSONGenerator struct {
	DatadogService
}

func (g *DashboardJSONGenerator) createResources(dashboards []datadogV1.DashboardSummaryDefinition) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, dashboard := range dashboards {
		resourceName := dashboard.GetId()
		resources = append(resources, g.createResource(resourceName))
	}

	return resources
}

func (g *DashboardJSONGenerator) createResource(dashboardID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		dashboardID,
		fmt.Sprintf("dashboard_json_%s", dashboardID),
		"datadog_dashboard_json",
		"datadog",
		DashboardJSONAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each dashboard_json create 1 TerraformResource.
// Need Dashboard ID as ID for terraform resource
func (g *DashboardJSONGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewDashboardsApi(datadogClient)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("dashboard_json") {
			for _, value := range filter.AcceptableValues {
				dashboard, _, err := api.GetDashboard(auth, value)
				if err != nil {
					return err
				}

				resources = append(resources, g.createResource(dashboard.GetId()))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	summary, _, err := api.ListDashboards(auth)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(summary.GetDashboards())
	return nil
}
