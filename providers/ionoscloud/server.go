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
			// skip servers that do not contain NICs or Volumes. They would create
			// invalid Terraform plans, as the server resource requires a NIC and a Volume
			if server.Entities.Nics == nil || server.Entities.Nics.Items == nil || len(*server.Entities.Nics.Items) == 0 {
				log.Printf("Server %s from datacenter %s contains no nics, moving on", *server.Id, *datacenter.Id)
				continue
			}
			if server.Entities.Volumes == nil || server.Entities.Volumes.Items == nil || len(*server.Entities.Volumes.Items) == 0 {
				log.Printf("Server %s from datacenter %s contains no volumes, moving on", *server.Id, *datacenter.Id)
				continue
			}
			_, apiResponse, err := cloudApiClient.LabelsApi.DatacentersServersLabelsFindByKey(context.TODO(), *datacenter.Id, *server.Id, "managedexternally").Execute()
			if err != nil {
				if !apiResponse.HttpNotFound() {
					return err
				}
			} else {
				//the server is managed externally(eg : k8s nodepool). This means we do not want to write the server to the tf plan.
				continue
			}
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
