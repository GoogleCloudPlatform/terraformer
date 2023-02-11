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

package heroku

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	heroku "github.com/heroku/heroku-go/v5"
)

type AddOnGenerator struct {
	HerokuService
}

func (g AddOnGenerator) createResources(addOnList []heroku.AddOn) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, addOn := range addOnList {
		resources = append(resources, terraformutils.NewSimpleResource(
			addOn.ID,
			addOn.Name,
			"heroku_addon",
			"heroku",
			[]string{}))
	}
	return resources
}

func (g *AddOnGenerator) InitResources() error {
	svc := g.generateService()
	ctx := context.Background()

	var output []heroku.AddOn

	if len(g.Filter) > 0 {
		for _, filter := range g.Filter {
			if filter.IsApplicable("app_id") {
				for _, appID := range filter.AcceptableValues {
					appAddons, err := svc.AddOnListByApp(ctx, appID, &heroku.ListRange{Field: "id", Max: 1000})
					if err != nil {
						return fmt.Errorf("Error filtering addons by app, querying for %s: %w", appID, err)
					}
					for _, addOn := range appAddons {
						output = append(output, addOn)
					}
				}
			}
		}
	} else {
		return fmt.Errorf("Heroku Addons must be filtered by app_id: --filter=app_id=<name or ID>")
	}

	g.Resources = g.createResources(output)
	return nil
}
