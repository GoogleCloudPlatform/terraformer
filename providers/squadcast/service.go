// service resource is yet to be implemented

package squadcast

import (
	"encoding/json"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ServiceGenerator struct {
	SquadcastService
}

type Service struct {
	ID   string `json:"id" tf:"id"`
	Name string `json:"name" tf:"name"`
}

var responseService struct {
	Data *[]Service `json:"data"`
}

type Services []Service

func (g *ServiceGenerator) createResources(services Services) []terraformutils.Resource {
	var serviceList []terraformutils.Resource
	for _, service := range services {
		serviceList = append(serviceList, terraformutils.NewSimpleResource(
			service.ID,
			"service_"+(service.Name),
			"squadcast_service",
			"squadcast",
			[]string{},
		))
	}
	return serviceList
}

func (g *ServiceGenerator) InitResources() error {
	body, err := g.generateRequest("https://api.squadcast.com/v3/services")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &responseService)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*responseService.Data)
	return nil
}
