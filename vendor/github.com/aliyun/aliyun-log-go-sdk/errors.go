package sls

import (
	"encoding/json"
)

// BadResponseError : special sls error, not valid json format
type BadResponseError struct {
	RespBody   string
	RespHeader map[string][]string
	HTTPCode   int
}

func (e BadResponseError) String() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func (e BadResponseError) Error() string {
	return e.String()
}

// NewBadResponseError ...
func NewBadResponseError(body string, header map[string][]string, httpCode int) *BadResponseError {
	return &BadResponseError{
		RespBody:   body,
		RespHeader: header,
		HTTPCode:   httpCode,
	}
}

// mockErrorRetry : for mock the error retry logic
type mockErrorRetry struct {
	Err      Error
	RetryCnt int // RetryCnt-- after each retry. When RetryCnt > 0, return Err, else return nil, if set it BigUint, it equivalents to always failing.
}

func (e mockErrorRetry) Error() string {
	return e.Err.String()
}
