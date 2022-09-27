package opal

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ResourceGenerator struct {
	OpalService
}

func (g *ResourceGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to list opal resources: %v", err)
	}

	resources, _, err := client.ResourcesApi.GetResources(context.TODO()).Execute()
	if err != nil {
		return fmt.Errorf("unable to list opal resources: %v", err)
	}

	for {
		for _, resource := range resources.Results {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resource.ResourceId,
				normalizeResourceName(*resource.Name),
				"opal_resource",
				"opal",
				[]string{},
			))
		}

		if !resources.HasNext() || resources.Next.Get() == nil {
			break
		}

		resources, _, err = client.ResourcesApi.GetResources(context.TODO()).Cursor(*resources.Next.Get()).Execute()
		if err != nil {
			return fmt.Errorf("unable to list opal resources: %v", err)
		}
	}

	return nil
}
