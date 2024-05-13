package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type DataPlatformNodePoolGenerator struct {
	Service
}

func (g *DataPlatformNodePoolGenerator) InitResources() error {
	client := g.generateClient()
	dataPlatformClient := client.DataPlatformAPIClient
	resourceType := "ionoscloud_dataplatform_node_pool"

	dpClusters, _, err := dataPlatformClient.DataPlatformClusterApi.ClustersGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if dpClusters.Items == nil {
		log.Printf("[WARNING] expected a response containing data platform clusters but received 'nil' instead.")
		return nil
	}
	for _, dpCluster := range *dpClusters.Items {
		dpNodePools, _, err := dataPlatformClient.DataPlatformNodePoolApi.ClustersNodepoolsGet(context.TODO(), *dpCluster.Id).Execute()
		if err != nil {
			return err
		}
		if dpNodePools.Items == nil {
			log.Printf("[WARNING] expected a response containing data platform node pools but received 'nil' instead, skipping search for data platform cluster with ID: %v", *dpCluster.Id)
			continue
		}
		for _, dpNodePool := range *dpNodePools.Items {
			if dpNodePool.Properties == nil || dpNodePool.Properties.Name == nil {
				log.Printf("[WARNING] 'nil' values in the response for data platform node pool with ID %v, cluster ID: %v, skipping this resource",
					*dpNodePool.Id,
					*dpCluster.Id)
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*dpNodePool.Id,
				*dpNodePool.Properties.Name+"-"+*dpNodePool.Id,
				resourceType,
				helpers.Ionos,
				map[string]string{helpers.ClusterID: *dpCluster.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
