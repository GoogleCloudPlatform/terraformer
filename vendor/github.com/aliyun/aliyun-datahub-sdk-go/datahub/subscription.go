package datahub

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Subscription
type Subscription struct {
	SubId          string            `json:"SubId"`
	TopicName      string            `json:"TopicName"`
	IsOwner        bool              `json:"IsOwner"`
	Type           SubscriptionType  `json:"Type"`
	State          SubscriptionState `json:"State,omitempty"`
	Comment        string            `json:"Comment,omitempty"`
	CreateTime     uint64            `json:"CreateTime"`
	LastModifyTime uint64            `json:"LastModifyTime"`
}

func (s *Subscription) String() string {
	sBytes, _ := json.Marshal(s)
	return string(sBytes)
}

func (s *Subscription) RequestBodyEncode(method string) ([]byte, error) {
	switch method {
	case http.MethodGet, http.MethodDelete:
		return nil, nil
	case http.MethodPost:
		reqMsg := struct {
			Action  string `json:"Action"`
			Comment string `json:"Comment"`
		}{
			Action:  "create",
			Comment: s.Comment,
		}
		return json.Marshal(reqMsg)
	case http.MethodPut:
		reqMsg := struct {
			State   int    `json:"State,omitempty"`
			Comment string `json:"Comment,omitempty"`
		}{
			State:   s.State.Value(),
			Comment: s.Comment,
		}
		return json.Marshal(reqMsg)
	default:
		return nil, errors.New(fmt.Sprintf("Subscription not support method %s", method))
	}
}

func (s *Subscription) ResponseBodyDecode(method string, body []byte) error {
	switch method {
	case http.MethodGet, http.MethodPost:
		return json.Unmarshal(body, s)
	}
	return nil
}

// Subscriptions for list subscriptions
type Subscriptions struct {
	Subscriptions []Subscription `json:"Subscriptions"`
	TotalCount    int            `json:"TotalCount"`
}

func (ss *Subscriptions) String() string {
	ssBytes, _ := json.Marshal(ss)
	return string(ssBytes)
}

func (ss *Subscriptions) RequestBodyEncode(method string) ([]byte, error) {
	switch method {
	case http.MethodPost:
		reqMsg := struct {
			Action string `json:"Action"`
		}{
			Action: "list",
		}
		return json.Marshal(reqMsg)
	default:
		return nil, errors.New(fmt.Sprintf("Subscriptions not support method %s", method))
	}
}

func (ss *Subscriptions) ResponseBodyDecode(method string, body []byte) error {
	switch method {
	case http.MethodPost:
		return json.Unmarshal(body, ss)
	default:
		return errors.New(fmt.Sprintf("Subscriptions not support method %s", method))
	}
}
