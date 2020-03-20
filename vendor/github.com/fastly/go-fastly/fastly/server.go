package fastly

import (
	"fmt"
	"sort"
	"time"
)

// ServerType represents a server response from the Fastly API.
type Server struct {
	ServiceID string `mapstructure:"service_id"`
	PoolID    string `mapstructure:"pool_id"`
	ID        string `mapstructure:"id"`

	Address      string     `mapstructure:"address"`
	Comment      string     `mapstructure:"comment"`
	Weight       uint       `mapstructure:"weight"`
	MaxConn      uint       `mapstructure:"max_conn"`
	Port         uint       `mapstructure:"port"`
	Disabled     bool       `mapstructure:"disabled"`
	OverrideHost string     `mapstructure:"override_host"`
	CreatedAt    *time.Time `mapstructure:"created_at"`
	DeletedAt    *time.Time `mapstructure:"deleted_at"`
	UpdatedAt    *time.Time `mapstructure:"updated_at"`
}

// serversByAddress is a sortable list of servers.
type serversByAddress []*Server

// Len, Swap, and Less implement the sortable interface.
func (s serversByAddress) Len() int      { return len(s) }
func (s serversByAddress) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s serversByAddress) Less(i, j int) bool {
	return s[i].Address < s[j].Address
}

// ListServersInput is used as input to the ListServers function.
type ListServersInput struct {
	// Service is the ID of the service (required).
	Service string

	// Pool is the ID of the pool (required).
	Pool string
}

// ListServers lists all servers for a particular service and pool.
func (c *Client) ListServers(i *ListServersInput) ([]*Server, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Pool == "" {
		return nil, ErrMissingPool
	}

	path := fmt.Sprintf("/service/%s/pool/%s/servers", i.Service, i.Pool)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var ss []*Server
	if err := decodeJSON(&ss, resp.Body); err != nil {
		return nil, err
	}
	sort.Stable(serversByAddress(ss))
	return ss, nil
}

// CreateServerInput is used as input to the CreateServer function.
type CreateServerInput struct {
	// Service is the ID of the service. Pool is the ID of the pool. Both
	// fields are required.
	Service string
	Pool    string

	// Address is the hostname or IP of the origin server (required).
	Address string `form:"address"`

	// Optional fields.
	Comment      *string `form:"comment,omitempty"`
	Weight       *uint   `form:"weight,omitempty"`
	MaxConn      *uint   `form:"max_conn,omitempty"`
	Port         *uint   `form:"port,omitempty"`
	Disabled     *bool   `form:"disabled,omitempty"`
	OverrideHost *string `form:"override_host,omitempty"`
}

// CreateServer creates a single server for a particular service and pool.
// Servers are versionless resources that are associated with a Pool.
func (c *Client) CreateServer(i *CreateServerInput) (*Server, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Pool == "" {
		return nil, ErrMissingPool
	}

	if i.Address == "" {
		return nil, ErrMissingAddress
	}

	path := fmt.Sprintf("/service/%s/pool/%s/server", i.Service, i.Pool)
	resp, err := c.PostForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var s *Server
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}
	return s, nil
}

// GetServerInput is used as input to the GetServer function.
type GetServerInput struct {
	// Service is the ID of the service. Pool is the ID of the pool. Server is
	// the ID of the server in the Pool. These are required fields.
	Service string
	Pool    string
	Server  string
}

// GetServer gets a single server for a particular service and pool.
func (c *Client) GetServer(i *GetServerInput) (*Server, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Pool == "" {
		return nil, ErrMissingPool
	}

	if i.Server == "" {
		return nil, ErrMissingServer
	}

	path := fmt.Sprintf("/service/%s/pool/%s/server/%s", i.Service, i.Pool, i.Server)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var s *Server
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}
	return s, nil
}

// UpdateServerInput is used as input to the UpdateServer function.
type UpdateServerInput struct {
	// Service is the ID of the service. Pool is the ID of the pool. Server is
	// the ID of the server in the Pool. These are required fields.
	Service string
	Pool    string
	Server  string

	// Optional fields.
	Address      *string `form:"address,omitempty"`
	Comment      *string `form:"comment,omitempty"`
	Weight       *uint   `form:"weight,omitempty"`
	MaxConn      *uint   `form:"max_conn,omitempty"`
	Port         *uint   `form:"port,omitempty"`
	Disabled     *bool   `form:"disabled,omitempty"`
	OverrideHost *string `form:"override_host,omitempty"`
}

// UpdateServer updates a single server for a particular service and pool.
func (c *Client) UpdateServer(i *UpdateServerInput) (*Server, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Pool == "" {
		return nil, ErrMissingPool
	}

	if i.Server == "" {
		return nil, ErrMissingServer
	}

	path := fmt.Sprintf("/service/%s/pool/%s/server/%s", i.Service, i.Pool, i.Server)
	resp, err := c.PutForm(path, i, nil)
	if err != nil {
		return nil, err
	}

	var s *Server
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}
	return s, nil
}

// DeleteServerInput is used as input to the DeleteServer function.
type DeleteServerInput struct {
	// Service is the ID of the service. Pool is the ID of the pool. Server is
	// the ID of the server in the Pool. These are required fields.
	Service string
	Pool    string
	Server  string
}

// DeleteServer deletes a single server for a particular service and pool.
func (c *Client) DeleteServer(i *DeleteServerInput) error {
	if i.Service == "" {
		return ErrMissingService
	}

	if i.Pool == "" {
		return ErrMissingPool
	}

	if i.Server == "" {
		return ErrMissingServer
	}

	path := fmt.Sprintf("/service/%s/pool/%s/server/%s", i.Service, i.Pool, i.Server)
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
