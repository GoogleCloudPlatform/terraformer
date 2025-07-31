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

type NotificationGroupGenerator struct {
	serviceName string
	MackerelService
}

func (g *NotificationGroupGenerator) createNotificationGroupGeneratorResources(client *mackerel.Client) error {
	notificationGroups, err := client.FindNotificationGroups()
	if err != nil {
		return err
	}
	notificationGroupByID := map[string]*mackerel.NotificationGroup{}
	for _, ng := range notificationGroups {
		notificationGroupByID[ng.ID] = ng
	}

	for id, notificationGroup := range notificationGroupByID {
		if len(notificationGroup.Services) == 0 {
			continue
		}

		isTarget := false
		for _, svc := range notificationGroup.Services {
			if svc.Name == g.serviceName {
				isTarget = true
				break
			}
		}
		if !isTarget {
			continue
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			id,
			fmt.Sprintf("notification_group_%s", notificationGroup.Name),
			"mackerel_notification_group",
			g.ProviderName,
			map[string]string{
				"name": notificationGroup.Name,
			},
			[]string{},
			map[string]interface{}{},
		))

		for _, childID := range notificationGroup.ChildNotificationGroupIDs {
			if id == childID {
				child := notificationGroupByID[id]
				g.Resources = append(g.Resources, terraformutils.NewResource(
					child.ID,
					fmt.Sprintf("notification_group_%s", child.Name),
					"mackerel_notification_group",
					g.ProviderName,
					map[string]string{
						"name": child.Name,
					},
					[]string{},
					map[string]interface{}{},
				))
			}
		}
	}
	return nil
}

// InitResources Generate TerraformResources from Mackerel API,
// from each notification group create 1 TerraformResource.
// Need Notification Group ID as ID for terraform resource
func (g *NotificationGroupGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*mackerel.Client) error{
		g.createNotificationGroupGeneratorResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
