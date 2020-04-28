package keycloak

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type IdentityProviderConfig struct {
	Key                              string                 `json:"key,omitempty"`
	HostIp                           string                 `json:"hostIp,omitempty"`
	UseJwksUrl                       KeycloakBoolQuoted     `json:"useJwksUrl,omitempty"`
	JwksUrl                          string                 `json:"jwksUrl,omitempty"`
	ClientId                         string                 `json:"clientId,omitempty"`
	ClientSecret                     string                 `json:"clientSecret,omitempty"`
	DisableUserInfo                  KeycloakBoolQuoted     `json:"disableUserInfo"`
	UserInfoUrl                      string                 `json:"userInfoUrl,omitempty"`
	HideOnLoginPage                  KeycloakBoolQuoted     `json:"hideOnLoginPage"`
	NameIDPolicyFormat               string                 `json:"nameIDPolicyFormat,omitempty"`
	SingleLogoutServiceUrl           string                 `json:"singleLogoutServiceUrl,omitempty"`
	SingleSignOnServiceUrl           string                 `json:"singleSignOnServiceUrl,omitempty"`
	SigningCertificate               string                 `json:"signingCertificate,omitempty"`
	SignatureAlgorithm               string                 `json:"signatureAlgorithm,omitempty"`
	XmlSignKeyInfoKeyNameTransformer string                 `json:"xmlSignKeyInfoKeyNameTransformer,omitempty"`
	PostBindingAuthnRequest          KeycloakBoolQuoted     `json:"postBindingAuthnRequest,omitempty"`
	PostBindingResponse              KeycloakBoolQuoted     `json:"postBindingResponse,omitempty"`
	PostBindingLogout                KeycloakBoolQuoted     `json:"postBindingLogout,omitempty"`
	ForceAuthn                       KeycloakBoolQuoted     `json:"forceAuthn,omitempty"`
	WantAuthnRequestsSigned          KeycloakBoolQuoted     `json:"wantAuthnRequestsSigned,omitempty"`
	WantAssertionsSigned             KeycloakBoolQuoted     `json:"wantAssertionsSigned,omitempty"`
	WantAssertionsEncrypted          KeycloakBoolQuoted     `json:"wantAssertionsEncrypted,omitempty"`
	BackchannelSupported             KeycloakBoolQuoted     `json:"backchannelSupported,omitempty"`
	ValidateSignature                KeycloakBoolQuoted     `json:"validateSignature,omitempty"`
	AuthorizationUrl                 string                 `json:"authorizationUrl,omitempty"`
	TokenUrl                         string                 `json:"tokenUrl,omitempty"`
	LoginHint                        string                 `json:"loginHint,omitempty"`
	UILocales                        KeycloakBoolQuoted     `json:"uiLocales,omitempty"`
	LogoutUrl                        string                 `json:"logoutUrl,omitempty"`
	DefaultScope                     string                 `json:"defaultScope,omitempty"`
	AcceptsPromptNoneForwFrmClt      KeycloakBoolQuoted     `json:"acceptsPromptNoneForwardFromClient,omitempty"`
	HostedDomain                     string                 `json:"hostedDomain,omitempty"`
	UserIp                           KeycloakBoolQuoted     `json:"userIp,omitempty"`
	OfflineAccess                    KeycloakBoolQuoted     `json:"offlineAccess,omitempty"`
	ExtraConfig                      map[string]interface{} `json:"-"`
}

type IdentityProvider struct {
	Realm                     string                  `json:"-"`
	InternalId                string                  `json:"internalId,omitempty"`
	Alias                     string                  `json:"alias"`
	DisplayName               string                  `json:"displayName"`
	ProviderId                string                  `json:"providerId"`
	Enabled                   bool                    `json:"enabled"`
	StoreToken                bool                    `json:"storeToken"`
	AddReadTokenRoleOnCreate  bool                    `json:"addReadTokenRoleOnCreate"`
	AuthenticateByDefault     bool                    `json:"authenticateByDefault"`
	LinkOnly                  bool                    `json:"linkOnly"`
	TrustEmail                bool                    `json:"trustEmail"`
	FirstBrokerLoginFlowAlias string                  `json:"firstBrokerLoginFlowAlias"`
	PostBrokerLoginFlowAlias  string                  `json:"postBrokerLoginFlowAlias"`
	Config                    *IdentityProviderConfig `json:"config"`
}

func (f *IdentityProviderConfig) UnmarshalJSON(data []byte) error {
	f.ExtraConfig = map[string]interface{}{}
	err := json.Unmarshal(data, &f.ExtraConfig)
	if err != nil {
		return err
	}
	v := reflect.ValueOf(f).Elem()
	for i := 0; i < v.NumField(); i++ {
		structField := v.Type().Field(i)
		jsonKey := strings.Split(structField.Tag.Get("json"), ",")[0]
		if jsonKey != "-" {
			value, ok := f.ExtraConfig[jsonKey]
			if ok {
				field := v.FieldByName(structField.Name)
				if field.IsValid() && field.CanSet() {
					if field.Kind() == reflect.String {
						field.SetString(value.(string))
					} else if field.Kind() == reflect.Bool {
						boolVal, err := strconv.ParseBool(value.(string))
						if err == nil {
							field.Set(reflect.ValueOf(KeycloakBoolQuoted(boolVal)))
						}
					}
					delete(f.ExtraConfig, jsonKey)
				}
			}
		}
	}
	return nil
}

func (f *IdentityProviderConfig) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}

	for k, v := range f.ExtraConfig {
		out[k] = v
	}
	v := reflect.ValueOf(f).Elem()
	for i := 0; i < v.NumField(); i++ {
		jsonKey := strings.Split(v.Type().Field(i).Tag.Get("json"), ",")[0]
		if jsonKey != "-" {
			field := v.Field(i)
			if field.IsValid() && field.CanSet() {
				if field.Kind() == reflect.String {
					out[jsonKey] = field.String()
				} else if field.Kind() == reflect.Bool {
					out[jsonKey] = KeycloakBoolQuoted(field.Bool())
				}
			}
		}
	}
	return json.Marshal(out)
}

func (keycloakClient *KeycloakClient) NewIdentityProvider(identityProvider *IdentityProvider) error {
	log.Printf("[WARN] Realm: %s", identityProvider.Realm)
	_, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/identity-provider/instances", identityProvider.Realm), identityProvider)
	if err != nil {
		return err
	}

	return nil
}

func (keycloakClient *KeycloakClient) GetIdentityProvider(realm, alias string) (*IdentityProvider, error) {
	var identityProvider IdentityProvider
	identityProvider.Realm = realm

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/identity-provider/instances/%s", realm, alias), &identityProvider, nil)
	if err != nil {
		return nil, err
	}

	return &identityProvider, nil
}

func (keycloakClient *KeycloakClient) UpdateIdentityProvider(identityProvider *IdentityProvider) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/identity-provider/instances/%s", identityProvider.Realm, identityProvider.Alias), identityProvider)
}

func (keycloakClient *KeycloakClient) DeleteIdentityProvider(realm, alias string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/identity-provider/instances/%s", realm, alias), nil)
}
