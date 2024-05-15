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
	VPC           string
}

func (p *IBMProvider) Init(args []string) error {
	p.ResourceGroup = args[0]
	p.Region = args[1]
	p.VPC = args[2]

	var err error
	if p.Region != NoRegion {
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
	return map[string]map[string][]string{
		"ibm_is_vpc_route": {
			"ibm_is_vpc": []string{"vpc", "id"},
		},
		"ibm_is_vpc_routing_table": {
			"ibm_is_vpc": []string{"vpc", "id"},
		},
		"ibm_is_subnet": {
			"ibm_is_vpc":            []string{"vpc", "id"},
			"ibm_is_public_gateway": []string{"public_gateway", "id"},
		},
		"ibm_is_vpc_address_prefix": {
			"ibm_is_vpc": []string{"vpc", "id"},
		},
		"ibm_is_lb": {
			"ibm_is_vpc":            []string{"vpc", "id"},
			"ibm_is_subnet":         []string{"subnets", "id"},
			"ibm_is_security_group": []string{"security_groups", "id"},
		},
		"ibm_is_instance": {
			"ibm_is_vpc":            []string{"vpc", "id"},
			"ibm_is_image":          []string{"image", "id"},
			"ibm_is_subnet":         []string{"primary_network_interface.subnet", "id"},
			"ibm_is_security_group": []string{"primary_network_interface.security_groups", "id"},
			"ibm_is_volume":         []string{"volumes", "id"},
		},
		"ibm_is_security_group": {
			"ibm_is_vpc": []string{"vpc", "id"},
		},
		"ibm_is_network_acl": {
			"ibm_is_vpc":    []string{"vpc", "id"},
			"ibm_is_subnet": []string{"subnets", "id"},
		},
		"ibm_is_public_gateway": {
			"ibm_is_vpc":         []string{"vpc", "id"},
			"ibm_is_floating_ip": []string{"floating_ip.id", "id"},
		},
		"ibm_container_vpc_cluster": {
			"ibm_is_vpc":    []string{"vpc_id", "id"},
			"ibm_is_subnet": []string{"zones.subnet_id", "id"},
		},
		"ibm_vpe_gateway": {
			"ibm_is_vpc":            []string{"vpc", "id"},
			"ibm_is_security_group": []string{"security_groups", "id"},
			"ibm_is_subnet":         []string{"ips.subnet", "id"},
		},
		"ibm_is_vpn_gateway": {
			"ibm_is_subnet": []string{"subnet", "id"},
		},
	}
}

func (p *IBMProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"ibm_kp":                            &KPGenerator{},
		"ibm_container_vpc_cluster":         &VPCClusterGenerator{},
		"ibm_container_cluster":             &ContainerClusterGenerator{},
		"ibm_cos":                           &COSGenerator{},
		"ibm_database_elasticsearch":        &DatabaseElasticSearchGenerator{},
		"ibm_database_etcd":                 &DatabaseETCDGenerator{},
		"ibm_database_mongo":                &DatabaseMongoGenerator{},
		"ibm_database_postgresql":           &DatabasePostgresqlGenerator{},
		"ibm_database_rabbitmq":             &DatabaseRabbitMQGenerator{},
		"ibm_database_redis":                &DatabaseRedisGenerator{},
		"ibm_iam":                           &IAMGenerator{},
		"ibm_is_instance_group":             &InstanceGroupGenerator{},
		"ibm_is_vpc":                        &VPCGenerator{},
		"ibm_is_vpc_address_prefix":         &VPCAddressPrefixGenerator{},
		"ibm_is_vpc_route":                  &VPCRouteGenerator{},
		"ibm_is_vpc_routing_table":          &VPCRoutingTableGenerator{},
		"ibm_is_subnet":                     &SubnetGenerator{},
		"ibm_is_instance":                   &InstanceGenerator{},
		"ibm_is_security_group":             &SecurityGroupGenerator{},
		"ibm_cis":                           &CISGenerator{},
		"ibm_is_network_acl":                &NetworkACLGenerator{},
		"ibm_is_public_gateway":             &PublicGatewayGenerator{},
		"ibm_is_volume":                     &VolumeGenerator{},
		"ibm_is_vpn_gateway":                &VPNGatewayGenerator{},
		"ibm_is_lb":                         &LBGenerator{},
		"ibm_is_ssh_key":                    &SSHKeyGenerator{},
		"ibm_is_floating_ip":                &FloatingIPGenerator{},
		"ibm_is_image":                      &ImageGenerator{},
		"ibm_is_ipsec_policy":               &IpsecGenerator{},
		"ibm_is_ike_policy":                 &IkeGenerator{},
		"ibm_is_flow_log":                   &FlowLogGenerator{},
		"ibm_is_instance_template":          &InstanceTemplateGenerator{},
		"ibm_function":                      &CloudFunctionGenerator{},
		"ibm_private_dns":                   &privateDNSTemplateGenerator{},
		"ibm_certificate_manager":           &CMGenerator{},
		"ibm_direct_link":                   &DLGenerator{},
		"ibm_transit_gateway":               &TGGenerator{},
		"ibm_vpe_gateway":                   &VPEGenerator{},
		"ibm_satellite_control_plane":       &SatelliteControlPlaneGenerator{},
		"ibm_satellite_data_plane":          &SatelliteDataPlaneGenerator{},
		"ibm_secrets_manager":               &SecretsManagerGenerator{},
		"ibm_continuous_delivery":           &ContinuousDeliveryGenerator{},
		"ibm_cloud_sysdig_monitor":          &MonitoringGenerator{},
		"ibm_cloud_logdna":                  &LogAnalysisGenerator{},
		"ibm_cloud_atracker":                &ActivityTrackerGenerator{},
		"ibm_cloud_watson_studio":           &WatsonStudioGenerator{},
		"ibm_cloud_watson_machine_learning": &WatsonMachineLearningGenerator{},
		"ibm_cloudant":                      &CloudantGenerator{},
		"ibm_code_engine":                   &CodeEngineGenerator{},
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
		"vpc":            p.VPC,
	})
	return nil
}
