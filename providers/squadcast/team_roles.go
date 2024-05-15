package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamRolesGenerator struct {
	SCService
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
	req := TRequest{
		URL:             fmt.Sprintf("/v3/teams/%s/roles", g.Args["team_id"].(string)),
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
	}
	response, _, err := Request[[]TeamRole](req)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
