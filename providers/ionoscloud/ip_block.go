package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type IPBlockGenerator struct {
	Service
}

func (g *IPBlockGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	resourceType := "ionoscloud_ipblock"

	ipBlockResponse, _, err := cloudAPIClient.IPBlocksApi.IpblocksGet(context.TODO()).Depth(1).Execute()
	if err != nil {
		return err
	}
	if ipBlockResponse.Items == nil {
		log.Printf("[WARNING] expected a response containing IP blocks but received 'nil' instead.")
		return nil
	}
	ipBlocks := *ipBlockResponse.Items
	for _, ipBlock := range ipBlocks {
		if ipBlock.Properties == nil || ipBlock.Properties.Name == nil {
			log.Printf(
				"[WARNING] 'nil' values in the response for IP block with ID %v, skipping this resource.\n",
				*ipBlock.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*ipBlock.Id,
			*ipBlock.Properties.Name+"-"+*ipBlock.Id,
			resourceType,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
