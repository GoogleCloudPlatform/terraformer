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

package kibana

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DashboardGenerator struct {
	KibanaService
}

func (g *DashboardGenerator) InitResources() error {
	client := g.generateClient()

	dashboards, err := client.Dashboard().List()
	if err != nil {
		return err
	}
	for _, dashboard := range dashboards {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			dashboard.Id,
			createSlug(dashboard.Attributes.Title),
			"kibana_dashboard",
			"kibana",
			[]string{},
		))
	}
	return nil
}
