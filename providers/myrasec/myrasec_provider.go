package myrasec

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

//
// MyrasecProvider
//
type MyrasecProvider struct {
	terraformutils.Provider
}

//
// Init
//
func (p *MyrasecProvider) Init(args []string) error {
	return nil
}

//
// GetName
//
func (p *MyrasecProvider) GetName() string {
	return "myrasec"
}

//
// GetProviderData
//
func (p *MyrasecProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

//
// GetResourceConnections
//
func (MyrasecProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

//
// GetSupportedService
//
func (p *MyrasecProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"domain": &DomainGenerator{},
		"dns":    &DNSGenerator{},

		"errorpage": &ErrorPageGenerator{},
	}
}

//
// InitService
//
func (p *MyrasecProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("myrasec: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())

	return nil
}
