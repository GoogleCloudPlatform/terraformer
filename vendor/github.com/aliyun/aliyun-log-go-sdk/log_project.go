package sls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/golang/glog"
)

const (
	httpScheme  = "http://"
	httpsScheme = "https://"
	ipRegexStr  = `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}.*`
)

var (
	ipRegex = regexp.MustCompile(ipRegexStr)
)

// this file is deprecated and no maintenance
// see client_project.go

// LogProject defines log project
type LogProject struct {
	Name           string `json:"projectName"`    // Project name
	Description    string `json:"description"`    // Project description
	Status         string `json:"status"`         // Normal
	Owner          string `json:"owner"`          // empty
	Region         string `json:"region"`         // region id, eg cn-shanghai
	CreateTime     string `json:"createTime"`     // unix time seconds, eg 1524539357
	LastModifyTime string `json:"lastModifyTime"` // unix time seconds, eg 1524539357

	Endpoint        string // IP or hostname of SLS endpoint
	AccessKeyID     string
	AccessKeySecret string
	SecurityToken   string
	UsingHTTP       bool   // default https
	UserAgent       string // default defaultLogUserAgent
	baseURL         string
	retryTimeout    time.Duration
	httpClient      *http.Client
}

// NewLogProject creates a new SLS project.
func NewLogProject(name, endpoint, accessKeyID, accessKeySecret string) (p *LogProject, err error) {
	p = &LogProject{
		Name:            name,
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		httpClient:      defaultHttpClient,
		retryTimeout:    defaultRetryTimeout,
	}
	p.parseEndpoint()
	return p, nil
}

// WithToken add token parameter
func (p *LogProject) WithToken(token string) (*LogProject, error) {
	p.SecurityToken = token
	return p, nil
}

// WithRequestTimeout with custom timeout for a request
func (p *LogProject) WithRequestTimeout(timeout time.Duration) *LogProject {
	p.httpClient = &http.Client{
		Timeout: timeout,
	}
	return p
}

// WithRetryTimeout with custom timeout for a operation
// each operation may send one or more HTTP requests in case of retry required.
func (p *LogProject) WithRetryTimeout(timeout time.Duration) *LogProject {
	p.retryTimeout = timeout
	return p
}

// ListLogStore returns all logstore names of project p.
func (p *LogProject) ListLogStore() ([]string, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}

	uri := fmt.Sprintf("/logstores")
	r, err := request(p, "GET", uri, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return nil, err
	}

	type Body struct {
		Count     int
		LogStores []string
	}
	body := &Body{}
	json.Unmarshal(buf, body)
	storeNames := body.LogStores
	return storeNames, nil
}

// GetLogStore returns logstore according by logstore name.
func (p *LogProject) GetLogStore(name string) (*LogStore, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}

	r, err := request(p, "GET", "/logstores/"+name, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return nil, err
	}

	s := &LogStore{}
	json.Unmarshal(buf, s)
	s.Name = name
	s.project = p
	return s, nil
}

// CreateLogStore creates a new logstore in SLS,
// where name is logstore name,
// and ttl is time-to-live(in day) of logs,
// and shardCnt is the number of shards,
// and autoSplit is auto split,
// and maxSplitShard is the max number of shard.
func (p *LogProject) CreateLogStore(name string, ttl, shardCnt int, autoSplit bool, maxSplitShard int) error {
	type Body struct {
		Name          string `json:"logstoreName"`
		TTL           int    `json:"ttl"`
		ShardCount    int    `json:"shardCount"`
		AutoSplit     bool   `json:"autoSplit"`
		MaxSplitShard int    `json:"maxSplitShard"`
	}
	store := &Body{
		Name:          name,
		TTL:           ttl,
		ShardCount:    shardCnt,
		AutoSplit:     autoSplit,
		MaxSplitShard: maxSplitShard,
	}
	body, err := json.Marshal(store)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}

	r, err := request(p, "POST", "/logstores", h, body)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

// CreateLogStoreV2 creates a new logstore in SLS
func (p *LogProject) CreateLogStoreV2(logstore *LogStore) error {
	body, err := json.Marshal(logstore)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}

	r, err := request(p, "POST", "/logstores", h, body)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

// DeleteLogStore deletes a logstore according by logstore name.
func (p *LogProject) DeleteLogStore(name string) (err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}

	r, err := request(p, "DELETE", "/logstores/"+name, h, nil)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

// UpdateLogStore updates a logstore according by logstore name,
// obviously we can't modify the logstore name itself.
func (p *LogProject) UpdateLogStore(name string, ttl, shardCnt int) (err error) {
	type Body struct {
		Name       string `json:"logstoreName"`
		TTL        int    `json:"ttl"`
		ShardCount int    `json:"shardCount"`
	}
	store := &Body{
		Name:       name,
		TTL:        ttl,
		ShardCount: shardCnt,
	}
	body, err := json.Marshal(store)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}
	r, err := request(p, "PUT", "/logstores/"+name, h, body)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

// UpdateLogStoreV2 updates a logstore according by logstore name
// obviously we can't modify the logstore name itself.
func (p *LogProject) UpdateLogStoreV2(logstore *LogStore) (err error) {
	body, err := json.Marshal(logstore)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}
	r, err := request(p, "PUT", "/logstores/"+logstore.Name, h, body)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

// ListMachineGroup returns machine group name list and the total number of machine groups.
// The offset starts from 0 and the size is the max number of machine groups could be returned.
func (p *LogProject) ListMachineGroup(offset, size int) (m []string, total int, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	if size <= 0 {
		size = 500
	}
	uri := fmt.Sprintf("/machinegroups?offset=%v&size=%v", offset, size)
	r, err := request(p, "GET", uri, h, nil)
	if err != nil {
		return nil, 0, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return nil, 0, err
	}

	type Body struct {
		MachineGroups []string
		Count         int
		Total         int
	}
	body := &Body{}
	json.Unmarshal(buf, body)
	m = body.MachineGroups
	total = body.Total
	return m, total, nil
}

// CheckLogstoreExist check logstore exist or not
func (p *LogProject) CheckLogstoreExist(name string) (bool, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	r, err := request(p, "GET", "/logstores/"+name, h, nil)

	if err != nil {
		if _, ok := err.(*Error); ok {
			slsErr := err.(*Error)
			if slsErr.Code == "LogStoreNotExist" {
				return false, nil
			}
			return false, slsErr
		}
		return false, err
	}
	defer r.Body.Close()
	return true, nil
}

// CheckMachineGroupExist check machine group exist or not
func (p *LogProject) CheckMachineGroupExist(name string) (bool, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	r, err := request(p, "GET", "/machinegroups/"+name, h, nil)
	if err != nil {
		if _, ok := err.(*Error); ok {
			slsErr := err.(*Error)
			if slsErr.Code == "MachineGroupNotExist" {
				return false, nil
			}
			return false, slsErr
		}
		return false, err
	}
	defer r.Body.Close()
	return true, nil
}

// GetMachineGroup retruns machine group according by machine group name.
func (p *LogProject) GetMachineGroup(name string) (m *MachineGroup, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	r, err := request(p, "GET", "/machinegroups/"+name, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return nil, err
	}

	m = new(MachineGroup)
	json.Unmarshal(buf, m)
	return m, nil
}

// CreateMachineGroup creates a new machine group in SLS.
func (p *LogProject) CreateMachineGroup(m *MachineGroup) error {
	body, err := json.Marshal(m)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}
	r, err := request(p, "POST", "/machinegroups", h, body)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

// UpdateMachineGroup updates a machine group.
func (p *LogProject) UpdateMachineGroup(m *MachineGroup) (err error) {
	body, err := json.Marshal(m)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}
	r, err := request(p, "PUT", "/machinegroups/"+m.Name, h, body)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

// DeleteMachineGroup deletes machine group according machine group name.
func (p *LogProject) DeleteMachineGroup(name string) (err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	r, err := request(p, "DELETE", "/machinegroups/"+name, h, nil)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}

	return nil
}

// ListConfig returns config names list and the total number of configs.
// The offset starts from 0 and the size is the max number of configs could be returned.
func (p *LogProject) ListConfig(offset, size int) (cfgNames []string, total int, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	if size <= 0 {
		size = 100
	}
	uri := fmt.Sprintf("/configs?offset=%v&size=%v", offset, size)
	r, err := request(p, "GET", uri, h, nil)
	if err != nil {
		return nil, 0, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return nil, 0, err
	}

	type Body struct {
		Total   int
		Configs []string
	}
	body := &Body{}
	json.Unmarshal(buf, body)
	cfgNames = body.Configs
	total = body.Total
	return cfgNames, total, nil
}

// CheckConfigExist check config exist or not
func (p *LogProject) CheckConfigExist(name string) (bool, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	r, err := request(p, "GET", "/configs/"+name, h, nil)
	if err != nil {
		if _, ok := err.(*Error); ok {
			slsErr := err.(*Error)
			if slsErr.Code == "ConfigNotExist" {
				return false, nil
			}
			return false, slsErr
		}
		return false, err
	}
	defer r.Body.Close()
	return true, nil
}

// GetConfig returns config according by config name.
func (p *LogProject) GetConfig(name string) (c *LogConfig, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	r, err := request(p, "GET", "/configs/"+name, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return nil, err
	}

	c = &LogConfig{}
	json.Unmarshal(buf, c)

	if glog.V(4) {
		glog.Info("Get logtail config, result", *c)
	}

	return c, nil
}

// UpdateConfig updates a config.
func (p *LogProject) UpdateConfig(c *LogConfig) (err error) {
	body, err := json.Marshal(c)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}
	r, err := request(p, "PUT", "/configs/"+c.Name, h, body)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

// CreateConfig creates a new config in SLS.
func (p *LogProject) CreateConfig(c *LogConfig) (err error) {
	body, err := json.Marshal(c)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}
	r, err := request(p, "POST", "/configs", h, body)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, err = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

// DeleteConfig deletes a config according by config name.
func (p *LogProject) DeleteConfig(name string) (err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	r, err := request(p, "DELETE", "/configs/"+name, h, nil)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}

	return nil
}

// GetAppliedMachineGroups returns applied machine group names list according config name.
func (p *LogProject) GetAppliedMachineGroups(confName string) (groupNames []string, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/configs/%v/machinegroups", confName)
	r, err := request(p, "GET", uri, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return nil, err
	}

	type Body struct {
		Count         int
		Machinegroups []string
	}
	body := &Body{}
	json.Unmarshal(buf, body)
	groupNames = body.Machinegroups
	return groupNames, nil
}

// GetAppliedConfigs returns applied config names list according machine group name groupName.
func (p *LogProject) GetAppliedConfigs(groupName string) (confNames []string, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/machinegroups/%v/configs", groupName)
	r, err := request(p, "GET", uri, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return nil, err
	}

	type Cfg struct {
		Count   int      `json:"count"`
		Configs []string `json:"configs"`
	}
	body := &Cfg{}
	json.Unmarshal(buf, body)
	confNames = body.Configs
	return confNames, nil
}

// ApplyConfigToMachineGroup applies config to machine group.
func (p *LogProject) ApplyConfigToMachineGroup(confName, groupName string) (err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/machinegroups/%v/configs/%v", groupName, confName)
	r, err := request(p, "PUT", uri, h, nil)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return err
	}

	return nil
}

// RemoveConfigFromMachineGroup removes config from machine group.
func (p *LogProject) RemoveConfigFromMachineGroup(confName, groupName string) (err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/machinegroups/%v/configs/%v", groupName, confName)
	r, err := request(p, "DELETE", uri, h, nil)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return err
	}

	return nil
}

func (p *LogProject) CreateEtlMeta(etlMeta *EtlMeta) (err error) {
	body, err := json.Marshal(etlMeta)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate",
	}
	r, err := request(p, "POST", fmt.Sprintf("/%v", EtlMetaURI), h, body)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, err = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

func (p *LogProject) UpdateEtlMeta(etlMeta *EtlMeta) (err error) {
	body, err := json.Marshal(etlMeta)
	if err != nil {
		return NewClientError(err)
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate",
	}
	r, err := request(p, "PUT", fmt.Sprintf("/%v", EtlMetaURI), h, body)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

func (p *LogProject) DeleteEtlMeta(etlMetaName, etlMetaKey string) (err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/%v?etlMetaName=%v&etlMetaKey=%v&etlMetaTag=%v", EtlMetaURI, etlMetaName, etlMetaKey, EtlMetaAllTagMatch)
	r, err := request(p, "DELETE", uri, h, nil)
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(body, err)
		return err
	}
	return nil
}

func (p *LogProject) listEtlMeta(etlMetaName, etlMetaKey, etlMetaTag string, offset, size int) (total int, count int, etlMeta []*EtlMeta, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/%v?offset=%v&size=%v&etlMetaName=%v&etlMetaKey=%v&etlMetaTag=%v", EtlMetaURI, offset, size, etlMetaName, etlMetaKey, etlMetaTag)
	r, err := request(p, "GET", uri, h, nil)
	if err != nil {
		return 0, 0, nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return 0, 0, nil, err
	}
	type BodyMeta struct {
		MetaName  string `json:"etlMetaName"`
		MetaKey   string `json:"etlMetaKey"`
		MetaTag   string `json:"etlMetaTag"`
		MetaValue string `json:"etlMetaValue"`
	}
	type Body struct {
		Total    int         `json:"total"`
		Count    int         `json:"count"`
		MetaList []*BodyMeta `json:"etlMetaList"`
	}
	body := &Body{}
	json.Unmarshal(buf, body)
	if body.Count == 0 || len(body.MetaList) == 0 {
		return body.Total, body.Count, nil, nil
	}
	var etlMetaList []*EtlMeta = make([]*EtlMeta, len(body.MetaList))
	for index, value := range body.MetaList {
		var metaValueMap map[string]string
		err := json.Unmarshal([]byte(value.MetaValue), &metaValueMap)
		if err != nil {
			return 0, 0, nil, NewClientError(err)
		}
		etlMetaList[index] = &EtlMeta{
			MetaName:  value.MetaName,
			MetaKey:   value.MetaKey,
			MetaTag:   value.MetaTag,
			MetaValue: metaValueMap,
		}
	}
	return body.Total, body.Count, etlMetaList, nil
}

func (p *LogProject) GetEtlMeta(etlMetaName, etlMetaKey string) (etlMeta *EtlMeta, err error) {
	_, count, etlMetaList, err := p.listEtlMeta(etlMetaName, etlMetaKey, EtlMetaAllTagMatch, 0, 1)
	if err != nil {
		return nil, err
	} else if count == 0 {
		return nil, nil
	}
	return etlMetaList[0], nil
}

func (p *LogProject) ListEtlMeta(etlMetaName string, offset, size int) (total int, count int, etlMetaList []*EtlMeta, err error) {
	return p.listEtlMeta(etlMetaName, "", EtlMetaAllTagMatch, offset, size)
}

func (p *LogProject) ListEtlMetaWithTag(etlMetaName, etlMetaTag string, offset, size int) (total int, count int, etlMetaList []*EtlMeta, err error) {
	return p.listEtlMeta(etlMetaName, "", etlMetaTag, offset, size)
}

func (p *LogProject) ListEtlMetaName(offset, size int) (total int, count int, etlMetaNameList []string, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/%v?offset=%v&size=%v", EtlMetaNameURI, offset, size)
	r, err := request(p, "GET", uri, h, nil)
	if err != nil {
		return 0, 0, nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		json.Unmarshal(buf, err)
		return 0, 0, nil, err
	}
	type Body struct {
		Total        int      `json:"total"`
		Count        int      `json:"count"`
		MetaNameList []string `json:"etlMetaNameList"`
	}
	body := &Body{}
	json.Unmarshal(buf, body)
	return body.Total, body.Count, body.MetaNameList, nil
}

func (p *LogProject) init() {
	if p.retryTimeout == time.Duration(0) {
		p.httpClient = defaultHttpClient
		p.retryTimeout = defaultRetryTimeout
		p.parseEndpoint()
	}
}

func (p *LogProject) getBaseURL() string {
	if len(p.baseURL) > 0 {
		return p.baseURL
	}
	p.parseEndpoint()
	return p.baseURL
}

func (p *LogProject) parseEndpoint() {
	scheme := httpScheme // default to http scheme
	host := p.Endpoint

	if strings.HasPrefix(p.Endpoint, httpScheme) {
		scheme = httpScheme
		host = strings.TrimPrefix(p.Endpoint, scheme)
	} else if strings.HasPrefix(p.Endpoint, httpsScheme) {
		scheme = httpsScheme
		host = strings.TrimPrefix(p.Endpoint, scheme)
	}

	if GlobalForceUsingHTTP || p.UsingHTTP {
		scheme = httpScheme
	}
	if ipRegex.MatchString(host) { // ip format
		if len(p.Name) == 0 {
			p.baseURL = fmt.Sprintf("%s%s", scheme, host)
		} else {
			p.baseURL = fmt.Sprintf("%s%s/%s", scheme, host, p.Name)
		}

	} else {
		if len(p.Name) == 0 {
			p.baseURL = fmt.Sprintf("%s%s", scheme, host)
		} else {
			p.baseURL = fmt.Sprintf("%s%s.%s", scheme, p.Name, host)
		}

	}
}
