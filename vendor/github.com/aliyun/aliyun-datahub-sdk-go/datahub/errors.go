package datahub

import (
	"fmt"
)

/*
datahub errors
*/

// Error codes
const (
	// 4xx
	Unauthorized               = "Unauthorized"
	NoPermission               = "NoPermission"
	InvalidParameter           = "InvalidParameter"
	InvalidSubscription        = "InvalidSubscription"
	ListSubscriptionOutofRange = "ListSubscriptionOutofRange"
	MalformedRecord            = "MalformedRecord"
	InvalidCursor              = "InvalidCursor"
	NoSuchProject              = "NoSuchProject"
	NoSuchTopic                = "NoSuchTopic"
	NoSuchShard                = "NoSuchShard"
	NoSuchSubscription         = "NoSuchSubscription"
	ProjectAlreadyExist        = "ProjectAlreadyExist"
	TopicAlreadyExist          = "TopicAlreadyExist"
	MethodNotAllowed           = "MethodNotAllowed"
	InvalidShardOperation      = "InvalidShardOperation"

	// 5xx
	ServiceUnavailable  = "ServiceUnavailable"
	ShardNotReady       = "ShardNotReady"
	InternalServerError = "InternalServerError"
	LimitExceeded       = "LimitExceeded"
)

// DatahubError struct
type DatahubError struct {
	StatusCode int    `json:"StatusCode"`   // Http status code
	RequestId  string `json:"RequestId"`    // Request-id to trace the request
	Code       string `json:"ErrorCode"`    // Datahub error code
	Message    string `json:"ErrorMessage"` // Error msg of the error code
}

func (err DatahubError) Error() string {
	return fmt.Sprintf("statusCode: %d, requestId: %s, errCode: %s, errMsg: %s",
		err.StatusCode, err.RequestId, err.Code, err.Message)
}

func NewError(statusCode int, requestId, code, message string) error {
	return DatahubError{
		StatusCode: statusCode,
		RequestId:  requestId,
		Code:       code,
		Message:    message,
	}
}
