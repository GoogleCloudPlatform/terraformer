package squadcast

import (
	"encoding/json"
	"errors"
	"fmt"
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

	var host string
	switch g.Args["region"] {
	case "us":
		host = "squadcast.com"
	case "eu":
		host = "eu.squadcast.com"
	case "internal":
		host = "squadcast.xyz"
	case "staging":
		host = "squadcast.tech"
	case "dev":
		host = "localhost"
	default:
		return errors.New("unknown region")
	}

	body, err := g.generateRequest(fmt.Sprintf("https://api.%s/v3/users", host))
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
