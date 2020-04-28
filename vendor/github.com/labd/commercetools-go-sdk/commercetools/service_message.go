// Automatically generated, do not edit

package commercetools

import "strings"

// AbstractMessageURLPath is the commercetools API path.
const AbstractMessageURLPath = "messages"

// AbstractMessageQuery allows querying for type Message
func (client *Client) AbstractMessageQuery(input *QueryInput) (result *MessagePagedQueryResponse, err error) {
	err = client.Query(AbstractMessageURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AbstractMessageGetWithID for type Message
func (client *Client) AbstractMessageGetWithID(ID string) (result *Message, err error) {
	err = client.Get(strings.Replace("messages/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
