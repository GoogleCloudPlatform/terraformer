package ionoscloud

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type PrivateCrossConnectGenerator struct {
	Service
}

func (g *PrivateCrossConnectGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	resourceType := "ionoscloud_private_crossconnect"

	pccsResponse, _, err := cloudAPIClient.PrivateCrossConnectsApi.PccsGet(context.TODO()).Depth(1).Execute()
	if err != nil {
		return err
	}
	if pccsResponse.Items == nil {
		log.Printf("[WARNING] expected a response containing PCCs but received 'nil' instead.\n")
		return nil
	}
	pccs := *pccsResponse.Items
	for _, pcc := range pccs {
		if pcc.Properties == nil || pcc.Properties.Name == nil {
			log.Printf("[WARNING] 'nil' values in the response for PCC with ID %v, skipping this resource.\n", *pcc.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*pcc.Id,
			*pcc.Properties.Name+"-"+*pcc.Id,
			resourceType,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
