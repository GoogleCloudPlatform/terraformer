package keycloak

import (
	"fmt"
)

type SamlUserAttributeProtocolMapper struct {
	Id            string
	Name          string
	RealmId       string
	ClientId      string
	ClientScopeId string

	UserAttribute           string
	FriendlyName            string
	SamlAttributeName       string
	SamlAttributeNameFormat string
}

func (mapper *SamlUserAttributeProtocolMapper) convertToGenericProtocolMapper() *protocolMapper {
	return &protocolMapper{
		Id:             mapper.Id,
		Name:           mapper.Name,
		Protocol:       "saml",
		ProtocolMapper: "saml-user-attribute-mapper",
		Config: map[string]string{
			attributeNameField:       mapper.SamlAttributeName,
			attributeNameFormatField: mapper.SamlAttributeNameFormat,
			friendlyNameField:        mapper.FriendlyName,
			userAttributeField:       mapper.UserAttribute,
		},
	}
}

func (protocolMapper *protocolMapper) convertToSamlUserAttributeProtocolMapper(realmId, clientId, clientScopeId string) *SamlUserAttributeProtocolMapper {
	return &SamlUserAttributeProtocolMapper{
		Id:            protocolMapper.Id,
		Name:          protocolMapper.Name,
		RealmId:       realmId,
		ClientId:      clientId,
		ClientScopeId: clientScopeId,

		UserAttribute:           protocolMapper.Config[userAttributeField],
		FriendlyName:            protocolMapper.Config[friendlyNameField],
		SamlAttributeName:       protocolMapper.Config[attributeNameField],
		SamlAttributeNameFormat: protocolMapper.Config[attributeNameFormatField],
	}
}

func (keycloakClient *KeycloakClient) GetSamlUserAttributeProtocolMapper(realmId, clientId, clientScopeId, mapperId string) (*SamlUserAttributeProtocolMapper, error) {
	var protocolMapper *protocolMapper

	err := keycloakClient.get(individualProtocolMapperPath(realmId, clientId, clientScopeId, mapperId), &protocolMapper, nil)
	if err != nil {
		return nil, err
	}

	return protocolMapper.convertToSamlUserAttributeProtocolMapper(realmId, clientId, clientScopeId), nil
}

func (keycloakClient *KeycloakClient) DeleteSamlUserAttributeProtocolMapper(realmId, clientId, clientScopeId, mapperId string) error {
	return keycloakClient.delete(individualProtocolMapperPath(realmId, clientId, clientScopeId, mapperId), nil)
}

func (keycloakClient *KeycloakClient) NewSamlUserAttributeProtocolMapper(mapper *SamlUserAttributeProtocolMapper) error {
	path := protocolMapperPath(mapper.RealmId, mapper.ClientId, mapper.ClientScopeId)

	_, location, err := keycloakClient.post(path, mapper.convertToGenericProtocolMapper())
	if err != nil {
		return err
	}

	mapper.Id = getIdFromLocationHeader(location)

	return nil
}

func (keycloakClient *KeycloakClient) UpdateSamlUserAttributeProtocolMapper(mapper *SamlUserAttributeProtocolMapper) error {
	path := individualProtocolMapperPath(mapper.RealmId, mapper.ClientId, mapper.ClientScopeId, mapper.Id)

	return keycloakClient.put(path, mapper.convertToGenericProtocolMapper())
}

func (keycloakClient *KeycloakClient) ValidateSamlUserAttributeProtocolMapper(mapper *SamlUserAttributeProtocolMapper) error {
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
