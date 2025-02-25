// Copyright 2021 The Terraformer Authors.
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
	sumologic "github.com/SumoLogic/sumologic-go-sdk"
	"github.com/iancoleman/strcase"
)

type MetricsSearchV2Generator struct {
	SumoLogicService
}

func (g *MetricsSearchV2Generator) createResources(searches []sumologic.MetricsSearchResponse) []terraformutils.Resource {
	resources := make([]terraformutils.Resource, len(searches))

	for i, search := range searches {
		title := strcase.ToSnake(replaceSpaceAndDash(search.Title))
		resource := terraformutils.NewSimpleResource(
			*search.Id,
			fmt.Sprintf("%s-%s", title, *search.Id),
			"sumologic_metrics_search_v2",
			g.ProviderName,
			[]string{})
		resources[i] = resource
	}

	return resources
}

func (g *MetricsSearchV2Generator) InitResources() error {
	client := g.Client()
	req := client.MetricsSearchesManagementV2API.ListMetricsSearches(g.AuthCtx())
	req = req.Limit(100).Mode("allViewableByUser")

	respBody, _, err := client.MetricsSearchesManagementV2API.ListMetricsSearchesExecute(req)
	if err != nil {
		return err
	}
	searches := respBody.MetricsSearches
	for respBody.Next != nil {
		req = req.Token(respBody.GetNext())
		respBody, _, err = client.MetricsSearchesManagementV2API.ListMetricsSearchesExecute(req)
		if err != nil {
			return err
		}
		searches = append(searches, respBody.MetricsSearches...)
	}

	resources := g.createResources(searches)
	g.Resources = resources
	return nil
}
