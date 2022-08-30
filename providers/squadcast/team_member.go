package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamMemberGenerator struct {
	SquadcastService
}

var getTeamResponse struct {
	Data *Team `json:"data"`
}

func (g *TeamMemberGenerator) createResources(team Team) []terraformutils.Resource {
	var teamMemberList []terraformutils.Resource
	for _, member := range team.Members {
		teamMemberList = append(teamMemberList, terraformutils.NewResource(
			member.UserID,
			"squadcast_team_member_"+(member.UserID),
			"squadcast_team_member",
			g.ProviderName,
			map[string]string{
				"team_id": team.ID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return teamMemberList
}

func (g *TeamMemberGenerator) InitResources() error {
	teamName := g.Args["team_name"].(string)
	if len(teamName) == 0 {
		return errors.New("--team-name is required")
	}

	escapedTeamName := url.QueryEscape(teamName)
	getTeamURL := fmt.Sprintf("/v3/teams/by-name?name=%s", escapedTeamName)
	
	body, err := g.generateRequest(getTeamURL)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &getTeamResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getTeamResponse.Data)

	return nil
}
