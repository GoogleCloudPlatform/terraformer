// service resource is yet to be implemented

package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SLOGenerator struct {
	SquadcastService
	teamID string
}

type SLO struct {
	ID   int `json:"id"`
	Name string `json:"name"`
}

var getSLOsResponse struct {
	Data struct {
		SLOs *[]SLO `json:"slos"`
	} `json:"data"`
}

func (g *SLOGenerator) createResources(slo []SLO) []terraformutils.Resource {
	var SLOList []terraformutils.Resource
	for _, s := range slo {
		SLOList = append(SLOList, terraformutils.NewResource(
			fmt.Sprintf("%d", s.ID),
			"slo_"+(s.Name),
			"squadcast_slo",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.teamID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return SLOList
}

func (g *SLOGenerator) InitResources() error {
	teamName := g.Args["team_name"].(string)
	if len(teamName) == 0 {
		return errors.New("--team-name is required")
	}

	escapedTeamName := url.QueryEscape(teamName)
	getTeamURL := fmt.Sprintf("/v3/teams/by-name?name=%s", escapedTeamName)
	team, err := g.generateRequest(getTeamURL)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(team, &getTeamResponse)
	if err != nil {
		return err
	}
	g.teamID = getTeamResponse.Data.ID

	getSLOsURL := "/v3/slo?owner_id=" + g.teamID
	body, err := g.generateRequest(getSLOsURL)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(body, &getSLOsResponse)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(*getSLOsResponse.Data.SLOs)
	return nil
}
