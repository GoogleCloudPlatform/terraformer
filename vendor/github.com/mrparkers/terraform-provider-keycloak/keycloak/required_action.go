package keycloak

import "fmt"

type RequiredAction struct {
	Id            string              `json:"-"`
	RealmId       string              `json:"-"`
	Alias         string              `json:"alias"`
	Name          string              `json:"name"`
	Enabled       bool                `json:"enabled"`
	DefaultAction bool                `json:"defaultAction"`
	Priority      int                 `json:"priority"`
	Config        map[string][]string `json:"config"`
}

func (requiredActions *RequiredAction) getConfig(val string) string {
	if len(requiredActions.Config[val]) == 0 {
		return ""
	}
	return requiredActions.Config[val][0]
}

func (requiredActions *RequiredAction) getConfigOk(val string) (string, bool) {
	if v, ok := requiredActions.Config[val]; ok {
		return v[0], true
	}
	return "", false
}

func (keycloakClient *KeycloakClient) GetRequiredActions(realmId string) ([]*RequiredAction, error) {
	var requiredActions []*RequiredAction

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/authentication/required-actions", realmId), &requiredActions, nil)
	if err != nil {
		return nil, err
	}

	for _, requiredAction := range requiredActions {
		requiredAction.RealmId = realmId
	}

	return requiredActions, nil
}

func (keycloakClient *KeycloakClient) GetRequiredAction(realmId string, alias string) (*RequiredAction, error) {
	var requiredAction RequiredAction

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/authentication/required-actions/%s", realmId, alias), &requiredAction, nil)
	if err != nil {
		return nil, err
	}
	requiredAction.RealmId = realmId
	return &requiredAction, nil
}

func (keycloakClient *KeycloakClient) CreateRequiredAction(requiredAction *RequiredAction) error {
	requiredAction.Id = fmt.Sprintf("%s/%s", requiredAction.RealmId, requiredAction.Alias)
	return keycloakClient.UpdateRequiredAction(requiredAction)
}

func (keycloakClient *KeycloakClient) UpdateRequiredAction(requiredAction *RequiredAction) error {

	err := keycloakClient.ValidateRequiredAction(requiredAction)
	if err != nil {
		return err
	}

	return keycloakClient.put(fmt.Sprintf("/realms/%s/authentication/required-actions/%s", requiredAction.RealmId, requiredAction.Alias), requiredAction)
}

func (keycloakClient *KeycloakClient) DeleteRequiredAction(realmName string, alias string) error {
	err := keycloakClient.delete(fmt.Sprintf("/realms/%s/authentication/required-actions/%s", realmName, alias), nil)
	if err != nil {
		// For whatever reason, this fails sometimes with a 500 during acceptance tests. try again
		return keycloakClient.delete(fmt.Sprintf("/realms/%s/authentication/required-actions/%s", realmName, alias), nil)
	}

	return nil
}

func (keycloakClient *KeycloakClient) ValidateRequiredAction(requiredAction *RequiredAction) error {
	serverInfo, err := keycloakClient.GetServerInfo()
	if err != nil {
		return err
	}

	if requiredAction.DefaultAction && !requiredAction.Enabled {
		return fmt.Errorf("validation error: a 'default' required action should be enabled, set 'defaultAction' to 'false' or set 'enabled' to 'true'")
	}

	if !serverInfo.providerInstalled("required-action", requiredAction.Alias) {
		return fmt.Errorf("validation error: required action \"%s\" does not exist on the server, installed providers: %s", requiredAction.Alias, serverInfo.getInstalledProvidersNames("required-action"))
	}

	return nil
}
