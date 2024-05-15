package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DataPlatformClusterGenerator struct {
	Service
}

func (g *DataPlatformClusterGenerator) InitResources() error {
	client := g.generateClient()
	dataPlatformClient := client.DataPlatformAPIClient
	resourceType := "ionoscloud_dataplatform_cluster"

	response, _, err := dataPlatformClient.DataPlatformClusterApi.ClustersGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if response.Items == nil {
		log.Printf("[WARNING] expected a response containing data platform clusters, but received 'nil' instead.")
		return nil
	}
	clusters := *response.Items
	for _, cluster := range clusters {
		if cluster.Properties == nil || cluster.Properties.Name == nil {
			log.Printf("[WARNING] 'nil' values in the response for data platform cluster with ID %v, skipping this resource.", *cluster.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*cluster.Id,
			*cluster.Properties.Name+"-"+*cluster.Id,
			resourceType,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
