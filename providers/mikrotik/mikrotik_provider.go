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

package mikrotik

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/ddelnano/terraform-provider-mikrotik/client"
)

type MikrotikProvider struct { //nolint
	terraformutils.Provider
	client.Mikrotik
}

func (p *MikrotikProvider) Init(args []string) error {
	// The mikrotik provider gets its credentials through environment variables
	// and therefore nothing needs to be done here
	return nil
}

func (p *MikrotikProvider) GetName() string {
	return "mikrotik"
}

func (p *MikrotikProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"mikrotik": map[string]interface{}{
				"host": p.Host,
				"user": p.Username,
			},
		},
	}
}

func (MikrotikProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *MikrotikProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"dhcp_lease": &DhcpLeaseGenerator{},
	}
}

func (p *MikrotikProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("mikrotik: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"host":           p.Host,
		"user":           p.Username,
		"password":       p.Password,
		"tls":            p.TLS,
		"ca_certificate": p.CA,
		"insecure":       p.Insecure,
	})
	return nil
}
