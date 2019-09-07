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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	githubAPI "github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

type TeamsGenerator struct {
	GithubService
}

func (g *TeamsGenerator) createTeamsResources(ctx context.Context, teams []*githubAPI.Team, client *githubAPI.Client) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, team := range teams {
		resources = append(resources, terraform_utils.NewResource(
			strconv.FormatInt(team.GetID(), 10),
			team.GetName(),
			"github_team",
			"github",
			map[string]string{},
			[]string{},
			map[string]string{},
		))
		resources = append(resources, g.createTeamMembersResources(ctx, team, client)...)
		resources = append(resources, g.createTeamRepositoriesResources(ctx, team, client)...)
	}
	return resources
}

func (g *TeamsGenerator) createTeamMembersResources(ctx context.Context, team *githubAPI.Team, client *githubAPI.Client) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	members, _, err := client.Teams.ListTeamMembers(ctx, team.GetID(), nil)
	if err != nil {
		log.Println(err)
	}
	for _, member := range members {
		resources = append(resources, terraform_utils.NewResource(
			strconv.FormatInt(team.GetID(), 10)+":"+member.GetLogin(),
			team.GetName()+"_"+member.GetLogin(),
			"github_team_membership",
			"github",
			map[string]string{},
			[]string{},
			map[string]string{},
		))
	}
	return resources
}

func (g *TeamsGenerator) createTeamRepositoriesResources(ctx context.Context, team *githubAPI.Team, client *githubAPI.Client) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	repos, _, err := client.Teams.ListTeamRepos(ctx, team.GetID(), nil)
	if err != nil {
		log.Println(err)
	}
	for _, repo := range repos {
		resources = append(resources, terraform_utils.NewResource(
			strconv.FormatInt(team.GetID(), 10)+":"+repo.GetName(),
			team.GetName()+"_"+repo.GetName(),
			"github_team_repository",
			"github",
			map[string]string{},
			[]string{},
			map[string]string{},
		))
	}
	return resources
}

// Generate TerraformResources from Github API,
func (g *TeamsGenerator) InitResources() error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.Args["token"].(string)},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := githubAPI.NewClient(tc)

	teams, _, err := client.Teams.ListTeams(ctx, g.Args["organization"].(string), nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	g.Resources = g.createTeamsResources(ctx, teams, client)
	g.PopulateIgnoreKeys()
	return nil
}

// PostGenerateHook for connect between team and members
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
