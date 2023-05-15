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

package azuredevops

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type AzureDevOpsProvider struct { //nolint
	terraformutils.Provider
	organizationURL     string
	personalAccessToken string
}

func (p *AzureDevOpsProvider) setEnvConfig() error {

	organizationURL := os.Getenv("AZDO_ORG_SERVICE_URL")
	if organizationURL == "" {
		return errors.New("environment variable AZDO_ORG_SERVICE_URL missing")
	}
	personalAccessToken := os.Getenv("AZDO_PERSONAL_ACCESS_TOKEN")
	if personalAccessToken == "" {
		return errors.New("environment variable AZDO_PERSONAL_ACCESS_TOKEN missing")
	}
	p.organizationURL = organizationURL
	p.personalAccessToken = personalAccessToken
	return nil
}

func (p *AzureDevOpsProvider) Init(args []string) error {
	err := p.setEnvConfig()
	if err != nil {
		return err
	}
	return nil
}

func (p *AzureDevOpsProvider) GetName() string {
	return "azuredevops"
}

func (p *AzureDevOpsProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (p AzureDevOpsProvider) GetResourceConnections() map[string]map[string][]string {
	supported := p.GetSupportedService()
	connections := make(map[string]map[string][]string)
	for serviceName, service := range supported {
		if service2, ok := service.(AzureDevOpsServiceGenerator); ok {
			if conn := service2.GetResourceConnections(); conn != nil {
				connections[serviceName] = conn
			}
		}
	}
	return connections
}

func (p *AzureDevOpsProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"project":        &ProjectGenerator{},
		"group":          &GroupGenerator{},
		"git_repository": &GitRepositoryGenerator{},
	}
}

func (p *AzureDevOpsProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("azuredevops: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"organizationURL":     p.organizationURL,
		"personalAccessToken": p.personalAccessToken,
	})
	return nil
}
