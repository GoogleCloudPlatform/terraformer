package azure

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/authentication"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2019-08-01/web"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type AppServiceGenerator struct {
	AzureService
}

func (g AppServiceGenerator) listApps() ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()

	appServiceClient := web.NewAppsClient(g.Args["config"].(authentication.Config).SubscriptionID)
	appServiceClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	apps, err := appServiceClient.List(ctx)
	if err != nil {
		return nil, err
	}
	for _, site := range apps.Values() {
		resources = append(resources, terraformutils.NewSimpleResource(
			*site.ID,
			*site.Name,
			"azurerm_app_service",
			g.ProviderName,
			[]string{}))
	}

	return resources, nil
}

func (g *AppServiceGenerator) InitResources() error {
	resources, err := g.listApps()
	if err != nil {
		return err
	}

	g.Resources = append(g.Resources, resources...)

	return nil
}
