package keycloak

import (
	"fmt"
)

// this is only used when creating an execution on a flow.
// other fields can be provided to the API but they are ignored
// POST /realms/${realmId}/authentication/flows/${flowAlias}/executions/execution
type authenticationExecutionCreate struct {
	Provider string `json:"provider"` //authenticator of the execution
}

type authenticationExecutionRequirementUpdate struct {
	RealmId         string `json:"-"`
	ParentFlowAlias string `json:"-"`
	Id              string `json:"id"`
	Requirement     string `json:"requirement"`
}

// this type is returned by GET /realms/${realmId}/authentication/flows/${flowAlias}/executions
type AuthenticationExecution struct {
	Id                   string `json:"id"`
	RealmId              string `json:"-"`
	ParentFlowAlias      string `json:"-"`
	Authenticator        string `json:"authenticator"` //can be any authenticator from GET realms/{realm}/authentication/authenticator-providers OR GET realms/{realm}/authentication/client-authenticator-providers OR GET realms/{realm}/authentication/form-action-providers
	AuthenticationConfig string `json:"authenticationConfig"`
	AuthenticationFlow   bool   `json:"authenticationFlow"`
	FlowId               string `json:"flowId"`
	ParentFlowId         string `json:"parentFlow"`
	Priority             int    `json:"priority"`
	Requirement          string `json:"requirement"`
}

// another model is used for GET /realms/${realmId}/authentication/executions/${executionId}, but I am going to try to avoid using this API
type AuthenticationExecutionInfo struct {
	Id                   string `json:"id"`
	RealmId              string `json:"-"`
	ParentFlowAlias      string `json:"-"`
	Alias                string `json:"alias"`
	AuthenticationConfig string `json:"authenticationConfig"`
	AuthenticationFlow   bool   `json:"authenticationFlow"`
	Configurable         bool   `json:"configurable"`
	FlowId               string `json:"flowId"`
	Index                int    `json:"index"`
	Level                int    `json:"level"`
	ProviderId           string `json:"providerId"`
	Requirement          string `json:"requirement"`
}

type AuthenticationExecutionList []*AuthenticationExecutionInfo

func (list AuthenticationExecutionList) Len() int {
	return len(list)
}

func (list AuthenticationExecutionList) Less(i, j int) bool {
	return list[i].Index < list[j].Index
}

func (list AuthenticationExecutionList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (keycloakClient *KeycloakClient) ListAuthenticationExecutions(realmId, parentFlowAlias string) (AuthenticationExecutionList, error) {
	var authenticationExecutions []*AuthenticationExecutionInfo

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/authentication/flows/%s/executions", realmId, parentFlowAlias), &authenticationExecutions, nil)
	if err != nil {
		return nil, err
	}

	return authenticationExecutions, err
}

func (keycloakClient *KeycloakClient) NewAuthenticationExecution(execution *AuthenticationExecution) error {
	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/authentication/flows/%s/executions/execution", execution.RealmId, execution.ParentFlowAlias), &authenticationExecutionCreate{Provider: execution.Authenticator})
	if err != nil {
		return err
	}

	execution.Id = getIdFromLocationHeader(location)

	err = keycloakClient.UpdateAuthenticationExecution(execution)
	if err != nil {
		return err
	}

	return nil
}

func (keycloakClient *KeycloakClient) GetAuthenticationExecution(realmId, parentFlowAlias, id string) (*AuthenticationExecution, error) {
	var authenticationExecution AuthenticationExecution

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/authentication/executions/%s", realmId, id), &authenticationExecution, nil)
	if err != nil {
		return nil, err
	}

	authenticationExecution.RealmId = realmId
	authenticationExecution.ParentFlowAlias = parentFlowAlias

	return &authenticationExecution, nil
}

func (keycloakClient *KeycloakClient) UpdateAuthenticationExecution(execution *AuthenticationExecution) error {
	authenticationExecutionUpdateRequirement := &authenticationExecutionRequirementUpdate{
		RealmId:         execution.RealmId,
		ParentFlowAlias: execution.ParentFlowAlias,
		Id:              execution.Id,
		Requirement:     execution.Requirement,
	}
	return keycloakClient.UpdateAuthenticationExecutionRequirement(authenticationExecutionUpdateRequirement)
}

func (keycloakClient *KeycloakClient) UpdateAuthenticationExecutionRequirement(executionRequirementUpdate *authenticationExecutionRequirementUpdate) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/authentication/flows/%s/executions", executionRequirementUpdate.RealmId, executionRequirementUpdate.ParentFlowAlias), executionRequirementUpdate)
}

func (keycloakClient *KeycloakClient) DeleteAuthenticationExecution(realmId, id string) error {
	err := keycloakClient.delete(fmt.Sprintf("/realms/%s/authentication/executions/%s", realmId, id), nil)
	if err != nil {
		// For whatever reason, this fails sometimes with a 500 during acceptance tests. try again
		return keycloakClient.delete(fmt.Sprintf("/realms/%s/authentication/executions/%s", realmId, id), nil)
	}

	return nil
}

func (keycloakClient *KeycloakClient) RaiseAuthenticationExecutionPriority(realmId, id string) error {
	_, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/authentication/executions/%s/raise-priority", realmId, id), nil)
	if err != nil {
		return err
	}
	return nil
}

func (keycloakClient *KeycloakClient) LowerAuthenticationExecutionPriority(realmId, id string) error {
	_, _, err := keycloakClient.post(fmt.Sprintf("/realms/%s/authentication/executions/%s/lower-priority", realmId, id), nil)
	if err != nil {
		return err
	}
	return nil
}
