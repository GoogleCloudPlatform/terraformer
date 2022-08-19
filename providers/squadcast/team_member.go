package squadcast

import (
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"net/url"
)

type TeamMemberGenerator struct {
	SquadcastService
	teamID string
}

var ResponseTeam struct {
	Data *Team `json:"data"`
}

func (g *TeamMemberGenerator) createResources(team Team) []terraformutils.Resource {
	var teamMemberList []terraformutils.Resource
	g.teamID = team.ID
	for _, member := range team.Members {
		teamMemberList = append(teamMemberList, terraformutils.NewResource(
			member.UserID,
			"squadcast_team_member_"+(member.UserID),
			"squadcast_team_member",
			g.ProviderName,
			map[string]string{
				"team_id": g.teamID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return teamMemberList
}

func (g *TeamMemberGenerator) InitResources() error {
	body, err := g.generateRequest(fmt.Sprintf("/v3/teams/by-name?name=%s", url.QueryEscape(g.Args["team_name"].(string))))
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ResponseTeam)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*ResponseTeam.Data)

	return nil
}
