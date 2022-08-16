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

var response struct {
	Data *[]User `json:"data"`
}

type Users []User

var UserAllowEmptyValues = []string{}

func (g UserGenerator) createResources(users Users) []terraformutils.Resource {
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
	// body, err := g.generateRequest("https://api.squadcast.com/v3/organizations/5a4262733c36823b3ed91fb9/users")
	body, err := g.generateRequest("https://api.squadcast.com/v3/users")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response.Data)
	return nil
}
