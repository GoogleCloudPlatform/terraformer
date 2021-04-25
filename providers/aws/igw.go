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

var IgwAllowEmptyValues = []string{"tags."}

type IgwGenerator struct {
	AWSService
}

// Generate TerraformResources from AWS API,
// from each Internet gateway create 1 TerraformResource.
// Need InternetGatewayId as ID for terraform resource
func (g *IgwGenerator) createResources(igws *ec2.DescribeInternetGatewaysOutput) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, internetGateway := range igws.InternetGateways {
		if len(internetGateway.Attachments) == 0 {
			continue
		}
		resources = append(resources, terraformutils.NewSimpleResource(
			StringValue(internetGateway.InternetGatewayId),
			StringValue(internetGateway.InternetGatewayId),
			"aws_internet_gateway",
			"aws",
			IgwAllowEmptyValues,
		))
	}
	return resources
}

func (g *IgwGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.NewFromConfig(config)
	p := ec2.NewDescribeInternetGatewaysPaginator(svc, &ec2.DescribeInternetGatewaysInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, g.createResources(page)...)
	}
	return nil
}
