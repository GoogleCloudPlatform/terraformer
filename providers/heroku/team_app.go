// Copyright 2022 The Terraformer Authors.
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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	heroku "github.com/heroku/heroku-go/v5"
)

type TeamAppGenerator struct {
	HerokuService
}

func (g *TeamAppGenerator) InitResources() error {
	svc := g.generateService()
	ctx := context.Background()

	// ensure HEROKU_TEAM is set
	team := g.Args["team"].(string)
	if team == "" {
		log.Fatalln("missing exported value for HEROKU_TEAM")
		return nil
	}

	// filter if necessary
	if len(g.Filter) > 0 {
		var teamApps []heroku.TeamApp
		for _, filter := range g.Filter {
			if filter.IsApplicable("team_app") {
				for _, app_id := range filter.AcceptableValues {
					app, err := svc.TeamAppInfo(ctx, app_id)
					if err != nil {
						return err
					}

					teamApps = append(teamApps, *app)
				}
			}
		}

		g.Resources = g.createResources(ctx, teamApps)
		return nil
	}

	// otherwise, return all team apps
	allTeamApps, err := svc.TeamAppListByTeam(ctx, team, &heroku.ListRange{Field: "id", Max: 1000})
	if err != nil {
		return err
	}

	g.Resources = g.createResources(ctx, allTeamApps)
	return nil
}

func (g TeamAppGenerator) createResources(ctx context.Context, appList []heroku.TeamApp) []terraformutils.Resource {
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
