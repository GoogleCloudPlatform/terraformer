// ApplicationServiceGenerator
package azuread

import (
	"context"
	"log"

	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

type ApplicationServiceGenerator struct {
	AzureADService
}

func (az *ApplicationServiceGenerator) listResources() ([]msgraph.Application, error) {
	client, fail := az.getApplicationsClient()
	client.BaseClient.DisableRetries = true

	var resources []msgraph.Application

	if fail != nil {
		return nil, fail
	}
	ctx := context.Background()

	applications, _, err := client.List(ctx, odata.Query{})
	if err != nil {
		return nil, err
	}

	for _, application := range *applications {
		resources = append(resources, application)
	}

	return resources, nil
}

func (az *ApplicationServiceGenerator) appendResource(resource *msgraph.Application) {
	id := resource.ID
	az.appendSimpleResource(*id, *resource.DisplayName, "azuread_application")
}

func (az *ApplicationServiceGenerator) InitResources() error {

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

func (az *ApplicationServiceGenerator) GetResourceConnections() map[string][]string {

	return map[string][]string{
		"application": {"id"},
	}
}
