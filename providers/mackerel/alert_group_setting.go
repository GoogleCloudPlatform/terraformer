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

package mackerel

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mackerelio/mackerel-client-go"
)

// AlertGroupSettingGenerator ...
type AlertGroupSettingGenerator struct {
	MackerelService
}

func (g *AlertGroupSettingGenerator) createResources(alertGroupSettings []*mackerel.AlertGroupSetting) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, alertGroupSetting := range alertGroupSettings {
		resources = append(resources, g.createResource(alertGroupSetting.ID))
	}
	return resources
}

func (g *AlertGroupSettingGenerator) createResource(alertGroupSettingID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		alertGroupSettingID,
		fmt.Sprintf("alert_group_setting_%s", alertGroupSettingID),
		"mackerel_alert_group_setting",
		"mackerel",
		[]string{},
	)
}

// InitResources Generate TerraformResources from Mackerel API,
// from each alert group setting create 1 TerraformResource.
// Need Alert Group Setting ID as ID for terraform resource
func (g *AlertGroupSettingGenerator) InitResources() error {
	client := g.Args["mackerelClient"].(*mackerel.Client)
	alertGroupSettings, err := client.FindAlertGroupSettings()
	if err != nil {
		return err
	}
	g.Resources = append(g.Resources, g.createResources(alertGroupSettings)...)
	return nil
}
