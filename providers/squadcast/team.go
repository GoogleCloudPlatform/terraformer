package squadcast

import (
	"encoding/json"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamGenerator struct {
	SquadcastService
}

type Team struct {
	ID          string            `json:"id" tf:"id"`
	Name        string            `json:"name" tf:"name"`
	Description string            `json:"description" tf:"description"`
	Default     bool              `json:"default" tf:"default"`
	Members     []*DataTeamMember `json:"members" tf:"-"`
	Roles       []*TeamRole       `json:"roles" tf:"-"`
}

type DataTeamMember struct {
	UserID  string   `json:"user_id" tf:"user_id"`
	RoleIDs []string `json:"role_ids" tf:"role_ids"`
}

var responseTeams struct {
	Data *[]Team `json:"data"`
}

type Teams []Team

func (g *TeamGenerator) createResources(teams Teams) []terraformutils.Resource {
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
	body, err := g.generateRequest("/v3/teams")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &responseTeams)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*responseTeams.Data)

	return nil
}
