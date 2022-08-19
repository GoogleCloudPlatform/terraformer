package honeycombio

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DatasetGenerator struct {
	HoneycombService
}

func (g *DatasetGenerator) InitResources() error {
	_, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to initialize Honeycomb client: %v", err)
	}

	for _, dataset := range g.datasets {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			dataset.Slug,
			dataset.Slug,
			"honeycombio_dataset",
			"honeycombio",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		))
	}

	return nil
}
