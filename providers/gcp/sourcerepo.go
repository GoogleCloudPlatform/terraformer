package gcp

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/sourcerepo/v1"
)

var sourcerepoAllowEmptyValues = []string{""}

var sourcerepoAdditionalFields = map[string]interface{}{}

type SourceRepoGenerator struct {
	GCPService
}

func (g *SourceRepoGenerator) InitResources() error {
	ctx := context.Background()
	sourcerepoService, err := sourcerepo.NewService(ctx)
	if err != nil {
		return err
	}
	project := g.GetArgs()["project"].(string)
	listCall := sourcerepoService.Projects.Repos.List("projects/" + project)

	if repos, err := listCall.Do(); err == nil {
		for _, repo := range repos.Repos {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				repo.Name,
				repo.Name,
				"google_sourcerepo_repository",
				g.GetProviderName(),
				map[string]string{
					"name":    repo.Name,
					"project": project,
				},
				sourcerepoAllowEmptyValues,
				sourcerepoAdditionalFields,
			))
		}
	} else {
		return err
	}

	return nil
}
