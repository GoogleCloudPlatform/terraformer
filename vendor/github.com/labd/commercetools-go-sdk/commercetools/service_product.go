// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// ProductURLPath is the commercetools API path.
const ProductURLPath = "products"

// ProductCreate creates a new instance of type Product
func (client *Client) ProductCreate(draft *ProductDraft) (result *Product, err error) {
	err = client.Create(ProductURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductQuery allows querying for type Product
func (client *Client) ProductQuery(input *QueryInput) (result *ProductPagedQueryResponse, err error) {
	err = client.Query(ProductURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDeleteWithKey for type Product
func (client *Client) ProductDeleteWithKey(key string, version int) (result *Product, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("products/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductGetWithKey for type Product
func (client *Client) ProductGetWithKey(key string) (result *Product, err error) {
	err = client.Get(strings.Replace("products/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductUpdateWithKeyInput is input for function ProductUpdateWithKey
type ProductUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ProductUpdateAction
}

// ProductUpdateWithKey for type Product
func (client *Client) ProductUpdateWithKey(input *ProductUpdateWithKeyInput) (result *Product, err error) {
	err = client.Update(strings.Replace("products/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDeleteWithID for type Product
func (client *Client) ProductDeleteWithID(ID string, version int) (result *Product, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("products/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductGetWithID for type Product
func (client *Client) ProductGetWithID(ID string) (result *Product, err error) {
	err = client.Get(strings.Replace("products/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductUpdateWithIDInput is input for function ProductUpdateWithID
type ProductUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ProductUpdateAction
}

// ProductUpdateWithID for type Product
func (client *Client) ProductUpdateWithID(input *ProductUpdateWithIDInput) (result *Product, err error) {
	err = client.Update(strings.Replace("products/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
