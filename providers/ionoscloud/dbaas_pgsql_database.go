package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DBaaSPgSQLDatabaseGenerator struct {
	Service
}

func (g *DBaaSPgSQLDatabaseGenerator) InitResources() error {
	client := g.generateClient()
	dbaasPgSQLClient := client.DBaaSPgSQLApiClient
	resourceType := "ionoscloud_pg_database"

	response, _, err := dbaasPgSQLClient.ClustersApi.ClustersGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if response.Items == nil {
		log.Printf("[WARNING] expected a response containing PgSQL DB clusters but received 'nil' instead")
	}
	clusters := *response.Items
	for _, cluster := range clusters {
		databasesResponse, _, err := dbaasPgSQLClient.DatabasesApi.DatabasesList(context.TODO(), *cluster.Id).Execute()
		if err != nil {
			return err
		}
		if databasesResponse.Items == nil {
			log.Printf("[WARNING] expected a response containing PgSQL databases but received 'nil' instead, skipping search for PgSQL cluster with ID: %v", *cluster.Id)
			continue
		}
		databases := *databasesResponse.Items
		for _, database := range databases {
			if database.Properties == nil || database.Properties.Name == nil {
				log.Printf("[WARNING] 'nil' values in the response for PgSQL database with ID: %v, skipping this resource", *database.Id)
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*database.Id,
				*database.Properties.Name+"-"+*database.Id,
				resourceType,
				helpers.Ionos,
				map[string]string{helpers.ClusterID: *cluster.Id, helpers.NameArg: *database.Properties.Name},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
