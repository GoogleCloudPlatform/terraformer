package authentication

import (
	"fmt"
	"log"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/hashicorp/go-multierror"
)

type managedServiceIdentityAuth struct {
	endpoint string
}

func (a managedServiceIdentityAuth) build(b Builder) (authMethod, error) {
	endpoint := b.MsiEndpoint
	if endpoint == "" {
		msiEndpoint, err := adal.GetMSIVMEndpoint()
		if err != nil {
			return nil, fmt.Errorf("Error determining MSI Endpoint: ensure the VM has MSI enabled, or configure the MSI Endpoint. Error: %s", err)
		}
		endpoint = msiEndpoint
	}

	log.Printf("[DEBUG] Using MSI endpoint %q", endpoint)

	auth := managedServiceIdentityAuth{
		endpoint: endpoint,
	}
	return auth, nil
}

func (a managedServiceIdentityAuth) isApplicable(b Builder) bool {
	return b.SupportsManagedServiceIdentity
}

func (a managedServiceIdentityAuth) name() string {
	return "Managed Service Identity"
}

func (a managedServiceIdentityAuth) getAuthorizationToken(oauthConfig *adal.OAuthConfig, endpoint string) (*autorest.BearerAuthorizer, error) {
	spt, err := adal.NewServicePrincipalTokenFromMSI(a.endpoint, endpoint)
	if err != nil {
		return nil, err
	}
	auth := autorest.NewBearerAuthorizer(spt)
	return auth, nil
}

func (a managedServiceIdentityAuth) populateConfig(c *Config) error {
	// nothing to populate back
	return nil
}

func (a managedServiceIdentityAuth) validate() error {
	var err *multierror.Error

	if a.endpoint == "" {
		err = multierror.Append(err, fmt.Errorf("An MSI Endpoint must be configured"))
	}

	return err.ErrorOrNil()
}
