package grafana

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	gapi "github.com/grafana/grafana-api-golang-client"
)

type FolderGenerator struct {
	GrafanaService
}

func (g *FolderGenerator) InitResources() error {
	client, err := g.buildClient()
	if err != nil {
		return fmt.Errorf("unable to build grafana client: %v", err)
	}

	err = g.createFolderResources(client)
	if err != nil {
		return err
	}

	return nil
}

func (g *FolderGenerator) createFolderResources(client *gapi.Client) error {
	folders, err := client.Folders()
	if err != nil {
		return fmt.Errorf("unable to list grafana folders: %v", err)
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
