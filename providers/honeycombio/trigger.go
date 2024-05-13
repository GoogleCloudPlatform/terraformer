package honeycombio

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TriggerGenerator struct {
	HoneycombService
}

func (g *TriggerGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to initialize Honeycomb client: %v", err)
	}

	for _, dataset := range g.datasets {
		if dataset.Slug == environmentWideDatasetSlug {
			// environment-wide Triggers are not supported
			continue
		}
		triggers, err := client.Triggers.List(context.TODO(), dataset.Slug)
		if err != nil {
			return fmt.Errorf("unable to list Honeycomb triggers for dataset %s: %v", dataset.Slug, err)
		}

		for _, trigger := range triggers {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				trigger.ID,
				trigger.ID,
				"honeycombio_trigger",
				"honeycombio",
				map[string]string{
					"dataset": dataset.Name,
				},
				[]string{},
				map[string]interface{}{},
			))
		}
	}

	return nil
}
