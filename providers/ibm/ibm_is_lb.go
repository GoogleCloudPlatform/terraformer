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
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// LBGenerator ...
type LBGenerator struct {
	IBMService
}

func (g LBGenerator) createLBResources(lbID, lbName string) terraformutils.Resource {
	resource := terraformutils.NewResource(
		lbID,
		normalizeResourceName(lbName, true),
		"ibm_is_lb",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})

	// Deprecated parameters
	resource.IgnoreKeys = append(resource.IgnoreKeys,
		"^profile$",
	)
	return resource
}

func (g LBGenerator) createLBPoolResources(lbID, lbPoolID, lbPoolName string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", lbID, lbPoolID),
		normalizeResourceName(lbPoolName, true),
		"ibm_is_lb_pool",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})
	return resources
}

func (g LBGenerator) createLBPoolMemberResources(lbID, lbPoolID, lbPoolMemberID, lbPoolMemberName string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s", lbID, lbPoolID, lbPoolMemberID),
		normalizeResourceName(lbPoolMemberName, true),
		"ibm_is_lb_pool_member",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})
	return resources
}

func (g LBGenerator) createLBListenerResources(lbID, lbListenerID, lbListenerName string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s", lbID, lbListenerID),
		normalizeResourceName(lbListenerName, true),
		"ibm_is_lb_listener",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})
	return resources
}

func (g LBGenerator) createLBListenerPolicyResources(lbID, lbListenerID, lbListenerPolicyID, lbListenerPolicyName string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s", lbID, lbListenerID, lbListenerPolicyID),
		normalizeResourceName(lbListenerPolicyName, true),
		"ibm_is_lb_listener_policy",
		"ibm",
		map[string]string{
			"target_http_status_code": "302",
		},
		[]string{},
		map[string]interface{}{})
	return resources
}

func (g LBGenerator) createLBListenerPolicyRuleResources(lbID, lbListenerID, lbListenerPolicyID, lbListenerPolicyRuleID, lbListenerPolicyName string) terraformutils.Resource {
	resources := terraformutils.NewResource(
		fmt.Sprintf("%s/%s/%s/%s", lbID, lbListenerID, lbListenerPolicyID, lbListenerPolicyRuleID),
		normalizeResourceName(lbListenerPolicyName, true),
		"ibm_is_lb_listener_policy_rule",
		"ibm",
		map[string]string{},
		[]string{},
		map[string]interface{}{})
	return resources
}

// InitResources ...
func (g *LBGenerator) InitResources() error {
	region := g.Args["region"].(string)
	apiKey := os.Getenv("IC_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("no API key set")
	}

	rg := g.Args["resource_group"]
	if rg != nil {
		_ = rg.(string)
	}

	isURL := GetVPCEndPoint(region)
	iamURL := GetAuthEndPoint()
	vpcoptions := &vpcv1.VpcV1Options{
		URL: isURL,
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
			URL:    iamURL,
		},
	}
	vpcclient, err := vpcv1.NewVpcV1(vpcoptions)
	if err != nil {
		return err
	}
	var allrecs []vpcv1.LoadBalancer

	listLoadBalancersOptions := &vpcv1.ListLoadBalancersOptions{}
	lbs, response, err := vpcclient.ListLoadBalancers(listLoadBalancersOptions)
	if err != nil {
		return fmt.Errorf("Error Fetching vpcs %s\n%s", err, response)
	}
	allrecs = append(allrecs, lbs.LoadBalancers...)

	for _, lb := range allrecs {
		g.Resources = append(g.Resources, g.createLBResources(*lb.ID, *lb.Name))

		listLoadBalancerPoolsOptions := &vpcv1.ListLoadBalancerPoolsOptions{
			LoadBalancerID: lb.ID,
		}
		lbPools, response, err := vpcclient.ListLoadBalancerPools(listLoadBalancerPoolsOptions)
		if err != nil {
			return fmt.Errorf("Error Fetching Load Balancer Pools %s\n%s", err, response)
		}
		for _, lbPool := range lbPools.Pools {
			g.Resources = append(g.Resources, g.createLBPoolResources(*lb.ID, *lbPool.ID, *lbPool.Name))
			listLoadBalancerPoolMembersOptions := &vpcv1.ListLoadBalancerPoolMembersOptions{
				LoadBalancerID: lb.ID,
				PoolID:         lbPool.ID,
			}
			lbPoolMembers, response, err := vpcclient.ListLoadBalancerPoolMembers(listLoadBalancerPoolMembersOptions)
			if err != nil {
				return fmt.Errorf("Error Fetching Load Balancer Pool Members %s\n%s", err, response)
			}
			for _, lbPoolMember := range lbPoolMembers.Members {
				g.Resources = append(g.Resources, g.createLBPoolMemberResources(*lb.ID, *lbPool.ID, *lbPoolMember.ID, *lbPool.Name))
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
			g.Resources = append(g.Resources, g.createLBListenerResources(*lb.ID, *lbListener.ID, *lbListener.ID))
			listLoadBalancerListenerPoliciesOptions := &vpcv1.ListLoadBalancerListenerPoliciesOptions{
				LoadBalancerID: lb.ID,
				ListenerID:     lbListener.ID,
			}
			lbListenerPolicies, response, err := vpcclient.ListLoadBalancerListenerPolicies(listLoadBalancerListenerPoliciesOptions)
			if err != nil {
				return fmt.Errorf("Error Fetching Load Balancer Listener Policies %s\n%s", err, response)
			}
			for _, lbListenerPolicy := range lbListenerPolicies.Policies {
				g.Resources = append(g.Resources, g.createLBListenerPolicyResources(*lb.ID, *lbListener.ID, *lbListenerPolicy.ID, *lbListenerPolicy.Name))
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
					g.Resources = append(g.Resources, g.createLBListenerPolicyRuleResources(*lb.ID, *lbListener.ID, *lbListenerPolicy.ID, *lbListenerPolicyRule.ID, *lbListenerPolicyRule.ID))

				}
			}
		}
	}
	return nil
}

func (g *LBGenerator) PostConvertHook() error {
	for _, r := range g.Resources {
		if r.InstanceInfo.Type != "ibm_is_lb" {
			continue
		}

		for i, pool := range g.Resources {
			if pool.InstanceInfo.Type != "ibm_is_lb_pool" {
				continue
			}
			if pool.InstanceState.Attributes["lb"] == r.InstanceState.Attributes["id"] {
				g.Resources[i].Item["lb"] = "${ibm_is_lb." + r.ResourceName + ".id}"
			}
			for i, poolMember := range g.Resources {
				if poolMember.InstanceInfo.Type != "ibm_is_lb_pool_member" {
					continue
				}

				poolID := strings.Split(pool.InstanceState.Attributes["id"], "/")[1]
				if poolMember.InstanceState.Attributes["pool"] == poolID {
					g.Resources[i].Item["pool"] = "${ibm_is_lb_pool." + pool.ResourceName + ".id}"
				}
			}
		}

		for i, poolMember := range g.Resources {
			if poolMember.InstanceInfo.Type != "ibm_is_lb_pool_member" {
				continue
			}
			if poolMember.InstanceState.Attributes["lb"] == r.InstanceState.Attributes["id"] {
				g.Resources[i].Item["lb"] = "${ibm_is_lb." + r.ResourceName + ".id}"
			}
		}

		for i, listener := range g.Resources {
			if listener.InstanceInfo.Type != "ibm_is_lb_listener" {
				continue
			}
			if listener.InstanceState.Attributes["lb"] == r.InstanceState.Attributes["id"] {
				g.Resources[i].Item["lb"] = "${ibm_is_lb." + r.ResourceName + ".id}"
			}
		}

		for i, listenerPolicy := range g.Resources {
			if listenerPolicy.InstanceInfo.Type != "ibm_is_lb_listener_policy" {
				continue
			}
			if listenerPolicy.InstanceState.Attributes["lb"] == r.InstanceState.Attributes["id"] {
				g.Resources[i].Item["lb"] = "${ibm_is_lb." + r.ResourceName + ".id}"
			}
			for i, listenerPolicyRule := range g.Resources {
				if listenerPolicyRule.InstanceInfo.Type != "ibm_is_lb_listener_policy_rule" {
					continue
				}

				if listenerPolicyRule.InstanceState.Attributes["listener"] == listenerPolicy.InstanceState.Attributes["id"] {
					g.Resources[i].Item["listener"] = "${ibm_is_lb_listener_policy." + listenerPolicy.ResourceName + ".id}"
				}
			}
		}

		for i, listenerPolicyRule := range g.Resources {
			if listenerPolicyRule.InstanceInfo.Type != "ibm_is_lb_listener_policy_rule" {
				continue
			}
			if listenerPolicyRule.InstanceState.Attributes["lb"] == r.InstanceState.Attributes["id"] {
				g.Resources[i].Item["lb"] = "${ibm_is_lb." + r.ResourceName + ".id}"
			}
		}

	}

	return nil
}
