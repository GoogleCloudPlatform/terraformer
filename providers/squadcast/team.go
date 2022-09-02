package squadcast

import (
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
	response, err := Request[[]Team](getTeamsURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
