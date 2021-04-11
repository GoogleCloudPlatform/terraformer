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
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var SubnetAllowEmptyValues = []string{"tags."}

type SubnetGenerator struct {
	AWSService
}

func (SubnetGenerator) createResources(subnets *ec2.DescribeSubnetsOutput) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, subnet := range subnets.Subnets {
		resource := terraformutils.NewSimpleResource(
			StringValue(subnet.SubnetId),
			StringValue(subnet.SubnetId),
			"aws_subnet",
			"aws",
			SubnetAllowEmptyValues,
		)
		resource.IgnoreKeys = append(resource.IgnoreKeys, "availability_zone")
		resources = append(resources, resource)
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each subnet create 1 TerraformResource.
// Need SubnetId as ID for terraform resource
func (g *SubnetGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.NewFromConfig(config)
	p := ec2.NewDescribeSubnetsPaginator(svc, &ec2.DescribeSubnetsInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, g.createResources(page)...)
	}
	return nil
}
