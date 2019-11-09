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

var rtbAllowEmptyValues = []string{"tags."}

type RouteTableGenerator struct {
	AWSService
}

func (g RouteTableGenerator) createRouteTablesResources(svc *ec2.EC2) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	err := svc.DescribeRouteTablesPages(
		&ec2.DescribeRouteTablesInput{},
		func(tables *ec2.DescribeRouteTablesOutput, lastPage bool) bool {
			for _, table := range tables.RouteTables {
				// route table
				resources = append(resources, terraform_utils.NewSimpleResource(
					aws.StringValue(table.RouteTableId),
					aws.StringValue(table.RouteTableId),
					"aws_route_table",
					"aws",
					rtbAllowEmptyValues,
				))

				for _, assoc := range table.Associations {
					if aws.BoolValue(assoc.Main) {
						// main route table association
						resources = append(resources, terraform_utils.NewResource(
							aws.StringValue(assoc.RouteTableAssociationId),
							aws.StringValue(table.VpcId),
							"aws_main_route_table_association",
							"aws",
							map[string]string{
								"vpc_id":         aws.StringValue(table.VpcId),
								"route_table_id": aws.StringValue(table.RouteTableId),
							},
							rtbAllowEmptyValues,
							map[string]interface{}{},
						))
					} else {
						// subnet-specific route table association
						resources = append(resources, terraform_utils.NewResource(
							aws.StringValue(assoc.RouteTableAssociationId),
							aws.StringValue(assoc.SubnetId),
							"aws_route_table_association",
							"aws",
							map[string]string{
								"subnet_id":      aws.StringValue(assoc.SubnetId),
								"route_table_id": aws.StringValue(table.RouteTableId),
							},
							rtbAllowEmptyValues,
							map[string]interface{}{},
						))
					}
				}
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
// create terraform resource for each route tables
func (g *RouteTableGenerator) InitResources() error {
	sess := g.generateSession()
	svc := ec2.New(sess)

	g.Resources = g.createRouteTablesResources(svc)
	return nil
}
