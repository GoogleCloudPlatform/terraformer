// Copyright 2022 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tencentcloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TencentCloudService struct { //nolint
	terraformutils.Service
}

const RequestClient = "TENCENTCLOUD_API_REQUEST_CLIENT"

var ReqClient = "Terraformer-latest"

type LogRoundTripper struct {
	Verbose bool
}

func (me *LogRoundTripper) RoundTrip(request *http.Request) (response *http.Response, errRet error) {
	var inBytes, outBytes []byte
	var start = time.Now()

	if me.Verbose {
		defer func() { me.log(inBytes, outBytes, errRet, start) }()
	}

	bodyReader, errRet := request.GetBody()
	if errRet != nil {
		return
	}
	var headName = "X-TC-Action"

	if envReqClient := os.Getenv(RequestClient); envReqClient != "" {
		ReqClient = envReqClient
	}

	request.Header.Set("X-TC-RequestClient", ReqClient)
	inBytes = []byte(fmt.Sprintf("%s, request: ", request.Header[headName]))
	requestBody, errRet := ioutil.ReadAll(bodyReader)
	if errRet != nil {
		return
	}
	inBytes = append(inBytes, requestBody...)

	headName = "X-TC-Region"
	appendMessage := []byte(fmt.Sprintf(
		", (host %+v, region:%+v)",
		request.Header["Host"],
		request.Header[headName],
	))

	inBytes = append(inBytes, appendMessage...)

	response, errRet = http.DefaultTransport.RoundTrip(request)
	if errRet != nil {
		return
	}
	outBytes, errRet = ioutil.ReadAll(response.Body)
	if errRet != nil {
		return
	}
	response.Body = ioutil.NopCloser(bytes.NewBuffer(outBytes))

	return
}

func (me *LogRoundTripper) log(in []byte, out []byte, err error, start time.Time) {
	var buf bytes.Buffer
	buf.WriteString("#########")
	tag := "[DEBUG]"
	if err != nil {
		tag = "[CRITICAL]"
	}
	buf.WriteString(tag)
	if len(in) > 0 {
		buf.WriteString("tencentcloud-sdk-go: ")
		buf.Write(in)
	}
	if len(out) > 0 {
		buf.WriteString("; response:")
		err := json.Compact(&buf, out)
		if err != nil {
			out := bytes.Replace(out,
				[]byte("\n"),
				[]byte(""),
				-1)
			out = bytes.Replace(out,
				[]byte(" "),
				[]byte(""),
				-1)
			buf.Write(out)
		}
	}

	if err != nil {
		buf.WriteString("; error:")
		buf.WriteString(err.Error())
	}

	costFormat := fmt.Sprintf(",cost %s", time.Since(start).String())
	buf.WriteString(costFormat)

	fmt.Println(buf.String())
}
