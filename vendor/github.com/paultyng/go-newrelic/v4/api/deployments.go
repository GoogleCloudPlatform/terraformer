package api

import (
	"fmt"
	"net/url"
)

// ListDeployments returns deployments by newrelic applicationID.
func (c *Client) ListDeployments(id int) ([]Deployment, error) {
	deployments := []Deployment{}
	reqURL, err := url.Parse(fmt.Sprintf("/applications/%v/deployments.json", id))
	if err != nil {
		return nil, err
	}
	nextPath := reqURL.String()

	for nextPath != "" {
		resp := struct {
			Deployments []Deployment `json:"deployments,omitempty"`
		}{}

		nextPath, err = c.Do("GET", nextPath, nil, &resp)
		if err != nil {
			return nil, err
		}

		deployments = append(deployments, resp.Deployments...)
	}

	return deployments, nil

}

// CreateDeployment creates a deployment for an application.
func (c *Client) CreateDeployment(applicationID int, deployment Deployment) (*Deployment, error) {
	req := struct {
		Deployment Deployment `json:"deployment"`
	}{
		Deployment: deployment,
	}

	resp := struct {
		Deployment Deployment `json:"deployment,omitempty"`
	}{}

	u := &url.URL{Path: fmt.Sprintf("/applications/%v/deployments.json", applicationID)}
	_, err := c.Do("POST", u.String(), req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Deployment, nil
}

// DeleteDeployment deletes an application deployment from an application.
func (c *Client) DeleteDeployment(applicationID, deploymentID int) error {
	u := &url.URL{Path: fmt.Sprintf("/applications/%v/deployments/%v.json", applicationID, deploymentID)}
	_, err := c.Do("DELETE", u.String(), nil, nil)
	return err
}
