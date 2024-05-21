package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type WebformsGenerator struct {
	SCService
}

type Webform struct {
	ID uint `json:"id"`
}

func (g *WebformsGenerator) createResources(webforms []Webform) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, webform := range webforms {
		resourceList = append(resourceList, terraformutils.NewSimpleResource(
			fmt.Sprintf("%d", webform.ID),
			fmt.Sprintf("webform_%d", webform.ID),
			"squadcast_webform",
			g.GetProviderName(),
			[]string{},
		))
	}
	return resourceList
}

func (g *WebformsGenerator) InitResources() error {
	req := TRequest{
		URL:             fmt.Sprintf("/v3/webforms?owner_id=%s", g.Args["team_id"].(string)),
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
	}
	response, _, err := Request[[]Webform](req)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
