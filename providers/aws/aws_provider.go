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

package aws

import (
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/pkg/errors"
)

type AWSProvider struct {
	terraform_utils.Provider
	region string
}

const awsProviderVersion = ">1.56.0"

func (p AWSProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{
		"subnet":         {"vpc": []string{"vpc_id", "id"}},
		"vpn_gateway":    {"vpc": []string{"vpc_id", "id"}},
		"vpn_connection": {"vpn_gateway": []string{"vpn_gateway_id", "id"}},
		"rds": {
			"subnet": []string{"subnet_ids", "id"},
			"sg":     []string{"vpc_security_group_ids", "id"},
		},
		"nacl": {
			"subnet": []string{"subnet_ids", "id"},
			"vpc":    []string{"vpc_id", "id"},
		},
		"igw": {"vpc": []string{"vpc_id", "id"}},
		"elasticache": {
			"vpc":    []string{"vpc_id", "id"},
			"subnet": []string{"subnet_ids", "id"},
			"sg":     []string{"security_group_ids", "id"},
		},
		"alb": {
			"sg":     []string{"security_groups", "id"},
			"subnet": []string{"subnets", "id"},
			"vpc":    []string{"vpc_id", "id"},
		},
		"elb": {
			"sg":     []string{"security_groups", "id"},
			"subnet": []string{"subnets", "id"},
		},
		"auto_scaling": {
			"sg":     []string{"security_groups", "id"},
			"subnet": []string{"vpc_zone_identifier", "id"},
		},
	}
}
func (p AWSProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"aws": map[string]interface{}{
				"version": awsProviderVersion,
				"region":  p.region,
			},
		},
	}
}

// check projectName in env params
func (p *AWSProvider) Init(args []string) error {
	p.region = args[0]
	// terraform work with env param AWS_DEFAULT_REGION
	err := os.Setenv("AWS_DEFAULT_REGION", p.region)
	if err != nil {
		return err
	}
	return nil
}

func (p *AWSProvider) GetName() string {
	return "aws"
}

func (p *AWSProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("aws: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"region": p.region,
	})
	return nil
}

// GetAWSSupportService return map of support service for AWS
func (p *AWSProvider) GetSupportedService() map[string]terraform_utils.ServiceGenerator {
	return map[string]terraform_utils.ServiceGenerator{
		"vpc":            &VpcGenerator{},
		"sg":             &SecurityGenerator{},
		"subnet":         &SubnetGenerator{},
		"igw":            &IgwGenerator{},
		"vpn_gateway":    &VpnGatewayGenerator{},
		"nacl":           &NaclGenerator{},
		"vpn_connection": &VpnConnectionGenerator{},
		"s3":             &S3Generator{},
		"elb":            &ElbGenerator{},
		"iam":            &IamGenerator{},
		"route53":        &Route53Generator{},
		"auto_scaling":   &AutoScalingGenerator{},
		"rds":            &RDSGenerator{},
		"elasticache":    &ElastiCacheGenerator{},
		"alb":            &AlbGenerator{},
		"acm":            &ACMGenerator{},
		"cloudfront":     &CloudFrontGenerator{},
		"ec2_instance":   &Ec2Generator{},
	}
}
