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

type OrganizationProjectGenerator struct {
	GithubService
}

// Generate TerraformResources from Github API,
func (g *OrganizationProjectGenerator) InitResources() error {
	ctx := context.Background()
	client, err := g.createClient()
	if err != nil {
		return err
	}

	owner := g.Args["owner"].(string)
	g.Resources = append(g.Resources, createOrganizationProjects(ctx, client, owner)...)

	return nil
}

func createOrganizationProjects(ctx context.Context, client *githubAPI.Client, owner string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}

	opt := &githubAPI.ProjectListOptions{
		ListOptions: githubAPI.ListOptions{PerPage: 100},
	}

	// List all organization projects for the authenticated user
	for {
		projects, resp, err := client.Organizations.ListProjects(ctx, owner, opt)
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
