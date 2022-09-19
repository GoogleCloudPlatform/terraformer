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

// PlacementGroupGenerator ...
type PlacementGroupGenerator struct {
	IBMService
}

func (g PlacementGroupGenerator) createPlacementGroupResources(pGroupID, pGroupName string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		pGroupID,
		normalizeResourceName(pGroupName, false),
		"ibm_is_placement_group",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})
	return resources
}

// InitResources ...
func (g *PlacementGroupGenerator) InitResources() error {
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

	listPlacementGroupsOptions := &vpcv1.ListPlacementGroupsOptions{}
	start := ""
	allPlacementGroupsRecs := []vpcv1.PlacementGroup{}
	for {
		if start != "" {
			listPlacementGroupsOptions.Start = &start
		}
		placementGroupCollection, response, err := vpcclient.ListPlacementGroups(listPlacementGroupsOptions)
		if err != nil {
			return fmt.Errorf("list PlacementGroups with Context failed %s\n%s", err, response)
		}
		start = GetNext(placementGroupCollection.Next)
		allPlacementGroupsRecs = append(allPlacementGroupsRecs, placementGroupCollection.PlacementGroups...)
		if start == "" {
			break
		}
	}

	for _, pg := range allPlacementGroupsRecs {
		g.Resources = append(g.Resources, g.createPlacementGroupResources(*pg.ID, *pg.Name))
	}

	return nil
}
