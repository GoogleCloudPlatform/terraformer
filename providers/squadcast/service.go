// service resource is yet to be implemented

package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ServiceGenerator struct {
	SquadcastService
	teamID string
}

type Service struct {
	ID   string `json:"id" tf:"id"`
	Name string `json:"name" tf:"name"`
}

type Services []Service

var getServicesResponse struct {
	Data *[]Service `json:"data"`
}

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
	teamName := g.Args["team_name"].(string)
	if len(teamName) == 0 {
		return errors.New("--team-name is required")
	}

	escapedTeamName := url.QueryEscape(teamName)
	getTeamURL := fmt.Sprintf("/v3/teams/by-name?name=%s", escapedTeamName)
	team, err := g.generateRequest(getTeamURL)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(team, &getTeamResponse)
	if err != nil {
		return err
	}
	g.teamID = getTeamResponse.Data.ID

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
