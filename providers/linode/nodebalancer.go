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

package linode

import (
	"context"
	"strconv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/linode/linodego"
)

type NodeBalancerGenerator struct {
	LinodeService
}

func (g *NodeBalancerGenerator) loadNodeBalancers(client linodego.Client) ([]linodego.NodeBalancer, error) {
	nodeBalancerList, err := client.ListNodeBalancers(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	for _, nodeBalancer := range nodeBalancerList {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			strconv.Itoa(nodeBalancer.ID),
			strconv.Itoa(nodeBalancer.ID),
			"linode_nodebalancer",
			"linode",
			[]string{}))
	}
	return nodeBalancerList, nil
}

func (g *NodeBalancerGenerator) loadNodeBalancerConfigs(client linodego.Client, nodebalancerID int) ([]linodego.NodeBalancerConfig, error) {
	nodeBalancerConfigList, err := client.ListNodeBalancerConfigs(context.Background(), nodebalancerID, nil)
	if err != nil {
		return nil, err
	}
	for _, nodeBalancerConfig := range nodeBalancerConfigList {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			strconv.Itoa(nodeBalancerConfig.ID),
			strconv.Itoa(nodeBalancerConfig.ID),
			"linode_nodebalancer_config",
			"linode",
			map[string]string{"nodebalancer_id": strconv.Itoa(nodebalancerID)},
			[]string{},
			map[string]interface{}{}))
	}
	return nodeBalancerConfigList, nil
}

func (g *NodeBalancerGenerator) loadNodeBalancerNodes(client linodego.Client, nodebalancerID int, nodebalancerConfigID int) error {
	nodeBalancerNodeList, err := client.ListNodeBalancerNodes(context.Background(), nodebalancerID, nodebalancerConfigID, nil)
	if err != nil {
		return err
	}
	for _, nodeBalancerNode := range nodeBalancerNodeList {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			strconv.Itoa(nodeBalancerNode.ID),
			strconv.Itoa(nodeBalancerNode.ID),
			"linode_nodebalancer_node",
			"linode",
			map[string]string{
				"nodebalancer_id": strconv.Itoa(nodebalancerID),
				"config_id":       strconv.Itoa(nodebalancerConfigID),
			},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}

func (g *NodeBalancerGenerator) InitResources() error {
	client := g.generateClient()
	nodeBalancerList, err := g.loadNodeBalancers(client)
	if err != nil {
		return err
	}
	for _, nodeBalancer := range nodeBalancerList {
		nodeBalancerConfigList, err := g.loadNodeBalancerConfigs(client, nodeBalancer.ID)
		if err != nil {
			return err
		}
		for _, nodeBalancerConfig := range nodeBalancerConfigList {
			err := g.loadNodeBalancerNodes(client, nodeBalancer.ID, nodeBalancerConfig.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
