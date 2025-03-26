// Copyright 2022 The Terraformer Authors.
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

package launchdarkly

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	launchdarkly "github.com/launchdarkly/api-client-go"
)

type ProjectGenerator struct {
	LaunchDarklyService
}

func getProjects(ctx context.Context, client *launchdarkly.APIClient) (launchdarkly.Projects, error) {
	projects, _, err := client.ProjectsApi.GetProjects(ctx)
	return projects, err
}

func (g *ProjectGenerator) loadProjects(ctx context.Context, client *launchdarkly.APIClient) error {
	projects, err := getProjects(ctx, client)
	if err != nil {
		return err
	}
	for _, project := range projects.Items {
		resource := terraformutils.NewResource(
			project.Key,
			project.Key,
			"launchdarkly_project",
			"launchdarkly",
			map[string]string{
				"key": project.Key,
			},
			[]string{},
			map[string]interface{}{})
		resource.IgnoreKeys = append(resource.IgnoreKeys, "include_in_snippet")
		g.Resources = append(g.Resources, resource)
	}
	return nil
}

func (g *ProjectGenerator) InitResources() error {
	if err := g.loadProjects(g.GetArgs()["ctx"].(context.Context), g.GetArgs()["client"].(*launchdarkly.APIClient)); err != nil {
		return err
	}

	return nil
}
