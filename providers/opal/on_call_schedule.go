package opal

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type OnCallScheduleGenerator struct {
	OpalService
}

func (g *OnCallScheduleGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to list opal on call schedules: %v", err)
	}

	onCallSchedules, _, err := client.OnCallSchedulesApi.GetOnCallSchedules(context.TODO()).Execute()
	if err != nil {
		return fmt.Errorf("unable to list opal on call schedules: %v", err)
	}

	countByName := make(map[string]int)

	for _, onCallSchedule := range onCallSchedules.OnCallSchedules {
		name := normalizeResourceName(*onCallSchedule.Name)
		if count, ok := countByName[name]; ok {
			countByName[name] = count + 1
			name = normalizeResourceName(fmt.Sprintf("%s_%d", *onCallSchedule.Name, count+1))
		} else {
			countByName[name] = 1
		}

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*onCallSchedule.OnCallScheduleId,
			name,
			"opal_on_call_schedule",
			"opal",
			[]string{},
		))
	}

	return nil
}
