package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type NicGenerator struct {
	IonosCloudService
}

func (g *NicGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	datacenters, err := helpers.GetAllDatacenters(*cloudApiClient)
	if err != nil {
		return err
	}
	if datacenters != nil {
		for _, datacenter := range datacenters {
			servers, _, err := cloudApiClient.ServersApi.DatacentersServersGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
			if err != nil {
				return err
			}
			if servers.Items != nil {
				for _, server := range *servers.Items {
					nics, _, err := cloudApiClient.NetworkInterfacesApi.DatacentersServersNicsGet(context.TODO(), *datacenter.Id, *server.Id).Depth(2).Execute()
					if err != nil {
						return err
					}
					for _, nic := range *nics.Items {
						if nic.Properties != nil && nic.Properties.Name != nil && *nic.Properties.Lan != 1 {
							g.Resources = append(g.Resources, terraformutils.NewResource(
								*nic.Id,
								*nic.Properties.Name+"-"+*nic.Id,
								"ionoscloud_nic",
								helpers.Ionos,
								map[string]string{helpers.DcId: *datacenter.Id,
									"server_id": *server.Id},
								[]string{},
								map[string]interface{}{}))
						}
					}
				}
			}
		}
	}
	return nil
}
