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
		return fmt.Errorf("unable to list Honeycomb boards: %v", err)
	}

	boards, err := client.Boards.List(context.TODO())
	if err != nil {
		return err
	}

	for _, board := range boards {
		onlyValidDatasets := true
		for _, query := range board.Queries {
			_, present := g.datasetMap[query.Dataset]
			if !present {
				onlyValidDatasets = false
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
