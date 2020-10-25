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

package datadog

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type DatadogProvider struct { //nolint
	terraformutils.Provider
	apiKey          string
	appKey          string
	apiURL          string
	authV1          context.Context
	datadogClientV1 *datadogV1.APIClient
}

// Init check env params and initialize API Client
func (p *DatadogProvider) Init(args []string) error {
	if args[0] != "" {
		p.apiKey = args[0]
	} else {
		if apiKey := os.Getenv("DATADOG_API_KEY"); apiKey != "" {
			p.apiKey = apiKey
		} else {
			return errors.New("api-key requirement")
		}
	}

	if args[1] != "" {
		p.appKey = args[1]
	} else {
		if appKey := os.Getenv("DATADOG_APP_KEY"); appKey != "" {
			p.appKey = appKey
		} else {
			return errors.New("app-key requirement")
		}
	}

	if args[2] != "" {
		p.apiURL = args[2]
	} else if v := os.Getenv("DATADOG_HOST"); v != "" {
		p.apiURL = v
	}

	// Initialize the Datadog API client
	authV1 := context.WithValue(
		context.Background(),
		datadogV1.ContextAPIKeys,
		map[string]datadogV1.APIKey{
			"apiKeyAuth": {
				Key: p.apiKey,
			},
			"appKeyAuth": {
				Key: p.appKey,
			},
		},
	)
	if p.apiURL != "" {
		parsedAPIURL, parseErr := url.Parse(p.apiURL)
		if parseErr != nil {
			return fmt.Errorf(`invalid API Url : %v`, parseErr)
		}
		if parsedAPIURL.Host == "" || parsedAPIURL.Scheme == "" {
			return fmt.Errorf(`missing protocol or host : %v`, p.apiURL)
		}
		// If api url is passed, set and use the api name and protocol on ServerIndex{1}
		authV1 = context.WithValue(authV1, datadogV1.ContextServerIndex, 1)
		authV1 = context.WithValue(authV1, datadogV1.ContextServerVariables, map[string]string{
			"name":     parsedAPIURL.Host,
			"protocol": parsedAPIURL.Scheme,
		})
	}
	p.authV1 = authV1

	configV1 := datadogV1.NewConfiguration()
	datadogClientV1 := datadogV1.NewAPIClient(configV1)
	p.datadogClientV1 = datadogClientV1

	return nil
}

// GetName return string of provider name for Datadog
func (p *DatadogProvider) GetName() string {
	return "datadog"
}

// GetConfig return map of provider config for Datadog
func (p *DatadogProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"api_key": cty.StringVal(p.apiKey),
		"app_key": cty.StringVal(p.appKey),
		"api_url": cty.StringVal(p.apiURL),
	})
}

// InitService ...
func (p *DatadogProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"api-key":         p.apiKey,
		"app-key":         p.appKey,
		"api-url":         p.apiURL,
		"authV1":          p.authV1,
		"datadogClientV1": p.datadogClientV1,
	})
	return nil
}

// GetSupportedService return map of support service for Datadog
func (p *DatadogProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"dashboard":   &DashboardGenerator{},
		"downtime":    &DowntimeGenerator{},
		"monitor":     &MonitorGenerator{},
		"screenboard": &ScreenboardGenerator{},
		"synthetics":  &SyntheticsGenerator{},
		"timeboard":   &TimeboardGenerator{},
		"user":        &UserGenerator{},
	}
}

// GetResourceConnections return map of resource connections for Datadog
func (DatadogProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

// GetProviderData return map of provider data for Datadog
func (p DatadogProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}
