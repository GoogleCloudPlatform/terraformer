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
	newrelic "github.com/newrelic/newrelic-client-go/newrelic"
)

type AlertChannelGenerator struct {
	NewRelicService
}

func (g *AlertChannelGenerator) createAlertChannelResources(client *newrelic.NewRelic) error {
	alertChannels, err := client.Alerts.ListChannels()
	if err != nil {
		return err
	}

	for _, channel := range alertChannels {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			fmt.Sprintf("%d", channel.ID),
			fmt.Sprintf("%s-%d", normalizeResourceName(channel.Name), channel.ID),
			"newrelic_alert_channel",
			g.ProviderName,
			[]string{},
		))
	}

	return nil
}

func (g *AlertChannelGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	err = g.createAlertChannelResources(client)
	if err != nil {
		return err
	}

	return nil
}
