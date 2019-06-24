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

var IgwAllowEmptyValues = []string{"tags."}

type IgwGenerator struct {
	AWSService
}

// Generate TerraformResources from AWS API,
// from each Internet gateway create 1 TerraformResource.
// Need InternetGatewayId as ID for terraform resource
func (g IgwGenerator) createResources(igws *ec2.DescribeInternetGatewaysOutput) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, internetGateway := range igws.InternetGateways {
		resources = append(resources, terraform_utils.NewResource(
			aws.StringValue(internetGateway.InternetGatewayId),
			aws.StringValue(internetGateway.InternetGatewayId),
			"aws_internet_gateway",
			"aws",
			map[string]string{},
			IgwAllowEmptyValues,
			map[string]string{},
		))
	}
	return resources
}

func (g *IgwGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
	svc := ec2.New(sess)
	igws, err := svc.DescribeInternetGateways(&ec2.DescribeInternetGatewaysInput{})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(igws)
	g.PopulateIgnoreKeys()
	return nil

}
