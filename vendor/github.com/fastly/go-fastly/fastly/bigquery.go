package fastly

import (
	"fmt"
	"net/url"
	"sort"
	"time"
)

// BigQuery represents a BigQuery response from the Fastly API.
type BigQuery struct {
	ServiceID string `mapstructure:"service_id"`
	Version   int    `mapstructure:"version"`

	Name              string     `mapstructure:"name"`
	Format            string     `mapstructure:"format"`
	User              string     `mapstructure:"user"`
	ProjectID         string     `mapstructure:"project_id"`
	Dataset           string     `mapstructure:"dataset"`
	Table             string     `mapstructure:"table"`
	Template          string     `mapstructure:"template_suffix"`
	SecretKey         string     `mapstructure:"secret_key"`
	ResponseCondition string     `mapstructure:"response_condition"`
	Placement         string     `mapstructure:"placement"`
	FormatVersion     uint       `mapstructure:"format_version"`
	CreatedAt         *time.Time `mapstructure:"created_at"`
	UpdatedAt         *time.Time `mapstructure:"updated_at"`
	DeletedAt         *time.Time `mapstructure:"deleted_at"`
}

// bigQueriesByName is a sortable list of BigQueries.
type bigQueriesByName []*BigQuery

// Len, Swap, and Less implement the sortable interface.
func (s bigQueriesByName) Len() int      { return len(s) }
func (s bigQueriesByName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s bigQueriesByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// ListBigQueriesInput is used as input to the ListBigQueries function.
type ListBigQueriesInput struct {
	// Service is the ID of the service (required).
	Service string

	// Version is the specific configuration version (required).
	Version int
}

// ListBigQueries returns the list of BigQueries for the configuration version.
func (c *Client) ListBigQueries(i *ListBigQueriesInput) ([]*BigQuery, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/bigquery", i.Service, i.Version)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var bigQueries []*BigQuery
	if err := decodeJSON(&bigQueries, resp.Body); err != nil {
		return nil, err
	}
	sort.Stable(bigQueriesByName(bigQueries))
	return bigQueries, nil
}

// CreateBigQueryInput is used as input to the CreateBigQuery function.
type CreateBigQueryInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	Name              string `form:"name,omitempty"`
	ProjectID         string `form:"project_id,omitempty"`
	Dataset           string `form:"dataset,omitempty"`
	Table             string `form:"table,omitempty"`
	Template          string `form:"template_suffix,omitempty"`
	User              string `form:"user,omitempty"`
	SecretKey         string `form:"secret_key,omitempty"`
	Format            string `form:"format,omitempty"`
	ResponseCondition string `form:"response_condition,omitempty"`
	Placement         string `form:"placement,omitempty"`
	FormatVersion     uint   `form:"format_version,omitempty"`
}

// CreateBigQuery creates a new Fastly BigQuery.
func (c *Client) CreateBigQuery(i *CreateBigQueryInput) (*BigQuery, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/bigquery", i.Service, i.Version)
	resp, err := c.PostForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var bigQuery *BigQuery
	if err := decodeJSON(&bigQuery, resp.Body); err != nil {
		return nil, err
	}
	return bigQuery, nil
}

// GetBigQueryInput is used as input to the GetBigQuery function.
type GetBigQueryInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the BigQuery to fetch.
	Name string
}

// GetBigQuery gets the BigQuery configuration with the given parameters.
func (c *Client) GetBigQuery(i *GetBigQueryInput) (*BigQuery, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/bigquery/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var bigQuery *BigQuery
	if err := decodeJSON(&bigQuery, resp.Body); err != nil {
		return nil, err
	}
	return bigQuery, nil
}

// UpdateBigQueryInput is used as input to the UpdateBigQuery function.
type UpdateBigQueryInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the BigQuery to update.
	Name string

	NewName           string `form:"name,omitempty"`
	ProjectID         string `form:"project_id,omitempty"`
	Dataset           string `form:"dataset,omitempty"`
	Table             string `form:"table,omitempty"`
	Template          string `form:"template_suffix,omitempty"`
	User              string `form:"user,omitempty"`
	SecretKey         string `form:"secret_key,omitempty"`
	Format            string `form:"format,omitempty"`
	ResponseCondition string `form:"response_condition,omitempty"`
	Placement         string `form:"placement,omitempty"`
	FormatVersion     uint   `form:"format_version,omitempty"`
}

// UpdateBigQuery updates a specific BigQuery.
func (c *Client) UpdateBigQuery(i *UpdateBigQueryInput) (*BigQuery, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/bigquery/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.PutForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var bigQuery *BigQuery
	if err := decodeJSON(&bigQuery, resp.Body); err != nil {
		return nil, err
	}
	return bigQuery, nil
}

// DeleteBigQueryInput is the input parameter to DeleteBigQuery.
type DeleteBigQueryInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the BigQuery to delete (required).
	Name string
}

// DeleteBigQuery deletes the given BigQuery version.
func (c *Client) DeleteBigQuery(i *DeleteBigQueryInput) error {
	if i.Service == "" {
		return ErrMissingService
	}

	if i.Version == 0 {
		return ErrMissingVersion
	}

	if i.Name == "" {
		return ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/logging/bigquery/%s", i.Service, i.Version, url.PathEscape(i.Name))
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
