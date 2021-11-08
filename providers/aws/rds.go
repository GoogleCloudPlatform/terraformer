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

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

var RDSAllowEmptyValues = []string{"tags."}

type RDSGenerator struct {
	AWSService
}

func (g *RDSGenerator) loadDBClusters(svc *rds.Client) error {
	p := rds.NewDescribeDBClustersPaginator(svc, &rds.DescribeDBClustersInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, cluster := range page.DBClusters {
			resourceName := StringValue(cluster.DBClusterIdentifier)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_rds_cluster",
				"aws",
				RDSAllowEmptyValues,
			))
		}
	}
	return nil
}

func (g *RDSGenerator) loadDBProxies(svc *rds.Client) error {
	p := rds.NewDescribeDBProxiesPaginator(svc, &rds.DescribeDBProxiesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, db := range page.DBProxies {
			resourceName := StringValue(db.DBProxyName)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_db_proxy",
				"aws",
				RDSAllowEmptyValues,
			))
		}
	}
	return nil

}
func (g *RDSGenerator) loadDBInstances(svc *rds.Client) error {
	p := rds.NewDescribeDBInstancesPaginator(svc, &rds.DescribeDBInstancesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, db := range page.DBInstances {
			resourceName := StringValue(db.DBInstanceIdentifier)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_db_instance",
				"aws",
				RDSAllowEmptyValues,
			))
		}
	}
	return nil
}

func (g *RDSGenerator) loadDBParameterGroups(svc *rds.Client) error {
	p := rds.NewDescribeDBParameterGroupsPaginator(svc, &rds.DescribeDBParameterGroupsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, parameterGroup := range page.DBParameterGroups {
			resourceName := StringValue(parameterGroup.DBParameterGroupName)
			if strings.Contains(resourceName, ".") {
				continue // skip default Default ParameterGroups like default.mysql5.6
			}
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_db_parameter_group",
				"aws",
				RDSAllowEmptyValues,
			))
		}
	}
	return nil
}

func (g *RDSGenerator) loadDBSubnetGroups(svc *rds.Client) error {
	p := rds.NewDescribeDBSubnetGroupsPaginator(svc, &rds.DescribeDBSubnetGroupsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, subnet := range page.DBSubnetGroups {
			resourceName := StringValue(subnet.DBSubnetGroupName)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_db_subnet_group",
				"aws",
				RDSAllowEmptyValues,
			))
		}
	}
	return nil
}

func (g *RDSGenerator) loadOptionGroups(svc *rds.Client) error {
	p := rds.NewDescribeOptionGroupsPaginator(svc, &rds.DescribeOptionGroupsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, optionGroup := range page.OptionGroupsList {
			resourceName := StringValue(optionGroup.OptionGroupName)
			if strings.Contains(resourceName, ".") || strings.Contains(resourceName, ":") {
				continue // skip default Default OptionGroups like default.mysql5.6
			}
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_db_option_group",
				"aws",
				RDSAllowEmptyValues,
			))
		}
	}
	return nil
}

func (g *RDSGenerator) loadEventSubscription(svc *rds.Client) error {
	p := rds.NewDescribeEventSubscriptionsPaginator(svc, &rds.DescribeEventSubscriptionsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, eventSubscription := range page.EventSubscriptionsList {
			resourceName := StringValue(eventSubscription.CustomerAwsId)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_db_event_subscription",
				"aws",
				RDSAllowEmptyValues,
			))
		}
	}
	return nil
}

// Generate TerraformResources from AWS API,
// from each database create 1 TerraformResource.
// Need only database name as ID for terraform resource
// AWS api support paging
func (g *RDSGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := rds.NewFromConfig(config)

	if err := g.loadDBClusters(svc); err != nil {
		return err
	}
	if err := g.loadDBInstances(svc); err != nil {
		return err
	}
	if err := g.loadDBProxies(svc); err != nil {
		return err
	}
	if err := g.loadDBParameterGroups(svc); err != nil {
		return err
	}
	if err := g.loadDBSubnetGroups(svc); err != nil {
		return err
	}
	if err := g.loadOptionGroups(svc); err != nil {
		return err
	}

	if err := g.loadEventSubscription(svc); err != nil {
		return err
	}

	return nil
}

func (g *RDSGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type == "aws_db_instance" || r.InstanceInfo.Type == "aws_rds_cluster" {

			for _, parameterGroup := range g.Resources {
				if parameterGroup.InstanceInfo.Type != "aws_db_parameter_group" {
					continue
				}
				if parameterGroup.InstanceState.Attributes["name"] == r.InstanceState.Attributes["parameter_group_name"] {
					g.Resources[i].Item["parameter_group_name"] = "${aws_db_parameter_group." + parameterGroup.ResourceName + ".name}"
				}
			}

			for _, subnet := range g.Resources {
				if subnet.InstanceInfo.Type != "aws_db_subnet_group" {
					continue
				}
				if subnet.InstanceState.Attributes["name"] == r.InstanceState.Attributes["db_subnet_group_name"] {
					g.Resources[i].Item["db_subnet_group_name"] = "${aws_db_subnet_group." + subnet.ResourceName + ".name}"
				}
			}

			for _, optionGroup := range g.Resources {
				if optionGroup.InstanceInfo.Type != "aws_db_option_group" {
					continue
				}
				if optionGroup.InstanceState.Attributes["name"] == r.InstanceState.Attributes["option_group_name"] {
					g.Resources[i].Item["option_group_name"] = "${aws_db_option_group." + optionGroup.ResourceName + ".name}"
				}
			}
		} else {
			continue
		}
	}
	return nil
}
