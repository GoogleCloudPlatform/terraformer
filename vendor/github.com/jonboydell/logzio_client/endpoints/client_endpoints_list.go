package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
	"io/ioutil"
	"net/http"
)

const listEndpointsServiceUrl string = endpointServiceEndpoint
const listEndpointsServiceMethod string = http.MethodGet
const listEndpointsMethodSuccess int = 200

const errorListEndpointApiCallFailed = "API call ListEndpoints failed with status code:%d, data:%s"

func (c *EndpointsClient) buildListEndpointsApiRequest(apiToken string) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(listEndpointsServiceMethod, fmt.Sprintf(listEndpointsServiceUrl, baseUrl), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Lists all the current endpoints in an array, returns an error otherwise
func (c *EndpointsClient) ListEndpoints() ([]Endpoint, error) {
	req, _ := c.buildListEndpointsApiRequest(c.ApiToken)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{listEndpointsMethodSuccess}) {
		return nil, fmt.Errorf(errorListEndpointApiCallFailed, resp.StatusCode, jsonBytes)
	}

	var arr []Endpoint
	var jsonResponse []interface{}
	err = json.Unmarshal([]byte(jsonBytes), &jsonResponse)

	for _, json := range jsonResponse {
		jsonEndpoint := json.(map[string]interface{})
		endpoint := jsonEndpointToEndpoint(jsonEndpoint)
		arr = append(arr, endpoint)
	}

	return arr, nil
}
