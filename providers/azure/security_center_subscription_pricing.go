package azure

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/authentication"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SecurityCenterSubscriptionPricingGenerator struct {
	AzureService
}

func (g SecurityCenterSubscriptionPricingGenerator) listSubscriptionPricing() ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID

	securityCenterPricingClient := security.NewPricingsClient(subscriptionID, "")
	securityCenterPricingClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	pricingList, err := securityCenterPricingClient.List(ctx)
	if err != nil {
		return resources, err
	}

	for _, pricing := range *pricingList.Value {
		resources = append(resources, terraformutils.NewSimpleResource(
			*pricing.ID,
			*pricing.Name,
			"azurerm_security_center_subscription_pricing",
			g.ProviderName,
			[]string{}))
	}

	return resources, nil
}

func (g *SecurityCenterSubscriptionPricingGenerator) InitResources() error {
	resources, err := g.listSubscriptionPricing()
	if err != nil {
		return err
	}

	g.Resources = append(g.Resources, resources...)

	return nil
}
