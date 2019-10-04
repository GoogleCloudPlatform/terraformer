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

package alicloud

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/zclconf/go-cty/cty"

	"github.com/pkg/errors"
)

type AlicloudProvider struct {
	terraform_utils.Provider
	region  string
	profile string
}

const alicloudProviderVersion = ">1.56.0"

// GetConfig return map of provider config for Datadog
func (p *AlicloudProvider) GetConfig() cty.Value {
	config, _ := LoadConfigFromProfile()

	return cty.ObjectVal(map[string]cty.Value{
		// "name":               cty.StringVal("default"),
		"AccessKey": cty.StringVal(config.AccessKey),
		"SecretKey": cty.StringVal(config.SecretKey),
		// "SecurityToken":      cty.StringVal(config.SecurityToken),
		"RamRoleArn":         cty.StringVal(config.RamRoleArn),
		"RamRoleSessionName": cty.StringVal(config.RamRoleSessionName),
		"RegionId":           cty.StringVal(config.RegionId),
		// "ram_role_name":     cty.StringVal(config.RamRoleSessionName),
		// "mode":              cty.StringVal("RamRoleArn"),
		// "private_key":       cty.StringVal(""),
		// "key_pair_name":     cty.StringVal(""),
		// "expired_seconds":   cty.NumberIntVal(900),
		// "verified":          cty.StringVal(""),
		// "output_format":     cty.StringVal("json"),
		// "language":          cty.StringVal("en"),
		// "site":              cty.StringVal(""),
		// "retry_timeout":     cty.NumberIntVal(0),
		// "retry_count":       cty.NumberIntVal(0),
	})
}

func (p AlicloudProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		// "subnet":         {"vpc": []string{"vpc_id", "id"}},
		// "vpn_gateway":    {"vpc": []string{"vpc_id", "id"}},
		// "vpn_connection": {"vpn_gateway": []string{"vpn_gateway_id", "id"}},
		// "rds": {
		// 	"subnet": []string{"subnet_ids", "id"},
		// 	"sg":     []string{"vpc_security_group_ids", "id"},
		// },
		// "nacl": {
		// 	"subnet": []string{"subnet_ids", "id"},
		// 	"vpc":    []string{"vpc_id", "id"},
		// },
		// "igw": {"vpc": []string{"vpc_id", "id"}},
		// "elasticache": {
		// 	"vpc":    []string{"vpc_id", "id"},
		// 	"subnet": []string{"subnet_ids", "id"},
		// 	"sg":     []string{"security_group_ids", "id"},
		// },
		// "alb": {
		// 	"sg":     []string{"security_groups", "id"},
		// 	"subnet": []string{"subnets", "id"},
		// 	"vpc":    []string{"vpc_id", "id"},
		// },
		// "elb": {
		// 	"sg":     []string{"security_groups", "id"},
		// 	"subnet": []string{"subnets", "id"},
		// },
		// "auto_scaling": {
		// 	"sg":     []string{"security_groups", "id"},
		// 	"subnet": []string{"vpc_zone_identifier", "id"},
		// },
		// "ec2_instance": {
		// 	"sg":     []string{"vpc_security_group_ids", "id"},
		// 	"subnet": []string{"subnet_id", "id"},
		// },
		// "route_table": {
		// 	"vpc": []string{"vpc_id", "id"},
		// },
	}
}

func (p AlicloudProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"alicloud": map[string]interface{}{
				"version": alicloudProviderVersion,
				"region":  p.region,
			},
		},
	}
}

// check projectName in env params
func (p *AlicloudProvider) Init(args []string) error {
	p.region = args[0]
	p.profile = args[1]
	// terraform work with env params ALICLOUD_DEFAULT_REGION
	err := os.Setenv("ALICLOUD_DEFAULT_REGION", p.region)
	if err != nil {
		return err
	}
	return nil
}

func (p *AlicloudProvider) GetName() string {
	return "alicloud"
}

func (p *AlicloudProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("alicloud: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"region":  p.region,
		"profile": p.profile,
	})
	s, _ := json.MarshalIndent(p.Service, "", "\t")
	fmt.Println("InitService" + string(s))
	return nil
}

func (p *AlicloudProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		// "vpc":            &VpcGenerator{},
		// "sg":             &SecurityGenerator{},
		// "subnet":         &SubnetGenerator{},
		// "igw":            &IgwGenerator{},
		// "vpn_gateway":    &VpnGatewayGenerator{},
		// "nacl":           &NaclGenerator{},
		// "vpn_connection": &VpnConnectionGenerator{},
		"oss": &OSSGenerator{},
		// "elb": &ElbGenerator{},
		// "iam":            &IamGenerator{},
		// "route53":        &Route53Generator{},
		// "auto_scaling":   &AutoScalingGenerator{},
		// "rds":            &RDSGenerator{},
		// "elasticache":    &ElastiCacheGenerator{},
		// "alb":            &AlbGenerator{},
		// "acm":            &ACMGenerator{},
		// "cloudfront":     &CloudFrontGenerator{},
		"ecs": &EcsGenerator{},
		// "firehose":       &FirehoseGenerator{},
		// "glue":           &GlueGenerator{},
		// "route_table":    &RouteTableGenerator{},
	}
}
