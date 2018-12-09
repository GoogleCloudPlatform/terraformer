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

package elb

import (
	"waze/terraformer/aws_terraforming/aws_generator"
	"waze/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
)

var ignoreKey = map[string]bool{
	"^id$":                      true,
	"^arn":                      true,
	"^dns_name":                 true,
	"^source_security_group_id": true,
	"^zone_id":                  true,
	"^instances":                true, //dynamic value
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

type ElbGenerator struct {
	aws_generator.BasicGenerator
}

// Generate TerraformResources from AWS API,
// from each ELB create 1 TerraformResource.
// Need only ELB name as ID for terraform resource
// AWS api support paging
func (ElbGenerator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := elb.New(sess)
	resources := []terraform_utils.TerraformResource{}
	err := svc.DescribeLoadBalancersPages(&elb.DescribeLoadBalancersInput{}, func(loadBalancers *elb.DescribeLoadBalancersOutput, lastPage bool) bool {
		for _, loadBalancer := range loadBalancers.LoadBalancerDescriptions {
			resourceName := aws.StringValue(loadBalancer.LoadBalancerName)
			resources = append(resources, terraform_utils.NewTerraformResource(
				aws.StringValue(loadBalancer.LoadBalancerName),
				resourceName,
				"aws_elb",
				"aws",
				nil,
				map[string]string{}))
		}
		return !lastPage
	})
	if err != nil {
		return []terraform_utils.TerraformResource{}, map[string]terraform_utils.ResourceMetaData{}, err
	}
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, map[string]string{})
	return resources, metadata, nil

}
