// service resource is yet to be implemented

package squadcast

import (
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SLOGenerator struct {
	SquadcastService
}

type SLO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type getSLOsResponse struct {
	SLOs *[]SLO `json:"slos"`
}

func (g *SLOGenerator) createResources(slo []SLO) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, s := range slo {
		resourceList = append(resourceList, terraformutils.NewResource(
			fmt.Sprintf("%d", s.ID),
			fmt.Sprintf("slo_%s", s.Name),
			"squadcast_slo",
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

func (g *SLOGenerator) InitResources() error {
	if len(g.Args["team_name"].(string)) == 0 {
		return errors.New("--team-name is required")
	}

	getSLOsURL := fmt.Sprintf("/v3/slo?owner_id=%s", g.Args["team_id"].(string))
	response, err := Request[getSLOsResponse](getSLOsURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response.SLOs)
	return nil
}
