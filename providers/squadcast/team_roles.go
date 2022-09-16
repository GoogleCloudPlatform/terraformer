package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamRolesGenerator struct {
	SquadcastService
}

type TeamRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (g *TeamRolesGenerator) createResources(teamRoles []TeamRole) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, role := range teamRoles {
		resourceList = append(resourceList, terraformutils.NewResource(
			role.ID,
			fmt.Sprintf("team_role_%s", role.Name),
			"squadcast_team_role",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.Args["team_id"].(string),
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resourceList
}

func (g *TeamRolesGenerator) InitResources() error {
	getTeamRolesURL := fmt.Sprintf("/v3/teams/%s/roles", g.Args["team_id"].(string))
	response, err := Request[[]TeamRole](getTeamRolesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
