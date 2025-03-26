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

type ExtensionGenerator struct {
	PagerDutyService
}

func (g *ExtensionGenerator) createExtensionResources(client *pagerduty.Client) error {
	var offset = 0
	options := pagerduty.ListExtensionsOptions{}
	for {
		options.Offset = offset
		resp, _, err := client.Extensions.List(&options)
		if err != nil {
			return err
		}

		for _, extension := range resp.Extensions {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				extension.ID,
				fmt.Sprintf("extension_%s_%s", strings.Replace(extension.Name, " ", "_", -1), extension.ID),
				"pagerduty_extension",
				g.ProviderName,
				map[string]string{
					"extension_schema": extension.ExtensionSchema.ID,
				},
				[]string{},
				map[string]interface{}{},
			))
		}

		if !resp.More {
			break
		}
		offset += resp.Limit
	}

	return nil
}

func (g *ExtensionGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*pagerduty.Client) error{
		g.createExtensionResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}
