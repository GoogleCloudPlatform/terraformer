GOPATH=$(shell go env | grep GOPATH | sed 's/GOPATH=//' | sed 's/"//g')
gogoproto="$(GOPATH)/src/github.com/gogo/protobuf/gogoproto"
protobuf="$(GOPATH)/src/github.com/gogo/protobuf/protobuf"

all: log.pb.go
	go build

log.pb.go: log.proto
	go get github.com/gogo/protobuf/proto
	go get github.com/gogo/protobuf/protoc-gen-gogo
	go get github.com/gogo/protobuf/gogoproto
	protoc --proto_path=.:${gogoproto}:${protobuf} --gogo_out=. log.proto

test:
	go test

