package keycloak

import (
	"fmt"
	"log"
	"net/url"
)

type Role struct {
	Id          string `json:"id,omitempty"`
	RealmId     string `json:"-"`
	ClientId    string `json:"-"`
	RoleId      string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ClientRole  bool   `json:"clientRole"`
	ContainerId string `json:"containerId"`
	Composite   bool   `json:"composite"`
}

type UsersInRole struct {
	Role  *Role
	Users *[]User
}

/*
 * Realm roles: /realms/${realm_id}/roles
 * Client roles: /realms/${realm_id}/clients/${client_id}/roles
 */
func roleByNameUrl(realmId, clientId string) string {
	if clientId == "" {
		return fmt.Sprintf("/realms/%s/roles", realmId)
	}

	return fmt.Sprintf("/realms/%s/clients/%s/roles", realmId, clientId)
}

func (keycloakClient *KeycloakClient) CreateRole(role *Role) error {
	roleUrl := roleByNameUrl(role.RealmId, role.ClientId)

	if role.ClientId != "" {
		role.ContainerId = role.ClientId
		role.ClientRole = true
	}

	_, _, err := keycloakClient.post(roleUrl, role)
	if err != nil {
		return err
	}

	var createdRole Role
	var roleName = url.PathEscape(role.Name)

	err = keycloakClient.get(fmt.Sprintf("%s/%s", roleUrl, roleName), &createdRole, nil)
	if err != nil {
		return err
	}

	role.Id = createdRole.Id

	return nil
}

func (keycloakClient *KeycloakClient) GetRealmRoles(realmId string) ([]*Role, error) {
	var roles []*Role

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/roles", realmId), &roles, nil)
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		role.RealmId = realmId
	}

	return roles, nil
}

func (keycloakClient *KeycloakClient) GetClientRoles(realmId string, clients []*OpenidClient) ([]*Role, error) {
	var roles []*Role

	for _, client := range clients {
		var rolesClient []*Role

		err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/roles", realmId, client.Id), &rolesClient, nil)
		if err != nil {
			return nil, err
		}

		for _, roleClient := range rolesClient {
			roleClient.RealmId = realmId
			roleClient.ClientId = client.Id
		}

		roles = append(roles, rolesClient...)
	}

	return roles, nil
}

func (keycloakClient *KeycloakClient) GetClientRoleUsers(realmId string, roles []*Role) (*[]UsersInRole, error) {
	var usersInRoles []UsersInRole

	for _, role := range roles {
		var usersInRole UsersInRole

		usersInRole.Role = role
		err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/roles/%s/users", realmId, role.ClientId, role.Name), &usersInRole.Users, nil)
		if usersInRole.Users == nil {
			continue
		}
		if err != nil {
			return nil, err
		}

		usersInRoles = append(usersInRoles, usersInRole)
	}

	return &usersInRoles, nil
}

func (keycloakClient *KeycloakClient) GetRole(realmId, id string) (*Role, error) {
	var role Role
	err := keycloakClient.get(fmt.Sprintf("/realms/%s/roles-by-id/%s", realmId, id), &role, nil)
	if err != nil {
		return nil, err
	}

	role.RealmId = realmId

	if role.ClientRole {
		role.ClientId = role.ContainerId
	}

	return &role, nil
}

func (keycloakClient *KeycloakClient) GetRoleByName(realmId, clientId, name string) (*Role, error) {
	var role Role
	var roleName = url.PathEscape(name)

	err := keycloakClient.get(fmt.Sprintf("%s/%s", roleByNameUrl(realmId, clientId), roleName), &role, nil)
	if err != nil {
		return nil, err
	}

	role.RealmId = realmId

	if role.ClientRole {
		role.ClientId = role.ContainerId
	}

	return &role, nil
}

func (keycloakClient *KeycloakClient) UpdateRole(role *Role) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/roles-by-id/%s", role.RealmId, role.Id), role)
}

func (keycloakClient *KeycloakClient) DeleteRole(realmId, id string) error {
	err := keycloakClient.delete(fmt.Sprintf("/realms/%s/roles-by-id/%s", realmId, id), nil)
	if err != nil {
		log.Printf("[DEBUG] Failed to delete role with id %s. Trying again...", id)

		return keycloakClient.delete(fmt.Sprintf("/realms/%s/roles-by-id/%s", realmId, id), nil)
	}

	return nil
}

func (keycloakClient *KeycloakClient) AddCompositesToRole(role *Role, compositeRoles []*Role) error {
	_, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/roles-by-id/%s/composites", role.RealmId, role.Id), compositeRoles)
	if err != nil {
		return err
	}

	return nil
}

func (keycloakClient *KeycloakClient) RemoveCompositesFromRole(role *Role, compositeRoles []*Role) error {
	err := keycloakClient.delete(fmt.Sprintf("/realms/%s/roles-by-id/%s/composites", role.RealmId, role.Id), compositeRoles)
	if err != nil {
		return err
	}

	return nil
}

func (keycloakClient *KeycloakClient) GetRoleComposites(role *Role) ([]*Role, error) {
	var composites []*Role

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/roles-by-id/%s/composites", role.RealmId, role.Id), &composites, nil)
	if err != nil {
		return nil, err
	}

	return composites, nil
}
