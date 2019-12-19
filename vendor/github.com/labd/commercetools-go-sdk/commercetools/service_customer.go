// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// CustomerURLPath is the commercetools API path.
const CustomerURLPath = "customers"

// CustomerCreate creates a new instance of type Customer
func (client *Client) CustomerCreate(draft *CustomerDraft) (result *Customer, err error) {
	err = client.Create(CustomerURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerQuery allows querying for type Customer
func (client *Client) CustomerQuery(input *QueryInput) (result *CustomerPagedQueryResponse, err error) {
	err = client.Query(CustomerURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerDeleteWithID for type Customer
func (client *Client) CustomerDeleteWithID(ID string, version int, dataErasure bool) (result *Customer, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("customers/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGetWithID for type Customer
func (client *Client) CustomerGetWithID(ID string) (result *Customer, err error) {
	err = client.Get(strings.Replace("customers/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerUpdateWithIDInput is input for function CustomerUpdateWithID
type CustomerUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CustomerUpdateAction
}

// CustomerUpdateWithID for type Customer
func (client *Client) CustomerUpdateWithID(input *CustomerUpdateWithIDInput) (result *Customer, err error) {
	err = client.Update(strings.Replace("customers/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
