package squadcast

import (
	"fmt"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamMemberGenerator struct {
	SCService
}

type TeamMember struct {
	UserID string `json:"user_id"`
}

func (g *TeamMemberGenerator) createResources(team Team) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, member := range team.Members {
		resourceList = append(resourceList, terraformutils.NewResource(
			member.UserID,
			fmt.Sprintf("squadcast_team_member_%s", member.UserID),
			"squadcast_team_member",
			g.GetProviderName(),
			map[string]string{
				"team_id": team.ID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resourceList
}

func (g *TeamMemberGenerator) InitResources() error {
	req := TRequest{
		URL:             fmt.Sprintf("/v3/teams/by-name?name=%s", url.QueryEscape(g.Args["team_name"].(string))),
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
	}

	response, _, err := Request[Team](req)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
