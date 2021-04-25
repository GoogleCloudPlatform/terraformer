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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	heroku "github.com/heroku/heroku-go/v5"
)

type AppFeatureGenerator struct {
	HerokuService
}

func (g AppFeatureGenerator) createResources(svc *heroku.Service, appList []heroku.App) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, app := range appList {
		output, err := svc.AppFeatureList(context.TODO(), app.ID, &heroku.ListRange{Field: "id"})
		if err != nil {
			log.Println(err)
		}
		for _, appFeature := range output {
			resources = append(resources, terraformutils.NewSimpleResource(
				fmt.Sprintf("%s:%s", app.Name, appFeature.ID),
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
	output, err := svc.AppList(context.TODO(), &heroku.ListRange{Field: "id"})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(svc, output)
	return nil
}
