// Copyright 2020 The Terraformer Authors.
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

type OrganizationGenerator struct {
	GithubService
}

// Generate TerraformResources from Github API
func (g *OrganizationGenerator) InitResources() error {
	ctx := context.Background()
	client, err := g.createClient()
	if err != nil {
		return err
	}

	g.Resources = append(g.Resources, g.createMembershipsResources(ctx, client)...)
	g.Resources = append(g.Resources, g.createOrganizationBlocksResources(ctx, client)...)
	g.Resources = append(g.Resources, g.createOrganizationProjects(ctx, client)...)

	return nil
}

func (g *OrganizationGenerator) createOrganizationProjects(ctx context.Context, client *githubAPI.Client) []terraformutils.Resource {
	resources := []terraformutils.Resource{}

	opt := &githubAPI.ProjectListOptions{
		ListOptions: githubAPI.ListOptions{PerPage: 100},
	}

	// List all organization projects for the authenticated user
	for {
		projects, resp, err := client.Organizations.ListProjects(ctx, g.Args["owner"].(string), opt)
		if err != nil {
			log.Println(err)
			return nil
		}

		for _, project := range projects {
			resource := terraformutils.NewSimpleResource(
				strconv.FormatInt(project.GetID(), 10),
				strconv.FormatInt(project.GetID(), 10),
				"github_organization_project",
				"github",
				[]string{},
			)
			resource.SlowQueryRequired = true
			resources = append(resources, resource)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return resources
}

func (g *OrganizationGenerator) createOrganizationBlocksResources(ctx context.Context, client *githubAPI.Client) []terraformutils.Resource {
	resources := []terraformutils.Resource{}

	opt := &githubAPI.ListOptions{PerPage: 100}

	// List all organization blocks for the authenticated user
	for {
		blocks, resp, err := client.Organizations.ListBlockedUsers(ctx, g.Args["owner"].(string), opt)
		if err != nil {
			log.Println(err)
			return nil
		}

		for _, block := range blocks {
			resource := terraformutils.NewSimpleResource(
				block.GetLogin(),
				block.GetLogin(),
				"github_organization_block",
				"github",
				[]string{},
			)
			resource.SlowQueryRequired = true

			resources = append(resources, resource)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return resources
}

func (g *OrganizationGenerator) createMembershipsResources(ctx context.Context, client *githubAPI.Client) []terraformutils.Resource {
	resources := []terraformutils.Resource{}

	opt := &githubAPI.ListMembersOptions{
		ListOptions: githubAPI.ListOptions{PerPage: 100},
	}

	// List all organization members for the authenticated user
	for {
		members, resp, err := client.Organizations.ListMembers(ctx, g.Args["owner"].(string), opt)
		if err != nil {
			log.Println(err)
			return nil
		}

		for _, member := range members {
			resource := terraformutils.NewSimpleResource(
				g.Args["owner"].(string)+":"+member.GetLogin(),
				member.GetLogin(),
				"github_membership",
				"github",
				[]string{},
			)
			resource.SlowQueryRequired = true

			resources = append(resources, resource)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return resources
}
