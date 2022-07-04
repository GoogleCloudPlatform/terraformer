package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
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
	if datacenters != nil {
		for _, datacenter := range datacenters {
			lans, _, err := cloudApiClient.LANsApi.DatacentersLansGet(context.TODO(), *datacenter.Id).Depth(10).Execute()
			if err != nil {
				return err
			}
			if lans.Items != nil {
				for _, lan := range *lans.Items {
					g.Resources = append(g.Resources, terraformutils.NewResource(
						*lan.Id,
						*lan.Properties.Name+"-"+*lan.Id,
						"ionoscloud_lan",
						"ionoscloud",
						map[string]string{},
						[]string{},
						map[string]interface{}{}))
				}
			}
		}
	}
	return nil
}
