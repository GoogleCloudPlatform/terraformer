package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SquadGenerator struct {
	SquadcastService
	teamID string
}

type Squad struct {
	ID string `json:"id" tf:"id"`
	Name string `json:"name" tf:"name"`
}

var getSquadsResponse struct {
	Data *[]Squad `json:"data"`
}

type Squads []Squad

func (g *SquadGenerator) createResources(squads Squads) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, squad := range squads {
		resources = append(resources, terraformutils.NewResource(
			squad.ID,
			squad.Name,
			"squadcast_squad",
			"squadcast",
			map[string]string{
				"team_id": g.teamID,
				"name": squad.Name,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resources
}

func (g *SquadGenerator) InitResources() error {
	teamName := g.Args["team_name"].(string)
	if(teamName == "") {
		return errors.New("--team-name is required")
	}
	
	escapedTeamName := url.QueryEscape(teamName)
	getTeamURL := fmt.Sprintf("/v3/teams/by-name?name=%s",escapedTeamName)

	team, err := g.generateRequest(getTeamURL)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(team, &getTeamResponse)
	if err != nil {
		return err
	}

	g.teamID = getTeamResponse.Data.ID

	getSquadsURL := fmt.Sprintf("/v3/squads?owner_id=%s", g.teamID);
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
