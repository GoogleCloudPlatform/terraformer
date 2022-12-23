package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type LoadBalancerGenerator struct {
	IonosCloudService
}

func (g *LoadBalancerGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resourceType := "ionoscloud_loadbalancer"

	datacenters, err := helpers.GetAllDatacenters(*cloudApiClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		loadBalancerResponse, _, err := cloudApiClient.LoadBalancersApi.DatacentersLoadbalancersGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		if loadBalancerResponse.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing load balancers but received 'nil' instead, skipping search for datacenter with ID: %v",
				*datacenter.Id)
			continue
		}
		loadBalancers := *loadBalancerResponse.Items
		for _, loadBalancer := range loadBalancers {
			if loadBalancer.Properties == nil || loadBalancer.Properties.Name == nil {
				log.Printf(
					"[WARNING] 'nil' values in the response for load balancer with ID %v, datacenter ID: %v, skipping this resource",
					*loadBalancer.Id,
					*datacenter.Id)
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*loadBalancer.Id,
				*loadBalancer.Properties.Name+"-"+*loadBalancer.Id,
				resourceType,
				helpers.Ionos,
				map[string]string{helpers.DcId: *datacenter.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
