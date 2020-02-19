package octopusdeploy

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

type TagSetGenerator struct {
	OctopusDeployService
}

func (g *TagSetGenerator) InitResources() error {
	client, err := g.Client()
	if err != nil {
		return err
	}

	funcs := []func(*octopusdeploy.Client) error{
		g.createTagSetResources,
	}

	for _, f := range funcs {
		err := f(client)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *TagSetGenerator) createTagSetResources(client *octopusdeploy.Client) error {
	tagSets, err := client.TagSet.GetAll()
	if err != nil {
		return err
	}

	for _, tagSet := range *tagSets {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			fmt.Sprintf("%s", tagSet.ID),
			fmt.Sprintf("%s", tagSet.Name),
			"octopusdeploy_tag_set",
			g.ProviderName,
			[]string{},
		))
	}

	return nil
}
