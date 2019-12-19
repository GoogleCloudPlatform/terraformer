// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// ShoppingListURLPath is the commercetools API path.
const ShoppingListURLPath = "shopping-lists"

// ShoppingListCreate creates a new instance of type ShoppingList
func (client *Client) ShoppingListCreate(draft *ShoppingListDraft) (result *ShoppingList, err error) {
	err = client.Create(ShoppingListURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListQuery allows querying for type ShoppingList
func (client *Client) ShoppingListQuery(input *QueryInput) (result *ShoppingListPagedQueryResponse, err error) {
	err = client.Query(ShoppingListURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListDeleteWithKey for type ShoppingList
func (client *Client) ShoppingListDeleteWithKey(key string, version int, dataErasure bool) (result *ShoppingList, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("shopping-lists/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListGetWithKey for type ShoppingList
func (client *Client) ShoppingListGetWithKey(key string) (result *ShoppingList, err error) {
	err = client.Get(strings.Replace("shopping-lists/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListUpdateWithKeyInput is input for function ShoppingListUpdateWithKey
type ShoppingListUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ShoppingListUpdateAction
}

// ShoppingListUpdateWithKey for type ShoppingList
func (client *Client) ShoppingListUpdateWithKey(input *ShoppingListUpdateWithKeyInput) (result *ShoppingList, err error) {
	err = client.Update(strings.Replace("shopping-lists/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListDeleteWithID for type ShoppingList
func (client *Client) ShoppingListDeleteWithID(ID string, version int, dataErasure bool) (result *ShoppingList, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("shopping-lists/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListGetWithID for type ShoppingList
func (client *Client) ShoppingListGetWithID(ID string) (result *ShoppingList, err error) {
	err = client.Get(strings.Replace("shopping-lists/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListUpdateWithIDInput is input for function ShoppingListUpdateWithID
type ShoppingListUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ShoppingListUpdateAction
}

// ShoppingListUpdateWithID for type ShoppingList
func (client *Client) ShoppingListUpdateWithID(input *ShoppingListUpdateWithIDInput) (result *ShoppingList, err error) {
	err = client.Update(strings.Replace("shopping-lists/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
