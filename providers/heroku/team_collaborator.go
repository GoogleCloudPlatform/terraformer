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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	heroku "github.com/heroku/heroku-go/v5"
)

type TeamCollaboratorGenerator struct {
	HerokuService
}

func (g TeamCollaboratorGenerator) createResources(svc *heroku.Service, teamList []heroku.Team) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, team := range teamList {
		apps, err := svc.TeamAppListByTeam(context.TODO(), team.ID, &heroku.ListRange{Field: "id"})
		if err != nil {
			log.Println(err)
		}
		for _, app := range apps {
			collaborators, err := svc.TeamAppCollaboratorList(context.TODO(), app.ID, &heroku.ListRange{Field: "id"})
			if err != nil {
				log.Println(err)
			}
			for _, collaborator := range collaborators {
				resources = append(resources, terraformutils.NewResource(
					collaborator.ID,
					collaborator.ID,
					"heroku_team_collaborator",
					"heroku",
					map[string]string{"app": app.Name},
					[]string{},
					map[string]interface{}{}))
			}
		}
	}
	return resources
}

func (g *TeamCollaboratorGenerator) InitResources() error {
	svc := g.generateService()
	output, err := svc.TeamList(context.TODO(), &heroku.ListRange{Field: "id"})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(svc, output)
	return nil
}
