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

package grafana

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/pkg/errors"
	"github.com/zclconf/go-cty/cty"
)

type GrafanaProvider struct { //nolint
	terraformutils.Provider
	auth               string
	url                string
	organization       int
	tlsKey             string
	tlsCert            string
	caCert             string
	insecureSkipVerify bool
}

func (p GrafanaProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"grafana_dashboard": {
			"folder": []string{"grafana_folder", "id"},
		},
	}
}

func (p GrafanaProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"grafana": map[string]interface{}{
				"org_id":               p.organization,
				"url":                  p.url,
				"auth":                 p.auth,
				"tls_key":              p.tlsKey,
				"tls_cert":             p.tlsCert,
				"ca_cert":              p.caCert,
				"insecure_skip_verify": p.insecureSkipVerify,
			},
		},
	}
}

func (p *GrafanaProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"org_id":               cty.NumberIntVal(int64(p.organization)),
		"url":                  cty.StringVal(p.url),
		"auth":                 cty.StringVal(p.auth),
		"tls_key":              cty.StringVal(p.tlsKey),
		"tls_cert":             cty.StringVal(p.tlsCert),
		"ca_cert":              cty.StringVal(p.caCert),
		"insecure_skip_verify": cty.BoolVal(p.insecureSkipVerify),
	})
}

// Init GrafanaProvider with API apiToken
func (p *GrafanaProvider) Init(args []string) error {
	p.auth = args[0]
	p.url = args[1]
	p.organization = 1
	p.insecureSkipVerify = true
	return nil
}

func (p *GrafanaProvider) GetName() string {
	return "grafana"
}

func (p *GrafanaProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"org_id":               p.organization,
		"url":                  p.url,
		"auth":                 p.auth,
		"tls_key":              p.tlsKey,
		"tls_cert":             p.tlsCert,
		"ca_cert":              p.caCert,
		"insecure_skip_verify": p.insecureSkipVerify,
	})
	return nil
}

func (p *GrafanaProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"grafana_dashboard": &DashboardGenerator{},
		"grafana_folder":    &FolderGenerator{},
	}
}
