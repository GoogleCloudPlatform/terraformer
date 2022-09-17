// Copyright 2022 The Terraformer Authors.
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

package sumologic

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type SumoLogicProvider struct {
	terraformutils.Provider
	AccessId    string
	AccessKey   string
	Environment string
	baseUrl     string
}

func (p *SumoLogicProvider) Init(args []string) error {
	if args[0] != "" {
		p.AccessId = args[0]
	} else if accessId := os.Getenv("SUMOLOGIC_ACCESS_ID"); accessId != "" {
		p.AccessId = os.Getenv("SUMOLOGIC_ACCESS_ID")
	} else {
		return errors.New("accessId is not set")
	}

	if args[1] != "" {
		p.AccessKey = args[1]
	} else if accessKey := os.Getenv("SUMOLOGIC_ACCESS_KEY"); accessKey != "" {
		p.AccessKey = os.Getenv("SUMOLOGIC_ACCESS_KEY")
	} else {
		return errors.New("accessKey is not set")
	}

	if args[2] != "" {
		p.Environment = args[2]
	} else if environment := os.Getenv("SUMOLOGIC_ENVIRONMENT"); environment != "" {
		p.Environment = environment
	} else if baseUrl := os.Getenv("SUMOLOGIC_BASE_URL"); baseUrl != "" {
		p.baseUrl = baseUrl
	} else {
		return errors.New("environment is not set")
	}

	return nil
}

func (p *SumoLogicProvider) GetName() string {
	return "sumologic"
}

func (p *SumoLogicProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"access_id":   cty.StringVal(p.AccessId),
		"access_key":  cty.StringVal(p.AccessKey),
		"environment": cty.StringVal(p.Environment),
		"base_url":    cty.StringVal(p.baseUrl),
	})
}

func (p *SumoLogicProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"role": &RoleGenerator{},
		"user": &UserGenerator{},
	}
}

func (p *SumoLogicProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"accessId":    p.AccessId,
		"accessKey":   p.AccessKey,
		"environment": p.Environment,
		"baseUrl":     p.baseUrl,
	})

	return nil
}

func (p *SumoLogicProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"sumologic": map[string]interface{}{
				"accessId":    p.AccessId,
				"accessKey":   p.AccessKey,
				"environment": p.Environment,
			},
		},
	}
}

func (SumoLogicProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}
