package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
)

type TargetGroupGenerator struct {
	IonosCloudService
}

func (g *TargetGroupGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resource_type := "ionoscloud_target_group"

	targetGroupResponse, _, err := cloudApiClient.TargetGroupsApi.TargetgroupsGet(context.TODO()).Depth(1).Execute()
	if err != nil {
		return err
	}
	if targetGroupResponse.Items == nil {
		log.Printf("[WARNING] expected a response containing target groups but received 'nil' instead.")
		return nil
	}
	targetGroups := *targetGroupResponse.Items
	for _, targetGroup := range targetGroups {
		if targetGroup.Properties == nil || targetGroup.Properties.Name == nil {
			log.Printf(
				"[WARNING] 'nil' values in the response for target group with ID %v, skipping this resource.\n",
				*targetGroup.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*targetGroup.Id,
			*targetGroup.Properties.Name+"-"+*targetGroup.Id,
			resource_type,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
