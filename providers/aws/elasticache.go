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

package aws

import (
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

var elastiCacheAllowEmptyValues = []string{"tags."}

type ElastiCacheGenerator struct {
	AWSService
}

func (g *ElastiCacheGenerator) loadCacheClusters(svc *elasticache.Client) error {
	p := elasticache.NewDescribeCacheClustersPaginator(svc, &elasticache.DescribeCacheClustersInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, cluster := range page.CacheClusters {
			resourceName := StringValue(cluster.CacheClusterId)
			resource := terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_elasticache_cluster",
				"aws",
				elastiCacheAllowEmptyValues,
			)
			// redis only - if cluster has Replication Group not need next attributes.
			// terraform-aws provider has ConflictsWith on ReplicationGroupId with all next attributes,
			// but return all attributes on refresh :(
			// https://github.com/terraform-providers/terraform-provider-aws/blob/master/aws/resource_aws_elasticache_cluster.go#L167
			if StringValue(cluster.ReplicationGroupId) != "" {
				resource.IgnoreKeys = append(resource.IgnoreKeys,
					"^availability_zones$",
					"^az_mode$",
					"^engine_version$",
					"^engine$",
					"^maintenance_window$",
					"^node_type$",
					"^notification_topic_arn$",
					"^num_cache_nodes$",
					"^parameter_group_name$",
					"^port$",
					"^security_group_ids.(.*)",
					"^security_group_names$",
					"^snapshot_arns$",
					"^snapshot_name$",
					"^snapshot_retention_limit$",
					"^snapshot_window$",
					"^subnet_group_name$",
				)
			}
			g.Resources = append(g.Resources, resource)
		}
	}
	return nil
}

func (g *ElastiCacheGenerator) loadParameterGroups(svc *elasticache.Client) error {
	p := elasticache.NewDescribeCacheParameterGroupsPaginator(svc, &elasticache.DescribeCacheParameterGroupsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, parameterGroup := range page.CacheParameterGroups {
			resourceName := StringValue(parameterGroup.CacheParameterGroupName)
			if strings.Contains(resourceName, ".") {
				continue // skip default Default ParameterGroups like default.redis5.0
			}
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_elasticache_parameter_group",
				"aws",
				elastiCacheAllowEmptyValues,
			))
		}
	}
	return nil
}

func (g *ElastiCacheGenerator) loadSubnetGroups(svc *elasticache.Client) error {
	p := elasticache.NewDescribeCacheSubnetGroupsPaginator(svc, &elasticache.DescribeCacheSubnetGroupsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, subnet := range page.CacheSubnetGroups {
			resourceName := StringValue(subnet.CacheSubnetGroupName)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_elasticache_subnet_group",
				"aws",
				elastiCacheAllowEmptyValues,
			))
		}
	}
	return nil
}

func (g *ElastiCacheGenerator) loadReplicationGroups(svc *elasticache.Client) error {
	p := elasticache.NewDescribeReplicationGroupsPaginator(svc, &elasticache.DescribeReplicationGroupsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, replicationGroup := range page.ReplicationGroups {
			resourceName := StringValue(replicationGroup.ReplicationGroupId)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_elasticache_replication_group",
				"aws",
				elastiCacheAllowEmptyValues,
			))
		}
	}
	return nil
}

// Generate TerraformResources from AWS API,
// from each database create 1 TerraformResource.
// Need only database name as ID for terraform resource
// AWS api support paging
func (g *ElastiCacheGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := elasticache.NewFromConfig(config)

	if err := g.loadCacheClusters(svc); err != nil {
		return err
	}
	if err := g.loadParameterGroups(svc); err != nil {
		return err
	}
	if err := g.loadReplicationGroups(svc); err != nil {
		return err
	}
	if err := g.loadSubnetGroups(svc); err != nil {
		return err
	}

	return nil
}

func (g *ElastiCacheGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_elasticache_cluster" {
			continue
		}
		for _, parameterGroup := range g.Resources {
			if parameterGroup.InstanceInfo.Type != "aws_elasticache_parameter_group" {
				continue
			}
			if parameterGroup.InstanceState.Attributes["name"] == r.InstanceState.Attributes["parameter_group_name"] {
				if strings.HasPrefix(parameterGroup.InstanceState.Attributes["family"], r.InstanceState.Attributes["engine"]) {
					g.Resources[i].Item["parameter_group_name"] = "${aws_elasticache_parameter_group." + parameterGroup.ResourceName + ".name}"
				}
			}
		}

		for _, subnet := range g.Resources {
			if subnet.InstanceInfo.Type != "aws_elasticache_subnet_group" {
				continue
			}
			if subnet.InstanceState.Attributes["name"] == r.Item["subnet_group_name"] {
				g.Resources[i].Item["subnet_group_name"] = "${aws_elasticache_subnet_group." + subnet.ResourceName + ".name}"
			}
		}

		for _, replicationGroup := range g.Resources {
			if replicationGroup.InstanceInfo.Type != "aws_elasticache_replication_group" {
				continue
			}
			if replicationGroup.InstanceState.Attributes["replication_group_id"] == r.InstanceState.Attributes["replication_group_id"] {
				g.Resources[i].Item["replication_group_id"] = "${aws_elasticache_replication_group." + replicationGroup.ResourceName + ".replication_group_id}"
			}
		}
	}
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_elasticache_replication_group" {
			continue
		}
		for _, subnet := range g.Resources {
			if subnet.InstanceInfo.Type != "aws_elasticache_subnet_group" {
				continue
			}
			if subnet.InstanceState.Attributes["name"] == r.InstanceState.Attributes["subnet_group_name"] {
				g.Resources[i].Item["subnet_group_name"] = "${aws_elasticache_subnet_group." + subnet.ResourceName + ".name}"
			}
		}
	}
	return nil
}
