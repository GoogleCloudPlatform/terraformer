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

package pagerduty

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	pagerduty "github.com/heimweh/go-pagerduty/pagerduty"
)

type ScheduleGenerator struct {
	PagerDutyService
}

func (g *ScheduleGenerator) createScheduleResources(client *pagerduty.Client) error {
	var offset = 0
	options := pagerduty.ListSchedulesOptions{}
	for {
		options.Offset = offset
		resp, _, err := client.Schedules.List(&options)
		if err != nil {
			return err
		}

		for _, schedule := range resp.Schedules {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				schedule.ID,
				fmt.Sprintf("schedule_%s", schedule.Name),
				"pagerduty_schedule",
				g.ProviderName,
				[]string{},
			))
		}
		if !resp.More {
			break
		}

		offset += resp.Limit
	}

	return nil
}

func (g *ScheduleGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createScheduleResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
