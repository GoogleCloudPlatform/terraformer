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

package vpc

import (
	"waze/terraformer/aws_terraforming/aws_generator"
	"waze/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

type VpcGenerator struct {
	aws_generator.BasicGenerator
}

func (VpcGenerator) createResources(vpcs *ec2.DescribeVpcsOutput) []terraform_utils.TerraformResource {
	resoures := []terraform_utils.TerraformResource{}
	for _, vpc := range vpcs.Vpcs {
		resourceName := ""
		if len(vpc.Tags) > 0 {
			for _, tag := range vpc.Tags {
				if aws.StringValue(tag.Key) == "Name" {
					resourceName = aws.StringValue(tag.Value)
					break
				}
			}
		}
		resoures = append(resoures, terraform_utils.NewTerraformResource(aws.StringValue(vpc.VpcId), resourceName, "aws_vpc", "aws", nil, map[string]string{}))
	}
	return resoures
}

// Generate TerraformResources from AWS API,
// from each vpc create 1 TerraformResource.
// Need VpcId as ID for terraform resource
func (g VpcGenerator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := ec2.New(sess)
	vpcs, err := svc.DescribeVpcs(&ec2.DescribeVpcsInput{})
	if err != nil {
		return []terraform_utils.TerraformResource{}, map[string]terraform_utils.ResourceMetaData{}, err
	}
	resources := g.createResources(vpcs)
	metadata := terraform_utils.NewResourcesMetaData(resources, g.IgnoreKeys(resources), allowEmptyValues, map[string]string{})
	return resources, metadata, nil
}
