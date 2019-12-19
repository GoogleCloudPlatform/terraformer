// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// CartURLPath is the commercetools API path.
const CartURLPath = "carts"

// CartCreate creates a new instance of type Cart
func (client *Client) CartCreate(draft *CartDraft) (result *Cart, err error) {
	err = client.Create(CartURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartQuery allows querying for type Cart
func (client *Client) CartQuery(input *QueryInput) (result *CartPagedQueryResponse, err error) {
	err = client.Query(CartURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDeleteWithID for type Cart
func (client *Client) CartDeleteWithID(ID string, version int, dataErasure bool) (result *Cart, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("carts/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartGetWithID for type Cart
func (client *Client) CartGetWithID(ID string) (result *Cart, err error) {
	err = client.Get(strings.Replace("carts/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartUpdateWithIDInput is input for function CartUpdateWithID
type CartUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CartUpdateAction
}

// CartUpdateWithID for type Cart
func (client *Client) CartUpdateWithID(input *CartUpdateWithIDInput) (result *Cart, err error) {
	err = client.Update(strings.Replace("carts/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
