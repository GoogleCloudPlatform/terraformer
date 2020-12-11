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

package openstack

import (
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/pkg/errors"
)

type OpenStackProvider struct { //nolint
	terraformutils.Provider
	region string
}

func (p OpenStackProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p OpenStackProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"openstack": map[string]interface{}{
				"region": p.region,
			},
		},
	}
}

// check projectName in env params
func (p *OpenStackProvider) Init(args []string) error {
	p.region = args[0]
	// terraform work with env param OS_REGION_NAME
	err := os.Setenv("OS_REGION_NAME", p.region)
	if err != nil {
		return err
	}
	return nil
}

func (p *OpenStackProvider) GetName() string {
	return "openstack"
}

func (p *OpenStackProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("openstack: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"region": p.region,
	})
	return nil
}

// GetOpenStackSupportService return map of support service for OpenStack
func (p *OpenStackProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"blockstorage": &BlockStorageGenerator{},
		"compute":      &ComputeGenerator{},
		"networking":   &NetworkingGenerator{},
	}
}
