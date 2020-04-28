package keycloak

import (
	"fmt"
	"strconv"
	"strings"
)

type LdapUserFederation struct {
	Id      string
	Name    string
	RealmId string

	Enabled  bool
	Priority int

	ImportEnabled     bool
	EditMode          string // can be "READ_ONLY", "WRITABLE", or "UNSYNCED"
	SyncRegistrations bool   // I think this field controls whether or not BatchSizeForSync, FullSyncPeriod, and ChangedSyncPeriod are needed

	Vendor                 string // can be "other", "edirectory", "ad", "rhds", or "tivoli". honestly I don't think this field actually does anything
	UsernameLDAPAttribute  string
	RdnLDAPAttribute       string
	UuidLDAPAttribute      string
	UserObjectClasses      []string // api expects comma + space separated for some reason
	ConnectionUrl          string
	UsersDn                string
	BindDn                 string
	BindCredential         string
	CustomUserSearchFilter string // must start with '(' and end with ')'
	SearchScope            string // api expects "1" or "2", but that means "One Level" or "Subtree"

	ValidatePasswordPolicy bool
	UseTruststoreSpi       string // can be "ldapsOnly", "always", or "never"
	ConnectionTimeout      string // duration string (ex: 1h30m)
	ReadTimeout            string // duration string (ex: 1h30m)
	Pagination             bool

	BatchSizeForSync  int
	FullSyncPeriod    int // either a number, in milliseconds, or -1 if full sync is disabled
	ChangedSyncPeriod int // either a number, in milliseconds, or -1 if changed sync is disabled

	CachePolicy string
}

func convertFromLdapUserFederationToComponent(ldap *LdapUserFederation) (*component, error) {
	componentConfig := map[string][]string{
		"cachePolicy": {
			ldap.CachePolicy,
		},
		"enabled": {
			strconv.FormatBool(ldap.Enabled),
		},
		"priority": {
			strconv.Itoa(ldap.Priority),
		},
		"importEnabled": {
			strconv.FormatBool(ldap.ImportEnabled),
		},
		"editMode": {
			ldap.EditMode,
		},
		"syncRegistrations": {
			strconv.FormatBool(ldap.SyncRegistrations),
		},
		"vendor": {
			strings.ToLower(ldap.Vendor),
		},
		"usernameLDAPAttribute": {
			ldap.UsernameLDAPAttribute,
		},
		"rdnLDAPAttribute": {
			ldap.RdnLDAPAttribute,
		},
		"uuidLDAPAttribute": {
			ldap.UuidLDAPAttribute,
		},
		"userObjectClasses": {
			strings.Join(ldap.UserObjectClasses, ", "),
		},
		"connectionUrl": {
			ldap.ConnectionUrl,
		},
		"usersDn": {
			ldap.UsersDn,
		},
		"searchScope": {
			ldap.SearchScope,
		},
		"validatePasswordPolicy": {
			strconv.FormatBool(ldap.ValidatePasswordPolicy),
		},
		"pagination": {
			strconv.FormatBool(ldap.Pagination),
		},
		"batchSizeForSync": {
			strconv.Itoa(ldap.BatchSizeForSync),
		},
		"fullSyncPeriod": {
			strconv.Itoa(ldap.FullSyncPeriod),
		},
		"changedSyncPeriod": {
			strconv.Itoa(ldap.ChangedSyncPeriod),
		},
	}

	if ldap.BindDn != "" && ldap.BindCredential != "" {
		componentConfig["bindDn"] = []string{ldap.BindDn}
		componentConfig["bindCredential"] = []string{ldap.BindCredential}

		componentConfig["authType"] = []string{"simple"}
	} else {
		componentConfig["authType"] = []string{"none"}
	}

	if ldap.SearchScope == "ONE_LEVEL" {
		componentConfig["searchScope"] = []string{"1"}
	} else {
		componentConfig["searchScope"] = []string{"2"}
	}

	if ldap.CustomUserSearchFilter != "" {
		componentConfig["customUserSearchFilter"] = []string{ldap.CustomUserSearchFilter}
	}

	if ldap.UseTruststoreSpi == "ONLY_FOR_LDAPS" {
		componentConfig["useTruststoreSpi"] = []string{"ldapsOnly"}
	} else {
		componentConfig["useTruststoreSpi"] = []string{strings.ToLower(ldap.UseTruststoreSpi)}
	}

	if ldap.ConnectionTimeout != "" {
		connectionTimeoutMs, err := getMillisecondsFromDurationString(ldap.ConnectionTimeout)
		if err != nil {
			return nil, err
		}

		componentConfig["connectionTimeout"] = []string{connectionTimeoutMs}
	} else {
		componentConfig["connectionTimeout"] = []string{} // the keycloak API will not unset this unless the config is present with an empty array
	}

	if ldap.ReadTimeout != "" {
		readTimeoutMs, err := getMillisecondsFromDurationString(ldap.ReadTimeout)
		if err != nil {
			return nil, err
		}

		componentConfig["readTimeout"] = []string{readTimeoutMs}
	} else {
		componentConfig["readTimeout"] = []string{} // the keycloak API will not unset this unless the config is present with an empty array
	}

	return &component{
		Id:           ldap.Id,
		Name:         ldap.Name,
		ProviderId:   "ldap",
		ProviderType: userStorageProviderType,
		ParentId:     ldap.RealmId,
		Config:       componentConfig,
	}, nil
}

func convertFromComponentToLdapUserFederation(component *component) (*LdapUserFederation, error) {
	enabled, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("enabled"))
	if err != nil {
		return nil, err
	}

	priority, err := strconv.Atoi(component.getConfig("priority"))
	if err != nil {
		return nil, err
	}

	importEnabled, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("importEnabled"))
	if err != nil {
		return nil, err
	}

	syncRegistrations, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("syncRegistrations"))
	if err != nil {
		return nil, err
	}

	userObjectClasses := strings.Split(component.getConfig("userObjectClasses"), ", ")

	validatePasswordPolicy, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("validatePasswordPolicy"))
	if err != nil {
		return nil, err
	}

	pagination, err := parseBoolAndTreatEmptyStringAsFalse(component.getConfig("pagination"))
	if err != nil {
		return nil, err
	}

	batchSizeForSync, err := strconv.Atoi(component.getConfig("batchSizeForSync"))
	if err != nil {
		return nil, err
	}

	fullSyncPeriod, err := strconv.Atoi(component.getConfig("fullSyncPeriod"))
	if err != nil {
		return nil, err
	}

	changedSyncPeriod, err := strconv.Atoi(component.getConfig("changedSyncPeriod"))
	if err != nil {
		return nil, err
	}

	ldap := &LdapUserFederation{
		Id:      component.Id,
		Name:    component.Name,
		RealmId: component.ParentId,

		Enabled:  enabled,
		Priority: priority,

		ImportEnabled:     importEnabled,
		EditMode:          component.getConfig("editMode"),
		SyncRegistrations: syncRegistrations,

		Vendor:                 strings.ToUpper(component.getConfig("vendor")),
		UsernameLDAPAttribute:  component.getConfig("usernameLDAPAttribute"),
		RdnLDAPAttribute:       component.getConfig("rdnLDAPAttribute"),
		UuidLDAPAttribute:      component.getConfig("uuidLDAPAttribute"),
		UserObjectClasses:      userObjectClasses,
		ConnectionUrl:          component.getConfig("connectionUrl"),
		UsersDn:                component.getConfig("usersDn"),
		BindDn:                 component.getConfig("bindDn"),
		BindCredential:         component.getConfig("bindCredential"),
		CustomUserSearchFilter: component.getConfig("customUserSearchFilter"),
		SearchScope:            component.getConfig("searchScope"),

		ValidatePasswordPolicy: validatePasswordPolicy,
		UseTruststoreSpi:       component.getConfig("useTruststoreSpi"),
		Pagination:             pagination,

		BatchSizeForSync:  batchSizeForSync,
		FullSyncPeriod:    fullSyncPeriod,
		ChangedSyncPeriod: changedSyncPeriod,

		CachePolicy: component.getConfig("cachePolicy"),
	}

	if bindDn := component.getConfig("bindDn"); bindDn != "" {
		ldap.BindDn = bindDn
	}

	if bindCredential := component.getConfig("bindCredential"); bindCredential != "" {
		ldap.BindCredential = bindCredential
	}

	if customUserSearchFilter := component.getConfig("customUserSearchFilter"); customUserSearchFilter != "" {
		ldap.CustomUserSearchFilter = customUserSearchFilter
	}

	if component.getConfig("searchScope") == "1" {
		ldap.SearchScope = "ONE_LEVEL"
	} else {
		ldap.SearchScope = "SUBTREE"
	}

	if useTruststoreSpi := component.getConfig("useTruststoreSpi"); useTruststoreSpi == "ldapsOnly" {
		ldap.UseTruststoreSpi = "ONLY_FOR_LDAPS"
	} else {
		ldap.UseTruststoreSpi = strings.ToUpper(useTruststoreSpi)
	}

	if connectionTimeout, ok := component.getConfigOk("connectionTimeout"); ok {
		connectionTimeoutDurationString, err := GetDurationStringFromMilliseconds(connectionTimeout)
		if err != nil {
			return nil, err
		}

		ldap.ConnectionTimeout = connectionTimeoutDurationString
	}

	if readTimeout, ok := component.getConfigOk("readTimeout"); ok {
		readTimeoutDurationString, err := GetDurationStringFromMilliseconds(readTimeout)
		if err != nil {
			return nil, err
		}

		ldap.ReadTimeout = readTimeoutDurationString
	}

	return ldap, nil
}

func (keycloakClient *KeycloakClient) ValidateLdapUserFederation(ldap *LdapUserFederation) error {
	if (ldap.BindDn == "" && ldap.BindCredential != "") || (ldap.BindDn != "" && ldap.BindCredential == "") {
		return fmt.Errorf("validation error: authentication requires both BindDN and BindCredential to be set")
	}

	return nil
}

func (keycloakClient *KeycloakClient) NewLdapUserFederation(ldapUserFederation *LdapUserFederation) error {
	component, err := convertFromLdapUserFederationToComponent(ldapUserFederation)
	if err != nil {
		return err
	}

	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/components", ldapUserFederation.RealmId), component)
	if err != nil {
		return err
	}

	ldapUserFederation.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetLdapUserFederation(realmId, id string) (*LdapUserFederation, error) {
	var component *component

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/components/%s", realmId, id), &component, nil)
	if err != nil {
		return nil, err
	}

	return convertFromComponentToLdapUserFederation(component)
}

func (keycloakClient *KeycloakClient) GetLdapUserFederationMappers(realmId, id string) (*[]interface{}, error) {
	var components []*component
	var ldapUserFederationMappers []interface{}

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/components?parent=%s&type=org.keycloak.storage.ldap.mappers.LDAPStorageMapper", realmId, id), &components, nil)
	if err != nil {
		return nil, err
	}
	for _, component := range components {
		switch component.ProviderId {
		case "full-name-ldap-mapper":
			mapper, err := convertFromComponentToLdapFullNameMapper(component, realmId)
			if err != nil {
				return nil, err
			}
			ldapUserFederationMappers = append(ldapUserFederationMappers, mapper)
		case "group-ldap-mapper":
			mapper, err := convertFromComponentToLdapGroupMapper(component, realmId)
			if err != nil {
				return nil, err
			}
			ldapUserFederationMappers = append(ldapUserFederationMappers, mapper)
		case "hardcoded-ldap-group-mapper":
			mapper := convertFromComponentToLdapHardcodedGroupMapper(component, realmId)
			ldapUserFederationMappers = append(ldapUserFederationMappers, mapper)
		case "hardcoded-ldap-role-mapper":
			mapper := convertFromComponentToLdapHardcodedRoleMapper(component, realmId)
			ldapUserFederationMappers = append(ldapUserFederationMappers, mapper)
		case "msad-lds-user-account-control-mapper":
			mapper, err := convertFromComponentToLdapMsadLdsUserAccountControlMapper(component, realmId)
			if err != nil {
				return nil, err
			}
			ldapUserFederationMappers = append(ldapUserFederationMappers, mapper)
		case "msad-user-account-control-mapper":
			mapper, err := convertFromComponentToLdapMsadUserAccountControlMapper(component, realmId)
			if err != nil {
				return nil, err
			}
			ldapUserFederationMappers = append(ldapUserFederationMappers, mapper)
		case "user-attribute-ldap-mapper":
			mapper, err := convertFromComponentToLdapUserAttributeMapper(component, realmId)
			if err != nil {
				return nil, err
			}
			ldapUserFederationMappers = append(ldapUserFederationMappers, mapper)
		case "role-ldap-mapper":
			mapper, err := convertFromComponentToLdapRoleMapper(component, realmId)
			if err != nil {
				return nil, err
			}
			ldapUserFederationMappers = append(ldapUserFederationMappers, mapper)
		}
	}

	return &ldapUserFederationMappers, nil
}

func (keycloakClient *KeycloakClient) UpdateLdapUserFederation(ldapUserFederation *LdapUserFederation) error {
	component, err := convertFromLdapUserFederationToComponent(ldapUserFederation)
	if err != nil {
		return err
	}

	return keycloakClient.put(fmt.Sprintf("/realms/%s/components/%s", ldapUserFederation.RealmId, ldapUserFederation.Id), component)
}

func (keycloakClient *KeycloakClient) DeleteLdapUserFederation(realmId, id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/components/%s", realmId, id), nil)
}
