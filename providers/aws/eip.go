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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var eipAllowEmptyValues = []string{"tags."}

type ElasticIPGenerator struct {
	AWSService
}

func (g *ElasticIPGenerator) createElasticIpsResources(svc *ec2.Client) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	addresses, err := svc.DescribeAddresses(context.TODO(), &ec2.DescribeAddressesInput{})

	if err != nil {
		log.Println(err)
		return resources
	}

	for _, eip := range addresses.Addresses {
		resources = append(resources, terraformutils.NewSimpleResource(
			StringValue(eip.AllocationId),
			StringValue(eip.AllocationId),
			"aws_eip",
			"aws",
			eipAllowEmptyValues,
		))
	}

	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each elastic IPs
func (g *ElasticIPGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.NewFromConfig(config)

	g.Resources = g.createElasticIpsResources(svc)
	return nil
}
