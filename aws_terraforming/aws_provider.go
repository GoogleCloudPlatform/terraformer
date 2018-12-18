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

package aws_terraforming

import (
	"fmt"
	"log"
	"os"
	"waze/terraformer/terraform_utils"

	"github.com/pkg/errors"
)

const GenerateFilesFolderPath = "%s/generated/aws/%s/%s/"

type AWSProvider struct {
	terraform_utils.Provider
	region string
}

func (p AWSProvider) RegionResource() map[string]interface{} {
	return map[string]interface{}{
		"aws": map[string]interface{}{
			"region": p.region,
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

func (p *AWSProvider) GenerateOutputPath() error {
	if err := os.MkdirAll(p.CurrentPath(), os.ModePerm); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (p *AWSProvider) CurrentPath() string {
	rootPath, _ := os.Getwd()
	return fmt.Sprintf(GenerateFilesFolderPath, rootPath, p.region, p.Service.GetName())
}

func (p *AWSProvider) InitService(serviceName string) error {
	var isSupported bool
	if _, isSupported = p.GetAWSSupportService()[serviceName]; !isSupported {
		return errors.New("aws: " + serviceName + " not supported service")
	}
	p.Service = p.GetAWSSupportService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]string{
		"region": p.region,
	})
	return nil
}

// GetAWSSupportService return map of support service for AWS
func (p *AWSProvider) GetAWSSupportService() map[string]terraform_utils.ServiceGenerator {
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
	}
}
