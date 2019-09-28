package fc

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	accountPath = "/account-settings"
)

type accountSettings struct {
	AvailableAZs []string `json:"availableAZs"`
}

// GetAccountSettingsInput defines get account settings intput.
type GetAccountSettingsInput struct {
}

func NewGetAccountSettingsInput() *GetAccountSettingsInput {
	return new(GetAccountSettingsInput)
}

func (o *GetAccountSettingsInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (o *GetAccountSettingsInput) GetPath() string {
	return accountPath
}

func (o *GetAccountSettingsInput) GetHeaders() Header {
	return make(Header, 0)
}

func (o *GetAccountSettingsInput) GetPayload() interface{} {
	return nil
}

func (o *GetAccountSettingsInput) Validate() error {
	return nil
}

// GetAccountSettingsOutput defines get account settings output.
type GetAccountSettingsOutput struct {
	Header http.Header
	accountSettings
}

func (o GetAccountSettingsOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o GetAccountSettingsOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}
