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
	"log"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

type VPCGenerator struct {
	IBMService
}

func (g VPCGenerator) createVPCResources(vpcID, vpcName string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		vpcID,
		vpcName,
		"ibm_is_vpc",
		"ibm",
		[]string{})
	return resources
}

func (g VPCGenerator) createVPCAddressPrefixResources(vpcID, addPrefixID, addPrefixName string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s/%s", vpcID, addPrefixID),
		addPrefixName,
		"ibm_is_vpc_address_prefix",
		"ibm",
		[]string{})
	return resources
}

func (g VPCGenerator) createVPCRouteResources(vpcID, routeID, routeName string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		fmt.Sprintf("%s/%s", vpcID, routeID),
		routeName,
		"ibm_is_vpc_route",
		"ibm",
		[]string{})
	return resources
}

func (g *VPCGenerator) InitResources() error {
	var resoureGroup string
	region := envFallBack([]string{"IC_REGION"}, "us-south")
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		log.Fatal("No API key set")
	}

	rg := g.Args["resource_group"]
	if rg != nil {
		resoureGroup = rg.(string)
	}

	vpcurl := fmt.Sprintf("https://%s.iaas.cloud.ibm.com/v1", region)
	vpcoptions := &vpcv1.VpcV1Options{
		URL: envFallBack([]string{"IBMCLOUD_IS_API_ENDPOINT"}, vpcurl),
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	}
	vpcclient, err := vpcv1.NewVpcV1(vpcoptions)
	if err != nil {
		return err
	}
	start := ""
	allrecs := []vpcv1.VPC{}
	for {
		listVpcsOptions := &vpcv1.ListVpcsOptions{}
		if start != "" {
			listVpcsOptions.Start = &start
		}
		if resoureGroup != "" {
			listVpcsOptions.ResourceGroupID = &resoureGroup
		}
		vpcs, response, err := vpcclient.ListVpcs(listVpcsOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching vpcs %s\n%s", err, response)
		}
		start = GetNext(vpcs.Next)
		allrecs = append(allrecs, vpcs.Vpcs...)
		if start == "" {
			break
		}
	}

	for _, vpc := range allrecs {
		g.Resources = append(g.Resources, g.createVPCResources(*vpc.ID, *vpc.Name))
		listVPCAddressPrefixesOptions := &vpcv1.ListVPCAddressPrefixesOptions{
			VPCID: vpc.ID,
		}
		addprefixes, response, err := vpcclient.ListVPCAddressPrefixes(listVPCAddressPrefixesOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching vpc address prefixes %s\n%s", err, response)
		}
		for _, addprefix := range addprefixes.AddressPrefixes {
			g.Resources = append(g.Resources, g.createVPCAddressPrefixResources(*vpc.ID, *addprefix.ID, *addprefix.Name))
		}
		listVPCRoutesOptions := &vpcv1.ListVPCRoutesOptions{
			VPCID: vpc.ID,
		}
		routes, response, err := vpcclient.ListVPCRoutes(listVPCRoutesOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching vpc routes %s\n%s", err, response)
		}
		for _, route := range routes.Routes {
			g.Resources = append(g.Resources, g.createVPCRouteResources(*vpc.ID, *route.ID, *route.Name))
		}

	}
	return nil
}
