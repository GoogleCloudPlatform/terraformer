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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var NaclAllowEmptyValues = []string{"tags."}

type NaclGenerator struct {
	AWSService
}

func (NaclGenerator) createResources(nacls *ec2.DescribeNetworkAclsOutput) []terraformutils.Resource {
	resources := []terraformutils.Resource{}
	for _, nacl := range nacls.NetworkAcls {
		resources = append(resources, terraformutils.NewSimpleResource(
			aws.StringValue(nacl.NetworkAclId),
			aws.StringValue(nacl.NetworkAclId),
			"aws_network_acl",
			"aws",
			NaclAllowEmptyValues))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each network ACL create 1 TerraformResource.
// Need NetworkAclId as ID for terraform resource
func (g *NaclGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.New(config)
	p := ec2.NewDescribeNetworkAclsPaginator(svc.DescribeNetworkAclsRequest(&ec2.DescribeNetworkAclsInput{}))
	for p.Next(context.Background()) {
		g.Resources = append(g.Resources, g.createResources(p.CurrentPage())...)
	}
	return p.Err()
}
