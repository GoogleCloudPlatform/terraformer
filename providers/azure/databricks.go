package azure

import (
	"context"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/databricks/mgmt/2018-04-01/databricks"
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/authentication"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DatabricksGenerator struct {
	AzureService
}

func (g *DatabricksGenerator) appendResourceAs(resources []terraformutils.Resource, itemID string, itemName string, resourceType string) []terraformutils.Resource {
	prefix := strings.ReplaceAll(resourceType, resourceType, "dbw")
	suffix := strings.ReplaceAll(itemName, "-", "_")
	resourceName := prefix + "_" + suffix
	res := terraformutils.NewSimpleResource(itemID, resourceName, resourceType, g.ProviderName, []string{})
	resources = append(resources, res)
	return resources
}

func (g *DatabricksGenerator) getArgsProperties() (subscriptionID string, authorizer autorest.Authorizer) {
	subs := g.Args["config"].(authentication.Config).SubscriptionID
	auth := g.Args["authorizer"].(autorest.Authorizer)
	return subs, auth
}

func (g *DatabricksGenerator) listWorkspaces() ([]databricks.Workspace, error) {
	subscriptionID, authorizer := g.getArgsProperties()
	client := databricks.NewWorkspacesClient(subscriptionID)
	client.Authorizer = authorizer
	var (
		iterator databricks.WorkspaceListResultIterator
		err      error
	)
	ctx := context.Background()
	if rg := g.Args["resource_group"].(string); rg != "" {
		iterator, err = client.ListByResourceGroupComplete(ctx, rg)
	} else {
		iterator, err = client.ListBySubscriptionComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	var resources []databricks.Workspace
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

func (g *DatabricksGenerator) createDatabricksWorkspaces(workspaces []databricks.Workspace) ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	for _, item := range workspaces {
		resources = g.appendResourceAs(resources, *item.ID, *item.Name, "azurerm_databricks_workspace")
	}
	return resources, nil
}

func (g *DatabricksGenerator) InitResources() error {

	workspaces, err := g.listWorkspaces()
	if err != nil {
		return err
	}

	workspacesFunctions := []func([]databricks.Workspace) ([]terraformutils.Resource, error){
		g.createDatabricksWorkspaces,
	}

	for _, f := range workspacesFunctions {
		resources, ero := f(workspaces)
		if ero != nil {
			return ero
		}
		g.Resources = append(g.Resources, resources...)
	}
	return nil
}
