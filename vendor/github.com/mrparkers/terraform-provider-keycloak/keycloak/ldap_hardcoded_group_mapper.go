package keycloak

import "fmt"

type LdapHardcodedGroupMapper struct {
	Id                   string
	Name                 string
	RealmId              string
	LdapUserFederationId string
	Group                string
}

func convertFromLdapHardcodedGroupMapperToComponent(ldapMapper *LdapHardcodedGroupMapper) *component {
	return &component{
		Id:           ldapMapper.Id,
		Name:         ldapMapper.Name,
		ProviderId:   "hardcoded-ldap-group-mapper",
		ProviderType: "org.keycloak.storage.ldap.mappers.LDAPStorageMapper",
		ParentId:     ldapMapper.LdapUserFederationId,

		Config: map[string][]string{
			"group": {
				ldapMapper.Group,
			},
		},
	}
}

func convertFromComponentToLdapHardcodedGroupMapper(component *component, realmId string) *LdapHardcodedGroupMapper {
	return &LdapHardcodedGroupMapper{
		Id:                   component.Id,
		Name:                 component.Name,
		RealmId:              realmId,
		LdapUserFederationId: component.ParentId,

		Group: component.getConfig("group"),
	}
}

func (keycloakClient *KeycloakClient) ValidateLdapHardcodedGroupMapper(ldapMapper *LdapHardcodedGroupMapper) error {
	if len(ldapMapper.Group) == 0 {
		return fmt.Errorf("validation error: hardcoded group name must not be empty")
	}
	return nil
}

func (keycloakClient *KeycloakClient) NewLdapHardcodedGroupMapper(ldapMapper *LdapHardcodedGroupMapper) error {
	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/components", ldapMapper.RealmId), convertFromLdapHardcodedGroupMapperToComponent(ldapMapper))
	if err != nil {
		return err
	}

	ldapMapper.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetLdapHardcodedGroupMapper(realmId, id string) (*LdapHardcodedGroupMapper, error) {
	var component *component

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/components/%s", realmId, id), &component, nil)
	if err != nil {
		return nil, err
	}

	return convertFromComponentToLdapHardcodedGroupMapper(component, realmId), nil
}

func (keycloakClient *KeycloakClient) UpdateLdapHardcodedGroupMapper(ldapMapper *LdapHardcodedGroupMapper) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/components/%s", ldapMapper.RealmId, ldapMapper.Id), convertFromLdapHardcodedGroupMapperToComponent(ldapMapper))
}

func (keycloakClient *KeycloakClient) DeleteLdapHardcodedGroupMapper(realmId, id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/components/%s", realmId, id), nil)
}
