package authlete

import (
	"context"
	"errors"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	authlete "github.com/authlete/openapi-for-go"
)

type ClientGenerator struct {
	AuthleteService
}

func (c *ClientGenerator) InitResources() error {
	if c.Resources == nil {
		c.Resources = []terraformutils.Resource{}
	}

	authleteClient := c.getClient()

	apiKey := c.GetArgs()["api_key"].(string)
	apiSecret := c.GetArgs()["api_secret"].(string)
	auth := context.WithValue(context.Background(), authlete.ContextBasicAuth, authlete.BasicAuth{
		UserName: apiKey,
		Password: apiSecret,
	})

	end := int32(0)
	total := int32(10)
	for end < total {
		listClients, _, err := authleteClient.ClientManagementApi.ClientGetListApi(auth).Start(end).End(total).Execute()
		if err != nil {
			return errors.New("could not fetch the client list:  " + err.Error())
		}
		total = *listClients.TotalCount
		end = *listClients.End
		c.Resources = append(c.Resources, mapClients(listClients, apiKey, apiSecret)...)
	}

	return nil
}

func mapClients(clients *authlete.ClientGetListResponse, apiKey string, apiSecret string) []terraformutils.Resource {
	result := []terraformutils.Resource{}

	for _, client := range clients.Clients {
		newResource := terraformutils.NewResource(
			strconv.FormatInt(client.GetClientId(), 10),
			strconv.FormatInt(client.GetClientId(), 10),
			"authlete_client",
			"authlete",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		)
		result = append(result, newResource)
	}

	return result
}
