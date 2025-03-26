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
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// LogsMetricAllowEmptyValues ...
	LogsMetricAllowEmptyValues = []string{}
)

// LogsMetricGenerator ...
type LogsMetricGenerator struct {
	DatadogService
}

func (g *LogsMetricGenerator) createResources(logsMetrics []datadogV2.LogsMetricResponseData) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, logsMetric := range logsMetrics {
		resourceName := logsMetric.GetId()
		resources = append(resources, g.createResource(resourceName))
	}

	return resources
}

func (g *LogsMetricGenerator) createResource(logsMetricName string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		logsMetricName,
		fmt.Sprintf("logs_metric_%s", logsMetricName),
		"datadog_logs_metric",
		"datadog",
		LogsMetricAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each log's metric create 1 TerraformResource.
// Need LogsMetric Name as ID for terraform resource
func (g *LogsMetricGenerator) InitResources() error {
	datadogClient := g.Args["datadogClient"].(*datadog.APIClient)
	auth := g.Args["auth"].(context.Context)
	api := datadogV2.NewLogsMetricsApi(datadogClient)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("logs_metric") {
			for _, value := range filter.AcceptableValues {
				logsMetric, _, err := api.GetLogsMetric(auth, value)
				if err != nil {
					return err
				}

				resources = append(resources, g.createResource(logsMetric.Data.GetId()))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	logsMetrics, _, err := api.ListLogsMetrics(auth)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(logsMetrics.GetData())
	return nil
}
