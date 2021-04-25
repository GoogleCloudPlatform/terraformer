package connectivity

import (
	"context"
	"strings"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"golang.org/x/oauth2/clientcredentials"
)

func (c *Config) NewClient() *commercetools.Client {
	oauth2Config := &clientcredentials.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		Scopes:       strings.Split(c.ClientScope, " "),
		TokenURL:     c.TokenURL,
	}

	httpClient := oauth2Config.Client(context.TODO())

	return commercetools.New(&commercetools.Config{
		ProjectKey:  c.ProjectKey,
		URL:         c.BaseURL,
		HTTPClient:  httpClient,
		LibraryName: "terraformer",
	})
}
