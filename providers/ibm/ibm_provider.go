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

package ibm

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type IBMProvider struct {
	terraformutils.Provider
	ResourceGroup string
}

func (p *IBMProvider) Init(args []string) error {
	if args[0] != "" {
		p.ResourceGroup = args[0]
	}
	return nil
}

func (p *IBMProvider) GetName() string {
	return "ibm"
}

func (p *IBMProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"ibm": map[string]interface{}{},
		},
	}
}

func (IBMProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *IBMProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"ibm_kp":                     &KPGenerator{},
		"ibm_container_vpc_cluster":  &VPCClusterGenerator{},
		"ibm_container_cluster":      &ContainerClusterGenerator{},
		"ibm_cos":                    &COSGenerator{},
		"ibm_database_elasticsearch": &DatabaseElasticSearchGenerator{},
		"ibm_database_etcd":          &DatabaseETCDGenerator{},
		"ibm_database_mongo":         &DatabaseMongoGenerator{},
		"ibm_database_postgresql":    &DatabasePostgresqlGenerator{},
		"ibm_database_rabbitmq":      &DatabaseRabbitMQGenerator{},
		"ibm_database_redis":         &DatabaseRedisGenerator{},
		"ibm_iam":                    &IAMGenerator{},
		"ibm_is_instance_group":      &InstanceGroupGenerator{},
		"ibm_is_vpc":                 &VPCGenerator{},
		"ibm_is_subnet":              &SubnetGenerator{},
		"ibm_is_instance":            &InstanceGenerator{},
		"ibm_is_security_group":      &SecurityGroupGenerator{},
		"ibm_cis":                    &CISGenerator{},
	}
}

func (p *IBMProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("IBM: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())

	p.Service.SetArgs(map[string]interface{}{
		"resource_group": p.ResourceGroup,
	})
	return nil
}
