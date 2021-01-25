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

	datadogV2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

var (
	// LogsArchiveAllowEmptyValues ...
	LogsArchiveAllowEmptyValues = []string{"path", "query"}
)

// LogsArchiveGenerator ...
type LogsArchiveGenerator struct {
	DatadogService
}

func (g *LogsArchiveGenerator) createResources(logsArchives []datadogV2.LogsArchiveDefinition) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, logsArchive := range logsArchives {
		logsArchiveID := logsArchive.GetId()
		resources = append(resources, g.createResource(logsArchiveID))
	}

	return resources
}

func (g *LogsArchiveGenerator) createResource(logsArchiveID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		logsArchiveID,
		fmt.Sprintf("logs_archive_%s", logsArchiveID),
		"datadog_logs_archive",
		"datadog",
		LogsArchiveAllowEmptyValues,
	)
}

// InitResources Generate TerraformResources from Datadog API,
// from each archive create 1 TerraformResource.
// Need LogsArchive ID as ID for terraform resource
func (g *LogsArchiveGenerator) InitResources() error {
	datadogClientV2 := g.Args["datadogClientV2"].(*datadogV2.APIClient)
	authV2 := g.Args["authV2"].(context.Context)

	resources := []terraformutils.Resource{}
	for _, filter := range g.Filter {
		if filter.FieldPath == "id" && filter.IsApplicable("logs_archive") {
			for _, value := range filter.AcceptableValues {
				resp, _, err := datadogClientV2.LogsArchivesApi.GetLogsArchive(authV2, value).Execute()
				if err != nil {
					return err
				}
				logsArchiveData := resp.GetData()
				resources = append(resources, g.createResource(logsArchiveData.GetId()))
			}
		}
	}

	if len(resources) > 0 {
		g.Resources = resources
		return nil
	}

	logsArchiveListResp, _, err := datadogClientV2.LogsArchivesApi.ListLogsArchives(authV2).Execute()
	logsArchiveList := logsArchiveListResp.GetData()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(logsArchiveList)
	return nil
}
