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

package logzio

import (
	"regexp"
	"strings"

	"github.com/zclconf/go-cty/cty"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/pkg/errors"
)

type LogzioProvider struct {
	terraform_utils.Provider
	token   string
	baseURL string
}

var (
	disallowedChars = regexp.MustCompile(`[^A-Za-z0-9-]`)
)

const logzioProviderVersion = "~>v1.1.1"

func (p LogzioProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"alerts": {"alert_notification_endpoints": []string{"alert_notification_endpoints", "id"}},
	}
}

func (p LogzioProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"logzio": map[string]interface{}{
				"version": logzioProviderVersion,
			},
		},
	}
}

func (p *LogzioProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"token":   cty.StringVal(p.token),
		"baseURL": cty.StringVal(p.baseURL),
	})
}

// Init LogzioProvider with API token
func (p *LogzioProvider) Init(args []string) error {
	p.token = args[0]
	p.baseURL = args[1]
	return nil
}

func (p *LogzioProvider) GetName() string {
	return "logzio"
}

func (p *LogzioProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"token":   p.token,
		"baseURL": p.baseURL,
	})
	return nil
}

// GetSupportedService return map of support service for Logzio
func (p *LogzioProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"alerts":                       &AlertsGenerator{},
		"alert_notification_endpoints": &AlertNotificationEndpointsGenerator{},
	}
}

func createSlug(s string) string {
	s = strings.ToLower(s)

	return disallowedChars.ReplaceAllString(s, "-")
}
