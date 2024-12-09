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
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type InstanceGenerator struct {
	ScalewayService
}

func (g InstanceGenerator) ListInstances(client *scw.Client) ([]*instance.Server, error) {
	list := []*instance.Server{}
	instanceApi := instance.NewAPI(client)

	page := int32(1)
	perpage := uint32(100)
	opt := &instance.ListServersRequest{
		Page:    &page,
		PerPage: &perpage,
	}

	for _, zone := range scw.AllZones {
		opt.Zone = zone
		for {
			resp, err := instanceApi.ListServers(opt)
			if err != nil {
				return nil, err
			}
			for _, server := range resp.Servers {
				if server != nil {
					list = append(list, server)
				}
			}
			// Exit loop when we are on the last page.
			if resp.TotalCount < *opt.PerPage {
				break
			}
			*opt.Page++
		}
	}

	return list, nil
}

func (g InstanceGenerator) createResources(instanceList []*instance.Server) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, instance := range instanceList {
		managed := false
		for _, tag := range instance.Tags {
			// Don't want to include managed instances, because they are managed by something else
			if tag == "managed=true" {
				managed = true
				break
			}
		}
		if managed {
			break
		}
		resources = append(resources, terraformutils.NewSimpleResource(
			string(instance.Zone)+"/"+instance.ID,
			instance.Name,
			"scaleway_instance_server",
			"scaleway",
			[]string{}))
	}
	return resources
}

func (g *InstanceGenerator) InitResources() error {
	client := g.generateClient()
	output, err := g.ListInstances(client)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
