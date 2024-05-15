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

package bizflycloud

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type BizflyCloudProvider struct { //nolint
	terraformutils.Provider
	authMethod string
	region string
	email string
	password string
	appCredID string
	appCredSecret string
	projectID string
}

func (p *BizflyCloudProvider) Init(args []string) error {
	if os.Getenv("BIZFLYCLOUD_AUTH_METHOD") == "" {
		p.authMethod = "password"
	}

	if os.Getenv("BIZFLYCLOUD_PROJECT_ID") == "" {
		return errors.New("set BIZFLYCLOUD_PROJECT_ID env var")
	}
	p.projectID =  os.Getenv("BIZFLYCLOUD_PROJECT_ID")
	
	if os.Getenv("BIZFLYCLOUD_REGION") == "" {
		return errors.New("set BIZFLYCLOUD_REGION env var")
	}
	p.region = os.Getenv("BIZFLYCLOUD_REGION")

	p.appCredID = ""
	p.appCredSecret = ""
	p.email = ""
	p.password = ""
	if os.Getenv("BIZFLYCLOUD_EMAIL") == "" && p.authMethod == "password" {
		return errors.New("set BIZFLYCLOUD_EMAIL env var")
	}
	p.email = os.Getenv("BIZFLYCLOUD_EMAIL")

	if os.Getenv("BIZFLYCLOUD_PASSWORD") == "" && p.authMethod == "password" {
		return errors.New("set BIZFLYCLOUD_PASSWORD env var")
	}
	p.password = os.Getenv("BIZFLYCLOUD_PASSWORD")



	if os.Getenv("BIZFLYCLOUD_APP_CRED_ID") == "" && p.authMethod == "application_credential" {
		return errors.New("set BIZFLYCLOUD_APP_CRED_ID env var")
	}
	p.appCredID = os.Getenv("BIZFLYCLOUD_APP_CRED_ID")

	if os.Getenv("BIZFLYCLOUD_APP_CRED_SECRET") == "" && p.authMethod == "application_credential" {
		return errors.New("set BIZFLYCLOUD_APP_CRED_SECRET env var")
	}
	p.appCredSecret = os.Getenv("BIZFLYCLOUD_APP_CRED_SECRET")

	return nil
}

func (p *BizflyCloudProvider) GetName() string {
	return "bizflycloud"
}

func (p *BizflyCloudProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (BizflyCloudProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *BizflyCloudProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"server":            &ServerGenerator{},
	}
}

func (p *BizflyCloudProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("bizflycloud: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"auth_method": p.authMethod,
		"email": p.email,
		"password": p.password,
		"app_credential_id": p.appCredID,
		"app_credential_secret": p.appCredSecret,
		"project_id": p.projectID,
		"region_name": p.region,
	})
	return nil
}
