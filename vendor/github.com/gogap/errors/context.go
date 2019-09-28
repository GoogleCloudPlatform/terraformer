package errors

import (
	"encoding/json"
)

type ErrorContext map[string]interface{}

func (p ErrorContext) String() string {
	if p == nil {
		return ""
	}

	if bJson, e := json.Marshal(p); e == nil {
		return string(bJson)
	}
	return ""
}
