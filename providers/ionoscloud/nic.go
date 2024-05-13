package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type NicGenerator struct {
	Service
}

func (g *NicGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	datacenters, err := helpers.GetAllDatacenters(*cloudAPIClient)
	if err != nil {
		return err
	}
	for _, datacenter := range datacenters {
		servers, _, err := cloudAPIClient.ServersApi.DatacentersServersGet(context.TODO(), *datacenter.Id).Execute()
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
			nics, _, err := cloudAPIClient.NetworkInterfacesApi.DatacentersServersNicsGet(context.TODO(), *datacenter.Id, *server.Id).Depth(1).Execute()
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
			lastNicIdx := len(*nics.Items) - 1
			for idx, nic := range *nics.Items {
				// skip the last nic from the list, as it will be added to the server separately.
				if idx == lastNicIdx {
					continue
				}
				if nic.Properties == nil || nic.Properties.Name == nil {
					log.Printf(
						"[WARNING] 'nil' values in the response for NIC with ID %v, server ID: %v, datacenter ID: %v, skipping this resource.\n",
						*nic.Id,
						*server.Id,
						*datacenter.Id,
					)
					continue
				}
				g.Resources = append(g.Resources, terraformutils.NewResource(
					*nic.Id,
					*nic.Properties.Name+"-"+*nic.Id,
					"ionoscloud_nic",
					helpers.Ionos,
					map[string]string{helpers.DcID: *datacenter.Id,
						helpers.ServerID: *server.Id},
					[]string{},
					map[string]interface{}{}))
			}
		}
	}
	return nil
}
