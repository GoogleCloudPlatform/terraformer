package azuredevpos

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
)

type GitRepositoryGenerator struct {
	AzureDevOpsService
}

func (az *GitRepositoryGenerator) listResources() ([]git.GitRepository, error) {

	client, err := az.getGitClient()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	projects, err := client.GetRepositories(ctx, git.GetRepositoriesArgs{})
	if err != nil {
		return nil, err
	}
	return *projects, nil
}

func (az *GitRepositoryGenerator) appendResource(project *git.GitRepository) {

	az.appendSimpleResource((*project.Id).String(), *project.Name, "azuredevops_git_repository")
}

func (az *GitRepositoryGenerator) InitResources() error {

	projects, err := az.listResources()
	if err != nil {
		return err
	}
	for _, project := range projects {
		az.appendResource(&project)
	}
	return nil
}

func (az *GitRepositoryGenerator) GetResourceConnections() map[string][]string {

	return map[string][]string{
		"project": {"project_id", "id"},
	}
}
