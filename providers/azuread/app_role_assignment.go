// AppRoleAssignmentServiceGenerator
package azuread

import (
	"context"
	"fmt"
	"log"

	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

type AppRoleAssignmentServiceGenerator struct {
	AzureADService
}

func (az *AppRoleAssignmentServiceGenerator) listResources() ([]msgraph.AppRoleAssignment, error) {
	client, fail := az.getAppRoleAssignmentsClient()
	servicePrincipalsClient, err := az.getServicePrincipalsClient()
	if err != nil {
		return nil, err
	}
	client.BaseClient.DisableRetries = true

	var resources []msgraph.AppRoleAssignment

	if fail != nil {
		return nil, fail
	}
	ctx := context.Background()

	servicePrincipals, _, spErr := servicePrincipalsClient.List(ctx, odata.Query{})
	if spErr != nil {
		return nil, spErr
	}

	for _, sp := range *servicePrincipals {
		appRoleAssignments, _, araErr := client.List(ctx, *sp.ID, odata.Query{})
		if araErr != nil {
			return nil, araErr
		}
		if appRoleAssignments == nil {
			continue
		}
		for _, assignment := range *appRoleAssignments {
			if *assignment.PrincipalType != "ServicePrincipal" {
				continue
			}
			if assignment.Id != nil {
				resources = append(resources, assignment)
			}
		}
	}

	return resources, nil
}

func (az *AppRoleAssignmentServiceGenerator) appendResource(resource *msgraph.AppRoleAssignment) {
	// {objectId}/{type}/{subId}
	id := fmt.Sprintf("%s/appRoleAssignment/%s", *resource.PrincipalId, *resource.Id)
	az.appendSimpleResource(id, *resource.PrincipalDisplayName+"-"+id, "azuread_app_role_assignment")
}

func (az *AppRoleAssignmentServiceGenerator) InitResources() error {

	resources, err := az.listResources()
	if err != nil {
		return err
	}
	for _, resource := range resources {
		log.Println(*resource.PrincipalDisplayName)
		az.appendResource(&resource)
	}
	return nil
}

func (az *AppRoleAssignmentServiceGenerator) GetResourceConnections() map[string][]string {

	return map[string][]string{
		"app_role_assignment": {"id"},
	}
}
