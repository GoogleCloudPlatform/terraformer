// Copyright 2023 The Terraformer Authors.
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

var VpcEndpointAllowEmptyValues = []string{"tags."}

type VpcEndpointGenerator struct {
	AWSService
}

func (g *VpcEndpointGenerator) createResources(vpceps *ec2.DescribeVpcEndpointsOutput) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, vpcEndpoint := range vpceps.VpcEndpoints {
		resources = append(resources, terraformutils.NewSimpleResource(
			StringValue(vpcEndpoint.VpcEndpointId),
			StringValue(vpcEndpoint.VpcEndpointId),
			"aws_vpc_endpoint",
			"aws",
			VpcAllowEmptyValues,
		))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each vpc endpoint create 1 TerraformResource.
// Need VpcEndpointId as ID for terraform resource
func (g *VpcEndpointGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.NewFromConfig(config)
	vpceps, err := svc.DescribeVpcEndpoints(context.TODO(), &ec2.DescribeVpcEndpointsInput{})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(vpceps)
	return nil
}
