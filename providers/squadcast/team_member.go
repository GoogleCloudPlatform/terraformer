package squadcast

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamMemberGenerator struct {
	SquadcastService
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
	if len(g.Args["team_name"].(string)) == 0 {
		return errors.New("--team-name is required")
	}

	getTeamURL := fmt.Sprintf("/v3/teams/by-name?name=%s", url.QueryEscape(g.Args["team_name"].(string)))
	response, err := Request[Team](getTeamURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
