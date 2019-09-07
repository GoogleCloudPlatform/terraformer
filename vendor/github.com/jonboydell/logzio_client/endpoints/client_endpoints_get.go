package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
	"strings"
)

const getEndpointsServiceUrl string = endpointServiceEndpoint + "/%d"
const getEndpointsServiceMethod string = http.MethodGet
const getEndpointsMethodSuccess int = 200

const apiGetEndpointNoEndpoint = "The endpoint doesn't exist"

const errorGetEndpointApiCallFailed = "API call GetEndpoint failed with status code:%d, data:%s"
const errorGetEndpointDoesntExist = "API call GetEndpoint failed as endpoint with id:%d doesn't exist, data:%s"

func (c *EndpointsClient) buildGetEnpointApiRequest(apiToken string, notificationId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(getEndpointsServiceMethod, fmt.Sprintf(getEndpointsServiceUrl, baseUrl, notificationId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Returns an endpoint, given it's name.  Returns nil (and an error) if an endpoint with the specified name can't be found
func (c *EndpointsClient) GetEndpointByName(endpointName string) (*Endpoint, error) {
	list, err := c.ListEndpoints()
	if err != nil {
		return nil, err
	}

	for _, endpoint := range list {
		if endpoint.Title == endpointName {
			return &endpoint, nil
		}
	}

	return nil, err
}

// Returns an endpoint, given it's identity.  Returns nul (and an error) if an endpoint with the specified id can't be found
func (c *EndpointsClient) GetEndpoint(endpointId int64) (*Endpoint, error) {
	req, _ := c.buildGetEnpointApiRequest(c.ApiToken, endpointId)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// @todo: this isn't the idiomatic way to read a response body
	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{getEndpointsMethodSuccess}) {
		return nil, fmt.Errorf(errorGetEndpointApiCallFailed, resp.StatusCode, jsonBytes)
	}

	// sometimes logz.io returns a string rather than a json object (and a 200 status code), even though the call has failed
	// @todo: can this be refactored?
	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, apiGetEndpointNoEndpoint) {
		return nil, fmt.Errorf(errorGetEndpointDoesntExist, endpointId, str)
	}

	var jsonEndpoint map[string]interface{}
	err = json.Unmarshal([]byte(jsonBytes), &jsonEndpoint)
	if err != nil {
		return nil, err
	}

	endpoint := jsonEndpointToEndpoint(jsonEndpoint)

	return &endpoint, nil
}
