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
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-03-01/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type VirtualMachineGenerator struct {
	AzureService
}

func (g VirtualMachineGenerator) createResources(virtualMachineListResultPage compute.VirtualMachineListResultPage) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for virtualMachineListResultPage.NotDone() {
		vms := virtualMachineListResultPage.Values()
		for _, vm := range vms {
			var newResource terraformutils.Resource
			if vm.VirtualMachineProperties.OsProfile.WindowsConfiguration != nil {
				newResource = terraformutils.NewSimpleResource(
					*vm.ID,
					*vm.Name,
					"azurerm_windows_virtual_machine",
					"azurerm",
					[]string{})
			} else {
				newResource = terraformutils.NewSimpleResource(
					*vm.ID,
					*vm.Name,
					"azurerm_linux_virtual_machine",
					"azurerm",
					[]string{})
			}

			resources = append(resources, newResource)
		}
		if err := virtualMachineListResultPage.Next(); err != nil {
			log.Println(err)
			break
		}
	}
	return resources
}

func (g *VirtualMachineGenerator) InitResources() error {
	ctx := context.Background()
	vmClient := compute.NewVirtualMachinesClient(g.Args["config"].(authentication.Config).SubscriptionID)

	vmClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	output, err := vmClient.ListAll(ctx)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
