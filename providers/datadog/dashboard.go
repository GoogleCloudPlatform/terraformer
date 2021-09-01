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

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// DashboardAllowEmptyValues ...
	DashboardAllowEmptyValues = []string{"tags.", "manage_status_definition.*.query"}
)

// DashboardGenerator ...
type DashboardGenerator struct {
	DatadogService
}

func (g *DashboardGenerator) createResources(dashboards []datadogV1.DashboardSummaryDefinition) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, dashboard := range dashboards {
		resourceName := dashboard.GetId()
		resources = append(resources, g.createResource(resourceName))
	}

	return resources
}

func (g *DashboardGenerator) createResource(dashboardID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		dashboardID,
		fmt.Sprintf("dashboard_%s", dashboardID),
		"datadog_dashboard",
		"datadog",
		DashboardAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each dashboard create 1 TerraformResource.
// Need Dashboard ID as ID for terraform resource
func (g *DashboardGenerator) InitResources() error {
	datadogClientV1 := g.Args["datadogClientV1"].(*datadogV1.APIClient)
	authV1 := g.Args["authV1"].(context.Context)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("dashboard") {
			for _, value := range filter.AcceptableValues {
				dashboard, _, err := datadogClientV1.DashboardsApi.GetDashboard(authV1, value)
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

	summary, _, err := datadogClientV1.DashboardsApi.ListDashboards(authV1)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(summary.GetDashboards())
	return nil
}
