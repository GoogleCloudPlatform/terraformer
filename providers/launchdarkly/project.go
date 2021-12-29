// Copyright 2021 The Terraformer Authors.
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
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	launchdarkly "github.com/launchdarkly/api-client-go"
)

type ProjectGenerator struct {
	LaunchDarklyService
}

func (g *ProjectGenerator) loadProjects(ctx context.Context, client *launchdarkly.APIClient) error {
	projects, _, err := client.ProjectsApi.GetProjects(ctx)
	if err != nil {
		return err
	}

	for _, project := range projects.Items {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			project.Key,
			project.Key,
			"launchdarkly_project",
			"launchdarkly",
			[]string{}))
	}
	return nil
}

func (g *ProjectGenerator) InitResources() error {
	apiKey := os.Getenv("LAUNCHDARKLY_ACCESS_TOKEN")
	cfg := &launchdarkly.Configuration{
		BasePath:      basePath,
		DefaultHeader: make(map[string]string),
		UserAgent:     fmt.Sprintf("launchdarkly-terraform-provider/%s", version),
	}
	cfg.AddDefaultHeader("LD-API-Version", APIVersion)

	client := launchdarkly.NewAPIClient(cfg)

	ctx := context.WithValue(context.Background(), launchdarkly.ContextAPIKey, launchdarkly.APIKey{
		Key: apiKey,
	})

	if err := g.loadProjects(ctx, client); err != nil {
		return err
	}

	return nil
}
