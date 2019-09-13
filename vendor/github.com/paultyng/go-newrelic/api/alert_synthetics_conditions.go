package api

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) queryAlertSyntheticsConditions(policyID int) ([]AlertSyntheticsCondition, error) {
	conditions := []AlertSyntheticsCondition{}

	reqURL, err := url.Parse("/alerts_synthetics_conditions.json")
	if err != nil {
		return nil, err
	}

	qs := reqURL.Query()
	qs.Set("policy_id", strconv.Itoa(policyID))

	reqURL.RawQuery = qs.Encode()

	nextPath := reqURL.String()

	for nextPath != "" {
		resp := struct {
			SyntheticsConditions []AlertSyntheticsCondition `json:"synthetics_conditions,omitempty"`
		}{}

		nextPath, err = c.Do("GET", nextPath, nil, &resp)
		if err != nil {
			return nil, err
		}

		for _, c := range resp.SyntheticsConditions {
			c.PolicyID = policyID
		}

		conditions = append(conditions, resp.SyntheticsConditions...)
	}

	return conditions, nil
}

// GetAlertSyntheticsCondition gets information about a Synthetics alert condition given an ID and policy ID.
func (c *Client) GetAlertSyntheticsCondition(policyID int, id int) (*AlertSyntheticsCondition, error) {
	conditions, err := c.queryAlertSyntheticsConditions(policyID)
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

// ListAlertSyntheticsConditions returns Synthetics alert conditions for the specified policy.
func (c *Client) ListAlertSyntheticsConditions(policyID int) ([]AlertSyntheticsCondition, error) {
	return c.queryAlertSyntheticsConditions(policyID)
}

// CreateAlertSyntheticsCondition creates an Synthetics alert condition given the passed configuration.
func (c *Client) CreateAlertSyntheticsCondition(condition AlertSyntheticsCondition) (*AlertSyntheticsCondition, error) {
	policyID := condition.PolicyID

	req := struct {
		Condition AlertSyntheticsCondition `json:"synthetics_condition"`
	}{
		Condition: condition,
	}

	resp := struct {
		Condition AlertSyntheticsCondition `json:"synthetics_condition,omitempty"`
	}{}

	u := &url.URL{Path: fmt.Sprintf("/alerts_synthetics_conditions/policies/%v.json", policyID)}
	_, err := c.Do("POST", u.String(), req, &resp)
	if err != nil {
		return nil, err
	}

	resp.Condition.PolicyID = policyID

	return &resp.Condition, nil
}

// UpdateAlertSyntheticsCondition updates a Synthetics alert condition with the specified changes.
func (c *Client) UpdateAlertSyntheticsCondition(condition AlertSyntheticsCondition) (*AlertSyntheticsCondition, error) {
	policyID := condition.PolicyID
	id := condition.ID

	req := struct {
		Condition AlertSyntheticsCondition `json:"synthetics_condition"`
	}{
		Condition: condition,
	}

	resp := struct {
		Condition AlertSyntheticsCondition `json:"synthetics_condition,omitempty"`
	}{}

	u := &url.URL{Path: fmt.Sprintf("/alerts_synthetics_conditions/%v.json", id)}
	_, err := c.Do("PUT", u.String(), req, &resp)
	if err != nil {
		return nil, err
	}

	resp.Condition.PolicyID = policyID

	return &resp.Condition, nil
}

// DeleteAlertSyntheticsCondition removes the Synthetics alert condition given the specified ID and policy ID.
func (c *Client) DeleteAlertSyntheticsCondition(policyID int, id int) error {
	u := &url.URL{Path: fmt.Sprintf("/alerts_synthetics_conditions/%v.json", id)}
	_, err := c.Do("DELETE", u.String(), nil, nil)
	return err
}
