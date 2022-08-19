package squadcast

import (
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
	"net/url"
)

type TeamRolesGenerator struct {
	SquadcastService
	teamID string
}

type TeamRole struct {
	ID      string `json:"id" tf:"id"`
	Name    string `json:"name" tf:"name"`
	Slug    string `json:"slug" tf:"-"`
	Default bool   `json:"default" tf:"default"`
}

type TeamRoles []TeamRole

var responseTeamRoles struct {
	Data *[]TeamRole `json:"data"`
}

func (g *TeamRolesGenerator) createResources(teamRoles TeamRoles) []terraformutils.Resource {
	var teamRolesList []terraformutils.Resource
	for _, role := range teamRoles {
		teamRolesList = append(teamRolesList, terraformutils.NewResource(
			role.ID,
			"team_role_"+(role.ID),
			"squadcast_team_role",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.teamID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return teamRolesList
}

func (g *TeamRolesGenerator) InitResources() error {
	if len(g.Args["team_name"].(string)) == 0 {
		log.Fatal("team_name is required")
	}
	team, err := g.generateRequest(fmt.Sprintf("/v3/teams/by-name?name=%s", url.QueryEscape(g.Args["team_name"].(string))))
	if err != nil {
		return err
	}
	err = json.Unmarshal(team, &ResponseTeam)
	if err != nil {
		return err
	}
	g.teamID = ResponseTeam.Data.ID

	body, err := g.generateRequest(fmt.Sprintf("/v3/teams/%s/roles", g.teamID))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &responseTeamRoles)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(*responseTeamRoles.Data)

	return nil
}
