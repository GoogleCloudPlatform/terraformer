package keycloak

import (
	"fmt"
)

type AuthenticationFlow struct {
	Id          string `json:"id,omitempty"`
	RealmId     string `json:"-"`
	Alias       string `json:"alias"`
	Description string `json:"description"`
	ProviderId  string `json:"providerId"` // "basic-flow" or "client-flow"
	TopLevel    bool   `json:"topLevel"`
	BuiltIn     bool   `json:"builtIn"`
}

func (keycloakClient *KeycloakClient) NewAuthenticationFlow(authenticationFlow *AuthenticationFlow) error {
	authenticationFlow.TopLevel = true
	authenticationFlow.BuiltIn = false

	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/authentication/flows", authenticationFlow.RealmId), authenticationFlow)
	if err != nil {
		return err
	}
	authenticationFlow.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetAuthenticationFlow(realmId, id string) (*AuthenticationFlow, error) {
	var authenticationFlow AuthenticationFlow
	err := keycloakClient.get(fmt.Sprintf("/realms/%s/authentication/flows/%s", realmId, id), &authenticationFlow, nil)
	if err != nil {
		return nil, err
	}

	authenticationFlow.RealmId = realmId
	return &authenticationFlow, nil
}

func (keycloakClient *KeycloakClient) UpdateAuthenticationFlow(authenticationFlow *AuthenticationFlow) error {
	authenticationFlow.TopLevel = true
	authenticationFlow.BuiltIn = false

	return keycloakClient.put(fmt.Sprintf("/realms/%s/authentication/flows/%s", authenticationFlow.RealmId, authenticationFlow.Id), authenticationFlow)
}

func (keycloakClient *KeycloakClient) DeleteAuthenticationFlow(realmId, id string) error {
	err := keycloakClient.delete(fmt.Sprintf("/realms/%s/authentication/flows/%s", realmId, id), nil)
	if err != nil {
		// For whatever reason, this fails sometimes with a 500 during acceptance tests. try again
		return keycloakClient.delete(fmt.Sprintf("/realms/%s/authentication/flows/%s", realmId, id), nil)
	}
	return nil
}
