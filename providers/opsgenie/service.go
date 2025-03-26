package opsgenie

import (
	"context"
	"fmt"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk-v2/service"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ServiceGenerator struct {
	OpsgenieService
}

func (g *ServiceGenerator) InitResources() error {
	client, err := g.ServiceClient()
	if err != nil {
		return err
	}

	limit := 50
	offset := 0

	var services []service.Service

	for {
		result, err := func(limit, offset int) (*service.ListResult, error) {
			ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancelFunc()

			return client.List(ctx, &service.ListRequest{Limit: limit, Offset: offset})
		}(limit, offset)

		if err != nil {
			return err
		}

		if len(result.Services) == 0 {
			break
		}

		services = append(services, result.Services...)
		offset += limit
	}

	g.Resources = g.createResources(services)
	return nil
}

func (g *ServiceGenerator) createResources(services []service.Service) []terraformutils.Resource {
	var resources []terraformutils.Resource

	for _, s := range services {
		resources = append(resources, terraformutils.NewResource(
			s.Id,
			fmt.Sprintf("%s-%s", s.Id, s.Name),
			"opsgenie_service",
			g.ProviderName,
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}
