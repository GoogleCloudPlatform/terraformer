package keycloak

import (
	"fmt"
)

type OpenidClientServiceAccountRealmRole struct {
	Id                   string `json:"id"`
	RealmId              string `json:"-"`
	ServiceAccountUserId string `json:"-"`
	Name                 string `json:"name,omitempty"`
	Description          string `json:"description"`
}

func (keycloakClient *KeycloakClient) NewOpenidClientServiceAccountRealmRole(serviceAccountRole *OpenidClientServiceAccountRealmRole) error {
	serviceAccountRoles := []OpenidClientServiceAccountRealmRole{*serviceAccountRole}

	_, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/users/%s/role-mappings/realm", serviceAccountRole.RealmId, serviceAccountRole.ServiceAccountUserId), serviceAccountRoles)

	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) DeleteOpenidClientServiceAccountRealmRole(realm, serviceAccountUserId, roleId string) error {
	serviceAccountRole, err := keycloakClient.GetOpenidClientServiceAccountRealmRole(realm, serviceAccountUserId, roleId)
	if err != nil {
		return err
	}
	serviceAccountRoles := []OpenidClientServiceAccountRealmRole{*serviceAccountRole}
	err = keycloakClient.delete(fmt.Sprintf("/realms/%s/users/%s/role-mappings/realm", realm, serviceAccountUserId), &serviceAccountRoles)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) GetOpenidClientServiceAccountRealmRole(realm, serviceAccountUserId, roleId string) (*OpenidClientServiceAccountRealmRole, error) {
	serviceAccountRoles := []OpenidClientServiceAccountRealmRole{
		{
			Id:                   roleId,
			RealmId:              realm,
			ServiceAccountUserId: serviceAccountUserId,
		},
	}
	err := keycloakClient.get(fmt.Sprintf("/realms/%s/users/%s/role-mappings/realm", realm, serviceAccountUserId), &serviceAccountRoles, nil)
	if err != nil {
		return nil, err
	}
	for _, serviceAccountRole := range serviceAccountRoles {
		if serviceAccountRole.Id == roleId {
			serviceAccountRole.RealmId = realm
			serviceAccountRole.ServiceAccountUserId = serviceAccountUserId
			return &serviceAccountRole, nil
		}
	}
	return &OpenidClientServiceAccountRealmRole{}, nil
}
