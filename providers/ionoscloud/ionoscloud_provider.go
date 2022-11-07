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

package ionoscloud

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type IonosCloudProvider struct { //nolint
	terraformutils.Provider
	username string
	password string
	token    string
	url      string
}

func (p *IonosCloudProvider) Init(args []string) error {
	username := os.Getenv(helpers.IonosUsername)
	password := os.Getenv(helpers.IonosPassword)
	token := os.Getenv(helpers.IonosToken)
	url := os.Getenv(helpers.IonosApiUrl)

	if (username == "" || password == "") && token == "" {
		return errors.New(helpers.CredentialsError)
	}

	if username != "" && password != "" {
		p.username = username
		p.password = password
	}

	if token != "" {
		p.token = token
	}

	p.url = url

	return nil
}

func (p *IonosCloudProvider) GetName() string {
	return helpers.ProviderName
}

func (p *IonosCloudProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (IonosCloudProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"compute": {
			"datacenter": []string{
				"lan", "id",
				"server", "id"},
		},
	}
}

func (p *IonosCloudProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"datacenter":               &DatacenterGenerator{},
		"lan":                      &LanGenerator{},
		"nic":                      &NicGenerator{},
		"server":                   &ServerGenerator{},
		"volume":                   &VolumeGenerator{},
		"pg_cluster":               &DBaaSClusterGenerator{},
		"backup_unit":              &BackupUnitGenerator{},
		"ipblock":                  &IPBlockGenerator{},
		"k8s_cluster":              &KubernetesClusterGenerator{},
		"k8s_node_pool":            &KubernetesNodePoolGenerator{},
		"target_group":             &TargetGroupGenerator{},
		"networkloadbalancer":      &NetworkLoadBalancerGenerator{},
		"natgateway":               &NATGatewayGenerator{},
		"group":                    &GroupGenerator{},
		"application_loadbalancer": &ApplicationLoadBalancer{},
	}
}

func (p *IonosCloudProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(helpers.Ionos + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"username": p.username,
		"password": p.password,
		"token":    p.token,
		"url":      p.url,
	})
	return nil
}
