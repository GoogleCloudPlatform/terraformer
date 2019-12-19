package fastly

import (
	"fmt"
	"net/url"
	"sort"
	"time"
)

// Splunk represents a splunk response from the Fastly API.
type Splunk struct {
	ServiceID string `mapstructure:"service_id"`
	Version   int    `mapstructure:"version"`

	Name              string     `mapstructure:"name"`
	URL               string     `mapstructure:"url"`
	Format            string     `mapstructure:"format"`
	FormatVersion     uint       `mapstructure:"format_version"`
	ResponseCondition string     `mapstructure:"response_condition"`
	Placement         string     `mapstructure:"placement"`
	Token             string     `mapstructure:"token"`
	CreatedAt         *time.Time `mapstructure:"created_at"`
	UpdatedAt         *time.Time `mapstructure:"updated_at"`
	DeletedAt         *time.Time `mapstructure:"deleted_at"`
}

// splunkByName is a sortable list of splunks.
type splunkByName []*Splunk

// Len, Swap, and Less implement the sortable interface.
func (s splunkByName) Len() int      { return len(s) }
func (s splunkByName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s splunkByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// ListSplunksInput is used as input to the ListSplunks function.
type ListSplunksInput struct {
	// Service is the ID of the service (required).
	Service string

	// Version is the specific configuration version (required).
	Version int
}

// ListSplunks returns the list of splunks for the configuration version.
func (c *Client) ListSplunks(i *ListSplunksInput) ([]*Splunk, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/splunk", i.Service, i.Version)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var ss []*Splunk
	if err := decodeJSON(&ss, resp.Body); err != nil {
		return nil, err
	}
	sort.Stable(splunkByName(ss))
	return ss, nil
}

// CreateSplunkInput is used as input to the CreateSplunk function.
type CreateSplunkInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	Name              string `form:"name,omitempty"`
	URL               string `form:"url,omitempty"`
	Format            string `form:"format,omitempty"`
	FormatVersion     uint   `form:"format_version,omitempty"`
	ResponseCondition string `form:"response_condition,omitempty"`
	Placement         string `form:"placement,omitempty"`
	Token             string `form:"token,omitempty"`
}

// CreateSplunk creates a new Fastly splunk.
func (c *Client) CreateSplunk(i *CreateSplunkInput) (*Splunk, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/splunk", i.Service, i.Version)
	resp, err := c.PostForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var s *Splunk
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}
	return s, nil
}

// GetSplunkInput is used as input to the GetSplunk function.
type GetSplunkInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the splunk to fetch.
	Name string
}

// GetSplunk gets the splunk configuration with the given parameters.
func (c *Client) GetSplunk(i *GetSplunkInput) (*Splunk, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/splunk/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var s *Splunk
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}
	return s, nil
}

// UpdateSplunkInput is used as input to the UpdateSplunk function.
type UpdateSplunkInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the splunk to update.
	Name string

	NewName           string `form:"name,omitempty"`
	URL               string `form:"url,omitempty"`
	Format            string `form:"format,omitempty"`
	FormatVersion     uint   `form:"format_version,omitempty"`
	ResponseCondition string `form:"response_condition,omitempty"`
	Placement         string `form:"placement,omitempty"`
	Token             string `form:"token,omitempty"`
}

// UpdateSplunk updates a specific splunk.
func (c *Client) UpdateSplunk(i *UpdateSplunkInput) (*Splunk, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/splunk/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.PutForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var s *Splunk
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}
	return s, nil
}

// DeleteSplunkInput is the input parameter to DeleteSplunk.
type DeleteSplunkInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the splunk to delete (required).
	Name string
}

// DeleteSplunk deletes the given splunk version.
func (c *Client) DeleteSplunk(i *DeleteSplunkInput) error {
	if i.Service == "" {
		return ErrMissingService
	}

	if i.Version == 0 {
		return ErrMissingVersion
	}

	if i.Name == "" {
		return ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/splunk/%s", i.Service, i.Version, url.PathEscape(i.Name))
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
