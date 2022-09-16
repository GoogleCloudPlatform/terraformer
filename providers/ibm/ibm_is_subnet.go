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

// SubnetGenerator ...
type SubnetGenerator struct {
	IBMService
}

func (g SubnetGenerator) createSubnetResources(subnetID, subnetName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		subnetID,
		normalizeResourceName(subnetName, true),
		"ibm_is_subnet",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})

	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^total_ipv4_address_count$",
	)
	return resource
}

// InitResources ...
func (g *SubnetGenerator) InitResources() error {
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
	var allrecs []vpcv1.VPC
	for {
		listVpcsOptions := &vpcv1.ListVpcsOptions{}
		if start != "" {
			listVpcsOptions.Start = &start
		}
		if rg := g.Args["resource_group"].(string); rg != "" {
			rg, err = GetResourceGroupID(apiKey, rg, region)
			if err != nil {
				return fmt.Errorf("Error Fetching Resource Group Id %s", err)
			}
			listVpcsOptions.ResourceGroupID = &rg
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
		start = ""
		var allSubNetRecs []vpcv1.Subnet
		for {
			options := &vpcv1.ListSubnetsOptions{}
			if start != "" {
				options.Start = &start
			}

			subnets, response, err := vpcclient.ListSubnets(options)
			if err != nil {
				return fmt.Errorf("Error Fetching subnets %s\n%s", err, response)
			}
			start = GetNext(subnets.Next)
			allSubNetRecs = append(allSubNetRecs, subnets.Subnets...)
			if start == "" {
				break
			}
		}

		for _, subnet := range allSubNetRecs {
			if *vpc.ID == *subnet.VPC.ID {
				g.Resources = append(g.Resources, g.createSubnetResources(*subnet.ID, *subnet.Name))
			}
		}
	}

	return nil
}
