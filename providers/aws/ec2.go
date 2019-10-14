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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var ec2AllowEmptyValues = []string{"tags."}

type Ec2Generator struct {
	AWSService
}

func (g *Ec2Generator) InitResources() error {
	sess := g.generateSession()
	svc := ec2.New(sess)
	var filters []*ec2.Filter
	for _, filter := range g.Filter {
		if strings.HasPrefix(filter.FieldPath, "tags.") && filter.ResourceName == "aws_instance" {
			filters = append(filters, &ec2.Filter{
				Name:   aws.String("tag:" + strings.TrimPrefix(filter.FieldPath, "tags.")),
				Values: aws.StringSlice(filter.AcceptableValues),
			})
		}
	}
	input := ec2.DescribeInstancesInput{
		Filters: filters,
	}
	err := svc.DescribeInstancesPages(&input, func(instances *ec2.DescribeInstancesOutput, lastPage bool) bool {
		for _, reservation := range instances.Reservations {
			for _, instance := range reservation.Instances {
				name := ""
				for _, tag := range instance.Tags {
					if strings.ToLower(aws.StringValue(tag.Key)) == "name" {
						name = aws.StringValue(tag.Value)
					}
				}
				attr, err := svc.DescribeInstanceAttribute(&ec2.DescribeInstanceAttributeInput{
					Attribute:  aws.String(ec2.InstanceAttributeNameUserData),
					InstanceId: instance.InstanceId,
				})
				userDataBase64 := ""
				if err == nil && attr.UserData != nil && attr.UserData.Value != nil {
					userDataBase64 = aws.StringValue(attr.UserData.Value)
				}
				r := terraform_utils.NewResource(
					aws.StringValue(instance.InstanceId),
					aws.StringValue(instance.InstanceId)+"_"+name,
					"aws_instance",
					"aws",
					map[string]string{
						"user_data_base64":  userDataBase64,
						"source_dest_check": "true",
					},
					ec2AllowEmptyValues,
					map[string]interface{}{},
				)
				r.IgnoreKeys = append(r.IgnoreKeys, "^ebs_block_device.(.*)")
				g.Resources = append(g.Resources, r)
			}

		}
		return true
	})
	if err != nil {
		return err
	}
	return nil

}
