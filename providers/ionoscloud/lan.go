package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type LanGenerator struct {
	Service
}

func (g *LanGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	datacenters, err := helpers.GetAllDatacenters(*cloudAPIClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		lans, _, err := cloudAPIClient.LANsApi.DatacentersLansGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		if lans.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing LANs but received 'nil' instead, skipping search for datacenter with ID: %v",
				*datacenter.Id)
			continue
		}
		for _, lan := range *lans.Items {
			if lan.Properties == nil || lan.Properties.Name == nil {
				log.Printf(
					"[WARNING] 'nil' values in the response for LAN with ID %v, datacenter ID: %v, skipping this resource",
					*lan.Id,
					*datacenter.Id,
				)
				continue
			}
			if lan.Properties != nil && lan.Properties.Name != nil {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					*lan.Id,
					*lan.Properties.Name+"-"+*lan.Id,
					"ionoscloud_lan",
					helpers.Ionos,
					map[string]string{helpers.DcID: *datacenter.Id},
					[]string{},
					map[string]interface{}{}))
			}
		}
	}
	return nil
}
