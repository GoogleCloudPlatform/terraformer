// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vultr

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type VultrProvider struct { //nolint
	terraformutils.Provider
	apiKey string
}

func (p *VultrProvider) Init(args []string) error {
	if os.Getenv("VULTR_API_KEY") == "" {
		return errors.New("set VULTR_API_KEY env var")
	}
	p.apiKey = os.Getenv("VULTR_API_KEY")

	return nil
}

func (p *VultrProvider) GetName() string {
	return "vultr"
}

func (p *VultrProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (VultrProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *VultrProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"bare_metal_server": &BareMetalServerGenerator{},
		"block_storage":     &BlockStorageGenerator{},
		"dns_domain":        &DNSDomainGenerator{},
		"firewall_group":    &FirewallGroupGenerator{},
		"network":           &NetworkGenerator{},
		"reserved_ip":       &ReservedIPGenerator{},
		"server":            &ServerGenerator{},
		"snapshot":          &SnapshotGenerator{},
		"ssh_key":           &SSHKeyGenerator{},
		"startup_script":    &StartupScriptGenerator{},
		"user":              &UserGenerator{},
	}
}

func (p *VultrProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("vultr: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"api_key": p.apiKey,
	})
	return nil
}
