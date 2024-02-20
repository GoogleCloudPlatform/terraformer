// Copyright 2021 The Terraformer Authors.
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

package mackerel

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/mackerelio/mackerel-client-go"
	"github.com/zclconf/go-cty/cty"
)

type MackerelProvider struct {
	terraformutils.Provider
	apiKey      string
	apiBase     string
	serviceName string
	client      *mackerel.Client
}

func (m *MackerelProvider) Init(args []string) error {
	if args[2] == "" {
		return errors.New("you should specify mackerel service name")
	}

	if apiKey := os.Getenv("MACKEREL_API_KEY"); apiKey != "" {
		m.apiKey = apiKey
	}
	if base := os.Getenv("MACKEREL_API_BASE"); base != "" {
		m.apiBase = base
	}

	if args[0] != "" {
		m.apiKey = args[0]
	}

	if args[1] != "" {
		m.apiBase = args[1]
	}

	m.serviceName = args[2]
	return nil
}

func (m *MackerelProvider) GetName() string {
	return "mackerel"
}

func (m *MackerelProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"api_key": cty.StringVal(m.apiKey),
	})
}

func (m *MackerelProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"mackerel": map[string]interface{}{},
		},
	}
}

func (MackerelProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (m *MackerelProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"alert_group_setting": &AlertGroupSettingGenerator{
			serviceName: m.serviceName,
		},
		"downtime": &DowntimeGenerator{
			serviceName: m.serviceName,
		},
		"monitor": &MonitorGenerator{
			serviceName: m.serviceName,
		},
		"channel": &ChannelGenerator{
			serviceName: m.serviceName,
		},
		"notification_group": &NotificationGroupGenerator{
			serviceName: m.serviceName,
		},
		"role": &RoleGenerator{
			serviceName: m.serviceName,
		},
		"service": &ServiceGenerator{
			serviceName: m.serviceName,
		},
		"role_metadata": &RoleMetadataGenerator{
			serviceName: m.serviceName,
		},
		"service_metadata": &ServiceMetadataGenerator{
			serviceName: m.serviceName,
		},
	}
}

func (m *MackerelProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = m.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(m.GetName() + ": " + serviceName + " not supported service")
	}
	m.Service = m.GetSupportedService()[serviceName]
	m.Service.SetName(serviceName)
	m.Service.SetVerbose(verbose)
	m.Service.SetProviderName(m.GetName())
	m.Service.SetArgs(map[string]interface{}{
		"api_key":  m.apiKey,
		"api_base": m.apiBase,
	})
	return nil
}
