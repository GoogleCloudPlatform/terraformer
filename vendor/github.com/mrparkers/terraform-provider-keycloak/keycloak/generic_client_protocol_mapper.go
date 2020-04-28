package keycloak

import (
	"fmt"
)

type GenericClientProtocolMapper struct {
	ClientId       string            `json:"-"`
	ClientScopeId  string            `json:"-"`
	Config         map[string]string `json:"config"`
	Id             string            `json:"id,omitempty"`
	Name           string            `json:"name"`
	Protocol       string            `json:"protocol"`
	ProtocolMapper string            `json:"protocolMapper"`
	RealmId        string            `json:"-"`
}

type OpenidClientWithGenericClientProtocolMappers struct {
	OpenidClient
	ProtocolMappers []*GenericClientProtocolMapper
}

func (keycloakClient *KeycloakClient) NewGenericClientProtocolMapper(genericClientProtocolMapper *GenericClientProtocolMapper) error {
	path := protocolMapperPath(genericClientProtocolMapper.RealmId, genericClientProtocolMapper.ClientId, genericClientProtocolMapper.ClientScopeId)

	_, location, err := keycloakClient.post(path, genericClientProtocolMapper)
	if err != nil {
		return err
	}

	genericClientProtocolMapper.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) GetGenericClientProtocolMappers(realmId string, clientId string) (*OpenidClientWithGenericClientProtocolMappers, error) {
	var openidClientWithGenericClientProtocolMappers OpenidClientWithGenericClientProtocolMappers

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/clients/%s", realmId, clientId), &openidClientWithGenericClientProtocolMappers, nil)
	if err != nil {
		return nil, err
	}

	openidClientWithGenericClientProtocolMappers.RealmId = realmId
	openidClientWithGenericClientProtocolMappers.ClientId = clientId

	for _, protocolMapper := range openidClientWithGenericClientProtocolMappers.ProtocolMappers {
		protocolMapper.RealmId = realmId
		protocolMapper.ClientId = clientId
	}

	return &openidClientWithGenericClientProtocolMappers, nil

}

func (keycloakClient *KeycloakClient) GetGenericClientProtocolMapper(realmId string, clientId string, clientScopeId string, mapperId string) (*GenericClientProtocolMapper, error) {
	var genericClientProtocolMapper GenericClientProtocolMapper

	err := keycloakClient.get(individualProtocolMapperPath(realmId, clientId, clientScopeId, mapperId), &genericClientProtocolMapper, nil)
	if err != nil {
		return nil, err
	}

	// these values are not provided by the keycloak API
	genericClientProtocolMapper.ClientId = clientId
	genericClientProtocolMapper.ClientScopeId = clientScopeId
	genericClientProtocolMapper.RealmId = realmId

	return &genericClientProtocolMapper, nil
}

func (keycloakClient *KeycloakClient) UpdateGenericClientProtocolMapper(genericClientProtocolMapper *GenericClientProtocolMapper) error {
	path := individualProtocolMapperPath(genericClientProtocolMapper.RealmId, genericClientProtocolMapper.ClientId, genericClientProtocolMapper.ClientScopeId, genericClientProtocolMapper.Id)

	return keycloakClient.put(path, genericClientProtocolMapper)
}

func (keycloakClient *KeycloakClient) DeleteGenericClientProtocolMapper(realmId string, clientId string, clientScopeId string, mapperId string) error {
	return keycloakClient.delete(individualProtocolMapperPath(realmId, clientId, clientScopeId, mapperId), nil)
}

func (mapper *GenericClientProtocolMapper) Validate(keycloakClient *KeycloakClient) error {
	if mapper.ClientId == "" && mapper.ClientScopeId == "" {
		return fmt.Errorf("validation error: one of ClientId or ClientScopeId must be set")
	}
	if mapper.ClientId != "" && mapper.ClientScopeId != "" {
		return fmt.Errorf("validation error: only one of ClientId or ClientScopeId must be set")
	}

	protocolMappers, err := keycloakClient.listGenericProtocolMappers(mapper.RealmId, mapper.ClientId, mapper.ClientScopeId)
	if err != nil {
		return err
	}

	for _, protocolMapper := range protocolMappers {
		if protocolMapper.Name == mapper.Name {
			return fmt.Errorf("validation error: a protocol mapper with name %s already exists for this client", mapper.Name)
		}
	}

	return nil
}
