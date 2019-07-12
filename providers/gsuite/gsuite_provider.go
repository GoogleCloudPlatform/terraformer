// Copyright 2018 The Terraformer Authors.
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

package gsuite

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

const gsuiteProviderVersion = ">0.1.21"

type GSuiteProvider struct {
	terraform_utils.Provider
	credentials           string
	impersonatedUserEmail string
}

// Init check env params
func (p *GSuiteProvider) Init(args []string) error {
	if args[0] != "" {
		p.credentials = args[0]
	} else {
		if credentials := os.Getenv("GOOGLE_CREDENTIALS"); credentials != "" {
			p.credentials = credentials
		} else {
			return errors.New("GOOGLE_CREDENTIALS requirement")
		}
	}

	if args[1] != "" {
		p.impersonatedUserEmail = args[1]
	} else {
		if impersonatedUserEmail := os.Getenv("IMPERSONATED_USER_EMAIL"); impersonatedUserEmail != "" {
			p.impersonatedUserEmail = impersonatedUserEmail
		} else {
			return errors.New("IMPERSONATED_USER_EMAIL requirement")
		}
	}

	return nil
}

// GetName return string of provider name for GSuite
func (p *GSuiteProvider) GetName() string {
	return "gsuite"
}

// GetConfig return map of provider config for GSuite
func (p *GSuiteProvider) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"credentials":             p.credentials,
		"impersonated_user_email": p.impersonatedUserEmail,
	}
}

// InitService ...
func (p *GSuiteProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]string{
		"credentials":             p.credentials,
		"impersonated_user_email": p.impersonatedUserEmail,
	})
	return nil
}

// GetSupportedService return map of support service for GSuite
func (p *GSuiteProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		//"gsuite_domain":         resourceDomain(),
		//"gsuite_group":          resourceGroup(),
		//"gsuite_group_member":   resourceGroupMember(),
		//"gsuite_group_members":  resourceGroupMembers(),
		//"gsuite_group_settings": resourceGroupSettings(),
		"users": &UsersGenerator{},
		//"gsuite_user_schema":    resourceUserSchema(),
	}
}

// GetResourceConnections return map of resource connections for GSuite
func (GSuiteProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

// GetProviderData return map of provider data for GSuite
func (p GSuiteProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			p.GetName(): map[string]interface{}{
				"version": gsuiteProviderVersion,
			},
		},
	}
}
