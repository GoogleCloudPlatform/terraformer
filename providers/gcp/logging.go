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

package gcp

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/iterator"

	"cloud.google.com/go/logging/logadmin"
)

var loggingAllowEmptyValues = []string{}

var loggingAdditionalFields = map[string]interface{}{}

type LoggingGenerator struct {
	GCPService
}

func (g *LoggingGenerator) loadLoggingMetrics(ctx context.Context, client *logadmin.Client) error {
	metricIterator := client.Metrics(ctx)

	for {
		metric, err := metricIterator.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			metric.ID,
			metric.ID,
			"google_logging_metric",
			g.ProviderName,
			map[string]string{
				"name":    metric.ID,
				"project": g.GetArgs()["project"].(string),
			},
			loggingAllowEmptyValues,
			loggingAdditionalFields,
		))
	}
	return nil
}

// Generate TerraformResources from GCP API
func (g *LoggingGenerator) InitResources() error {
	project := g.GetArgs()["project"].(string)
	ctx := context.Background()
	client, err := logadmin.NewClient(ctx, project)
	if err != nil {
		return err
	}

	if err := g.loadLoggingMetrics(ctx, client); err != nil {
		return err
	}

	return nil
}
