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

package commercetools

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/pkg/errors"
)

type CommercetoolsProvider struct { //nolint
	terraformutils.Provider
	clientID     string
	clientSecret string
	clientScope  string
	projectKey   string
	baseURL      string
	tokenURL     string
}

func (p CommercetoolsProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p CommercetoolsProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

// Init CommerectoolsProvider
func (p *CommercetoolsProvider) Init(args []string) error {
	p.clientID = args[0]
	p.clientScope = args[1]
	p.clientSecret = args[2]
	p.projectKey = args[3]
	p.baseURL = args[4]
	p.tokenURL = args[5]
	return nil
}

func (p *CommercetoolsProvider) GetName() string {
	return "commercetools"
}

func (p *CommercetoolsProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"client_id":     p.clientID,
		"client_secret": p.clientSecret,
		"client_scope":  p.clientScope,
		"project_key":   p.projectKey,
		"base_url":      p.baseURL,
		"token_url":     p.tokenURL,
	})
	return nil
}

// GetSupportedService return map of support service for Logzio
func (p *CommercetoolsProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"api_extension":   &APIExtensionGenerator{},
		"channel":         &ChannelGenerator{},
		"custom_object":   &CustomObjectGenerator{},
		"product_type":    &ProductTypeGenerator{},
		"shipping_zone":   &ShippingZoneGenerator{},
		"shipping_method": &ShippingMethodGenerator{},
		"state":           &StateGenerator{},
		"store":           &StoreGenerator{},
		"subscription":    &SubscriptionGenerator{},
		"tax_category":    &TaxCategoryGenerator{},
		"types":           &TypesGenerator{},
	}
}
