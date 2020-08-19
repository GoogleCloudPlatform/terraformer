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
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-08-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/hashicorp/go-azure-helpers/authentication"
)

type SubnetGenerator struct {
	AzureService
}

func (g SubnetGenerator) createResources(ctx context.Context, iter network.SubnetListResultIterator) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	for iter.NotDone() {
		subnet := iter.Value()
		subnetID, err := ParseAzureResourceID(*subnet.ID)
		if err != nil {
			return nil, err
		}
		resources = append(resources, terraformutils.NewSimpleResource(
			*subnet.ID,
			fmt.Sprintf("%s-%s-%s", subnetID.ResourceGroup, subnetID.Path["virtualNetworks"], *subnet.Name),
			"azurerm_subnet",
			"azurerm",
			[]string{}))
		if err := iter.NextWithContext(ctx); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func (g *SubnetGenerator) PostConvertHook() error {
	for _, resource := range g.Resources {
		if resource.InstanceInfo.Type != "azurerm_subnet" {
			continue
		}
		delete(resource.Item, "address_prefix")
	}
	return nil
}

func (g *SubnetGenerator) InitResources() error {
	ctx := context.Background()
	client := network.NewSubnetsClient(g.Args["config"].(authentication.Config).SubscriptionID)
	client.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	vnetClient := network.NewVirtualNetworksClient(g.Args["config"].(authentication.Config).SubscriptionID)
	vnetClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	vnets, err := vnetClient.ListAllComplete(ctx)
	if err != nil {
		return err
	}

	var subnetResources []terraformutils.Resource
	for vnets.NotDone() {
		vnet := vnets.Value()

		vnetID, err := ParseAzureResourceID(*vnet.ID)
		if err != nil {
			return err
		}
		subnetIter, err := client.ListComplete(ctx, vnetID.ResourceGroup, *vnet.Name)
		if err != nil {
			return err
		}
		subnets, err := g.createResources(ctx, subnetIter)
		if err != nil {
			return err
		}
		subnetResources = append(subnetResources, subnets...)

		if err := vnets.NextWithContext(ctx); err != nil {
			return err
		}
	}
	g.Resources = subnetResources
	return nil
}
