package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
)

type KubernetesClusterGenerator struct {
	IonosCloudService
}

func (g *KubernetesClusterGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resource_type := "ionoscloud_k8s_cluster"

	kubernetesClusterResponse, _, err := cloudApiClient.KubernetesApi.K8sGet(context.TODO()).Depth(1).Execute()
	if err != nil {
		return err
	}
	if kubernetesClusterResponse.Items == nil {
		log.Printf("[WARNING] expected a response containing k8s clusters but received 'nil' instead.")
		return nil
	}
	kubernetesClusters := *kubernetesClusterResponse.Items
	for _, kubernetesCluster := range kubernetesClusters {
		if kubernetesCluster.Properties == nil || kubernetesCluster.Properties.Name == nil {
			log.Printf(
				"[WARNING] 'nil' values in the response for k8s cluster with ID %v, skipping this resource.\n",
				*kubernetesCluster.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*kubernetesCluster.Id,
			*kubernetesCluster.Properties.Name+"-"+*kubernetesCluster.Id,
			resource_type,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
