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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var AsgAllowEmptyValues = []string{"tags."}

type AutoScalingGenerator struct {
	AWSService
}

func (g *AutoScalingGenerator) loadAutoScalingGroups(svc *autoscaling.Client) error {
	p := autoscaling.NewDescribeAutoScalingGroupsPaginator(svc.DescribeAutoScalingGroupsRequest(&autoscaling.DescribeAutoScalingGroupsInput{}))
	for p.Next(context.Background()) {
		for _, asg := range p.CurrentPage().AutoScalingGroups {
			resourceName := aws.StringValue(asg.AutoScalingGroupName)
			g.Resources = append(g.Resources, terraformutils.NewResource(
				resourceName,
				resourceName,
				"aws_autoscaling_group",
				"aws",
				map[string]string{
					"force_delete":              "false",
					"metrics_granularity":       "1Minute",
					"wait_for_capacity_timeout": "10m",
				},
				AsgAllowEmptyValues,
				map[string]interface{}{},
			))
		}
	}
	return p.Err()
}

func (g *AutoScalingGenerator) loadLaunchConfigurations(svc *autoscaling.Client) error {
	p := autoscaling.NewDescribeLaunchConfigurationsPaginator(svc.DescribeLaunchConfigurationsRequest(&autoscaling.DescribeLaunchConfigurationsInput{}))
	for p.Next(context.Background()) {
		for _, lc := range p.CurrentPage().LaunchConfigurations {
			resourceName := aws.StringValue(lc.LaunchConfigurationName)
			attributes := map[string]string{}
			// only for LaunchConfigurations with userdata, we want get user_data_base64
			if aws.StringValue(lc.UserData) != "" {
				attributes["user_data_base64"] = "=" // need set not empty string to get user_data_base64 from provider
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				resourceName,
				resourceName,
				"aws_launch_configuration",
				"aws",
				attributes,
				AsgAllowEmptyValues,
				map[string]interface{}{},
			))
		}
	}
	return p.Err()
}

func (g *AutoScalingGenerator) loadLaunchTemplates(config aws.Config) error {
	ec2svc := ec2.New(config)

	p := ec2.NewDescribeLaunchTemplatesPaginator(ec2svc.DescribeLaunchTemplatesRequest(&ec2.DescribeLaunchTemplatesInput{}))
	for p.Next(context.Background()) {
		for _, lt := range p.CurrentPage().LaunchTemplates {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				aws.StringValue(lt.LaunchTemplateId),
				aws.StringValue(lt.LaunchTemplateName),
				"aws_launch_template",
				"aws",
				AsgAllowEmptyValues,
			))
		}
	}
	return p.Err()
}

// Generate TerraformResources from AWS API,
// from each ASG create 1 TerraformResource.
// Need only ASG name as ID for terraform resource
// AWS api support paging
func (g *AutoScalingGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := autoscaling.New(config)
	if err := g.loadAutoScalingGroups(svc); err != nil {
		return err
	}
	if err := g.loadLaunchConfigurations(svc); err != nil {
		return err
	}
	if err := g.loadLaunchTemplates(config); err != nil {
		return err
	}
	return nil
}

func (g *AutoScalingGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_autoscaling_group" {
			continue
		}
		if lcName, exist := r.InstanceState.Attributes["launch_configuration"]; exist {
			for _, lc := range g.Resources {
				if lc.InstanceInfo.Type != "aws_launch_configuration" {
					continue
				}
				if lcName == lc.InstanceState.Attributes["name"] {
					g.Resources[i].Item["launch_configuration"] = "${aws_launch_configuration." + lc.ResourceName + ".name}"
					continue
				}
			}
		}
		// TODO add LaunchTemplate and mix policy connection naming
	}
	// TODO fix tfVar value
	/*
		templateFiles := []terraformutils.Resource{}
		for i, r := range g.Resources {
			if r.InstanceInfo.Type != "aws_launch_configuration" {
				continue
			}
			if userDataBase64, exist := r.InstanceState.Attributes["user_data_base64"]; exist {
				userData, err := base64.StdEncoding.DecodeString(userDataBase64)
				if err != nil {
					continue
				}
				fileName := "userdata-" + r.ServiceName + ".txt"
				err = ioutil.WriteFile(fileName, userData, os.ModePerm) // TODO write files in tf file path
				if err != nil {
					continue
				}
				userDataFile := terraformutils.NewResource(
					r.ServiceName+"_userdata",
					r.ServiceName+"_userdata",
					"template_file",
					"",
					map[string]string{},
					[]string{},
					map[string]string{},
				)
				tfVar := strings.Replace(fmt.Sprintf("${base64decode(file(\"%s\"))}", fileName), "\\\"", "\"", -1)
				userDataFile.Item = map[string]interface{}{
					"template": tfVar,
				}

				delete(g.Resources[i].Item, "user_data_base64")
				g.Resources[i].Item["user_data"] = "${template_file." + userDataFile.ServiceName + ".rendered}"
				templateFiles = append(templateFiles, userDataFile)
			}
		}
		g.Resources = append(g.Resources, templateFiles...)
	*/
	return nil
}
