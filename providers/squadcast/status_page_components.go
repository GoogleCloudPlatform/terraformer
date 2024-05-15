// service resource is yet to be implemented

package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type StatusPageComponentsGenerator struct {
	SCService
}

type StatusPageComponent struct {
	ID     uint `json:"id"`
	PageID uint `json:"pageID"`
}

func (g *StatusPageComponentsGenerator) createResources(statusPageComponents []StatusPageComponent) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, spc := range statusPageComponents {
		resourceList = append(resourceList, terraformutils.NewResource(
			fmt.Sprintf("%d", spc.ID),
			fmt.Sprintf("status_page_component_%d", spc.ID),
			"squadcast_status_page_component",
			g.GetProviderName(),
			map[string]string{
				"status_page_id": fmt.Sprintf("%d", spc.PageID),
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resourceList
}

func (g *StatusPageComponentsGenerator) InitResources() error {
	req := TRequest{
		URL:             fmt.Sprintf("/v4/statuspages?teamID=%s", g.Args["team_id"].(string)),
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
	}
	response, _, err := Request[[]StatusPage](req)
	if err != nil {
		return err
	}
	for _, sp := range *response {
		req := TRequest{
			URL:             fmt.Sprintf("/v4/statuspages/%d/components", sp.ID),
			AccessToken:     g.Args["access_token"].(string),
			Region:          g.Args["region"].(string),
			IsAuthenticated: true,
		}
		response, _, err := Request[[]StatusPageComponent](req)
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, g.createResources(*response)...)
	}
	return nil
}
