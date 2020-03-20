package fastly

import (
	"fmt"
	"net/http"
	"sort"
	"time"
)

// TokenScope is used to match possible authorization scopes
type TokenScope string

const (
	// GlobalScope is the default scope covering all supported capabilities.
	GlobalScope TokenScope = "global"
	// PurgeSelectScope allows purging with surrogate key and URL, disallows purging with purge all.
	PurgeSelectScope TokenScope = "purge_select"
	// PurgeAllScope allows purging an entire service via purge_all.
	PurgeAllScope TokenScope = "purge_all"
	// GlobalReadScope allows read-only access to account information, configuration, and stats.
	GlobalReadScope TokenScope = "global:read"
)

// Token represents an API token which are used to authenticate requests to the
// Fastly API.
type Token struct {
	ID          string     `mapstructure:"id"`
	Name        string     `mapstructure:"name"`
	UserID      string     `mapstructure:"user_id"`
	Services    []string   `mapstructure:"services"`
	AccessToken string     `mapstructure:"access_token"`
	Scope       TokenScope `mapstructure:"scope"`
	IP          string     `mapstructure:"ip"`
	CreatedAt   *time.Time `mapstructure:"created_at"`
	LastUsedAt  *time.Time `mapstructure:"last_used_at"`
	ExpiresAt   *time.Time `mapstructure:"expires_at"`
}

// tokensByName is a sortable list of tokens.
type tokensByName []*Token

// Len, Swap, and Less implement the sortable interface.
func (s tokensByName) Len() int      { return len(s) }
func (s tokensByName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s tokensByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// ListTokens returns the full list of tokens belonging to the currently
// authenticated user.
func (c *Client) ListTokens() ([]*Token, error) {
	resp, err := c.Get("/tokens", nil)
	if err != nil {
		return nil, err
	}

	var t []*Token
	if err := decodeJSON(&t, resp.Body); err != nil {
		return nil, err
	}
	sort.Stable(tokensByName(t))
	return t, nil
}

// ListCustomerTokensInput is used as input to the ListCustomerTokens function.
type ListCustomerTokensInput struct {
	ID string
}

// ListCustomerTokens returns the full list of tokens belonging to a specific
// customer.
func (c *Client) ListCustomerTokens(i *ListCustomerTokensInput) ([]*Token, error) {
	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/customer/%s/tokens", i.ID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var t []*Token
	if err := decodeJSON(&t, resp.Body); err != nil {
		return nil, err
	}
	sort.Stable(tokensByName(t))
	return t, nil
}

// GetTokenSelf retrieves the token information for the the access_token used
// used to authenticate the request. Returns a 401 if the token has expired
// and a 403 for invalid access token.
func (c *Client) GetTokenSelf() (*Token, error) {
	resp, err := c.Get("/tokens/self", nil)
	if err != nil {
		return nil, err
	}

	var t *Token
	if err := decodeJSON(&t, resp.Body); err != nil {
		return nil, err
	}

	return t, nil
}

// CreateTokenInput is used as input to the Token function.
type CreateTokenInput struct {
	Name      string     `form:"name,omitempty"`
	Scope     TokenScope `form:"scope,omitempty"`
	Username  string     `form:"username,omitempty"`
	Password  string     `form:"password,omitempty"`
	Services  []string   `form:"services,omitempty"`
	ExpiresAt *time.Time `form:"expires_at,omitempty"`
}

// CreateToken creates a new API token with the given information.
func (c *Client) CreateToken(i *CreateTokenInput) (*Token, error) {
	_, err := c.PostForm("/sudo", i, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.PostForm("/tokens", i, nil)
	if err != nil {
		return nil, err
	}

	var t *Token
	if err := decodeJSON(&t, resp.Body); err != nil {
		return nil, err
	}
	return t, nil
}

// DeleteTokenInput is used as input to the DeleteToken function.
type DeleteTokenInput struct {
	ID string
}

// DeleteToken revokes a specific token by its ID.
func (c *Client) DeleteToken(i *DeleteTokenInput) error {
	if i.ID == "" {
		return ErrMissingID
	}

	path := fmt.Sprintf("/tokens/%s", i.ID)
	resp, err := c.Delete(path, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Not Ok")
	}
	return nil
}

// DeleteTokenSelf revokes the token used to authorise the request.
func (c *Client) DeleteTokenSelf() error {
	resp, err := c.Delete("/tokens/self", nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Not Ok")
	}
	return nil
}
