package squadcast

import (
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"net/url"
)

type TeamMemberGenerator struct {
	SquadcastService
}

var responseTeam struct {
	Data *Team `json:"data"`
}

func (g *TeamMemberGenerator) createResources(team Team) []terraformutils.Resource {
	var teamMemberList []terraformutils.Resource
	for _, member := range team.Members {
		teamMemberList = append(teamMemberList, terraformutils.NewSimpleResource(
			member.UserID,
			"squadcast_team_member_"+(member.UserID),
			"squadcast_team",
			"squadcast",
			[]string{},
		))
	}
	return teamMemberList
}

func (g *TeamMemberGenerator) InitResources() error {
	body, err := g.generateRequest(fmt.Sprintf("/v3/teams/by-name?name=%s", url.QueryEscape(g.Args["team_name"].(string))))
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &responseTeam)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*responseTeam.Data)

	return nil
}

//type TeamMemberGenerator struct {
//	SquadcastService
//	team Team
//}
//
//type TeamMember struct {
//	ID      string   `json:"id" tf:"id"`
//	UserID  string   `json:"user_id" tf:"user_id"`
//	RoleIDs []string `json:"role_ids" tf:"role_ids"`
//}
//
//type TeamMembers []TeamMember
//
//func (g *TeamMemberGenerator) filterTeam(teams Teams) {
//	for _, team := range teams {
//		if team.Name == g.Args["team_name"] {
//			g.team = team
//			return
//		}
//	}
//
//	log.Fatal("team not found")
//}
//
//func (g *TeamMemberGenerator) createResources() []terraformutils.Resource {
//	var teamMemberList []terraformutils.Resource
//	for _, member := range g.team.Members {
//		teamMemberList = append(teamMemberList, terraformutils.NewSimpleResource(
//			member.UserID,
//			"team_member_"+(member.UserID),
//			"squadcast_team_member",
//			"squadcast",
//			[]string{},
//		))
//	}
//	return teamMemberList
//}
//
//func (g *TeamMemberGenerator) InitResources() error {
//	body, err := g.generateRequest("/v3/teams")
//	if err != nil {
//		return err
//	}
//
//	err = json.Unmarshal(body, &ResponseTeam)
//	if err != nil {
//		return err
//	}
//
//	g.filterTeam(*ResponseTeam.Data)
//	g.Resources = g.createResources()
//
//	return nil
//}
