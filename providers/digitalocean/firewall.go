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

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/digitalocean/godo"
)

type FirewallGenerator struct {
	DigitalOceanService
}

func (g FirewallGenerator) listFirewalls(ctx context.Context, client *godo.Client) ([]godo.Firewall, error) {
	list := []godo.Firewall{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		firewalls, resp, err := client.Firewalls.List(ctx, opt)
		if err != nil {
			return nil, err
		}

		for _, firewall := range firewalls {
			list = append(list, firewall)
		}

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

func (g FirewallGenerator) createResources(firewallList []godo.Firewall) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, firewall := range firewallList {
		resources = append(resources, terraform_utils.NewSimpleResource(
			firewall.ID,
			firewall.Name,
			"digitalocean_firewall",
			"digitalocean",
			[]string{}))
	}
	return resources
}

func (g *FirewallGenerator) InitResources() error {
	client := g.generateClient()
	output, err := g.listFirewalls(context.TODO(), client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
