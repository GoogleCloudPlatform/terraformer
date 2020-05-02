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
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var peeringAllowEmptyValues = []string{"tags."}

type VpcPeeringConnectionGenerator struct {
	AWSService
}

func (g *VpcPeeringConnectionGenerator) createResources(peerings *ec2.DescribeVpcPeeringConnectionsOutput) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, peering := range peerings.VpcPeeringConnections {
		resources = append(resources, terraformutils.NewSimpleResource(
			aws.StringValue(peering.VpcPeeringConnectionId),
			aws.StringValue(peering.VpcPeeringConnectionId),
			"aws_vpc_peering_connection",
			"aws",
			peeringAllowEmptyValues,
		))
	}

	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each VPC Peering Connection
func (g *VpcPeeringConnectionGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.New(config)
	p := ec2.NewDescribeVpcPeeringConnectionsPaginator(svc.DescribeVpcPeeringConnectionsRequest(&ec2.DescribeVpcPeeringConnectionsInput{}))
	for p.Next(context.Background()) {
		g.Resources = append(g.Resources, g.createResources(p.CurrentPage())...)
	}
	return p.Err()
}
