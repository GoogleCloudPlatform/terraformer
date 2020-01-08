// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// PaymentURLPath is the commercetools API path.
const PaymentURLPath = "payments"

// PaymentCreate creates a new instance of type Payment
func (client *Client) PaymentCreate(draft *PaymentDraft) (result *Payment, err error) {
	err = client.Create(PaymentURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentQuery allows querying for type Payment
func (client *Client) PaymentQuery(input *QueryInput) (result *PaymentPagedQueryResponse, err error) {
	err = client.Query(PaymentURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentDeleteWithKey for type Payment
func (client *Client) PaymentDeleteWithKey(key string, version int, dataErasure bool) (result *Payment, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("payments/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentGetWithKey for type Payment
func (client *Client) PaymentGetWithKey(key string) (result *Payment, err error) {
	err = client.Get(strings.Replace("payments/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentUpdateWithKeyInput is input for function PaymentUpdateWithKey
type PaymentUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []PaymentUpdateAction
}

// PaymentUpdateWithKey for type Payment
func (client *Client) PaymentUpdateWithKey(input *PaymentUpdateWithKeyInput) (result *Payment, err error) {
	err = client.Update(strings.Replace("payments/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentDeleteWithID for type Payment
func (client *Client) PaymentDeleteWithID(ID string, version int, dataErasure bool) (result *Payment, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(strings.Replace("payments/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentGetWithID for type Payment
func (client *Client) PaymentGetWithID(ID string) (result *Payment, err error) {
	err = client.Get(strings.Replace("payments/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaymentUpdateWithIDInput is input for function PaymentUpdateWithID
type PaymentUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []PaymentUpdateAction
}

// PaymentUpdateWithID for type Payment
func (client *Client) PaymentUpdateWithID(input *PaymentUpdateWithIDInput) (result *Payment, err error) {
	err = client.Update(strings.Replace("payments/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
