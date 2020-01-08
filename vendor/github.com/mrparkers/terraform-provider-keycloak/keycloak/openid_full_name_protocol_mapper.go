package keycloak

import (
	"fmt"
	"strconv"
)

type OpenIdFullNameProtocolMapper struct {
	Id            string
	Name          string
	RealmId       string
	ClientId      string
	ClientScopeId string

	AddToIdToken     bool
	AddToAccessToken bool
	AddToUserInfo    bool
}

func (mapper *OpenIdFullNameProtocolMapper) convertToGenericProtocolMapper() *protocolMapper {
	return &protocolMapper{
		Id:             mapper.Id,
		Name:           mapper.Name,
		Protocol:       "openid-connect",
		ProtocolMapper: "oidc-full-name-mapper",
		Config: map[string]string{
			addToIdTokenField:     strconv.FormatBool(mapper.AddToIdToken),
			addToAccessTokenField: strconv.FormatBool(mapper.AddToAccessToken),
			addToUserInfoField:    strconv.FormatBool(mapper.AddToUserInfo),
		},
	}
}

func (protocolMapper *protocolMapper) convertToOpenIdFullNameProtocolMapper(realmId, clientId, clientScopeId string) (*OpenIdFullNameProtocolMapper, error) {
	idTokenClaim, err := strconv.ParseBool(protocolMapper.Config[addToIdTokenField])
	if err != nil {
		return nil, err
	}

	accessTokenClaim, err := strconv.ParseBool(protocolMapper.Config[addToAccessTokenField])
	if err != nil {
		return nil, err
	}

	userinfoTokenClaim, err := strconv.ParseBool(protocolMapper.Config[addToUserInfoField])
	if err != nil {
		return nil, err
	}

	return &OpenIdFullNameProtocolMapper{
		Id:            protocolMapper.Id,
		Name:          protocolMapper.Name,
		RealmId:       realmId,
		ClientId:      clientId,
		ClientScopeId: clientScopeId,

		AddToIdToken:     idTokenClaim,
		AddToAccessToken: accessTokenClaim,
		AddToUserInfo:    userinfoTokenClaim,
	}, nil
}

func (keycloakClient *KeycloakClient) GetOpenIdFullNameProtocolMapper(realmId, clientId, clientScopeId, mapperId string) (*OpenIdFullNameProtocolMapper, error) {
	var protocolMapper *protocolMapper

	err := keycloakClient.get(individualProtocolMapperPath(realmId, clientId, clientScopeId, mapperId), &protocolMapper, nil)
	if err != nil {
		return nil, err
	}

	return protocolMapper.convertToOpenIdFullNameProtocolMapper(realmId, clientId, clientScopeId)
}

func (keycloakClient *KeycloakClient) DeleteOpenIdFullNameProtocolMapper(realmId, clientId, clientScopeId, mapperId string) error {
	return keycloakClient.delete(individualProtocolMapperPath(realmId, clientId, clientScopeId, mapperId), nil)
}

func (keycloakClient *KeycloakClient) NewOpenIdFullNameProtocolMapper(mapper *OpenIdFullNameProtocolMapper) error {
	path := protocolMapperPath(mapper.RealmId, mapper.ClientId, mapper.ClientScopeId)

	_, location, err := keycloakClient.post(path, mapper.convertToGenericProtocolMapper())
	if err != nil {
		return err
	}

	mapper.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) UpdateOpenIdFullNameProtocolMapper(mapper *OpenIdFullNameProtocolMapper) error {
	path := individualProtocolMapperPath(mapper.RealmId, mapper.ClientId, mapper.ClientScopeId, mapper.Id)

	return keycloakClient.put(path, mapper.convertToGenericProtocolMapper())
}

func (keycloakClient *KeycloakClient) ValidateOpenIdFullNameProtocolMapper(mapper *OpenIdFullNameProtocolMapper) error {
	if mapper.ClientId == "" && mapper.ClientScopeId == "" {
		return fmt.Errorf("validation error: one of ClientId or ClientScopeId must be set")
	}

	protocolMappers, err := keycloakClient.listGenericProtocolMappers(mapper.RealmId, mapper.ClientId, mapper.ClientScopeId)
	if err != nil {
		return err
	}

	for _, protocolMapper := range protocolMappers {
		if protocolMapper.Name == mapper.Name && protocolMapper.Id != mapper.Id {
			return fmt.Errorf("validation error: a protocol mapper with name %s already exists for this client", mapper.Name)
		}
	}

	return nil
}
