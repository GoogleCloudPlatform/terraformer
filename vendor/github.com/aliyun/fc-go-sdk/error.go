package fc

import (
	"encoding/json"
	"errors"
)

var (
	ErrUnknownTriggerType = errors.New("unknown trigger type")
)

// ServiceError defines error from fc
type ServiceError struct {
	HTTPStatus   int    `json:"HttpStatus"`
	RequestID    string `json:"RequestId"`
	ErrorCode    string `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
}

func (e ServiceError) String() string {
	b, err := json.MarshalIndent(e, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (e ServiceError) Error() string {
	return e.String()
}
