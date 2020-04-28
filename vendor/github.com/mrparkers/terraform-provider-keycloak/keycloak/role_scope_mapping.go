package keycloak

import (
	"fmt"
)

func roleScopeMappingUrl(realmId, clientId string, clientScopeId string, role *Role) string {
	if clientId != "" {
		return fmt.Sprintf("/realms/%s/clients/%s/scope-mappings/clients/%s", realmId, clientId, role.ClientId)
	} else {
		return fmt.Sprintf("/realms/%s/client-scopes/%s/scope-mappings/clients/%s", realmId, clientScopeId, role.ClientId)
	}
}

func (keycloakClient *KeycloakClient) CreateRoleScopeMapping(realmId string, clientId string, clientScopeId string, role *Role) error {
	roleUrl := roleScopeMappingUrl(realmId, clientId, clientScopeId, role)

	_, _, err := keycloakClient.post(roleUrl, []Role{*role})
	if err != nil {
		return err
	}

	return nil
}

func (keycloakClient *KeycloakClient) GetRoleScopeMapping(realmId string, clientId string, clientScopeId string, role *Role) (*Role, error) {
	roleUrl := roleScopeMappingUrl(realmId, clientId, clientScopeId, role)
	var roles []Role

	err := keycloakClient.get(roleUrl, &roles, nil)
	if err != nil {
		return nil, err
	}

	for _, mappedRole := range roles {
		if mappedRole.Id == role.Id {
			return role, nil
		}
	}

	return nil, nil
}

func (keycloakClient *KeycloakClient) DeleteRoleScopeMapping(realmId string, clientId string, clientScopeId string, role *Role) error {
	roleUrl := roleScopeMappingUrl(realmId, clientId, clientScopeId, role)
	return keycloakClient.delete(roleUrl, nil)
}
