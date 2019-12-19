package fastly

import (
	"fmt"
	"net/url"
	"sort"
	"time"
)

const (
	// SnippetTypeInit sets the type to init
	SnippetTypeInit SnippetType = "init"

	// SnippetTypeRecv sets the type to recv
	SnippetTypeRecv SnippetType = "recv"

	// SnippetTypeHash sets the type to hash
	SnippetTypeHash SnippetType = "hash"

	// SnippetTypeHit sets the type to hit
	SnippetTypeHit SnippetType = "hit"

	// SnippetTypeMiss sets the type to miss
	SnippetTypeMiss SnippetType = "miss"

	// SnippetTypePass sets the type to pass
	SnippetTypePass SnippetType = "pass"

	// SnippetTypeFetch sets the type to fetch
	SnippetTypeFetch SnippetType = "fetch"

	// SnippetTypeError sets the type to error
	SnippetTypeError SnippetType = "error"

	// SnippetTypeDeliver sets the type to deliver
	SnippetTypeDeliver SnippetType = "deliver"

	// SnippetTypeLog sets the type to log
	SnippetTypeLog SnippetType = "log"

	// SnippetTypeNone sets the type to none
	SnippetTypeNone SnippetType = "none"
)

// SnippetType is the type of VCL Snippet
type SnippetType string

// Snippet is the Fastly Snippet object
type Snippet struct {
	ServiceID string `mapstructure:"service_id"`
	Version   int    `mapstructure:"version"`

	Name      string      `mapstructure:"name"`
	ID        string      `mapstructure:"id"`
	Priority  int         `mapstructure:"priority"`
	Dynamic   int         `mapstructure:"dynamic"`
	Content   string      `mapstructure:"content"`
	Type      SnippetType `mapstructure:"type"`
	CreatedAt *time.Time  `mapstructure:"created_at"`
	UpdatedAt *time.Time  `mapstructure:"updated_at"`
	DeletedAt *time.Time  `mapstructure:"deleted_at"`
}

// CreateSnippetInput is the input for CreateSnippet
type CreateSnippetInput struct {
	// Service is the ID of the Service to add the snippet to.
	Service string

	// Version is the editable version of the service
	Version int

	// Name is the name for the snippet.
	Name string `form:"name"`

	// Priority determines the ordering for multiple snippets. Lower numbers execute first.
	Priority int `form:"priority"`

	// Dynamic sets the snippet version to regular (0) or dynamic (1).
	Dynamic int `form:"dynamic"`

	// Content is the VCL code that specifies exactly what the snippet does.
	Content string `form:"content"`

	// Type is the location in generated VCL where the snippet should be placed.
	Type SnippetType `form:"type"`
}

// CreateSnippet creates a new snippet or dynamic snippet on a unlocked version
func (c *Client) CreateSnippet(i *CreateSnippetInput) (*Snippet, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	if i.Dynamic == 0 && i.Content == "" {
		return nil, ErrMissingContent
	}

	path := fmt.Sprintf("/service/%s/version/%d/snippet", i.Service, i.Version)
	resp, err := c.PostForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var snippet *Snippet
	if err := decodeJSON(&snippet, resp.Body); err != nil {
		return nil, err
	}
	return snippet, err
}

// UpdateSnippetInput is the input for UpdateSnippet
type UpdateSnippetInput struct {
	// Service is the ID of the Service to add the snippet to.
	Service string

	// Version is the editable version of the service
	Version int

	Name string

	// Name is the name for the snippet.
	NewName string `form:"name"`

	// Priority determines the ordering for multiple snippets. Lower numbers execute first.
	Priority int `form:"priority"`

	// Dynamic sets the snippet version to regular (0) or dynamic (1).
	Dynamic int `form:"dynamic"`

	// Content is the VCL code that specifies exactly what the snippet does.
	Content string `form:"content"`

	// Type is the location in generated VCL where the snippet should be placed.
	Type SnippetType `form:"type"`
}

// UpdateSnippet updates a snippet on a unlocked version
func (c *Client) UpdateSnippet(i *UpdateSnippetInput) (*Snippet, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	if i.Dynamic == 0 && i.Content == "" {
		return nil, ErrMissingContent
	}

	path := fmt.Sprintf("/service/%s/version/%d/snippet/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.PutForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var snippet *Snippet
	if err := decodeJSON(&snippet, resp.Body); err != nil {
		return nil, err
	}
	return snippet, err
}

// DynamicSnippet is the object returned when updating or retrieving a Dynamic Snippet
type DynamicSnippet struct {
	ServiceID string `mapstructure:"service_id"`
	ID        string `mapstructure:"snippet_id"`

	Content   string     `mapstructure:"content"`
	CreatedAt *time.Time `mapstructure:"created_at"`
	UpdatedAt *time.Time `mapstructure:"updated_at"`
}

// UpdateDynamicSnippetInput is the input for UpdateDynamicSnippet
type UpdateDynamicSnippetInput struct {
	// Service is the ID of the Service to add the snippet to.
	Service string

	// ID is the ID of the Snippet to modify
	ID string

	// Content is the VCL code that specifies exactly what the snippet does.
	Content string `form:"content"`
}

// UpdateDynamicSnippet replaces the content of a Dynamic Snippet
func (c *Client) UpdateDynamicSnippet(i *UpdateDynamicSnippetInput) (*DynamicSnippet, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/service/%s/snippet/%s", i.Service, i.ID)
	resp, err := c.PutForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var updateSnippet *DynamicSnippet
	if err := decodeJSON(&updateSnippet, resp.Body); err != nil {
		return nil, err
	}
	return updateSnippet, err
}

type DeleteSnippetInput struct {
	// Service is the ID of the Service to add the snippet to.
	Service string

	// Name is the Name of the Snippet to Delete
	Name string

	// Version is the editable version of the service
	Version int
}

func (c *Client) DeleteSnippet(i *DeleteSnippetInput) error {
	if i.Service == "" {
		return ErrMissingService
	}

	if i.Version == 0 {
		return ErrMissingVersion
	}

	if i.Name == "" {
		return ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/snippet/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.Delete(path, nil)
	if err != nil {
		return err
	}

	var r *statusResp
	if err := decodeJSON(&r, resp.Body); err != nil {
		return err
	}
	if !r.Ok() {
		return fmt.Errorf("Not Ok")
	}
	return nil
}

// ListSnippetsInput is used as input to the ListSnippets function.
type ListSnippetsInput struct {
	// Service is the ID of the service (required).
	Service string

	// Version is the specific configuration version (required).
	Version int
}

// snippetsByName is a sortable list of Snippets.
type snippetsByName []*Snippet

// Len, Swap, and Less implement the sortable interface.
func (s snippetsByName) Len() int      { return len(s) }
func (s snippetsByName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s snippetsByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// ListSnippets returns the list of Snippets for the configuration version. Content is not displayed for Dynmanic Snippets due to them being
// versionless, use the GetDynamicSnippet function to show current content.
func (c *Client) ListSnippets(i *ListSnippetsInput) ([]*Snippet, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/snippet", i.Service, i.Version)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var snippets []*Snippet
	if err := decodeJSON(&snippets, resp.Body); err != nil {
		return nil, err
	}
	sort.Stable(snippetsByName(snippets))
	return snippets, nil
}

// GetSnippetInput is used as input to the GetSnippet function.
type GetSnippetInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the Snippet to fetch.
	Name string
}

// GetSnippet gets the Snippet configuration with the given parameters. Dynamic Snippets will not show content due to them
// being versionless, use GetDynamicSnippet to see content.
func (c *Client) GetSnippet(i *GetSnippetInput) (*Snippet, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/snippet/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var snippet *Snippet
	if err := decodeJSON(&snippet, resp.Body); err != nil {
		return nil, err
	}
	return snippet, nil
}

// GetDynamicSnippetInput is used as input to the GetDynamicSnippet function.
type GetDynamicSnippetInput struct {
	// Service is the ID of the service.
	Service string

	// ID is the ID of the Snippet to fetch.
	ID string
}

// GetDynamicSnippet gets the Snippet configuration with the given parameters. This will show the current content
// associated with a Dynamic Snippet.
func (c *Client) GetDynamicSnippet(i *GetDynamicSnippetInput) (*DynamicSnippet, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/service/%s/snippet/%s", i.Service, i.ID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var snippet *DynamicSnippet
	if err := decodeJSON(&snippet, resp.Body); err != nil {
		return nil, err
	}
	return snippet, nil
}
