package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
)

type NATGatewayRuleGenerator struct {
	IonosCloudService
}

func (g *NATGatewayRuleGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resourceType := "ionoscloud_natgateway_rule"

	datacenters, err := helpers.GetAllDatacenters(*cloudApiClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		natGatewaysResponse, _, err := cloudApiClient.NATGatewaysApi.DatacentersNatgatewaysGet(context.TODO(), *datacenter.Id).Execute()
		if err != nil {
			return err
		}
		if natGatewaysResponse.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing NAT gateways but received 'nil' instead, skipping search for datacenter with ID: %v",
				*datacenter.Id)
			continue
		}
		natGateways := *natGatewaysResponse.Items
		for _, natGateway := range natGateways {
			rulesResponse, _, err := cloudApiClient.NATGatewaysApi.DatacentersNatgatewaysRulesGet(context.TODO(), *datacenter.Id, *natGateway.Id).Depth(1).Execute()
			if err != nil {
				return err
			}
			if rulesResponse.Items == nil {
				log.Printf(
					"[WARNING] expected a response containing NAT gateway rules but received 'nil' instead, skipping search for NAT Gateway with ID: %v, datacenter ID: %v.",
					*natGateway.Id,
					*datacenter.Id)
				continue
			}
			rules := *rulesResponse.Items
			for _, rule := range rules {
				if rule.Properties == nil || rule.Properties.Name == nil {
					log.Printf(
						"[WARNING] 'nil' values in the response for NAT gateway rule with ID: %v, NAT gateway ID: %v, datacenter ID: %v",
						*rule.Id,
						*natGateway.Id,
						*datacenter.Id)
					continue
				}
				g.Resources = append(g.Resources, terraformutils.NewResource(
					*rule.Id,
					*rule.Properties.Name+"-"+*rule.Id,
					resourceType,
					helpers.Ionos,
					map[string]string{helpers.DcId: *datacenter.Id, "natgateway_id": *natGateway.Id},
					[]string{},
					map[string]interface{}{},
				))
			}
		}
	}
	return nil
}
