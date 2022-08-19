// service resource is yet to be implemented

package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"net/url"
)

type ServiceGenerator struct {
	SquadcastService
	teamID string
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
		serviceList = append(serviceList, terraformutils.NewResource(
			service.ID,
			"service_"+(service.Name),
			"squadcast_service",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.teamID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return serviceList
}

func (g *ServiceGenerator) InitResources() error {
	if len(g.Args["team_name"].(string)) == 0 {
		return errors.New("--team-name is required")
	}
	team, err := g.generateRequest(fmt.Sprintf("/v3/teams/by-name?name=%s", url.QueryEscape(g.Args["team_name"].(string))))
	if err != nil {
		return err
	}
	err = json.Unmarshal(team, &ResponseTeam)
	if err != nil {
		return err
	}
	g.teamID = ResponseTeam.Data.ID

	body, err := g.generateRequest("/v3/services")
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
