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

type AppConfigAssociationGenerator struct {
	HerokuService
}

func (g AppConfigAssociationGenerator) createResources(svc *heroku.Service, appList []heroku.App) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, app := range appList {
		output, err := svc.ConfigVarInfoForApp(context.TODO(), app.ID)
		if err != nil {
			log.Println(err)
		}
		resources = append(resources, terraformutils.NewResource(
			fmt.Sprintf("%s-config-association", app.Name),
			fmt.Sprintf("%s-config-association", app.Name),
			"heroku_app_config_association",
			"heroku",
			map[string]string{
				"app_id": app.ID,
			},
			[]string{},
			map[string]interface{}{
				// since Heroku does not distinguish between non-sensitive and sensitive config variables, set output to vars
				"vars":           output,
				"sensitive_vars": map[string]interface{}{},
			}))
	}
	return resources
}

func (g *AppConfigAssociationGenerator) InitResources() error {
	svc := g.generateService()
	output, err := svc.AppList(context.TODO(), &heroku.ListRange{Field: "id"})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(svc, output)
	return nil
}
