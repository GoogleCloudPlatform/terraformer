package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type KubernetesNodePoolGenerator struct {
	IonosCloudService
}

func (g *KubernetesNodePoolGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resource_type := "ionoscloud_k8s_node_pool"

	kubernetesClusters, _, err := cloudApiClient.KubernetesApi.K8sGet(context.TODO()).Depth(1).Execute()
	if err != nil {
		return err
	}
	for _, kubernetesCluster := range *kubernetesClusters.Items {
		kubernetesNodePools, _, err := cloudApiClient.KubernetesApi.K8sNodepoolsGet(context.TODO(), *kubernetesCluster.Id).Depth(1).Execute()
		if err != nil {
			return err
		}
		for _, kubernetesNodePool := range *kubernetesNodePools.Items {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				*kubernetesNodePool.Id,
				*kubernetesNodePool.Properties.Name+"-"+*kubernetesNodePool.Id,
				resource_type,
				helpers.Ionos,
				map[string]string{helpers.K8sClusterId: *kubernetesCluster.Id},
				[]string{},
				map[string]interface{}{}))
		}
	}
	return nil
}
