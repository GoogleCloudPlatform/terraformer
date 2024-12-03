// Copyright 2024 The Terraformer Authors.
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

package scaleway

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ScalewayProvider struct { //nolint
	terraformutils.Provider
	accesskey    string
	secretkey    string
	region       string
	organization string
}

func (p *ScalewayProvider) Init(args []string) error {
	if os.Getenv("SCW_ACCESS_KEY") == "" {
		return errors.New("set SCW_ACCESS_KEY env var")
	}
	p.accesskey = os.Getenv("SCW_ACCESS_KEY")

	if os.Getenv("SCW_SECRET_KEY") == "" {
		return errors.New("set SCW_SECRET_KEY env var")
	}
	p.secretkey = os.Getenv("SCW_SECRET_KEY")

	if os.Getenv("SCW_REGION") == "" {
		return errors.New("set SCW_REGION env var")
	}
	p.region = os.Getenv("SCW_REGION")

	if os.Getenv("SCW_DEFAULT_ORGANIZATION_ID") == "" {
		return errors.New("set SCW_DEFAULT_ORGANIZATION_ID env var")
	}
	p.organization = os.Getenv("SCW_DEFAULT_ORGANIZATION_ID")

	return nil
}

func (p *ScalewayProvider) GetName() string {
	return "scaleway"
}

func (p *ScalewayProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (ScalewayProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *ScalewayProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"instances":  &InstanceGenerator{},
		"kubernetes": &KubernetesGenerator{},
		"vpc":        &VpcGenerator{},
	}
}

func (p *ScalewayProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("scaleway: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"accesskey":    p.accesskey,
		"secretkey":    p.secretkey,
		"region":       p.region,
		"organization": p.organization,
	})
	return nil
}
