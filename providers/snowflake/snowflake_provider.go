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

package snowflake

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

type SnowflakeProvider struct {
	terraform_utils.Provider
	account      string
	username     string
	browser_auth string
	region       string
	role         string
}

func (p *SnowflakeProvider) GetName() string {
	return "snowflake-provider"
}

func (p *SnowflakeProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"snowflake": map[string]interface{}{
				"account":      p.account,
				"username":     p.username,
				"browser_auth": p.browser_auth,
				"region":       p.region,
				"role":         p.role,
			},
		},
	}
}

func (SnowflakeProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *SnowflakeProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"snowflake_database": &DatabaseGenerator{},
	}
}

func (p *SnowflakeProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("snowflake: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"account":      p.account,
		"username":     p.username,
		"browser_auth": p.browser_auth,
		"region":       p.region,
		"role":         p.role,
	})
	return nil
}
