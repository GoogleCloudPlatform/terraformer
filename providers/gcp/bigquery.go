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
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"google.golang.org/api/bigquery/v2"
)

var bigQueryAllowEmptyValues = []string{""}

var bigQueryAdditionalFields = map[string]string{}

type BigQueryGenerator struct {
	GCPService
}

// Run on datasetsList and create for each TerraformResource
func (g BigQueryGenerator) createResources(dataSetsList *bigquery.DatasetsListCall, ctx context.Context, bigQueryService *bigquery.Service) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := dataSetsList.Pages(ctx, func(page *bigquery.DatasetList) error {
		for _, dataset := range page.Datasets {
			name := dataset.FriendlyName
			if name == "" {
				name = dataset.Id
			}
			ID := strings.Split(dataset.Id, ":")[1]
			resources = append(resources, terraform_utils.NewResource(
				dataset.Id,
				name,
				"google_bigquery_dataset",
				"google",
				map[string]string{},
				bigQueryAllowEmptyValues,
				bigQueryAdditionalFields,
			))
			resources = append(resources, g.createResourcesTables(ID, ctx, bigQueryService)...)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

func (g *BigQueryGenerator) createResourcesTables(datasetID string, ctx context.Context, bigQueryService *bigquery.Service) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	tableList := bigQueryService.Tables.List(g.Args["project"].(string), datasetID)
	if err := tableList.Pages(ctx, func(page *bigquery.TableList) error {
		for _, table := range page.Tables {
			name := table.FriendlyName
			if name == "" {
				name = table.Id
			}
			resources = append(resources, terraform_utils.NewResource(
				table.Id,
				name,
				"google_bigquery_table",
				"google",
				map[string]string{},
				bigQueryAllowEmptyValues,
				bigQueryAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
func (g *BigQueryGenerator) InitResources() error {
	ctx := context.Background()
	bigQueryService, err := bigquery.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	datasetsList := bigQueryService.Datasets.List(g.GetArgs()["project"].(string))

	g.Resources = g.createResources(datasetsList, ctx, bigQueryService)
	g.PopulateIgnoreKeys()
	return nil
}

// PostGenerateHook for convert schema json as heredoc
func (g *BigQueryGenerator) PostConvertHook() error {
	for i, dataset := range g.Resources {
		if dataset.InstanceInfo.Type != "google_bigquery_dataset" {
			continue
		}
		if val, ok := dataset.Item["default_table_expiration_ms"].(string); ok { // TODO zero int issue
			if val == "0" {
				delete(g.Resources[i].Item, "default_table_expiration_ms")
			}
		}
		for j, table := range g.Resources {
			if table.InstanceInfo.Type != "google_bigquery_table" {
				continue
			}
			if table.InstanceState.Attributes["dataset_id"] == dataset.InstanceState.Attributes["dataset_id"] {
				g.Resources[j].Item["dataset_id"] = "${google_bigquery_dataset." + dataset.ResourceName + ".dataset_id}"
			}
		}
	}

	return nil
}
