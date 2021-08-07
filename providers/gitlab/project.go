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

package gitlab

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
	"strconv"
	"strings"

	gitLabAPI "github.com/xanzy/go-gitlab"
)

type ProjectGenerator struct {
	GitLabService
}

// Generate TerraformResources from gitlab API,
func (g *ProjectGenerator) InitResources() error {
	ctx := context.Background()
	client, err := g.createClient()
	if err != nil {
		return err
	}

	group := g.Args["group"].(string)
	g.Resources = append(g.Resources, createProjects(ctx, client, group)...)

	return nil
}

// PostConvertHook for add policy json as heredoc
func (g *ProjectGenerator) PostConvertHook() error {
	for i, resource := range g.Resources {
		if resource.InstanceInfo.Type == "gitlab_project_variable" {
			if val, ok := g.Resources[i].Item["value"]; ok {
				g.Resources[i].Item["value"] = fmt.Sprintf(`<<PROJECTVARIABLE
%s
PROJECTVARIABLE`, val.(string))
			}
		}
	}
	return nil
}

func createProjects(ctx context.Context, client *gitLabAPI.Client, group string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	opt := &gitLabAPI.ListGroupProjectsOptions{
		ListOptions: gitLabAPI.ListOptions{
			PerPage: 100,
		},
	}

	// List all organization projects for the authenticated user
	for {
		projects, resp, err := client.Groups.ListGroupProjects(group, opt, gitLabAPI.WithContext(ctx))
		if err != nil {
			log.Println(err)
			return nil
		}

		for _, project := range projects {
			resource := terraformutils.NewSimpleResource(
				strconv.FormatInt(int64(project.ID), 10),
				getProjectResourceName(project),
				"gitlab_project",
				"gitlab",
				[]string{},
			)

			//mirror fields from API doesn't match with the ones from terraform provider
			resource.IgnoreKeys = []string{"mirror_trigger_builds", "only_mirror_protected_branches", "mirror", "mirror_overwrites_diverged_branches"}

			resource.SlowQueryRequired = true
			resources = append(resources, resource)
			resources = append(resources, createProjectVariables(ctx, client, project)...)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return resources
}
func createProjectVariables(ctx context.Context, client *gitLabAPI.Client, project *gitLabAPI.Project) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	opt := &gitLabAPI.ListProjectVariablesOptions{}

	// List all organization projects for the authenticated user
	for {
		projectVariables, resp, err := client.ProjectVariables.ListVariables(project.ID, opt, gitLabAPI.WithContext(ctx))
		if err != nil {
			log.Println(err)
			return nil
		}

		for _, projectVariable := range projectVariables {

			resource := terraformutils.NewSimpleResource(
				fmt.Sprintf("%d:%s:%s", project.ID, projectVariable.Key, projectVariable.EnvironmentScope),
				fmt.Sprintf("%s___%s___%s", getProjectResourceName(project), projectVariable.Key, projectVariable.EnvironmentScope),
				"gitlab_project_variable",
				"gitlab",
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

func getProjectResourceName(project *gitLabAPI.Project) string {
	return fmt.Sprintf("%d___%s", project.ID, strings.ReplaceAll(project.PathWithNamespace, "/", "__"))
}
