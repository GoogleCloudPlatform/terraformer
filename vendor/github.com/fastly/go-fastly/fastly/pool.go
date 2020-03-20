package fastly

import (
	"fmt"
	"net/url"
	"sort"
	"time"
)

const (
	// PoolTypeRandom is a pool that does random direction.
	PoolTypeRandom PoolType = "random"

	// PoolTypeHash is a pool that does hash direction.
	PoolTypeHash PoolType = "hash"

	// PoolTypeClient ins a pool that does client direction.
	PoolTypeClient PoolType = "client"
)

// PoolType is a type of pool.
type PoolType string

// Pool represents a pool response from the Fastly API.
type Pool struct {
	ServiceID string `mapstructure:"service_id"`
	Version   int    `mapstructure:"version"`

	ID               string     `mapstructure:"id"`
	Name             string     `mapstructure:"name"`
	Comment          string     `mapstructure:"comment"`
	Shield           string     `mapstructure:"shield"`
	RequestCondition string     `mapstructure:"request_condition"`
	MaxConnDefault   uint       `mapstructure:"max_conn_default"`
	ConnectTimeout   uint       `mapstructure:"connect_timeout"`
	FirstByteTimeout uint       `mapstructure:"first_byte_timeout"`
	Quorum           uint       `mapstructure:"quorum"`
	UseTLS           bool       `mapstructure:"use_tls"`
	TLSCACert        string     `mapstructure:"tls_ca_cert"`
	TLSCiphers       string     `mapstructure:"tls_ciphers"`
	TLSClientKey     string     `mapstructure:"tls_client_key"`
	TLSClientCert    string     `mapstructure:"tls_client_cert"`
	TLSSNIHostname   string     `mapstructure:"tls_sni_hostname"`
	TLSCheckCert     bool       `mapstructure:"tls_check_cert"`
	TLSCertHostname  string     `mapstructure:"tls_cert_hostname"`
	MinTLSVersion    string     `mapstructure:"min_tls_version"`
	MaxTLSVersion    string     `mapstructure:"max_tls_version"`
	Healthcheck      string     `mapstructure:"healthcheck"`
	Type             PoolType   `mapstructure:"type"`
	OverrideHost     string     `mapstructure:"override_host"`
	CreatedAt        *time.Time `mapstructure:"created_at"`
	DeletedAt        *time.Time `mapstructure:"deleted_at"`
	UpdatedAt        *time.Time `mapstructure:"updated_at"`
}

// poolsByName is a sortable list of pools.
type poolsByName []*Pool

// Len, Swap, and Less implement the sortable interface.
func (s poolsByName) Len() int      { return len(s) }
func (s poolsByName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s poolsByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// ListPoolsInput is used as input to the ListPools function.
type ListPoolsInput struct {
	// Service is the ID of the service (required).
	Service string

	// Version is the specific configuration version (required).
	Version int
}

// ListPools lists all pools for a particular service and version.
func (c *Client) ListPools(i *ListPoolsInput) ([]*Pool, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/pool", i.Service, i.Version)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var ps []*Pool
	if err := decodeJSON(&ps, resp.Body); err != nil {
		return nil, err
	}
	sort.Stable(poolsByName(ps))
	return ps, nil
}

// CreatePoolInput is used as input to the CreatePool function.
type CreatePoolInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the pool to create (required).
	Name string `form:"name"`

	// Optional fields.
	Comment          *string      `form:"comment,omitempty"`
	Shield           *string      `form:"shield,omitempty"`
	RequestCondition *string      `form:"request_condition,omitempty"`
	MaxConnDefault   *uint        `form:"max_conn_default,omitempty"`
	ConnectTimeout   *uint        `form:"connect_timeout,omitempty"`
	FirstByteTimeout *uint        `form:"first_byte_timeout,omitempty"`
	Quorum           *uint        `form:"quorum,omitempty"`
	UseTLS           *Compatibool `form:"use_tls,omitempty"`
	TLSCACert        *string      `form:"tls_ca_cert,omitempty"`
	TLSCiphers       *string      `form:"tls_ciphers,omitempty"`
	TLSClientKey     *string      `form:"tls_client_key,omitempty"`
	TLSClientCert    *string      `form:"tls_client_cert,omitempty"`
	TLSSNIHostname   *string      `form:"tls_sni_hostname,omitempty"`
	TLSCheckCert     *Compatibool `form:"tls_check_cert,omitempty"`
	TLSCertHostname  *string      `form:"tls_cert_hostname,omitempty"`
	MinTLSVersion    *string      `form:"min_tls_version,omitempty"`
	MaxTLSVersion    *string      `form:"max_tls_version,omitempty"`
	Healthcheck      *string      `form:"healthcheck,omitempty"`
	Type             PoolType     `form:"type,omitempty"`
	OverrideHost     *string      `form:"override_host,omitempty"`
}

// CreatePool creates a pool for a particular service and version.
func (c *Client) CreatePool(i *CreatePoolInput) (*Pool, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/pool", i.Service, i.Version)
	resp, err := c.PostForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var p *Pool
	if err := decodeJSON(&p, resp.Body); err != nil {
		return nil, err
	}
	return p, nil
}

// GetPoolInput is used as input to the GetPool function.
type GetPoolInput struct {
	// Service is the ID of the service (required).
	Service string

	// Version is the specific configuration version (required).
	Version int

	// Name is the name of the pool of interest (required).
	Name string
}

// GetPool gets a single pool for a particular service and version.
func (c *Client) GetPool(i *GetPoolInput) (*Pool, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/pool/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var p *Pool
	if err := decodeJSON(&p, resp.Body); err != nil {
		return nil, err
	}
	return p, nil
}

// UpdatePoolInput is used as input to the UpdatePool function.
type UpdatePoolInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Name is the name of the pool to update. All three fields
	// are required.
	Service string
	Version int
	Name    string

	// Optional fields.
	NewName          *string      `form:"name,omitempty"`
	Comment          *string      `form:"comment,omitempty"`
	Shield           *string      `form:"shield,omitempty"`
	RequestCondition *string      `form:"request_condition,omitempty"`
	MaxConnDefault   *uint        `form:"max_conn_default,omitempty"`
	ConnectTimeout   *uint        `form:"connect_timeout,omitempty"`
	FirstByteTimeout *uint        `form:"first_byte_timeout,omitempty"`
	Quorum           *uint        `form:"quorum,omitempty"`
	UseTLS           *Compatibool `form:"use_tls,omitempty"`
	TLSCACert        *string      `form:"tls_ca_cert,omitempty"`
	TLSCiphers       *string      `form:"tls_ciphers,omitempty"`
	TLSClientKey     *string      `form:"tls_client_key,omitempty"`
	TLSClientCert    *string      `form:"tls_client_cert,omitempty"`
	TLSSNIHostname   *string      `form:"tls_sni_hostname,omitempty"`
	TLSCheckCert     *Compatibool `form:"tls_check_cert,omitempty"`
	TLSCertHostname  *string      `form:"tls_cert_hostname,omitempty"`
	MinTLSVersion    *string      `form:"min_tls_version,omitempty"`
	MaxTLSVersion    *string      `form:"max_tls_version,omitempty"`
	Healthcheck      *string      `form:"healthcheck,omitempty"`
	Type             PoolType     `form:"type,omitempty"`
	OverrideHost     *string      `form:"override_host,omitempty"`
}

// UpdatePool updates a specufic pool for a particular service and version.
func (c *Client) UpdatePool(i *UpdatePoolInput) (*Pool, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/pool/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.PutForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var p *Pool
	if err := decodeJSON(&p, resp.Body); err != nil {
		return nil, err
	}
	return p, nil
}

// DeletePoolInput is used as input to the DeletePool function.
type DeletePoolInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Name is the name of the pool to delete. All three fields
	// are required.
	Service string
	Version int
	Name    string
}

// DeletePool deletes a specific pool for a particular service and version.
func (c *Client) DeletePool(i *DeletePoolInput) error {
	if i.Service == "" {

		return ErrMissingService
	}

	if i.Version == 0 {
		return ErrMissingVersion
	}

	if i.Name == "" {
		return ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/pool/%s", i.Service, i.Version, url.PathEscape(i.Name))
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
