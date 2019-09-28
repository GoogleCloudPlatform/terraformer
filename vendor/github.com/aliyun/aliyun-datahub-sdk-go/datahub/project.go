package datahub

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Project
type Project struct {
	CreateTime     uint64 `json:"CreateTime"`
	LastModifyTime uint64 `json:"LastModifyTime"`
	Comment        string `json:"Comment"`
}

func (p *Project) String() string {
	pbytes, _ := json.Marshal(p)
	return string(pbytes)
}

func (p *Project) RequestBodyEncode(method string) ([]byte, error) {
	switch method {
	case http.MethodGet, http.MethodDelete:
		return nil, nil
	case http.MethodPost, http.MethodPut:
		reqMsg := struct {
			Comment string `json:"Comment"`
		}{
			Comment: p.Comment,
		}
		return json.Marshal(reqMsg)
	default:
		return nil, errors.New(fmt.Sprintf("Project not support method %s", method))
	}
}

func (p *Project) ResponseBodyDecode(method string, body []byte) error {
	switch method {
	case http.MethodGet:
		return json.Unmarshal(body, p)
	}
	return nil
}

// Projects for list projects
type Projects struct {
	Names []string `json:"ProjectNames"`
}

func (ps *Projects) String() string {
	psbytes, _ := json.Marshal(ps)
	return string(psbytes)
}

func (ps *Projects) RequestBodyEncode(method string) ([]byte, error) {
	switch method {
	case http.MethodGet:
		return nil, nil
	default:
		return nil, errors.New(fmt.Sprintf("Projects not support method %s", method))
	}
}

func (ps *Projects) ResponseBodyDecode(method string, body []byte) error {
	switch method {
	case http.MethodGet:
		return json.Unmarshal(body, ps)
	default:
		return errors.New(fmt.Sprintf("Projects not support method %s", method))
	}
}
