// Copyright 2025 The Terraformer Authors.
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

package keboola

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type KeboolaProvider struct {
	terraformutils.Provider
	apiKey string
	apiURL string
}

func (p *KeboolaProvider) Init(args []string) error {
	if os.Getenv("KEBOOLA_TOKEN") == "" {
		return errors.New("set KEBOOLA_TOKEN env var")
	}
	p.apiKey = os.Getenv("KEBOOLA_TOKEN")

	if os.Getenv("KEBOOLA_HOST") != "" {
		p.apiURL = os.Getenv("KEBOOLA_HOST")
	} else {
		p.apiURL = "https://connection.keboola.com" // Default URL, can be changed
	}

	return nil
}

func (p *KeboolaProvider) GetName() string {
	return "keboola"
}

func (p *KeboolaProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"keboola": map[string]interface{}{
				"host":  p.apiURL,
				"token": p.apiKey,
			},
		},
	}
}

func (p KeboolaProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"component_configuration": {
			"scheduler": []string{"id"},
		},
		"scheduler": {
			"component_configuration": []string{"configuration_id"},
		},
	}
}

func (p *KeboolaProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"component_configuration": &ComponentConfigurationGenerator{},
		"scheduler":               &SchedulerGenerator{},
	}
}

func (p *KeboolaProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("keboola: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"host":  p.apiURL,
		"token": p.apiKey,
	})
	return nil
}
