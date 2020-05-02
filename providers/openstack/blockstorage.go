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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/pkg/errors"
)

var resourceType = map[string]string{
	"volumev3": "openstack_blockstorage_volume_v3",
	"volumev2": "openstack_blockstorage_volume_v2",
	"volume":   "openstack_blockstorage_volume_v1",
}

type BlockStorageGenerator struct {
	OpenStackService
}

// createResources iterate on all openstack_blockstorage_volume
func (g *BlockStorageGenerator) createResources(list *pagination.Pager, clientType string) []terraformutils.Resource {
	resources := []terraformutils.Resource{}

	err := list.EachPage(func(page pagination.Page) (bool, error) {
		volumes, err := volumes.ExtractVolumes(page)
		if err != nil {
			return false, err
		}

		for _, v := range volumes {
			// Use volume ID as a name if the volume doesn't have a name
			name := v.Name
			if v.Name == "" {
				name = v.ID
			}

			resource := terraformutils.NewSimpleResource(
				v.ID,
				name,
				resourceType[clientType],
				"openstack",
				[]string{},
			)

			resources = append(resources, resource)
		}

		return true, nil
	})
	if err != nil {
		log.Println(err)
	}

	return resources
}

// Creates a BlockStorage ServiceClient
func newBlockStorageClent(provider *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	// Choose v3 client firstly
	if client, err := openstack.NewBlockStorageV3(provider, eo); err == nil {
		log.Println("Using BlockStorage API v3")
		return client, nil
	}

	// if it can't initialize v3 client, try to initialize v2 client
	if client, err := openstack.NewBlockStorageV2(provider, eo); err == nil {
		log.Println("Using BlockStorage API v2")
		return client, nil
	}

	// if it can't initialize v2 client, try to initialize v1 client
	if client, err := openstack.NewBlockStorageV1(provider, eo); err == nil {
		log.Println("Using BlockStorage API v1")
		return client, nil
	}

	return nil, errors.New("Failed to initialize BlockStorage client")
}

// Generate TerraformResources from OpenStack API,
func (g *BlockStorageGenerator) InitResources() error {
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return err
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return err
	}

	eo := gophercloud.EndpointOpts{
		Region: g.GetArgs()["region"].(string),
	}

	client, err := newBlockStorageClent(provider, eo)
	if err != nil {
		return err
	}

	list := volumes.List(client, nil)

	g.Resources = g.createResources(&list, client.Type)

	return nil
}
