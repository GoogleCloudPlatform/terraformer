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

func (g AppGenerator) createResources(appList []heroku.App) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	var resourcesEmpty []terraformutils.Resource

	for _, app := range appList {
		configVars, err := g.getSettableConfigVars(app.ID)
		if err != nil {
			return resourcesEmpty, fmt.Errorf("Error in getSettableConfigVars for '%s': %w", app.ID, err)
		}
		resources = append(resources, terraformutils.NewResource(
			app.ID,
			app.Name,
			"heroku_app",
			"heroku",
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"config_vars": configVars,
			}))
	}
	return resources, nil
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

	resources, err := g.createResources(output)
	if err != nil {
		return fmt.Errorf("Error creating app resources: %w", err)
	}
	g.Resources = resources

	return nil
}

func (g AppGenerator) getSettableConfigVars(appID string) (map[string]string, error) {
	svc := g.generateService()
	ctx := context.Background()
	output := map[string]string{}
	emptyOutput := map[string]string{}

	vars, err := svc.ConfigVarInfoForApp(ctx, appID)
	if err != nil {
		return emptyOutput, fmt.Errorf("Error querying ConfigVarInfoForApp '%s': %w", appID, err)
	}

	for k, v := range vars {
		if v != nil {
			output[k] = *v
		}
	}

	appAddons, err := svc.AddOnListByApp(ctx, appID, &heroku.ListRange{Field: "id", Max: 1000})
	if err != nil {
		return emptyOutput, fmt.Errorf("Error querying AddOnListByApp '%s': %w", appID, err)
	}
	for _, addOn := range appAddons {
		for _, addOnConfigVar := range addOn.ConfigVars {
			delete(output, addOnConfigVar)
		}
	}

	return output, nil
}
