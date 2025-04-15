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

type DowntimeGenerator struct {
	serviceName string
	MackerelService
}

func (g *DowntimeGenerator) isDowntimeTarget(d *mackerel.Downtime) bool {
	if len(d.ServiceScopes) > 0 {
		for _, svc := range d.ServiceScopes {
			if svc == g.serviceName {
				return true
			}
		}
	}

	if len(d.ServiceExcludeScopes) > 0 {
		for _, svc := range d.ServiceExcludeScopes {
			if svc == g.serviceName {
				return false
			}
		}
	}

	if len(d.RoleScopes) > 0 {
		for _, role := range d.RoleScopes {
			sp := strings.Split(role, ":")
			if sp[0] == g.serviceName {
				return true
			}
		}
	}

	if len(d.RoleExcludeScopes) > 0 {
		for _, role := range d.RoleExcludeScopes {
			sp := strings.Split(role, ":")
			if sp[0] == g.serviceName {
				return false
			}
		}
	}

	return false
}

func (g *DowntimeGenerator) createDowntimeResources(client *mackerel.Client) error {
	downtimes, err := client.FindDowntimes()
	if err != nil {
		return err
	}

	for _, downtime := range downtimes {
		if !g.isDowntimeTarget(downtime) {
			continue
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			downtime.ID,
			fmt.Sprintf("downtime_%s", downtime.Name),
			"mackerel_downtime",
			g.ProviderName,
			map[string]string{
				"name": downtime.Name,
				"memo": downtime.Memo,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return nil
}

// InitResources Generate TerraformResources from Mackerel API,
// from each downtime create 1 TerraformResource.
// Need Downtime ID as ID for terraform resource
func (g *DowntimeGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*mackerel.Client) error{
		g.createDowntimeResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
