// Copyright 2019 The Terraformer Authors.
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

package newrelic

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	newrelic "github.com/paultyng/go-newrelic/v4/api"
)

type DashboardGenerator struct {
	NewRelicService
}

func (g *DashboardGenerator) createDashboardResources(client *newrelic.Client) error {
	dashboards, err := client.ListDashboards()
	if err != nil {
		return err
	}

	for _, dashboard := range dashboards {
		resource := terraformutils.NewSimpleResource(
			fmt.Sprintf("%d", dashboard.ID),
			fmt.Sprintf("%s-%d", normalizeResourceName(dashboard.Title), dashboard.ID),
			"newrelic_dashboard",
			g.ProviderName,
			[]string{})
		resource.SlowQueryRequired = true
		g.Resources = append(g.Resources, resource)
	}

	return nil
}

func (g *DashboardGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	err = g.createDashboardResources(client)
	if err != nil {
		return err
	}

	return nil
}

func (g *DashboardGenerator) PostConvertHook() error {
	// Widget's title is a `required` field
	// Setting title to empty string for those resources missing this property
	for i, resource := range g.Resources {
		widgets := resource.Item["widget"]
		if widgets == nil {
			continue
		}
		for wIdx, w := range resource.Item["widget"].([]interface{}) {
			if _, ok := w.(map[string]interface{})["title"]; !ok {
				g.Resources[i].Item["widget"].([]interface{})[wIdx].(map[string]interface{})["title"] = ""
			}
		}
	}

	return nil
}
