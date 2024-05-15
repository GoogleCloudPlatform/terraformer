package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type UserGenerator struct {
	SCService
}

type User struct {
	ID string `json:"id"`
}

func (g *UserGenerator) createResources(users []User) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, user := range users {
		resourceList = append(resourceList, terraformutils.NewSimpleResource(
			user.ID,
			fmt.Sprintf("user_%s", user.ID),
			"squadcast_user",
			g.GetProviderName(),
			[]string{},
		))
	}
	return resourceList
}

func (g *UserGenerator) InitResources() error {
	req := TRequest{
		URL:             "/v3/users",
		AccessToken:     g.Args["access_token"].(string),
		Region:          g.Args["region"].(string),
		IsAuthenticated: true,
	}
	response, _, err := Request[[]User](req)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
