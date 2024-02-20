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

type ChannelGenerator struct {
	serviceName string
	MackerelService
}

func (g *ChannelGenerator) createChannelResources(client *mackerel.Client) error {
	channels, err := client.FindChannels()
	if err != nil {
		return err
	}

	channelByID := map[string]*mackerel.Channel{}
	for _, c := range channels {
		channelByID[c.ID] = c
	}

	targetChannelIDs := map[string]bool{}
	notificationGroups, err := client.FindNotificationGroups()
	if err != nil {
		return err
	}

	for _, notificationGroup := range notificationGroups {
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
		for _, chanID := range notificationGroup.ChildChannelIDs {
			targetChannelIDs[chanID] = true
		}
	}

	for id := range targetChannelIDs {
		channel, ok := channelByID[id]
		if !ok {
			// error?
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			channel.ID,
			fmt.Sprintf("channel_%s", channel.Name),
			"mackerel_channel",
			g.ProviderName,
			map[string]string{
				"name": channel.Name,
			},
			[]string{},
			map[string]interface{}{},
		))

	}

	return nil
}

// InitResources Generate TerraformResources from Mackerel API,
// from each channel create 1 TerraformResource.
// Need Channel ID as ID for terraform resource
func (g *ChannelGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*mackerel.Client) error{
		g.createChannelResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
