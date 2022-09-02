package squadcast

import (
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SquadGenerator struct {
	SquadcastService
}

type Squad struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

var getSquadsResponse struct {
	Data *[]Squad `json:"data"`
}

func (g *SquadGenerator) createResources(squads []Squad) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, squad := range squads {
		resources = append(resources, terraformutils.NewResource(
			squad.ID,
			squad.Name,
			"squadcast_squad",
			"squadcast",
			map[string]string{
				"team_id": g.Args["team_id"].(string),
				"name": squad.Name,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}

func (g *SquadGenerator) InitResources() error {
	teamID := g.Args["team_id"].(string)
	if teamID == "" {
		return errors.New("--team-name is required")
	}
	getSquadsURL := fmt.Sprintf("/v3/squads?owner_id=%s", teamID);

	response, err := Request[[]Squad](getSquadsURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(*response)
	return nil
}
