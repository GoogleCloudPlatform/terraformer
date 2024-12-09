// Copyright 2024 The Terraformer Authors.
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

package scaleway

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/scaleway/scaleway-sdk-go/api/account/v3"
	"github.com/scaleway/scaleway-sdk-go/api/vpc/v2"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type VpcGenerator struct {
	ScalewayService
}

func ProjectName(client *scw.Client, projectID string) (string, error) {
	accountAPI := account.NewProjectAPI(client)

	resp, err := accountAPI.GetProject(&account.ProjectAPIGetProjectRequest{
		ProjectID: projectID,
	})
	if err != nil {
		return "", err
	}
	return resp.Name, nil
}

func (g VpcGenerator) ListVPCs(client *scw.Client) ([]*vpc.VPC, error) {
	list := []*vpc.VPC{}
	vpcApi := vpc.NewAPI(client)

	page := int32(1)
	pagesize := uint32(100)
	opt := &vpc.ListVPCsRequest{
		Page:     &page,
		PageSize: &pagesize,
	}

	for {
		resp, err := vpcApi.ListVPCs(opt)
		if err != nil {
			return nil, err
		}
		for _, vpc := range resp.Vpcs {
			if vpc != nil {
				list = append(list, vpc)
			}
		}
		// Exit loop when we are on the last page.
		if resp.TotalCount < *opt.PageSize {
			break
		}
		*opt.Page++
	}

	return list, nil
}

func (g VpcGenerator) createResources(vpcList []*vpc.VPC, client *scw.Client) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, vpc := range vpcList {
		project_name, _ := ProjectName(client, vpc.ProjectID)
		resources = append(resources, terraformutils.NewSimpleResource(
			string(vpc.Region)+"/"+vpc.ID,
			string(project_name)+"-"+vpc.Name,
			"scaleway_vpc",
			"scaleway",
			[]string{}))
	}
	return resources
}

func (g *VpcGenerator) InitResources() error {
	client := g.generateClient()
	output, err := g.ListVPCs(client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output, client)
	return nil
}
