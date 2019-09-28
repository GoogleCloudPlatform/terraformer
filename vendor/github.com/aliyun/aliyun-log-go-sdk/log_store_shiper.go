package sls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// CreateShipper ...
func (s *LogStore) CreateShipper(shipper *Shipper) error {
	body, err := json.Marshal(shipper)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}

	uri := fmt.Sprintf("/logstores/%s/shipper", s.Name)
	r, err := request(s.project, "POST", uri, h, body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

// UpdateShipper ...
func (s *LogStore) UpdateShipper(shipper *Shipper) error {
	body, err := json.Marshal(shipper)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}

	uri := fmt.Sprintf("/logstores/%s/shipper/%s", s.Name, shipper.ShipperName)
	r, err := request(s.project, "PUT", uri, h, body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

// DeleteShipper ...
func (s *LogStore) DeleteShipper(shipperName string) error {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/logstores/%s/shipper/%s", s.Name, shipperName)
	r, err := request(s.project, "DELETE", uri, h, nil)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

// GetShipper ...
func (s *LogStore) GetShipper(shipperName string) (*Shipper, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}

	uri := fmt.Sprintf("/logstores/%s/shipper/%s", s.Name, shipperName)
	r, err := request(s.project, "GET", uri, h, nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	shipper := &Shipper{}
	err = json.Unmarshal(buf, shipper)
	if err != nil {
		return nil, NewBadResponseError(string(buf), r.Header, r.StatusCode)
	}
	return shipper, nil
}
