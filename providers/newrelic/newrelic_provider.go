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
	"os"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type NewRelicProvider struct { //nolint
	terraformutils.Provider
	accountID int
	APIKey    string
	Region    string
}

func (p *NewRelicProvider) Init(args []string) error {
	if apiKey := os.Getenv("NEW_RELIC_API_KEY"); apiKey != "" {
		p.APIKey = os.Getenv("NEW_RELIC_API_KEY")
	}
	if accountIDs := os.Getenv("NEW_RELIC_ACCOUNT_ID"); accountIDs != "" {
		accountID, err := strconv.Atoi(accountIDs)
		if err != nil {
			return err

		}
		p.accountID = accountID
	}
	if len(args) > 0 {
		p.APIKey = args[0]
	}
	if len(args) > 1 {
		accountID, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		p.accountID = accountID
	}
	if len(args) > 1 {
		p.Region = args[2]
	}
	if p.Region == "" {
		p.Region = "US"
	}
	return nil
}

func (p *NewRelicProvider) GetName() string {
	return "newrelic"
}

func (p *NewRelicProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"account_id": cty.NumberIntVal(int64(p.accountID)),
		"api_key":    cty.StringVal(p.APIKey),
		"region":     cty.StringVal(p.Region),
	})
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
	p.Service.SetArgs(map[string]interface{}{"apiKey": p.APIKey})
	p.Service.SetProviderName(p.GetName())

	return nil
}
