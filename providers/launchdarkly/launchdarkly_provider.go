// Copyright 2020 The Terraformer Authors.
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

package launchdarkly

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type LaunchDarklyProvider struct { //nolint
	terraformutils.Provider
	apiKey string
}

const (
	basePath   = "https://app.launchdarkly.com/api/v2"
	version    = "0.0.1"
	APIVersion = "20191212"
)

func (p *LaunchDarklyProvider) Init(args []string) error {
	if os.Getenv("LAUNCHDARKLY_ACCESS_TOKEN") == "" {
		return errors.New("set LAUNCHDARKLY_ACCESS_TOKEN env var")
	}
	p.apiKey = os.Getenv("LAUNCHDARKLY_ACCESS_TOKEN")

	return nil
}

func (p *LaunchDarklyProvider) GetName() string {
	return "launchdarkly"
}

func (p *LaunchDarklyProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"launchdarkly": map[string]interface{}{
				"api_key": p.apiKey,
			},
		},
	}
}

func (LaunchDarklyProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *LaunchDarklyProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"project": &ProjectGenerator{},
	}
}

func (p *LaunchDarklyProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("launchdarkly: " + serviceName + " not supported service")
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
