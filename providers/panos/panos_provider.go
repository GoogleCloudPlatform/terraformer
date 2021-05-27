// Copyright 2018 The Terraformer Authors.
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

package panos

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type PanosProvider struct { //nolint
	terraformutils.Provider
	hostname string
	username string
	password string
	vsys     string
}

func (p *PanosProvider) Init(args []string) error {
	p.hostname = args[0]
	p.username = args[1]
	p.password = args[2]
	p.vsys = args[3]

	return nil
}

func (p *PanosProvider) GetName() string {
	return "panos"
}

func (p *PanosProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (p *PanosProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"hostname": cty.StringVal(p.hostname),
		"username": cty.StringVal(p.username),
		"password": cty.StringVal(p.password),
	})
}

func (p *PanosProvider) GetBasicConfig() cty.Value {
	return p.GetConfig()
}

func (p *PanosProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}

	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"hostname": p.hostname,
		"username": p.username,
		"password": p.password,
		"vsys":     p.vsys,
	})

	return nil
}

func (p *PanosProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {

	return map[string]terraformutils.ServiceGenerator{
		"device_config":       &DeviceConfigGenerator{},
		"firewall_networking": &FirewallNetworkingGenerator{},
		"firewall_objects":    &FirewallObjectsGenerator{},
		"firewall_policy":     &FirewallPolicyGenerator{},
	}
}

func (PanosProvider) GetResourceConnections() map[string]map[string][]string {

	return map[string]map[string][]string{}
}
