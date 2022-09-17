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
	"time"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/iancoleman/strcase"
	sumologic "github.com/sumovishal/sumologic-go-sdk/api"
)

type MonitorGenerator struct {
	SumoLogicService
}

func (g *MonitorGenerator) createResources(monitors []sumologic.MonitorsLibraryBaseResponse) []terraformutils.Resource {
	resources := make([]terraformutils.Resource, len(monitors))

	for i, monitor := range monitors {
		name := strcase.ToSnake(replaceSpaceAndDash(monitor.Name))
		resource := terraformutils.NewSimpleResource(
			monitor.Id,
			fmt.Sprintf("%s-%s", name, monitor.Id),
			"sumologic_monitor",
			g.ProviderName,
			[]string{},
		)
		resources[i] = resource
	}

	return resources
}

func (g *MonitorGenerator) InitResources() error {
	client := g.Client()

	searchReq := client.MonitorsLibraryManagementApi.MonitorsSearch(
		g.AuthCtx()).Query(fmt.Sprintf("createdAfter:%d", time.Now().Unix()))
	rspBody, _, err := client.MonitorsLibraryManagementApi.MonitorsSearchExecute(searchReq)
	if err != nil {
		return err
	}

	var monitors []sumologic.MonitorsLibraryBaseResponse
	for _, itemWithPath := range rspBody {
		monitors = append(monitors, itemWithPath.Item)
	}

	resources := g.createResources(monitors)
	g.Resources = resources
	return nil
}
