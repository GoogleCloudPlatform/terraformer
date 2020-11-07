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

package ns1

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type Ns1Provider struct { //nolint
	terraformutils.Provider
	apiKey string
}

func (p *Ns1Provider) Init(args []string) error {
	if os.Getenv("NS1_APIKEY") == "" {
		return errors.New("set NS1_APIKEY env var")
	}
	p.apiKey = os.Getenv("NS1_APIKEY")

	return nil
}

func (p *Ns1Provider) GetName() string {
	return "ns1"
}

func (p *Ns1Provider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (Ns1Provider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *Ns1Provider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"monitoringjob": &MonitoringJobGenerator{},
		"team":          &TeamGenerator{},
		"zone":          &ZoneGenerator{},
	}
}

func (p *Ns1Provider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("ns1: " + serviceName + " not supported service")
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
