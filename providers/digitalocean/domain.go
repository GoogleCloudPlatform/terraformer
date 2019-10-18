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

type DomainGenerator struct {
	DigitalOceanService
}

func (g DomainGenerator) listDomains(ctx context.Context, client *godo.Client) ([]godo.Domain, error) {
	list := []godo.Domain{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		domains, resp, err := client.Domains.List(ctx, opt)
		if err != nil {
			return nil, err
		}

		for _, domain := range domains {
			list = append(list, domain)
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

func (g DomainGenerator) createResources(domainList []godo.Domain) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, domain := range domainList {
		resources = append(resources, terraform_utils.NewSimpleResource(
			domain.Name,
			domain.Name,
			"digitalocean_domain",
			"digitalocean",
			[]string{}))
	}
	return resources
}

func (g *DomainGenerator) InitResources() error {
	client := g.generateClient()
	output, err := g.listDomains(context.TODO(), client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
