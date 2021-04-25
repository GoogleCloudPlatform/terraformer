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
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
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
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := elasticloadbalancing.NewFromConfig(config)
	p := elasticloadbalancing.NewDescribeLoadBalancersPaginator(svc, &elasticloadbalancing.DescribeLoadBalancersInput{})
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, loadBalancer := range page.LoadBalancerDescriptions {
			resourceName := StringValue(loadBalancer.LoadBalancerName)
			resource := terraformutils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_elb",
				"aws",
				ElbAllowEmptyValues,
			)
			resource.IgnoreKeys = append(resource.IgnoreKeys, "^instances\\.(.*)") // don't import current connect instances to ELB
			g.Resources = append(g.Resources, resource)
		}
	}
	return nil
}
