package ionoscloud

import (
	"context"
	"log"

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
	for _, datacenter := range datacenters {
		servers, _, err := cloudApiClient.ServersApi.DatacentersServersGet(context.TODO(), *datacenter.Id).Depth(1).Execute()
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
			nics, _, err := cloudApiClient.NetworkInterfacesApi.DatacentersServersNicsGet(context.TODO(), *datacenter.Id, *server.Id).Depth(2).Execute()
			if err != nil {
				return err
			}
			if nics.Items == nil {
				log.Printf(
					"[WARNING] expected a response containing NICs but received 'nil' instead, skipping search for server with ID: %v, datacenter ID: %v.\n",
					*server.Id,
					*datacenter.Id)
				continue
			}
			for _, nic := range *nics.Items {
				if nic.Properties == nil || nic.Properties.Name == nil {
					log.Printf(
						"[WARNING] 'nil' values in the response for NIC with ID %v, server ID: %v, datacenter ID: %v, skipping this resource.\n",
						*nic.Id,
						*server.Id,
						*datacenter.Id,
					)
					continue
				}
				// Check if the NIC is not attached to a server, this is required in order to avoid NIC duplicates in
				// plans.
				if *nic.Properties.Lan != 1 {
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
	return nil
}
