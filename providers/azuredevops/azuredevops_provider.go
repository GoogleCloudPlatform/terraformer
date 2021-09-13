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

package azuredevpos

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils/providerwrapper"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
)

type AzureDevOpsProvider struct { //nolint
	terraformutils.Provider
	organizationUrl     string
	personalAccessToken string
	connection          *azuredevops.Connection
}

func (p *AzureDevOpsProvider) setEnvConfig() error {

	organizationUrl := os.Getenv("AZDO_ORG_SERVICE_URL")
	if organizationUrl == "" {
		return errors.New("set AZDO_ORG_SERVICE_URL env var")
	}
	personalAccessToken := os.Getenv("AZDO_PERSONAL_ACCESS_TOKEN")
	if personalAccessToken == "" {
		return errors.New("set AZDO_PERSONAL_ACCESS_TOKEN env var")
	}
	p.organizationUrl = organizationUrl
	p.personalAccessToken = personalAccessToken
	p.connection = azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
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
	return "azuredevpos"
}

func (p *AzureDevOpsProvider) GetProviderData(arg ...string) map[string]interface{} {
	version := providerwrapper.GetProviderVersion(p.GetName())
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"azuredevpos": map[string]interface{}{
				"version":  version,
				"features": map[string]interface{}{},
			},
		},
	}
}

func (AzureDevOpsProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"git_repository": {
			"project": []string{"project_id", "id"},
		},
	}
}

func (p *AzureDevOpsProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"project": &ProjectGenerator{},
	}
}

func (p *AzureDevOpsProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("azuredevpos: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"url":   p.organizationUrl,
		"token": p.personalAccessToken,
	})
	p.Service.(*AzureDevOpsService).connection = p.connection
	return nil
}
