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
	"sort"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/volumeattach"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

type ComputeGenerator struct {
	OpenStackService
}

// createResources iterate on all openstack_compute_instance_v2
func (g *ComputeGenerator) createResources(list *pagination.Pager, client_block *gophercloud.ServiceClient, client *gophercloud.ServiceClient) []terraform_utils.Resource {
	var importblockdevicestoinstance = g.GetArgs()["importblockdevicestoinstance"]
	resources := []terraform_utils.Resource{}

	var done_volume_ids = []string{}
	var volumes_list =[]map[string]interface{}{}

	if importblockdevicestoinstance == "true" {
		volumelist := volumes.List(client_block,nil)
		volumelist.EachPage(func(page pagination.Page) (bool, error) {
			volume_list, err := volumes.ExtractVolumes(page)
			if err != nil {
				return false, err
			}
			for vnumber,v := range volume_list {

				volumes_list = append(volumes_list,map[string]interface{}{"vnumber":vnumber,"volume":v,"id":v.ID})
			}
			return true,nil
		})
	}

	list.EachPage(func(page pagination.Page) (bool, error) {
		server_list, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}

		for i, s := range server_list {
			var abds =[]map[string]interface{}{}
			if importblockdevicestoinstance == "true" {
				clientType := client_block.Type
				var boot_index = 0
				var depends_on = ""
				attachdetaillist := volumeattach.List(client, s.ID)

				attachdetaillist.EachPage(func(page pagination.Page) (bool, error) {
					attachments, err := volumeattach.ExtractVolumeAttachments(page)
					if err != nil {
						return false, err
					}

					sort.SliceStable(attachments, func(i, j int) bool {
							return attachments[i].Device  < attachments[j].Device
						})
					for _,attachment := range attachments {
						var v volumes.Volume
						var vnumber = 0

						for _,vol  := range volumes_list {
							if (vol["id"].(string) == attachment.ID) {
								v = vol["volume"].(volumes.Volume)
								vnumber = vol["vnumber"].(int)

								name := "server" + strconv.Itoa(i+1) + strings.Replace(attachment.Device,"/dev/","",-1)
								tv := map[string]interface{}{"volume_resource":  "${"+resourceType[clientType]+".tfer--vol" + strconv.Itoa(vnumber+1)+".id}",
								"server_resource": "${openstack_compute_instance_v2.tfer--" + "server" + strconv.Itoa(i+1)+".id}"}
								if depends_on != "" {
									tv["depends_on"] = []string{depends_on}
								}

								if v.Bootable != "true" {
									rid := s.ID + "/" + attachment.VolumeID
									resource := terraform_utils.NewResource(
									rid,
									name,
									"openstack_compute_volume_attach_v2",
									"openstack",
									map[string]string{},
									[]string{},
									tv,
									)
									resources = append(resources, resource)
									depends_on = "openstack_compute_volume_attach_v2.tfer--" + name
								} else if v.Bootable == "true" {
									var abd = map[string]interface{}{}
									if v.VolumeImageMetadata["image_id"] != "" {
										abd["source_type"] = "image"
										abd["uuid"] = v.VolumeImageMetadata["image_id"]
									} else {
										abd["source_type"] = "blank"
									}
									abd["volume_size"] = strconv.Itoa(v.Size)
									abd["boot_index"] = strconv.Itoa(boot_index)
									abd["destination_type"] = "volume"
									abd["delete_on_termination"] = "false"
									delete(abd, "id")
									boot_index++
									abds = append(abds, abd)
									done_volume_ids = append(done_volume_ids,v.ID)
								}
							}
						}
					}

				return true, nil

				})
				for _,vol := range volumes_list {
					v := vol["volume"].(volumes.Volume)
					seen := false
					for _,dv := range done_volume_ids {
						if dv == v.ID {
							seen = true
						}
					}
					if !seen  {
						resource := terraform_utils.NewSimpleResource(
						v.ID,
						"vol"+strconv.Itoa(vol["vnumber"].(int) +1),
						resourceType[clientType],
						"openstack",
						[]string{},
						)

						resources = append(resources, resource)
						done_volume_ids = append(done_volume_ids,v.ID)

					}
				}
			}
			if importblockdevicestoinstance == "true" && len(abds)>0 {
				t := map[string]interface{}{"block_device": abds}

				resource := terraform_utils.NewResource(
					s.ID,
					"server"+strconv.Itoa(i+1),
					"openstack_compute_instance_v2",
					"openstack",
					map[string]string{},
					[]string{},
					t,
				)
				resources = append(resources, resource)

			} else {
				name := s.Name
				if importblockdevicestoinstance == "true" {
				name = "server"+strconv.Itoa(i+1)
				}
				resource := terraform_utils.NewSimpleResource(
					s.ID,
					name,
					"openstack_compute_instance_v2",
					"openstack",
					[]string{},
				)

				resources = append(resources, resource)
			}

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
	var client_block *gophercloud.ServiceClient = nil
	if g.GetArgs()["importblockdevicestoinstance"]=="true" {
		client_block, err = newBlockStorageClent(provider, gophercloud.EndpointOpts{
			Region: g.GetArgs()["region"].(string),
		})
	}
	if err != nil {
		return err
	}

	list := servers.List(client, nil)
	g.Resources = g.createResources(&list, client_block, client)

	return nil
}

func (g *ComputeGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type == "openstack_compute_volume_attach_v2" {
			g.Resources[i].Item["volume_id"] = r.AdditionalFields["volume_resource"].(string)
			g.Resources[i].Item["instance_id"] = r.AdditionalFields["server_resource"].(string)
			delete(g.Resources[i].Item, "volume_resource")
			delete(g.Resources[i].Item, "server_resource")
			delete(g.Resources[i].Item, "device")
		}
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
		var importblockdevicestoinstance = g.GetArgs()["importblockdevicestoinstance"]
		if importblockdevicestoinstance == "true" && r.AdditionalFields["block_device"] != nil {
			bds := r.AdditionalFields["block_device"].([]map[string]interface{})
			for bi, bd := range bds {
				for k, v := range bd {
					g.Resources[i].InstanceState.Attributes["block_device."+strconv.Itoa(bi)+"."+k] = v.(string)
				}
			}

			g.Resources[i].InstanceState.Attributes["block_device.#"] = strconv.Itoa(len(bds))
		}
	}

	return nil
}
