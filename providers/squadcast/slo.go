// service resource is yet to be implemented

package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SLOGenerator struct {
	SquadcastService
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
				"team_id": g.Args["team_id"].(string),
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return SLOList
}

func (g *SLOGenerator) InitResources() error {
	teamID := g.Args["team_id"].(string)
	if teamID == "" {
		return errors.New("--team-name is required")
	}
	getSLOsURL := fmt.Sprintf("/v3/slo?owner_id=%s", teamID)
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
