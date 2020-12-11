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

package linode

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type LinodeProvider struct { //nolint
	terraformutils.Provider
	token string
}

func (p *LinodeProvider) Init(args []string) error {
	if os.Getenv("LINODE_TOKEN") == "" {
		return errors.New("set LINODE_TOKEN env var")
	}
	p.token = os.Getenv("LINODE_TOKEN")

	return nil
}

func (p *LinodeProvider) GetName() string {
	return "linode"
}

func (p *LinodeProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (LinodeProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *LinodeProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"domain":       &DomainGenerator{},
		"image":        &ImageGenerator{},
		"instance":     &InstanceGenerator{},
		"nodebalancer": &NodeBalancerGenerator{},
		"rdns":         &RDNSGenerator{},
		"sshkey":       &SSHKeyGenerator{},
		"stackscript":  &StackScriptGenerator{},
		"token":        &TokenGenerator{},
		"volume":       &VolumeGenerator{},
	}
}

func (p *LinodeProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("linode: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"token": p.token,
	})
	return nil
}
