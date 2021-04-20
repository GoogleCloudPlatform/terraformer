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
	datadogClientV1 := g.Args["datadogClientV1"].(*datadogV1.APIClient)
	authV1 := g.Args["authV1"].(context.Context)

	var slos []datadogV1.ServiceLevelObjective
	resp, _, err := datadogClientV1.ServiceLevelObjectivesApi.ListSLOs(authV1).Execute()
	if err != nil {
		return err
	}

	slos = append(slos, resp.GetData()...)
	g.Resources = g.createResources(slos)
	return nil
}
