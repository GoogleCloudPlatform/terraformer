package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
)

type NATGatewayGenerator struct {
	IonosCloudService
}

func (g *NATGatewayGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resource_type := "ionoscloud_natgateway"
	datacenters, err := helpers.GetAllDatacenters(*cloudApiClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		natGatewayResponse, _, err := cloudApiClient.NATGatewaysApi.DatacentersNatgatewaysGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		if natGatewayResponse.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing NAT gateways but received 'nil' instead, skipping search for datacenter with ID: %v.\n",
				*datacenter.Id)
			continue
		}
		natGateways := *natGatewayResponse.Items
		for _, natGateway := range natGateways {
			if natGateway.Properties == nil || natGateway.Properties.Name == nil {
				log.Printf(
					"[WARNING] 'nil' values in the response for NAT gateway with ID %v, datacenter ID: %v, skipping this resource.\n",
					*natGateway.Id,
					*datacenter.Id,
				)
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*natGateway.Id,
				*natGateway.Properties.Name+"-"+*natGateway.Id,
				resource_type,
				helpers.Ionos,
				map[string]string{helpers.DcId: *datacenter.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
