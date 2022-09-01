package squadcast

import (
	"encoding/json"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchedulesGenerator struct {
	SquadcastService
}

type Schedule struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

var getSchedulesResponse struct {
	Data *[]Schedule `json:"data"`
}

func (g *SchedulesGenerator) createResources(schedules []Schedule) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, schedule := range schedules {
		resourceList = append(resourceList, terraformutils.NewResource(
			schedule.ID,
			schedule.Name,
			"squadcast_schedule",
			g.GetProviderName(),
			map[string]string{
				"team_id":  g.Args["team_id"].(string),
				"name": 	g.Args["schedule_name"].(string),
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resourceList
}

func (g *SchedulesGenerator) InitResources() error {
	getSchedules := "/v3/schedules"
	body, err := g.generateRequest(getSchedules)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &getSchedulesResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getSchedulesResponse.Data)

	return nil
}
