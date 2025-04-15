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

type MonitorGenerator struct {
	serviceName string
	MackerelService
}

const (
	monitorTypeConnectivity     = "connectivity"
	monitorTypeHostMetric       = "host"
	monitorTypeServiceMetric    = "service"
	monitorTypeExternalHTTP     = "external"
	monitorTypeExpression       = "expression"
	monitorTypeAnomalyDetection = "anomalyDetection"
)

func (g *MonitorGenerator) isMonitorTarget(serviceName string, scopes, excludeScopes []string) bool {
	if serviceName == g.serviceName {
		return true
	}

	isTarget := false
	for _, scope := range scopes {
		sp := strings.Split(scope, ":")
		if sp[0] == g.serviceName {
			isTarget = true
			continue
		}
	}
	if len(scopes) > 0 && isTarget {
		return true
	}

	isTarget = true
	for _, scope := range excludeScopes {
		sp := strings.Split(scope, ":")
		if sp[0] == g.serviceName {
			isTarget = false
			break
		}
	}
	if len(excludeScopes) > 0 && isTarget {
		return true
	}
	return false
}

func (g *MonitorGenerator) createMonitorResources(client *mackerel.Client) error {
	monitors, err := client.FindMonitors()
	if err != nil {
		return err
	}

	countByMonitorName := map[string]int{}
	for _, monitor := range monitors {
		var mService string
		var scopes, excludeScopes []string
		switch monitor.MonitorType() {
		case monitorTypeConnectivity:
			scopes = monitor.(*mackerel.MonitorConnectivity).Scopes
			excludeScopes = monitor.(*mackerel.MonitorConnectivity).ExcludeScopes
		case monitorTypeHostMetric:
			scopes = monitor.(*mackerel.MonitorHostMetric).Scopes
			excludeScopes = monitor.(*mackerel.MonitorHostMetric).ExcludeScopes
		case monitorTypeServiceMetric:
			mService = monitor.(*mackerel.MonitorServiceMetric).Service
		case monitorTypeExternalHTTP:
			mService = monitor.(*mackerel.MonitorExternalHTTP).Service
		case monitorTypeExpression:
			// nothing to do
		case monitorTypeAnomalyDetection:
			scopes = monitor.(*mackerel.MonitorAnomalyDetection).Scopes
		default:
			return fmt.Errorf("unsupported monitor type: %s(%s)", monitor.MonitorType(), monitor.MonitorName())
		}

		if len(scopes) == 0 && len(excludeScopes) == 0 && len(mService) == 0 {
			// all of scope and excludes and service are empty, we can't detect which service use it.
			continue
		}

		if !g.isMonitorTarget(mService, scopes, excludeScopes) {
			continue
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			monitor.MonitorID(),
			fmt.Sprintf("monitor_%s-%d", monitor.MonitorName(), countByMonitorName[monitor.MonitorName()]),
			"mackerel_monitor",
			g.ProviderName,
			map[string]string{
				"name": monitor.MonitorName(),
			},
			[]string{},
			map[string]interface{}{},
		))
		countByMonitorName[monitor.MonitorName()]++

	}
	return nil
}

// InitResources Generate TerraformResources from Mackerel API,
// from each monitor create 1 TerraformResource.
// Need Monitor ID as ID for terraform resource
func (g *MonitorGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*mackerel.Client) error{
		g.createMonitorResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
