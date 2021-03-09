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
		configJson, err := json.Marshal(dash.Model)
		if err != nil {
			return errors.Wrapf(err, "unable to marshal configuration for dashboard %s", dashboard.Title)
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			dash.Meta.Slug,
			dashboard.Title,
			"grafana_dashboard",
			"grafana",
			map[string]string{},
			[]string{},
			map[string]interface{}{
				"config_json":  string(configJson),
				"folder":       dashboard.FolderID,
				"slug":         dash.Meta.Slug,
				"dashboard_id": fmt.Sprint(dashboard.ID),
			},
		))
	}

	return nil
}
