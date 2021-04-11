// Copyright 2021 The Terraformer Authors.
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
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type OpsworksGenerator struct {
	AWSService
}

func (g *OpsworksGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := opsworks.NewFromConfig(config)

	e = g.fetchStacks(svc)
	if e != nil {
		return e
	}

	e = g.fetchUserProfile(svc)
	if e != nil {
		return e
	}

	return nil
}

func (g *OpsworksGenerator) fetchApps(stackID *string, svc *opsworks.Client) error {
	apps, err := svc.DescribeApps(context.TODO(), &opsworks.DescribeAppsInput{
		StackId: stackID,
	})
	if err != nil {
		return err
	}
	for _, app := range apps.Apps {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			StringValue(app.AppId),
			StringValue(app.AppId),
			"aws_opsworks_application",
			"aws",
			[]string{"tags."},
		))
	}
	return nil
}

func (g *OpsworksGenerator) fetchLayers(stackID *string, svc *opsworks.Client) error {
	apps, err := svc.DescribeLayers(context.TODO(), &opsworks.DescribeLayersInput{
		StackId: stackID,
	})
	if err != nil {
		return err
	}
	for _, layer := range apps.Layers {
		switch layer.Type {
		case types.LayerTypeCustom:
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				StringValue(layer.LayerId),
				StringValue(layer.LayerId),
				"aws_opsworks_custom_layer",
				"aws",
				[]string{"tags."},
			))
		case types.LayerTypePhpApp:
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				StringValue(layer.LayerId),
				StringValue(layer.LayerId),
				"aws_opsworks_php_app_layer",
				"aws",
				[]string{"tags."},
			))
		case types.LayerTypeJavaApp:
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				StringValue(layer.LayerId),
				StringValue(layer.LayerId),
				"aws_opsworks_java_app_layer",
				"aws",
				[]string{"tags."},
			))
		case types.LayerTypeWeb:
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				StringValue(layer.LayerId),
				StringValue(layer.LayerId),
				"aws_opsworks_static_web_layer",
				"aws",
				[]string{"tags."},
			))
		}
	}
	return nil
}

func (g *OpsworksGenerator) fetchInstances(stackID *string, svc *opsworks.Client) error {
	apps, err := svc.DescribeInstances(context.TODO(), &opsworks.DescribeInstancesInput{
		StackId: stackID,
	})
	if err != nil {
		return err
	}
	for _, instances := range apps.Instances {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			StringValue(instances.InstanceId),
			StringValue(instances.InstanceId),
			"aws_opsworks_instance",
			"aws",
			[]string{"tags."},
		))
	}
	return nil
}
func (g *OpsworksGenerator) fetchRdsInstances(stackID *string, svc *opsworks.Client) error {
	apps, err := svc.DescribeRdsDbInstances(context.TODO(), &opsworks.DescribeRdsDbInstancesInput{
		StackId: stackID,
	})
	if err != nil {
		return err
	}
	for _, rdsDbInstance := range apps.RdsDbInstances {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			StringValue(rdsDbInstance.RdsDbInstanceArn),
			StringValue(rdsDbInstance.RdsDbInstanceArn),
			"aws_opsworks_instance",
			"aws",
			map[string]string{
				"rds_db_instance_arn": StringValue(rdsDbInstance.RdsDbInstanceArn),
				"stack_id":            StringValue(stackID),
			},
			[]string{"tags."},
			map[string]interface{}{},
		))
	}
	return nil
}

func (g *OpsworksGenerator) fetchStacks(svc *opsworks.Client) error {
	apps, err := svc.DescribeStacks(context.TODO(), &opsworks.DescribeStacksInput{})
	if err != nil {
		return err
	}
	for _, stack := range apps.Stacks {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			StringValue(stack.StackId),
			StringValue(stack.StackId),
			"aws_opsworks_stack",
			"aws",
			[]string{"tags."},
		))

		e := g.fetchApps(stack.StackId, svc)
		if e != nil {
			log.Println(err)
		}

		e = g.fetchInstances(stack.StackId, svc)
		if e != nil {
			log.Println(err)
		}

		e = g.fetchRdsInstances(stack.StackId, svc)
		if e != nil {
			log.Println(err)
		}

		e = g.fetchLayers(stack.StackId, svc)
		if e != nil {
			log.Println(err)
		}
	}
	return nil
}

func (g *OpsworksGenerator) fetchUserProfile(svc *opsworks.Client) error {
	apps, err := svc.DescribeUserProfiles(context.TODO(), &opsworks.DescribeUserProfilesInput{})
	if err != nil {
		return err
	}
	for _, userProfile := range apps.UserProfiles {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			StringValue(userProfile.IamUserArn),
			StringValue(userProfile.IamUserArn),
			"aws_opsworks_user_profile",
			"aws",
			[]string{"tags."},
		))
	}
	return nil
}
