package keycloak

import "fmt"

type GenericClient struct {
	Id       string `json:"id,omitempty"`
	ClientId string `json:"clientId"`
	RealmId  string `json:"-"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`

	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
}

func (keycloakClient *KeycloakClient) listGenericClients(realmId string) ([]*GenericClient, error) {
	var clients []*GenericClient

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients", realmId), &clients, nil)
	if err != nil {
		return nil, err
	}

	for _, client := range clients {
		client.RealmId = realmId
	}

	return clients, nil
}

func (keycloakClient *KeycloakClient) GetGenericClient(realmId, id string) (*GenericClient, error) {
	var client GenericClient

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s", realmId, id), &client, nil)
	if err != nil {
		return nil, err
	}

	client.RealmId = realmId

	return &client, nil
}

func (keycloakClient *KeycloakClient) GetGenericClientByClientId(realmId, clientId string) (*GenericClient, error) {
	var clients []GenericClient

	params := map[string]string{
		"clientId": clientId,
	}

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients", realmId), &clients, params)
	if err != nil {
		return nil, err
	}

	if len(clients) == 0 {
		return nil, fmt.Errorf("generic client with name %s does not exist", clientId)
	}

	client := clients[0]

	client.RealmId = realmId

	return &client, nil
}
