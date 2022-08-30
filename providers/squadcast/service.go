// service resource is yet to be implemented

package squadcast

import (
	"encoding/json"
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ServiceGenerator struct {
	SquadcastService
}

type Service struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var getServicesResponse struct {
	Data *[]Service `json:"data"`
}

func (g *ServiceGenerator) createResources(services []Service) []terraformutils.Resource {
	var serviceList []terraformutils.Resource
	for _, service := range services {
		serviceList = append(serviceList, terraformutils.NewResource(
			service.ID,
			"service_"+(service.Name),
			"squadcast_service",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.Args["team_id"].(string),
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return serviceList
}

func (g *ServiceGenerator) InitResources() error {
	if g.Args["team_id"].(string) == "" {
		return errors.New("--team-name is required")
	}
	getServicesURL := "/v3/services"
	body, err := g.generateRequest(getServicesURL)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(body, &getServicesResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getServicesResponse.Data)
	return nil
}
