package keycloak

import "fmt"

type RealmEventsConfig struct {
	AdminEventsDetailsEnabled bool     `json:"adminEventsDetailsEnabled"`
	AdminEventsEnabled        bool     `json:"adminEventsEnabled"`
	EnabledEventTypes         []string `json:"enabledEventTypes"`
	EventsEnabled             bool     `json:"eventsEnabled"`
	EventsExpiration          int      `json:"eventsExpiration"`
	EventsListeners           []string `json:"eventsListeners,omitempty"`
}

func (keycloakClient *KeycloakClient) GetRealmEventsConfig(realmId string) (*RealmEventsConfig, error) {
	var realmEventsConfig RealmEventsConfig

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/events/config", realmId), &realmEventsConfig, nil)
	if err != nil {
		return nil, err
	}

	return &realmEventsConfig, nil
}

func (keycloakClient *KeycloakClient) UpdateRealmEventsConfig(realmId string, realmEventsConfig *RealmEventsConfig) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s/events/config", realmId), realmEventsConfig)
}
