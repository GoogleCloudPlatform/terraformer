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
	az.AppendSimpleResource(*subnet.ID, *subnet.Name, "azurerm_subnet", "snet")
}

func (az *SubnetGenerator) appendRouteTable(subnet *network.Subnet) {
	if props := subnet.SubnetPropertiesFormat; props != nil {
		if prop := props.RouteTable; prop != nil {
			named := *subnet.Name + "_RouteTable"
			az.appendSimpleAssociation(
				*subnet.ID, named,
				"azurerm_subnet_route_table_association", "snetrt",
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
			named := *subnet.Name + "_NetworkSecurityGroup"
			az.appendSimpleAssociation(
				*subnet.ID, named,
				"azurerm_subnet_network_security_group_association", "snetsg",
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
			named := *subnet.Name + "_NatGateway"
			az.appendSimpleAssociation(
				*subnet.ID, named,
				"azurerm_subnet_nat_gateway_association", "snetnat",
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
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_subnet_service_endpoint_storage_policy", "snetpol")
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
