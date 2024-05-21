// service resource is yet to be implemented

package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type RunbookGenerator struct {
	SCService
}

type Runbook struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (g *RunbookGenerator) createResources(runbooks []Runbook) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, runbook := range runbooks {
		resourceList = append(resourceList, terraformutils.NewResource(
			runbook.ID,
			fmt.Sprintf("runbook_%s", runbook.Name),
			"squadcast_runbook",
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

func (g *RunbookGenerator) InitResources() error {
	req := TRequest{
		URL:             "/v3/runbooks",
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
	}
	response, _, err := Request[[]Runbook](req)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
