package fastly

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
)

// ErrMissingService is an error that is returned when an input struct requires
// a "Service" key, but one was not set.
var ErrMissingService = errors.New("Missing required field 'Service'")

// ErrMissingStatus is an error that is returned when an input struct requires
// a "Status" key, but one was not set.
var ErrMissingStatus = errors.New("Missing required field 'Status'")

// ErrMissingTag is an error that is returned when an input struct requires
// a "Tag" key, but one was not set.
var ErrMissingTag = errors.New("Missing required field 'Tag'")

// ErrMissingVersion is an error that is returned when an input struct requires
// a "Version" key, but one was not set.
var ErrMissingVersion = errors.New("Missing required field 'Version'")

// ErrMissingContent is an error that is returned when an input struct requires a
// "Content" key, but one was not set.
var ErrMissingContent = errors.New("Missing required field 'Content'")

// ErrMissingName is an error that is returned when an input struct requires a
// "Name" key, but one was not set.
var ErrMissingName = errors.New("Missing required field 'Name'")

// ErrMissingKey is an error that is returned when an input struct requires a
// "Key" key, but one was not set.
var ErrMissingKey = errors.New("Missing required field 'Key'")

// ErrMissingURL is an error that is returned when an input struct requires a
// "URL" key, but one was not set.
var ErrMissingURL = errors.New("Missing required field 'URL'")

// ErrMissingID is an error that is returned when an input struct requires an
// "ID" key, but one was not set.
var ErrMissingID = errors.New("Missing required field 'ID'")

// ErrMissingDictionary is an error that is returned when an input struct
// requires a "Dictionary" key, but one was not set.
var ErrMissingDictionary = errors.New("Missing required field 'Dictionary'")

// ErrMissingItemKey is an error that is returned when an input struct
// requires a "ItemKey" key, but one was not set.
var ErrMissingItemKey = errors.New("Missing required field 'ItemKey'")

// ErrMissingFrom is an error that is returned when an input struct
// requires a "From" key, but one was not set.
var ErrMissingFrom = errors.New("Missing required field 'From'")

// ErrMissingTo is an error that is returned when an input struct
// requires a "To" key, but one was not set.
var ErrMissingTo = errors.New("Missing required field 'To'")

// ErrMissingDirector is an error that is returned when an input struct
// requires a "From" key, but one was not set.
var ErrMissingDirector = errors.New("Missing required field 'Director'")

// ErrMissingBackend is an error that is returned when an input struct
// requires a "Backend" key, but one was not set.
var ErrMissingBackend = errors.New("Missing required field 'Backend'")

// ErrMissingYear is an error that is returned when an input struct
// requires a "Year" key, but one was not set.
var ErrMissingYear = errors.New("Missing required field 'Year'")

// ErrMissingMonth is an error that is returned when an input struct
// requires a "Month" key, but one was not set.
var ErrMissingMonth = errors.New("Missing required field 'Month'")

// ErrMissingNewName is an erorr that is returned when an input struct
// requires a "NewName" key, but one was not set
var ErrMissingNewName = errors.New("Missing required field 'NewName'")

// ErrMissingAcl is an error that is returned when an unout struct
// required an "Acl" key, but one is not set
var ErrMissingACL = errors.New("Missing required field 'ACL'")

// ErrMissingIP is an error that is returned when an unout struct
// required an "IP" key, but one is not set
var ErrMissingIP = errors.New("Missing requried field 'IP'")

// ErrMissingEventID is an error that is returned was an input struct
// requires a "EventID" key, but one was not set
var ErrMissingEventID = errors.New("Missing required field 'EventID'")

// ErrMissingWafID is an error that is returned was an input struct
// requires a "WafID" key, but one was not set
var ErrMissingWAFID = errors.New("Missing required field 'WAFID'")

// ErrMissingOWASPID is an error that is returned was an input struct
// requires a "OWASPID" key, but one was not set
var ErrMissingOWASPID = errors.New("Missing required field 'OWASPID'")

// ErrMissingRuleID is an error that is returned was an input struct
// requires a "RuleID" key, but one was not set
var ErrMissingRuleID = errors.New("Missing required field 'RuleID'")

// ErrMissingConfigSetID is an error that is returned was an input struct
// requires a "ConfigSetID" key, but one was not set
var ErrMissingConfigSetID = errors.New("Missing required field 'ConfigSetID'")

// ErrMissingWAFList is an error that is returned was an input struct
// requires a list of WAF id's, but it is empty
var ErrMissingWAFList = errors.New("WAF slice is empty")

// ErrBatchUpdateMaximumItemsExceeded is an error that indicates that too many batch operations are being executed.
// The Fastly API specifies an maximum limit.
var ErrBatchUpdateMaximumOperationsExceeded = errors.New("batch modify maximum operations exceeded")

// Ensure HTTPError is, in fact, an error.
var _ error = (*HTTPError)(nil)

// HTTPError is a custom error type that wraps an HTTP status code with some
// helper functions.
type HTTPError struct {
	// StatusCode is the HTTP status code (2xx-5xx).
	StatusCode int

	Errors []*ErrorObject `mapstructure:"errors"`
}

// ErrorObject is a single error.
type ErrorObject struct {
	ID     string `mapstructure:"id"`
	Title  string `mapstructure:"title"`
	Detail string `mapstructure:"detail"`
	Status string `mapstructure:"status"`
	Code   string `mapstructure:"code"`

	Meta *map[string]interface{} `mapstructure:"meta"`
}

// legacyError represents the older-style errors from Fastly. It is private
// because it is automatically converted to a jsonapi error.
type legacyError struct {
	Message string `mapstructure:"msg"`
	Detail  string `mapstructure:"detail"`
}

// NewHTTPError creates a new HTTP error from the given code.
func NewHTTPError(resp *http.Response) *HTTPError {
	var e HTTPError
	e.StatusCode = resp.StatusCode

	if resp.Body == nil {
		return &e
	}

	// If this is a jsonapi response, decode it accordingly
	if resp.Header.Get("Content-Type") == jsonapi.MediaType {
		if err := decodeJSON(&e, resp.Body); err != nil {
			panic(err)
		}
	} else {
		var lerr *legacyError
		decodeJSON(&lerr, resp.Body)
		if lerr != nil {
			e.Errors = append(e.Errors, &ErrorObject{
				Title:  lerr.Message,
				Detail: lerr.Detail,
			})
		}
	}

	return &e
}

// Error implements the error interface and returns the string representing the
// error text that includes the status code and the corresponding status text.
func (e *HTTPError) Error() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "%d - %s:", e.StatusCode, http.StatusText(e.StatusCode))

	for _, e := range e.Errors {
		fmt.Fprintf(&b, "\n")

		if e.ID != "" {
			fmt.Fprintf(&b, "\n    ID:     %s", e.ID)
		}

		if e.Title != "" {
			fmt.Fprintf(&b, "\n    Title:  %s", e.Title)
		}

		if e.Detail != "" {
			fmt.Fprintf(&b, "\n    Detail: %s", e.Detail)
		}

		if e.Code != "" {
			fmt.Fprintf(&b, "\n    Code:   %s", e.Code)
		}

		if e.Meta != nil {
			fmt.Fprintf(&b, "\n    Meta:   %v", *e.Meta)
		}
	}

	return b.String()
}

// String implements the stringer interface and returns the string representing
// the string text that includes the status code and corresponding status text.
func (e *HTTPError) String() string {
	return e.Error()
}

// IsNotFound returns true if the HTTP error code is a 404, false otherwise.
func (e *HTTPError) IsNotFound() bool {
	return e.StatusCode == 404
}
