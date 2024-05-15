// service resource is yet to be implemented

package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type StatusPagesGenerator struct {
	SCService
}

type StatusPage struct {
	ID uint `json:"id"`
}

func (g *StatusPagesGenerator) createResources(statusPages []StatusPage) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, sp := range statusPages {
		resourceList = append(resourceList, terraformutils.NewResource(
			fmt.Sprintf("%d", sp.ID),
			fmt.Sprintf("status_page_%d", sp.ID),
			"squadcast_status_page",
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

func (g *StatusPagesGenerator) InitResources() error {
	req := TRequest{
		URL:             fmt.Sprintf("/v4/statuspages?teamID=%s", g.Args["team_id"].(string)),
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
	}
	response, err := Request[[]StatusPage](req)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
