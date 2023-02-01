package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type NATGatewayGenerator struct {
	Service
}

func (g *NATGatewayGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	resourceType := "ionoscloud_natgateway"
	datacenters, err := helpers.GetAllDatacenters(*cloudAPIClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		natGatewayResponse, _, err := cloudAPIClient.NATGatewaysApi.DatacentersNatgatewaysGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
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
				resourceType,
				helpers.Ionos,
				map[string]string{helpers.DcID: *datacenter.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
