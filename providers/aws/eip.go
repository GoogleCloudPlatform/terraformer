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
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var eipAllowEmptyValues = []string{"tags."}

type ElasticIpGenerator struct {
	AWSService
}

func (g ElasticIpGenerator) createElasticIpsResources(svc *ec2.EC2) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	addresses, err := svc.DescribeAddresses(&ec2.DescribeAddressesInput{})

	if err != nil {
		log.Println(err)
		return resources
	}

	for _, eip := range addresses.Addresses {
		resources = append(resources, terraform_utils.NewSimpleResource(
			aws.StringValue(eip.AllocationId),
			aws.StringValue(eip.AllocationId),
			"aws_eip",
			"aws",
			eipAllowEmptyValues,
		))
	}

	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each elastic IPs
func (g *ElasticIpGenerator) InitResources() error {
	sess := g.generateSession()
	svc := ec2.New(sess)

	g.Resources = g.createElasticIpsResources(svc)
	return nil
}
