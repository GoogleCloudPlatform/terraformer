package keycloak

import (
	"fmt"
)

type OpenidClientAuthorizationPolicy struct {
	Id               string   `json:"id,omitempty"`
	RealmId          string   `json:"-"`
	ResourceServerId string   `json:"-"`
	Name             string   `json:"name"`
	Owner            string   `json:"owner"`
	DecisionStrategy string   `json:"decisionStrategy"`
	Logic            string   `json:"logic"`
	Policies         []string `json:"policies"`
	Resources        []string `json:"resources"`
	Scopes           []string `json:"scopes"`
	Type             string   `json:"type"`
}

func (keycloakClient *KeycloakClient) GetClientAuthorizationPolicyByName(realmId, resourceServerId, name string) (*OpenidClientAuthorizationPolicy, error) {
	policies := []OpenidClientAuthorizationPolicy{}
	params := map[string]string{"name": name}
	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s/authz/resource-server/policy", realmId, resourceServerId), &policies, params)
	if err != nil {
		return nil, err
	}
	policy := policies[0]
	policy.RealmId = realmId
	policy.ResourceServerId = resourceServerId
	policy.Name = name
	return &policy, nil
}
