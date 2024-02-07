package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DBaaSPgSQLUserGenerator struct {
	Service
}

func (g *DBaaSPgSQLUserGenerator) InitResources() error {
	client := g.generateClient()
	dbaasPgSQLClient := client.DBaaSPgSQLApiClient
	resourceType := "ionoscloud_pg_user"

	response, _, err := dbaasPgSQLClient.ClustersApi.ClustersGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if response.Items == nil {
		log.Printf("[WARNING] expected a response containing PgSQL DB clusters but received 'nil' instead")
	}
	clusters := *response.Items
	for _, cluster := range clusters {
		usersResponse, _, err := dbaasPgSQLClient.UsersApi.UsersList(context.TODO(), *cluster.Id).Execute()
		if err != nil {
			return err
		}
		if usersResponse.Items == nil {
			log.Printf("[WARNING] expected a response containing PgSQL users but received 'nil' instead, skipping search for PgSQL cluster with ID: %v", *cluster.Id)
			continue
		}
		users := *usersResponse.Items
		for _, user := range users {
			if user.Properties == nil || user.Properties.Username == nil {
				log.Printf("[WARNING] 'nil' values in the response for PgSQL user, skipping this resource")
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*user.Id,
				*user.Properties.Username+"-"+*user.Id,
				resourceType,
				helpers.Ionos,
				map[string]string{helpers.ClusterID: *cluster.Id, helpers.UsernameArg: *user.Properties.Username},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
