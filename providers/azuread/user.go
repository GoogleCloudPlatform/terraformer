// UserServiceGenerator
package azuread

import (
	"context"
	"log"

	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

type UserServiceGenerator struct {
	AzureADService
}

func (az *UserServiceGenerator) listResources() ([]msgraph.User, error) {
	client, fail := az.getUserClient()
	client.BaseClient.DisableRetries = true

	var resources []msgraph.User

	if fail != nil {
		return nil, fail
	}
	ctx := context.Background()

	users, _, err := client.List(ctx, odata.Query{})
	if err != nil {
		return nil, err
	}

	for _, user := range *users {
		resources = append(resources, user)
	}

	return resources, nil
}

func (az *UserServiceGenerator) appendResource(resource *msgraph.User) {
	id := resource.ID
	az.appendSimpleResource(*id, *resource.DisplayName, "azuread_user")
}

func (az *UserServiceGenerator) InitResources() error {

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

func (az *UserServiceGenerator) GetResourceConnections() map[string][]string {

	return map[string][]string{
		"user": {"id"},
	}
}
