package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
)

type ApplicationLoadBalancerGenerator struct {
	IonosCloudService
}

func (g *ApplicationLoadBalancerGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resource_type := "ionoscloud_application_loadbalancer"
	datacenters, err := helpers.GetAllDatacenters(*cloudApiClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		applicationLoadBalancerResponse, _, err := cloudApiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		if applicationLoadBalancerResponse.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing application load balancers but received 'nil' instead, skipping search for datacenter with ID: %v.\n",
				*datacenter.Id)
			continue
		}
		applicationLoadBalancers := *applicationLoadBalancerResponse.Items
		for _, applicationLoadBalancer := range applicationLoadBalancers {
			if applicationLoadBalancer.Properties == nil || applicationLoadBalancer.Properties.Name == nil {
				log.Printf(
					"[WARNING] 'nil' values in the response for application load balancer with ID %v, datacenter ID: %v, skipping this resource.\n",
					*applicationLoadBalancer.Id,
					*datacenter.Id,
				)
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*applicationLoadBalancer.Id,
				*applicationLoadBalancer.Properties.Name+"-"+*applicationLoadBalancer.Id,
				resource_type,
				helpers.Ionos,
				map[string]string{helpers.DcId: *datacenter.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
