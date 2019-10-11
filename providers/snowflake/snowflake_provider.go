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
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

type SnowflakeProvider struct {
	terraform_utils.Provider
	account  string
	username string
	region   string
	role     string
}

func (p *SnowflakeProvider) Init(args []string) error {
	fmt.Println("snowflake_provider init")
	if os.Getenv("SNOWFLAKE_ACCOUNT") == "" {
		return errors.New("set SNOWFLAKE_ACCOUNT env var")
	}
	p.account = os.Getenv("SNOWFLAKE_ACCOUNT")

	if os.Getenv("SNOWFLAKE_USERNAME") == "" {
		return errors.New("set SNOWFLAKE_USERNAME env var")
	}
	p.account = os.Getenv("SNOWFLAKE_USERNAME")

	if os.Getenv("SNOWFLAKE_REGION") == "" {
		return errors.New("set SNOWFLAKE_REGION env var")
	}
	p.region = os.Getenv("SNOWFLAKE_REGION")

	if os.Getenv("SNOWFLAKE_ROLE") == "" {
		return errors.New("set SNOWFLAKE_ROLE env var")
	}
	p.role = os.Getenv("SNOWFLAKE_ROLE")

	return nil
}

func (p *SnowflakeProvider) GetName() string {
	fmt.Println("getname")
	return "snowflake-provider"
}

func (p *SnowflakeProvider) GetProviderData(arg ...string) map[string]interface{} {
	fmt.Println("getproviderdata")
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"snowflake": map[string]interface{}{
				"account":  p.account,
				"username": p.username,
				"region":   p.region,
				"role":     p.role,
			},
		},
	}
}

func (SnowflakeProvider) GetResourceConnections() map[string]map[string][]string {
	fmt.Println("getresourceconnections")
	return map[string]map[string][]string{}
}

func (p *SnowflakeProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	fmt.Println("getsupportedservice")
	return map[string]terraform_utils.ServiceGenerator{
		"snowflake_database": &DatabaseGenerator{},
	}
}

func (p *SnowflakeProvider) InitService(serviceName string) error {
	fmt.Println("initservice")
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("snowflake: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"account":  p.account,
		"username": p.username,
		"region":   p.region,
		"role":     p.role,
	})
	return nil
}
