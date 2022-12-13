package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
)

type GroupGenerator struct {
	IonosCloudService
}

func (g *GroupGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resource_type := "ionoscloud_group"

	groupResponse, _, err := cloudApiClient.UserManagementApi.UmGroupsGet(context.TODO()).Depth(1).Execute()
	if err != nil {
		return err
	}
	if groupResponse.Items == nil {
		log.Printf("[WARNING] expected a response containing groups but received 'nil' instead.")
		return nil
	}
	groups := *groupResponse.Items
	for _, group := range groups {
		if group.Properties == nil || group.Properties.Name == nil {
			log.Printf(
				"[WARNING] 'nil' values in the response for group with ID %v, skipping this resource.\n",
				*group.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*group.Id,
			*group.Properties.Name+"-"+*group.Id,
			resource_type,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}