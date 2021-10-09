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
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const DefaultRegion = "us-south"
const NoRegion = ""

type IBMProvider struct { //nolint
	terraformutils.Provider
	ResourceGroup string
	Region        string
	CIS           string
}

func (p *IBMProvider) Init(args []string) error {
	p.ResourceGroup = args[0]
	p.Region = args[1]
	p.CIS = args[2]

	var err error
	if p.Region != DefaultRegion && p.Region != NoRegion {
		err = os.Setenv("IC_REGION", p.Region)
	} else {
		p.Region = DefaultRegion
		err = os.Setenv("IC_REGION", DefaultRegion)
	}
	if err != nil {
		return err
	}
	return nil
}

func (p *IBMProvider) GetName() string {
	return "ibm"
}

func (p *IBMProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"ibm": map[string]interface{}{
				"region": p.Region,
			},
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
		"ibm_is_network_acl":         &NetworkACLGenerator{},
		"ibm_is_public_gateway":      &PublicGatewayGenerator{},
		"ibm_is_volume":              &VolumeGenerator{},
		"ibm_is_vpn_gateway":         &VPNGatewayGenerator{},
		"ibm_is_lb":                  &LBGenerator{},
		"ibm_is_ssh_key":             &SSHKeyGenerator{},
		"ibm_is_floating_ip":         &FloatingIPGenerator{},
		"ibm_is_image":               &ImageGenerator{},
		"ibm_is_ipsec_policy":        &IpsecGenerator{},
		"ibm_is_ike_policy":          &IkeGenerator{},
		"ibm_is_flow_log":            &FlowLogGenerator{},
		"ibm_is_instance_template":   &InstanceTemplateGenerator{},
		"ibm_function":               &CloudFunctionGenerator{},
		"ibm_private_dns":            &privateDNSTemplateGenerator{},
		"ibm_certificate_manager":    &CMGenerator{},
		"ibm_direct_link":            &DLGenerator{},
		"ibm_transit_gateway":        &TGGenerator{},
		"ibm_vpe_gateway":            &VPEGenerator{},
		"ibm_satellite":              &SatelliteGenerator{},
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
		"region":         p.Region,
		"cis":            p.CIS,
	})
	return nil
}
