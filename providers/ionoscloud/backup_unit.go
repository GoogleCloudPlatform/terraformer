package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type BackupUnitGenerator struct {
	IonosCloudService
}

func (g *BackupUnitGenerator) InitResources() error {
	client := g.generateClient()
	cloudApiClient := client.CloudApiClient
	resource_type := "ionoscloud_backup_unit"

	backupUnitResponse, _, err := cloudApiClient.BackupUnitsApi.BackupunitsGet(context.TODO()).Depth(1).Execute()
	if err != nil {
		return err
	}
	backupUnits := *backupUnitResponse.Items
	for _, backupUnit := range backupUnits {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*backupUnit.Id,
			*backupUnit.Properties.Name+"-"+*backupUnit.Id,
			resource_type,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
