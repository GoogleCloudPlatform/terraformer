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

	countByName := make(map[string]int)

	for {
		for _, resource := range resources.Results {
			name := normalizeResourceName(*resource.Name)
			if count, ok := countByName[name]; ok {
				countByName[name] = count + 1
				name = normalizeResourceName(fmt.Sprintf("%s_%d", *resource.Name, count+1))
			} else {
				countByName[name] = 1
			}

			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resource.ResourceId,
				name,
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
