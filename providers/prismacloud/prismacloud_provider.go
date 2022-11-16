package prismacloud

import (
	"errors"
	"github.com/zclconf/go-cty/cty"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type PrismaCloudProvider struct { //nolint
	terraformutils.Provider
	configFilePath string
}

func (p *PrismaCloudProvider) Init([]string) error {
	if os.Getenv("CONFIG_FILE_PATH") == "" {
		return errors.New("set CONFIG_FILE_PATH env var")
	}
	p.configFilePath = os.Getenv("CONFIG_FILE_PATH")

	return nil
}

func (p *PrismaCloudProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"config_file_path": p.configFilePath,
	})
	return nil
}

func (p *PrismaCloudProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{})
}

func (p *PrismaCloudProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (p *PrismaCloudProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *PrismaCloudProvider) GetName() string {
	return "prismacloud"
}

func (p *PrismaCloudProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"policy": &PolicyGenerator{},
	}
}
