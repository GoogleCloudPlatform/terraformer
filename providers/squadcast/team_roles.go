package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
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

var getTeamRolesResponse struct {
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
	getTeamRolesURL := fmt.Sprintf("/v3/teams/%s/roles", g.teamID)

	body, err := g.generateRequest(getTeamRolesURL)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(body, &getTeamRolesResponse)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(*getTeamRolesResponse.Data)

	return nil
}
