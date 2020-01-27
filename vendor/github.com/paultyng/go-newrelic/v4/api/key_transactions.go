package api

import (
	"fmt"
	"net/url"
)

func (c *Client) queryKeyTransactions() ([]KeyTransaction, error) {
	transactions := []KeyTransaction{}

	reqURL, err := url.Parse("/key_transactions.json")
	if err != nil {
		return nil, err
	}

	nextPath := reqURL.String()

	for nextPath != "" {
		resp := struct {
			Transactions []KeyTransaction `json:"key_transactions,omitempty"`
		}{}

		nextPath, err = c.Do("GET", nextPath, nil, &resp)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, resp.Transactions...)
	}

	return transactions, nil
}

// GetKeyTransaction returns a specific key transaction by ID.
func (c *Client) GetKeyTransaction(id int) (*KeyTransaction, error) {
	reqURL, err := url.Parse(fmt.Sprintf("/key_transactions/%v.json", id))
	if err != nil {
		return nil, err
	}

	resp := struct {
		Transaction KeyTransaction `json:"key_transaction,omitempty"`
	}{}

	_, err = c.Do("GET", reqURL.String(), nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Transaction, nil
}

// ListKeyTransactions returns all key transactions for the account.
func (c *Client) ListKeyTransactions() ([]KeyTransaction, error) {
	return c.queryKeyTransactions()
}
