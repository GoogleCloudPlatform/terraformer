package api

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) queryAlertPluginsConditions(policyID int) ([]AlertPluginsCondition, error) {
	conditions := []AlertPluginsCondition{}

	reqURL, err := url.Parse("/alerts_plugins_conditions.json")
	if err != nil {
		return nil, err
	}

	qs := reqURL.Query()
	qs.Set("policy_id", strconv.Itoa(policyID))

	reqURL.RawQuery = qs.Encode()

	nextPath := reqURL.String()

	for nextPath != "" {
		resp := struct {
			PluginConditions []AlertPluginsCondition `json:"plugins_conditions,omitempty"`
		}{}

		nextPath, err = c.Do("GET", nextPath, nil, &resp)
		if err != nil {
			return nil, err
		}

		for _, c := range resp.PluginConditions {
			c.PolicyID = policyID
		}

		conditions = append(conditions, resp.PluginConditions...)
	}

	return conditions, nil
}

// GetAlertPluginsCondition gets information about a plugin alert condition given an ID and policy ID.
func (c *Client) GetAlertPluginsCondition(policyID, id int) (*AlertPluginsCondition, error) {
	conditions, err := c.queryAlertPluginsConditions(policyID)
	if err != nil {
		return nil, err
	}

	for _, condition := range conditions {
		if condition.ID == id {
			return &condition, nil
		}
	}

	return nil, ErrNotFound
}

// ListAlertPluginsConditions returns Plugin alert conditions for the specified policy.
func (c *Client) ListAlertPluginsConditions(policyID int) ([]AlertPluginsCondition, error) {
	return c.queryAlertPluginsConditions(policyID)
}

// CreateAlertPluginsCondition creates an Plugin alert condition given the passed configuration.
func (c *Client) CreateAlertPluginsCondition(condition AlertPluginsCondition) (*AlertPluginsCondition, error) {
	policyID := condition.PolicyID

	req := struct {
		Condition AlertPluginsCondition `json:"plugins_condition"`
	}{
		Condition: condition,
	}

	resp := struct {
		Condition AlertPluginsCondition `json:"plugins_condition,omitempty"`
	}{}

	u := &url.URL{Path: fmt.Sprintf("/alerts_plugins_conditions/policies/%v.json", policyID)}
	_, err := c.Do("POST", u.String(), req, &resp)
	if err != nil {
		return nil, err
	}

	resp.Condition.PolicyID = policyID

	return &resp.Condition, nil
}

// UpdateAlertPluginsCondition updates a Plugin alert condition with the specified changes.
func (c *Client) UpdateAlertPluginsCondition(condition AlertPluginsCondition) (*AlertPluginsCondition, error) {
	policyID := condition.PolicyID
	id := condition.ID

	req := struct {
		Condition AlertPluginsCondition `json:"plugins_condition"`
	}{
		Condition: condition,
	}

	resp := struct {
		Condition AlertPluginsCondition `json:"plugins_condition,omitempty"`
	}{}

	u := &url.URL{Path: fmt.Sprintf("/alerts_plugins_conditions/%v.json", id)}
	_, err := c.Do("PUT", u.String(), req, &resp)
	if err != nil {
		return nil, err
	}

	resp.Condition.PolicyID = policyID

	return &resp.Condition, nil
}

// DeleteAlertPluginsCondition removes the Plugin alert condition given the specified ID and policy ID.
func (c *Client) DeleteAlertPluginsCondition(policyID, id int) error {
	u := &url.URL{Path: fmt.Sprintf("/alerts_plugins_conditions/%v.json", id)}
	_, err := c.Do("DELETE", u.String(), nil, nil)
	return err
}
