package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	dbaas "github.com/ionos-cloud/sdk-go-dbaas-postgres"
)

type DBaaSPgSQLClusterGenerator struct {
	Service
}

func (g DBaaSPgSQLClusterGenerator) createResources(
	clustersList []dbaas.ClusterResponse,
) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, cluster := range clustersList {
		if cluster.Properties == nil || cluster.Properties.DisplayName == nil {
			log.Printf(
				"[WARNING] 'nil' values in the response for db cluster with ID %v, skipping this resource.\n",
				*cluster.Id,
			)
			continue
		}
		resources = append(resources, terraformutils.NewResource(
			*cluster.Id,
			*cluster.Properties.DisplayName+"-"+*cluster.Id,
			"ionoscloud_pg_cluster",
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return resources
}

func (g *DBaaSPgSQLClusterGenerator) InitResources() error {
	client := g.generateClient()
	dbaasAPIClient := client.DBaaSPgSQLApiClient
	output, _, err := dbaasAPIClient.ClustersApi.ClustersGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if output.Items != nil {
		g.Resources = g.createResources(*output.Items)
	} else {
		log.Printf("[WARNING] expected a response containing db clusters but received 'nil' instead.")
	}
	return nil
}
