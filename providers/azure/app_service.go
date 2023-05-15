package azure

import (
	"context"
	"log"

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

	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint
	appServiceClient := web.NewAppsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	appServiceClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)
	var (
		appsIterator web.AppCollectionIterator
		err          error
	)
	if rg := g.Args["resource_group"].(string); rg != "" {
		appsIterator, err = appServiceClient.ListByResourceGroupComplete(ctx, rg, nil)
	} else {
		appsIterator, err = appServiceClient.ListComplete(ctx)
	}
	if err != nil {
		return nil, err
	}
	for appsIterator.NotDone() {
		site := appsIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*site.ID,
			*site.Name,
			"azurerm_app_service",
			g.ProviderName,
			[]string{}))

		if err := appsIterator.NextWithContext(ctx); err != nil {
			log.Println(err)
			return resources, err
		}
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
