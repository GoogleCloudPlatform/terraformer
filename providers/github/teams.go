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

package github

import (
	"context"
	"log"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	githubAPI "github.com/google/go-github/v35/github"
)

type TeamsGenerator struct {
	GithubService
}

func (g *TeamsGenerator) createTeamsResources(ctx context.Context, teams []*githubAPI.Team, client *githubAPI.Client) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, team := range teams {
		resource := terraformutils.NewSimpleResource(
			strconv.FormatInt(team.GetID(), 10),
			team.GetName(),
			"github_team",
			"github",
			[]string{},
		)
		resource.SlowQueryRequired = true
		resources = append(resources, resource)
		resources = append(resources, g.createTeamMembersResources(ctx, team, client)...)
		resources = append(resources, g.createTeamRepositoriesResources(ctx, team, client)...)
	}
	return resources
}

func (g *TeamsGenerator) createTeamMembersResources(ctx context.Context, team *githubAPI.Team, client *githubAPI.Client) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	members, _, err := client.Teams.ListTeamMembersBySlug(ctx, g.Args["owner"].(string), team.GetSlug(), nil)
	if err != nil {
		log.Println(err)
	}
	for _, member := range members {
		resources = append(resources, terraformutils.NewSimpleResource(
			strconv.FormatInt(team.GetID(), 10)+":"+member.GetLogin(),
			team.GetName()+"_"+member.GetLogin(),
			"github_team_membership",
			"github",
			[]string{},
		))
	}
	return resources
}

func (g *TeamsGenerator) createTeamRepositoriesResources(ctx context.Context, team *githubAPI.Team, client *githubAPI.Client) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	repos, _, err := client.Teams.ListTeamReposBySlug(ctx, g.Args["owner"].(string), team.GetSlug(), nil)
	if err != nil {
		log.Println(err)
	}
	for _, repo := range repos {
		resources = append(resources, terraformutils.NewSimpleResource(
			strconv.FormatInt(team.GetID(), 10)+":"+repo.GetName(),
			team.GetName()+"_"+repo.GetName(),
			"github_team_repository",
			"github",
			[]string{},
		))
	}
	return resources
}

// InitResources generates TerraformResources from Github API,
func (g *TeamsGenerator) InitResources() error {
	ctx := context.Background()
	client, err := g.createClient()
	if err != nil {
		return err
	}

	opt := &githubAPI.ListOptions{PerPage: 1}

	for {
		teams, resp, err := client.Teams.ListTeams(ctx, g.Args["owner"].(string), opt)
		if err != nil {
			log.Println(err)
			return nil
		}

		g.Resources = append(g.Resources, g.createTeamsResources(ctx, teams, client)...)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return nil
}

// PostConvertHook for connect between team and members
func (g *TeamsGenerator) PostConvertHook() error {
	for _, team := range g.Resources {
		if team.InstanceInfo.Type != "github_team" {
			continue
		}
		for i, member := range g.Resources {
			if member.InstanceInfo.Type != "github_team_membership" {
				continue
			}
			if member.InstanceState.Attributes["team_id"] == team.InstanceState.Attributes["id"] {
				g.Resources[i].Item["team_id"] = "${github_team." + team.ResourceName + ".id}"
			}
		}
		for i, repo := range g.Resources {
			if repo.InstanceInfo.Type != "github_team_repository" {
				continue
			}
			if repo.InstanceState.Attributes["team_id"] == team.InstanceState.Attributes["id"] {
				g.Resources[i].Item["team_id"] = "${github_team." + team.ResourceName + ".id}"
			}
		}
	}
	return nil
}
