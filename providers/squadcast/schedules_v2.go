// service resource is yet to be implemented

package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SchedulesGenerator struct {
	SCService
}

type ScheduleQueryStruct struct {
	Schedules []*Schedule `graphql:"schedules(filters:  { teamID: $teamID })"`
}

type Schedule struct {
	ID        int         `graphql:"ID" json:"ID"`
	Rotations []*Rotation `graphql:"rotations" json:"rotations"`
}

type Rotation struct {
	ID int `graphql:"ID" json:"id"`
}

func (g *SchedulesGenerator) createResources(schedules []*Schedule) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, sch := range schedules {
		resourceList = append(resourceList, terraformutils.NewResource(
			fmt.Sprintf("%d", sch.ID),
			fmt.Sprintf("schedule_v2_%d", sch.ID),
			"squadcast_schedule_v2",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.Args["team_id"].(string),
			},
			[]string{},
			map[string]interface{}{},
		))

		for _, rot := range sch.Rotations {
			resourceList = append(resourceList, terraformutils.NewResource(
				fmt.Sprintf("%d", rot.ID),
				fmt.Sprintf("rotation_v2_%d", rot.ID),
				"squadcast_schedule_rotation_v2",
				g.GetProviderName(),
				map[string]string{
					"schedule_id": fmt.Sprintf("%d", sch.ID),
				},
				[]string{},
				map[string]interface{}{},
			))
		}
	}

	return resourceList
}

func (g *SchedulesGenerator) InitResources() error {
	var payload ScheduleQueryStruct

	variables := map[string]interface{}{
		"teamID": g.Args["team_id"].(string),
	}

	response, err := GraphQLRequest[ScheduleQueryStruct]("query", g.Args["access_token"].(string), g.Args["region"].(string), &payload, variables)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(response.Schedules)
	return nil
}
