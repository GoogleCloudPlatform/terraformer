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

// FloatingIPGenerator ...
type FloatingIPGenerator struct {
	IBMService
}

func (g FloatingIPGenerator) createFloatingIPResources(fipID, fipName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		fipID,
		normalizeResourceName(fipName, false),
		"ibm_is_floating_ip",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})

	// Conflict parameters
	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^zone$",
	)
	return resource
}

// InitResources ...
func (g *FloatingIPGenerator) InitResources() error {
	region := g.Args["region"].(string)
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		log.Fatal("No API key set")
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
	var allrecs []vpcv1.FloatingIP
	for {
		options := &vpcv1.ListFloatingIpsOptions{}
		if start != "" {
			options.Start = &start
		}
		if rg := g.Args["resource_group"].(string); rg != "" {
			rg, err = GetResourceGroupID(apiKey, rg, region)
			if err != nil {
				return fmt.Errorf("Error Fetching Resource Group Id %s", err)
			}
			options.ResourceGroupID = &rg
		}
		fips, response, err := vpcclient.ListFloatingIps(options)
		if err != nil {
			return fmt.Errorf("Error Fetching Floating IPs %s\n%s", err, response)
		}
		start = GetNext(fips.Next)
		allrecs = append(allrecs, fips.FloatingIps...)
		if start == "" {
			break
		}
	}

	for _, fip := range allrecs {
		g.Resources = append(g.Resources, g.createFloatingIPResources(*fip.ID, *fip.Name))
	}
	return nil
}
