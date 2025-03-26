package opsgenie

import (
	"context"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk-v2/team"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamGenerator struct {
	OpsgenieService
}

func (g *TeamGenerator) InitResources() error {
	client, err := g.TeamClient()
	if err != nil {
		return err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	result, err := client.List(ctx, &team.ListTeamRequest{})
	if err != nil {
		return err
	}

	g.Resources = g.createResources(result.Teams)
	return nil
}

func (g *TeamGenerator) createResources(teams []team.ListedTeams) []terraformutils.Resource {
	var resources []terraformutils.Resource

	for _, t := range teams {
		resources = append(resources, terraformutils.NewResource(
			t.Id,
			t.Name,
			"opsgenie_team",
			g.ProviderName,
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}
