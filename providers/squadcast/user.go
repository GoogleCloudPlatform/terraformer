package squadcast

import (
	"encoding/json"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type UserGenerator struct {
	SquadcastService
}

type User struct {
	ID string `json:"id" tf:"id"`
}

var getUsersResponse struct {
	Data *[]User `json:"data"`
}

type Users []User

func (g *UserGenerator) createResources(users Users) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, user := range users {
		resources = append(resources, terraformutils.NewSimpleResource(
			user.ID,
			"user_"+(user.ID),
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
