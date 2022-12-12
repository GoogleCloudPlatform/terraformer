package honeycombio

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ColumnGenerator struct {
	HoneycombService
}

func (g *ColumnGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to initialize Honeycomb client: %v", err)
	}

	for _, dataset := range g.datasets {
		if dataset.Slug == environmentWideDatasetSlug {
			continue
		}
		columns, err := client.Columns.List(context.TODO(), dataset.Slug)
		if err != nil {
			return fmt.Errorf("unable to list Honeycomb columns for dataset %s: %v", dataset.Slug, err)
		}

		for _, column := range columns {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				column.ID,
				fmt.Sprintf("%s_%s", dataset.Name, column.KeyName),
				"honeycombio_column",
				"honeycombio",
				map[string]string{
					"dataset":  dataset.Name,
					"key_name": column.KeyName,
				},
				[]string{"hidden", "type"},
				map[string]interface{}{},
			))
		}
	}

	return nil
}
