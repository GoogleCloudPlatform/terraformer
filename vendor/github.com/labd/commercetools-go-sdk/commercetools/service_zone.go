// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// ZoneURLPath is the commercetools API path.
const ZoneURLPath = "zones"

// ZoneCreate creates a new instance of type Zone
func (client *Client) ZoneCreate(draft *ZoneDraft) (result *Zone, err error) {
	err = client.Create(ZoneURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneQuery allows querying for type Zone
func (client *Client) ZoneQuery(input *QueryInput) (result *ZonePagedQueryResponse, err error) {
	err = client.Query(ZoneURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneDeleteWithKey for type Zone
func (client *Client) ZoneDeleteWithKey(key string, version int) (result *Zone, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("zones/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneGetWithKey for type Zone
func (client *Client) ZoneGetWithKey(key string) (result *Zone, err error) {
	err = client.Get(strings.Replace("zones/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneUpdateWithKeyInput is input for function ZoneUpdateWithKey
type ZoneUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ZoneUpdateAction
}

// ZoneUpdateWithKey for type Zone
func (client *Client) ZoneUpdateWithKey(input *ZoneUpdateWithKeyInput) (result *Zone, err error) {
	err = client.Update(strings.Replace("zones/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneDeleteWithID for type Zone
func (client *Client) ZoneDeleteWithID(ID string, version int) (result *Zone, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("zones/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneGetWithID for type Zone
func (client *Client) ZoneGetWithID(ID string) (result *Zone, err error) {
	err = client.Get(strings.Replace("zones/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneUpdateWithIDInput is input for function ZoneUpdateWithID
type ZoneUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ZoneUpdateAction
}

// ZoneUpdateWithID for type Zone
func (client *Client) ZoneUpdateWithID(input *ZoneUpdateWithIDInput) (result *Zone, err error) {
	err = client.Update(strings.Replace("zones/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
