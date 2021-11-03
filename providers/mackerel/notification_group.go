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

// NotificationGroupGenerator ...
type NotificationGroupGenerator struct {
	MackerelService
}

func (g *NotificationGroupGenerator) createResources(notificationGroups []*mackerel.NotificationGroup) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, notificationGroup := range notificationGroups {
		resources = append(resources, g.createResource(notificationGroup.ID))
	}
	return resources
}

func (g *NotificationGroupGenerator) createResource(notificationGroupID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		notificationGroupID,
		fmt.Sprintf("notification_group_%s", notificationGroupID),
		"mackerel_notification_group",
		"mackerel",
		[]string{},
	)
}

// InitResources Generate TerraformResources from Mackerel API,
// from each notification group create 1 TerraformResource.
// Need Notification Group ID as ID for terraform resource
func (g *NotificationGroupGenerator) InitResources() error {
	client := g.Args["mackerelClient"].(*mackerel.Client)
	notificationGroups, err := client.FindNotificationGroups()
	if err != nil {
		return err
	}
	g.Resources = append(g.Resources, g.createResources(notificationGroups)...)
	return nil
}
