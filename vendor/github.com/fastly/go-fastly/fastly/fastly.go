package fastly

import (
	"bytes"
	"encoding"
)

type BatchOperation string

const (
	CreateBatchOperation BatchOperation = "create"
	UpdateBatchOperation BatchOperation = "update"
	UpsertBatchOperation BatchOperation = "upsert"
	DeleteBatchOperation BatchOperation = "delete"

	// Represents the maximum number of operations that can be sent within a single batch request.
	// This is currently not documented in the API.
	BatchModifyMaximumOperations = 1000

	// Represents the maximum number of items that can be placed within an Edge Dictionary.
	MaximumDictionarySize = 10000

	// Represents the maximum number of entries that can be placed within an ACL.
	MaximumACLSize = 10000
)

type statusResp struct {
	Status string
	Msg    string
}

func (t *statusResp) Ok() bool {
	return t.Status == "ok"
}

// Ensure Compatibool implements the proper interfaces.
var (
	_ encoding.TextMarshaler   = new(Compatibool)
	_ encoding.TextUnmarshaler = new(Compatibool)
)

// Helper function to get a pointer to bool
func CBool(b bool) *Compatibool {
	c := Compatibool(b)
	return &c
}

// Compatibool is a boolean value that marshalls to 0/1 instead of true/false
// for compatibility with Fastly's API.
type Compatibool bool

// MarshalText implements the encoding.TextMarshaler interface.
func (b Compatibool) MarshalText() ([]byte, error) {
	if b {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (b *Compatibool) UnmarshalText(t []byte) error {
	if bytes.Equal(t, []byte("1")) {
		*b = Compatibool(true)
	}
	return nil
}

// String is a helper that returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// Uint is a helper that returns a pointer to the uint value passed in.
func Uint(v uint) *uint {
	return &v
}

// Bool is a helper that returns a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}
