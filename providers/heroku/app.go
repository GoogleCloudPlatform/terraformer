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
	return resources
}

func (g *AppGenerator) InitResources() error {
	svc := g.generateService()
	ctx := context.Background()
	team := g.GetArgs()["team"].(string)

	var output []heroku.App
	var hasRequiredFilter bool

	if len(g.Filter) > 0 {
		for _, filter := range g.Filter {
			if filter.IsApplicable("app") {
				hasRequiredFilter = true
				for _, appID := range filter.AcceptableValues {
					app, err := svc.AppInfo(ctx, appID)
					if err != nil {
						return fmt.Errorf("Error filtering apps by app '%s': %w", appID, err)
					}
					output = append(output, heroku.App{ID: app.ID, Name: app.Name})
				}
			}
		}
	}

	if team != "" {
		hasRequiredFilter = true
		teamApps, err := svc.TeamAppListByTeam(ctx, team, &heroku.ListRange{Field: "id", Max: 1000})
		if err != nil {
			return fmt.Errorf("Error querying apps by team '%s': %w", team, err)
		}
		for _, app := range teamApps {
			output = append(output, heroku.App{ID: app.ID, Name: app.Name})
		}
	}

	if !hasRequiredFilter {
		return fmt.Errorf("Heroku Apps must be scoped by team or filtered by app: --team=<name> or --filter=app=<ID>")
	}

	g.Resources = g.createResources(output)
	return nil
}
