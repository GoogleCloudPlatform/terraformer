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

package fastly

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils/provider_wrapper"
)

type FastlyProvider struct {
	terraform_utils.Provider
	apiKey string
}

func (p *FastlyProvider) Init(args []string) error {
	if os.Getenv("FASTLY_API_KEY") == "" {
		return errors.New("set FASTLY_API_KEY env var")
	}
	p.apiKey = os.Getenv("FASTLY_API_KEY")

	return nil
}

func (p *FastlyProvider) GetName() string {
	return "fastly"
}

func (p *FastlyProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"fastly": map[string]interface{}{
				"version": provider_wrapper.GetProviderVersion(p.GetName()),
				"api_key": p.apiKey,
			},
		},
	}
}

func (FastlyProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *FastlyProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"service_v1": &ServiceV1Generator{},
	}
}

func (p *FastlyProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("fastly: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"api_key": p.apiKey,
	})
	return nil
}
