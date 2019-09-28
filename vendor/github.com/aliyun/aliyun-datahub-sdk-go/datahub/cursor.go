package datahub

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Cursor struct {
	Id         string     `json:"Cursor"`
	Sequence   uint64     `json:"Sequence"`
	RecordTime uint64     `json:"RecordTime"`
	Type       CursorType `json:"Type"`
	SystemTime uint64     `json:"SystemTime"`
}

func (c *Cursor) String() string {
	cBytes, _ := json.Marshal(c)
	return string(cBytes)
}

func (c *Cursor) RequestBodyEncode(method string) ([]byte, error) {
	switch method {
	case http.MethodPost:
		if !ValidateCursorType(c.Type) {
			return nil, errors.New(fmt.Sprintf("cursor type %q not support", c.Type))
		}
		reqMsg := struct {
			Action     string `json:"Action"`
			SystemTime uint64 `json:"SystemTime"`
			Type       string `json:"Type"`
		}{
			Action:     "cursor",
			SystemTime: c.SystemTime,
			Type:       c.Type.String(),
		}
		return json.Marshal(reqMsg)
	default:
		return nil, errors.New(fmt.Sprintf("Cursor not support method %s", method))
	}
}

func (c *Cursor) ResponseBodyDecode(method string, body []byte) error {
	switch method {
	case http.MethodPost:
		return json.Unmarshal(body, c)
	default:
		return errors.New(fmt.Sprintf("Cursor not support method %s", method))
	}
}
