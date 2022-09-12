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

// VPCGenerator ...
type VPCAddressPrefixGenerator struct {
	IBMService
}

func (g VPCAddressPrefixGenerator) createVPCAddressPrefixResources(vpcID, addPrefixID, addPrefixName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", vpcID, addPrefixID),
		normalizeResourceName(addPrefixName, false),
		"ibm_is_vpc_address_prefix",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})

	return resource
}

// InitResources ...
func (g *VPCAddressPrefixGenerator) InitResources() error {
	region := g.Args["region"].(string)
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("no API key set")
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

		// address prefix
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

	}
	return nil
}
