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
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type SnowflakeConfig struct {
	Account  string `envconfig:"SNOWFLAKE_ACCOUNT" required:"true"`
	Username string `envconfig:"SNOWFLAKE_USER" required:"true"`
	Password string `envconfig:"SNOWFLAKE_PASSWORD" required:"true"`
	Region   string `envconfig:"SNOWFLAKE_REGION" default:""`
	Role     string `envconfig:"SNOWFLAKE_ROLE" required:"true"`
}

type SnowflakeProvider struct {
	terraformutils.Provider

	Config *SnowflakeConfig
}

func (p *SnowflakeProvider) Init(args []string) error {
	p.Config = &SnowflakeConfig{}
	err := envconfig.Process("", p.Config)
	if err != nil {
		return errors.Wrap(err, "")
	}

	// us-west-2 is their default region, but if you actually specify that it won't trigger their default code
	//  https://github.com/snowflakedb/gosnowflake/blob/52137ce8c32eaf93b0bd22fc5c7297beff339812/dsn.go#L61
	if p.Config.Region == "us-west-2" {
		p.Config.Region = ""
	}
	return nil
}

func (p *SnowflakeProvider) GetName() string {
	return "snowflake"
}

func (p *SnowflakeProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"snowflake": map[string]interface{}{
				"account":  p.Config.Account,
				"username": p.Config.Username,
				"region":   p.Config.Region,
				"role":     p.Config.Role,
				"password": p.Config.Password,
			},
		},
	}
}

func (SnowflakeProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *SnowflakeProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"database":        &DatabaseGenerator{},
		"database_grant":  &DatabaseGrantGenerator{},
		"role":            &RoleGenerator{},
		"role_grant":      &RoleGrantGenerator{},
		"schema":          &SchemaGenerator{},
		"schema_grant":    &SchemaGrantGenerator{},
		"user":            &UserGenerator{},
		"view":            &ViewGenerator{},
		"warehouse":       &WarehouseGenerator{},
		"warehouse_grant": &WarehouseGrantGenerator{},
	}
}

func (p *SnowflakeProvider) InitService(serviceName string) error {
	if _, isSupported := p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(fmt.Sprintf("snowflake: %s not a supported service", serviceName))
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"account":  p.Config.Account,
		"username": p.Config.Username,
		"region":   p.Config.Region,
		"role":     p.Config.Role,
		"password": p.Config.Password,
	})
	return nil
}
