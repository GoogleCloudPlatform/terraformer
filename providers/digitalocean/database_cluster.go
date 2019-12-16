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
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/digitalocean/godo"
)

type DatabaseClusterGenerator struct {
	DigitalOceanService
}

func (g *DatabaseClusterGenerator) loadDatabaseClusters(ctx context.Context, client *godo.Client) ([]godo.Database, error) {
	list := []godo.Database{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		clusters, resp, err := client.Databases.List(ctx, opt)
		if err != nil {
			return nil, err
		}

		for _, cluster := range clusters {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				cluster.ID,
				cluster.Name,
				"digitalocean_database_cluster",
				"digitalocean",
				[]string{}))
			list = append(list, cluster)
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

func (g *DatabaseClusterGenerator) loadDatabaseConnectionPools(ctx context.Context, client *godo.Client, clusterID string) error {
	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		pools, resp, err := client.Databases.ListPools(ctx, clusterID, opt)
		if err != nil {
			return err
		}

		for _, pool := range pools {
			g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
				fmt.Sprintf("%s/%s", clusterID, pool.Name),
				pool.Name,
				"digitalocean_database_connection_pool",
				"digitalocean",
				[]string{}))
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

func (g *DatabaseClusterGenerator) InitResources() error {
	client := g.generateClient()
	clusters, err := g.loadDatabaseClusters(context.TODO(), client)
	if err != nil {
		return err
	}
	for _, cluster := range clusters {
		err := g.loadDatabaseConnectionPools(context.TODO(), client, cluster.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
