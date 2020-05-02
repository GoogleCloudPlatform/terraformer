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

type TeamMemberGenerator struct {
	HerokuService
}

func (g TeamMemberGenerator) createResources(svc *heroku.Service, teamList []heroku.Team) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, team := range teamList {
		output, err := svc.TeamMemberList(context.TODO(), team.ID, &heroku.ListRange{Field: "id"})
		if err != nil {
			log.Println(err)
		}
		for _, member := range output {
			resources = append(resources, terraformutils.NewSimpleResource(
				fmt.Sprintf("%s:%s", team.ID, member.Email),
				member.ID,
				"heroku_team_member",
				"heroku",
				[]string{}))
		}
	}
	return resources
}

func (g *TeamMemberGenerator) InitResources() error {
	svc := g.generateService()
	output, err := svc.TeamList(context.TODO(), &heroku.ListRange{Field: "id"})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(svc, output)
	return nil
}
