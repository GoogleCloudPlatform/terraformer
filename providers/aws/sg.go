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
	"log"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const maxResults = 100

var SgAllowEmptyValues = []string{"tags."}

type SecurityGenerator struct {
	AWSService
}

func (SecurityGenerator) createResources(securityGroups []*ec2.SecurityGroup) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, sg := range securityGroups {
		if sg.VpcId == nil {
			continue
		}
		resources = append(resources, terraform_utils.NewResource(
			aws.StringValue(sg.GroupId),
			strings.Trim(aws.StringValue(sg.GroupName)+"_"+aws.StringValue(sg.GroupId), " "),
			"aws_security_group",
			"aws",
			map[string]string{},
			SgAllowEmptyValues,
			map[string]string{}))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each security group create 1 TerraformResource.
// Need GroupId as ID for terraform resource
// AWS support pagination with NextToken pattern
func (g *SecurityGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
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
	g.Resources = g.createResources(securityGroups)
	g.PopulateIgnoreKeys()
	return nil
}

// PostGenerateHook - replace sg-xxxxx string to terraform ID in all security group
func (g *SecurityGenerator) PostConvertHook() error {
	for j, resource := range g.Resources {
		for _, typeOfRule := range []string{"ingress", "egress"} {
			if _, exist := resource.Item[typeOfRule]; !exist {
				continue
			}
			for i, k := range resource.Item[typeOfRule].([]interface{}) {
				ingresses := k.(map[string]interface{})
				for key, ingress := range ingresses {
					if key != "security_groups" {
						continue
					}
					securityGroups := ingress.([]interface{})
					renamedSecurityGroups := []string{}
					for _, securityGroup := range securityGroups {
						found := false
						for _, i := range g.Resources {
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
					g.Resources[j].Item[typeOfRule].([]interface{})[i].(map[string]interface{})["security_groups"] = renamedSecurityGroups
				}
			}
		}
	}
	return nil
}
