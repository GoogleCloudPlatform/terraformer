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

package azure

import (
	"errors"
	"fmt"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils/providerwrapper"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
)

type AzureProvider struct { //nolint
	terraformutils.Provider
	config     authentication.Config
	authorizer autorest.Authorizer
}

func (p *AzureProvider) setEnvConfig() error {
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		return errors.New("set ARM_SUBSCRIPTION_ID env var")
	}
	builder := &authentication.Builder{
		ClientID:                 os.Getenv("ARM_CLIENT_ID"),
		SubscriptionID:           subscriptionID,
		TenantID:                 os.Getenv("ARM_TENANT_ID"),
		Environment:              os.Getenv("ARM_ENVIRONMENT"),
		ClientSecret:             os.Getenv("ARM_CLIENT_SECRET"),
		SupportsAzureCliToken:    true,
		SupportsClientSecretAuth: true,
		SupportsClientCertAuth:   true,
		ClientCertPath:           os.Getenv("ARM_CLIENT_CERTIFICATE_PATH"),
		ClientCertPassword:       os.Getenv("ARM_CLIENT_CERTIFICATE_PASSWORD"),

		/*
		   // Managed Service Identity Auth
		   SupportsManagedServiceIdentity bool
		   MsiEndpoint                    string
		*/
	}

	if builder.Environment == "" {
		builder.Environment = "public"
	}
	config, err := builder.Build()
	if err != nil {
		return nil
	}
	p.config = *config

	return nil
}

func (p *AzureProvider) getAuthorizer() (autorest.Authorizer, error) {
	env, err := authentication.DetermineEnvironment(p.config.Environment)
	if err != nil {
		return nil, err
	}

	oauthConfig, err := p.config.BuildOAuthConfig(env.ActiveDirectoryEndpoint)
	if err != nil {
		return nil, err
	}

	if oauthConfig == nil {
		return nil, fmt.Errorf("Unable to configure OAuthConfig for tenant %s", p.config.TenantID)
	}

	sender := sender.BuildSender("AzureRM")

	auth, err := p.config.GetAuthorizationToken(sender, oauthConfig, env.ResourceManagerEndpoint)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (p *AzureProvider) Init(args []string) error {
	err := p.setEnvConfig()
	if err != nil {
		return err
	}

	authorizer, err := p.getAuthorizer()
	if err != nil {
		return err
	}
	p.authorizer = authorizer

	return nil
}

func (p *AzureProvider) GetName() string {
	return "azurerm"
}

func (p *AzureProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"azurerm": map[string]interface{}{
				"version": providerwrapper.GetProviderVersion(p.GetName()),
			},
		},
	}
}

func (AzureProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *AzureProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"app_service":                          &AppServiceGenerator{},
		"analysis":                             &AnalysisGenerator{},
		"database":                             &DatabasesGenerator{},
		"disk":                                 &DiskGenerator{},
		"network_interface":                    &NetworkInterfaceGenerator{},
		"network_security_group":               &NetworkSecurityGroupGenerator{},
		"resource_group":                       &ResourceGroupGenerator{},
		"security_center_contact":              &SecurityCenterContactGenerator{},
		"security_center_subscription_pricing": &SecurityCenterSubscriptionPricingGenerator{},
		"storage_account":                      &StorageAccountGenerator{},
		"virtual_machine":                      &VirtualMachineGenerator{},
		"virtual_network":                      &VirtualNetworkGenerator{},
	}
}

func (p *AzureProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("azurerm: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"config":     p.config,
		"authorizer": p.authorizer,
	})
	return nil
}
