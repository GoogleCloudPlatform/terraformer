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

	"golang.org/x/exp/slices"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/bizflycloud/gobizfly"
)

type KubernetesClusterGenerator struct {
	BizflyCloudService
}

func (g *KubernetesClusterGenerator) loadKubernetesClusters(ctx context.Context, client *gobizfly.Client) ([]*gobizfly.FullCluster, error) {
	list := []*gobizfly.FullCluster{}

	// create options. initially, these will be blank
	opt := &gobizfly.ListOptions{}
	clusters, err := client.KubernetesEngine.List(ctx, opt)
	if err != nil {
		return nil, err
	}

	networkIDs := []string{}
	networkNames := map[string]string{}
	for _, _cluster := range clusters {
		cluster, err := client.KubernetesEngine.Get(ctx, _cluster.UID)
		if err != nil {
			return nil, err
		}

		if !slices.Contains(networkIDs, cluster.VPCNetworkID) {
			net, err := client.CloudServer.VPCNetworks().Get(ctx, cluster.VPCNetworkID)
			if err != nil {
				return nil, err
			}

			networkIDs = append(networkIDs, cluster.VPCNetworkID)
			networkNames[net.ID] = net.Name
			networkName := net.Name
			if networkName == "" {
				networkName = net.ID
			}

			g.Resources = append(g.Resources, terraformutils.NewResource(
				net.ID,
				networkName,
				"bizflycloud_vpc_network",
				"bizflycloud",
				map[string]string{},
				[]string{},
				map[string]interface{}{"is_default": net.IsDefault}))
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			cluster.UID,
			cluster.Name,
			"bizflycloud_kubernetes",
			"bizflycloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{"vpc_network_id": "${bizflycloud_vpc_network.tfer--" + networkNames[cluster.VPCNetworkID] + ".id}"}))
		list = append(list, cluster)
	}

	return list, nil
}

func (g *KubernetesClusterGenerator) InitResources() error {
	client := g.generateClient()
	_, err := g.loadKubernetesClusters(context.TODO(), client)
	if err != nil {
		return err
	}

	return nil
}
