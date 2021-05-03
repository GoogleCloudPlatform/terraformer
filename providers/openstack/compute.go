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
	"github.com/zclconf/go-cty/cty"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

type ComputeGenerator struct {
	OpenStackService
}

// createResources iterate on all openstack_compute_instance_v2
func (g *ComputeGenerator) createResources(list *pagination.Pager, volclient *gophercloud.ServiceClient) []terraformutils.Resource {
	resources := []terraformutils.Resource{}

	err := list.EachPage(func(page pagination.Page) (bool, error) {
		servers, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}

		for _, s := range servers {
			var bds = []map[string]interface{}{}
			var vol []volumes.Volume
			t := map[string]interface{}{}
			if volclient != nil {
				for _, av := range s.AttachedVolumes {
					onevol, err := volumes.Get(volclient, av.ID).Extract()
					if err == nil {
						vol = append(vol, *onevol)
					}
				}

				sort.SliceStable(vol, func(i, j int) bool {
					return vol[i].Attachments[0].Device < vol[j].Attachments[0].Device
				})

				var bindex = 0
				var dependsOn = ""
				for _, v := range vol {
					if v.Bootable == "true" && v.VolumeImageMetadata != nil {
						bds = append(bds, map[string]interface{}{
							"source_type":           "image",
							"uuid":                  v.VolumeImageMetadata["image_id"],
							"volume_size":           strconv.Itoa(v.Size),
							"boot_index":            strconv.Itoa(bindex),
							"destination_type":      "volume",
							"delete_on_termination": "false",
						})
						bindex++
					} else {
						tv := map[string]interface{}{}
						if dependsOn != "" {
							tv["depends_on"] = []string{dependsOn}
						}

						name := s.Name + strings.ReplaceAll(v.Attachments[0].Device, "/dev/", "")
						rid := s.ID + "/" + v.ID
						resource := terraformutils.NewResource(
							rid,
							name,
							"openstack_compute_volume_attach_v2",
							"openstack",
							map[string]string{},
							[]string{},
							tv,
						)
						dependsOn = "openstack_compute_volume_attach_v2." + terraformutils.TfSanitize(name)
						tv["instance_name"] = terraformutils.TfSanitize(s.Name)
						if v.Name == "" {
							v.Name = v.ID
						}
						tv["volume_name"] = terraformutils.TfSanitize(v.Name)
						resources = append(resources, resource)
					}
				}
			}

			if len(bds) > 0 {
				t = map[string]interface{}{"block_device": bds}
			}

			resource := terraformutils.NewResource(
				s.ID,
				s.Name,
				"openstack_compute_instance_v2",
				"openstack",
				map[string]string{},
				[]string{},
				t,
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
	volclient, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{
		Region: g.GetArgs()["region"].(string)})
	if err != nil {
		log.Println("VolumeImageMetadata requires blockStorage API v3")
		volclient = nil
	}
	g.Resources = g.createResources(&list, volclient)

	return nil
}

func (g *ComputeGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.Address.Type == "openstack_compute_volume_attach_v2" {
			r.SetStateAttr("volume_id", cty.StringVal("${openstack_blockstorage_volume_v3."+r.AdditionalFields["volume_name"].(string)+".id}"))
			r.SetStateAttr("instance_id", cty.StringVal("${openstack_compute_instance_v2."+r.AdditionalFields["instance_name"].(string)+".id}"))
			r.DeleteStateAttr("volume_name")
			r.DeleteStateAttr("instance_name")
			r.DeleteStateAttr("device")
		}
		if r.Address.Type != "openstack_compute_instance_v2" {
			continue
		}

		// Copy "all_metadata" to "metadata"
		if r.HasStateAttr("all_metadata") {
			metadata := r.GetStateAttrMap("metadata")
			for k, v := range r.GetStateAttrMap("all_metadata") {
				metadata[k] = v
			}
			r.SetStateAttr("metadata", cty.ObjectVal(metadata))
		}
		if r.AdditionalFields["block_device"] != nil {
			bds := r.AdditionalFields["block_device"].([]map[string]interface{})
			var blockDevices []cty.Value
			for _, bd := range bds {
				var blockDeviceMap = map[string]cty.Value{}
				for k, v := range bd {
					blockDeviceMap[k] = cty.StringVal(v.(string))
				}
				blockDevices = append(blockDevices, cty.ObjectVal(blockDeviceMap))
			}
			r.SetStateAttr("block_device", cty.ListVal(blockDevices))
		}
	}

	return nil
}
