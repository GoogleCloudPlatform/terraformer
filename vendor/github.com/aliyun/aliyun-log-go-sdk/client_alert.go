package sls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// SavedSearch ...
type SavedSearch struct {
	SavedSearchName string `json:"savedsearchName"`
	SearchQuery     string `json:"searchQuery"`
	Logstore        string `json:"logstore"`
	Topic           string `json:"topic"`
	DisplayName     string `json:"displayName"`
}

// AlertDetail ...
type AlertDetail struct {
	AlertKey   string `json:"alertKey"`
	AlertValue string `json:"alertValue"`
	Comparator string `json:"comparator"`
}

// ActionDetail ...
type ActionDetail struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	MNSParam    string `json:"param,omitempty"`
	Message     string `json:"message,omitempty"`
	Webhook     string `json:"webhook,omitempty"`
}

const (
	ActionTypeSMS      = "sms"
	ActionTypeMNS      = "mns"
	ActionTypeWebhook  = "webhook"
	ActionTypeDingtalk = "dingtalk"
	ActionTypeOthers   = "unknown"
)

// Alert ...
type Alert struct {
	AlertName       string       `json:"alertName"`
	SavedSearchName string       `json:"savedsearchName"`
	From            string       `json:"from"`
	To              string       `json:"to"`
	RoleArn         string       `json:"roleArn"`
	CheckInterval   int          `json:"checkInterval"`
	Count           int          `json:"count"`
	AlertDetail     AlertDetail  `json:"alertDetail"`
	ActionType      string       `json:"actionType"`
	ActionDetail    ActionDetail `json:"actionDetail"`
	DisplayName     string       `json:"displayName"`
}

func (c *Client) CreateSavedSearch(project string, savedSearch *SavedSearch) error {
	body, err := json.Marshal(savedSearch)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
	}

	uri := "/savedsearches"
	r, err := c.request(project, "POST", uri, h, body)
	if err != nil {
		return err
	}
	r.Body.Close()
	return nil
}

func (c *Client) UpdateSavedSearch(project string, savedSearch *SavedSearch) error {
	body, err := json.Marshal(savedSearch)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
	}

	uri := "/savedsearches/" + savedSearch.SavedSearchName
	r, err := c.request(project, "PUT", uri, h, body)
	if err != nil {
		return err
	}
	r.Body.Close()
	return nil
}

func (c *Client) DeleteSavedSearch(project string, savedSearchName string) error {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Content-Type":      "application/json",
	}

	uri := "/savedsearches/" + savedSearchName
	r, err := c.request(project, "DELETE", uri, h, nil)
	if err != nil {
		return err
	}
	r.Body.Close()
	return nil
}

func (c *Client) GetSavedSearch(project string, savedSearchName string) (*SavedSearch, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Content-Type":      "application/json",
	}

	uri := "/savedsearches/" + savedSearchName
	r, err := c.request(project, "GET", uri, h, nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	savedSearch := &SavedSearch{}
	if err = json.Unmarshal(buf, savedSearch); err != nil {
		err = NewClientError(err)
	}
	return savedSearch, err
}

func (c *Client) ListSavedSearch(project string, savedSearchName string, offset, size int) (savedSearches []string, total int, count int, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Content-Type":      "application/json",
		"savedsearchName":   savedSearchName,
		"offset":            strconv.Itoa(offset),
		"size":              strconv.Itoa(size),
	}

	uri := "/savedsearches"
	r, err := c.request(project, "GET", uri, h, nil)
	if err != nil {
		return nil, 0, 0, err
	}
	defer r.Body.Close()

	type ListSavedSearch struct {
		Total         int      `json:"total"`
		Count         int      `json:"count"`
		Savedsearches []string `json:"savedsearches"`
	}

	buf, _ := ioutil.ReadAll(r.Body)
	listSavedSearch := &ListSavedSearch{}
	if err = json.Unmarshal(buf, listSavedSearch); err != nil {
		err = NewClientError(err)
	}
	return listSavedSearch.Savedsearches, listSavedSearch.Total, listSavedSearch.Count, err
}

func (c *Client) CreateAlert(project string, alert *Alert) error {
	body, err := json.Marshal(alert)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
	}

	uri := "/alerts"
	r, err := c.request(project, "POST", uri, h, body)
	if err != nil {
		return err
	}
	r.Body.Close()
	return nil
}

func (c *Client) UpdateAlert(project string, alert *Alert) error {
	body, err := json.Marshal(alert)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
	}

	uri := "/alerts/" + alert.AlertName
	r, err := c.request(project, "PUT", uri, h, body)
	if err != nil {
		return err
	}
	r.Body.Close()
	return nil
}

func (c *Client) DeleteAlert(project string, alertName string) error {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Content-Type":      "application/json",
	}

	uri := "/alerts/" + alertName
	r, err := c.request(project, "DELETE", uri, h, nil)
	if err != nil {
		return err
	}
	r.Body.Close()
	return nil
}

func (c *Client) GetAlert(project string, alertName string) (*Alert, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Content-Type":      "application/json",
	}

	uri := "/alerts/" + alertName
	r, err := c.request(project, "GET", uri, h, nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	alert := &Alert{}
	if err = json.Unmarshal(buf, alert); err != nil {
		err = NewClientError(err)
	}
	return alert, err
}

func (c *Client) ListAlert(project string, alertName string, offset, size int) (alerts []string, total int, count int, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Content-Type":      "application/json",
		"AlertName":         alertName,
		"offset":            strconv.Itoa(offset),
		"size":              strconv.Itoa(size),
	}

	uri := "/alerts"
	r, err := c.request(project, "GET", uri, h, nil)
	if err != nil {
		return nil, 0, 0, err
	}
	defer r.Body.Close()

	type ListAlert struct {
		Total   int      `json:"total"`
		Count   int      `json:"count"`
		Alertes []string `json:"alerts"`
	}

	buf, _ := ioutil.ReadAll(r.Body)
	listAlert := &ListAlert{}
	if err = json.Unmarshal(buf, listAlert); err != nil {
		err = NewClientError(err)
	}
	return listAlert.Alertes, listAlert.Total, listAlert.Count, err
}
