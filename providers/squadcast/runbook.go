// service resource is yet to be implemented

package squadcast

import (
	"encoding/json"
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type RunbookGenerator struct {
	SquadcastService
}

type Runbook struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}


var getRunbooksResponse struct {
	Data *[]Runbook `json:"data"`
}

func (g *RunbookGenerator) createResources(runbooks []Runbook) []terraformutils.Resource {
	var runbookList []terraformutils.Resource
	for _, runbook := range runbooks {
		runbookList = append(runbookList, terraformutils.NewResource(
			runbook.ID,
			"runbook_"+(runbook.Name),
			"squadcast_runbook",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.Args["team_id"].(string),
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return runbookList
}

func (g *RunbookGenerator) InitResources() error {
	if g.Args["team_id"].(string) == "" {
		return errors.New("--team-name is required")
	}
	getRunbooksURL := "/v3/runbooks"
	body, err := g.generateRequest(getRunbooksURL)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(body, &getRunbooksResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getRunbooksResponse.Data)
	return nil
}
