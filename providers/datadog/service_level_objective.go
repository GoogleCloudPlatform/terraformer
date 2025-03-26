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
	// ServiceLevelObjectiveAllowEmptyValues ...
	ServiceLevelObjectiveAllowEmptyValues = []string{"tags."}
)

// ServiceLevelObjectiveGenerator ...
type ServiceLevelObjectiveGenerator struct {
	DatadogService
}

func (g *ServiceLevelObjectiveGenerator) createResources(sloList []datadogV1.ServiceLevelObjective) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, slo := range sloList {
		resourceID := slo.GetId()
		resources = append(resources, g.createResource(resourceID))
	}

	return resources
}

func (g *ServiceLevelObjectiveGenerator) createResource(sloID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		sloID,
		fmt.Sprintf("service_level_objective_%s", sloID),
		"datadog_service_level_objective",
		"datadog",
		ServiceLevelObjectiveAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each service_level_objective create 1 TerraformResource.
// Need ServiceLevelObjective ID as ID for terraform resource
func (g *ServiceLevelObjectiveGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewServiceLevelObjectivesApi(datadogClient)

	var slos []datadogV1.ServiceLevelObjective
	resp, _, err := api.ListSLOs(auth)
	if err != nil {
		return err
	}

	slos = append(slos, resp.GetData()...)
	g.Resources = g.createResources(slos)
	return nil
}
