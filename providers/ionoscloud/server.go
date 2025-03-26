package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

type ServerGenerator struct {
	Service
}

func (g *ServerGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	datacenters, err := helpers.GetAllDatacenters(*cloudAPIClient)
	if err != nil {
		return err
	}

	for _, datacenter := range datacenters {
		servers, _, err := cloudAPIClient.ServersApi.DatacentersServersGet(context.TODO(), *datacenter.Id).Depth(4).Execute()
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
			if !isServerValid(server, *datacenter.Id) {
				continue
			}
			_, apiResponse, err := cloudAPIClient.LabelsApi.DatacentersServersLabelsFindByKey(context.TODO(), *datacenter.Id, *server.Id, "managedexternally").Execute()
			if err != nil {
				if !apiResponse.HttpNotFound() {
					return err
				}
			} else {
				// The server is managed externally(eg : k8s nodepool).
				// This means we do not want to write the server to the tf plan.
				continue
			}

			resourceType := getServerResourceType(*server.Properties.Type)
			if resourceType == "" {
				log.Printf("[WARNING] unknown server type: %v for server with ID: %v, skipping this server", *server.Properties.Type, *server.Id)
				continue
			}

			g.Resources = append(g.Resources, terraformutils.NewResource(
				*server.Id,
				*server.Properties.Name+"-"+*server.Id,
				resourceType,
				helpers.Ionos,
				map[string]string{helpers.DcID: *datacenter.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}

// isServerValid skips servers that would not create a valid tf plan.
func isServerValid(server ionoscloud.Server, datacenterID string) bool {

	if server.Properties == nil || server.Properties.Name == nil {
		log.Printf(
			"[WARNING] 'nil' values in the response for server with ID %v, datacenter ID: %v, skipping this resource.\n",
			*server.Id,
			datacenterID,
		)
		return false
	}
	if server.Entities.Nics == nil || len(*server.Entities.Nics.Items) == 0 {
		log.Printf("Server %s, datacenter ID: %v  contains no nics, moving on", *server.Id, datacenterID)
		return false
	}
	if server.Entities.Volumes == nil || len(*server.Entities.Volumes.Items) == 0 {
		log.Printf("Server %s, datacenter ID: %v contains no volumes, moving on", *server.Id, datacenterID)
		return false
	}
	if server.Properties.BootVolume == nil {
		log.Printf("Server %s, datacenter ID: %v  contains no boot volume, moving on", *server.Id, datacenterID)
		return false
	}

	return true
}

func getServerResourceType(serverType string) string {
	resourceType := ""
	switch serverType {
	case "ENTERPRISE":
		resourceType = "ionoscloud_server"
	case "CUBE":
		resourceType = "ionoscloud_cube_server"
	case "VCPU":
		resourceType = "ionoscloud_vcpu_server"
	}
	return resourceType
}
