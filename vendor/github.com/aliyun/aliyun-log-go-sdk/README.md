This is a Golang SDK for Alibaba Cloud [Log Service](https://sls.console.aliyun.com/).

API Reference :

* [Chinese](https://help.aliyun.com/document_detail/29007.html)
* [English](https://intl.aliyun.com/help/doc-detail/29007.htm)

[![Build Status](https://travis-ci.org/aliyun/aliyun-log-go-sdk.svg?branch=master)](https://travis-ci.org/aliyun/aliyun-log-go-sdk)
[![Coverage Status](https://coveralls.io/repos/github/aliyun/aliyun-log-go-sdk/badge.svg?branch=master)](https://coveralls.io/github/aliyun/aliyun-log-go-sdk?branch=master)


# Install Instruction

### LogHub Golang SDK

```
go get github.com/aliyun/aliyun-log-go-sdk
```

# Example 

### Write and Read LogHub

[loghub_sample.go](example/loghub/loghub_sample.go)

### Use Index on LogHub (SLS)

[index_sample.go](example/index/index_sample.go)

### Create Config for Logtail

[log_config_sample.go](example/config/log_config_sample.go)

### Create Machine Group for Logtail

[machine_group_sample.go](example/machine_group/machine_group_sample.go)

# For developer
### Update log protobuf
`protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gofast_out=. log.proto`
