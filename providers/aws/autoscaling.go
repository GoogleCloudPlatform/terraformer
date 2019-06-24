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
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var AsgAllowEmptyValues = []string{"tags."}

type AutoScalingGenerator struct {
	AWSService
}

func (g *AutoScalingGenerator) loadAutoScalingGroups(svc *autoscaling.AutoScaling) error {
	err := svc.DescribeAutoScalingGroupsPages(&autoscaling.DescribeAutoScalingGroupsInput{}, func(asgs *autoscaling.DescribeAutoScalingGroupsOutput, lastPage bool) bool {
		for _, asg := range asgs.AutoScalingGroups {
			resourceName := aws.StringValue(asg.AutoScalingGroupName)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
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
				map[string]string{},
			))
		}
		return !lastPage
	})
	return err
}

func (g *AutoScalingGenerator) loadLaunchConfigurations(svc *autoscaling.AutoScaling) error {
	err := svc.DescribeLaunchConfigurationsPages(&autoscaling.DescribeLaunchConfigurationsInput{}, func(lcs *autoscaling.DescribeLaunchConfigurationsOutput, lastPage bool) bool {
		for _, lc := range lcs.LaunchConfigurations {
			resourceName := aws.StringValue(lc.LaunchConfigurationName)
			attributes := map[string]string{}
			// only for LaunchConfigurations with userdata, we want get user_data_base64
			if aws.StringValue(lc.UserData) != "" {
				attributes["user_data_base64"] = "=" //need set not empty string to get user_data_base64 from provider
			}
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_launch_configuration",
				"aws",
				attributes,
				AsgAllowEmptyValues,
				map[string]string{},
			))
		}
		return !lastPage
	})
	return err
}

func (g *AutoScalingGenerator) loadLaunchTemplates(sess *session.Session) error {
	ec2svc := ec2.New(sess)
	firstRun := true
	var err error
	launchTemplatesOutput := &ec2.DescribeLaunchTemplatesOutput{}
	for {
		if firstRun || launchTemplatesOutput.NextToken != nil {
			firstRun = false
			launchTemplatesOutput, err = ec2svc.DescribeLaunchTemplates(&ec2.DescribeLaunchTemplatesInput{
				MaxResults: aws.Int64(maxResults),
				NextToken:  launchTemplatesOutput.NextToken,
			})
			for _, lt := range launchTemplatesOutput.LaunchTemplates {
				g.Resources = append(g.Resources, terraform_utils.NewResource(
					aws.StringValue(lt.LaunchTemplateId),
					aws.StringValue(lt.LaunchTemplateName),
					"aws_launch_template",
					"aws",
					map[string]string{},
					AsgAllowEmptyValues,
					map[string]string{},
				))
			}
		} else {
			break
		}
	}
	return err
}

// Generate TerraformResources from AWS API,
// from each ASG create 1 TerraformResource.
// Need only ASG name as ID for terraform resource
// AWS api support paging
func (g *AutoScalingGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
	svc := autoscaling.New(sess)
	if err := g.loadAutoScalingGroups(svc); err != nil {
		return err
	}
	if err := g.loadLaunchConfigurations(svc); err != nil {
		return err
	}
	if err := g.loadLaunchTemplates(sess); err != nil {
		return err
	}
	g.PopulateIgnoreKeys()
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
		templateFiles := []terraform_utils.Resource{}
		for i, r := range g.Resources {
			if r.InstanceInfo.Type != "aws_launch_configuration" {
				continue
			}
			if userDataBase64, exist := r.InstanceState.Attributes["user_data_base64"]; exist {
				userData, err := base64.StdEncoding.DecodeString(userDataBase64)
				if err != nil {
					continue
				}
				fileName := "userdata-" + r.ResourceName + ".txt"
				err = ioutil.WriteFile(fileName, userData, os.ModePerm) // TODO write files in tf file path
				if err != nil {
					continue
				}
				userDataFile := terraform_utils.NewResource(
					r.ResourceName+"_userdata",
					r.ResourceName+"_userdata",
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
				g.Resources[i].Item["user_data"] = "${template_file." + userDataFile.ResourceName + ".rendered}"
				templateFiles = append(templateFiles, userDataFile)
			}
		}
		g.Resources = append(g.Resources, templateFiles...)
	*/
	return nil
}
