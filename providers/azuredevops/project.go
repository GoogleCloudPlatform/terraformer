package azuredevops

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
		fetched := *pages
		items := fetched.Value
		resources = append(resources, items...)
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
	id := *resource.Id
	az.appendSimpleResource(id.String(), *resource.Name, "azuredevops_project")
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
