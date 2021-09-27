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

package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"
)

type SubnetGenerator struct {
	AzureService
}

func (az *SubnetGenerator) lisSubnets() ([]network.Subnet, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	subnetClient := network.NewSubnetsClient(subscriptionID)
	subnetClient.Authorizer = authorizer
	vnetClient := network.NewVirtualNetworksClient(subscriptionID)
	vnetClient.Authorizer = authorizer
	var (
		vnetIter   network.VirtualNetworkListResultIterator
		subnetIter network.SubnetListResultIterator
		err        error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		vnetIter, err = vnetClient.ListComplete(ctx, resourceGroup)
	} else {
		vnetIter, err = vnetClient.ListAllComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []network.Subnet
	for vnetIter.NotDone() {
		vnet := vnetIter.Value()
		vnetID, err := ParseAzureResourceID(*vnet.ID)
		if err != nil {
			return nil, err
		}
		subnetIter, err = subnetClient.ListComplete(ctx, vnetID.ResourceGroup, *vnet.Name)
		if err != nil {
			return nil, err
		}
		for subnetIter.NotDone() {
			item := subnetIter.Value()
			resources = append(resources, item)
			if err := subnetIter.NextWithContext(ctx); err != nil {
				log.Println(err)
				return resources, err
			}
		}
		if err := vnetIter.NextWithContext(ctx); err != nil {
			log.Println(err)
			return resources, err
		}
	}
	return resources, nil
}

func (az *SubnetGenerator) AppendSubnet(subnet *network.Subnet) {
	az.AppendSimpleResource(*subnet.ID, *subnet.Name, "azurerm_subnet")
}

func (az *SubnetGenerator) appendRouteTable(subnet *network.Subnet) {
	if props := subnet.SubnetPropertiesFormat; props != nil {
		if prop := props.RouteTable; prop != nil {
			az.appendSimpleAssociation(
				*subnet.ID, *subnet.Name, prop.Name,
				"azurerm_subnet_route_table_association",
				map[string]string{
					"subnet_id":      *subnet.ID,
					"route_table_id": *prop.ID,
				})
		}
	}
}

func (az *SubnetGenerator) appendNetworkSecurityGroupAssociation(subnet *network.Subnet) {
	if props := subnet.SubnetPropertiesFormat; props != nil {
		if prop := props.NetworkSecurityGroup; prop != nil {
			az.appendSimpleAssociation(
				*subnet.ID, *subnet.Name, prop.Name,
				"azurerm_subnet_network_security_group_association",
				map[string]string{
					"subnet_id":                 *subnet.ID,
					"network_security_group_id": *prop.ID,
				})
		}
	}
}

func (az *SubnetGenerator) appendNatGateway(subnet *network.Subnet) {
	if props := subnet.SubnetPropertiesFormat; props != nil {
		if prop := props.NatGateway; prop != nil {
			az.appendSimpleAssociation(
				*subnet.ID, *subnet.Name, nil,
				"azurerm_subnet_nat_gateway_association",
				map[string]string{
					"subnet_id":      *subnet.ID,
					"nat_gateway_id": *prop.ID,
				})
		}
	}
}

func (az *SubnetGenerator) appendServiceEndpointPolicies() error {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := network.NewServiceEndpointPoliciesClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator network.ServiceEndpointPolicyListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListByResourceGroupComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListComplete(ctx)
	}
	if err != nil {
		return err
	}

	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_subnet_service_endpoint_storage_policy")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *SubnetGenerator) InitResources() error {

	subnets, err := az.lisSubnets()
	if err != nil {
		return err
	}
	for _, subnet := range subnets {
		az.AppendSubnet(&subnet)
		az.appendRouteTable(&subnet)
		az.appendNetworkSecurityGroupAssociation(&subnet)
		az.appendNatGateway(&subnet)
	}
	if err := az.appendServiceEndpointPolicies(); err != nil {
		return err
	}
	return nil
}

func (az *SubnetGenerator) PostConvertHook() error {
	for _, resource := range az.Resources {
		if resource.InstanceInfo.Type != "azurerm_subnet" {
			continue
		}
		delete(resource.Item, "address_prefix")
	}
	return nil
}
