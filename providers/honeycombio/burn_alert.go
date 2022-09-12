package honeycombio

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type BurnAlertGenerator struct {
	HoneycombService
}

func (g *BurnAlertGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to initialize Honeycomb client: %v", err)
	}

	ctx := context.TODO()

	for _, dataset := range g.datasets {
		slos, err := client.SLOs.List(ctx, dataset.Slug)
		if err != nil {
			return fmt.Errorf("unable to list Honeycomb SLOs for dataset %s: %v", dataset.Slug, err)
		}

		for _, slo := range slos {
			bas, _ := client.BurnAlerts.ListForSLO(ctx, dataset.Slug, slo.ID)
			for _, ba := range bas {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					ba.ID,
					ba.ID,
					"honeycombio_burn_alert",
					"honeycombio",
					map[string]string{
						"dataset": dataset.Name,
						"slo_id":  slo.ID,
					},
					[]string{"recipient"},
					map[string]interface{}{},
				))
			}
		}
	}

	return nil
}
