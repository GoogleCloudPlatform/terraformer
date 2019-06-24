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
	"github.com/aws/aws-sdk-go/service/ec2"
)

var NaclAllowEmptyValues = []string{"tags."}

type NaclGenerator struct {
	AWSService
}

func (NaclGenerator) createResources(nacls *ec2.DescribeNetworkAclsOutput) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, nacl := range nacls.NetworkAcls {
		resources = append(resources, terraform_utils.NewResource(
			aws.StringValue(nacl.NetworkAclId),
			aws.StringValue(nacl.NetworkAclId),
			"aws_network_acl",
			"aws",
			nil,
			NaclAllowEmptyValues,
			map[string]string{}))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each network ACL create 1 TerraformResource.
// Need NetworkAclId as ID for terraform resource
func (g *NaclGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
	svc := ec2.New(sess)
	nacls, err := svc.DescribeNetworkAcls(&ec2.DescribeNetworkAclsInput{})
	if err != nil {
		return err

	}
	g.Resources = g.createResources(nacls)
	g.PopulateIgnoreKeys()
	return nil

}
