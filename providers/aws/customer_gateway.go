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

package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var customerGatewayAllowEmptyValues = []string{"tags."}

type CustomerGatewayGenerator struct {
	AWSService
}

func (CustomerGatewayGenerator) createResources(cgws *ec2.DescribeCustomerGatewaysOutput) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, cgws := range cgws.CustomerGateways {
		resources = append(resources, terraformutils.NewSimpleResource(
			StringValue(cgws.CustomerGatewayId),
			StringValue(cgws.CustomerGatewayId),
			"aws_customer_gateway",
			"aws",
			customerGatewayAllowEmptyValues,
		))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each customer gateway create 1 TerraformResource.
// Need CustomerGatewayId as ID for terraform resource
func (g *CustomerGatewayGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.NewFromConfig(config)
	cgws, err := svc.DescribeCustomerGateways(context.TODO(), &ec2.DescribeCustomerGatewaysInput{})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(cgws)
	return nil
}
