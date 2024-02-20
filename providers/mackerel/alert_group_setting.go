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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mackerelio/mackerel-client-go"
)

type AlertGroupSettingGenerator struct {
	serviceName string
	MackerelService
}

func (g *AlertGroupSettingGenerator) isAlertGroupSettingTarget(alertGroup *mackerel.AlertGroupSetting) bool {
	for _, svc := range alertGroup.ServiceScopes {
		if svc == g.serviceName {
			return true
		}
	}

	for _, role := range alertGroup.RoleScopes {
		sp := strings.Split(role, ":")
		if sp[0] == g.serviceName {
			return true
		}
	}
	return false
}

func (g *AlertGroupSettingGenerator) createAlertGroupSettingResources(client *mackerel.Client) error {
	alertGroups, err := client.FindAlertGroupSettings()
	if err != nil {
		return err
	}

	for _, alertGroup := range alertGroups {
		if !g.isAlertGroupSettingTarget(alertGroup) {
			continue
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			alertGroup.ID,
			fmt.Sprintf("alert_group_setting_%s", alertGroup.Name),
			"mackerel_alert_group_setting",
			g.ProviderName,
			map[string]string{
				"name": alertGroup.Name,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return nil
}

// InitResources Generate TerraformResources from Mackerel API,
// from each service create 1 TerraformResource.
// Need Service Name as ID for terraform resource
func (g *AlertGroupSettingGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*mackerel.Client) error{
		g.createAlertGroupSettingResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
