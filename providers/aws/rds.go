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

	"github.com/aws/aws-sdk-go-v2/aws"
)

var RDSAllowEmptyValues = []string{"tags."}

type RDSGenerator struct {
	AWSService
}

func (g *RDSGenerator) loadDBInstances(svc *rds.Client) error {
	p := rds.NewDescribeDBInstancesPaginator(svc.DescribeDBInstancesRequest(&rds.DescribeDBInstancesInput{}))
	for p.Next(context.Background()) {
		for _, db := range p.CurrentPage().DBInstances {
			resourceName := aws.StringValue(db.DBInstanceIdentifier)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_db_instance",
				"aws",
				RDSAllowEmptyValues,
			))
		}
	}
	return p.Err()
}

func (g *RDSGenerator) loadDBParameterGroups(svc *rds.Client) error {
	p := rds.NewDescribeDBParameterGroupsPaginator(svc.DescribeDBParameterGroupsRequest(&rds.DescribeDBParameterGroupsInput{}))
	for p.Next(context.Background()) {
		for _, parameterGroup := range p.CurrentPage().DBParameterGroups {
			resourceName := aws.StringValue(parameterGroup.DBParameterGroupName)
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
	return p.Err()
}

func (g *RDSGenerator) loadDBSubnetGroups(svc *rds.Client) error {
	p := rds.NewDescribeDBSubnetGroupsPaginator(svc.DescribeDBSubnetGroupsRequest(&rds.DescribeDBSubnetGroupsInput{}))
	for p.Next(context.Background()) {
		for _, subnet := range p.CurrentPage().DBSubnetGroups {
			resourceName := aws.StringValue(subnet.DBSubnetGroupName)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_db_subnet_group",
				"aws",
				RDSAllowEmptyValues,
			))
		}
	}
	return p.Err()
}

func (g *RDSGenerator) loadOptionGroups(svc *rds.Client) error {
	p := rds.NewDescribeOptionGroupsPaginator(svc.DescribeOptionGroupsRequest(&rds.DescribeOptionGroupsInput{}))
	for p.Next(context.Background()) {
		for _, optionGroup := range p.CurrentPage().OptionGroupsList {
			resourceName := aws.StringValue(optionGroup.OptionGroupName)
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
	return p.Err()
}

func (g *RDSGenerator) loadEventSubscription(svc *rds.Client) error {
	p := rds.NewDescribeEventSubscriptionsPaginator(svc.DescribeEventSubscriptionsRequest(&rds.DescribeEventSubscriptionsInput{}))
	for p.Next(context.Background()) {
		for _, eventSubscription := range p.CurrentPage().EventSubscriptionsList {
			resourceName := aws.StringValue(eventSubscription.CustomerAwsId)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_db_event_subscription",
				"aws",
				RDSAllowEmptyValues,
			))
		}
	}
	return p.Err()
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
	svc := rds.New(config)

	if err := g.loadDBInstances(svc); err != nil {
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
		if r.InstanceInfo.Type != "aws_db_instance" {
			continue
		}
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
	}
	return nil
}
