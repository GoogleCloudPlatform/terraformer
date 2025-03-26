// ServicePrincipalServiceGenerator
package azuread

import (
	"context"
	"log"

	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

type ServicePrincipalServiceGenerator struct {
	AzureADService
}

func (az *ServicePrincipalServiceGenerator) listResources() ([]msgraph.ServicePrincipal, error) {
	client, fail := az.getServicePrincipalsClient()
	client.BaseClient.DisableRetries = true

	var resources []msgraph.ServicePrincipal

	if fail != nil {
		return nil, fail
	}
	ctx := context.Background()

	servicePrincipal, _, err := client.List(ctx, odata.Query{})
	if err != nil {
		return nil, err
	}

	for _, sp := range *servicePrincipal {
		resources = append(resources, sp)
	}

	return resources, nil
}

func (az *ServicePrincipalServiceGenerator) appendResource(resource *msgraph.ServicePrincipal) {
	id := resource.ID
	az.appendSimpleResource(*id, *resource.DisplayName+"-"+*id, "azuread_service_principal")
}

func (az *ServicePrincipalServiceGenerator) InitResources() error {

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

func (az *ServicePrincipalServiceGenerator) GetResourceConnections() map[string][]string {

	return map[string][]string{
		"servicePrincipal": {"id"},
	}
}
