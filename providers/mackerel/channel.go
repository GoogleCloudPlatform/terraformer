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

// ChannelGenerator ...
type ChannelGenerator struct {
	MackerelService
}

func (g *ChannelGenerator) createResources(channels []*mackerel.Channel) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, channel := range channels {
		if channel.Type != "email" && channel.Type != "slack" && channel.Type != "webhook" {
			continue
		}

		if channel.Type == "email" {
			if channel.Events != nil {
				events := *channel.Events
				for _, event := range events {
					if event != "alert" && event != "alertGroup" {
						continue
					}
				}
			} else {
				continue
			}
		}

		resources = append(resources, g.createResource(channel.ID))
	}
	return resources
}

func (g *ChannelGenerator) createResource(channelID string) terraformutils.Resource {
	return terraformutils.NewSimpleResource(
		channelID,
		fmt.Sprintf("channel_%s", channelID),
		"mackerel_channel",
		"mackerel",
		[]string{},
	)
}

// InitResources Generate TerraformResources from Mackerel API,
// from each channel create 1 TerraformResource.
// Need Channel ID as ID for terraform resource
func (g *ChannelGenerator) InitResources() error {
	client := g.Args["mackerelClient"].(*mackerel.Client)
	channels, err := client.FindChannels()
	if err != nil {
		return err
	}
	g.Resources = append(g.Resources, g.createResources(channels)...)
	return nil
}
