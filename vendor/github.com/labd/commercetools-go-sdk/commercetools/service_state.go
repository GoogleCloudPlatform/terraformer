// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// StateURLPath is the commercetools API path.
const StateURLPath = "states"

// StateCreate creates a new instance of type State
func (client *Client) StateCreate(draft *StateDraft) (result *State, err error) {
	err = client.Create(StateURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateQuery allows querying for type State
func (client *Client) StateQuery(input *QueryInput) (result *StatePagedQueryResponse, err error) {
	err = client.Query(StateURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateDeleteWithID for type State
func (client *Client) StateDeleteWithID(ID string, version int) (result *State, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("states/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateGetWithID for type State
func (client *Client) StateGetWithID(ID string) (result *State, err error) {
	err = client.Get(strings.Replace("states/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateUpdateWithIDInput is input for function StateUpdateWithID
type StateUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []StateUpdateAction
}

// StateUpdateWithID for type State
func (client *Client) StateUpdateWithID(input *StateUpdateWithIDInput) (result *State, err error) {
	err = client.Update(strings.Replace("states/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
