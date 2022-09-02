package squadcast

import (
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamRolesGenerator struct {
	SquadcastService
}

type TeamRole struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Default bool   `json:"default"`
}

func (g *TeamRolesGenerator) createResources(teamRoles []TeamRole) []terraformutils.Resource {
	var teamRolesList []terraformutils.Resource
	for _, role := range teamRoles {
		teamRolesList = append(teamRolesList, terraformutils.NewResource(
			role.ID,
			fmt.Sprintf("team_role_%s", role.ID),
			"squadcast_team_role",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.Args["team_id"].(string),
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return teamRolesList
}

func (g *TeamRolesGenerator) InitResources() error {
	teamID := g.Args["team_id"].(string)
	if teamID == "" {
		return errors.New("--team-name is required")
	}
	getTeamRolesURL := fmt.Sprintf("/v3/teams/%s/roles", teamID)
	response, err := Request[[]TeamRole](getTeamRolesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
