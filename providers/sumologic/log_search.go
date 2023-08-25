// Copyright 2023 The Terraformer Authors.
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

package sumologic

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/iancoleman/strcase"
	sumologic "github.com/sumovishal/sumologic-go-sdk/api"
)

type LogSearchGenerator struct {
	SumoLogicService
}

func (g *LogSearchGenerator) createResources(logSearches []sumologic.LogSearch) []terraformutils.Resource {
	resources := make([]terraformutils.Resource, len(logSearches))

	for i, logSearch := range logSearches {
		title := strcase.ToSnake(replaceSpaceAndDash(logSearch.Name))

		resource := terraformutils.NewSimpleResource(
			logSearch.Id,
			fmt.Sprintf("%s-%s", title, logSearch.Id),
			"sumologic_log_search",
			g.ProviderName,
			[]string{})
		resources[i] = resource
	}

	return resources
}

func (g *LogSearchGenerator) InitResources() error {
	client := g.Client()

	var resources []terraformutils.Resource
	var logSearches []sumologic.LogSearch

	req := client.LogSearchesManagementApi.ListLogSearches(g.AuthCtx())
	req = req.Limit(100)

	respBody, _, err := client.LogSearchesManagementApi.ListLogSearchesExecute(req)
	if err != nil {
		return err
	}
	logSearches = respBody.LogSearches
	for respBody.Token != nil {
		req = req.Token(respBody.GetToken())
		respBody, _, err = client.LogSearchesManagementApi.ListLogSearchesExecute(req)
		if err != nil {
			return err
		}
		logSearches = append(logSearches, respBody.LogSearches...)
	}

	resources = g.createResources(logSearches)
	g.Resources = resources
	return nil
}
