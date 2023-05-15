// Copyright 2019 The Terraformer Authors.
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

package azuread

import (
	"errors"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type AzureADProvider struct { //nolint
	terraformutils.Provider
	tenantID     string
	clientID     string
	clientSecret string
}

func (p *AzureADProvider) setEnvConfig() error {

	tenantID := os.Getenv("ARM_TENANT_ID")
	if tenantID == "" {
		return errors.New("please set ARM_TENANT_ID in your environment")
	}
	clientID := os.Getenv("ARM_CLIENT_ID")
	if clientID == "" {
		return errors.New("please set ARM_CLIENT_ID in your environment")
	}
	clientSecret := os.Getenv("ARM_CLIENT_SECRET")
	if clientSecret == "" {
		return errors.New("please set ARM_CLIENT_SECRET in your environment")
	}
	p.tenantID = tenantID
	p.clientID = clientID
	p.clientSecret = clientSecret
	return nil
}

func (p *AzureADProvider) Init(args []string) error {
	err := p.setEnvConfig()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (p *AzureADProvider) GetName() string {
	return "azuread"
}

func (p *AzureADProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (AzureADProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *AzureADProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"user":                &UserServiceGenerator{},
		"application":         &ApplicationServiceGenerator{},
		"group":               &GroupServiceGenerator{},
		"service_principal":   &ServicePrincipalServiceGenerator{},
		"app_role_assignment": &AppRoleAssignmentServiceGenerator{},
	}
}

func (p *AzureADProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("azuread: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"tenant_id":     p.tenantID,
		"client_id":     p.clientID,
		"client_secret": p.clientSecret,
	})
	return nil
}
