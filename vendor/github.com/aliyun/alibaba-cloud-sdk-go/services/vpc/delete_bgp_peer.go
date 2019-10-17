package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteBgpPeer invokes the vpc.DeleteBgpPeer API synchronously
// api document: https://help.aliyun.com/api/vpc/deletebgppeer.html
func (client *Client) DeleteBgpPeer(request *DeleteBgpPeerRequest) (response *DeleteBgpPeerResponse, err error) {
	response = CreateDeleteBgpPeerResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBgpPeerWithChan invokes the vpc.DeleteBgpPeer API asynchronously
// api document: https://help.aliyun.com/api/vpc/deletebgppeer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBgpPeerWithChan(request *DeleteBgpPeerRequest) (<-chan *DeleteBgpPeerResponse, <-chan error) {
	responseChan := make(chan *DeleteBgpPeerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBgpPeer(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteBgpPeerWithCallback invokes the vpc.DeleteBgpPeer API asynchronously
// api document: https://help.aliyun.com/api/vpc/deletebgppeer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBgpPeerWithCallback(request *DeleteBgpPeerRequest, callback func(response *DeleteBgpPeerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBgpPeerResponse
		var err error
		defer close(result)
		response, err = client.DeleteBgpPeer(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteBgpPeerRequest is the request struct for api DeleteBgpPeer
type DeleteBgpPeerRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	BgpPeerId            string           `position:"Query" name:"BgpPeerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteBgpPeerResponse is the response struct for api DeleteBgpPeer
type DeleteBgpPeerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteBgpPeerRequest creates a request to invoke DeleteBgpPeer API
func CreateDeleteBgpPeerRequest() (request *DeleteBgpPeerRequest) {
	request = &DeleteBgpPeerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteBgpPeer", "vpc", "openAPI")
	return
}

// CreateDeleteBgpPeerResponse creates a response to parse from DeleteBgpPeer response
func CreateDeleteBgpPeerResponse() (response *DeleteBgpPeerResponse) {
	response = &DeleteBgpPeerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
