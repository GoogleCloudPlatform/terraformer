package honeycombio

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type QueryAnnotationGenerator struct {
	HoneycombService
}

func (g *QueryAnnotationGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to initialize Honeycomb client: %v", err)
	}

	boards, err := client.Boards.List(context.TODO())
	if err != nil {
		return err
	}

	for _, board := range boards {
		for _, query := range board.Queries {
			_, datasetSelected := g.datasetMap[query.Dataset]
			if datasetSelected && query.QueryAnnotationID != "" {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					query.QueryAnnotationID,
					query.QueryAnnotationID,
					"honeycombio_query_annotation",
					"honeycombio",
					map[string]string{
						"query_id": query.QueryID,
						"dataset":  query.Dataset,
					},
					[]string{},
					map[string]interface{}{},
				))
			}
		}
	}

	return nil
}
