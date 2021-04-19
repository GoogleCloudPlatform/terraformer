package grafana

import (
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/pkg/errors"
)

type DashboardGenerator struct {
	GrafanaService
}

func (g *DashboardGenerator) InitResources() error {
	client, err := g.buildClient()
	if err != nil {
		return errors.Wrap(err, "unable to build grafana client")
	}

	dashboards, err := client.Dashboards()
	if err != nil {
		return errors.Wrap(err, "unable to list dashboards")
	}

	for _, dashboard := range dashboards {
		// search result doesn't include slug, so need to look up dashboard.
		dash, err := client.DashboardByUID(dashboard.UID)
		if err != nil {
			return errors.Wrapf(err, "unable to read dashboard %s", dashboard.Title)
		}
		configJson, err := json.MarshalIndent(dash.Model, "", "  ")
		if err != nil {
			return errors.Wrapf(err, "unable to marshal configuration for dashboard %s", dashboard.Title)
		}

		filename := fmt.Sprintf("dashboard-%s.json", dash.Meta.Slug)
		resource := terraformutils.NewResource(
			dash.Meta.Slug,
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
			filename: configJson,
		}
		g.Resources = append(g.Resources, resource)
	}

	return nil
}
