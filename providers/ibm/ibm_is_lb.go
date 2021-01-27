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

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// LBGenerator ...
type LBGenerator struct {
	IBMService
}

func (g LBGenerator) createLBResources(lbID, lbName string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		lbID,
		lbName,
		"ibm_is_lb",
		"ibm",
		[]string{})
	return resources
}

func (g LBGenerator) createLBPoolResources(lbID, lbPoolID, lbPoolName string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s/%s", lbID, lbPoolID),
		lbPoolName,
		"ibm_is_lb_pool",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g LBGenerator) createLBPoolMemberResources(lbID, lbPoolID, lbPoolMemberID, lbPoolMemberName string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s", lbID, lbPoolID, lbPoolMemberID),
		lbPoolMemberName,
		"ibm_is_lb_pool_member",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g LBGenerator) createLBListenerResources(lbID, lbListenerID, lbListenerName string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s/%s", lbID, lbListenerID),
		lbListenerName,
		"ibm_is_lb_listener",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g LBGenerator) createLBListenerPolicyResources(lbID, lbListenerID, lbListenerPolicyID, lbListenerPolicyName string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s", lbID, lbListenerID, lbListenerPolicyID),
		lbListenerPolicyName,
		"ibm_is_lb_listener_policy",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

func (g LBGenerator) createLBListenerPolicyRuleResources(lbID, lbListenerID, lbListenerPolicyID, lbListenerPolicyRuleID, lbListenerPolicyName string, dependsOn []string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s/%s", lbID, lbListenerID, lbListenerPolicyID, lbListenerPolicyRuleID),
		lbListenerPolicyName,
		"ibm_is_lb_listener_policy_rule",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{
			"depends_on": dependsOn,
		})
	return resources
}

// InitResources ...
func (g *LBGenerator) InitResources() error {
	region := envFallBack([]string{"IC_REGION"}, "us-south")
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("No API key set")
	}

	rg := g.Args["resource_group"]
	if rg != nil {
		_ = rg.(string)
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
	allrecs := []vpcv1.LoadBalancer{}

	listLoadBalancersOptions := &vpcv1.ListLoadBalancersOptions{}
	lbs, response, err := vpcclient.ListLoadBalancers(listLoadBalancersOptions)
	if err != nil {
		return fmt.Errorf("Error Fetching vpcs %s\n%s", err, response)
	}
	allrecs = append(allrecs, lbs.LoadBalancers...)

	for _, lb := range allrecs {
		var dependsOn []string
		dependsOn = append(dependsOn,
			"ibm_is_lb."+terraformutils.TfSanitize(*lb.Name))
		g.Resources = append(g.Resources, g.createLBResources(*lb.ID, *lb.Name))
		listLoadBalancerPoolsOptions := &vpcv1.ListLoadBalancerPoolsOptions{
			LoadBalancerID: lb.ID,
		}
		lbPools, response, err := vpcclient.ListLoadBalancerPools(listLoadBalancerPoolsOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching Load Balancer Pools %s\n%s", err, response)
		}
		for _, lbPool := range lbPools.Pools {
			g.Resources = append(g.Resources, g.createLBPoolResources(*lb.ID, *lbPool.ID, *lbPool.Name, dependsOn))
			var dependsOn1 []string
			dependsOn1 = append(dependsOn,
				"ibm_is_lb_pool."+terraformutils.TfSanitize(*lbPool.Name))
			listLoadBalancerPoolMembersOptions := &vpcv1.ListLoadBalancerPoolMembersOptions{
				LoadBalancerID: lb.ID,
				PoolID:         lbPool.ID,
			}
			lbPoolMembers, response, err := vpcclient.ListLoadBalancerPoolMembers(listLoadBalancerPoolMembersOptions)
			if err != nil {
				return fmt.Errorf("Error Fetching Load Balancer Pool Members %s\n%s", err, response)
			}
			for _, lbPoolMember := range lbPoolMembers.Members {
				g.Resources = append(g.Resources, g.createLBPoolMemberResources(*lb.ID, *lbPool.ID, *lbPoolMember.ID, *lbPool.Name, dependsOn1))
			}
		}

		listLoadBalancerListenersOptions := &vpcv1.ListLoadBalancerListenersOptions{
			LoadBalancerID: lb.ID,
		}
		lbListeners, response, err := vpcclient.ListLoadBalancerListeners(listLoadBalancerListenersOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching Load Balancer Listeners %s\n%s", err, response)
		}
		for _, lbListener := range lbListeners.Listeners {
			g.Resources = append(g.Resources, g.createLBListenerResources(*lb.ID, *lbListener.ID, *lbListener.ID, dependsOn))
			var dependsOn2 []string
			dependsOn2 = append(dependsOn,
				"ibm_is_lb_listener."+terraformutils.TfSanitize(*lbListener.ID))
			listLoadBalancerListenerPoliciesOptions := &vpcv1.ListLoadBalancerListenerPoliciesOptions{
				LoadBalancerID: lb.ID,
				ListenerID:     lbListener.ID,
			}
			lbListenerPolicies, response, err := vpcclient.ListLoadBalancerListenerPolicies(listLoadBalancerListenerPoliciesOptions)
			if err != nil {
				return fmt.Errorf("Error Fetching Load Balancer Listener Policies %s\n%s", err, response)
			}
			for _, lbListenerPolicy := range lbListenerPolicies.Policies {
				g.Resources = append(g.Resources, g.createLBListenerPolicyResources(*lb.ID, *lbListener.ID, *lbListenerPolicy.ID, *lbListenerPolicy.Name, dependsOn2))
				dependsOn2 = append(dependsOn2,
					"ibm_is_lb_listener_policy."+terraformutils.TfSanitize(*lbListenerPolicy.Name))
				listLoadBalancerListenerPolicyRulesOptions := &vpcv1.ListLoadBalancerListenerPolicyRulesOptions{
					LoadBalancerID: lb.ID,
					ListenerID:     lbListener.ID,
					PolicyID:       lbListenerPolicy.ID,
				}
				lbListenerPolicyRules, response, err := vpcclient.ListLoadBalancerListenerPolicyRules(listLoadBalancerListenerPolicyRulesOptions)
				if err != nil {
					return fmt.Errorf("Error Fetching Load Balancer Listener Policy Rules %s\n%s", err, response)
				}
				for _, lbListenerPolicyRule := range lbListenerPolicyRules.Rules {
					g.Resources = append(g.Resources, g.createLBListenerPolicyRuleResources(*lb.ID, *lbListener.ID, *lbListenerPolicy.ID, *lbListenerPolicyRule.ID, *lbListenerPolicyRule.ID, dependsOn2))

				}
			}
		}
	}
	return nil
}
