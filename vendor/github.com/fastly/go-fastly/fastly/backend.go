package fastly

import (
	"fmt"
	"net/url"
	"sort"
	"time"
)

// Backend represents a backend response from the Fastly API.
type Backend struct {
	ServiceID string `mapstructure:"service_id"`
	Version   int    `mapstructure:"version"`

	Name                string     `mapstructure:"name"`
	Comment             string     `mapstructure:"comment"`
	Address             string     `mapstructure:"address"`
	Port                uint       `mapstructure:"port"`
	OverrideHost        string     `mapstructure:"override_host"`
	ConnectTimeout      uint       `mapstructure:"connect_timeout"`
	MaxConn             uint       `mapstructure:"max_conn"`
	ErrorThreshold      uint       `mapstructure:"error_threshold"`
	FirstByteTimeout    uint       `mapstructure:"first_byte_timeout"`
	BetweenBytesTimeout uint       `mapstructure:"between_bytes_timeout"`
	AutoLoadbalance     bool       `mapstructure:"auto_loadbalance"`
	Weight              uint       `mapstructure:"weight"`
	RequestCondition    string     `mapstructure:"request_condition"`
	HealthCheck         string     `mapstructure:"healthcheck"`
	Hostname            string     `mapstructure:"hostname"`
	Shield              string     `mapstructure:"shield"`
	UseSSL              bool       `mapstructure:"use_ssl"`
	SSLCheckCert        bool       `mapstructure:"ssl_check_cert"`
	SSLCACert           string     `mapstructure:"ssl_ca_cert"`
	SSLClientCert       string     `mapstructure:"ssl_client_cert"`
	SSLClientKey        string     `mapstructure:"ssl_client_key"`
	SSLHostname         string     `mapstructure:"ssl_hostname"`
	SSLCertHostname     string     `mapstructure:"ssl_cert_hostname"`
	SSLSNIHostname      string     `mapstructure:"ssl_sni_hostname"`
	MinTLSVersion       string     `mapstructure:"min_tls_version"`
	MaxTLSVersion       string     `mapstructure:"max_tls_version"`
	SSLCiphers          []string   `mapstructure:"ssl_ciphers"`
	CreatedAt           *time.Time `mapstructure:"created_at"`
	UpdatedAt           *time.Time `mapstructure:"updated_at"`
	DeletedAt           *time.Time `mapstructure:"deleted_at"`
}

// backendsByName is a sortable list of backends.
type backendsByName []*Backend

// Len, Swap, and Less implement the sortable interface.
func (s backendsByName) Len() int      { return len(s) }
func (s backendsByName) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s backendsByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// ListBackendsInput is used as input to the ListBackends function.
type ListBackendsInput struct {
	// Service is the ID of the service (required).
	Service string

	// Version is the specific configuration version (required).
	Version int
}

// ListBackends returns the list of backends for the configuration version.
func (c *Client) ListBackends(i *ListBackendsInput) ([]*Backend, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/backend", i.Service, i.Version)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var bs []*Backend
	if err := decodeJSON(&bs, resp.Body); err != nil {
		return nil, err
	}
	sort.Stable(backendsByName(bs))
	return bs, nil
}

// CreateBackendInput is used as input to the CreateBackend function.
type CreateBackendInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	Name                string       `form:"name,omitempty"`
	Comment             string       `form:"comment,omitempty"`
	Address             string       `form:"address,omitempty"`
	Port                uint         `form:"port,omitempty"`
	OverrideHost        string       `form:"override_host,omitempty"`
	ConnectTimeout      uint         `form:"connect_timeout,omitempty"`
	MaxConn             uint         `form:"max_conn,omitempty"`
	ErrorThreshold      uint         `form:"error_threshold,omitempty"`
	FirstByteTimeout    uint         `form:"first_byte_timeout,omitempty"`
	BetweenBytesTimeout uint         `form:"between_bytes_timeout,omitempty"`
	AutoLoadbalance     *Compatibool `form:"auto_loadbalance,omitempty"`
	Weight              uint         `form:"weight,omitempty"`
	RequestCondition    string       `form:"request_condition,omitempty"`
	HealthCheck         string       `form:"healthcheck,omitempty"`
	Shield              string       `form:"shield,omitempty"`
	UseSSL              *Compatibool `form:"use_ssl,omitempty"`
	SSLCheckCert        *Compatibool `form:"ssl_check_cert,omitempty"`
	SSLCACert           string       `form:"ssl_ca_cert,omitempty"`
	SSLClientCert       string       `form:"ssl_client_cert,omitempty"`
	SSLClientKey        string       `form:"ssl_client_key,omitempty"`
	SSLHostname         string       `form:"ssl_hostname,omitempty"`
	SSLCertHostname     string       `form:"ssl_cert_hostname,omitempty"`
	SSLSNIHostname      string       `form:"ssl_sni_hostname,omitempty"`
	MinTLSVersion       string       `form:"min_tls_version,omitempty"`
	MaxTLSVersion       string       `form:"max_tls_version,omitempty"`
	SSLCiphers          []string     `form:"ssl_ciphers,omitempty"`
}

// CreateBackend creates a new Fastly backend.
func (c *Client) CreateBackend(i *CreateBackendInput) (*Backend, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/backend", i.Service, i.Version)
	resp, err := c.PostForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var b *Backend
	if err := decodeJSON(&b, resp.Body); err != nil {
		return nil, err
	}
	return b, nil
}

// GetBackendInput is used as input to the GetBackend function.
type GetBackendInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the backend to fetch.
	Name string
}

// GetBackend gets the backend configuration with the given parameters.
func (c *Client) GetBackend(i *GetBackendInput) (*Backend, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/backend/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var b *Backend
	if err := decodeJSON(&b, resp.Body); err != nil {
		return nil, err
	}
	return b, nil
}

// UpdateBackendInput is used as input to the UpdateBackend function.
type UpdateBackendInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the backend to update.
	Name string

	NewName             string       `form:"name,omitempty"`
	Comment             string       `form:"comment,omitempty"`
	Address             string       `form:"address,omitempty"`
	Port                uint         `form:"port,omitempty"`
	OverrideHost        string       `form:"override_host,omitempty"`
	ConnectTimeout      uint         `form:"connect_timeout,omitempty"`
	MaxConn             uint         `form:"max_conn,omitempty"`
	ErrorThreshold      uint         `form:"error_threshold,omitempty"`
	FirstByteTimeout    uint         `form:"first_byte_timeout,omitempty"`
	BetweenBytesTimeout uint         `form:"between_bytes_timeout,omitempty"`
	AutoLoadbalance     *Compatibool `form:"auto_loadbalance,omitempty"`
	Weight              uint         `form:"weight,omitempty"`
	RequestCondition    string       `form:"request_condition,omitempty"`
	HealthCheck         string       `form:"healthcheck,omitempty"`
	Shield              string       `form:"shield,omitempty"`
	UseSSL              *Compatibool `form:"use_ssl,omitempty"`
	SSLCheckCert        *Compatibool `form:"ssl_check_cert,omitempty"`
	SSLCACert           string       `form:"ssl_ca_cert,omitempty"`
	SSLClientCert       string       `form:"ssl_client_cert,omitempty"`
	SSLClientKey        string       `form:"ssl_client_key,omitempty"`
	SSLHostname         string       `form:"ssl_hostname,omitempty"`
	SSLCertHostname     string       `form:"ssl_cert_hostname,omitempty"`
	SSLSNIHostname      string       `form:"ssl_sni_hostname,omitempty"`
	MinTLSVersion       string       `form:"min_tls_version,omitempty"`
	MaxTLSVersion       string       `form:"max_tls_version,omitempty"`
	SSLCiphers          []string     `form:"ssl_ciphers,omitempty"`
}

// UpdateBackend updates a specific backend.
func (c *Client) UpdateBackend(i *UpdateBackendInput) (*Backend, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.Name == "" {
		return nil, ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/backend/%s", i.Service, i.Version, url.PathEscape(i.Name))
	resp, err := c.PutForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var b *Backend
	if err := decodeJSON(&b, resp.Body); err != nil {
		return nil, err
	}
	return b, nil
}

// DeleteBackendInput is the input parameter to DeleteBackend.
type DeleteBackendInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// Name is the name of the backend to delete (required).
	Name string
}

// DeleteBackend deletes the given backend version.
func (c *Client) DeleteBackend(i *DeleteBackendInput) error {
	if i.Service == "" {
		return ErrMissingService
	}

	if i.Version == 0 {
		return ErrMissingVersion
	}

	if i.Name == "" {
		return ErrMissingName
	}

	path := fmt.Sprintf("/service/%s/version/%d/backend/%s", i.Service, i.Version, url.PathEscape(i.Name))
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
