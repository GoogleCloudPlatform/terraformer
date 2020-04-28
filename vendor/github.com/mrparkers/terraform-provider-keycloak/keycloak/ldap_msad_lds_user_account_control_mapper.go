package keycloak

import (
	"fmt"
)

type LdapMsadLdsUserAccountControlMapper struct {
	Id                   string
	Name                 string
	RealmId              string
	LdapUserFederationId string
}

func convertFromLdapMsadLdsUserAccountControlMapperToComponent(ldapMsadLdsUserAccountControlMapper *LdapMsadLdsUserAccountControlMapper) *component {
	return &component{
		Id:           ldapMsadLdsUserAccountControlMapper.Id,
		Name:         ldapMsadLdsUserAccountControlMapper.Name,
		ProviderId:   "msad-lds-user-account-control-mapper",
		ProviderType: "org.keycloak.storage.ldap.mappers.LDAPStorageMapper",
		ParentId:     ldapMsadLdsUserAccountControlMapper.LdapUserFederationId,
	}
}

func convertFromComponentToLdapMsadLdsUserAccountControlMapper(component *component, realmId string) (*LdapMsadLdsUserAccountControlMapper, error) {
	return &LdapMsadLdsUserAccountControlMapper{
		Id:                   component.Id,
		Name:                 component.Name,
		RealmId:              realmId,
		LdapUserFederationId: component.ParentId,
	}, nil
}

func (keycloakClient *KeycloakClient) NewLdapMsadLdsUserAccountControlMapper(ldapMsadLdsUserAccountControlMapper *LdapMsadLdsUserAccountControlMapper) error {
	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/components", ldapMsadLdsUserAccountControlMapper.RealmId), convertFromLdapMsadLdsUserAccountControlMapperToComponent(ldapMsadLdsUserAccountControlMapper))
	if err != nil {
		return err
	}

	ldapMsadLdsUserAccountControlMapper.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetLdapMsadLdsUserAccountControlMapper(realmId, id string) (*LdapMsadLdsUserAccountControlMapper, error) {
	var component *component

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/components/%s", realmId, id), &component, nil)
	if err != nil {
		return nil, err
	}

	return convertFromComponentToLdapMsadLdsUserAccountControlMapper(component, realmId)
}

func (keycloakClient *KeycloakClient) UpdateLdapMsadLdsUserAccountControlMapper(ldapMsadLdsUserAccountControlMapper *LdapMsadLdsUserAccountControlMapper) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/components/%s", ldapMsadLdsUserAccountControlMapper.RealmId, ldapMsadLdsUserAccountControlMapper.Id), convertFromLdapMsadLdsUserAccountControlMapperToComponent(ldapMsadLdsUserAccountControlMapper))
}

func (keycloakClient *KeycloakClient) DeleteLdapMsadLdsUserAccountControlMapper(realmId, id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/components/%s", realmId, id), nil)
}
