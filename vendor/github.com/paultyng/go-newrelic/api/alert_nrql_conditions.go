package api

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) queryAlertNrqlConditions(policyID int) ([]AlertNrqlCondition, error) {
	conditions := []AlertNrqlCondition{}

	reqURL, err := url.Parse("/alerts_nrql_conditions.json")
	if err != nil {
		return nil, err
	}

	qs := reqURL.Query()
	qs.Set("policy_id", strconv.Itoa(policyID))

	reqURL.RawQuery = qs.Encode()

	nextPath := reqURL.String()

	for nextPath != "" {
		resp := struct {
			NrqlConditions []AlertNrqlCondition `json:"nrql_conditions,omitempty"`
		}{}

		nextPath, err = c.Do("GET", nextPath, nil, &resp)
		if err != nil {
			return nil, err
		}

		for _, c := range resp.NrqlConditions {
			c.PolicyID = policyID
		}

		conditions = append(conditions, resp.NrqlConditions...)
	}

	return conditions, nil
}

// GetAlertNrqlCondition gets information about a NRQL alert condition given an ID and policy ID.
func (c *Client) GetAlertNrqlCondition(policyID int, id int) (*AlertNrqlCondition, error) {
	conditions, err := c.queryAlertNrqlConditions(policyID)
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

// ListAlertNrqlConditions returns NRQL alert conditions for the specified policy.
func (c *Client) ListAlertNrqlConditions(policyID int) ([]AlertNrqlCondition, error) {
	return c.queryAlertNrqlConditions(policyID)
}

// CreateAlertNrqlCondition creates an NRQL alert condition given the passed configuration.
func (c *Client) CreateAlertNrqlCondition(condition AlertNrqlCondition) (*AlertNrqlCondition, error) {
	policyID := condition.PolicyID

	req := struct {
		Condition AlertNrqlCondition `json:"nrql_condition"`
	}{
		Condition: condition,
	}

	resp := struct {
		Condition AlertNrqlCondition `json:"nrql_condition,omitempty"`
	}{}

	u := &url.URL{Path: fmt.Sprintf("/alerts_nrql_conditions/policies/%v.json", policyID)}
	_, err := c.Do("POST", u.String(), req, &resp)
	if err != nil {
		return nil, err
	}

	resp.Condition.PolicyID = policyID

	return &resp.Condition, nil
}

// UpdateAlertNrqlCondition updates a NRQL alert condition with the specified changes.
func (c *Client) UpdateAlertNrqlCondition(condition AlertNrqlCondition) (*AlertNrqlCondition, error) {
	policyID := condition.PolicyID
	id := condition.ID

	req := struct {
		Condition AlertNrqlCondition `json:"nrql_condition"`
	}{
		Condition: condition,
	}

	resp := struct {
		Condition AlertNrqlCondition `json:"nrql_condition,omitempty"`
	}{}

	u := &url.URL{Path: fmt.Sprintf("/alerts_nrql_conditions/%v.json", id)}
	_, err := c.Do("PUT", u.String(), req, &resp)
	if err != nil {
		return nil, err
	}

	resp.Condition.PolicyID = policyID

	return &resp.Condition, nil
}

// DeleteAlertNrqlCondition removes the NRQL alert condition given the specified ID and policy ID.
func (c *Client) DeleteAlertNrqlCondition(policyID int, id int) error {
	u := &url.URL{Path: fmt.Sprintf("/alerts_nrql_conditions/%v.json", id)}
	_, err := c.Do("DELETE", u.String(), nil, nil)
	return err
}
