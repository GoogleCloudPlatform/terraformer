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

package auth0

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type Auth0Provider struct { //nolint
	terraformutils.Provider
	domain       string
	clientID     string
	clientSecret string
}

func (p *Auth0Provider) Init(args []string) error {
	orgName := os.Getenv("AUTH0_DOMAIN")
	if orgName == "" {
		return errors.New("set AUTH0_DOMAIN env var")
	}
	p.domain = orgName

	baseURL := os.Getenv("AUTH0_CLIENT_ID")
	if baseURL == "" {
		return errors.New("set AUTH0_CLIENT_ID env var")
	}
	p.clientID = baseURL

	apiToken := os.Getenv("AUTH0_CLIENT_SECRET")
	if apiToken == "" {
		return errors.New("set AUTH0_CLIENT_SECRET env var")
	}
	p.clientSecret = apiToken

	return nil
}

func (p *Auth0Provider) GetName() string {
	return "auth0"
}

func (p *Auth0Provider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"domain":        cty.StringVal(p.domain),
		"client_id":     cty.StringVal(p.clientID),
		"client_secret": cty.StringVal(p.clientSecret),
	})
}

func (p *Auth0Provider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"domain":        p.domain,
		"client_id":     p.clientID,
		"client_secret": p.clientSecret,
	})
	return nil
}

func (p *Auth0Provider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"auth0_action":          &ActionGenerator{},
		"auth0_client":          &ClientGenerator{},
		"auth0_client_grant":    &ClientGrantGenerator{},
		"auth0_hook":            &HookGenerator{},
		"auth0_resource_server": &ResourceServerGenerator{},
		"auth0_role":            &RoleGenerator{},
		"auth0_rule":            &RuleGenerator{},
		"auth0_rule_config":     &RuleConfigGenerator{},
		"auth0_trigger":         &TriggerBindingGenerator{},
		"auth0_user":            &UserGenerator{},
	}
}

func (p Auth0Provider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p Auth0Provider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}
