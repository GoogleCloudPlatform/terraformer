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
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/provider_wrapper"
)

type AzureProvider struct {
	terraform_utils.Provider
	subscription string
}

func (p *AzureProvider) Init(args []string) error {
	// Assert ENV Vars for Azure Terraform are set
	if os.Getenv("ARM_CLIENT_ID") == "" {
		return errors.New("set ARM_CLIENT_ID env var")
	}

	if os.Getenv("ARM_CLIENT_SECRET") == "" {
		return errors.New("set ARM_CLIENT_SECRET env var")
	}

	if os.Getenv("ARM_SUBSCRIPTION_ID") == "" {
		return errors.New("set ARM_SUBSCRIPTION_ID env var")
	}
	p.subscription = os.Getenv("ARM_SUBSCRIPTION_ID")

	if os.Getenv("ARM_TENANT_ID") == "" {
		return errors.New("set ARM_TENANT_ID env var")
	}

	// Assert ENV Vars for Azure Go SDK are set
	if os.Getenv("AZURE_CLIENT_ID") == "" {
		return errors.New("set AZURE_CLIENT_ID env var")
	}

	if os.Getenv("AZURE_CLIENT_SECRET") == "" {
		return errors.New("set AZURE_CLIENT_SECRET env var")
	}

	if os.Getenv("AZURE_TENANT_ID") == "" {
		return errors.New("set AZURE_TENANT_ID env var")
	}

	return nil
}

func (p *AzureProvider) GetName() string {
	return "azurerm"
}

func (p *AzureProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"azurerm": map[string]interface{}{
				"version": provider_wrapper.GetProviderVersion(p.GetName()),
			},
		},
	}
}

func (AzureProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *AzureProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"disk":                   &DiskGenerator{},
		"network_interface":      &NetworkInterfaceGenerator{},
		"network_security_group": &NetworkSecurityGroupGenerator{},
		"resource_group":         &ResourceGroupGenerator{},
		"storage_account":        &StorageAccountGenerator{},
		"virtual_machine":        &VirtualMachineGenerator{},
		"virtual_network":        &VirtualNetworkGenerator{},
	}
}

func (p *AzureProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("azurerm: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"subscription": p.subscription,
	})
	return nil
}
