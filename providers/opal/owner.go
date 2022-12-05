package opal

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type OwnerGenerator struct {
	OpalService
}

func (g *OwnerGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to list opal owners: %v", err)
	}

	owners, _, err := client.OwnersApi.GetOwners(context.TODO()).Execute()
	if err != nil {
		return fmt.Errorf("unable to list opal owners: %v", err)
	}

	countByName := make(map[string]int)

	for {
		for _, owner := range owners.Results {
			name := normalizeResourceName(*owner.Name)
			if count, ok := countByName[name]; ok {
				countByName[name] = count + 1
				name = normalizeResourceName(fmt.Sprintf("%s_%d", *owner.Name, count+1))
			} else {
				countByName[name] = 1
			}

			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				owner.OwnerId,
				name,
				"opal_owner",
				"opal",
				[]string{},
			))
		}

		if !owners.HasNext() || owners.Next.Get() == nil {
			break
		}

		owners, _, err = client.OwnersApi.GetOwners(context.TODO()).Cursor(*owners.Next.Get()).Execute()
		if err != nil {
			return fmt.Errorf("unable to list opal owners: %v", err)
		}
	}

	return nil
}
