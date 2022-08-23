package squadcast

import (
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type UserGenerator struct {
	SquadcastService
}

type User struct {
	ID string `json:"id"`
}

var getUsersResponse struct {
	Data *[]User `json:"data"`
}


func (g *UserGenerator) createResources(users []User) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, user := range users {
		resources = append(resources, terraformutils.NewSimpleResource(
			user.ID,
			fmt.Sprintf("user_%s", user.ID),
			"squadcast_user",
			"squadcast",
			[]string{},
		))
	}
	return resources
}

func (g *UserGenerator) InitResources() error {
	getUsersURL := "/v3/users"
	body, err := g.generateRequest(getUsersURL)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &getUsersResponse)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*getUsersResponse.Data)
	return nil
}
