// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// StoreURLPath is the commercetools API path.
const StoreURLPath = "stores"

// StoreCreate creates a new instance of type Store
func (client *Client) StoreCreate(draft *StoreDraft) (result *Store, err error) {
	err = client.Create(StoreURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreQuery allows querying for type Store
func (client *Client) StoreQuery(input *QueryInput) (result *StorePagedQueryResponse, err error) {
	err = client.Query(StoreURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreDeleteWithKey for type Store
func (client *Client) StoreDeleteWithKey(key string, version int) (result *Store, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("stores/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreGetWithKey for type Store
func (client *Client) StoreGetWithKey(key string) (result *Store, err error) {
	err = client.Get(strings.Replace("stores/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreUpdateWithKeyInput is input for function StoreUpdateWithKey
type StoreUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []StoreUpdateAction
}

// StoreUpdateWithKey for type Store
func (client *Client) StoreUpdateWithKey(input *StoreUpdateWithKeyInput) (result *Store, err error) {
	err = client.Update(strings.Replace("stores/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreDeleteWithID for type Store
func (client *Client) StoreDeleteWithID(ID string, version int) (result *Store, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("stores/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreGetWithID for type Store
func (client *Client) StoreGetWithID(ID string) (result *Store, err error) {
	err = client.Get(strings.Replace("stores/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreUpdateWithIDInput is input for function StoreUpdateWithID
type StoreUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []StoreUpdateAction
}

// StoreUpdateWithID for type Store
func (client *Client) StoreUpdateWithID(input *StoreUpdateWithIDInput) (result *Store, err error) {
	err = client.Update(strings.Replace("stores/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
