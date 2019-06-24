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

var VpnConnectionAllowEmptyValues = []string{"tags."}

type VpnConnectionGenerator struct {
	AWSService
}

func (VpnConnectionGenerator) createResources(vpncs *ec2.DescribeVpnConnectionsOutput) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	for _, vpnc := range vpncs.VpnConnections {
		resources = append(resources, terraform_utils.NewResource(
			aws.StringValue(vpnc.VpnConnectionId),
			aws.StringValue(vpnc.VpnConnectionId),
			"aws_vpn_connection",
			"aws",
			map[string]string{},
			VpnConnectionAllowEmptyValues,
			map[string]string{},
		))
	}
	return resources
}

// Generate TerraformResources from AWS API,
// from each vpn connection create 1 TerraformResource.
// Need VpnConnectionId as ID for terraform resource
func (g *VpnConnectionGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
	svc := ec2.New(sess)
	vpncs, err := svc.DescribeVpnConnections(&ec2.DescribeVpnConnectionsInput{})
	if err != nil {
		return err
	}
	g.Resources = g.createResources(vpncs)
	g.PopulateIgnoreKeys()
	return nil

}
