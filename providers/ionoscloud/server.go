package ionoscloud

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

type ServerGenerator struct {
	IonosCloudService
}

func (g ServerGenerator) createResources(serversList []ionoscloud.Server) []terraformutils.Resource {
	var resources []terraformutils.Resource
	fmt.Printf("servers LIST %v", serversList)
	for _, server := range serversList {
		fmt.Printf("server ID: %v\n", *server.Id)

		resources = append(resources, terraformutils.NewResource(
			*server.Id,
			*server.Properties.Name+"-"+*server.Id,
			"ionoscloud_server",
			"ionoscloud",
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return resources
}

func (g *ServerGenerator) InitResources() error {
	var serversOuput []ionoscloud.Server
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	datacenters, err := helpers.GetAllDatacenters(*cloudApiClient)
	if err != nil {
		return err
	}
	if datacenters != nil {
		for _, datacenter := range datacenters {
			fmt.Printf("Datacenters: %v \n", *datacenter.Id)
			servers, _, err := cloudApiClient.ServersApi.DatacentersServersGet(context.TODO(), *datacenter.Id).Depth(5).Execute()
			if err != nil {
				return err
			}
			serversToAdd := *servers.Items
			fmt.Printf("servers to add %v \n", serversToAdd)
			for _, server := range serversToAdd {
				if server.Id != nil {
					serversOuput = append(serversOuput, server)
				}
			}

		}
	}
	g.Resources = g.createResources(serversOuput)
	return nil
}
