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

package vpn_connection

import (
	"waze/terraformer/aws_terraforming/aws_generator"
	"waze/terraformer/terraform_utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var ignoreKey = map[string]bool{
	"^tunnel2_vgw_inside_address": true,
	"^id$":                        true,
	"^tunnel2_cgw_inside_address": true,
	"^tunnel2_bgp_holdtime":       true,
	"^tunnel2_bgp_asn":            true,
	"^tunnel2_address":            true,
	"^tunnel1_vgw_inside_address": true,
	"^tunnel1_cgw_inside_address": true,
	"^tunnel1_bgp_holdtime":       true,
	"^tunnel1_bgp_asn":            true,
	"^tunnel1_address":            true,
}

var allowEmptyValues = map[string]bool{
	"tags.": true,
}

type VpnConnectionGenerator struct {
	aws_generator.BasicGenerator
}

func (VpnConnectionGenerator) createResources(vpncs *ec2.DescribeVpnConnectionsOutput) []terraform_utils.TerraformResource {
	resoures := []terraform_utils.TerraformResource{}
	for _, vpnc := range vpncs.VpnConnections {
		resourceName := ""
		if len(vpnc.Tags) > 0 {
			for _, tag := range vpnc.Tags {
				if aws.StringValue(tag.Key) == "Name" {
					resourceName = aws.StringValue(tag.Value)
					break
				}
			}
		}
		resoures = append(resoures, terraform_utils.TerraformResource{
			ResourceType: "aws_vpn_connection",
			ResourceName: resourceName,
			Item:         nil,
			ID:           aws.StringValue(vpnc.VpnConnectionId),
			Provider:     "aws",
		})
	}
	return resoures
}

// Generate TerraformResources from AWS API,
// from each vpn connection create 1 TerraformResource.
// Need VpnConnectionId as ID for terraform resource
func (g VpnConnectionGenerator) Generate(region string) ([]terraform_utils.TerraformResource, map[string]terraform_utils.ResourceMetaData, error) {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(region)})
	svc := ec2.New(sess)
	vpncs, err := svc.DescribeVpnConnections(&ec2.DescribeVpnConnectionsInput{})
	if err != nil {
		return []terraform_utils.TerraformResource{}, map[string]terraform_utils.ResourceMetaData{}, err
	}
	resources := g.createResources(vpncs)
	metadata := terraform_utils.NewResourcesMetaData(resources, ignoreKey, allowEmptyValues, map[string]string{})
	return resources, metadata, nil

}
