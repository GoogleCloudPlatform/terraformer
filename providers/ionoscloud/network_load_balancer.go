package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type NetworkLoadBalancerGenerator struct {
	Service
}

func (g *NetworkLoadBalancerGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	resourceType := "ionoscloud_networkloadbalancer"
	datacenters, err := helpers.GetAllDatacenters(*cloudAPIClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		networkLoadBalancerResponse, _, err := cloudAPIClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
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
				resourceType,
				helpers.Ionos,
				map[string]string{helpers.DcID: *datacenter.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
