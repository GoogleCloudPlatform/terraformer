package keycloak

import (
	"encoding/json"
	"fmt"
)

type OpenidClientAuthorizationTimePolicy struct {
	Id               string `json:"id,omitempty"`
	RealmId          string `json:"-"`
	ResourceServerId string `json:"-"`
	Name             string `json:"name"`
	DecisionStrategy string `json:"decisionStrategy"`
	Logic            string `json:"logic"`
	Type             string `json:"type"`
	NotBefore        string `json:"notBefore"`
	NotOnOrAfter     string `json:"notOnOrAfter"`
	DayMonth         string `json:"dayMonth"`
	DayMonthEnd      string `json:"dayMonthEnd"`
	Month            string `json:"month"`
	MonthEnd         string `json:"monthEnd"`
	Year             string `json:"year"`
	YearEnd          string `json:"yearEnd"`
	Hour             string `json:"hour"`
	HourEnd          string `json:"hourEnd"`
	Minute           string `json:"minute"`
	MinuteEnd        string `json:"minuteEnd"`
	Description      string `json:"description"`
}

func (keycloakClient *KeycloakClient) NewOpenidClientAuthorizationTimePolicy(policy *OpenidClientAuthorizationTimePolicy) error {
	body, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/time", policy.RealmId, policy.ResourceServerId), policy)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &policy)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) UpdateOpenidClientAuthorizationTimePolicy(policy *OpenidClientAuthorizationTimePolicy) error {
	err := keycloakClient.put(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/time/%s", policy.RealmId, policy.ResourceServerId, policy.Id), policy)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) DeleteOpenidClientAuthorizationTimePolicy(realmId, resourceServerId, policyId string) error {
	return keycloakClient.delete(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/time/%s", realmId, resourceServerId, policyId), nil)
}

func (keycloakClient *KeycloakClient) GetOpenidClientAuthorizationTimePolicy(realmId, resourceServerId, policyId string) (*OpenidClientAuthorizationTimePolicy, error) {

	policy := OpenidClientAuthorizationTimePolicy{
		Id:               policyId,
		ResourceServerId: resourceServerId,
		RealmId:          realmId,
	}
	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy/time/%s", realmId, resourceServerId, policyId), &policy, nil)
	if err != nil {
		return nil, err
	}

	return &policy, nil
}
