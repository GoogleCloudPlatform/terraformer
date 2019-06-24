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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/service/rds"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var RDSAllowEmptyValues = []string{"tags."}

type RDSGenerator struct {
	AWSService
}

func (g *RDSGenerator) loadDBInstances(svc *rds.RDS) error {
	return svc.DescribeDBInstancesPages(&rds.DescribeDBInstancesInput{}, func(dbInstances *rds.DescribeDBInstancesOutput, lastPage bool) bool {
		for _, db := range dbInstances.DBInstances {
			resourceName := aws.StringValue(db.DBInstanceIdentifier)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_db_instance",
				"aws",
				map[string]string{},
				RDSAllowEmptyValues,
				map[string]string{},
			))
		}
		return !lastPage
	})

}

func (g *RDSGenerator) loadDBParameterGroups(svc *rds.RDS) error {
	return svc.DescribeDBParameterGroupsPages(&rds.DescribeDBParameterGroupsInput{}, func(parameterGroups *rds.DescribeDBParameterGroupsOutput, lastPage bool) bool {
		for _, parameterGroup := range parameterGroups.DBParameterGroups {
			resourceName := aws.StringValue(parameterGroup.DBParameterGroupName)
			if strings.Contains(resourceName, ".") {
				continue // skip default Default ParameterGroups like default.mysql5.6
			}
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_db_parameter_group",
				"aws",
				map[string]string{},
				RDSAllowEmptyValues,
				map[string]string{},
			))
		}
		return !lastPage
	})
}

func (g *RDSGenerator) loadDBSubnetGroups(svc *rds.RDS) error {
	return svc.DescribeDBSubnetGroupsPages(&rds.DescribeDBSubnetGroupsInput{}, func(subnets *rds.DescribeDBSubnetGroupsOutput, lastPage bool) bool {
		for _, subnet := range subnets.DBSubnetGroups {
			resourceName := aws.StringValue(subnet.DBSubnetGroupName)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_db_subnet_group",
				"aws",
				map[string]string{},
				RDSAllowEmptyValues,
				map[string]string{},
			))
		}
		return !lastPage
	})
}

func (g *RDSGenerator) loadOptionGroups(svc *rds.RDS) error {
	return svc.DescribeOptionGroupsPages(&rds.DescribeOptionGroupsInput{}, func(optionGroups *rds.DescribeOptionGroupsOutput, lastPage bool) bool {
		for _, optionGroup := range optionGroups.OptionGroupsList {
			resourceName := aws.StringValue(optionGroup.OptionGroupName)
			if strings.Contains(resourceName, ".") || strings.Contains(resourceName, ":") {
				continue // skip default Default OptionGroups like default.mysql5.6
			}
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_db_option_group",
				"aws",
				map[string]string{},
				RDSAllowEmptyValues,
				map[string]string{},
			))
		}
		return !lastPage
	})
}

func (g *RDSGenerator) loadEventSubscription(svc *rds.RDS) error {
	return svc.DescribeEventSubscriptionsPages(&rds.DescribeEventSubscriptionsInput{}, func(eventSubscriptions *rds.DescribeEventSubscriptionsOutput, lastPage bool) bool {
		for _, eventSubscription := range eventSubscriptions.EventSubscriptionsList {
			resourceName := aws.StringValue(eventSubscription.CustomerAwsId)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_db_event_subscription",
				"aws",
				map[string]string{},
				RDSAllowEmptyValues,
				map[string]string{},
			))
		}
		return !lastPage
	})
}

// Generate TerraformResources from AWS API,
// from each database create 1 TerraformResource.
// Need only database name as ID for terraform resource
// AWS api support paging
func (g *RDSGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
	svc := rds.New(sess)

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

	g.PopulateIgnoreKeys()
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
