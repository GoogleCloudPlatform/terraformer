package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
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
	kubernetesClusters := *kubernetesClusterResponse.Items
	for _, kubernetesCluster := range kubernetesClusters {
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
