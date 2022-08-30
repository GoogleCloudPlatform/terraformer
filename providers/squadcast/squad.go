package squadcast

import (
	"encoding/json"
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
	body, err := g.generateRequest(getSquadsURL)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &getSquadsResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getSquadsResponse.Data)
	return nil
}
