package opsgenie

import (
	"context"
	"fmt"
	"time"

	"github.com/opsgenie/opsgenie-go-sdk-v2/user"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type UserGenerator struct {
	OpsgenieService
}

func (g *UserGenerator) InitResources() error {
	client, err := g.UserClient()
	if err != nil {
		return err
	}

	limit := 50
	offset := 0

	var users []user.User

	for {
		result, err := func(limit, offset int) (*user.ListResult, error) {
			ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancelFunc()

			return client.List(ctx, &user.ListRequest{Limit: limit, Offset: offset})
		}(limit, offset)

		if err != nil {
			return err
		}

		users = append(users, result.Users...)
		offset += limit

		if offset >= result.TotalCount {
			break
		}
	}

	g.Resources = g.createResources(users)
	return nil
}

func (g *UserGenerator) createResources(users []user.User) []terraformutils.Resource {
	var resources []terraformutils.Resource

	for _, u := range users {
		resources = append(resources, terraformutils.NewResource(
			u.Id,
			fmt.Sprintf("%s-%s", u.Id, u.Username),
			"opsgenie_user",
			g.ProviderName,
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		))
	}

	return resources
}
