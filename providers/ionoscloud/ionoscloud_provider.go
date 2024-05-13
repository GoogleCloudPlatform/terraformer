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
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type IonosCloudProvider struct { //nolint
	terraformutils.Provider
	username string
	password string
	token    string
	url      string
}

func (p *IonosCloudProvider) Init(_ []string) error {
	username := os.Getenv(ionoscloud.IonosUsernameEnvVar)
	password := os.Getenv(ionoscloud.IonosPasswordEnvVar)
	token := os.Getenv(ionoscloud.IonosTokenEnvVar)
	url := os.Getenv(ionoscloud.IonosApiUrlEnvVar)

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

func (p *IonosCloudProvider) GetProviderData(_ ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (IonosCloudProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"server": {
			"datacenter": []string{helpers.DcID, "id"},
		},
		"nic": {
			"datacenter": []string{helpers.DcID, "id"},
			"server":     []string{helpers.ServerID, "id"},
		},
		"volume": {
			"datacenter": []string{helpers.DcID, "id"},
			"server":     []string{helpers.ServerID, "id"},
		},
		"firewall": {
			"datacenter": []string{helpers.DcID, "id"},
			"server":     []string{helpers.ServerID, "id"},
			"nic":        []string{helpers.NicID, "id"},
		},
		"k8s_node_pool": {
			"datacenter":  []string{helpers.DcID, "id"},
			"k8s_cluster": []string{helpers.K8sClusterID, "id"},
		},
		"networkloadbalancer": {
			"datacenter": []string{helpers.DcID, "id"},
		},
		"natgateway": {
			"datacenter": []string{helpers.DcID, "id"},
		},
		"application_loadbalancer": {
			"datacenter": []string{helpers.DcID, "id"},
		},
		"networkloadbalancer_forwardingrule": {
			"datacenter":   []string{helpers.DcID, "id"},
			"loadbalancer": []string{"networkloadbalancer_id", "id"},
		},
		"loadbalancer": {
			"datacenter": []string{helpers.DcID, "id"},
		},
		"natgateway_rule": {
			"datacenter": []string{helpers.DcID, "id"},
			"natgateway": []string{"natgateway_id", "id"},
		},
		"s3_key": {
			"user": []string{helpers.UserID, "id"},
		},
		"share": {
			"group": []string{helpers.GroupID, "id"},
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
		"pg_cluster":               &DBaaSPgSQLClusterGenerator{},
		"pg_user":                  &DBaaSPgSQLUserGenerator{},
		"pg_database":              &DBaaSPgSQLDatabaseGenerator{},
		"mongo_cluster":            &DBaaSMongoClusterGenerator{},
		"mongo_user":               &DBaaSMongoUserGenerator{},
		"backup_unit":              &BackupUnitGenerator{},
		"ipblock":                  &IPBlockGenerator{},
		"k8s_cluster":              &KubernetesClusterGenerator{},
		"k8s_node_pool":            &KubernetesNodePoolGenerator{},
		"target_group":             &TargetGroupGenerator{},
		"networkloadbalancer":      &NetworkLoadBalancerGenerator{},
		"natgateway":               &NATGatewayGenerator{},
		"group":                    &GroupGenerator{},
		"application_loadbalancer": &ApplicationLoadBalancerGenerator{},
		"application_loadbalancer_forwardingrule": &ALBForwardingRuleGenerator{},
		"firewall":                           &FirewallGenerator{},
		"networkloadbalancer_forwardingrule": &NetworkLoadBalancerForwardingRuleGenerator{},
		"loadbalancer":                       &LoadBalancerGenerator{},
		"natgateway_rule":                    &NATGatewayRuleGenerator{},
		"certificate":                        &CertificateGenerator{},
		"private_crossconnect":               &PrivateCrossConnectGenerator{},
		"s3_key":                             &S3KeyGenerator{},
		"container_registry":                 &ContainerRegistryGenerator{},
		"dataplatform_cluster":               &DataPlatformClusterGenerator{},
		"dataplatform_node_pool":             &DataPlatformNodePoolGenerator{},
		"share":                              &ShareGenerator{},
		"user":                               &UserGenerator{},
		"container_registry_token":           &ContainerRegistryTokenGenerator{},
		"dns_zone":                           &DNSZoneGenerator{},
		"dns_record":                         &DNSRecordGenerator{},
		"logging_pipeline":                   &LoggingPipelineGenerator{},
		"ipfailover":                         &IPFailoverGenerator{},
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
