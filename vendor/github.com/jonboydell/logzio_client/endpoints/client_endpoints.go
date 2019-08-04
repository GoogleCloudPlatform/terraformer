package endpoints

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jonboydell/logzio_client"
	"github.com/jonboydell/logzio_client/client"
)

const (
	endpointServiceEndpoint = "%s/v1/endpoints"
)

const (
	fldEndpointId            string = "id"
	fldEndpointType          string = "endpointType"
	fldEndpointTitle         string = "title"
	fldEndpointDescription   string = "description"
	fldEndpointUrl           string = "url"
	fldEndpointMethod        string = "method"
	fldEndpointHeaders       string = "headers"
	fldEndpointBodyTemplate  string = "bodyTemplate"
	fldEndpointServiceKey    string = "serviceKey"
	fldEndpointApiToken      string = "apiToken"
	fldEndpointAppKey        string = "appKey"
	fldEndpointApiKey        string = "apiKey"
	fldEndpointRoutingKey    string = "routingKey"
	fldEndpointMessageType   string = "messageType"
	fldEndpointServiceApiKey string = "serviceApiKey"
)

const (
	EndpointTypeSlack     endpointType = "Slack"
	EndpointTypeCustom    endpointType = "Custom"
	EndpointTypePagerDuty endpointType = "PagerDuty"
	EndpointTypeBigPanda  endpointType = "BigPanda"
	EndpointTypeDataDog   endpointType = "Datadog"
	EndpointTypeVictorOps endpointType = "VictorOps"
)

type (
	endpointType string
)

type Endpoint struct {
	Id            int64             // all
	EndpointType  endpointType      // all
	Title         string            // all
	Description   string            // all
	Url           string            // custom & slack
	Method        string            // custom
	Headers       map[string]string // custom
	BodyTemplate  interface{}       // custom
	Message       string            // n.b. this is a hack to determine if there was an error (despite a 200 being returned)
	ServiceKey    string            // pager-duty
	ApiToken      string            // big-panda
	AppKey        string            // big-panda
	ApiKey        string            // data-dog
	RoutingKey    string            // victorops
	MessageType   string            // victorops
	ServiceApiKey string            // victorops
}

func jsonEndpointToEndpoint(jsonEndpoint map[string]interface{}) Endpoint {
	t := jsonEndpoint[fldEndpointType].(string)

	endpoint := Endpoint{
		Id:           int64(jsonEndpoint[fldEndpointId].(float64)),
		EndpointType: endpointType(t),
		Title:        jsonEndpoint[fldEndpointTitle].(string),
	}

	if jsonEndpoint[fldEndpointDescription] != nil {
		endpoint.Description = jsonEndpoint[fldEndpointDescription].(string)
	}

	switch endpoint.EndpointType {
	case EndpointTypeSlack:
		endpoint.Url = jsonEndpoint[fldEndpointUrl].(string)
	case EndpointTypeCustom:
		endpoint.Url = jsonEndpoint[fldEndpointUrl].(string)
		endpoint.BodyTemplate = jsonEndpoint[fldEndpointBodyTemplate]
		headerMap := make(map[string]string)
		headerString := jsonEndpoint[fldEndpointHeaders].(string)
		headers := strings.Split(headerString, ",")
		for _, header := range headers {
			kv := strings.Split(header, "=")
			headerMap[kv[0]] = kv[1]
		}
		endpoint.Headers = headerMap
		endpoint.Method = jsonEndpoint[fldEndpointMethod].(string)
	case EndpointTypePagerDuty:
		endpoint.ServiceKey = jsonEndpoint[fldEndpointServiceKey].(string)
	case EndpointTypeBigPanda:
		endpoint.ApiToken = jsonEndpoint[fldEndpointApiToken].(string)
		endpoint.AppKey = jsonEndpoint[fldEndpointAppKey].(string)
	case EndpointTypeDataDog:
		endpoint.ApiKey = jsonEndpoint[fldEndpointApiKey].(string)
	case EndpointTypeVictorOps:
		endpoint.RoutingKey = jsonEndpoint[fldEndpointRoutingKey].(string)
		endpoint.MessageType = jsonEndpoint[fldEndpointMessageType].(string)
		endpoint.ServiceApiKey = jsonEndpoint[fldEndpointServiceApiKey].(string)
	default:
		panic(fmt.Sprintf("unsupported endpoint type %s", endpoint.EndpointType))
	}

	return endpoint
}

type EndpointsClient struct {
	*client.Client
}

func New(apiToken, baseUrl string) (*EndpointsClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}
	c := &EndpointsClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}

type endpointValidator = func(e Endpoint) bool
type endpointBuilder = func(a string, t endpointType, e Endpoint) (*http.Request, error)
type endpointChecker = func(b []byte) error

func (c *EndpointsClient) makeEndpointRequest(endpoint interface{}, validator endpointValidator, builder endpointBuilder, checker endpointChecker) ([]byte, error, bool) {
	e := endpoint.(Endpoint)
	if !validator(e) {
		return nil, errors.New("the passed in endpoint is not valid"), false
	}
	req, _ := builder(c.ApiToken, e.EndpointType, e)
	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err, false
	}
	defer resp.Body.Close()
	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	if !logzio_client.CheckValidStatus(resp, []int{http.StatusOK}) {
		return nil, fmt.Errorf(errorCreateEndpointApiCallFailed, resp.StatusCode, jsonBytes), false
	}
	err = checker(jsonBytes)
	if err != nil {
		return nil, err, false
	}
	return jsonBytes, nil, true
}

func (c *EndpointsClient) getURLByType(t endpointType) string {
	switch t {
	case EndpointTypeSlack:
		return "slack"
	case EndpointTypeCustom:
		return "custom"
	case EndpointTypePagerDuty:
		return "pager-duty"
	case EndpointTypeBigPanda:
		return "big-panda"
	case EndpointTypeDataDog:
		return "data-dog"
	case EndpointTypeVictorOps:
		return "victorops"
	default:
		panic(fmt.Sprintf("unsupported endpoint type %s", t))
	}
}
