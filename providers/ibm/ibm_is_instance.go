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

// InstanceGenerator ...
type InstanceGenerator struct {
	IBMService
}

func (g InstanceGenerator) createInstanceResources(instanceID, instanceName, instanceImgID string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		instanceID,
		normalizeResourceName(instanceName, false),
		"ibm_is_instance",
		"ibm",
		map[string]string{
			"image": instanceImgID,
		},
		[]string{},
		map[string]interface{}{
			"keys": []string{},
		})

	// Deprecated parameters
	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^port$",
	)
	return resource
}

// InitResources ...
func (g *InstanceGenerator) InitResources() error {
	region := g.Args["region"].(string)
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("No API key set")
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
	var allrecs []vpcv1.Instance
	for {
		options := &vpcv1.ListInstancesOptions{}
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
		instances, response, err := vpcclient.ListInstances(options)
		if err != nil {
			return fmt.Errorf("Error Fetching Instances %s\n%s", err, response)
		}
		start = GetNext(instances.Next)
		allrecs = append(allrecs, instances.Instances...)
		if start == "" {
			break
		}
	}

	for _, instance := range allrecs {
		g.Resources = append(g.Resources, g.createInstanceResources(*instance.ID, *instance.Name, *instance.Image.ID))
	}
	return nil
}
