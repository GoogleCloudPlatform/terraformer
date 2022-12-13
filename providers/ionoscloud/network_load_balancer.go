package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
)

type NetworkLoadBalancerGenerator struct {
	IonosCloudService
}

func (g *NetworkLoadBalancerGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resource_type := "ionoscloud_networkloadbalancer"
	datacenters, err := helpers.GetAllDatacenters(*cloudApiClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		networkLoadBalancerResponse, _, err := cloudApiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		if networkLoadBalancerResponse.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing network load balancers but received 'nil' instead, skipping search for datacenter with ID: %v.\n",
				*datacenter.Id)
			continue
		}
		networkLoadBalancers := *networkLoadBalancerResponse.Items
		for _, networkLoadBalancer := range networkLoadBalancers {
			if networkLoadBalancer.Properties == nil || networkLoadBalancer.Properties.Name == nil {
				log.Printf(
					"[WARNING] 'nil' values in the response for network load balancer with ID %v, datacenter ID: %v, skipping this resource.\n",
					*networkLoadBalancer.Id,
					*datacenter.Id,
				)
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*networkLoadBalancer.Id,
				*networkLoadBalancer.Properties.Name+"-"+*networkLoadBalancer.Id,
				resource_type,
				helpers.Ionos,
				map[string]string{helpers.DcId: *datacenter.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}