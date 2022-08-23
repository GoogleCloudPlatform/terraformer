package squadcast

import (
	"encoding/json"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamGenerator struct {
	SquadcastService
}

type Team struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Default     bool              `json:"default"`
	Members     []*DataTeamMember `json:"members"`
	Roles       []*TeamRole       `json:"roles"`
}

type DataTeamMember struct {
	UserID  string   `json:"user_id"`
	RoleIDs []string `json:"role_ids"`
}

var getTeamsResponse struct {
	Data *[]Team `json:"data"`
}


func (g *TeamGenerator) createResources(teams []Team) []terraformutils.Resource {
	var teamList []terraformutils.Resource
	for _, team := range teams {
		teamList = append(teamList, terraformutils.NewSimpleResource(
			team.ID,
			team.Name,
			"squadcast_team",
			"squadcast",
			[]string{},
		))
	}
	return teamList
}

func (g *TeamGenerator) InitResources() error {
	getTeamsURL := "/v3/teams"
	body, err := g.generateRequest(getTeamsURL)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &getTeamsResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getTeamsResponse.Data)

	return nil
}
