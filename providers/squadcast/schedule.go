package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchedulesGenerator struct {
	SquadcastService
}

type Schedule struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (g *SchedulesGenerator) createResources(schedules []Schedule) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, schedule := range schedules {
		resourceList = append(resourceList, terraformutils.NewResource(
			schedule.ID,
			fmt.Sprintf("schedule_%s", schedule.Name),
			"squadcast_schedule",
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

func (g *SchedulesGenerator) InitResources() error {
	getSchedulesURL := "/v3/schedules"
	response, err := Request[[]Schedule](getSchedulesURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
