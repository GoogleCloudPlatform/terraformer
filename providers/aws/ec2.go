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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var ec2AllowEmptyValues = []string{"tags."}

type Ec2Generator struct {
	AWSService
}

func (g *Ec2Generator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.New(config)
	var filters []ec2.Filter
	for _, filter := range g.Filter {
		if strings.HasPrefix(filter.FieldPath, "tags.") && filter.IsApplicable("instance") {
			filters = append(filters, ec2.Filter{
				Name:   aws.String("tag:" + strings.TrimPrefix(filter.FieldPath, "tags.")),
				Values: filter.AcceptableValues,
			})
		}
	}
	p := ec2.NewDescribeInstancesPaginator(svc.DescribeInstancesRequest(&ec2.DescribeInstancesInput{
		Filters: filters,
	}))
	for p.Next(context.Background()) {
		for _, reservation := range p.CurrentPage().Reservations {
			for _, instance := range reservation.Instances {
				name := ""
				for _, tag := range instance.Tags {
					if strings.ToLower(aws.StringValue(tag.Key)) == "name" {
						name = aws.StringValue(tag.Value)
					}
				}
				attr, err := svc.DescribeInstanceAttributeRequest(&ec2.DescribeInstanceAttributeInput{
					Attribute:  ec2.InstanceAttributeNameUserData,
					InstanceId: instance.InstanceId,
				}).Send(context.Background())
				userDataBase64 := ""
				if err == nil && attr.UserData != nil && attr.UserData.Value != nil {
					userDataBase64 = aws.StringValue(attr.UserData.Value)
				}
				r := terraformutils.NewResource(
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
				g.Resources = append(g.Resources, r)
			}
		}
	}
	return p.Err()
}
