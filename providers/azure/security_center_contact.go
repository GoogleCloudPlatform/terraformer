package azure

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/authentication"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type SecurityCenterContactGenerator struct {
	AzureService
}

func (g SecurityCenterContactGenerator) listContacts() ([]terraformutils.Resource, error) {
	var resources []terraformutils.Resource
	ctx := context.Background()
	subscriptionID := g.Args["config"].(authentication.Config).SubscriptionID
	resourceManagerEndpoint := g.Args["config"].(authentication.Config).CustomResourceManagerEndpoint

	securityCenterContactClient := security.NewContactsClientWithBaseURI(resourceManagerEndpoint, subscriptionID, "")
	securityCenterContactClient.Authorizer = g.Args["authorizer"].(autorest.Authorizer)

	if rg := g.Args["resource_group"].(string); rg != "" {
		return resources, nil
	}
	contactsIterator, err := securityCenterContactClient.ListComplete(ctx)
	if err != nil {
		return resources, err
	}

	for contactsIterator.NotDone() {
		contact := contactsIterator.Value()
		resources = append(resources, terraformutils.NewSimpleResource(
			*contact.ID,
			*contact.Name,
			"azurerm_security_center_contact",
			g.ProviderName,
			[]string{}))

		if err := contactsIterator.NextWithContext(ctx); err != nil {
			return resources, err
		}
	}

	return resources, nil
}

func (g *SecurityCenterContactGenerator) InitResources() error {
	resources, err := g.listContacts()
	if err != nil {
		return err
	}

	g.Resources = append(g.Resources, resources...)

	return nil
}
