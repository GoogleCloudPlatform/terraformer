// Copyright 2018 The Terraformer Authors.
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

package openstack

import (
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

type ComputeGenerator struct {
	OpenStackService
}

// createResources iterate on all openstack_compute_instance_v2
func (g *ComputeGenerator) createResources(list *pagination.Pager) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}

	list.EachPage(func(page pagination.Page) (bool, error) {
		servers, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}

		for _, s := range servers {
			resource := terraform_utils.NewResource(
				s.ID,
				s.Name,
				"openstack_compute_instance_v2",
				"openstack",
				map[string]string{},
				[]string{},
				map[string]string{},
			)

			resources = append(resources, resource)
		}

		return true, nil
	})

	return resources
}

// Generate TerraformResources from OpenStack API,
func (g *ComputeGenerator) InitResources() error {
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return err
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return err
	}

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: g.GetArgs()["region"].(string),
	})
	if err != nil {
		return err
	}

	list := servers.List(client, nil)

	g.Resources = g.createResources(&list)
	g.PopulateIgnoreKeys()

	return nil
}

func (g *ComputeGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "openstack_compute_instance_v2" {
			continue
		}

		// Copy "all_metadata.%" to "metadata.%"
		for k, v := range g.Resources[i].InstanceState.Attributes {
			if strings.HasPrefix(k, "all_metadata") {
				newKey := strings.Replace(k, "all_metadata", "metadata", 1)
				g.Resources[i].InstanceState.Attributes[newKey] = v
			}
		}
		// Replace "all_metadata" to "metadata"
		// because "all_metadata" field cannot be set as resource argument
		for k, v := range g.Resources[i].Item {
			if strings.HasPrefix(k, "all_metadata") {
				newKey := strings.Replace(k, "all_metadata", "metadata", 1)
				g.Resources[i].Item[newKey] = v
				delete(g.Resources[i].Item, k)
			}
		}
	}

	return nil
}
