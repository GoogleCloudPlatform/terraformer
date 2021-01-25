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
	// LogsIndexAllowEmptyValues ...
	LogsIndexAllowEmptyValues = []string{"filter"}
)

// LogsIndexGenerator ...
type LogsIndexGenerator struct {
	DatadogService
}

func (g *LogsIndexGenerator) createResources(logsIndexes []datadogV1.LogsIndex) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, logsIndex := range logsIndexes {
		resourceName := logsIndex.GetName()
		resources = append(resources, g.createResource(resourceName))
	}

	return resources
}

func (g *LogsIndexGenerator) createResource(logsIndexName string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		logsIndexName,
		fmt.Sprintf("logs_index_%s", logsIndexName),
		"datadog_logs_index",
		"datadog",
		LogsIndexAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each index create 1 TerraformResource.
// Need LogsIndex Name as ID for terraform resource
func (g *LogsIndexGenerator) InitResources() error {
	datadogClientV1 := g.Args["datadogClientV1"].(*datadogV1.APIClient)
	authV1 := g.Args["authV1"].(context.Context)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("logs_index") {
			for _, value := range filter.AcceptableValues {
				logsIndex, _, err := datadogClientV1.LogsIndexesApi.GetLogsIndex(authV1, value).Execute()
				if err != nil {
					return err
				}

				resources = append(resources, g.createResource(logsIndex.GetName()))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	logsIndexList, _, err := datadogClientV1.LogsIndexesApi.ListLogIndexes(authV1).Execute()
	logsIndex := logsIndexList.GetIndexes()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(logsIndex)
	return nil
}
