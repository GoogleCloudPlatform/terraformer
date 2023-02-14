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

type AppFeatureGenerator struct {
	HerokuService
}

func (g AppFeatureGenerator) createResources(appFeatures map[string][]heroku.AppFeature) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for appID, appIDFeatures := range appFeatures {
		for _, appFeature := range appIDFeatures {
			resources = append(resources, terraformutils.NewSimpleResource(
				fmt.Sprintf("%s:%s", appID, appFeature.Name),
				appFeature.Name,
				"heroku_app_feature",
				"heroku",
				[]string{}))
		}
	}
	return resources
}

func (g *AppFeatureGenerator) InitResources() error {
	svc := g.generateService()
	ctx := context.Background()

	output := map[string][]heroku.AppFeature{}
	var hasRequiredFilter bool

	if len(g.Filter) > 0 {
		for _, filter := range g.Filter {
			if filter.IsApplicable("app") {
				hasRequiredFilter = true
				for _, appID := range filter.AcceptableValues {
					appFeatures, err := svc.AppFeatureList(ctx, appID, &heroku.ListRange{Field: "id", Max: 1000})
					if err != nil {
						return fmt.Errorf("Error listing for features for app '%s': %w", appID, err)
					}
					for _, appFeature := range appFeatures {
						if appFeature.Enabled {
							output[appID] = append(output[appID], appFeature)
						}
					}
				}
			}
		}
	}
	if !hasRequiredFilter {
		return fmt.Errorf("Heroku App Features must be filtered by app: --filter=app=<ID>")
	}

	g.Resources = g.createResources(output)
	return nil
}
