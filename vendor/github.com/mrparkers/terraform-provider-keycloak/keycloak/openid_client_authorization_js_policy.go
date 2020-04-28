package keycloak

import (
	"encoding/json"
	"fmt"
)

type OpenidClientAuthorizationJSPolicy struct {
	Id               string `json:"id,omitempty"`
	RealmId          string `json:"-"`
	ResourceServerId string `json:"-"`
	Name             string `json:"name"`
	DecisionStrategy string `json:"decisionStrategy"`
	Logic            string `json:"logic"`
	Type             string `json:"type"`
	Code             string `json:"code"`
	Description      string `json:"description"`
}

func (keycloakClient *KeycloakClient) NewOpenidClientAuthorizationJSPolicy(policy *OpenidClientAuthorizationJSPolicy) error {
	body, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/js", policy.RealmId, policy.ResourceServerId), policy)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &policy)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) UpdateOpenidClientAuthorizationJSPolicy(policy *OpenidClientAuthorizationJSPolicy) error {
	err := keycloakClient.put(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/js/%s", policy.RealmId, policy.ResourceServerId, policy.Id), policy)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) DeleteOpenidClientAuthorizationJSPolicy(realmId, resourceServerId, policyId string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/js/%s", realmId, resourceServerId, policyId), nil)
}

func (keycloakClient *KeycloakClient) GetOpenidClientAuthorizationJSPolicy(realmId, resourceServerId, policyId string) (*OpenidClientAuthorizationJSPolicy, error) {

	policy := OpenidClientAuthorizationJSPolicy{
		Id:               policyId,
		ResourceServerId: resourceServerId,
		RealmId:          realmId,
	}
	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/js/%s", realmId, resourceServerId, policyId), &policy, nil)
	if err != nil {
		return nil, err
	}

	return &policy, nil
}
