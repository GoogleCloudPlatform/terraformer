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
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

type SatelliteDataPlaneGenerator struct {
	IBMService
}

func (g SatelliteDataPlaneGenerator) loadVPCResources(vpcID, vpcName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		vpcID,
		vpcName,
		"ibm_is_vpc",
		"ibm",
		map[string]string{
			"address_prefix_management": "auto",
		},
		[]string{},
		map[string]interface{}{})

	return resource
}

func (g SatelliteDataPlaneGenerator) loadInstanceResources(instance vpcv1.Instance, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		*instance.ID,
		*instance.Name,
		"ibm_is_instance",
		"ibm",
		map[string]string{
			"vpc":                *instance.VPC.ID,
			"wait_before_delete": "true",
		},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
			"keys":       []string{},
		})

	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^port_speed$",
		"^primary_network_interface.[0-9].primary_ip.[0-9].address$",
		"^primary_network_interface.[0-9].primary_ip.[0-9].reserved_ip$",
	)

	return resource
}

func (g SatelliteDataPlaneGenerator) loadFloatingIPResources(floatingIPId, floatingIPName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		floatingIPId,
		floatingIPName,
		"ibm_is_floating_ip",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})

	// Conflicts with proxied attribute
	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^zone$",
	)

	return resource
}

func (g SatelliteDataPlaneGenerator) loadSecurityGroupResources(sgID, sgName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		sgID,
		sgName,
		"ibm_is_security_group",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})

	return resource
}

func (g SatelliteDataPlaneGenerator) loadSecurityGroupRuleResources(sgID, sgRuleID string, dependsOn []string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s.%s", sgID, sgRuleID),
		sgRuleID,
		"ibm_is_security_group_rule",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})

	return resources
}

func (g SatelliteDataPlaneGenerator) loadSubnetResources(subnetID, subnetName string, dependsOn []string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		subnetID,
		subnetName,
		"ibm_is_subnet",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})

	// Conflicts with proxied attribute
	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^total_ipv4_address_count$",
	)

	return resource
}

func contructEndpoint(subdomain, domain string) string {
	endpoint := fmt.Sprintf("https://%s.%s", subdomain, domain)
	return endpoint
}

func vpcClient(region string, sess *session.Session) (*vpcv1.VpcV1, error) {
	var cloudEndpoint = "cloud.ibm.com"

	bluemixToken := ""
	if strings.HasPrefix(sess.Config.IAMAccessToken, "Bearer") {
		bluemixToken = sess.Config.IAMAccessToken[7:len(sess.Config.IAMAccessToken)]
	} else {
		bluemixToken = sess.Config.IAMAccessToken
	}

	vpcurl := contructEndpoint(fmt.Sprintf("%s.iaas", region), fmt.Sprintf("%s/v1", cloudEndpoint))
	// if sess.Config.Visibility == "private" {
	// 	if region == "us-south" || region == "us-east" {
	// 		vpcurl = contructEndpoint(fmt.Sprintf("%s.private.iaas", region), fmt.Sprintf("%s/v1", cloudEndpoint))
	// 	} else {
	// 		return nil, fmt.Errorf("VPC supports private endpoints only in us-south and us-east")
	// 	}
	// }
	// if sess.Config.Visibility == "public-and-private" {
	// 	if region == "us-south" || region == "us-east" {
	// 		vpcurl = contructEndpoint(fmt.Sprintf("%s.private.iaas", region), fmt.Sprintf("%s/v1", cloudEndpoint))
	// 	}
	// 	vpcurl = contructEndpoint(fmt.Sprintf("%s.iaas", region), fmt.Sprintf("%s/v1", cloudEndpoint))
	// }

	vpcoptions := &vpcv1.VpcV1Options{
		URL: envFallBack([]string{"IBMCLOUD_IS_NG_API_ENDPOINT"}, vpcurl),
		Authenticator: &core.BearerTokenAuthenticator{
			BearerToken: bluemixToken,
		},
	}
	vpcclient, err := vpcv1.NewVpcV1(vpcoptions)
	if err != nil {
		return nil, fmt.Errorf("Error occured while configuring vpc service: %v ", err)
	}

	return vpcclient, nil
}

func (g *SatelliteDataPlaneGenerator) InitResources() error {
	vpcName := g.Args["vpc"].(string)
	if len(vpcName) == 0 {
		return fmt.Errorf("required VPC name missing, '-vpc=<vpcName>' flag not set")
	}

	region := g.Args["region"].(string)
	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
	}

	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}

	err = authenticateAPIKey(sess)
	if err != nil {
		return err
	}

	// VPC
	vpcObj, err := vpcClient(region, sess)
	if err != nil {
		log.Println("Error building VPC object: ", err)
		return err
	}

	start := ""
	allVPCrecs := []vpcv1.VPC{}
	for {
		listVpcsOptions := &vpcv1.ListVpcsOptions{}
		if start != "" {
			listVpcsOptions.Start = &start
		}
		vpcs, response, err := vpcObj.ListVpcs(listVpcsOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching vpcs %s\n%s", err, response)
		}

		start = GetNext(vpcs.Next)
		allVPCrecs = append(allVPCrecs, vpcs.Vpcs...)
		if start == "" {
			break
		}
	}

	// VPC & Instances
	for _, vpc := range allVPCrecs {
		if *vpc.Name == vpcName {

			var vpcDependsOn []string
			vpcDependsOn = append(vpcDependsOn,
				"ibm_is_vpc."+terraformutils.TfSanitize(*vpc.Name))

			g.Resources = append(g.Resources, g.loadVPCResources(*vpc.ID, *vpc.Name))

			start = ""
			var allrecs []vpcv1.Instance
			for {
				options := &vpcv1.ListInstancesOptions{}
				if start != "" {
					options.Start = &start
				}

				instances, response, err := vpcObj.ListInstances(options)
				if err != nil {
					return fmt.Errorf("Error Fetching Instances %s\n%s", err, response)
				}
				start = GetNext(instances.Next)
				allrecs = append(allrecs, instances.Instances...)
				if start == "" {
					break
				}
			}

			// Floating IP
			start := ""
			allFloatingIPs := []vpcv1.FloatingIP{}
			for {
				floatingIPOptions := &vpcv1.ListFloatingIpsOptions{}
				if start != "" {
					floatingIPOptions.Start = &start
				}
				floatingIPs, response, err := vpcObj.ListFloatingIps(floatingIPOptions)
				if err != nil {
					return fmt.Errorf("Error Fetching floating IPs %s\n%s", err, response)
				}
				start = GetNext(floatingIPs.Next)
				allFloatingIPs = append(allFloatingIPs, floatingIPs.FloatingIps...)
				if start == "" {
					break
				}
			}

			for _, instance := range allrecs {
				g.Resources = append(g.Resources, g.loadInstanceResources(instance, vpcDependsOn))

				for _, ip := range allFloatingIPs {
					target, _ := ip.Target.(*vpcv1.FloatingIPTarget)
					if *target.ID == *instance.PrimaryNetworkInterface.ID {
						g.Resources = append(g.Resources, g.loadFloatingIPResources(*ip.ID, *ip.Name))
					}
				}
			}

			// Security group
			start = ""
			var allSgRecs []vpcv1.SecurityGroup
			for {
				options := &vpcv1.ListSecurityGroupsOptions{
					VPCID: vpc.ID,
				}
				if start != "" {
					options.Start = &start
				}

				sgs, response, err := vpcObj.ListSecurityGroups(options)
				if err != nil {
					return fmt.Errorf("Error Fetching security Groups %s\n%s", err, response)
				}
				start = GetNext(sgs.Next)
				allSgRecs = append(allSgRecs, sgs.SecurityGroups...)
				if start == "" {
					break
				}
			}

			for _, group := range allSgRecs {
				var sgDependsOn []string
				sgDependsOn = append(sgDependsOn,
					"ibm_is_security_group."+terraformutils.TfSanitize(*group.Name))
				g.Resources = append(g.Resources, g.loadSecurityGroupResources(*group.ID, *group.Name))
				listSecurityGroupRulesOptions := &vpcv1.ListSecurityGroupRulesOptions{
					SecurityGroupID: group.ID,
				}
				rules, response, err := vpcObj.ListSecurityGroupRules(listSecurityGroupRulesOptions)
				if err != nil {
					return fmt.Errorf("Error Fetching security group rules %s\n%s", err, response)
				}
				for _, sgrule := range rules.Rules {
					switch reflect.TypeOf(sgrule).String() {
					case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp":
						{
							rule := sgrule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolIcmp)
							g.Resources = append(g.Resources, g.loadSecurityGroupRuleResources(*group.ID, *rule.ID, sgDependsOn))
						}

					case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll":
						{
							rule := sgrule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll)
							g.Resources = append(g.Resources, g.loadSecurityGroupRuleResources(*group.ID, *rule.ID, sgDependsOn))
						}

					case "*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp":
						{
							rule := sgrule.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp)
							g.Resources = append(g.Resources, g.loadSecurityGroupRuleResources(*group.ID, *rule.ID, sgDependsOn))
						}
					}
				}
			}

			// Subnet
			start = ""
			var allSubNetRecs []vpcv1.Subnet
			for {
				options := &vpcv1.ListSubnetsOptions{}
				if start != "" {
					options.Start = &start
				}

				subnets, response, err := vpcObj.ListSubnets(options)
				if err != nil {
					return fmt.Errorf("Error Fetching subnets %s\n%s", err, response)
				}
				start = GetNext(subnets.Next)
				allSubNetRecs = append(allSubNetRecs, subnets.Subnets...)
				if start == "" {
					break
				}
			}

			for _, subnet := range allSubNetRecs {
				if *subnet.VPC.Name == vpcName {
					g.Resources = append(g.Resources, g.loadSubnetResources(*subnet.ID, *subnet.Name, vpcDependsOn))
				}
			}
		}
	}
	return nil
}
