package gcp

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/artifactregistry/v1"
	"google.golang.org/api/compute/v1"
)

var artifactregistryAllowEmptyValues = []string{""}

type ArtifactregistryGenerator struct {
	GCPService
}

func (g *ArtifactregistryGenerator) InitResources() error {
	ctx := context.Background()

	artifactregistryService, err := artifactregistry.NewService(ctx)
	if err != nil {
		return err
	}

	project := g.GetArgs()["project"].(string)
	location := g.GetArgs()["region"].(compute.Region).Name

	repoListCall := artifactregistryService.Projects.Locations.Repositories.List("projects/" + project + "/locations/" + location)
	if repoList, err := repoListCall.Do(); err == nil {
		for _, repo := range repoList.Repositories {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				repo.Name,
				repo.Name,
				"google_artifact_registry_repository",
				g.GetProviderName(),
				map[string]string{
					"repository_id": repo.Name,
					"project":       project,
					"location":      location,
				},
				artifactregistryAllowEmptyValues,
				map[string]interface{}{
					"format": repo.Format,
				},
			))
		}
	} else {
		return err
	}

	return nil
}
