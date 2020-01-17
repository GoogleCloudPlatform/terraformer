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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	githubAPI "github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

type UsersGenerator struct {
	GithubService
}

// Generate TerraformResources from Github API,
func (g *UsersGenerator) InitResources() error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.Args["token"].(string)},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := githubAPI.NewClient(tc)

	opt := &githubAPI.UserListOptions{}

	// List all users for the authenticated user
	users, _, err := client.Users.ListAll(ctx, opt)
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, user := range users {
		log.Println("User:", user.GetName())
		opt := &githubAPI.ListOptions{PerPage: 100}

		// List all ssh keys for the user
		for {
			keys, resp, err := client.Users.ListKeys(ctx, user.GetName(), opt)
			if err != nil {
				log.Println(err)
				return nil
			}

			for _, key := range keys {
				log.Println("Key:", key.GetID())
				g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
					strconv.FormatInt(key.GetID(), 10),
					strconv.FormatInt(key.GetID(), 10),
					"github_user_ssh_key",
					"github",
					[]string{},
				))
			}

			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}

	return nil
}
