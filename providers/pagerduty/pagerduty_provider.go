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

package pagerduty

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type PagerDutyProvider struct { //nolint
	terraformutils.Provider
	token string
}

func (p *PagerDutyProvider) Init(args []string) error {
	if token := os.Getenv("PAGERDUTY_TOKEN"); token != "" {
		p.token = os.Getenv("PAGERDUTY_TOKEN")
	}
	if len(args) > 0 && args[0] != "" {
		p.token = args[0]
	}
	return nil
}

func (p *PagerDutyProvider) GetName() string {
	return "pagerduty"
}

func (p *PagerDutyProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"token": cty.StringVal(p.token),
	})
}

func (p *PagerDutyProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"pagerduty": map[string]interface{}{
				"token": p.token,
			},
		},
	}
}

func (PagerDutyProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *PagerDutyProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"business_service":  &BusinessServiceGenerator{},
		"escalation_policy": &EscalationPolicyGenerator{},
		"ruleset":           &RulesetGenerator{},
		"schedule":          &ScheduleGenerator{},
		"service":           &ServiceGenerator{},
		"team":              &TeamGenerator{},
		"user":              &UserGenerator{},
	}
}

func (p *PagerDutyProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"token": p.token,
	})
	return nil
}
