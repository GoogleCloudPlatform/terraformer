package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type VolumeGenerator struct {
	Service
}

func (g *VolumeGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	datacenters, err := helpers.GetAllDatacenters(*cloudAPIClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		servers, _, err := cloudAPIClient.ServersApi.DatacentersServersGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		if servers.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing servers but received 'nil' instead, skipping search for datacenter with ID: %v.\n",
				*datacenter.Id)
			continue
		}
		for _, server := range *servers.Items {
			volumes, _, err := cloudAPIClient.ServersApi.DatacentersServersVolumesGet(context.TODO(), *datacenter.Id, *server.Id).Depth(1).Execute()
			if err != nil {
				return err
			}
			if volumes.Items == nil {
				log.Printf(
					"[WARNING] expected a response containing volumes but received 'nil' instead, skipping search for server with ID: %v, datacenter ID: %v.\n",
					*server.Id,
					*datacenter.Id)
				continue
			}
			for _, volume := range *volumes.Items {
				if volume.Properties == nil || volume.Properties.Name == nil {
					log.Printf(
						"[WARNING] 'nil' values in the response for volume with ID %v, server ID: %v, datacenter ID: %v, skipping this resource.\n",
						*volume.Id,
						*server.Id,
						*datacenter.Id,
					)
					continue
				}
				// bootVolume will be included in the server
				if server.Properties.BootVolume != nil && *server.Properties.BootVolume.Id != *volume.Id {
					g.Resources = append(g.Resources, terraformutils.NewResource(
						*volume.Id,
						*volume.Properties.Name+"-"+*volume.Id,
						"ionoscloud_volume",
						helpers.Ionos,
						map[string]string{helpers.DcID: *datacenter.Id,
							helpers.ServerID: *server.Id},
						[]string{},
						map[string]interface{}{}))
				}
			}
		}
	}
	return nil
}
