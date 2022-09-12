// Copyright 2019 The Terraformer Authors.
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

package ibm

import (
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// VPEGenerator ...
type VPEGenerator struct {
	IBMService
}

func (g VPEGenerator) createVPEGatewayResources(gatewayID, gatewayName string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		gatewayID,
		normalizeResourceName(gatewayName, false),
		"ibm_is_virtual_endpoint_gateway",
		"ibm",
		[]string{})
	return resources
}

func (g VPEGenerator) createVPEGatewayIPResources(gatewayID, gatewayIPID, gatewayIPName string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", gatewayID, gatewayIPID),
		normalizeResourceName(gatewayIPName, false),
		"ibm_is_virtual_endpoint_gateway_ip",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})
	return resources
}

// InitResources ...
func (g *VPEGenerator) InitResources() error {
	region := g.Args["region"].(string)
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("No API key set")
	}

	isURL := GetVPCEndPoint(region)
	iamURL := GetAuthEndPoint()
	vpcoptions := &vpcv1.VpcV1Options{
		URL: isURL,
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
			URL:    iamURL,
		},
	}
	vpcclient, err := vpcv1.NewVpcV1(vpcoptions)
	if err != nil {
		return err
	}

	start := ""
	allrecs := []vpcv1.EndpointGateway{}
	for {
		listEndpointGatewaysOptions := &vpcv1.ListEndpointGatewaysOptions{}
		if start != "" {
			listEndpointGatewaysOptions.Start = &start
		}
		if rg := g.Args["resource_group"].(string); rg != "" {
			rg, err = GetResourceGroupID(apiKey, rg, region)
			if err != nil {
				return fmt.Errorf("Error Fetching Resource Group Id %s", err)
			}
			listEndpointGatewaysOptions.ResourceGroupID = &rg
		}
		gateways, response, err := vpcclient.ListEndpointGateways(listEndpointGatewaysOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching endpoint gateways %s\n%s", err, response)
		}
		start = GetNext(gateways.Next)
		allrecs = append(allrecs, gateways.EndpointGateways...)
		if start == "" {
			break
		}
	}

	for _, gateway := range allrecs {
		start := ""
		allrecs := []vpcv1.ReservedIP{}
		g.Resources = append(g.Resources, g.createVPEGatewayResources(*gateway.ID, *gateway.Name))
		listEndpointGatewayIpsOptions := &vpcv1.ListEndpointGatewayIpsOptions{
			EndpointGatewayID: gateway.ID,
		}
		if start != "" {
			listEndpointGatewayIpsOptions.Start = &start
		}
		ips, response, err := vpcclient.ListEndpointGatewayIps(listEndpointGatewayIpsOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching endpoint gateway ips %s\n%s", err, response)
		}
		start = GetNext(ips.Next)
		allrecs = append(allrecs, ips.Ips...)
		if start == "" {
			break
		}
		for _, ip := range allrecs {
			g.Resources = append(g.Resources, g.createVPEGatewayIPResources(*gateway.ID, *ip.ID, *ip.Name))
		}
	}
	return nil
}

func (g *VPEGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "ibm_is_virtual_endpoint_gateway" {
			continue
		}
		for _, gIP := range g.Resources {
			if gIP.InstanceInfo.Type != "ibm_is_virtual_endpoint_gateway_ip" {
				continue
			}
			if gIP.InstanceState.Attributes["gateway"] == r.InstanceState.Attributes["id"] {
				g.Resources[i].Item["gateway"] = "${ibm_is_virtual_endpoint_gateway." + r.ResourceName + ".id}"
			}
		}
	}

	return nil
}
