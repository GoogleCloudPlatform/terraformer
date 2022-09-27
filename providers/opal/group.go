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

	for {
		for _, group := range groups.Results {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				group.GroupId,
				normalizeResourceName(*group.Name),
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
