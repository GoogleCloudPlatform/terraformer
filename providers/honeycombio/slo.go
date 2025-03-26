package honeycombio

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SLOGenerator struct {
	HoneycombService
}

func (g *SLOGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to initialize Honeycomb client: %v", err)
	}

	ctx := context.TODO()

	for _, dataset := range g.datasets {
		if dataset.Slug == environmentWideDatasetSlug {
			// environment-wide SLOs are not supported
			continue
		}
		slos, err := client.SLOs.List(ctx, dataset.Slug)
		if err != nil {
			return fmt.Errorf("unable to list Honeycomb SLOs for dataset %s: %v", dataset.Slug, err)
		}

		for _, slo := range slos {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				slo.ID,
				slo.ID,
				"honeycombio_slo",
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
