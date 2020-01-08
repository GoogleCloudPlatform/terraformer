package fastly

import (
	"fmt"
	"sort"
	"time"
)

// Service represents a single service for the Fastly account.
type Service struct {
	ID            string     `mapstructure:"id"`
	Name          string     `mapstructure:"name"`
	Type          string     `mapstructure:"type"`
	Comment       string     `mapstructure:"comment"`
	CustomerID    string     `mapstructure:"customer_id"`
	CreatedAt     *time.Time `mapstructure:"created_at"`
	UpdatedAt     *time.Time `mapstructure:"updated_at"`
	DeletedAt     *time.Time `mapstructure:"deleted_at"`
	ActiveVersion uint       `mapstructure:"version"`
	Versions      []*Version `mapstructure:"versions"`
}

type ServiceDetail struct {
	ID            string     `mapstructure:"id"`
	Name          string     `mapstructure:"name"`
	Type          string     `mapstructure:"type"`
	Comment       string     `mapstructure:"comment"`
	CustomerID    string     `mapstructure:"customer_id"`
	ActiveVersion Version    `mapstructure:"active_version"`
	Version       Version    `mapstructure:"version"`
	Versions      []*Version `mapstructure:"versions"`
	CreatedAt     *time.Time `mapstructure:"created_at"`
	UpdatedAt     *time.Time `mapstructure:"updated_at"`
	DeletedAt     *time.Time `mapstructure:"deleted_at"`
}

type ServiceDomain struct {
	Locked    bool       `mapstructure:"locked"`
	Version   int64      `mapstructure:"version"`
	Name      string     `mapstructure:"name"`
	DeletedAt *time.Time `mapstructure:"deleted_at"`
	ServiceID string     `mapstructure:"service_id"`
	CreatedAt *time.Time `mapstructure:"created_at"`
	Comment   string     `mapstructure:"comment"`
	UpdatedAt *time.Time `mapstructure:"updated_at"`
}
type ServiceDomainsList []*ServiceDomain

// servicesByName is a sortable list of services.
type servicesByName []*Service

// Len, Swap, and Less implement the sortable interface.
func (s servicesByName) Len() int      { return len(s) }
func (s servicesByName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s servicesByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// ListServicesInput is used as input to the ListServices function.
type ListServicesInput struct{}

// ListServices returns the full list of services for the current account.
func (c *Client) ListServices(i *ListServicesInput) ([]*Service, error) {
	resp, err := c.Get("/service", nil)
	if err != nil {
		return nil, err
	}

	var s []*Service
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}
	sort.Stable(servicesByName(s))
	return s, nil
}

// CreateServiceInput is used as input to the CreateService function.
type CreateServiceInput struct {
	Name    string `form:"name,omitempty"`
	Type    string `form:"type,omitempty"`
	Comment string `form:"comment,omitempty"`
}

// CreateService creates a new service with the given information.
func (c *Client) CreateService(i *CreateServiceInput) (*Service, error) {
	resp, err := c.PostForm("/service", i, nil)
	if err != nil {
		return nil, err
	}

	var s *Service
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}
	return s, nil
}

// GetServiceInput is used as input to the GetService function.
type GetServiceInput struct {
	ID string
}

// GetService retrieves the service information for the service with the given
// id. If no service exists for the given id, the API returns a 400 response
// (not a 404).
func (c *Client) GetService(i *GetServiceInput) (*Service, error) {
	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/service/%s", i.ID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var s *Service
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}

	return s, nil
}

// GetService retrieves the details for the service with the given id. If no
// service exists for the given id, the API returns a 400 response (not a 404).
func (c *Client) GetServiceDetails(i *GetServiceInput) (*ServiceDetail, error) {
	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/service/%s/details", i.ID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var s *ServiceDetail
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}

	return s, nil
}

// UpdateServiceInput is used as input to the UpdateService function.
type UpdateServiceInput struct {
	ID string

	Name    string `form:"name,omitempty"`
	Comment string `form:"comment,omitempty"`
}

// UpdateService updates the service with the given input.
func (c *Client) UpdateService(i *UpdateServiceInput) (*Service, error) {
	if i.ID == "" {
		return nil, ErrMissingID
	}

	path := fmt.Sprintf("/service/%s", i.ID)
	resp, err := c.PutForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var s *Service
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}
	return s, nil
}

// DeleteServiceInput is used as input to the DeleteService function.
type DeleteServiceInput struct {
	ID string
}

// DeleteService updates the service with the given input.
func (c *Client) DeleteService(i *DeleteServiceInput) error {
	if i.ID == "" {
		return ErrMissingID
	}

	path := fmt.Sprintf("/service/%s", i.ID)
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

// SearchServiceInput is used as input to the SearchService function.
type SearchServiceInput struct {
	Name string
}

// SearchService gets a specific service by name. If no service exists by that
// name, the API returns a 400 response (not a 404).
func (c *Client) SearchService(i *SearchServiceInput) (*Service, error) {
	if i.Name == "" {
		return nil, ErrMissingName
	}

	resp, err := c.Get("/service/search", &RequestOptions{
		Params: map[string]string{
			"name": i.Name,
		},
	})
	if err != nil {
		return nil, err
	}

	var s *Service
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}

	return s, nil
}

type ListServiceDomainInput struct {
	ID string
}

// ListServiceDomains lists all domains associated with a given service
func (c *Client) ListServiceDomains(i *ListServiceDomainInput) (ServiceDomainsList, error) {
	if i.ID == "" {
		return nil, ErrMissingID
	}
	path := fmt.Sprintf("/service/%s/domain", i.ID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var ds ServiceDomainsList

	if err := decodeJSON(&ds, resp.Body); err != nil {
		return nil, err
	}

	return ds, nil
}
