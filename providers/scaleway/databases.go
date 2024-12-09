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
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type DatabaseGenerator struct {
	ScalewayService
}

func (g DatabaseGenerator) ListInstances(client *scw.Client) ([]*rdb.Instance, error) {
	list := []*rdb.Instance{}
	databaseApi := rdb.NewAPI(client)

	page := int32(1)
	perpage := uint32(100)
	opt := &rdb.ListInstancesRequest{
		Page:     &page,
		PageSize: &perpage,
	}

	for {
		resp, err := databaseApi.ListInstances(opt)
		if err != nil {
			return nil, err
		}
		for _, instance := range resp.Instances {
			if instance != nil {
				list = append(list, instance)
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

func (g DatabaseGenerator) createResources(databaseList []*rdb.Instance) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, instance := range databaseList {
		resources = append(resources, terraformutils.NewSimpleResource(
			string(instance.Region)+"/"+instance.ID,
			instance.Name,
			"scaleway_rdb_instance",
			"scaleway",
			[]string{}))
	}
	return resources
}

func (g *DatabaseGenerator) InitResources() error {
	client := g.generateClient()
	output, err := g.ListInstances(client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
