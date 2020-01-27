package api

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *InfraClient) queryAlertInfraConditions(policyID int) ([]AlertInfraCondition, error) {
	conditions := []AlertInfraCondition{}

	reqURL, err := url.Parse("/alerts/conditions")
	if err != nil {
		return nil, err
	}

	qs := reqURL.Query()
	qs.Set("policy_id", strconv.Itoa(policyID))

	reqURL.RawQuery = qs.Encode()

	nextPath := reqURL.String()

	for nextPath != "" {
		resp := struct {
			InfraConditions []AlertInfraCondition `json:"data,omitempty"`
		}{}

		nextPath, err = c.Do("GET", nextPath, nil, &resp)
		if err != nil {
			return nil, err
		}

		for _, c := range resp.InfraConditions {
			c.PolicyID = policyID
		}

		conditions = append(conditions, resp.InfraConditions...)
	}

	return conditions, nil
}

// GetAlertInfraCondition gets information about a Infra alert condition given an ID and policy ID.
func (c *InfraClient) GetAlertInfraCondition(policyID int, id int) (*AlertInfraCondition, error) {
	conditions, err := c.queryAlertInfraConditions(policyID)
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

// ListAlertInfraConditions returns Infra alert conditions for the specified policy.
func (c *InfraClient) ListAlertInfraConditions(policyID int) ([]AlertInfraCondition, error) {
	return c.queryAlertInfraConditions(policyID)
}

// CreateAlertInfraCondition creates an Infra alert condition given the passed configuration.
func (c *InfraClient) CreateAlertInfraCondition(condition AlertInfraCondition) (*AlertInfraCondition, error) {
	policyID := condition.PolicyID

	req := struct {
		Condition AlertInfraCondition `json:"data"`
	}{
		Condition: condition,
	}

	resp := struct {
		Condition AlertInfraCondition `json:"data,omitempty"`
	}{}

	u := &url.URL{Path: "/alerts/conditions"}
	_, err := c.Do("POST", u.String(), req, &resp)
	if err != nil {
		return nil, err
	}

	resp.Condition.PolicyID = policyID

	return &resp.Condition, nil
}

// UpdateAlertInfraCondition updates an Infra alert condition with the specified changes.
func (c *InfraClient) UpdateAlertInfraCondition(condition AlertInfraCondition) (*AlertInfraCondition, error) {
	policyID := condition.PolicyID
	id := condition.ID

	req := struct {
		Condition AlertInfraCondition `json:"data"`
	}{
		Condition: condition,
	}

	resp := struct {
		Condition AlertInfraCondition `json:"data,omitempty"`
	}{}

	u := &url.URL{Path: fmt.Sprintf("/alerts/conditions/%v", id)}
	_, err := c.Do("PUT", u.String(), req, &resp)
	if err != nil {
		return nil, err
	}

	resp.Condition.PolicyID = policyID

	return &resp.Condition, nil
}

// DeleteAlertInfraCondition removes the Infra alert condition given the specified ID and policy ID.
func (c *InfraClient) DeleteAlertInfraCondition(policyID int, id int) error {
	u := &url.URL{Path: fmt.Sprintf("/alerts/conditions/%v", id)}
	_, err := c.Do("DELETE", u.String(), nil, nil)
	return err
}
