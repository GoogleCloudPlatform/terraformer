package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type VolumeGenerator struct {
	IonosCloudService
}

func (g *VolumeGenerator) InitResources() error {
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
					volumes, _, err := cloudApiClient.ServersApi.DatacentersServersVolumesGet(context.TODO(), *datacenter.Id, *server.Id).Depth(2).Execute()
					if err != nil {
						return err
					}
					for _, volume := range *volumes.Items {
						//bootVolume will be included in the server
						if volume.Properties != nil && volume.Properties.Name != nil && (server.Properties.BootVolume != nil && *server.Properties.BootVolume.Id != *volume.Id) {
							g.Resources = append(g.Resources, terraformutils.NewResource(
								*volume.Id,
								*volume.Properties.Name+"-"+*volume.Id,
								"ionoscloud_volume",
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
