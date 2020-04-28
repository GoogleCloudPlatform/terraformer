package keycloak

import (
	"fmt"
	"strconv"
	"strings"
)

type LdapRoleMapper struct {
	Id                   string
	Name                 string
	RealmId              string
	LdapUserFederationId string

	LdapRolesDn                 string
	RoleNameLdapAttribute       string
	RoleObjectClasses           []string
	MembershipLdapAttribute     string
	MembershipAttributeType     string
	MembershipUserLdapAttribute string
	RolesLdapFilter             string
	Mode                        string
	UserRolesRetrieveStrategy   string
	MemberofLdapAttribute       string
	UseRealmRolesMapping        bool
	ClientId                    string
}

func convertFromLdapRoleMapperToComponent(ldapRoleMapper *LdapRoleMapper) *component {
	componentConfig := map[string][]string{
		"roles.dn": {
			ldapRoleMapper.LdapRolesDn,
		},
		"role.name.ldap.attribute": {
			ldapRoleMapper.RoleNameLdapAttribute,
		},
		"role.object.classes": {
			strings.Join(ldapRoleMapper.RoleObjectClasses, ","),
		},
		"membership.ldap.attribute": {
			ldapRoleMapper.MembershipLdapAttribute,
		},
		"membership.attribute.type": {
			ldapRoleMapper.MembershipAttributeType,
		},
		"membership.user.ldap.attribute": {
			ldapRoleMapper.MembershipUserLdapAttribute,
		},
		"mode": {
			ldapRoleMapper.Mode,
		},
		"user.roles.retrieve.strategy": {
			ldapRoleMapper.UserRolesRetrieveStrategy,
		},
		"memberof.ldap.attribute": {
			ldapRoleMapper.MemberofLdapAttribute,
		},
		"use.realm.roles.mapping": {
			strconv.FormatBool(ldapRoleMapper.UseRealmRolesMapping),
		},
	}

	if ldapRoleMapper.RolesLdapFilter != "" {
		componentConfig["roles.ldap.filter"] = []string{ldapRoleMapper.RolesLdapFilter}
	}

	if ldapRoleMapper.ClientId != "" {
		componentConfig["client.id"] = []string{ldapRoleMapper.ClientId}
	}

	return &component{
		Id:           ldapRoleMapper.Id,
		Name:         ldapRoleMapper.Name,
		ProviderId:   "role-ldap-mapper",
		ProviderType: "org.keycloak.storage.ldap.mappers.LDAPStorageMapper",
		ParentId:     ldapRoleMapper.LdapUserFederationId,
		Config:       componentConfig,
	}
}

func convertFromComponentToLdapRoleMapper(component *component, realmId string) (*LdapRoleMapper, error) {
	roleObjectClasses := strings.Split(component.getConfig("role.object.classes"), ",")
	for i, v := range roleObjectClasses {
		roleObjectClasses[i] = strings.TrimSpace(v)
	}

	useRealmRolesMapping, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("use.realm.roles.mapping"))
	if err != nil {
		return nil, err
	}

	ldapRoleMapper := &LdapRoleMapper{
		Id:                   component.Id,
		Name:                 component.Name,
		RealmId:              realmId,
		LdapUserFederationId: component.ParentId,

		LdapRolesDn:                 component.getConfig("roles.dn"),
		RoleNameLdapAttribute:       component.getConfig("role.name.ldap.attribute"),
		RoleObjectClasses:           roleObjectClasses,
		MembershipLdapAttribute:     component.getConfig("membership.ldap.attribute"),
		MembershipAttributeType:     component.getConfig("membership.attribute.type"),
		MembershipUserLdapAttribute: component.getConfig("membership.user.ldap.attribute"),
		Mode:                        component.getConfig("mode"),
		UserRolesRetrieveStrategy:   component.getConfig("user.roles.retrieve.strategy"),
		MemberofLdapAttribute:       component.getConfig("memberof.ldap.attribute"),
		UseRealmRolesMapping:        useRealmRolesMapping,
	}

	if rolesLdapFilter := component.getConfig("roles.ldap.filter"); rolesLdapFilter != "" {
		ldapRoleMapper.RolesLdapFilter = rolesLdapFilter
	}

	if clientId := component.getConfig("client.id"); clientId != "" {
		ldapRoleMapper.ClientId = clientId
	}

	return ldapRoleMapper, nil
}

func (keycloakClient *KeycloakClient) ValidateLdapRoleMapper(ldapRoleMapper *LdapRoleMapper) error {
	return nil
}

func (keycloakClient *KeycloakClient) NewLdapRoleMapper(ldapRoleMapper *LdapRoleMapper) error {
	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/components", ldapRoleMapper.RealmId), convertFromLdapRoleMapperToComponent(ldapRoleMapper))
	if err != nil {
		return err
	}

	ldapRoleMapper.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetLdapRoleMapper(realmId, id string) (*LdapRoleMapper, error) {
	var component *component

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/components/%s", realmId, id), &component, nil)
	if err != nil {
		return nil, err
	}

	return convertFromComponentToLdapRoleMapper(component, realmId)
}

func (keycloakClient *KeycloakClient) UpdateLdapRoleMapper(ldapRoleMapper *LdapRoleMapper) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/components/%s", ldapRoleMapper.RealmId, ldapRoleMapper.Id), convertFromLdapRoleMapperToComponent(ldapRoleMapper))
}

func (keycloakClient *KeycloakClient) DeleteLdapRoleMapper(realmId, id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/components/%s", realmId, id), nil)
}
