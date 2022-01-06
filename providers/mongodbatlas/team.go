// Copyright 2018 The Terraformer Authors.
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

package mongodbatlas

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"go.mongodb.org/atlas/mongodbatlas"
)

var (
	TeamAllowEmptyValues = []string{}
)

type TeamGenerator struct {
	MongoDBAtlasService
}

func (g TeamGenerator) createResources(teams []mongodbatlas.Team, orgID string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, team := range teams {
		resources = append(resources, terraformutils.NewResource(
			orgID+"-"+team.ID,
			orgID+"_"+team.ID,
			"mongodbatlas_teams",
			"mongodbatlas",
			map[string]string{
				"name":   team.Name,
				"org_id": orgID,
			},
			TeamAllowEmptyValues,
			map[string]interface{}{}))
	}
	return resources
}

func (g *TeamGenerator) InitResources() error {
	client := g.generateClient()
	orgID := g.Args["org_id"].(string)
	list := []mongodbatlas.Team{}
	opt := &mongodbatlas.ListOptions{}
	// Enable filtering by team
	for {
		teams, resp, err := client.Teams.List(context.TODO(), orgID, opt)
		if err != nil {
			return err
		}
		list = append(list, teams...)
		if resp.Links == nil || resp.IsLastPage() {
			break
		}
		page, err := resp.CurrentPage()
		if err != nil {
			return err
		}
		opt.PageNum = page + 1
	}
	g.Resources = g.createResources(list, orgID)
	return nil
}
