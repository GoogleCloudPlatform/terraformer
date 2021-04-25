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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// MetricMetadataAllowEmptyValues ...
	MetricMetadataAllowEmptyValues = []string{}
)

// MetricMetadataGenerator ...
type MetricMetadataGenerator struct {
	DatadogService
}

func (g *MetricMetadataGenerator) createResource(metricName string) terraformutils.Resource {
	return terraformutils.NewResource(
		metricName,
		fmt.Sprintf("metric_metadata_%s", metricName),
		"datadog_metric_metadata",
		"datadog",
		map[string]string{
			"metric": metricName,
		},
		MetricMetadataAllowEmptyValues,
		map[string]interface{}{},
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each metric create 1 TerraformResource.
// Need Metric Name as ID for terraform resource
func (g *MetricMetadataGenerator) InitResources() error {
	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("metric_metadata") {
			for _, value := range filter.AcceptableValues {
				resources = append(resources, g.createResource(value))
			}
		}
	}

	// Collecting all metrics_metadata can be an expensive task.
	// Hence, only allow collections of metrics passed via filter
	if len(resources) == 0 {
		log.Print("Filter(metric names as IDs) is required for importing datadog_metric_metadata resource")
		return nil
	}
	g.Resources = resources
	return nil
}
