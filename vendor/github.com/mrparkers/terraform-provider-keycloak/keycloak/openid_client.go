package keycloak

import (
	"fmt"
)

type OpenidClientRole struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	ScopeParamRequired bool   `json:"scopeParamRequired"`
	ClientRole         bool   `json:"clientRole"`
	ContainerId        string `json:"ContainerId"`
}

type OpenidClientSecret struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type OpenidClientAuthorizationSettings struct {
	PolicyEnforcementMode         string `json:"policyEnforcementMode,omitempty"`
	AllowRemoteResourceManagement bool   `json:"allowRemoteResourceManagement,omitempty"`
	KeepDefaults                  bool   `json:"-"`
}

type OpenidClient struct {
	Id                                 string                                   `json:"id,omitempty"`
	ClientId                           string                                   `json:"clientId"`
	RealmId                            string                                   `json:"-"`
	Name                               string                                   `json:"name"`
	Protocol                           string                                   `json:"protocol"`                // always openid-connect for this resource
	ClientAuthenticatorType            string                                   `json:"clientAuthenticatorType"` // always client-secret for now, don't have a need for JWT here
	ClientSecret                       string                                   `json:"secret,omitempty"`
	Enabled                            bool                                     `json:"enabled"`
	Description                        string                                   `json:"description"`
	PublicClient                       bool                                     `json:"publicClient"`
	BearerOnly                         bool                                     `json:"bearerOnly"`
	StandardFlowEnabled                bool                                     `json:"standardFlowEnabled"`
	ImplicitFlowEnabled                bool                                     `json:"implicitFlowEnabled"`
	DirectAccessGrantsEnabled          bool                                     `json:"directAccessGrantsEnabled"`
	ServiceAccountsEnabled             bool                                     `json:"serviceAccountsEnabled"`
	AuthorizationServicesEnabled       bool                                     `json:"authorizationServicesEnabled"`
	ValidRedirectUris                  []string                                 `json:"redirectUris"`
	WebOrigins                         []string                                 `json:"webOrigins"`
	AdminUrl                           string                                   `json:"adminUrl"`
	BaseUrl                            string                                   `json:"baseUrl"`
	FullScopeAllowed                   bool                                     `json:"fullScopeAllowed"`
	Attributes                         OpenidClientAttributes                   `json:"attributes"`
	AuthorizationSettings              *OpenidClientAuthorizationSettings       `json:"authorizationSettings,omitempty"`
	ConsentRequired                    bool                                     `json:"consentRequired"`
	AuthenticationFlowBindingOverrides OpenidAuthenticationFlowBindingOverrides `json:"authenticationFlowBindingOverrides,omitempty"`
}

type OpenidClientAttributes struct {
	PkceCodeChallengeMethod             string             `json:"pkce.code.challenge.method"`
	ExcludeSessionStateFromAuthResponse KeycloakBoolQuoted `json:"exclude.session.state.from.auth.response"`
	AccessTokenLifespan                 string             `json:"access.token.lifespan"`
}

type OpenidAuthenticationFlowBindingOverrides struct {
	BrowserId     string `json:"browser"`
	DirectGrantId string `json:"direct_grant"`
}

func (keycloakClient *KeycloakClient) GetOpenidClientServiceAccountUserId(realmId, clientId string) (*User, error) {
	var serviceAccountUser User

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/service-account-user", realmId, clientId), &serviceAccountUser, nil)
	if err != nil {
		return &serviceAccountUser, err
	}

	serviceAccountUser.RealmId = realmId

	return &serviceAccountUser, nil
}

func (keycloakClient *KeycloakClient) ValidateOpenidClient(client *OpenidClient) error {
	if client.BearerOnly && (client.StandardFlowEnabled || client.ImplicitFlowEnabled || client.DirectAccessGrantsEnabled || client.ServiceAccountsEnabled) {
		return fmt.Errorf("validation error: Keycloak cannot issue tokens for bearer-only clients; no oauth2 flows can be enabled for this client")
	}

	if (client.StandardFlowEnabled || client.ImplicitFlowEnabled) && len(client.ValidRedirectUris) == 0 {
		return fmt.Errorf("validation error: standard (authorization code) and implicit flows require at least one valid redirect uri")
	}

	if client.ServiceAccountsEnabled && client.PublicClient {
		return fmt.Errorf("validation error: service accounts (client credentials flow) cannot be enabled on public clients")
	}

	return nil
}

func (keycloakClient *KeycloakClient) NewOpenidClient(client *OpenidClient) error {
	client.Protocol = "openid-connect"
	client.ClientAuthenticatorType = "client-secret"

	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/clients", client.RealmId), client)
	if err != nil {
		return err
	}

	client.Id = getIdFromLocationHeader(location)

	if authorizationSettings := client.AuthorizationSettings; authorizationSettings != nil {
		if !(*authorizationSettings).KeepDefaults {
			resource, err := keycloakClient.GetOpenidClientAuthorizationResourceByName(client.RealmId, client.Id, "default")
			if err != nil {
				return err
			}
			err = keycloakClient.DeleteOpenidClientAuthorizationResource(resource.RealmId, resource.ResourceServerId, resource.Id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (keycloakClient *KeycloakClient) GetOpenidClients(realmId string, withSecrets bool) ([]*OpenidClient, error) {
	var clients []*OpenidClient
	var clientSecret OpenidClientSecret

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients", realmId), &clients, nil)
	if err != nil {
		return nil, err
	}

	for _, client := range clients {
		client.RealmId = realmId
		if !withSecrets {
			continue
		}

		err = keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/client-secret", realmId, client.Id), &clientSecret, nil)
		if err != nil {
			return nil, err
		}

		client.ClientSecret = clientSecret.Value
	}

	return clients, nil
}

func (keycloakClient *KeycloakClient) GetOpenidClient(realmId, id string) (*OpenidClient, error) {
	var client OpenidClient
	var clientSecret OpenidClientSecret

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s", realmId, id), &client, nil)
	if err != nil {
		return nil, err
	}

	err = keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/client-secret", realmId, id), &clientSecret, nil)
	if err != nil {
		return nil, err
	}

	client.RealmId = realmId
	client.ClientSecret = clientSecret.Value

	return &client, nil
}

func (keycloakClient *KeycloakClient) GetOpenidClientByClientId(realmId, clientId string) (*OpenidClient, error) {
	var clients []OpenidClient
	var clientSecret OpenidClientSecret

	params := map[string]string{
		"clientId": clientId,
	}

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients", realmId), &clients, params)
	if err != nil {
		return nil, err
	}

	if len(clients) == 0 {
		return nil, fmt.Errorf("openid client with name %s does not exist", clientId)
	}

	client := clients[0]

	err = keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/client-secret", realmId, client.Id), &clientSecret, nil)
	if err != nil {
		return nil, err
	}

	client.RealmId = realmId
	client.ClientSecret = clientSecret.Value

	return &client, nil
}

func (keycloakClient *KeycloakClient) UpdateOpenidClient(client *OpenidClient) error {
	client.Protocol = "openid-connect"
	client.ClientAuthenticatorType = "client-secret"

	return keycloakClient.put(fmt.Sprintf("/realms/%s/clients/%s", client.RealmId, client.Id), client)
}

func (keycloakClient *KeycloakClient) DeleteOpenidClient(realmId, id string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/clients/%s", realmId, id), nil)
}

func (keycloakClient *KeycloakClient) getOpenidClientScopes(realmId, clientId, t string) ([]*OpenidClientScope, error) {
	var scopes []*OpenidClientScope

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/%s-client-scopes", realmId, clientId, t), &scopes, nil)
	if err != nil {
		return nil, err
	}

	return scopes, nil
}

func (keycloakClient *KeycloakClient) GetOpenidClientDefaultScopes(realmId, clientId string) ([]*OpenidClientScope, error) {
	return keycloakClient.getOpenidClientScopes(realmId, clientId, "default")
}

func (keycloakClient *KeycloakClient) GetOpenidClientOptionalScopes(realmId, clientId string) ([]*OpenidClientScope, error) {
	return keycloakClient.getOpenidClientScopes(realmId, clientId, "optional")
}

func (keycloakClient *KeycloakClient) getRealmClientScopes(realmId, t string) ([]*OpenidClientScope, error) {
	var scopes []*OpenidClientScope

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/default-%s-client-scopes", realmId, t), &scopes, nil)
	if err != nil {
		return nil, err
	}

	return scopes, nil
}

func (keycloakClient *KeycloakClient) GetRealmDefaultClientScopes(realmId string) ([]*OpenidClientScope, error) {
	return keycloakClient.getRealmClientScopes(realmId, "default")
}

func (keycloakClient *KeycloakClient) GetRealmOptionalClientScopes(realmId string) ([]*OpenidClientScope, error) {
	return keycloakClient.getRealmClientScopes(realmId, "optional")
}

func (keycloakClient *KeycloakClient) attachOpenidClientScopes(realmId, clientId, t string, scopeNames []string) error {
	openidClient, err := keycloakClient.GetOpenidClient(realmId, clientId)
	if err != nil && ErrorIs404(err) {
		return fmt.Errorf("validation error: client with id %s does not exist", clientId)
	} else if err != nil {
		return err
	}

	if openidClient.BearerOnly {
		return fmt.Errorf("validation error: client with id %s uses access type BEARER-ONLY which does not use scopes", clientId)
	}

	allOpenidClientScopes, err := keycloakClient.ListOpenidClientScopesWithFilter(realmId, includeOpenidClientScopesMatchingNames(scopeNames))
	if err != nil {
		return err
	}

	var attachedClientScopes []*OpenidClientScope
	var duplicateScopeAssignmentErrorMessage string
	switch t {
	case "optional":
		attachedDefaultClientScopes, err := keycloakClient.GetOpenidClientDefaultScopes(realmId, clientId)
		if err != nil {
			return err
		}
		attachedClientScopes = append(attachedClientScopes, attachedDefaultClientScopes...)
		duplicateScopeAssignmentErrorMessage = "validation error: scope %s is already attached to client as a default scope"
	case "default":
		attachedOptionalClientScopes, err := keycloakClient.GetOpenidClientOptionalScopes(realmId, clientId)
		if err != nil {
			return err
		}
		attachedClientScopes = append(attachedClientScopes, attachedOptionalClientScopes...)
		duplicateScopeAssignmentErrorMessage = "validation error: scope %s is already attached to client as an optional scope"
	}

	for _, openidClientScope := range allOpenidClientScopes {
		for _, attachedClientScope := range attachedClientScopes {
			if openidClientScope.Id == attachedClientScope.Id {
				return fmt.Errorf(duplicateScopeAssignmentErrorMessage, attachedClientScope.Name)
			}
		}

		err := keycloakClient.put(fmt.Sprintf("/realms/%s/clients/%s/%s-client-scopes/%s", realmId, clientId, t, openidClientScope.Id), nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (keycloakClient *KeycloakClient) AttachOpenidClientDefaultScopes(realmId, clientId string, scopeNames []string) error {
	return keycloakClient.attachOpenidClientScopes(realmId, clientId, "default", scopeNames)
}

func (keycloakClient *KeycloakClient) AttachOpenidClientOptionalScopes(realmId, clientId string, scopeNames []string) error {
	return keycloakClient.attachOpenidClientScopes(realmId, clientId, "optional", scopeNames)
}

func (keycloakClient *KeycloakClient) detachOpenidClientScopes(realmId, clientId, t string, scopeNames []string) error {
	allOpenidClientScopes, err := keycloakClient.ListOpenidClientScopesWithFilter(realmId, includeOpenidClientScopesMatchingNames(scopeNames))
	if err != nil {
		return err
	}

	for _, openidClientScope := range allOpenidClientScopes {
		err := keycloakClient.delete(fmt.Sprintf("/realms/%s/clients/%s/%s-client-scopes/%s", realmId, clientId, t, openidClientScope.Id), nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (keycloakClient *KeycloakClient) DetachOpenidClientDefaultScopes(realmId, clientId string, scopeNames []string) error {
	return keycloakClient.detachOpenidClientScopes(realmId, clientId, "default", scopeNames)
}

func (keycloakClient *KeycloakClient) DetachOpenidClientOptionalScopes(realmId, clientId string, scopeNames []string) error {
	return keycloakClient.detachOpenidClientScopes(realmId, clientId, "optional", scopeNames)
}
