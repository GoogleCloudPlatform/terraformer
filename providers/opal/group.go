package opal

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type GroupGenerator struct {
	OpalService
}

func (g *GroupGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to list opal groups: %v", err)
	}

	groups, _, err := client.GroupsApi.GetGroups(context.TODO()).Execute()
	if err != nil {
		return fmt.Errorf("unable to list opal groups: %v", err)
	}

	countByName := make(map[string]int)

	for {
		for _, group := range groups.Results {
			name := normalizeResourceName(*group.Name)
			if count, ok := countByName[name]; ok {
				countByName[name] = count + 1
				name = normalizeResourceName(fmt.Sprintf("%s_%d", *group.Name, count+1))
			} else {
				countByName[name] = 1
			}

			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				group.GroupId,
				name,
				"opal_group",
				"opal",
				[]string{},
			))
		}

		if !groups.HasNext() || groups.Next.Get() == nil {
			break
		}

		groups, _, err = client.GroupsApi.GetGroups(context.TODO()).Cursor(*groups.Next.Get()).Execute()
		if err != nil {
			return fmt.Errorf("unable to list opal groups: %v", err)
		}
	}

	return nil
}
