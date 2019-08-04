package endpoints

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jonboydell/logzio_client"
)

const (
	deleteEndpointServiceUrl    string = endpointServiceEndpoint + "/%d"
	deleteEndpointServiceMethod string = http.MethodDelete
	deleteEndpointMethodSuccess int    = 200
)

const (
	errorDeleteEndpointDoesntExist = "API call DeleteEndpoint failed as endpoint with id:%d doesn't exist, data:%s"
)

func validateDeleteEndpoint(endpoint Endpoint) bool {
	return true
}

func (c *EndpointsClient) buildDeleteEndpointApiRequest(apiToken string, service endpointType, endpoint Endpoint) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(deleteEndpointServiceMethod, fmt.Sprintf(deleteEndpointServiceUrl, baseUrl, endpoint.Id), nil)
	logzio_client.AddHttpHeaders(apiToken, req)
	return req, err
}

// Deletes an endpoint with the given id, returns a non nil error otherwise
func (c *EndpointsClient) DeleteEndpoint(endpointId int64) error {
	if _, err, ok := c.makeEndpointRequest(Endpoint{Id: endpointId}, validateDeleteEndpoint, c.buildDeleteEndpointApiRequest, func(body []byte) error {
		if strings.Contains(fmt.Sprintf("%s", body), "endpoints/FORBIDDEN_OPERATION") {
			return fmt.Errorf(errorDeleteEndpointDoesntExist, endpointId, body)
		}
		if strings.Contains(fmt.Sprintf("%s", body), "endpoints/UNKNOWN_ENDPOINT") {
			return fmt.Errorf(errorDeleteEndpointDoesntExist, endpointId, body)
		}
		return nil
	}); !ok {
		return err
	}
	return nil
}
