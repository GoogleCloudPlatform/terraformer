package keycloak

import (
	"fmt"
	"strconv"
	"strings"
)

type LdapGroupMapper struct {
	Id                   string
	Name                 string
	RealmId              string
	LdapUserFederationId string

	LdapGroupsDn                string
	GroupNameLdapAttribute      string
	GroupObjectClasses          []string
	PreserveGroupInheritance    bool
	IgnoreMissingGroups         bool
	MembershipLdapAttribute     string
	MembershipAttributeType     string
	MembershipUserLdapAttribute string
	GroupsLdapFilter            string
	Mode                        string
	UserRolesRetrieveStrategy   string
	MemberofLdapAttribute       string
	MappedGroupAttributes       []string

	DropNonExistingGroupsDuringSync bool
}

func convertFromLdapGroupMapperToComponent(ldapGroupMapper *LdapGroupMapper) *component {
	componentConfig := map[string][]string{
		"groups.dn": {
			ldapGroupMapper.LdapGroupsDn,
		},
		"group.name.ldap.attribute": {
			ldapGroupMapper.GroupNameLdapAttribute,
		},
		"group.object.classes": {
			strings.Join(ldapGroupMapper.GroupObjectClasses, ","),
		},
		"preserve.group.inheritance": {
			strconv.FormatBool(ldapGroupMapper.PreserveGroupInheritance),
		},
		"ignore.missing.groups": {
			strconv.FormatBool(ldapGroupMapper.IgnoreMissingGroups),
		},
		"membership.ldap.attribute": {
			ldapGroupMapper.MembershipLdapAttribute,
		},
		"membership.attribute.type": {
			ldapGroupMapper.MembershipAttributeType,
		},
		"membership.user.ldap.attribute": {
			ldapGroupMapper.MembershipUserLdapAttribute,
		},
		"mode": {
			ldapGroupMapper.Mode,
		},
		"user.roles.retrieve.strategy": {
			ldapGroupMapper.UserRolesRetrieveStrategy,
		},
		"memberof.ldap.attribute": {
			ldapGroupMapper.MemberofLdapAttribute,
		},
		"drop.non.existing.groups.during.sync": {
			strconv.FormatBool(ldapGroupMapper.DropNonExistingGroupsDuringSync),
		},
	}

	if ldapGroupMapper.GroupsLdapFilter != "" {
		componentConfig["groups.ldap.filter"] = []string{ldapGroupMapper.GroupsLdapFilter}
	}

	if len(ldapGroupMapper.MappedGroupAttributes) != 0 {
		componentConfig["mapped.group.attributes"] = []string{strings.Join(ldapGroupMapper.MappedGroupAttributes, ",")}
	}

	return &component{
		Id:           ldapGroupMapper.Id,
		Name:         ldapGroupMapper.Name,
		ProviderId:   "group-ldap-mapper",
		ProviderType: "org.keycloak.storage.ldap.mappers.LDAPStorageMapper",
		ParentId:     ldapGroupMapper.LdapUserFederationId,
		Config:       componentConfig,
	}
}

func convertFromComponentToLdapGroupMapper(component *component, realmId string) (*LdapGroupMapper, error) {
	groupObjectClasses := strings.Split(component.getConfig("group.object.classes"), ",")
	for i, v := range groupObjectClasses {
		groupObjectClasses[i] = strings.TrimSpace(v)
	}

	preserveGroupInheritance, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("preserve.group.inheritance"))
	if err != nil {
		return nil, err
	}

	ignoreMissingGroups, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("ignore.missing.groups"))
	if err != nil {
		return nil, err
	}

	dropNonExistingGroupsDuringSync, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("drop.non.existing.groups.during.sync"))
	if err != nil {
		return nil, err
	}

	ldapGroupMapper := &LdapGroupMapper{
		Id:                   component.Id,
		Name:                 component.Name,
		RealmId:              realmId,
		LdapUserFederationId: component.ParentId,

		LdapGroupsDn:                    component.getConfig("groups.dn"),
		GroupNameLdapAttribute:          component.getConfig("group.name.ldap.attribute"),
		GroupObjectClasses:              groupObjectClasses,
		PreserveGroupInheritance:        preserveGroupInheritance,
		IgnoreMissingGroups:             ignoreMissingGroups,
		MembershipLdapAttribute:         component.getConfig("membership.ldap.attribute"),
		MembershipAttributeType:         component.getConfig("membership.attribute.type"),
		MembershipUserLdapAttribute:     component.getConfig("membership.user.ldap.attribute"),
		Mode:                            component.getConfig("mode"),
		UserRolesRetrieveStrategy:       component.getConfig("user.roles.retrieve.strategy"),
		MemberofLdapAttribute:           component.getConfig("memberof.ldap.attribute"),
		DropNonExistingGroupsDuringSync: dropNonExistingGroupsDuringSync,
	}

	if groupsLdapFilter := component.getConfig("groups.ldap.filter"); groupsLdapFilter != "" {
		ldapGroupMapper.GroupsLdapFilter = groupsLdapFilter
	}

	if mappedGroupAttributesString := component.getConfig("mapped.group.attributes"); mappedGroupAttributesString != "" {
		mappedGroupAttributes := strings.Split(mappedGroupAttributesString, ",")
		for i, v := range mappedGroupAttributes {
			mappedGroupAttributes[i] = strings.TrimSpace(v)
		}

		ldapGroupMapper.MappedGroupAttributes = mappedGroupAttributes
	}

	return ldapGroupMapper, nil
}

func (keycloakClient *KeycloakClient) ValidateLdapGroupMapper(ldapGroupMapper *LdapGroupMapper) error {
	if ldapGroupMapper.MembershipAttributeType == "UID" && ldapGroupMapper.PreserveGroupInheritance == true {
		return fmt.Errorf("validation error: group inheritance cannot be preserved while membership attribute type is UID")
	}

	return nil
}

func (keycloakClient *KeycloakClient) NewLdapGroupMapper(ldapGroupMapper *LdapGroupMapper) error {
	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/components", ldapGroupMapper.RealmId), convertFromLdapGroupMapperToComponent(ldapGroupMapper))
	if err != nil {
		return err
	}

	ldapGroupMapper.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetLdapGroupMapper(realmId, id string) (*LdapGroupMapper, error) {
	var component *component

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/components/%s", realmId, id), &component, nil)
	if err != nil {
		return nil, err
	}

	return convertFromComponentToLdapGroupMapper(component, realmId)
}

func (keycloakClient *KeycloakClient) UpdateLdapGroupMapper(ldapGroupMapper *LdapGroupMapper) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/components/%s", ldapGroupMapper.RealmId, ldapGroupMapper.Id), convertFromLdapGroupMapperToComponent(ldapGroupMapper))
}

func (keycloakClient *KeycloakClient) DeleteLdapGroupMapper(realmId, id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/components/%s", realmId, id), nil)
}
