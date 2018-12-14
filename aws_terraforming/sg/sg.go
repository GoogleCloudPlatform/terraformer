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

package sg

import (
	"log"
	"strings"

	"waze/terraformer/aws_terraforming/aws_generator"
	"waze/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const maxResults = 1000

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

type SecurityGenerator struct {
	aws_generator.BasicGenerator
}

func (SecurityGenerator) createResources(securityGroups []*ec2.SecurityGroup) []terraform_utils.TerraformResource {
	resources := []terraform_utils.TerraformResource{}
	for _, sg := range securityGroups {
		if sg.VpcId == nil {
			continue
		}
		resources = append(resources, terraform_utils.NewTerraformResource(
			aws.StringValue(sg.GroupId),
			strings.Trim(aws.StringValue(sg.GroupName), " "),
			"aws_security_group",
			"aws",
			nil,
			map[string]string{}))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each security group create 1 TerraformResource.
// Need GroupId as ID for terraform resource
// AWS support pagination with NextToken patter
func (g SecurityGenerator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := ec2.New(sess)
	var securityGroups []*ec2.SecurityGroup
	var err error
	firstRun := true
	securityGroupsOutput := &ec2.DescribeSecurityGroupsOutput{}
	for {
		if firstRun || securityGroupsOutput.NextToken != nil {
			firstRun = false
			securityGroupsOutput, err = svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
				MaxResults: aws.Int64(maxResults),
				NextToken:  securityGroupsOutput.NextToken,
			})
			securityGroups = append(securityGroups, securityGroupsOutput.SecurityGroups...)
			if err != nil {
				log.Println(err)
			}
		} else {
			break
		}
	}
	resources := g.createResources(securityGroups)
	metadata := terraform_utils.NewResourcesMetaData(resources, g.IgnoreKeys(resources, "aws"), allowEmptyValues, map[string]string{})
	return resources, metadata, nil
}

// PostGenerateHook - replace sg-xxxxx string to terraform ID in all security group
func (g SecurityGenerator) PostGenerateHook(resources []terraform_utils.TerraformResource) ([]terraform_utils.TerraformResource, error) {
	for _, resource := range resources {
		for _, typeOfRule := range []string{"ingress", "egress"} {
			item := resource.Item.(map[string]interface{})
			if _, exist := item[typeOfRule]; !exist {
				continue
			}
			for i, k := range item[typeOfRule].([]interface{}) {
				ingresses := k.(map[string]interface{})
				for key, ingress := range ingresses {
					if key != "security_groups" {
						continue
					}
					securityGroups := ingress.([]interface{})
					renamedSecurityGroups := []string{}
					for _, securityGroup := range securityGroups {
						found := false
						for _, i := range resources {
							if i.InstanceState.ID == securityGroup {
								renamedSecurityGroups = append(renamedSecurityGroups, "${"+i.InstanceInfo.Type+"."+i.ResourceName+".id}")
								found = true
								break
							}
						}
						if !found {
							renamedSecurityGroups = append(renamedSecurityGroups, securityGroup.(string))
						}
					}
					item[typeOfRule].([]interface{})[i].(map[string]interface{})["security_groups"] = renamedSecurityGroups
				}
			}
		}
	}
	return resources, nil
}
