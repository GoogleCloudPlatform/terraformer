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

package gmailfilter

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type GmailfilterProvider struct { //nolint
	terraformutils.Provider
	credentials           string
	impersonatedUserEmail string
}

func (p *GmailfilterProvider) Init(args []string) error {
	credentials := os.Getenv("GOOGLE_CREDENTIALS")
	if len(args) > 0 && args[0] != "" {
		credentials = args[0]
		os.Setenv("GOOGLE_CREDENTIALS", credentials)
	}
	email := os.Getenv("IMPERSONATED_USER_EMAIL")
	if len(args) > 1 && args[1] != "" {
		email = args[1]
		os.Setenv("IMPERSONATED_USER_EMAIL", email)
	}

	p.credentials = credentials
	p.impersonatedUserEmail = email

	return nil
}

func (p *GmailfilterProvider) GetName() string {
	return "gmailfilter"
}

func (p *GmailfilterProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("gmailfilter: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"credentials":           p.credentials,
		"impersonatedUserEmail": p.impersonatedUserEmail,
	})
	return nil
}

// GetGCPSupportService return map of support service for GCP
func (p *GmailfilterProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	services := make(map[string]terraformutils.ServiceGenerator)
	services["label"] = &LabelGenerator{}
	services["filter"] = &FilterGenerator{}
	return services
}

func (p *GmailfilterProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"filter": {
			"label": {
				"action.add_label_ids", "id",
			},
		},
	}
}

func (p *GmailfilterProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}
