package api

import (
	"fmt"
	"net/url"
)

func (c *Client) queryDashboards() ([]Dashboard, error) {
	dashboards := []Dashboard{}
	reqURL, err := url.Parse("/dashboards.json")

	if err != nil {
		return nil, err
	}
	nextPath := reqURL.String()

	for nextPath != "" {
		resp := struct {
			Dashboards []Dashboard `json:"dashboards,omitempty"`
		}{}

		nextPath, err = c.Do("GET", nextPath, nil, &resp)
		if err != nil {
			return nil, err
		}

		dashboards = append(dashboards, resp.Dashboards...)
	}
	return dashboards, nil
}

// GetDashboard returns a specific dashboard for the account.
func (c *Client) GetDashboard(id int) (*Dashboard, error) {
	reqURL, err := url.Parse(fmt.Sprintf("/dashboards/%v.json", id))

	if err != nil {
		return nil, err
	}

	resp := struct {
		Dashboard Dashboard `json:"dashboard,omitempty"`
	}{}

	_, err = c.Do("GET", reqURL.String(), nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Dashboard, nil
}

// ListDashboards returns all dashboards  for the account.
func (c *Client) ListDashboards() ([]Dashboard, error) {
	return c.queryDashboards()
}

// CreateDashboard creates dashboard given the passed configuration.
func (c *Client) CreateDashboard(dashboard Dashboard) (*Dashboard, error) {
	req := struct {
		Dashboard Dashboard `json:"dashboard"`
	}{
		Dashboard: dashboard,
	}

	resp := struct {
		Dashboard Dashboard `json:"dashboard,omitempty"`
	}{}

	u := &url.URL{Path: "/dashboards.json"}
	_, err := c.Do("POST", u.String(), req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Dashboard, nil
}

// UpdateDashboard updates a dashboard given the passed configuration
func (c *Client) UpdateDashboard(dashboard Dashboard) (*Dashboard, error) {
	id := dashboard.ID

	req := struct {
		Dashboard Dashboard `json:"dashboard"`
	}{
		Dashboard: dashboard,
	}

	resp := struct {
		Dashboard Dashboard `json:"dashboard,omitempty"`
	}{}

	u := &url.URL{Path: fmt.Sprintf("/dashboards/%v.json", id)}
	_, err := c.Do("PUT", u.String(), req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Dashboard, nil
}

// DeleteDashboard deletes an existing dashboard given the passed configuration
func (c *Client) DeleteDashboard(id int) error {
	u := &url.URL{Path: fmt.Sprintf("/dashboards/%v.json", id)}
	_, err := c.Do("DELETE", u.String(), nil, nil)
	return err
}
