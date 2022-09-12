package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type UserGenerator struct {
	SquadcastService
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
	getUsersURL := "/v3/users"
	response, err := Request[[]User](getUsersURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
