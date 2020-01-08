package keycloak

import (
	"encoding/json"
	"fmt"
)

type OpenidClientAuthorizationScope struct {
	Id               string `json:"id,omitempty"`
	RealmId          string `json:"-"`
	ResourceServerId string `json:"-"`
	Name             string `json:"name"`
	DisplayName      string `json:"displayName"`
	IconUri          string `json:"iconUri"`
}

func (keycloakClient *KeycloakClient) NewOpenidClientAuthorizationScope(scope *OpenidClientAuthorizationScope) error {
	body, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/scope", scope.RealmId, scope.ResourceServerId), scope)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &scope)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) GetOpenidClientAuthorizationScope(realm, resourceServerId, scopeId string) (*OpenidClientAuthorizationScope, error) {
	scope := OpenidClientAuthorizationScope{
		RealmId:          realm,
		ResourceServerId: resourceServerId,
	}
	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/scope/%s", realm, resourceServerId, scopeId), &scope, nil)
	if err != nil {
		return nil, err
	}
	return &scope, nil
}

func (keycloakClient *KeycloakClient) UpdateOpenidClientAuthorizationScope(scope *OpenidClientAuthorizationScope) error {
	err := keycloakClient.put(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/scope/%s", scope.RealmId, scope.ResourceServerId, scope.Id), scope)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) DeleteOpenidClientAuthorizationScope(realmId, resourceServerId, scopeId string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/scope/%s", realmId, resourceServerId, scopeId), nil)
}
