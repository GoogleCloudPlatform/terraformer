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

	for {
		for _, owner := range owners.Results {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				owner.OwnerId,
				normalizeResourceName(*owner.Name),
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
