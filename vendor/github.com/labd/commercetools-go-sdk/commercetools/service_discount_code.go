// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// DiscountCodeURLPath is the commercetools API path.
const DiscountCodeURLPath = "discount-codes"

// DiscountCodeCreate creates a new instance of type DiscountCode
func (client *Client) DiscountCodeCreate(draft *DiscountCodeDraft) (result *DiscountCode, err error) {
	err = client.Create(DiscountCodeURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DiscountCodeQuery allows querying for type DiscountCode
func (client *Client) DiscountCodeQuery(input *QueryInput) (result *DiscountCodePagedQueryResponse, err error) {
	err = client.Query(DiscountCodeURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DiscountCodeDeleteWithID for type DiscountCode
func (client *Client) DiscountCodeDeleteWithID(ID string, version int, dataErasure bool) (result *DiscountCode, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("discount-codes/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DiscountCodeGetWithID for type DiscountCode
func (client *Client) DiscountCodeGetWithID(ID string) (result *DiscountCode, err error) {
	err = client.Get(strings.Replace("discount-codes/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DiscountCodeUpdateWithIDInput is input for function DiscountCodeUpdateWithID
type DiscountCodeUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []DiscountCodeUpdateAction
}

// DiscountCodeUpdateWithID for type DiscountCode
func (client *Client) DiscountCodeUpdateWithID(input *DiscountCodeUpdateWithIDInput) (result *DiscountCode, err error) {
	err = client.Update(strings.Replace("discount-codes/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
