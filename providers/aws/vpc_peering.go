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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var ngwAllowEmptyValues = []string{"tags."}

type VpcPeeringConnectionGenerator struct {
	AWSService
}

func (g VpcPeeringConnectionGenerator) createVpcPeeringConnectionsResources(svc *ec2.EC2) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	err := svc.DescribeVpcPeeringConnectionsPages(
		&ec2.DescribeVpcPeeringConnectionsInput{},
		func(ngws *ec2.DescribeVpcPeeringConnectionsOutput, lastPage bool) bool {
			for _, ngw := range ngws.VpcPeeringConnections {
				resources = append(resources, terraform_utils.NewSimpleResource(
					aws.StringValue(ngw.VpcPeeringConnectionId),
					aws.StringValue(ngw.VpcPeeringConnectionId),
					"aws_vpc_peering_connection",
					"aws",
					ngwAllowEmptyValues,
				))
			}
			return true
		},
	)

	if err != nil {
		log.Println(err)
		return resources
	}

	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each VPC Peering Connection
func (g *VpcPeeringConnectionGenerator) InitResources() error {
	sess := g.generateSession()
	svc := ec2.New(sess)

	g.Resources = g.createVpcPeeringConnectionsResources(svc)
	return nil
}
