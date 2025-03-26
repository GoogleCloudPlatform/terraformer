package azuredevops

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/graph"
)

type GroupGenerator struct {
	AzureDevOpsService
}

func (az *GroupGenerator) listResources() ([]graph.GraphGroup, error) {

	client, fail := az.getGraphClient()
	if fail != nil {
		return nil, fail
	}
	ctx := context.Background()
	var resources []graph.GraphGroup
	pageArgs := graph.ListGroupsArgs{}
	pages, err := client.ListGroups(ctx, pageArgs)
	for ; err == nil; pages, err = client.ListGroups(ctx, pageArgs) {
		resources = append(resources, *pages.GraphGroups...)
		if pages.ContinuationToken == nil {
			return resources, nil
		}
		pageArgs = graph.ListGroupsArgs{
			ContinuationToken: &(*pages.ContinuationToken)[0],
		}
	}
	return nil, err
}

func (az *GroupGenerator) appendResource(resource *graph.GraphGroup) {

	resourceName := firstNonEmpty(resource.DisplayName, resource.MailAddress, resource.OriginId)
	az.appendSimpleResource(*resource.Descriptor, *resourceName, "azuredevops_group")
}

func (az *GroupGenerator) InitResources() error {

	resources, err := az.listResources()
	if err != nil {
		return err
	}
	for _, resource := range resources {
		az.appendResource(&resource)
	}
	return nil
}

func (az *GroupGenerator) GetResourceConnections() map[string][]string {

	return map[string][]string{
		"project": {"scope", "id"},
	}
}
