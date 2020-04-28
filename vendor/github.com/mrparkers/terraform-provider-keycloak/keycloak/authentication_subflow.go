package keycloak

import (
	"errors"
	"fmt"
)

type AuthenticationSubFlow struct {
	Id              string `json:"id,omitempty"`
	Alias           string `json:"alias"`
	RealmId         string `json:"-"`
	ParentFlowAlias string `json:"-"`
	ProviderId      string `json:"providerId"` // "basic-flow" or "client-flow" or form-flow see /keycloak/server-spi/src/main/java/org/keycloak/models/AuthenticationFlowModel.java
	TopLevel        bool   `json:"topLevel"`   // should only be false if this is a subflow
	BuiltIn         bool   `json:"builtIn"`    // this controls whether or not this flow can be edited from the console. it can be updated, but this provider will only set it to `true`
	Description     string `json:"description"`
	//execution part
	Authenticator string `json:"-"` //can be any authenticator see /auth/admin/master/console/#/server-info/providers (not limited to the authenticator spi section) for example could also be part of the form-action spi
	Priority      int    `json:"-"`
	Requirement   string `json:"-"`
}

//each subflow creates a flow and an execution under the covers
type authenticationSubFlowCreate struct {
	Alias       string `json:"alias"`
	Type        string `json:"type"`     //providerId of the flow
	Provider    string `json:"provider"` //authenticator of the execution
	Description string `json:"description"`
}

func (keycloakClient *KeycloakClient) NewAuthenticationSubFlow(authenticationSubFlow *AuthenticationSubFlow) error {
	authenticationSubFlow.TopLevel = false
	authenticationSubFlow.BuiltIn = false
	authenticationSubFlowCreate := &authenticationSubFlowCreate{
		Alias:       authenticationSubFlow.Alias,
		Type:        authenticationSubFlow.ProviderId,    //providerId of the flow
		Provider:    authenticationSubFlow.Authenticator, //seems this can be empty //authenticator of the execution
		Description: authenticationSubFlow.Description,
	}

	_, location, err := keycloakClient.post(fmt.Sprintf("/realms/%s/authentication/flows/%s/executions/flow", authenticationSubFlow.RealmId, authenticationSubFlow.ParentFlowAlias), authenticationSubFlowCreate)
	if err != nil {
		return err
	}
	authenticationSubFlow.Id = getIdFromLocationHeader(location)

	if authenticationSubFlow.Requirement != "DISABLED" {
		return keycloakClient.UpdateAuthenticationSubFlow(authenticationSubFlow)
	}
	return nil
}

func (keycloakClient *KeycloakClient) GetAuthenticationSubFlow(realmId, parentFlowAlias, id string) (*AuthenticationSubFlow, error) {
	var authenticationSubFlow AuthenticationSubFlow
	err := keycloakClient.get(fmt.Sprintf("/realms/%s/authentication/flows/%s", realmId, id), &authenticationSubFlow, nil)
	if err != nil {
		return nil, err
	}
	authenticationSubFlow.RealmId = realmId
	authenticationSubFlow.ParentFlowAlias = parentFlowAlias

	executionId, err := keycloakClient.getExecutionId(&authenticationSubFlow)
	if err != nil {
		return nil, err
	}

	subFlowExecution, err := keycloakClient.GetAuthenticationExecution(realmId, parentFlowAlias, executionId)
	if err != nil {
		return nil, err
	}
	authenticationSubFlow.Authenticator = subFlowExecution.Authenticator
	authenticationSubFlow.Requirement = subFlowExecution.Requirement

	return &authenticationSubFlow, nil
}

func (keycloakClient *KeycloakClient) getExecutionId(authenticationSubFlow *AuthenticationSubFlow) (string, error) {
	list, err := keycloakClient.ListAuthenticationExecutions(authenticationSubFlow.RealmId, authenticationSubFlow.ParentFlowAlias)
	if err != nil {
		return "", err
	}

	for _, ex := range list {
		if ex.FlowId == authenticationSubFlow.Id {
			return ex.Id, nil
		}
	}
	return "", errors.New("no execution id found for subflow")
}

func (keycloakClient *KeycloakClient) UpdateAuthenticationSubFlow(authenticationSubFlow *AuthenticationSubFlow) error {
	authenticationSubFlow.TopLevel = false
	authenticationSubFlow.BuiltIn = false

	err := keycloakClient.put(fmt.Sprintf("/realms/%s/authentication/flows/%s", authenticationSubFlow.RealmId, authenticationSubFlow.Id), authenticationSubFlow)

	if err != nil {
		return err
	}

	executionId, err := keycloakClient.getExecutionId(authenticationSubFlow)
	if err != nil {
		return err
	}

	//update requirement
	authenticationExecutionUpdateRequirement := &authenticationExecutionRequirementUpdate{
		RealmId:         authenticationSubFlow.RealmId,
		ParentFlowAlias: authenticationSubFlow.ParentFlowAlias,
		Id:              executionId,
		Requirement:     authenticationSubFlow.Requirement,
	}
	return keycloakClient.UpdateAuthenticationExecutionRequirement(authenticationExecutionUpdateRequirement)

}

func (keycloakClient *KeycloakClient) DeleteAuthenticationSubFlow(realmId, parentFlowAlias, id string) error {
	authenticationSubFlow := AuthenticationSubFlow{
		Id:              id,
		ParentFlowAlias: parentFlowAlias,
		RealmId:         realmId,
	}
	executionId, err := keycloakClient.getExecutionId(&authenticationSubFlow)
	if err != nil {
		return err
	}

	return keycloakClient.DeleteAuthenticationExecution(authenticationSubFlow.RealmId, executionId)
}

func (keycloakClient *KeycloakClient) RaiseAuthenticationSubFlowPriority(realmId, parentFlowAlias, id string) error {
	authenticationSubFlow := AuthenticationSubFlow{
		Id:              id,
		ParentFlowAlias: parentFlowAlias,
		RealmId:         realmId,
	}
	executionId, err := keycloakClient.getExecutionId(&authenticationSubFlow)
	if err != nil {
		return err
	}

	return keycloakClient.RaiseAuthenticationExecutionPriority(authenticationSubFlow.RealmId, executionId)
}

func (keycloakClient *KeycloakClient) LowerAuthenticationSubFlowPriority(realmId, parentFlowAlias, id string) error {
	authenticationSubFlow := AuthenticationSubFlow{
		Id:              id,
		ParentFlowAlias: parentFlowAlias,
		RealmId:         realmId,
	}
	executionId, err := keycloakClient.getExecutionId(&authenticationSubFlow)
	if err != nil {
		return err
	}

	return keycloakClient.LowerAuthenticationExecutionPriority(authenticationSubFlow.RealmId, executionId)
}
