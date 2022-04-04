// UserServiceGenerator
package azuread

import (
	"context"
	"log"

	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

type GroupServiceGenerator struct {
	AzureADService
}

func (az *GroupServiceGenerator) listResources() ([]msgraph.Group, error) {
	client, fail := az.getGroupsClient()
	client.BaseClient.DisableRetries = true

	var resources []msgraph.Group

	if fail != nil {
		return nil, fail
	}
	ctx := context.Background()

	groups, _, err := client.List(ctx, odata.Query{})
	if err != nil {
		return nil, err
	}

	for _, group := range *groups {
		resources = append(resources, group)
	}

	return resources, nil
}

func (az *GroupServiceGenerator) appendResource(resource *msgraph.Group) {
	id := resource.ID
	az.appendSimpleResource(*id, *resource.DisplayName+"-"+*id, "azuread_group")
}

func (az *GroupServiceGenerator) InitResources() error {

	resources, err := az.listResources()
	if err != nil {
		return err
	}
	for _, resource := range resources {
		log.Println(*resource.DisplayName)
		az.appendResource(&resource)
	}
	return nil
}

func (az *GroupServiceGenerator) GetResourceConnections() map[string][]string {

	return map[string][]string{
		"group": {"id"},
	}
}
