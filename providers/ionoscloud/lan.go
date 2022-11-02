package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
)

type LanGenerator struct {
	IonosCloudService
}

func (g *LanGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	datacenters, err := helpers.GetAllDatacenters(*cloudApiClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		lans, _, err := cloudApiClient.LANsApi.DatacentersLansGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		if lans.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing LANs but received 'nil' instead, skipping search for datacenter with ID: %v.\n",
				*datacenter.Id)
			continue
		}
		for _, lan := range *lans.Items {
			if lan.Properties == nil || lan.Properties.Name == nil {
				log.Printf(
					"[WARNING] 'nil' values in the response for LAN with ID %v, datacenter ID: %v, skipping this resource.\n",
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
					"ionoscloud",
					map[string]string{helpers.DcId: *datacenter.Id},
					[]string{},
					map[string]interface{}{}))
			}
		}
	}
	return nil
}
