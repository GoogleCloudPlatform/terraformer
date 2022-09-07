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
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewLogsIndexesApi(datadogClient)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("logs_index") {
			for _, value := range filter.AcceptableValues {
				logsIndex, _, err := api.GetLogsIndex(auth, value)
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

	logsIndexList, _, err := api.ListLogIndexes(auth)
	logsIndex := logsIndexList.GetIndexes()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(logsIndex)
	return nil
}
