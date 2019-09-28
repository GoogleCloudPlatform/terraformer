package sls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"github.com/pierrec/lz4"
)

// this file is deprecated and no maintenance
// see client_logstore.go

// LogStore defines LogStore struct
type LogStore struct {
	Name          string `json:"logstoreName"`
	TTL           int    `json:"ttl"`
	ShardCount    int    `json:"shardCount"`
	WebTracking   bool   `json:"enable_tracking"`
	AutoSplit     bool   `json:"autoSplit"`
	MaxSplitShard int    `json:"maxSplitShard"`
	AppendMeta    bool   `json:"appendMeta"`

	CreateTime     uint32 `json:"createTime,omitempty"`
	LastModifyTime uint32 `json:"lastModifyTime,omitempty"`

	project            *LogProject
	putLogCompressType int
}

// Shard defines shard struct
type Shard struct {
	ShardID           int    `json:"shardID"`
	Status            string `json:"status"`
	InclusiveBeginKey string `json:"inclusiveBeginKey"`
	ExclusiveBeginKey string `json:"exclusiveEndKey"`
	CreateTime        int    `json:"createTime"`
}

// NewLogStore ...
func NewLogStore(logStoreName string, project *LogProject) (*LogStore, error) {
	return &LogStore{
		Name:    logStoreName,
		project: project,
	}, nil
}

// SetPutLogCompressType set put log's compress type, default lz4
func (s *LogStore) SetPutLogCompressType(compressType int) error {
	if compressType < 0 || compressType >= Compress_Max {
		return InvalidCompressError
	}
	s.putLogCompressType = compressType
	return nil
}

// ListShards returns shard id list of this logstore.
func (s *LogStore) ListShards() (shardIDs []*Shard, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/logstores/%v/shards", s.Name)
	r, err := request(s.project, "GET", uri, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := &Error{}
		if jErr := json.Unmarshal(buf, err); jErr != nil {
			return nil, NewBadResponseError(string(buf), r.Header, r.StatusCode)
		}
		return nil, err
	}

	var shards []*Shard
	err = json.Unmarshal(buf, &shards)
	if err != nil {
		return nil, NewBadResponseError(string(buf), r.Header, r.StatusCode)
	}
	return shards, nil
}

func copyIncompressible(src, dst []byte) (int, error) {
	lLen, dn := len(src), len(dst)

	di := 0
	if lLen < 0xF {
		dst[di] = byte(lLen << 4)
	} else {
		dst[di] = 0xF0
		if di++; di == dn {
			return di, nil
		}
		lLen -= 0xF
		for ; lLen >= 0xFF; lLen -= 0xFF {
			dst[di] = 0xFF
			if di++; di == dn {
				return di, nil
			}
		}
		dst[di] = byte(lLen)
	}
	if di++; di+len(src) > dn {
		return di, nil
	}
	di += copy(dst[di:], src)
	return di, nil
}

// PutRawLog put raw log data to log service, no marshal
func (s *LogStore) PutRawLog(rawLogData []byte) (err error) {
	if len(rawLogData) == 0 {
		// empty log group
		return nil
	}

	var out []byte
	var h map[string]string
	var outLen int
	switch s.putLogCompressType {
	case Compress_LZ4:
		// Compresse body with lz4
		out = make([]byte, lz4.CompressBlockBound(len(rawLogData)))
		var hashTable [1 << 16]int
		n, err := lz4.CompressBlock(rawLogData, out, hashTable[:])
		if err != nil {
			return NewClientError(err)
		}
		// copy incompressible data as lz4 format
		if n == 0 {
			n, _ = copyIncompressible(rawLogData, out)
		}

		h = map[string]string{
			"x-log-compresstype": "lz4",
			"x-log-bodyrawsize":  strconv.Itoa(len(rawLogData)),
			"Content-Type":       "application/x-protobuf",
		}
		outLen = n
		break
	case Compress_None:
		// no compress
		out = rawLogData
		h = map[string]string{
			"x-log-bodyrawsize": strconv.Itoa(len(rawLogData)),
			"Content-Type":      "application/x-protobuf",
		}
		outLen = len(out)
	}

	uri := fmt.Sprintf("/logstores/%v", s.Name)
	r, err := request(s.project, "POST", uri, h, out[:outLen])
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		if jErr := json.Unmarshal(body, err); jErr != nil {
			return NewBadResponseError(string(body), r.Header, r.StatusCode)
		}
		return err
	}
	return nil
}

// PutLogs put logs into logstore.
// The callers should transform user logs into LogGroup.
func (s *LogStore) PutLogs(lg *LogGroup) (err error) {
	if len(lg.Logs) == 0 {
		// empty log group
		return nil
	}

	body, err := proto.Marshal(lg)
	if err != nil {
		return NewClientError(err)
	}

	var out []byte
	var h map[string]string
	var outLen int
	switch s.putLogCompressType {
	case Compress_LZ4:
		// Compresse body with lz4
		out = make([]byte, lz4.CompressBlockBound(len(body)))
		var hashTable [1 << 16]int
		n, err := lz4.CompressBlock(body, out, hashTable[:])
		if err != nil {
			return NewClientError(err)
		}
		// copy incompressible data as lz4 format
		if n == 0 {
			n, _ = copyIncompressible(body, out)
		}

		h = map[string]string{
			"x-log-compresstype": "lz4",
			"x-log-bodyrawsize":  strconv.Itoa(len(body)),
			"Content-Type":       "application/x-protobuf",
		}
		outLen = n
		break
	case Compress_None:
		// no compress
		out = body
		h = map[string]string{
			"x-log-bodyrawsize": strconv.Itoa(len(body)),
			"Content-Type":      "application/x-protobuf",
		}
		outLen = len(out)
	}

	uri := fmt.Sprintf("/logstores/%v", s.Name)
	r, err := request(s.project, "POST", uri, h, out[:outLen])
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		if jErr := json.Unmarshal(body, err); jErr != nil {
			return NewBadResponseError(string(body), r.Header, r.StatusCode)
		}
		return err
	}
	return nil
}

// PostLogStoreLogs put logs into Shard logstore by hashKey.
// The callers should transform user logs into LogGroup.
func (s *LogStore) PostLogStoreLogs(lg *LogGroup, hashKey *string) (err error) {
	if len(lg.Logs) == 0 {
		// empty log group or empty hashkey
		return nil
	}

	if hashKey == nil || *hashKey == "" {
		// empty hash call PutLogs
		return s.PutLogs(lg)
	}

	body, err := proto.Marshal(lg)
	if err != nil {
		return NewClientError(err)
	}

	var out []byte
	var h map[string]string
	var outLen int
	switch s.putLogCompressType {
	case Compress_LZ4:
		// Compresse body with lz4
		out = make([]byte, lz4.CompressBlockBound(len(body)))
		var hashTable [1 << 16]int
		n, err := lz4.CompressBlock(body, out, hashTable[:])
		if err != nil {
			return NewClientError(err)
		}
		// copy incompressible data as lz4 format
		if n == 0 {
			n, _ = copyIncompressible(body, out)
		}

		h = map[string]string{
			"x-log-compresstype": "lz4",
			"x-log-bodyrawsize":  strconv.Itoa(len(body)),
			"Content-Type":       "application/x-protobuf",
		}
		outLen = n
		break
	case Compress_None:
		// no compress
		out = body
		h = map[string]string{
			"x-log-bodyrawsize": strconv.Itoa(len(body)),
			"Content-Type":      "application/x-protobuf",
		}
		outLen = len(out)
	}

	uri := fmt.Sprintf("/logstores/%v/shards/route?key=%v", s.Name, *hashKey)
	r, err := request(s.project, "POST", uri, h, out[:outLen])
	if err != nil {
		return NewClientError(err)
	}
	defer r.Body.Close()
	body, _ = ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		if jErr := json.Unmarshal(body, err); jErr != nil {
			return NewBadResponseError(string(body), r.Header, r.StatusCode)
		}
		return err
	}
	return nil
}

// GetCursor gets log cursor of one shard specified by shardId.
// The from can be in three form: a) unix timestamp in seccond, b) "begin", c) "end".
// For more detail please read: https://help.aliyun.com/document_detail/29024.html
func (s *LogStore) GetCursor(shardID int, from string) (cursor string, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/logstores/%v/shards/%v?type=cursor&from=%v",
		s.Name, shardID, from)
	r, err := request(s.project, "GET", uri, h, nil)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	if r.StatusCode != http.StatusOK {
		errMsg := &Error{}
		err = json.Unmarshal(buf, errMsg)
		if err != nil {
			err = fmt.Errorf("failed to get cursor")
			dump, _ := httputil.DumpResponse(r, true)
			if glog.V(1) {
				glog.Error(string(dump))
			}
			return
		}
		err = fmt.Errorf("%v:%v", errMsg.Code, errMsg.Message)
		return
	}

	type Body struct {
		Cursor string
	}
	body := &Body{}

	err = json.Unmarshal(buf, body)
	if err != nil {
		return "", NewBadResponseError(string(buf), r.Header, r.StatusCode)
	}
	cursor = body.Cursor
	return cursor, nil
}

// GetLogsBytes gets logs binary data from shard specified by shardId according cursor and endCursor.
// The logGroupMaxCount is the max number of logGroup could be returned.
// The nextCursor is the next curosr can be used to read logs at next time.
func (s *LogStore) GetLogsBytes(shardID int, cursor, endCursor string,
	logGroupMaxCount int) (out []byte, nextCursor string, err error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Accept":            "application/x-protobuf",
		"Accept-Encoding":   "lz4",
	}

	uri := ""
	if endCursor == "" {
		uri = fmt.Sprintf("/logstores/%v/shards/%v?type=logs&cursor=%v&count=%v",
			s.Name, shardID, cursor, logGroupMaxCount)
	} else {
		uri = fmt.Sprintf("/logstores/%v/shards/%v?type=logs&cursor=%v&end_cursor=%v&count=%v",
			s.Name, shardID, cursor, endCursor, logGroupMaxCount)
	}

	r, err := request(s.project, "GET", uri, h, nil)
	if err != nil {
		return nil, "", err
	}
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, "", err
	}

	if r.StatusCode != http.StatusOK {
		errMsg := &Error{}
		err = json.Unmarshal(buf, errMsg)
		if err != nil {
			err = fmt.Errorf("failed to get cursor")
			dump, _ := httputil.DumpResponse(r, true)
			if glog.V(1) {
				glog.Error(string(dump))
			}
			return
		}
		err = fmt.Errorf("%v:%v", errMsg.Code, errMsg.Message)
		return
	}
	v, ok := r.Header["X-Log-Compresstype"]
	if !ok || len(v) == 0 {
		err = fmt.Errorf("can't find 'x-log-compresstype' header")
		return
	}
	if v[0] != "lz4" {
		err = fmt.Errorf("unexpected compress type:%v", v[0])
		return
	}

	v, ok = r.Header["X-Log-Cursor"]
	if !ok || len(v) == 0 {
		err = fmt.Errorf("can't find 'x-log-cursor' header")
		return
	}
	nextCursor = v[0]

	v, ok = r.Header["X-Log-Bodyrawsize"]
	if !ok || len(v) == 0 {
		err = fmt.Errorf("can't find 'x-log-bodyrawsize' header")
		return
	}
	bodyRawSize, err := strconv.Atoi(v[0])
	if err != nil {
		return nil, "", err
	}

	out = make([]byte, bodyRawSize)
	if bodyRawSize != 0 {
		len := 0
		if len, err = lz4.UncompressBlock(buf, out); err != nil || len != bodyRawSize {
			return
		}
	}
	return
}

// LogsBytesDecode decodes logs binary data returned by GetLogsBytes API
func LogsBytesDecode(data []byte) (gl *LogGroupList, err error) {

	gl = &LogGroupList{}
	err = proto.Unmarshal(data, gl)
	if err != nil {
		return nil, err
	}

	return gl, nil
}

// PullLogs gets logs from shard specified by shardId according cursor and endCursor.
// The logGroupMaxCount is the max number of logGroup could be returned.
// The nextCursor is the next cursor can be used to read logs at next time.
// @note if you want to pull logs continuous, set endCursor = ""
func (s *LogStore) PullLogs(shardID int, cursor, endCursor string,
	logGroupMaxCount int) (gl *LogGroupList, nextCursor string, err error) {

	out, nextCursor, err := s.GetLogsBytes(shardID, cursor, endCursor, logGroupMaxCount)
	if err != nil {
		return nil, "", err
	}

	gl, err = LogsBytesDecode(out)
	if err != nil {
		return nil, "", err
	}

	return gl, nextCursor, nil
}

// GetHistograms query logs with [from, to) time range
func (s *LogStore) GetHistograms(topic string, from int64, to int64, queryExp string) (*GetHistogramsResponse, error) {

	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Accept":            "application/json",
	}

	uri := fmt.Sprintf("/logstores/%v?type=histogram&topic=%v&from=%v&to=%v&query=%v", s.Name, topic, from, to, queryExp)

	r, err := request(s.project, "GET", uri, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		if jErr := json.Unmarshal(body, err); jErr != nil {
			return nil, NewBadResponseError(string(body), r.Header, r.StatusCode)
		}
		return nil, err
	}

	histograms := []SingleHistogram{}
	err = json.Unmarshal(body, &histograms)
	if err != nil {
		return nil, NewBadResponseError(string(body), r.Header, r.StatusCode)
	}

	count, err := strconv.ParseInt(r.Header[GetLogsCountHeader][0], 10, 64)
	if err != nil {
		return nil, err
	}
	getHistogramsResponse := GetHistogramsResponse{
		Progress:   r.Header[ProgressHeader][0],
		Count:      count,
		Histograms: histograms,
	}

	return &getHistogramsResponse, nil
}

// GetLogs query logs with [from, to) time range
func (s *LogStore) GetLogs(topic string, from int64, to int64, queryExp string,
	maxLineNum int64, offset int64, reverse bool) (*GetLogsResponse, error) {

	h := map[string]string{
		"x-log-bodyrawsize": "0",
		"Accept":            "application/json",
	}

	urlVal := url.Values{}
	urlVal.Add("type", "log")
	urlVal.Add("from", strconv.Itoa(int(from)))
	urlVal.Add("to", strconv.Itoa(int(to)))
	urlVal.Add("topic", topic)
	urlVal.Add("line", strconv.Itoa(int(maxLineNum)))
	urlVal.Add("offset", strconv.Itoa(int(offset)))
	urlVal.Add("reverse", strconv.FormatBool(reverse))
	urlVal.Add("query", queryExp)

	uri := fmt.Sprintf("/logstores/%s?%s", s.Name, urlVal.Encode())

	r, err := request(s.project, "GET", uri, h, nil)
	if err != nil {
		return nil, NewClientError(err)
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	if r.StatusCode != http.StatusOK {
		err := new(Error)
		if jErr := json.Unmarshal(body, err); jErr != nil {
			return nil, NewBadResponseError(string(body), r.Header, r.StatusCode)
		}
		return nil, err
	}

	logs := []map[string]string{}
	err = json.Unmarshal(body, &logs)
	if err != nil {
		return nil, NewBadResponseError(string(body), r.Header, r.StatusCode)
	}

	count, err := strconv.ParseInt(r.Header[GetLogsCountHeader][0], 10, 32)
	if err != nil {
		return nil, err
	}

	getLogsResponse := GetLogsResponse{
		Progress: r.Header[ProgressHeader][0],
		Count:    count,
		Logs:     logs,
	}

	return &getLogsResponse, nil
}

// CreateIndex ...
func (s *LogStore) CreateIndex(index Index) error {
	body, err := json.Marshal(index)
	if err != nil {
		return err
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}

	uri := fmt.Sprintf("/logstores/%s/index", s.Name)
	r, err := request(s.project, "POST", uri, h, body)
	if err != nil {
		return err
	}
	r.Body.Close()
	return nil
}

// UpdateIndex ...
func (s *LogStore) UpdateIndex(index Index) error {
	body, err := json.Marshal(index)
	if err != nil {
		return err
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}

	uri := fmt.Sprintf("/logstores/%s/index", s.Name)
	r, err := request(s.project, "PUT", uri, h, body)
	if r != nil {
		r.Body.Close()
	}
	return nil
}

// DeleteIndex ...
func (s *LogStore) DeleteIndex() error {
	type Body struct {
		project string `json:"projectName"`
		store   string `json:"logstoreName"`
	}

	body, err := json.Marshal(Body{
		project: s.project.Name,
		store:   s.Name,
	})
	if err != nil {
		return err
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}

	uri := fmt.Sprintf("/logstores/%s/index", s.Name)
	r, err := request(s.project, "DELETE", uri, h, body)
	if r != nil {
		r.Body.Close()
	}
	return nil
}

// GetIndex ...
func (s *LogStore) GetIndex() (*Index, error) {
	type Body struct {
		project string `json:"projectName"`
		store   string `json:"logstoreName"`
	}

	body, err := json.Marshal(Body{
		project: s.project.Name,
		store:   s.Name,
	})
	if err != nil {
		return nil, err
	}

	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}

	uri := fmt.Sprintf("/logstores/%s/index", s.Name)
	r, err := request(s.project, "GET", uri, h, body)
	if err != nil {
		return nil, err
	}
	index := &Index{}
	defer r.Body.Close()
	data, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(data, index)
	if err != nil {
		return nil, NewBadResponseError(string(data), r.Header, r.StatusCode)
	}

	return index, nil
}

// CheckIndexExist check index exist or not
func (s *LogStore) CheckIndexExist() (bool, error) {
	if _, err := s.GetIndex(); err != nil {
		if slsErr, ok := err.(*Error); ok {
			if slsErr.Code == "IndexConfigNotExist" {
				return false, nil
			}
			return false, slsErr
		}
		return false, err
	}

	return true, nil
}
