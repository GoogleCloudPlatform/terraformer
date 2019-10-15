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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/zclconf/go-cty/cty"
)

type AlicloudProvider struct {
	terraform_utils.Provider
	region  string
	profile string
}

func (p *AlicloudProvider) GetConfig() cty.Value {
	config, _ := LoadConfigFromProfile()

	val := cty.ObjectVal(map[string]cty.Value{
		"region": cty.StringVal(config.RegionId),
		"assume_role": cty.SetVal([]cty.Value{
			cty.ObjectVal(map[string]cty.Value{
				"role_arn": cty.StringVal(config.RamRoleArn),
			}),
		}),
	})

	return val
}

func (p AlicloudProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		// TODO: Not implemented
	}
}

func (p AlicloudProvider) GetProviderData(arg ...string) map[string]interface{} {
	conf, err := LoadConfigFromProfile()
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	region := p.region
	if region != "" {
		region = conf.RegionId
	}

	if conf.RamRoleArn != "" {
		return map[string]interface{}{
			"provider": map[string]interface{}{
				"alicloud": map[string]interface{}{
					"region": region,
					"assume_role": map[string]interface{}{
						"role_arn": conf.RamRoleArn,
					},
				},
			},
		}
	}
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"alicloud": map[string]interface{}{
				"region": region,
			},
		},
	}
}

// check projectName in env params
func (p *AlicloudProvider) Init(args []string) error {
	p.region = args[0]
	p.profile = args[1]
	return nil
}

func (p *AlicloudProvider) GetName() string {
	return "alicloud"
}

func (p *AlicloudProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("alicloud: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"region":  p.region,
		"profile": p.profile,
	})
	return nil
}

func (p *AlicloudProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"ecs": &EcsGenerator{},
	}
}
