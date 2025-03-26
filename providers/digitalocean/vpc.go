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

package digitalocean

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/digitalocean/godo"
)

type VPCGenerator struct {
	DigitalOceanService
}

func (g VPCGenerator) listVPCs(ctx context.Context, client *godo.Client) ([]*godo.VPC, error) {
	list := []*godo.VPC{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		vpcs, resp, err := client.VPCs.List(ctx, opt)
		if err != nil {
			return nil, err
		}

		list = append(list, vpcs...)

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}

	return list, nil
}

func (g VPCGenerator) createResources(vpcList []*godo.VPC) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, vpc := range vpcList {
		resources = append(resources, terraformutils.NewSimpleResource(
			vpc.ID,
			vpc.Name,
			"digitalocean_vpc",
			"digitalocean",
			[]string{}))
	}
	return resources
}

func (g *VPCGenerator) InitResources() error {
	client := g.generateClient()
	output, err := g.listVPCs(context.TODO(), client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
