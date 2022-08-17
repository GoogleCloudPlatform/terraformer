package squadcast

import (
	"encoding/json"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamGenerator struct {
	SquadcastService
}

type Team struct {
	ID   string `json:"id" tf:"id"`
	Name string `json:"name" tf:"name"`
}

var responseTeam struct {
	Data *[]Team `json:"data"`
}

type Teams []Team

func (g *TeamGenerator) createResources(teams Teams) []terraformutils.Resource {
	var teamList []terraformutils.Resource
	for _, team := range teams {
		teamList = append(teamList, terraformutils.NewSimpleResource(
			team.ID,
			(team.Name),
			"squadcast_team",
			"squadcast",
			[]string{},
		))
	}
	return teamList
}

func (g *TeamGenerator) InitResources() error {
	body, err := g.generateRequest("/v3/teams")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &responseTeam)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*responseTeam.Data)

	return nil
}
