package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
)

type BackupUnitGenerator struct {
	IonosCloudService
}

func (g *BackupUnitGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resourceType := "ionoscloud_backup_unit"

	backupUnitResponse, _, err := cloudApiClient.BackupUnitsApi.BackupunitsGet(context.TODO()).Depth(1).Execute()
	if err != nil {
		return err
	}
	if backupUnitResponse.Items == nil {
		log.Printf("[WARNING] expected a response containing backup units but received 'nil' instead.")
		return nil
	}
	backupUnits := *backupUnitResponse.Items
	for _, backupUnit := range backupUnits {
		if backupUnit.Properties == nil || backupUnit.Properties.Name == nil {
			log.Printf(
				"[WARNING] 'nil' values in the response for backup unit with ID %v, skipping this resource.\n",
				*backupUnit.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*backupUnit.Id,
			*backupUnit.Properties.Name+"-"+*backupUnit.Id,
			resourceType,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
