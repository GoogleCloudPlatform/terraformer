package grafana

import (
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	gapi "github.com/grafana/grafana-api-golang-client"
)

type DashboardGenerator struct {
	GrafanaService
}

func (g *DashboardGenerator) InitResources() error {
	client, err := g.buildClient()
	if err != nil {
		return fmt.Errorf("unable to build grafana client: %v", err)
	}

	err = g.createDashboardResources(client)
	if err != nil {
		return err
	}

	return nil
}

func (g *DashboardGenerator) createDashboardResources(client *gapi.Client) error {
	dashboards, err := client.Dashboards()
	if err != nil {
		return fmt.Errorf("unable to list grafana dashboards: %v", err)
	}

	for _, dashboard := range dashboards {
		// search result doesn't include slug, so need to look up dashboard.
		dash, err := client.DashboardByUID(dashboard.UID)
		if err != nil {
			return fmt.Errorf("unable to read grafana dashboard %s: %v", dashboard.Title, err)
		}

		configJSON, err := json.MarshalIndent(dash.Model, "", "  ")
		if err != nil {
			return fmt.Errorf("unable to marshal configuration for grafana dashboard %s: %v", dashboard.Title, err)
		}

		filename := fmt.Sprintf("dashboard-%s.json", dash.Meta.Slug)
		resource := terraformutils.NewResource(
			dashboard.UID,
			dashboard.Title,
			"grafana_dashboard",
			"grafana",
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"config_json": fmt.Sprintf("file(\"data/%s\")", filename),
				"folder":      dashboard.FolderID,
			},
		)
		resource.DataFiles = map[string][]byte{
			filename: configJSON,
		}
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
