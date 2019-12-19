package fastly

import (
	"fmt"
	"net/url"
	"sort"
	"time"
)

type ACL struct {
	ServiceID string `mapstructure:"service_id"`
	Version   int    `mapstructure:"version"`

	Name      string     `mapstructure:"name"`
	ID        string     `mapstructure:"id"`
	CreatedAt *time.Time `mapstructure:"created_at"`
	UpdatedAt *time.Time `mapstructure:"updated_at"`
	DeletedAt *time.Time `mapstructure:"deleted_at"`
}

// ACLsByName is a sortable list of ACLs.
type ACLsByName []*ACL

// Len, Swap, and Less implement the sortable interface.
func (s ACLsByName) Len() int      { return len(s) }
func (s ACLsByName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ACLsByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// ListACLsInput is used as input to the ListACLs function.
type ListACLsInput struct {
	// Service is the ID of the service (required).
	Service string

	// Version is the specific configuration version (required).
	Version int
}

// ListACLs returns the list of ACLs for the configuration version.
func (c *Client) ListACLs(i *ListACLsInput) ([]*ACL, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/acl", i.Service, i.Version)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var as []*ACL
	if err := decodeJSON(&as, resp.Body); err != nil {
		return nil, err
	}
	sort.Stable(ACLsByName(as))
	return as, nil
}

// CreateACLInput is used as input to the CreateACL function.
type CreateACLInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the ACL to create (required)
	Name string `form:"name"`
}

func (c *Client) CreateACL(i *CreateACLInput) (*ACL, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/acl", i.Service, i.Version)
	resp, err := c.PostForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var a *ACL
	if err := decodeJSON(&a, resp.Body); err != nil {
		return nil, err
	}
	return a, nil
}

// DeleteACLInput is the input parameter to DeleteACL function.
type DeleteACLInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the ACL to delete (required).
	Name string
}

// DeleteACL deletes the given ACL version.
func (c *Client) DeleteACL(i *DeleteACLInput) error {
	if i.Service == "" {
		return ErrMissingService
	}

	if i.Version == 0 {
		return ErrMissingVersion
	}

	if i.Name == "" {
		return ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/acl/%s", i.Service, i.Version, url.PathEscape(i.Name))
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

// GetACLInput is the input parameter to GetACL function.
type GetACLInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the ACL to get (required).
	Name string
}

// GetACL gets the ACL configuration with the given parameters.
func (c *Client) GetACL(i *GetACLInput) (*ACL, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/acl/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var a *ACL
	if err := decodeJSON(&a, resp.Body); err != nil {
		return nil, err
	}
	return a, nil
}

// UpdateACLInput is the input parameter to UpdateACL function.
type UpdateACLInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the ACL to update (required).
	Name string

	// NewName is the new name of the ACL to update (required).
	NewName string `form:"name"`
}

// UpdateACL updates the name of the ACL with the given parameters.
func (c *Client) UpdateACL(i *UpdateACLInput) (*ACL, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	if i.NewName == "" {
		return nil, ErrMissingNewName
	}

	path := fmt.Sprintf("/service/%s/version/%d/acl/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.PutForm(path, i, nil)

	if err != nil {
		return nil, err
	}

	var a *ACL
	if err := decodeJSON(&a, resp.Body); err != nil {
		return nil, err
	}

	return a, nil
}
