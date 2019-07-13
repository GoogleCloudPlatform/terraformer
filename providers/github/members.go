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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	githubAPI "github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

type MembersGenerator struct {
	GithubService
}

// Generate TerraformResources from Github API,
func (g *MembersGenerator) InitResources() error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.Args["token"].(string)},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := githubAPI.NewClient(tc)

	members, _, err := client.Organizations.ListMembers(ctx, g.Args["organization"].(string), nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	for _, member := range members {
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			g.Args["organization"].(string)+":"+member.GetLogin(),
			member.GetLogin(),
			"github_membership",
			"github",
			map[string]string{},
			[]string{},
			map[string]string{},
		))
	}
	g.PopulateIgnoreKeys()
	return nil
}
