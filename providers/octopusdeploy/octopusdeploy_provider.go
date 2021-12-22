package octopusdeploy

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type OctopusDeployProvider struct { //nolint
	terraformutils.Provider
	address string
	apiKey  string
}

func (p *OctopusDeployProvider) Init(args []string) error {
	if args[0] != "" {
		p.address = args[0]
	} else {
		if address := os.Getenv("OCTOPUS_CLI_SERVER"); address != "" {
			p.address = address
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
				"address": p.address,
			},
		},
	}
}

func (OctopusDeployProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *OctopusDeployProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"accounts": &GenericGenerator{APIService: "accounts"},
		// "channels":      &GenericGenerator{APIService: "channels"},
		"certificates":        &GenericGenerator{APIService: "certificates"},
		"environments":        &GenericGenerator{APIService: "environments"},
		"feeds":               &GenericGenerator{APIService: "feeds"},
		"libraryvariablesets": &GenericGenerator{APIService: "libraryvariablesets"},
		"lifecycles":          &GenericGenerator{APIService: "lifecycles"},
		"projects":            &GenericGenerator{APIService: "projects"},
		"projectgroups":       &GenericGenerator{APIService: "projectgroups"},
		"projecttriggers":     &GenericGenerator{APIService: "projecttriggers"},
		"tagsets":             &GenericGenerator{APIService: "tagsets"},
		// "variables":           &GenericGenerator{APIService: "variables"},
	}
}

func (p *OctopusDeployProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("octopusdeploy: " + serviceName + " not supported service, see list sub-command")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"api_key": p.apiKey,
		"address": p.address,
	})

	return nil
}

// GetConfig return map of provider config for OctopusDeployProvider
func (p *OctopusDeployProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"api_key": cty.StringVal(p.apiKey),
		"address": cty.StringVal(p.address),
	})
}

func (p *OctopusDeployProvider) GetBasicConfig() cty.Value {
	return p.GetConfig()
}
