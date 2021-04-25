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
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/digitalocean/godo"
)

type DomainGenerator struct {
	DigitalOceanService
}

func (g *DomainGenerator) loadDomains(ctx context.Context, client *godo.Client) ([]godo.Domain, error) {
	list := []godo.Domain{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		domains, resp, err := client.Domains.List(ctx, opt)
		if err != nil {
			return nil, err
		}

		for _, domain := range domains {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				domain.Name,
				domain.Name,
				"digitalocean_domain",
				"digitalocean",
				[]string{}))
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

func (g *DomainGenerator) loadRecords(ctx context.Context, client *godo.Client, domain string) error {
	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		records, resp, err := client.Domains.Records(ctx, domain, opt)
		if err != nil {
			return err
		}

		for _, record := range records {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				strconv.Itoa(record.ID),
				strconv.Itoa(record.ID),
				"digitalocean_record",
				"digitalocean",
				map[string]string{"domain": domain},
				[]string{},
				map[string]interface{}{}))
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}

	return nil
}

func (g *DomainGenerator) InitResources() error {
	client := g.generateClient()
	domains, err := g.loadDomains(context.TODO(), client)
	if err != nil {
		return err
	}
	for _, domain := range domains {
		err := g.loadRecords(context.TODO(), client, domain.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
