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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
)

var ElbAllowEmptyValues = []string{"tags."}

type ElbGenerator struct {
	AWSService
}

// Generate TerraformResources from AWS API,
// from each ELB create 1 TerraformResource.
// Need only ELB name as ID for terraform resource
// AWS api support paging
func (g *ElbGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
	svc := elb.New(sess)

	err := svc.DescribeLoadBalancersPages(&elb.DescribeLoadBalancersInput{}, func(loadBalancers *elb.DescribeLoadBalancersOutput, lastPage bool) bool {
		for _, loadBalancer := range loadBalancers.LoadBalancerDescriptions {
			resourceName := aws.StringValue(loadBalancer.LoadBalancerName)
			resource := terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_elb",
				"aws",
				map[string]string{},
				ElbAllowEmptyValues,
				map[string]string{},
			)
			resource.IgnoreKeys = append(resource.IgnoreKeys, "^instances\\.(.*)") // don't import current connect instances to ELB
			g.Resources = append(g.Resources, resource)
		}
		return !lastPage
	})
	if err != nil {
		return err
	}
	g.PopulateIgnoreKeys()
	return nil

}
