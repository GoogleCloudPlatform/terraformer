// Copyright 2021 The Terraformer Authors.
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

package equinixmetal

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type EquinixMetalProvider struct { //nolint
	terraformutils.Provider
	authToken string
	projectID string
}

func (p *EquinixMetalProvider) Init(args []string) error {
	if os.Getenv("PACKET_AUTH_TOKEN") == "" {
		return errors.New("set PACKET_AUTH_TOKEN env var")
	}
	p.authToken = os.Getenv("PACKET_AUTH_TOKEN")

	if os.Getenv("METAL_PROJECT_ID") == "" {
		return errors.New("set METAL_PROJECT_ID env var")
	}
	p.projectID = os.Getenv("METAL_PROJECT_ID")

	return nil
}

func (p *EquinixMetalProvider) GetName() string {
	return "metal"
}

func (p *EquinixMetalProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (EquinixMetalProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *EquinixMetalProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"device":            &DeviceGenerator{},
		"sshkey":            &SSHKeyGenerator{},
		"spotmarketrequest": &SpotMarketRequestGenerator{},
		"volume":            &VolumeGenerator{},
	}
}

func (p *EquinixMetalProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("equinixmetal: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"auth_token": p.authToken,
		"project_id": p.projectID,
	})
	return nil
}
