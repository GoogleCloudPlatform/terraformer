// service resource is yet to be implemented

package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ServiceGenerator struct {
	SCService
}

type Service struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (g *ServiceGenerator) createResources(services []Service) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, service := range services {
		resourceList = append(resourceList, terraformutils.NewResource(
			service.ID,
			fmt.Sprintf("service_%s", service.Name),
			"squadcast_service",
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

func (g *ServiceGenerator) InitResources() error {
	req := TRequest{
		URL:             "/v3/services",
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
	}
	response, err := Request[[]Service](req)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
