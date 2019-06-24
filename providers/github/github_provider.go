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
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/pkg/errors"
)

type GithubProvider struct {
	terraform_utils.Provider
	organization string
	token        string
}

const githubProviderVersion = "~>2.0.0"

func (p GithubProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p GithubProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"github": map[string]interface{}{
				"version":      githubProviderVersion,
				"organization": p.organization,
			},
		},
	}
}

func (p *GithubProvider) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"organization": p.organization,
		"token":        p.token,
	}
}

// Init GithubProvider with organization
func (p *GithubProvider) Init(args []string) error {
	p.organization = args[0]
	if len(args) < 2 {
		if os.Getenv("GITHUB_TOKEN") == "" {
			return errors.New("token requirement")
		} else {
			p.token = os.Getenv("GITHUB_TOKEN")
		}
	} else {
		p.token = args[1]
	}
	return nil
}

func (p *GithubProvider) GetName() string {
	return "github"
}

func (p *GithubProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"organization": p.organization,
		"token":        p.token,
	})
	return nil
}

// GetSupportedService return map of support service for Github
func (p *GithubProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"members":               &MembersGenerator{},
		"organization_webhooks": &OrganizationWebhooksGenerator{},
		"repositories":          &RepositoriesGenerator{},
		"teams":                 &TeamsGenerator{},
	}
}
