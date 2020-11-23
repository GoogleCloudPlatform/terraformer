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
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type MikrotikProvider struct { //nolint
	terraformutils.Provider
	host     string
	user     string
	password string
}

func (p *MikrotikProvider) Init(args []string) error {
	if os.Getenv("MIKROTIK_HOST") == "" {
		return errors.New("set MIKROTIK_HOST env var")
	}
	p.host = os.Getenv("MIKROTIK_HOST")

	if os.Getenv("MIKROTIK_USER") == "" {
		return errors.New("set MIKROTIK_USER env var")
	}
	p.user = os.Getenv("MIKROTIK_USER")

	if os.Getenv("MIKROTIK_PASSWORD") == "" {
		return errors.New("set MIKROTIK_PASSWORD env var")
	}
	p.password = os.Getenv("MIKROTIK_PASSWORD")

	return nil
}

func (p *MikrotikProvider) GetName() string {
	return "mikrotik"
}

func (p *MikrotikProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"mikrotik": map[string]interface{}{
				"host": p.host,
				"user": p.user,
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
		"host":     p.host,
		"user":     p.user,
		"password": p.password,
	})
	return nil
}
