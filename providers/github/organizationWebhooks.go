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

type OrganizationWebhooksGenerator struct {
	GithubService
}

// Generate TerraformResources from Github API,
func (g *OrganizationWebhooksGenerator) InitResources() error {
	ctx := context.Background()
	client, err := g.createClient()
	if err != nil {
		return err
	}

	opt := &githubAPI.ListOptions{PerPage: 100}

	// List all organization hooks for the authenticated user
	for {
		hooks, resp, err := client.Organizations.ListHooks(ctx, g.Args["owner"].(string), opt)
		if err != nil {
			log.Println(err)
			return nil
		}

		for _, hook := range hooks {
			resource := terraformutils.NewSimpleResource(
				strconv.FormatInt(hook.GetID(), 10),
				strconv.FormatInt(hook.GetID(), 10),
				"github_organization_webhook",
				"github",
				[]string{},
			)
			resource.SlowQueryRequired = true
			g.Resources = append(g.Resources, resource)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return nil
}
