package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jonboydell/logzio_client"
)

const (
	updateEndpointServiceUrl    string = endpointServiceEndpoint + "/%s/%d"
	updateEndpointServiceMethod string = http.MethodPut
	updateEndpointMethodSuccess int    = 200
)

const (
	errorUpdateEndpointApiCallFailed = "API call UpdateEndpoint failed with status code:%d, data:%s"
	errorUpdateEndpointDoesntExist   = "API call UpdateEndpoint failed as endpoint with id:%d doesn't exist, data:%s"
)

func (c *EndpointsClient) buildUpdateEndpointApiRequest(apiToken string, endpointType endpointType, endpoint Endpoint) (*http.Request, error) {
	jsonObject, err := buildUpdateEndpointRequest(endpoint)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(jsonObject)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	id := endpoint.Id
	req, err := http.NewRequest(updateEndpointServiceMethod, fmt.Sprintf(updateEndpointServiceUrl, baseUrl, strings.ToLower(c.getURLByType(endpointType)), id), bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

func buildUpdateEndpointRequest(endpoint Endpoint) (map[string]interface{}, error) {
	var updateEndpoint = map[string]interface{}{}

	updateEndpoint[fldEndpointTitle] = endpoint.Title
	updateEndpoint[fldEndpointDescription] = endpoint.Description

	switch endpoint.EndpointType {
	case EndpointTypeSlack:
		updateEndpoint[fldEndpointUrl] = endpoint.Url
	case EndpointTypeCustom:
		updateEndpoint[fldEndpointUrl] = endpoint.Url
		updateEndpoint[fldEndpointMethod] = endpoint.Method
		headers := endpoint.Headers
		headerStrings := []string{}
		for k, v := range headers {
			headerStrings = append(headerStrings, fmt.Sprintf("%s=%s", k, v))
		}
		headerString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(headerStrings)), ","), "[]")
		updateEndpoint[fldEndpointHeaders] = headerString
		updateEndpoint[fldEndpointBodyTemplate] = endpoint.BodyTemplate
	case EndpointTypePagerDuty:
		updateEndpoint[fldEndpointServiceKey] = endpoint.ServiceKey
	case EndpointTypeBigPanda:
		updateEndpoint[fldEndpointApiToken] = endpoint.ApiToken
		updateEndpoint[fldEndpointAppKey] = endpoint.AppKey
	case EndpointTypeDataDog:
		updateEndpoint[fldEndpointApiKey] = endpoint.ApiKey
	case EndpointTypeVictorOps:
		updateEndpoint[fldEndpointRoutingKey] = endpoint.RoutingKey
		updateEndpoint[fldEndpointMessageType] = endpoint.MessageType
		updateEndpoint[fldEndpointServiceApiKey] = endpoint.ServiceApiKey
	default:
		return nil, fmt.Errorf("don't recognise endpoint type %s", endpoint.EndpointType)
	}

	return updateEndpoint, nil
}

// Updates an existing endpoint, returns the updated endpoint if successful, an error otherwise
func (c *EndpointsClient) UpdateEndpoint(id int64, endpoint Endpoint) (*Endpoint, error) {

	endpoint.Id = id
	if jsonBytes, err, ok := c.makeEndpointRequest(endpoint, ValidateEndpointRequest, c.buildUpdateEndpointApiRequest, func(b []byte) error {
		if strings.Contains(fmt.Sprintf("%s", b), "Insufficient privileges") {
			return fmt.Errorf("API call %s failed for endpoint %d, data: %s", "UpdateEndpoint", id, b)
		}
		if strings.Contains(fmt.Sprintf("%s", b), "errorCode") {
			return fmt.Errorf("API call %s failed for endpoint %d, data: %s", "UpdateEndpoint", id, b)
		}

		return nil
	}); ok {
		var target Endpoint
		err = json.Unmarshal(jsonBytes, &target)
		if err != nil {
			return nil, err
		}
		return &endpoint, nil
	} else {
		return nil, err
	}
}
