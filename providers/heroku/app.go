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

	for _, app := range output {
		appFeatures, err := g.createAppFeatureResources(ctx, svc, app)
		if err != nil {
			return fmt.Errorf("Error creating app feature resources: %w", err)
		}
		g.Resources = append(g.Resources, appFeatures...)

		addons, err := g.createAddonResources(ctx, svc, app.ID)
		if err != nil {
			return fmt.Errorf("Error creating app addon resources: %w", err)
		}
		g.Resources = append(g.Resources, addons...)

		addonAttachments, err := g.createAddonAttachmentResources(ctx, svc, app.ID)
		if err != nil {
			return fmt.Errorf("Error creating app addon attachment resources: %w", err)
		}
		g.Resources = append(g.Resources, addonAttachments...)
	}

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

func (g AppGenerator) createAppFeatureResources(ctx context.Context, svc *heroku.Service, app heroku.App) ([]terraformutils.Resource, error) {
	list := []heroku.AppFeature{}

	appFeatures, err := svc.AppFeatureList(ctx, app.ID, &heroku.ListRange{Field: "id", Max: 1000})
	if err != nil {
		return []terraformutils.Resource{}, fmt.Errorf("Error listing for features for app '%s': %w", app.ID, err)
	}
	for _, appFeature := range appFeatures {
		if appFeature.Enabled {
			list = append(list, appFeature)
		}
	}
	var resources []terraformutils.Resource
	for _, appFeature := range list {
		resources = append(resources, terraformutils.NewSimpleResource(
			fmt.Sprintf("%s:%s", app.ID, appFeature.Name),
			fmt.Sprintf("%s-%s", app.Name, appFeature.Name),
			"heroku_app_feature",
			"heroku",
			[]string{}))
	}
	return resources, nil
}

func (g AppGenerator) createAddonResources(ctx context.Context, svc *heroku.Service, appID string) ([]terraformutils.Resource, error) {
	list := []heroku.AddOn{}

	appAddons, err := svc.AddOnListByApp(ctx, appID, &heroku.ListRange{Field: "id", Max: 1000})
	if err != nil {
		return []terraformutils.Resource{}, fmt.Errorf("Error listing addons by app '%s': %w", appID, err)
	}
	for _, addOn := range appAddons {
		list = append(list, addOn)
	}
	var resources []terraformutils.Resource
	for _, addOn := range list {
		resources = append(resources, terraformutils.NewSimpleResource(
			addOn.ID,
			addOn.Name,
			"heroku_addon",
			"heroku",
			[]string{}))
	}
	return resources, nil
}

func (g AppGenerator) createAddonAttachmentResources(ctx context.Context, svc *heroku.Service, appID string) ([]terraformutils.Resource, error) {
	list := []heroku.AddOnAttachment{}

	appAddons, err := svc.AddOnListByApp(ctx, appID, &heroku.ListRange{Field: "id", Max: 1000})
	if err != nil {
		return []terraformutils.Resource{}, fmt.Errorf("Error listing addons by app '%s': %w", appID, err)
	}
	for _, addOn := range appAddons {
		addonAttachments, err := svc.AddOnAttachmentListByAddOn(ctx, addOn.ID, &heroku.ListRange{Field: "id", Max: 1000})
		if err != nil {
			return []terraformutils.Resource{}, fmt.Errorf("Error listing addon attachments by addon '%s': %w", addOn.Name, err)
		}
		for _, attachment := range addonAttachments {
			list = append(list, attachment)
		}
	}
	var resources []terraformutils.Resource
	for _, addOnAttachment := range list {
		resources = append(resources, terraformutils.NewSimpleResource(
			addOnAttachment.ID,
			fmt.Sprintf("%s-%s", addOnAttachment.App.Name, addOnAttachment.Name),
			"heroku_addon_attachment",
			"heroku",
			[]string{}))
	}
	return resources, nil
}
