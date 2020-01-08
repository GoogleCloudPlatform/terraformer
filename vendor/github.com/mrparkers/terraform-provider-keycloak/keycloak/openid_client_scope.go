package keycloak

import (
	"fmt"
)

type OpenidClientScope struct {
	Id          string `json:"id,omitempty"`
	RealmId     string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Protocol    string `json:"protocol"`
	Attributes  struct {
		DisplayOnConsentScreen string `json:"display.on.consent.screen"` // boolean in string form
		ConsentScreenText      string `json:"consent.screen.text"`
	} `json:"attributes"`
}

type OpenidClientScopeFilterFunc func(*OpenidClientScope) bool

func (keycloakClient *KeycloakClient) NewOpenidClientScope(clientScope *OpenidClientScope) error {
	clientScope.Protocol = "openid-connect"

	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/client-scopes", clientScope.RealmId), clientScope)
	if err != nil {
		return err
	}

	clientScope.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetOpenidClientScope(realmId, id string) (*OpenidClientScope, error) {
	var clientScope OpenidClientScope

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/client-scopes/%s", realmId, id), &clientScope, nil)
	if err != nil {
		return nil, err
	}

	clientScope.RealmId = realmId

	return &clientScope, nil
}

func (keycloakClient *KeycloakClient) GetOpenidDefaultClientScopes(realmId, clientId string) (*[]OpenidClientScope, error) {
	var clientScopes []OpenidClientScope

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/default-client-scopes", realmId, clientId), &clientScopes, nil)
	if err != nil {
		return nil, err
	}

	for _, clientScope := range clientScopes {
		clientScope.RealmId = realmId
	}

	return &clientScopes, nil
}

func (keycloakClient *KeycloakClient) GetOpenidOptionalClientScopes(realmId, clientId string) (*[]OpenidClientScope, error) {
	var clientScopes []OpenidClientScope

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/optional-client-scopes", realmId, clientId), &clientScopes, nil)
	if err != nil {
		return nil, err
	}

	for _, clientScope := range clientScopes {
		clientScope.RealmId = realmId
	}

	return &clientScopes, nil
}

func (keycloakClient *KeycloakClient) UpdateOpenidClientScope(clientScope *OpenidClientScope) error {
	clientScope.Protocol = "openid-connect"

	return keycloakClient.put(fmt.Sprintf("/realms/%s/client-scopes/%s", clientScope.RealmId, clientScope.Id), clientScope)
}

func (keycloakClient *KeycloakClient) DeleteOpenidClientScope(realmId, id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/client-scopes/%s", realmId, id), nil)
}

func (keycloakClient *KeycloakClient) ListOpenidClientScopesWithFilter(realmId string, filter OpenidClientScopeFilterFunc) ([]*OpenidClientScope, error) {
	var clientScopes []OpenidClientScope
	var openidClientScopes []*OpenidClientScope

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/client-scopes", realmId), &clientScopes, nil)
	if err != nil {
		return nil, err
	}

	for _, clientScope := range clientScopes {
		if clientScope.Protocol == "openid-connect" && filter(&clientScope) {
			scope := new(OpenidClientScope)
			*scope = clientScope

			openidClientScopes = append(openidClientScopes, scope)
		}
	}

	return openidClientScopes, nil
}

func includeOpenidClientScopesMatchingNames(scopeNames []string) OpenidClientScopeFilterFunc {
	return func(scope *OpenidClientScope) bool {
		for _, scopeName := range scopeNames {
			if scopeName == scope.Name {
				return true
			}
		}

		return false
	}
}
