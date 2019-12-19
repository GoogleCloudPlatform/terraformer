package keycloak

import (
	"fmt"
	"strconv"
)

type LdapMsadUserAccountControlMapper struct {
	Id                   string
	Name                 string
	RealmId              string
	LdapUserFederationId string

	LdapPasswordPolicyHintsEnabled bool
}

func convertFromLdapMsadUserAccountControlMapperToComponent(ldapMsadUserAccountControlMapper *LdapMsadUserAccountControlMapper) *component {
	return &component{
		Id:           ldapMsadUserAccountControlMapper.Id,
		Name:         ldapMsadUserAccountControlMapper.Name,
		ProviderId:   "msad-user-account-control-mapper",
		ProviderType: "org.keycloak.storage.ldap.mappers.LDAPStorageMapper",
		ParentId:     ldapMsadUserAccountControlMapper.LdapUserFederationId,
		Config: map[string][]string{
			"ldap.password.policy.hints.enabled": {
				strconv.FormatBool(ldapMsadUserAccountControlMapper.LdapPasswordPolicyHintsEnabled),
			},
		},
	}
}

func convertFromComponentToLdapMsadUserAccountControlMapper(component *component, realmId string) (*LdapMsadUserAccountControlMapper, error) {
	ldapPasswordPolicyHintsEnabled, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("ldap.password.policy.hints.enabled"))
	if err != nil {
		return nil, err
	}

	return &LdapMsadUserAccountControlMapper{
		Id:                   component.Id,
		Name:                 component.Name,
		RealmId:              realmId,
		LdapUserFederationId: component.ParentId,

		LdapPasswordPolicyHintsEnabled: ldapPasswordPolicyHintsEnabled,
	}, nil
}

func (keycloakClient *KeycloakClient) NewLdapMsadUserAccountControlMapper(ldapMsadUserAccountControlMapper *LdapMsadUserAccountControlMapper) error {
	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/components", ldapMsadUserAccountControlMapper.RealmId), convertFromLdapMsadUserAccountControlMapperToComponent(ldapMsadUserAccountControlMapper))
	if err != nil {
		return err
	}

	ldapMsadUserAccountControlMapper.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetLdapMsadUserAccountControlMapper(realmId, id string) (*LdapMsadUserAccountControlMapper, error) {
	var component *component

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/components/%s", realmId, id), &component, nil)
	if err != nil {
		return nil, err
	}

	return convertFromComponentToLdapMsadUserAccountControlMapper(component, realmId)
}

func (keycloakClient *KeycloakClient) UpdateLdapMsadUserAccountControlMapper(ldapMsadUserAccountControlMapper *LdapMsadUserAccountControlMapper) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/components/%s", ldapMsadUserAccountControlMapper.RealmId, ldapMsadUserAccountControlMapper.Id), convertFromLdapMsadUserAccountControlMapperToComponent(ldapMsadUserAccountControlMapper))
}

func (keycloakClient *KeycloakClient) DeleteLdapMsadUserAccountControlMapper(realmId, id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/components/%s", realmId, id), nil)
}
