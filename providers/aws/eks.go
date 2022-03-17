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

package aws

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

var eksAllowEmptyValues = []string{"tags."}

type EksGenerator struct {
	AWSService
}

func (g *EksGenerator) getNodeGroups(clusterName string, svc *eks.Client) error {
	p := eks.NewListNodegroupsPaginator(svc, &eks.ListNodegroupsInput{
		ClusterName: &clusterName,
	})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, nodeGroupName := range page.Nodegroups {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				fmt.Sprintf("%s:%s", clusterName, nodeGroupName),
				nodeGroupName,
				"aws_eks_node_group",
				"aws",
				eksAllowEmptyValues,
			))
		}
	}
	return nil
}

func (g *EksGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := eks.NewFromConfig(config)
	p := eks.NewListClustersPaginator(svc, &eks.ListClustersInput{})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, clusterName := range page.Clusters {
			err := g.getNodeGroups(clusterName, svc)
			if err != nil {
				return err
			}
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				clusterName,
				clusterName,
				"aws_eks_cluster",
				"aws",
				eksAllowEmptyValues,
			))
		}
	}
	return nil
}

func (g *EksGenerator) PostConvertHook() error {
	for _, resource := range g.Resources {
		if resource.InstanceInfo.Type == "aws_eks_node_group" {
			if _, ok := resource.Item["launch_template"]; ok {
				delete(resource.Item["launch_template"].([]interface{})[0].(map[string]interface{}), "id")
			}
			if _, ok := resource.Item["update_config"]; ok {
				delete(resource.Item["update_config"].([]interface{})[0].(map[string]interface{}), "max_unavailable_percentage")
			}
			for cluster := range g.Resources {
				if g.Resources[cluster].InstanceInfo.Type == "aws_eks_cluster" {
					if g.Resources[cluster].Item["name"] == resource.Item["cluster_name"] {
						resource.Item["cluster_name"] = "${aws_eks_cluster." + g.Resources[cluster].InstanceInfo.ResourceAddress().Name + ".name}"
					}
				}
			}
		}
	}
	return nil
}
