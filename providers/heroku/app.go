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

type AppGenerator struct {
	HerokuService
}

func (g AppGenerator) createResources(appList []heroku.App) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, app := range appList {
		resources = append(resources, terraformutils.NewSimpleResource(
			app.ID,
			app.Name,
			"heroku_app",
			"heroku",
			[]string{}))
	}
	fmt.Printf("created app resources %+v", resources)
	return resources
}

func (g *AppGenerator) InitResources() error {
	svc := g.generateService()
	ctx := context.Background()

	var output []heroku.App

	if len(g.Filter) > 0 {
		for _, filter := range g.Filter {
			if filter.IsApplicable("app_id") {
				for _, appID := range filter.AcceptableValues {
					app, err := svc.AppInfo(ctx, appID)
					if err != nil {
						return fmt.Errorf("Error filtering apps by app, querying for %s: %w", appID, err)
					}
					output = append(output, heroku.App{ID: app.ID, Name: app.Name})
				}
			} else if filter.IsApplicable("team") {
				for _, team := range filter.AcceptableValues {
					teamApps, err := svc.TeamAppListByTeam(ctx, team, &heroku.ListRange{Field: "id", Max: 1000})
					if err != nil {
						return fmt.Errorf("Error filtering apps by team, querying for %s: %w", team, err)
					}
					for _, app := range teamApps {
						output = append(output, heroku.App{ID: app.ID, Name: app.Name})
					}
				}
			}
		}
	} else {
		return fmt.Errorf("Heroku Apps must be filtered by app_id or team: --filter=team=<name>")
	}

	g.Resources = g.createResources(output)
	return nil
}
