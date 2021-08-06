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

package gitlab

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/xanzy/go-gitlab"
)

const gitLabDefaultURL = "https://gitlab.com/api/v4/"

type GitLabService struct { //nolint
	terraformutils.Service
}

func (g *GitLabService) createClient() (*gitlab.Client, error) {
	if g.GetArgs()["base_url"].(string) == gitLabDefaultURL {
		return g.createRegularClient()
	}
	return g.createEnterpriseClient()
}

func (g *GitLabService) createRegularClient() (*gitlab.Client, error) {
	return gitlab.NewClient(g.Args["token"].(string))
}

func (g *GitLabService) createEnterpriseClient() (*gitlab.Client, error) {
	return gitlab.NewClient(g.Args["token"].(string), gitlab.WithBaseURL(g.GetArgs()["base_url"].(string)))
}
