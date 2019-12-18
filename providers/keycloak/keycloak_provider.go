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

package keycloak

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/provider_wrapper"
	"github.com/zclconf/go-cty/cty"
)

type KeycloakProvider struct {
	terraform_utils.Provider
	url          string
	clientID     string
	clientSecret string
	realm        string
}

func (p *KeycloakProvider) Init(args []string) error {
	p.url = args[0]
	p.clientID = args[1]
	p.clientSecret = args[2]
	p.realm = args[3]
	return nil
}

func (p *KeycloakProvider) GetName() string {
	return "keycloak"
}

func (p *KeycloakProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			p.GetName(): map[string]interface{}{
				"version": provider_wrapper.GetProviderVersion(p.GetName()),
			},
		},
	}
}

func (p *KeycloakProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"url":           cty.StringVal(p.url),
		"client_id":     cty.StringVal(p.clientID),
		"client_secret": cty.StringVal(p.clientSecret),
		"realm":         cty.StringVal(p.realm),
	})
}

func (p *KeycloakProvider) GetBasicConfig() cty.Value {
	return p.GetConfig()
}

func (p *KeycloakProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"url":           p.url,
		"client_id":     p.clientID,
		"client_secret": p.clientSecret,
		"realm":         p.realm,
	})
	return nil
}

func (p *KeycloakProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"openid_clients": &OpenIDClientGenerator{},
		"realms":         &RealmGenerator{},
		"users":          &UserGenerator{},
		"roles":          &RoleGenerator{},
		"scopes":         &ScopeGenerator{},
		"groups":         &GroupGenerator{},
	}
}

func (KeycloakProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"openid_clients": {
			"realms": []string{"realm_id", "self_link"},
		},
		"users": {
			"realms": []string{"realm_id", "self_link"},
		},
		"roles": {
			"openid_clients": []string{"client_id", "self_link"},
			"realms":         []string{"realm_id", "self_link"},
		},
		"scopes": {
			"openid_clients": []string{"client_id", "self_link"},
			"realms":         []string{"realm_id", "self_link"},
		},
		"groups": {
			"realms": []string{"realm_id", "self_link"},
			"roles":  []string{"role_ids", "self_link"},
			"users":  []string{"members", "username"},
		},
	}
}
