package search

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore/otsprotocol"
	"github.com/golang/protobuf/proto"
)

type ExistsQuery struct {
	FieldName string
}

func (q *ExistsQuery) Type() QueryType {
	return QueryType_ExistsQuery
}

func (q *ExistsQuery) Serialize() ([]byte, error) {
	query := &otsprotocol.ExistsQuery{}
	query.FieldName = &q.FieldName
	data, err := proto.Marshal(query)
	return data, err
}

func (q *ExistsQuery) ProtoBuffer() (*otsprotocol.Query, error) {
	return BuildPBForQuery(q)
}