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
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	pagerduty "github.com/heimweh/go-pagerduty/pagerduty"

	"fmt"
	"strings"
)

type AddonGenerator struct {
	PagerDutyService
}

func (g *AddonGenerator) createAddonResources(client *pagerduty.Client) error {
	resp, _, err := client.Addons.List(&pagerduty.ListAddonsOptions{})
	if err != nil {
		return err
	}

	for _, addon := range resp.Addons {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			addon.ID,
			fmt.Sprintf("addon_%s_%s", strings.Replace(addon.Name, " ", "_", -1), addon.ID),
			"pagerduty_addon",
			g.ProviderName,
			[]string{},
		))
	}

	return nil
}

func (g *AddonGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createAddonResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
