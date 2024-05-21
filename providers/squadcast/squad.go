package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SquadGenerator struct {
	SCService
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
	req := TRequest{
		URL:             fmt.Sprintf("/v3/squads?owner_id=%s", g.Args["team_id"].(string)),
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
	}
	response, _, err := Request[[]Squad](req)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
