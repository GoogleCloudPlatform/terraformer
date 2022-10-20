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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var tgwAllowEmptyValues = []string{"tags."}

type TransitGatewayGenerator struct {
	AWSService
}

func (g *TransitGatewayGenerator) getTransitGateways(svc *ec2.Client) error {
	p := ec2.NewDescribeTransitGatewaysPaginator(svc, &ec2.DescribeTransitGatewaysInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, tgw := range page.TransitGateways {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				StringValue(tgw.TransitGatewayId),
				StringValue(tgw.TransitGatewayId),
				"aws_ec2_transit_gateway",
				"aws",
				tgwAllowEmptyValues,
			))
		}
	}
	return nil
}

func (g *TransitGatewayGenerator) getTransitGatewayRouteTables(svc *ec2.Client) error {
	p := ec2.NewDescribeTransitGatewayRouteTablesPaginator(svc, &ec2.DescribeTransitGatewayRouteTablesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, tgwrt := range page.TransitGatewayRouteTables {
			// Default route table are automatically created on the tgw creation
			if *tgwrt.DefaultAssociationRouteTable {
				continue
			} else {
				g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
					StringValue(tgwrt.TransitGatewayRouteTableId),
					StringValue(tgwrt.TransitGatewayRouteTableId),
					"aws_ec2_transit_gateway_route_table",
					"aws",
					tgwAllowEmptyValues,
				))
			}
		}
	}
	return nil
}

func (g *TransitGatewayGenerator) getTransitGatewayVpcAttachments(svc *ec2.Client) error {
	p := ec2.NewDescribeTransitGatewayVpcAttachmentsPaginator(svc, &ec2.DescribeTransitGatewayVpcAttachmentsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, tgwa := range page.TransitGatewayVpcAttachments {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				StringValue(tgwa.TransitGatewayAttachmentId),
				StringValue(tgwa.TransitGatewayAttachmentId),
				"aws_ec2_transit_gateway_vpc_attachment",
				"aws",
				tgwAllowEmptyValues,
			))
		}
	}
	return nil
}

// Generate TerraformResources from AWS API,
// from each customer gateway create 1 TerraformResource.
// Need CustomerGatewayId as ID for terraform resource
func (g *TransitGatewayGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.NewFromConfig(config)
	g.Resources = []terraformutils.Resource{}
	err := g.getTransitGateways(svc)
	if err != nil {
		log.Println(err)
	}

	err = g.getTransitGatewayRouteTables(svc)
	if err != nil {
		log.Println(err)
	}

	err = g.getTransitGatewayVpcAttachments(svc)
	if err != nil {
		log.Println(err)
	}

	return nil
}
