package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DBaaSMongoClusterGenerator struct {
	Service
}

func (g *DBaaSMongoClusterGenerator) InitResources() error {
	client := g.generateClient()
	dbaasMongoClient := client.DBaaSMongoAPIClient
	resourceType := "ionoscloud_mongo_cluster"

	response, _, err := dbaasMongoClient.ClustersApi.ClustersGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if response.Items == nil {
		log.Printf("[WARNING] expected a response containing Mongo DB clusters but received 'nil' instead")
	}
	clusters := *response.Items
	for _, cluster := range clusters {
		if cluster.Properties == nil || cluster.Properties.DisplayName == nil {
			log.Printf("[WARNING] 'nil' values in the response for Mongo DB cluster with ID: %v, skipping search for this resource", *cluster.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*cluster.Id,
			*cluster.Properties.DisplayName+"-"+*cluster.Id,
			resourceType,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
