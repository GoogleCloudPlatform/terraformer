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

type SubnetGenerator struct {
	IBMService
}

func (g SubnetGenerator) createSubnetResources(subnetID, subnetName string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		subnetID,
		subnetName,
		"ibm_is_subnet",
		"ibm",
		[]string{})
	return resources
}

func (g *SubnetGenerator) InitResources() error {
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
	allrecs := []vpcv1.Subnet{}
	for {
		options := &vpcv1.ListSubnetsOptions{}
		if start != "" {
			options.Start = &start
		}
		if resoureGroup != "" {
			options.ResourceGroupID = &resoureGroup
		}
		subnets, response, err := vpcclient.ListSubnets(options)
		if err != nil {
			return fmt.Errorf("Error Fetching subnets %s\n%s", err, response)
		}
		start = GetNext(subnets.Next)
		allrecs = append(allrecs, subnets.Subnets...)
		if start == "" {
			break
		}
	}

	for _, subnet := range allrecs {
		g.Resources = append(g.Resources, g.createSubnetResources(*subnet.ID, *subnet.Name))
	}
	return nil
}
