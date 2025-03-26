package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type KubernetesNodePoolGenerator struct {
	Service
}

func (g *KubernetesNodePoolGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	resourceType := "ionoscloud_k8s_node_pool"

	kubernetesClusters, _, err := cloudAPIClient.KubernetesApi.K8sGet(context.TODO()).Execute()
	if err != nil {
		return err
	}
	if kubernetesClusters.Items == nil {
		log.Printf("[WARNING] expected a response containing k8s clusters but received 'nil' instead.")
		return nil
	}
	for _, kubernetesCluster := range *kubernetesClusters.Items {
		kubernetesNodePools, _, err := cloudAPIClient.KubernetesApi.K8sNodepoolsGet(context.TODO(), *kubernetesCluster.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		if kubernetesNodePools.Items == nil {
			log.Printf(
				"[WARNING] expected a response containing k8s node pools but received 'nil' instead, skipping search for k8s cluster with ID: %v.\n",
				*kubernetesCluster.Id)
			continue
		}
		for _, kubernetesNodePool := range *kubernetesNodePools.Items {
			if kubernetesNodePool.Properties == nil || kubernetesNodePool.Properties.Name == nil {
				log.Printf(
					"[WARNING] 'nil' values in the response for k8s node pool with ID %v, k8s cluster ID: %v, skipping this resource.\n",
					*kubernetesNodePool.Id,
					*kubernetesCluster.Id)
				continue
			}
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*kubernetesNodePool.Id,
				*kubernetesNodePool.Properties.Name+"-"+*kubernetesNodePool.Id,
				resourceType,
				helpers.Ionos,
				map[string]string{helpers.K8sClusterID: *kubernetesCluster.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
