package connectivity

import (
	"context"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"golang.org/x/oauth2/clientcredentials"
)

func (c *Config) NewClient() *commercetools.Client {
	oauth2Config := &clientcredentials.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		Scopes:       []string{c.ClientScope},
		TokenURL:     c.TokenURL,
	}

	httpClient := oauth2Config.Client(context.TODO())

	return commercetools.New(&commercetools.Config{
		ProjectKey:  "lca-dev",
		URL:         c.BaseURL,
		HTTPClient:  httpClient,
		LibraryName: "terraformer",
	})
}
