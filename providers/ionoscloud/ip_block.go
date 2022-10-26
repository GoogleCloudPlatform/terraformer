package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type IPBlockGenerator struct {
	IonosCloudService
}

func (g *IPBlockGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resource_type := "ionoscloud_ipblock"

	ipBlockResponse, _, err := cloudApiClient.IPBlocksApi.IpblocksGet(context.TODO()).Depth(1).Execute()
	if err != nil {
		return err
	}
	ipBlocks := *ipBlockResponse.Items
	for _, ipBlock := range ipBlocks {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*ipBlock.Id,
			*ipBlock.Properties.Name+"-"+*ipBlock.Id,
			resource_type,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
