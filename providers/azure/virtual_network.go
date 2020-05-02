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

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-08-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type VirtualNetworkGenerator struct {
	AzureService
}

func (g VirtualNetworkGenerator) createResources(virtualNetworkListResultPage network.VirtualNetworkListResultPage) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for virtualNetworkListResultPage.NotDone() {
		virtualNetworks := virtualNetworkListResultPage.Values()
		for _, virtualNetwork := range virtualNetworks {
			resources = append(resources, terraformutils.NewSimpleResource(
				*virtualNetwork.ID,
				*virtualNetwork.Name,
				"azurerm_virtual_network",
				"azurerm",
				[]string{}))
		}
		if err := virtualNetworkListResultPage.Next(); err != nil {
			log.Println(err)
			break
		}
	}
	return resources
}

func (g *VirtualNetworkGenerator) InitResources() error {
	ctx := context.Background()
	virtualNetworkClient := network.NewVirtualNetworksClient(g.Args["config"].(authentication.Config).SubscriptionID)

	virtualNetworkClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	output, err := virtualNetworkClient.ListAll(ctx)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
