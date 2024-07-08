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

package bizflycloud

import (
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/bizflycloud/gobizfly"
)

type CloudDatabaseGenerator struct {
	BizflyCloudService
}

func (g *CloudDatabaseGenerator) loadDatabaseInstances(ctx context.Context, client *gobizfly.Client) ([]*gobizfly.CloudDatabaseInstance, error) {
	list := []*gobizfly.CloudDatabaseInstance{}

	// create options. initially, these will be blank
	opt := &gobizfly.CloudDatabaseListOption{}
	for {
		instances, err := client.CloudDatabase.Instances().List(ctx, opt)
		if err != nil {
			return nil, err
		}

		nets, err := client.CloudServer.VPCNetworks().List(ctx)
		if err != nil {
			return nil, err
		}

		for _, instance := range instances {
			flavorName := ""
			availabilityZone := ""
			quantity := 0
			_availabilityZone := ""
			networkIDs := []string{}

			for _, node := range instance.Nodes {
				if node.Role == "primary" {
					availabilityZone = node.AvailabilityZone

					_flavor := strings.Replace(node.Flavor, "nix.", "", -1)
					_flavor = strings.Replace(_flavor, "_basic", "", -1)
					_flavor = strings.Replace(_flavor, "_dedicated", "", -1)
					_flavor = strings.Replace(_flavor, "_enterprise", "", -1)
					flavorName = _flavor

					for _, network := range node.Addresses.Private {
						for _, net := range nets {
							if net.Name == network.Network {
								networkIDs = append(networkIDs, net.ID)
							}
						}
					}
				}

				if node.Role != "primary" {
					quantity = quantity + 1
					_availabilityZone = node.AvailabilityZone
				}
			}

			additionalFields := map[string]interface{}{
				"flavor_name":       flavorName,
				"availability_zone": availabilityZone,
				"network_ids":       networkIDs,
			}
			if quantity > 0 {
				additionalFields["secondaries"] = map[string]interface{}{
					"availability_zone": _availabilityZone,
					"quantity":          quantity,
				}
			}

			instanceAutoScalingEnable := 0
			if instance.AutoScaling.Enable {
				instanceAutoScalingEnable = 1
			}
			additionalFields["autoscaling"] = map[string]interface{}{
				"enable":           instanceAutoScalingEnable,
				"volume_limited":   instance.AutoScaling.Volume.Limited,
				"volume_threshold": instance.AutoScaling.Volume.Threshold,
			}

			additionalFields["volume_size"] = instance.Volume.Size

			g.Resources = append(g.Resources, terraformutils.NewResource(
				instance.ID,
				instance.Name,
				"bizflycloud_cloud_database_instance",
				"bizflycloud",
				map[string]string{},
				[]string{},
				additionalFields))
			list = append(list, instance)
		}
		break
	}

	return list, nil
}

func (g *CloudDatabaseGenerator) InitResources() error {
	client := g.generateClient()
	_, err := g.loadDatabaseInstances(context.TODO(), client)
	if err != nil {
		return err
	}
	return nil
}
