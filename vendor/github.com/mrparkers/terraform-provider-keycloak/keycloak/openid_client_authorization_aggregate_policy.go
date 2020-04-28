package keycloak

import (
	"encoding/json"
	"fmt"
)

type OpenidClientAuthorizationAggregatePolicy struct {
	Id               string   `json:"id,omitempty"`
	RealmId          string   `json:"-"`
	ResourceServerId string   `json:"-"`
	Name             string   `json:"name"`
	DecisionStrategy string   `json:"decisionStrategy"`
	Logic            string   `json:"logic"`
	Policies         []string `json:"policies"`
	Type             string   `json:"type"`
	Description      string   `json:"description"`
}

func (keycloakClient *KeycloakClient) NewOpenidClientAuthorizationAggregatePolicy(policy *OpenidClientAuthorizationAggregatePolicy) error {
	body, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/aggregate", policy.RealmId, policy.ResourceServerId), policy)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &policy)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) UpdateOpenidClientAuthorizationAggregatePolicy(policy *OpenidClientAuthorizationAggregatePolicy) error {
	err := keycloakClient.put(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/aggregate/%s", policy.RealmId, policy.ResourceServerId, policy.Id), policy)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) DeleteOpenidClientAuthorizationAggregatePolicy(realmId, resourceServerId, policyId string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/aggregate/%s", realmId, resourceServerId, policyId), nil)
}

func (keycloakClient *KeycloakClient) GetOpenidClientAuthorizationAggregatePolicy(realmId, resourceServerId, policyId string) (*OpenidClientAuthorizationAggregatePolicy, error) {

	policy := OpenidClientAuthorizationAggregatePolicy{
		Id:               policyId,
		ResourceServerId: resourceServerId,
		RealmId:          realmId,
	}

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/aggregate/%s", realmId, resourceServerId, policyId), &policy, nil)
	if err != nil {
		return nil, err
	}

	var keycloakPolicies []map[string]interface{}
	errTwo := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/%s/associatedPolicies", realmId, resourceServerId, policyId), &keycloakPolicies, nil)
	if errTwo != nil {
		return nil, err
	}

	for i := 0; i < len(keycloakPolicies); i++ {
		policy.Policies = append(policy.Policies, keycloakPolicies[i]["id"].(string))
	}

	return &policy, nil
}
