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
package hashicups

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type HashicupsProvider struct {
	terraformutils.Provider
	username string
	password string
}

func (p HashicupsProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p HashicupsProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"hashicups": map[string]interface{}{
				"username": p.username,
				"password": p.password,
			},
		},
	}
}

func (p *HashicupsProvider) Init(args []string) error {
	p.username = args[0]
	p.password = args[1]
	os.Setenv("HASHICUPS_USERNAME", p.username)
	os.Setenv("HASHICUPS_PASSWORD", p.password)
	return nil
}

func (p *HashicupsProvider) GetName() string {
	return "hashicups"
}

func (p *HashicupsProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"username": p.username,
		"password": p.password,
	})
	return nil
}

func (p *HashicupsProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"order": &OrderGenerator{},
	}
}
