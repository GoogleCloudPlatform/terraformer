package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type ALBForwardingRuleGenerator struct {
	Service
}

func (g *ALBForwardingRuleGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	resourceType := "ionoscloud_application_loadbalancer_forwardingrule"
	datacenters, err := helpers.GetAllDatacenters(*cloudAPIClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		applicationLoadBalancerResponse, _, err := cloudAPIClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		if applicationLoadBalancerResponse.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing application load balancers but received 'nil' instead, skipping search for datacenter with ID: %v",
				*datacenter.Id)
			continue
		}
		applicationLoadBalancers := *applicationLoadBalancerResponse.Items
		for _, applicationLoadBalancer := range applicationLoadBalancers {
			if applicationLoadBalancer.Properties == nil || applicationLoadBalancer.Properties.Name == nil {
				log.Printf(
					"[WARNING] 'nil' values in the response for application load balancer with ID %v, datacenter ID: %v, skipping this resource",
					*applicationLoadBalancer.Id,
					*datacenter.Id,
				)
				continue
			}
			albForwardingRulesResponse, _, err := cloudAPIClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesGet(context.TODO(), *datacenter.Id, *applicationLoadBalancer.Id).Depth(1).Execute()
			if err != nil {
				return err
			}
			if albForwardingRulesResponse.Items == nil {
				log.Printf("[WARNING] expected a response containing ALB forwarding rules but received 'nil' instead, skipping search for ALB with ID: %v, datacenter ID: %v", *applicationLoadBalancer.Id, *datacenter.Id)
				continue
			}
			albForwardingRules := *albForwardingRulesResponse.Items
			for _, albForwardingRule := range albForwardingRules {
				if albForwardingRule.Properties == nil || albForwardingRule.Properties.Name == nil {
					log.Printf("[WARNING] 'nil' values in the response for ALB forwarding rule with ID: %v, ALB ID: %v, datacenter ID: %v, skipping this resource", *albForwardingRule.Id, *applicationLoadBalancer.Id, *datacenter.Id)
					continue
				}
				g.Resources = append(g.Resources, terraformutils.NewResource(
					*albForwardingRule.Id,
					*albForwardingRule.Properties.Name+"-"+*albForwardingRule.Id,
					resourceType,
					helpers.Ionos,
					map[string]string{"application_loadbalancer_id": *applicationLoadBalancer.Id, helpers.DcID: *datacenter.Id},
					[]string{},
					map[string]interface{}{}))
			}
		}
	}
	return nil
}
