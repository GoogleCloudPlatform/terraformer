package datahub

import (
	"encoding/json"
	"net/http"
)

type RestModel interface {
	// RequestBodyEncode encode request body base given method.
	// It returns []byte
	RequestBodyEncode(method string) ([]byte, error)

	// ResponseBodyDecode decode response body base given method
	ResponseBodyDecode(method string, body []byte) error
}

type CommonResponseResult struct {
	// StatusCode http return code
	StatusCode int

	// RequestId datahub request id return by server
	RequestId string
}

func NewCommonResponseResult(code int, header *http.Header, body []byte) (result *CommonResponseResult, err error) {
	result = &CommonResponseResult{
		StatusCode: code,
		RequestId:  header.Get("x-datahub-request-id"),
	}

	switch {
	case code >= 400:
		var datahubErr DatahubError
		err = json.Unmarshal(body, &datahubErr)
		if err == nil {
			err = NewError(code, header.Get("x-datahub-request-id"), datahubErr.Code, datahubErr.Message)
		}
	default:
		err = nil
	}
	return
}
