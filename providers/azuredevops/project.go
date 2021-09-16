package azuredevpos

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

type ProjectGenerator struct {
	AzureDevOpsService
}

func (az *ProjectGenerator) listResources() ([]core.TeamProjectReference, error) {
	client, fail := az.getCoreClient()
	if fail != nil {
		return nil, fail
	}
	ctx := context.Background()
	var resources []core.TeamProjectReference
	pageArgs := core.GetProjectsArgs{}
	pages, err := client.GetProjects(ctx, pageArgs)
	for ; err == nil; pages, err = client.GetProjects(ctx, pageArgs) {
		resources = append(resources, (*pages).Value...)
		if pages.ContinuationToken == "" {
			return resources, nil
		}
		pageArgs = core.GetProjectsArgs{
			ContinuationToken: &pages.ContinuationToken,
		}
	}
	return nil, err
}

func (az *ProjectGenerator) appendResource(resource *core.TeamProjectReference) {
	az.appendSimpleResource((*resource.Id).String(), *resource.Name, "azuredevops_project")
}

func (az *ProjectGenerator) InitResources() error {

	resources, err := az.listResources()
	if err != nil {
		return err
	}
	for _, resource := range resources {
		az.appendResource(&resource)
	}
	return nil
}
