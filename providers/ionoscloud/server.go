package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
)

type ServerGenerator struct {
	IonosCloudService
}

func (g *ServerGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	datacenters, err := helpers.GetAllDatacenters(*cloudApiClient)
	if err != nil {
		return err
	}

	for _, datacenter := range datacenters {
		servers, _, err := cloudApiClient.ServersApi.DatacentersServersGet(context.TODO(), *datacenter.Id).Depth(2).Execute()
		if err != nil {
			return err
		}
		if servers.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing servers but received 'nil' instead, skipping search for datacenter with ID: %v.\n",
				*datacenter.Id)
			continue
		}
		serversToAdd := *servers.Items
		for _, server := range serversToAdd {
			if server.Properties == nil || server.Properties.Name == nil {
				log.Printf(
					"[WARNING] 'nil' values in the response for server with ID %v, datacenter ID: %v, skipping this resource.\n",
					*server.Id,
					*datacenter.Id,
				)
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*server.Id,
				*server.Properties.Name+"-"+*server.Id,
				"ionoscloud_server",
				helpers.Ionos,
				map[string]string{helpers.DcId: *datacenter.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
