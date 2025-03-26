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
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
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
	svc := ec2.NewFromConfig(config)
	var filters []types.Filter
	for _, filter := range g.Filter {
		if strings.HasPrefix(filter.FieldPath, "tags.") && filter.IsApplicable("instance") {
			filters = append(filters, types.Filter{
				Name:   aws.String("tag:" + strings.TrimPrefix(filter.FieldPath, "tags.")),
				Values: filter.AcceptableValues,
			})
		}
	}
	p := ec2.NewDescribeInstancesPaginator(svc, &ec2.DescribeInstancesInput{
		Filters: filters,
	})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, reservation := range page.Reservations {
			for _, instance := range reservation.Instances {
				name := ""
				for _, tag := range instance.Tags {
					if strings.ToLower(*tag.Key) == "name" {
						name = *tag.Value
					}
				}
				attr, err := svc.DescribeInstanceAttribute(context.TODO(), &ec2.DescribeInstanceAttributeInput{
					Attribute:  types.InstanceAttributeNameUserData,
					InstanceId: instance.InstanceId,
				})
				userDataBase64 := ""
				if err == nil && attr.UserData != nil && attr.UserData.Value != nil {
					userDataBase64 = *attr.UserData.Value
				}
				r := terraformutils.NewResource(
					*instance.InstanceId,
					*instance.InstanceId+"_"+name,
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
	return nil
}

func (g *Ec2Generator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_instance" {
			continue
		}
		if r.Item["root_block_device"] == nil {
			continue
		}

		rootDeviceVolumeType := r.InstanceState.Attributes["root_block_device.0.volume_type"]
		if !(rootDeviceVolumeType == "io1" || rootDeviceVolumeType == "io2" || rootDeviceVolumeType == "gp3") {
			delete(r.Item["root_block_device"].([]interface{})[0].(map[string]interface{}), "iops")
		}
		if rootDeviceVolumeType != "gp3" {
			delete(r.Item["root_block_device"].([]interface{})[0].(map[string]interface{}), "throughput")
		}
	}

	return nil
}
