package keycloak

import (
	"encoding/json"
	"fmt"
)

type OpenidClientAuthorizationPermission struct {
	Id               string   `json:"id,omitempty"`
	RealmId          string   `json:"-"`
	ResourceServerId string   `json:"-"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	DecisionStrategy string   `json:"decisionStrategy"`
	Policies         []string `json:"policies"`
	Resources        []string `json:"resources"`
	Type             string   `json:"type"`
}

func (keycloakClient *KeycloakClient) GetOpenidClientAuthorizationPermission(realm, resourceServerId, id string) (*OpenidClientAuthorizationPermission, error) {
	permission := OpenidClientAuthorizationPermission{
		RealmId:          realm,
		ResourceServerId: resourceServerId,
		Id:               id,
	}

	policies := []OpenidClientAuthorizationPolicy{}
	resources := []OpenidClientAuthorizationResource{}

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/permission/resource/%s", realm, resourceServerId, id), &permission, nil)
	if err != nil {
		return nil, err
	}

	err = keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/%s/associatedPolicies", realm, resourceServerId, id), &policies, nil)
	if err != nil {
		return nil, err
	}

	err = keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/permission/%s/resources", realm, resourceServerId, id), &resources, nil)
	if err != nil {
		return nil, err
	}

	for _, policy := range policies {
		permission.Policies = append(permission.Policies, policy.Id)
	}

	for _, resource := range resources {
		permission.Resources = append(permission.Resources, resource.Id)
	}

	return &permission, nil
}

func (keycloakClient *KeycloakClient) NewOpenidClientAuthorizationPermission(permission *OpenidClientAuthorizationPermission) error {
	body, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/permission", permission.RealmId, permission.ResourceServerId), permission)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &permission)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) UpdateOpenidClientAuthorizationPermission(permission *OpenidClientAuthorizationPermission) error {
	err := keycloakClient.put(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/permission/resource/%s", permission.RealmId, permission.ResourceServerId, permission.Id), permission)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) DeleteOpenidClientAuthorizationPermission(realmId, resourceServerId, permissionId string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/permission/%s", realmId, resourceServerId, permissionId), nil)
}
