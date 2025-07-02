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

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

var dxAllowEmptyValues = []string{"tags."}

type DirectConnectGenerator struct {
	AWSService
}

func (g *DirectConnectGenerator) getDirectConnectGateways(svc *directconnect.Client) error {
	input := &directconnect.DescribeDirectConnectGatewaysInput{}
	for {
		// Fetch a page of results
		output, err := svc.DescribeDirectConnectGateways(context.TODO(), input)
		if err != nil {
			return err
		}

		// Process each DirectConnect Gateway
		for _, dx := range output.DirectConnectGateways {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				*dx.DirectConnectGatewayId, // Dereference the pointer
				*dx.DirectConnectGatewayId,
				"aws_dx_gateway",
				"aws",
				dxAllowEmptyValues,
			))
		}

		// Check if there are more pages
		if output.NextToken == nil {
			break
		}

		// Update the input token for the next page
		input.NextToken = output.NextToken
	}
	return nil
}

func (g *DirectConnectGenerator) getDirectConnectConnections(svc *directconnect.Client) error {
	input := &directconnect.DescribeConnectionsInput{}
	output, err := svc.DescribeConnections(context.TODO(), input)
	if err != nil {
		return err
	}

	for _, dx := range output.Connections {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*dx.ConnectionId, // Dereference the pointer
			*dx.ConnectionName,
			"aws_dx_connection",
			"aws",
			dxAllowEmptyValues,
		))
	}
	return nil
}

func (g *DirectConnectGenerator) getDirectConnectVritualInterfaces(svc *directconnect.Client) error {
	input := &directconnect.DescribeVirtualInterfacesInput{}
	output, err := svc.DescribeVirtualInterfaces(context.TODO(), input)
	if err != nil {
		return err
	}

	for _, vif := range output.VirtualInterfaces {
		var resourceType string

		if *vif.VirtualInterfaceType == "private" {
			resourceType = "aws_dx_private_virtual_interface"
		} else if *vif.VirtualInterfaceType == "public" {
			resourceType = "aws_dx_public_virtual_interface"
		} else {
			log.Printf("Unknown Virtual Interface Type: %s for ID: %s", *vif.VirtualInterfaceType, *vif.VirtualInterfaceId)
			continue
		}

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			*vif.VirtualInterfaceId,
			*vif.VirtualInterfaceName,
			resourceType,
			"aws",
			dxAllowEmptyValues,
		))
	}

	return nil

}

func (g *DirectConnectGenerator) InitResources() error {
	config, err := g.generateConfig()
	if err != nil {
		return err
	}
	svc := directconnect.NewFromConfig(config)
	if err := g.getDirectConnectGateways(svc); err != nil {
		log.Println(err)
		return err
	}

	err = g.getDirectConnectVritualInterfaces(svc)
	if err != nil {
		log.Println(err)
	}

	err = g.getDirectConnectConnections(svc)
	if err != nil {
		log.Println(err)
	}

	return nil
}
