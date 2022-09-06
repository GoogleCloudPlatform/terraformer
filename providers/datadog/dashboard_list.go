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
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// DashboardListAllowEmptyValues ...
	DashboardListAllowEmptyValues = []string{}
)

// DashboardListGenerator ...
type DashboardListGenerator struct {
	DatadogService
}

func (g *DashboardListGenerator) createResources(dashboardLists []datadogV1.DashboardList) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, dashboardList := range dashboardLists {
		resourceID := strconv.FormatInt(dashboardList.GetId(), 10)
		resources = append(resources, g.createResource(resourceID))
	}

	return resources
}

func (g *DashboardListGenerator) createResource(dashboardListID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		dashboardListID,
		fmt.Sprintf("dashboard_list_%s", dashboardListID),
		"datadog_dashboard_list",
		"datadog",
		DashboardListAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each dashboard_list create 1 TerraformResource.
// Need DashboardList ID as ID for terraform resource
func (g *DashboardListGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewDashboardListsApi(datadogClient)

	dlResponse, _, err := api.ListDashboardLists(auth)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(dlResponse.GetDashboardLists())
	return nil
}
