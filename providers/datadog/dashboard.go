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
	"fmt"

	datadog "github.com/zorkian/go-datadog-api"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

var (
	// DashboardAllowEmptyValues ...
	DashboardAllowEmptyValues = []string{"tags."}
)

// DashboardGenerator ...
type DashboardGenerator struct {
	DatadogService
}

func (DashboardGenerator) createResources(dashboards []datadog.BoardLite) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, dashboard := range dashboards {
		resourceName := dashboard.GetId()
		resources = append(resources, terraform_utils.NewSimpleResource(
			resourceName,
			fmt.Sprintf("dashboard_%s", resourceName),
			"datadog_dashboard",
			"datadog",
			DashboardAllowEmptyValues,
		))
	}

	return resources
}

// InitResources Generate TerraformResources from Datadog API,
// from each dashboard create 1 TerraformResource.
// Need Dashboard ID as ID for terraform resource
func (g *DashboardGenerator) InitResources() error {
	client := datadog.NewClient(g.Args["api-key"].(string), g.Args["app-key"].(string))
	_, err := client.Validate()
	if err != nil {
		return err
	}
	boards, err := client.GetBoards()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(boards)
	return nil
}
