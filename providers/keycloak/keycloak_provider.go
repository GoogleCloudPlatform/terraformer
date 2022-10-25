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
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type KeycloakProvider struct { //nolint
	terraformutils.Provider
	url                   string
	basePath              string
	clientID              string
	clientSecret          string
	realm                 string
	clientTimeout         int
	caCert                string
	tlsInsecureSkipVerify bool
	redHatSSO             bool
	target                string
}

func getArg(arg string) string {
	if arg == "-" {
		return ""
	}
	return arg
}

func (p *KeycloakProvider) Init(args []string) error {
	p.url = args[0]
	p.basePath = args[1]
	p.clientID = args[2]
	p.clientSecret = args[3]
	p.realm = args[4]
	p.clientTimeout, _ = strconv.Atoi(args[5])
	p.caCert = getArg(args[6])
	p.tlsInsecureSkipVerify, _ = strconv.ParseBool(args[7])
	p.redHatSSO, _ = strconv.ParseBool(args[8])
	p.target = getArg(args[9])
	return nil
}

func (p *KeycloakProvider) GetName() string {
	return "keycloak"
}

func (p *KeycloakProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (p *KeycloakProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"url":                      cty.StringVal(p.url),
		"base_path":                cty.StringVal(p.basePath),
		"client_id":                cty.StringVal(p.clientID),
		"client_secret":            cty.StringVal(p.clientSecret),
		"realm":                    cty.StringVal(p.realm),
		"client_timeout":           cty.NumberIntVal(int64(p.clientTimeout)),
		"root_ca_certificate":      cty.StringVal(p.caCert),
		"tls_insecure_skip_verify": cty.BoolVal(p.tlsInsecureSkipVerify),
		"red_hat_sso":              cty.BoolVal(p.redHatSSO),
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
		"url":                      p.url,
		"base_path":                p.basePath,
		"client_id":                p.clientID,
		"client_secret":            p.clientSecret,
		"realm":                    p.realm,
		"client_timeout":           p.clientTimeout,
		"root_ca_certificate":      p.caCert,
		"tls_insecure_skip_verify": p.tlsInsecureSkipVerify,
		"red_hat_sso":              p.redHatSSO,
		"target":                   p.target,
	})
	return nil
}

func (p *KeycloakProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"realms": &RealmGenerator{},
	}
}

func (KeycloakProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}
