package grafana

import (
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type FolderGenerator struct {
	GrafanaService
}

func (g *FolderGenerator) InitResources() error {
	client, err := g.buildClient()
	if err != nil {
		return err
	}

	folders, err := client.Folders()
	if err != nil {
		return err
	}

	for _, folder := range folders {
		g.Resources = append(g.Resources, terraformutils.NewResource(
			fmt.Sprint(folder.ID),
			folder.Title,
			"grafana_folder",
			"grafana",
			map[string]string{
				"uid": folder.UID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return nil
}
