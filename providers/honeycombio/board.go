package honeycombio

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type BoardGenerator struct {
	HoneycombService
}

func (g *BoardGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to initialize Honeycomb client: %v", err)
	}

	boards, err := client.Boards.List(context.TODO())
	if err != nil {
		return fmt.Errorf("unable to list Honeycomb boards: %v", err)
	}

	for _, board := range boards {
		// all of a board's queries must be in our list of target datasets or we don't import it
		onlyValidDatasets := true
		for _, query := range board.Queries {
			if query.Dataset == "" {
				// assume an unset dataset is an environment-wide query
				query.Dataset = environmentWideDatasetSlug
			}
			if _, exists := g.datasets[query.Dataset]; !exists {
				onlyValidDatasets = false
				break
			}
		}

		if onlyValidDatasets {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				board.ID,
				board.ID,
				"honeycombio_board",
				"honeycombio",
				[]string{},
			))
		}
	}

	return nil
}
