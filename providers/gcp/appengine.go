package gcp

import (
	"context"

	appengine "cloud.google.com/go/appengine/apiv1"
	"cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/iterator"
)

var appengineAllowEmptyValues = []string{}

type AppengineGenerator struct {
	GCPService
}

func (g *AppengineGenerator) InitResources() error {
	ctx := context.Background()
	servicesClient, err := appengine.NewServicesClient(ctx)
	if err != nil {
		return err
	}
	defer servicesClient.Close()

	req := &appenginepb.ListServicesRequest{Parent: "apps/" + g.GetArgs()["project"].(string)}

	it := servicesClient.ListServices(ctx, req)

	for {
		svc, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				return nil
			}
			return err
		}

		g.loadVersions(svc)
	}
}

func (g *AppengineGenerator) loadVersions(service *appenginepb.Service) error {
	ctx := context.Background()
	versionsClient, err := appengine.NewVersionsClient(ctx)
	if err != nil {
		return err
	}
	defer versionsClient.Close()

	req := &appenginepb.ListVersionsRequest{Parent: service.Name, View: appenginepb.VersionView_FULL}

	it := versionsClient.ListVersions(ctx, req)

	for {
		version, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				return nil
			}
			return err
		}

		if version.Env == "standard" {
			g.Resources = append(g.Resources, terraformutils.NewResource(
				version.Name,
				version.Name,
				"google_app_engine_standard_app_version",
				g.GetProviderName(),
				map[string]string{
					"service": service.Id,
					"project": g.GetArgs()["project"].(string),
				},
				appengineAllowEmptyValues,
				map[string]interface{}{
					"runtime":         version.Runtime,
					"deployment":      version.Deployment,
					"handlers":        version.Handlers,
					"version_id":      version.Id,
					"service_account": version.ServiceAccount,
				},
			))
		}
	}
}
