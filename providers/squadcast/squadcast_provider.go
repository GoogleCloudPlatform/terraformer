package squadcast

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SquadcastProvider struct {
	terraformutils.Provider
	accesstoken  string
	refreshtoken string
}

func (p *SquadcastProvider) Init(args []string) error {

	if accessToken := os.Getenv("SQUADCAST_ACCESS_TOKEN"); accessToken != "" {
		p.accesstoken = os.Getenv("SQUADCAST_ACCESS_TOKEN")
	}
	if p.accesstoken == "" {
		return errors.New("requred Access Token missing")
	}

	if refreshToken := os.Getenv("SQUADCAST_REFRESH_TOKEN"); refreshToken != "" {
		p.refreshtoken = os.Getenv("SQUADCAST_REFRESH_TOKEN")
	}
	if p.refreshtoken == "" {
		return errors.New("requred refresh Token missing")
	}

	return nil
}

func (p *SquadcastProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"access_token":  p.accesstoken,
		"refresh_token": p.refreshtoken,
	})

	return nil
}

func (p *SquadcastProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (p *SquadcastProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *SquadcastProvider) GetName() string {
	return "squadcast"
}

func (p *SquadcastProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"user": &UserGenerator{},
	}
}
