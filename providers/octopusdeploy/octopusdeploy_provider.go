package octopusdeploy

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/provider_wrapper"
)

type OctopusDeployProvider struct {
	terraform_utils.Provider
	server string
	apiKey string
}

func (p *OctopusDeployProvider) Init(args []string) error {
	if args[0] != "" {
		p.server = args[0]
	} else {
		if server := os.Getenv("OCTOPUS_CLI_SERVER"); server != "" {
			p.server = server
		} else {
			return errors.New("server requirement")
		}
	}

	if args[1] != "" {
		p.apiKey = args[1]
	} else {
		if apiKey := os.Getenv("OCTOPUS_CLI_API_KEY"); apiKey != "" {
			p.apiKey = apiKey
		} else {
			return errors.New("api-key requirement")
		}
	}

	return nil
}

func (p *OctopusDeployProvider) GetName() string {
	return "octopusdeploy"
}

func (p *OctopusDeployProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"octopusdeploy": map[string]interface{}{
				"version": provider_wrapper.GetProviderVersion(p.GetName()),
			},
		},
	}
}

func (OctopusDeployProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *OctopusDeployProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"tagsets": &TagSetGenerator{},
	}
}

func (p *OctopusDeployProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("octopusdeploy: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"apiKey": p.apiKey,
		"server": p.server,
	})

	return nil
}
