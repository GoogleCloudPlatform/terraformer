// Copyright 2019 The Terraformer Authors.
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

package ibm

import (
	"fmt"
	"os"
	"reflect"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// SecurityGroupGenerator ...
type SecurityGroupGenerator struct {
	IBMService
}

func (g SecurityGroupGenerator) createSecurityGroupResources(sgID, sgName string) terraformutils.Resource {
	resources := terraformutils.NewSimpleResource(
		sgID,
		normalizeResourceName(sgName, false),
		"ibm_is_security_group",
		"ibm",
		[]string{})
	return resources
}

func (g SecurityGroupGenerator) createSecurityGroupRuleResources(sgID, sgRuleID string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s.%s", sgID, sgRuleID),
		normalizeResourceName("ibm_is_security_group_rule", true),
		"ibm_is_security_group_rule",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

// InitResources ...
func (g *SecurityGroupGenerator) InitResources() error {
	region := g.Args["region"].(string)
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("No API key set")
	}

	vpcurl := fmt.Sprintf("https://%s.iaas.cloud.ibm.com/v1", region)
	vpcoptions := &vpcv1.VpcV1Options{
		URL: envFallBack([]string{"IBMCLOUD_IS_API_ENDPOINT"}, vpcurl),
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	}
	vpcclient, err := vpcv1.NewVpcV1(vpcoptions)
	if err != nil {
		return err
	}
	start := ""
	var allrecs []vpcv1.SecurityGroup
	for {
		options := &vpcv1.ListSecurityGroupsOptions{}
		if start != "" {
			options.Start = &start
		}
		if rg := g.Args["resource_group"].(string); rg != "" {
			rg, err = GetResourceGroupID(apiKey, rg, region)
			if err != nil {
				return fmt.Errorf("Error Fetching Resource Group Id %s", err)
			}
			options.ResourceGroupID = &rg
		}
		sgs, response, err := vpcclient.ListSecurityGroups(options)
		if err != nil {
			return fmt.Errorf("Error Fetching security Groups %s\n%s", err, response)
		}
		start = GetNext(sgs.Next)
		allrecs = append(allrecs, sgs.SecurityGroups...)
		if start == "" {
			break
		}
	}

	for _, group := range allrecs {
		var dependsOn []string

		g.Resources = append(g.Resources, g.createSecurityGroupResources(*group.ID, *group.Name))
		sgResourceName := g.Resources[len(g.Resources)-1:][0].ResourceName
		dependsOn = append(dependsOn,
			"ibm_is_security_group."+sgResourceName)
		listSecurityGroupRulesOptions := &vpcv1.ListSecurityGroupRulesOptions{
			SecurityGroupID: group.ID,
		}
		rules, response, err := vpcclient.ListSecurityGroupRules(listSecurityGroupRulesOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching security group rules %s\n%s", err, response)
		}
		for _, sgrule := range rules.Rules {
			switch reflect.TypeOf(sgrule).String() {
			case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp":
				{
					rule := sgrule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp)
					g.Resources = append(g.Resources, g.createSecurityGroupRuleResources(*group.ID, *rule.ID, dependsOn))
				}

			case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll":
				{
					rule := sgrule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll)
					g.Resources = append(g.Resources, g.createSecurityGroupRuleResources(*group.ID, *rule.ID, dependsOn))
				}

			case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp":
				{
					rule := sgrule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp)
					g.Resources = append(g.Resources, g.createSecurityGroupRuleResources(*group.ID, *rule.ID, dependsOn))
				}
			}
		}
	}
	return nil
}
