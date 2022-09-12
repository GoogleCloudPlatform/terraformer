package authlete

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type AuthleteProvider struct { // nolint
	terraformutils.Provider
	apiServer string
	soKey     string
	soSecret  string
	apiKey    string
	apiSecret string
}

func (p *AuthleteProvider) GetResourceConnections() map[string]map[string][]string {
	toReturn := make(map[string]map[string][]string)
	serviceMap := make(map[string][]string)
	serviceMap["authlete_service"] = []string{"id", "api_secret"}
	serviceMap["authlete_client"] = []string{"id", "client_secret"}
	toReturn["service"] = serviceMap
	return toReturn
}

func (p *AuthleteProvider) GetProviderData(arg ...string) map[string]interface{} {
	authleteConfig := map[string]interface{}{}

	authleteConfig["api_server"] = p.apiServer
	authleteConfig["service_owner_key"] = p.soKey
	authleteConfig["service_owner_secret"] = p.soSecret
	authleteConfig["api_key"] = p.apiKey
	authleteConfig["api_secret"] = p.apiSecret

	return map[string]interface{}{
		"provider": map[string]interface{}{
			"authlete": authleteConfig,
		},
	}
}

func (p *AuthleteProvider) Init(args []string) error {
	p.apiServer = args[0]
	p.soKey = args[1]
	p.soSecret = args[2]
	p.apiKey = args[3]
	p.apiSecret = args[4]
	return nil
}

func (p *AuthleteProvider) GetName() string {
	return "authlete"
}

func (p *AuthleteProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"api_server":           cty.StringVal(p.apiServer),
		"service_owner_key":    cty.StringVal(p.soKey),
		"service_owner_secret": cty.StringVal(p.soSecret),
		"api_key":              cty.StringVal(p.apiKey),
		"api_secret":           cty.StringVal(p.apiSecret),
	})
}

func (p *AuthleteProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"api_server":           p.apiServer,
		"service_owner_key":    p.soKey,
		"service_owner_secret": p.soSecret,
		"api_key":              p.apiKey,
		"api_secret":           p.apiSecret,
	})
	return nil
}

func (p *AuthleteProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"authlete_service": &ServiceGenerator{},
		"authlete_client":  &ClientGenerator{},
	}
}
