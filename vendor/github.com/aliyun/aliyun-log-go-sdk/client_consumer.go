package sls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/golang/glog"
)

// ConsumerGroup type define
type ConsumerGroup struct {
	ConsumerGroupName string `json:"consumerGroup"`
	Timeout           int    `json:"timeout"` // timeout seconds
	InOrder           bool   `json:"order"`
}

// ConsumerGroupCheckPoint type define
type ConsumerGroupCheckPoint struct {
	ShardID    int    `json:"shard"`
	CheckPoint string `json:"checkpoint"`
	UpdateTime int    `json:"updateTime"`
	Consumer   string `json:"consumer"`
}

// CreateConsumerGroup ...
func (c *Client) CreateConsumerGroup(project, logstore string, cg ConsumerGroup) (err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Content-Type":      "application/json",
	}

	body, err := json.Marshal(cg)
	uri := fmt.Sprintf("/logstores/%v/consumergroups", logstore)
	_, err = c.request(project, "POST", uri, h, body)
	if err != nil {
		return NewClientError(err)
	}
	return nil
}

// UpdateConsumerGroup ...
func (c *Client) UpdateConsumerGroup(project, logstore string, cg ConsumerGroup) (err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Content-Type":      "application/json",
	}

	body, err := json.Marshal(cg)
	uri := fmt.Sprintf("/logstores/%v/consumergroups/%v", logstore, cg.ConsumerGroupName)
	_, err = c.request(project, "PUT", uri, h, body)
	if err != nil {
		return NewClientError(err)
	}
	return nil
}

// DeleteConsumerGroup ...
func (c *Client) DeleteConsumerGroup(project, logstore string, cgName string) (err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}

	uri := fmt.Sprintf("/logstores/%v/consumergroups/%v", logstore, cgName)
	_, err = c.request(project, "DELETE", uri, h, nil)
	if err != nil {
		return NewClientError(err)
	}
	return nil
}

// ListConsumerGroup ...
func (c *Client) ListConsumerGroup(project, logstore string) (cgList []*ConsumerGroup, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}

	uri := fmt.Sprintf("/logstores/%v/consumergroups", logstore)
	r, err := c.request(project, "GET", uri, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, NewClientError(err)
	}

	if r.StatusCode != http.StatusOK {
		errMsg := &Error{}
		err = json.Unmarshal(buf, errMsg)
		if err != nil {
			err = fmt.Errorf("failed to split shards")
			if glog.V(1) {
				dump, _ := httputil.DumpResponse(r, true)
				glog.Error(string(dump))
			}
			return nil, NewClientError(err)
		}
		return nil, errMsg
	}

	type getConsumerGroup struct {
		ConsumerGroupName string `json:"name"`    // for getConsumerGroup, this is "name"
		Timeout           int    `json:"timeout"` // timeout seconds
		InOrder           bool   `json:"order"`
	}

	var cgListRaw []*getConsumerGroup

	err = json.Unmarshal(buf, &cgListRaw)
	for _, rawCg := range cgListRaw {
		cgList = append(cgList, &ConsumerGroup{
			ConsumerGroupName: rawCg.ConsumerGroupName,
			Timeout:           rawCg.Timeout,
			InOrder:           rawCg.InOrder,
		})
	}
	return
}

// HeartBeat ...
func (c *Client) HeartBeat(project, logstore string, cgName, consumer string, heartBeatShardIDs []int) (shardIDs []int, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Content-Type":      "application/json",
	}
	if heartBeatShardIDs == nil {
		heartBeatShardIDs = []int{}
	}
	body, err := json.Marshal(heartBeatShardIDs)
	if err != nil {
		return nil, NewClientError(err)
	}
	urlVal := url.Values{}
	urlVal.Add("type", "heartbeat")
	urlVal.Add("consumer", consumer)
	uri := fmt.Sprintf("/logstores/%v/consumergroups/%v?%v", logstore, cgName, urlVal.Encode())

	r, err := c.request(project, "POST", uri, h, body)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	var shards []int
	json.Unmarshal(buf, &shards)
	for _, v := range shards {
		shardIDs = append(shardIDs, v)
	}
	return shardIDs, nil
}

// UpdateCheckpoint ...
func (c *Client) UpdateCheckpoint(project, logstore string, cgName string, consumer string, shardID int, checkpoint string, forceSuccess bool) (err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Content-Type":      "application/json",
	}
	b := map[string]interface{}{
		"shard":      shardID,
		"checkpoint": checkpoint,
	}
	body, err := json.Marshal(b)
	if err != nil {
		return NewClientError(err)
	}
	urlVal := url.Values{}
	urlVal.Add("type", "checkpoint")
	urlVal.Add("consumer", consumer)
	if forceSuccess {
		urlVal.Add("forceSuccess", "true")
	} else {
		urlVal.Add("forceSuccess", "false")
	}
	fmt.Println(urlVal.Encode())
	uri := fmt.Sprintf("/logstores/%v/consumergroups/%v?%v", logstore, cgName, urlVal.Encode())
	_, err = c.request(project, "POST", uri, h, body)
	if err != nil {
		return NewClientError(err)
	}
	return nil
}

// GetCheckpoint ...
func (c *Client) GetCheckpoint(project, logstore string, cgName string) (checkPointList []*ConsumerGroupCheckPoint, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/logstores/%v/consumergroups/%v", logstore, cgName)
	r, err := c.request(project, "GET", uri, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(buf, &checkPointList)
	return
}
