package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-03-01/network"
)

type NetworkSecurityGroupGenerator struct {
	AzureService
}

func (az *NetworkSecurityGroupGenerator) listResources() ([]network.SecurityGroup, error) {
	subscriptionID, resourceGroup, authorizer := az.getClientArgs()
	client := network.NewSecurityGroupsClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator network.SecurityGroupListResultIterator
		err      error
	)
	ctx := context.Background()
	if resourceGroup != "" {
		iterator, err = client.ListComplete(ctx, resourceGroup)
	} else {
		iterator, err = client.ListAllComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []network.SecurityGroup
	for iterator.NotDone() {
		item := iterator.Value()
		resources = append(resources, item)
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return resources, err
		}
	}
	return resources, nil
}

func (az *NetworkSecurityGroupGenerator) appendResource(resource *network.SecurityGroup) {
	az.AppendSimpleResource(*resource.ID, *resource.Name, "azurerm_network_security_group")
}

func (az *NetworkSecurityGroupGenerator) appendRules(parent *network.SecurityGroup, resourceGroupID *ResourceID) error {
	subscriptionID, _, authorizer := az.getClientArgs()
	client := network.NewSecurityRulesClient(subscriptionID)
	client.Authorizer = authorizer
	ctx := context.Background()
	iterator, err := client.ListComplete(ctx, resourceGroupID.ResourceGroup, *parent.Name)
	if err != nil {
		return err
	}
	for iterator.NotDone() {
		item := iterator.Value()
		az.AppendSimpleResource(*item.ID, *item.Name, "azurerm_network_security_rule")
		if err := iterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (az *NetworkSecurityGroupGenerator) InitResources() error {

	resources, err := az.listResources()
	if err != nil {
		return err
	}
	for _, resource := range resources {
		az.appendResource(&resource)
		resourceGroupID, err := ParseAzureResourceID(*resource.ID)
		if err != nil {
			return err
		}
		err = az.appendRules(&resource, resourceGroupID)
		if err != nil {
			return err
		}
	}
	return nil
}
