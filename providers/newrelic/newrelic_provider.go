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

package newrelic

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type NewRelicProvider struct { //nolint
	terraformutils.Provider
}

func (p *NewRelicProvider) Init(args []string) error {
	return nil
}

func (p *NewRelicProvider) GetName() string {
	return "newrelic"
}

func (p *NewRelicProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (NewRelicProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *NewRelicProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"alert":      &AlertGenerator{},
		"dashboard":  &DashboardGenerator{},
		"infra":      &InfraGenerator{},
		"synthetics": &SyntheticsGenerator{},
	}
}

func (p *NewRelicProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("newrelic: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())

	return nil
}
