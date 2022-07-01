package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	dbaas "github.com/ionos-cloud/sdk-go-dbaas-postgres"
)

type DBaaSClusterGenerator struct {
	IonosCloudService
}

func (g DBaaSClusterGenerator) createResources(clustersList []dbaas.ClusterResponse) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, cluster := range clustersList {
		resources = append(resources, terraformutils.NewSimpleResource(
			*cluster.Id,
			*cluster.Properties.DisplayName+"-"+*cluster.Id,
			"ionoscloud_pg_cluster",
			"ionoscloud",
			[]string{}))
	}
	return resources
}

func (g *DBaaSClusterGenerator) InitResources() error {
	client := g.generateClient()
	dbaasApiClient := client.DBaaSApiClient
	output, _, err := dbaasApiClient.ClustersApi.ClustersGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	g.Resources = g.createResources(*output.Items)
	return nil
}
