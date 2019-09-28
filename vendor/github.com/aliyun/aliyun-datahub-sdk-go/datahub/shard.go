package datahub

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ShardAbstract struct {
	Id           string `json:"ShardId"`
	BeginHashKey string `json:"BeginHashKey"`
	EndHashKey   string `json:"EndHashKey"`
}

func (s ShardAbstract) String() string {
	sbytes, _ := json.Marshal(s)
	return string(sbytes)
}

type Shard struct {
	Id             string     `json:"ShardId"`
	State          ShardState `json:"State"`
	BeginHashKey   string     `json:"BeginHashKey"`
	EndHashKey     string     `json:"EndHashKey"`
	ClosedTime     uint64     `json:"ClosedTime"`
	ParentShardIds []string   `json:"ParentShardIds"`
	LeftShardId    string     `json:"LeftShardId"`
	RightShardId   string     `json:"RightShardId"`
}

func (s Shard) String() string {
	sbytes, _ := json.Marshal(s)
	return string(sbytes)
}

// Shards for list shard
type Shards struct {
	ShardList []Shard `json:"Shards"`
}

func (ss *Shards) String() string {
	ssbytes, _ := json.Marshal(ss)
	return string(ssbytes)
}

func (ss *Shards) RequestBodyEncode(method string) ([]byte, error) {
	switch method {
	case http.MethodGet:
		return nil, nil
	default:
		return nil, errors.New(fmt.Sprintf("Shards not support method %s", method))
	}
}

func (ss *Shards) ResponseBodyDecode(method string, body []byte) error {
	switch method {
	case http.MethodGet:
		return json.Unmarshal(body, ss)
	default:
		return errors.New(fmt.Sprintf("Shards not support method %s", method))
	}
}

// MergeShard
type MergeShard struct {
	Id              string        `json:"ShardId"`
	AdjacentShardId string        `json:"AdjacentShardId"`
	NewShard        ShardAbstract `json:"NewShard"`
}

func (ms *MergeShard) RequestBodyEncode(method string) ([]byte, error) {
	switch method {
	case http.MethodPost:
		reqMsg := struct {
			Action          string `json:"Action"`
			ShardId         string `json:"ShardId"`
			AdjacentShardId string `json:"AdjacentShardId"`
		}{
			Action:          "merge",
			ShardId:         ms.Id,
			AdjacentShardId: ms.AdjacentShardId,
		}
		return json.Marshal(reqMsg)
	default:
		return nil, errors.New(fmt.Sprintf("MergeShard not support method %s", method))
	}
}

func (ms *MergeShard) ResponseBodyDecode(method string, body []byte) error {
	switch method {
	case http.MethodPost:
		return json.Unmarshal(body, &ms.NewShard)
	default:
		return errors.New(fmt.Sprintf("MergeShard not support method %s", method))
	}
}

// SplitShard
type SplitShard struct {
	Id        string          `json:"ShardId"`
	SplitKey  string          `json:"SplitKey"`
	NewShards []ShardAbstract `json:"NewShards"`
}

func (ss *SplitShard) RequestBodyEncode(method string) ([]byte, error) {
	switch method {
	case http.MethodPost:
		reqMsg := struct {
			Action   string `json:"Action"`
			ShardId  string `json:"ShardId"`
			SplitKey string `json:"SplitKey"`
		}{
			Action:   "split",
			ShardId:  ss.Id,
			SplitKey: ss.SplitKey,
		}
		return json.Marshal(reqMsg)
	default:
		return nil, errors.New(fmt.Sprintf("SplitShard not support method %s", method))
	}
}

func (ss *SplitShard) ResponseBodyDecode(method string, body []byte) error {
	switch method {
	case http.MethodPost:
		return json.Unmarshal(body, ss)
	default:
		return errors.New(fmt.Sprintf("SplitShard not support method %s", method))
	}
}
