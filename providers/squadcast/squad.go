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
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (g *SquadGenerator) createResources(squads []Squad) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, squad := range squads {
		resourceList = append(resourceList, terraformutils.NewResource(
			squad.ID,
			fmt.Sprintf("squad_%s", squad.Name),
			"squadcast_squad",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.Args["team_id"].(string),
				"name":    squad.Name,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resourceList
}

func (g *SquadGenerator) InitResources() error {
	if len(g.Args["team_name"].(string)) == 0 {
		return errors.New("--team-name is required")
	}

	getSquadsURL := fmt.Sprintf("/v3/squads?owner_id=%s", g.Args["team_id"].(string))
	response, err := Request[[]Squad](getSquadsURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
