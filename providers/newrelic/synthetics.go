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

type SyntheticsGenerator struct {
	NewRelicService
}

func (g *SyntheticsGenerator) createSyntheticsMonitorResources(client *newrelic.NewRelic) error {
	allMonitors, err := client.Synthetics.ListMonitors()
	if err != nil {
		return err
	}

	for _, monitor := range allMonitors {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			fmt.Sprint(monitor.ID),
			fmt.Sprintf("%s-%s", normalizeResourceName(monitor.Name), monitor.ID),
			"newrelic_synthetics_monitor",
			g.ProviderName,
			[]string{}))
	}

	return nil
}

func (g *SyntheticsGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	err = g.createSyntheticsMonitorResources(client)
	if err != nil {
		return err
	}

	return nil
}
