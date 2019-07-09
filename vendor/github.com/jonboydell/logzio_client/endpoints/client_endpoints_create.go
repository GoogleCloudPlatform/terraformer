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
	createEndpointServiceUrl    string = endpointServiceEndpoint + "/%s"
	createEndpointServiceMethod string = http.MethodPost
)

const (
	errorCreateEndpointApiCallFailed = "API call CreateEndpoint failed with status code %d, data: %s"
)

func buildCreateEndpointRequest(endpoint Endpoint) map[string]interface{} {
	var createEndpoint = map[string]interface{}{}

	createEndpoint[fldEndpointTitle] = endpoint.Title
	createEndpoint[fldEndpointDescription] = endpoint.Description

	if endpoint.EndpointType == EndpointTypeSlack {
		createEndpoint[fldEndpointUrl] = endpoint.Url
	}

	if endpoint.EndpointType == EndpointTypeCustom {
		createEndpoint[fldEndpointUrl] = endpoint.Url
		createEndpoint[fldEndpointMethod] = endpoint.Method
		headers := endpoint.Headers
		headerStrings := []string{}
		for k, v := range headers {
			headerStrings = append(headerStrings, fmt.Sprintf("%s=%s", k, v))
		}
		headerString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(headerStrings)), ","), "[]")
		createEndpoint[fldEndpointHeaders] = headerString
		createEndpoint[fldEndpointBodyTemplate] = endpoint.BodyTemplate
	}

	if endpoint.EndpointType == EndpointTypePagerDuty {
		createEndpoint[fldEndpointServiceKey] = endpoint.ServiceKey
	}

	if endpoint.EndpointType == EndpointTypeBigPanda {
		createEndpoint[fldEndpointApiToken] = endpoint.ApiToken
		createEndpoint[fldEndpointAppKey] = endpoint.AppKey
	}

	if endpoint.EndpointType == EndpointTypeDataDog {
		createEndpoint[fldEndpointApiKey] = endpoint.ApiKey
	}

	if endpoint.EndpointType == EndpointTypeVictorOps {
		createEndpoint[fldEndpointRoutingKey] = endpoint.RoutingKey
		createEndpoint[fldEndpointMessageType] = endpoint.MessageType
		createEndpoint[fldEndpointServiceApiKey] = endpoint.ServiceApiKey
	}

	return createEndpoint
}

func (c *EndpointsClient) buildCreateEndpointApiRequest(apiToken string, endpointType endpointType, endpoint Endpoint) (*http.Request, error) {
	createEndpoint := buildCreateEndpointRequest(endpoint)

	jsonBytes, err := json.Marshal(createEndpoint)
	if err != nil {
		return nil, err
	}

	baseUrl := c.BaseUrl
	url := fmt.Sprintf(createEndpointServiceUrl, baseUrl, c.getURLByType(endpointType))
	req, err := http.NewRequest(createEndpointServiceMethod, url, bytes.NewBuffer(jsonBytes))
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Creates an endpoint, given the endpoint definition and the service to create the endpoint against
// Returns the endpoint object if successful (hopefully with an ID) and a non-nil error if not
func (c *EndpointsClient) CreateEndpoint(endpoint Endpoint) (*Endpoint, error) {
	if jsonBytes, err, ok := c.makeEndpointRequest(endpoint, ValidateEndpointRequest, c.buildCreateEndpointApiRequest, func(b []byte) error {
		var data map[string]interface{}
		json.Unmarshal(b, &data)

		if val, ok := data["errorCode"]; ok {
			return fmt.Errorf("%v", val)
		}

		if val, ok := data["message"]; ok {
			return fmt.Errorf("%v", val)
		}

		if strings.Contains(fmt.Sprintf("%s", b), errorCreateEndpointApiCallFailed) {
			return fmt.Errorf(errorCreateEndpointApiCallFailed, 200, errorCreateEndpointApiCallFailed)
		}
		return nil
	}); !ok {
		return nil, err
	} else {
		var target Endpoint
		err = json.Unmarshal(jsonBytes, &target)
		if err != nil {
			return nil, err
		}

		return &target, nil
	}
}
