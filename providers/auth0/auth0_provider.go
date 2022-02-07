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
	if args[0] != "" {
		p.domain = args[0]
	} else {
		if domain := os.Getenv("AUTH0_DOMAIN"); domain != "" {
			p.domain = domain
		} else {
			return errors.New("domain requirement")
		}
	}

	if args[1] != "" {
		p.clientID = args[1]
	} else {
		if clientID := os.Getenv("AUTH0_CLIENT_ID"); clientID != "" {
			p.clientID = clientID
		} else {
			return errors.New("clientID requirement")
		}
	}

	if args[2] != "" {
		p.clientSecret = args[2]
	} else {
		if clientSecret := os.Getenv("AUTH0_CLIENT_SECRET"); clientSecret != "" {
			p.clientSecret = clientSecret
		} else {
			return errors.New("clientSecret requirement")
		}
	}

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
		"action":          &ActionGenerator{},
		"client":          &ClientGenerator{},
		"client_grant":    &ClientGrantGenerator{},
		"hook":            &HookGenerator{},
		"resource_server": &ResourceServerGenerator{},
		"role":            &RoleGenerator{},
		"rule":            &RuleGenerator{},
		"rule_config":     &RuleConfigGenerator{},
		"trigger":         &TriggerBindingGenerator{},
		"user":            &UserGenerator{},
	}
}

func (p Auth0Provider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p Auth0Provider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}
