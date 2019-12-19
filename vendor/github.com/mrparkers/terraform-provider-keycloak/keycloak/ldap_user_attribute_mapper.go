package keycloak

import (
	"fmt"
	"strconv"
)

type LdapUserAttributeMapper struct {
	Id                   string
	Name                 string
	RealmId              string
	LdapUserFederationId string

	LdapAttribute           string
	IsMandatoryInLdap       bool
	ReadOnly                bool
	AlwaysReadValueFromLdap bool
	UserModelAttribute      string
}

func convertFromLdapUserAttributeMapperToComponent(ldapUserAttributeMapper *LdapUserAttributeMapper) *component {
	return &component{
		Id:           ldapUserAttributeMapper.Id,
		Name:         ldapUserAttributeMapper.Name,
		ProviderId:   "user-attribute-ldap-mapper",
		ProviderType: "org.keycloak.storage.ldap.mappers.LDAPStorageMapper",
		ParentId:     ldapUserAttributeMapper.LdapUserFederationId,
		Config: map[string][]string{
			"ldap.attribute": {
				ldapUserAttributeMapper.LdapAttribute,
			},
			"is.mandatory.in.ldap": {
				strconv.FormatBool(ldapUserAttributeMapper.IsMandatoryInLdap),
			},
			"read.only": {
				strconv.FormatBool(ldapUserAttributeMapper.ReadOnly),
			},
			"always.read.value.from.ldap": {
				strconv.FormatBool(ldapUserAttributeMapper.AlwaysReadValueFromLdap),
			},
			"user.model.attribute": {
				ldapUserAttributeMapper.UserModelAttribute,
			},
		},
	}
}

func convertFromComponentToLdapUserAttributeMapper(component *component, realmId string) (*LdapUserAttributeMapper, error) {
	isMandatoryInLdap, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("is.mandatory.in.ldap"))
	if err != nil {
		return nil, err
	}

	readOnly, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("read.only"))
	if err != nil {
		return nil, err
	}

	alwaysReadValueFromLdap, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("always.read.value.from.ldap"))
	if err != nil {
		return nil, err
	}

	return &LdapUserAttributeMapper{
		Id:                   component.Id,
		Name:                 component.Name,
		RealmId:              realmId,
		LdapUserFederationId: component.ParentId,

		LdapAttribute:           component.getConfig("ldap.attribute"),
		IsMandatoryInLdap:       isMandatoryInLdap,
		ReadOnly:                readOnly,
		AlwaysReadValueFromLdap: alwaysReadValueFromLdap,
		UserModelAttribute:      component.getConfig("user.model.attribute"),
	}, nil
}

func (keycloakClient *KeycloakClient) NewLdapUserAttributeMapper(ldapUserAttributeMapper *LdapUserAttributeMapper) error {
	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/components", ldapUserAttributeMapper.RealmId), convertFromLdapUserAttributeMapperToComponent(ldapUserAttributeMapper))
	if err != nil {
		return err
	}

	ldapUserAttributeMapper.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetLdapUserAttributeMapper(realmId, id string) (*LdapUserAttributeMapper, error) {
	var component *component

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/components/%s", realmId, id), &component, nil)
	if err != nil {
		return nil, err
	}

	return convertFromComponentToLdapUserAttributeMapper(component, realmId)
}

func (keycloakClient *KeycloakClient) UpdateLdapUserAttributeMapper(ldapUserAttributeMapper *LdapUserAttributeMapper) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/components/%s", ldapUserAttributeMapper.RealmId, ldapUserAttributeMapper.Id), convertFromLdapUserAttributeMapperToComponent(ldapUserAttributeMapper))
}

func (keycloakClient *KeycloakClient) DeleteLdapUserAttributeMapper(realmId, id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/components/%s", realmId, id), nil)
}
