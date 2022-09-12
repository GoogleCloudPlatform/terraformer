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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// LogsCustomPipelineAllowEmptyValues ...
	LogsCustomPipelineAllowEmptyValues = []string{"support_rules", "filter"}
)

// LogsCustomPipelineGenerator ...
type LogsCustomPipelineGenerator struct {
	DatadogService
}

func (g *LogsCustomPipelineGenerator) createResources(logsCustomPipelines []datadogV1.LogsPipeline) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, logsCustomPipeline := range logsCustomPipelines {
		// Import logs custom pipelines only
		if !logsCustomPipeline.GetIsReadOnly() {
			resourceName := logsCustomPipeline.GetId()
			resources = append(resources, g.createResource(resourceName))
		}
	}

	return resources
}

func (g *LogsCustomPipelineGenerator) createResource(logsCustomPipelineID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		logsCustomPipelineID,
		fmt.Sprintf("logs_custom_pipeline_%s", logsCustomPipelineID),
		"datadog_logs_custom_pipeline",
		"datadog",
		LogsCustomPipelineAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each custom pipeline create 1 TerraformResource.
// Need LogsPipeline ID as ID for terraform resource
func (g *LogsCustomPipelineGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV1.NewLogsPipelinesApi(datadogClient)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("logs_custom_pipeline") {
			for _, value := range filter.AcceptableValues {
				logsCustomPipeline, _, err := api.GetLogsPipeline(auth, value)
				if err != nil {
					return err
				}

				resources = append(resources, g.createResource(logsCustomPipeline.GetId()))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	logsCustomPipelines, _, err := api.ListLogsPipelines(auth)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(logsCustomPipelines)
	return nil
}

func (g *LogsCustomPipelineGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		for k, v := range r.Item {
			// Hack to properly escape `%{` used in pipeline processors
			if k == "processor" {
				var z interface{}
				jsonByte, err := json.Marshal(v)
				if err != nil {
					continue
				}
				jsonByte = []byte(strings.ReplaceAll(string(jsonByte), "%{", "%%{"))
				if err = json.Unmarshal(jsonByte, &z); err != nil {
					continue
				}
				g.Resources[i].Item[k] = z
			}
		}
	}
	return nil
}
