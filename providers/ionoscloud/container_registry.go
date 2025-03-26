package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ContainerRegistryGenerator struct {
	Service
}

func (g *ContainerRegistryGenerator) InitResources() error {
	client := g.generateClient()
	containerRegistryAPIClient := client.ContainerRegistryAPIClient
	resourceType := "ionoscloud_container_registry"

	response, _, err := containerRegistryAPIClient.RegistriesApi.RegistriesGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if response.Items == nil {
		log.Printf("[WARNING] expected a response containing registries but received 'nil' instead.")
		return nil
	}
	registries := *response.Items
	for _, registry := range registries {
		if registry.Properties == nil || registry.Properties.Name == nil {
			log.Printf("[WARNING] 'nil' values in the response for the registry with ID %v, skipping this resource.", *registry.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*registry.Id,
			*registry.Properties.Name+"-"+*registry.Id,
			resourceType,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
