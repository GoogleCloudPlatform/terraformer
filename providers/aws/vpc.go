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

var VpcAllowEmptyValues = []string{"tags."}

type VpcGenerator struct {
	AWSService
}

func (VpcGenerator) createResources(vpcs *ec2.DescribeVpcsOutput) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, vpc := range vpcs.Vpcs {
		resources = append(resources, terraform_utils.NewResource(
			aws.StringValue(vpc.VpcId),
			aws.StringValue(vpc.VpcId),
			"aws_vpc",
			"aws",
			map[string]string{},
			VpcAllowEmptyValues,
			map[string]string{},
		))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each vpc create 1 TerraformResource.
// Need VpcId as ID for terraform resource
func (g *VpcGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
	svc := ec2.New(sess)
	vpcs, err := svc.DescribeVpcs(&ec2.DescribeVpcsInput{})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(vpcs)
	g.PopulateIgnoreKeys()
	return nil
}
