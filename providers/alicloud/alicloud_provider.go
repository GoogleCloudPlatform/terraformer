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

package alicloud

import (
	"errors"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

// AliCloudProvider Provider for alicloud
type AliCloudProvider struct { //nolint
	terraformutils.Provider
	region  string
	profile string
}

// GetConfig Converts json config to go-cty
func (p *AliCloudProvider) GetConfig() cty.Value {
	args := p.Service.GetArgs()
	profile := args["profile"].(string)
	config, err := LoadConfigFromProfile(profile)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	region := p.region
	if region == "" {
		region = config.RegionID
	}

	var val cty.Value
	if config.RAMRoleArn != "" {
		val = cty.ObjectVal(map[string]cty.Value{
			"region":  cty.StringVal(region),
			"profile": cty.StringVal(profile),
			"assume_role": cty.SetVal([]cty.Value{
				cty.ObjectVal(map[string]cty.Value{
					"role_arn": cty.StringVal(config.RAMRoleArn),
				}),
			}),
		})
	} else {
		val = cty.ObjectVal(map[string]cty.Value{
			"region":  cty.StringVal(region),
			"profile": cty.StringVal(profile),
		})
	}

	return val
}

// GetResourceConnections Gets resource connections for alicloud
func (p AliCloudProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		// TODO: Not implemented
	}
}

// GetProviderData Used for generated HCL2 for the provider
func (p AliCloudProvider) GetProviderData(arg ...string) map[string]interface{} {
	args := p.Service.GetArgs()
	profile := args["profile"].(string)
	config, err := LoadConfigFromProfile(profile)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	region := p.region
	if region == "" {
		region = config.RegionID
	}

	if config.RAMRoleArn != "" {
		return map[string]interface{}{
			"provider": map[string]interface{}{
				"alicloud": map[string]interface{}{
					"region":  region,
					"profile": profile,
					"assume_role": map[string]interface{}{
						"role_arn": config.RAMRoleArn,
					},
				},
			},
		}
	}
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"alicloud": map[string]interface{}{
				"region":  region,
				"profile": profile,
			},
		},
	}
}

// Init Loads up command line arguments in the provider
func (p *AliCloudProvider) Init(args []string) error {
	p.region = args[0]
	p.profile = args[1]
	return nil
}

// GetName Gets name of provider
func (p *AliCloudProvider) GetName() string {
	return "alicloud"
}

// InitService Initializes the AliCloud service
func (p *AliCloudProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("alicloud: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"region":  p.region,
		"profile": p.profile,
	})
	return nil
}

// GetSupportedService Gets a list of all supported services
func (p *AliCloudProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"dns":     &DNSGenerator{},
		"ecs":     &EcsGenerator{},
		"keypair": &KeyPairGenerator{},
		"nat":     &NatGatewayGenerator{},
		"pvtz":    &PvtzGenerator{},
		"ram":     &RAMGenerator{},
		"rds":     &RdsGenerator{},
		"sg":      &SgGenerator{},
		"slb":     &SlbGenerator{},
		"vpc":     &VpcGenerator{},
		"vswitch": &VSwitchGenerator{},
	}
}
