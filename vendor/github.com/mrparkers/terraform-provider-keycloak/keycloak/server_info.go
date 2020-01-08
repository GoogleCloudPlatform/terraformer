package keycloak

type ComponentType struct {
	Id string `json:"id"`
}

type Provider struct {
}

type ProviderType struct {
	Internal  bool                `json:"internal"`
	Providers map[string]Provider `json:"providers"`
}

type ServerInfo struct {
	ComponentTypes map[string][]ComponentType `json:"componentTypes"`
	ProviderTypes  map[string]ProviderType    `json:"providers"`
	Themes         map[string][]Theme         `json:"themes"`
}

type Theme struct {
	Name    string   `json:"name"`
	Locales []string `json:"locales,omitempty"`
}

func (serverInfo *ServerInfo) ThemeIsInstalled(t, themeName string) bool {
	if themes, ok := serverInfo.Themes[t]; ok {
		for _, theme := range themes {
			if theme.Name == themeName {
				return true
			}
		}
	}

	return false
}

func (serverInfo *ServerInfo) ComponentTypeIsInstalled(componentType, componentTypeId string) bool {
	if componentTypes, ok := serverInfo.ComponentTypes[componentType]; ok {
		for _, componentType := range componentTypes {
			if componentType.Id == componentTypeId {
				return true
			}
		}
	}

	return false
}

func (serverInfo *ServerInfo) getInstalledProvidersNames(providerType string) []string {
	providers := serverInfo.ProviderTypes[providerType].Providers
	keys := make([]string, 0, len(providers))
	for p := range providers {
		keys = append(keys, p)
	}
	return keys
}

func (serverInfo *ServerInfo) providerInstalled(providerType, providerName string) bool {
	providers := serverInfo.ProviderTypes[providerType].Providers
	for p := range providers {
		if p == providerName {
			return true
		}
	}
	return false
}

func (keycloakClient *KeycloakClient) GetServerInfo() (*ServerInfo, error) {
	var serverInfo ServerInfo

	err := keycloakClient.get("/serverinfo", &serverInfo, nil)
	if err != nil {
		return nil, err
	}

	return &serverInfo, nil
}
