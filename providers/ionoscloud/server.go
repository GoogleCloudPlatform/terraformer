package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
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

	if datacenters != nil {
		for _, datacenter := range datacenters {
			servers, _, err := cloudApiClient.ServersApi.DatacentersServersGet(context.TODO(), *datacenter.Id).Depth(2).Execute()
			if err != nil {
				return err
			}
			if servers.Items != nil {
				serversToAdd := *servers.Items

				for _, server := range serversToAdd {
					if server.Id != nil && server.Properties != nil && server.Properties.Name != nil {
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
			}
		}
	}
	return nil
}
