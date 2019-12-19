// Automatically generated, do not edit

package commercetools

import "strings"

// MessageURLPath is the commercetools API path.
const MessageURLPath = "messages"

// MessageQuery allows querying for type Message
func (client *Client) MessageQuery(input *QueryInput) (result *MessagePagedQueryResponse, err error) {
	err = client.Query(MessageURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MessageGetWithID for type Message
func (client *Client) MessageGetWithID(ID string) (result *Message, err error) {
	err = client.Get(strings.Replace("messages/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
