package fc

import (
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/resty.v1"
)

// Client defines fc client
type Client struct {
	Config  *Config
	Connect *Connection
}

// NewClient new fc client
func NewClient(endpoint, apiVersion, accessKeyID, accessKeySecret string, opts ...ClientOption) (*Client, error) {
	config := NewConfig()
	config.APIVersion = apiVersion
	config.AccessKeyID = accessKeyID
	config.AccessKeySecret = accessKeySecret
	config.Endpoint, config.host = GetAccessPoint(endpoint)
	connect := NewConnection()
	client := &Client{config, connect}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

// GetAccountSettings returns account settings from fc
func (c *Client) GetAccountSettings(input *GetAccountSettingsInput) (*GetAccountSettingsOutput, error) {
	if input == nil {
		input = new(GetAccountSettingsInput)
	}

	var output = new(GetAccountSettingsOutput)
	httpResponse, err := c.sendRequest(input, http.MethodGet)
	if err != nil {
		return nil, err
	}

	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// GetService returns service metadata from fc
func (c *Client) GetService(input *GetServiceInput) (*GetServiceOutput, error) {
	if input == nil {
		input = new(GetServiceInput)
	}

	var output = new(GetServiceOutput)
	httpResponse, err := c.sendRequest(input, http.MethodGet)
	if err != nil {
		return nil, err
	}

	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// ListServices returns list of services from fc
func (c *Client) ListServices(input *ListServicesInput) (*ListServicesOutput, error) {
	if input == nil {
		input = new(ListServicesInput)
	}

	var output = new(ListServicesOutput)
	httpResponse, err := c.sendRequest(input, http.MethodGet)
	if err != nil {
		return nil, err
	}

	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// UpdateService updates service
func (c *Client) UpdateService(input *UpdateServiceInput) (*UpdateServiceOutput, error) {
	if input == nil {
		input = new(UpdateServiceInput)
	}

	var output = new(UpdateServiceOutput)
	httpResponse, err := c.sendRequest(input, http.MethodPut)
	if err != nil {
		return nil, err
	}
	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// CreateService creates service
func (c *Client) CreateService(input *CreateServiceInput) (*CreateServiceOutput, error) {
	if input == nil {
		input = new(CreateServiceInput)
	}

	var output = new(CreateServiceOutput)
	httpResponse, err := c.sendRequest(input, http.MethodPost)
	if err != nil {
		return nil, err
	}
	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// DeleteService deletes service
func (c *Client) DeleteService(input *DeleteServiceInput) (*DeleteServiceOutput, error) {
	if input == nil {
		input = new(DeleteServiceInput)
	}
	var output = new(DeleteServiceOutput)
	httpResponse, err := c.sendRequest(input, http.MethodDelete)
	if err != nil {
		return nil, err
	}
	output.Header = httpResponse.Header()
	return output, nil
}

// PublishServiceVersion publishes service version
func (c *Client) PublishServiceVersion(input *PublishServiceVersionInput) (*PublishServiceVersionOutput, error) {
       if input == nil {
               input = new(PublishServiceVersionInput)
       }
       var output = new(PublishServiceVersionOutput)
       httpResponse, err := c.sendRequest(input, http.MethodPost)
       if err != nil {
               return nil, err
       }
       output.Header = httpResponse.Header()
       json.Unmarshal(httpResponse.Body(), output)
       return output, nil
}

// ListServiceVersions returns list of service versions
func (c *Client) ListServiceVersions(input *ListServiceVersionsInput) (*ListServiceVersionsOutput, error) {
       if input == nil {
               input = new(ListServiceVersionsInput)
       }

       var output = new(ListServiceVersionsOutput)
       httpResponse, err := c.sendRequest(input, http.MethodGet)
       if err != nil {
               return nil, err
       }

       output.Header = httpResponse.Header()
       json.Unmarshal(httpResponse.Body(), output)
       return output, nil
}

// DeleteServiceVersion marks service version as deleted
func (c *Client) DeleteServiceVersion(input *DeleteServiceVersionInput) (*DeleteServiceVersionOutput, error) {
       if input == nil {
               input = new(DeleteServiceVersionInput)
       }
       var output = new(DeleteServiceVersionOutput)
       httpResponse, err := c.sendRequest(input, http.MethodDelete)
       if err != nil {
               return nil, err
       }
       output.Header = httpResponse.Header()
       return output, nil
}

// CreateAlias creates alias
func (c *Client) CreateAlias(input *CreateAliasInput) (*CreateAliasOutput, error) {
       if input == nil {
               input = new(CreateAliasInput)
       }

       var output = new(CreateAliasOutput)
       httpResponse, err := c.sendRequest(input, http.MethodPost)
       if err != nil {
               return nil, err
       }
       output.Header = httpResponse.Header()
       json.Unmarshal(httpResponse.Body(), output)
       return output, nil
}

// UpdateAlias updates alias
func (c *Client) UpdateAlias(input *UpdateAliasInput) (*UpdateAliasOutput, error) {
       if input == nil {
               input = new(UpdateAliasInput)
       }

       var output = new(UpdateAliasOutput)
       httpResponse, err := c.sendRequest(input, http.MethodPut)
       if err != nil {
               return nil, err
       }
       output.Header = httpResponse.Header()
       json.Unmarshal(httpResponse.Body(), output)
       return output, nil
}

// GetAlias returns alias metadata from fc
func (c *Client) GetAlias(input *GetAliasInput) (*GetAliasOutput, error) {
       if input == nil {
               input = new(GetAliasInput)
       }

       var output = new(GetAliasOutput)
       httpResponse, err := c.sendRequest(input, http.MethodGet)
       if err != nil {
               return nil, err
       }

       output.Header = httpResponse.Header()
       json.Unmarshal(httpResponse.Body(), output)
       return output, nil
}

// ListAliases returns list of aliases from fc
func (c *Client) ListAliases(input *ListAliasesInput) (*ListAliasesOutput, error) {
       if input == nil {
               input = new(ListAliasesInput)
       }

       var output = new(ListAliasesOutput)
       httpResponse, err := c.sendRequest(input, http.MethodGet)
       if err != nil {
               return nil, err
       }

       output.Header = httpResponse.Header()
       json.Unmarshal(httpResponse.Body(), output)
       return output, nil
}

// DeleteAlias deletes service
func (c *Client) DeleteAlias(input *DeleteAliasInput) (*DeleteAliasOutput, error) {
       if input == nil {
               input = new(DeleteAliasInput)
       }
       var output = new(DeleteAliasOutput)
       httpResponse, err := c.sendRequest(input, http.MethodDelete)
       if err != nil {
               return nil, err
       }
       output.Header = httpResponse.Header()
       return output, nil
}


// CreateFunction creates function
func (c *Client) CreateFunction(input *CreateFunctionInput) (*CreateFunctionOutput, error) {
	if input == nil {
		input = new(CreateFunctionInput)
	}
	var output = new(CreateFunctionOutput)
	httpResponse, err := c.sendRequest(input, http.MethodPost)
	if err != nil {
		return nil, err
	}

	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// DeleteFunction deletes function from service
func (c *Client) DeleteFunction(input *DeleteFunctionInput) (*DeleteFunctionOutput, error) {
	if input == nil {
		input = new(DeleteFunctionInput)
	}

	var output = new(DeleteFunctionOutput)
	httpResponse, err := c.sendRequest(input, http.MethodDelete)
	if err != nil {
		return nil, err
	}

	output.Header = httpResponse.Header()
	return output, nil
}

// GetFunction returns function metadata from service
func (c *Client) GetFunction(input *GetFunctionInput) (*GetFunctionOutput, error) {
	if input == nil {
		input = new(GetFunctionInput)
	}

	var output = new(GetFunctionOutput)
	httpResponse, err := c.sendRequest(input, http.MethodGet)
	if err != nil {
		return nil, err
	}
	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// GetFunctionCode returns function code
func (c *Client) GetFunctionCode(input *GetFunctionCodeInput) (*GetFunctionCodeOutput, error) {
	if input == nil {
		input = new(GetFunctionCodeInput)
	}

	var output = new(GetFunctionCodeOutput)
	httpResponse, err := c.sendRequest(input, http.MethodGet)
	if err != nil {
		return nil, err
	}
	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// ListFunctions returns list of functions
func (c *Client) ListFunctions(input *ListFunctionsInput) (*ListFunctionsOutput, error) {
	if input == nil {
		input = new(ListFunctionsInput)
	}

	var output = new(ListFunctionsOutput)
	httpResponse, err := c.sendRequest(input, http.MethodGet)
	if err != nil {
		return nil, err
	}
	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// UpdateFunction updates function
func (c *Client) UpdateFunction(input *UpdateFunctionInput) (*UpdateFunctionOutput, error) {
	if input == nil {
		input = new(UpdateFunctionInput)
	}

	httpResponse, err := c.sendRequest(input, http.MethodPut)
	if err != nil {
		return nil, err
	}
	var output = new(UpdateFunctionOutput)
	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// CreateTrigger creates trigger
func (c *Client) CreateTrigger(input *CreateTriggerInput) (*CreateTriggerOutput, error) {
	if input == nil {
		input = new(CreateTriggerInput)
	}

	var output = new(CreateTriggerOutput)
	httpResponse, err := c.sendRequest(input, http.MethodPost)
	if err != nil {
		return nil, err
	}

	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// GetTrigger returns trigger metadata
func (c *Client) GetTrigger(input *GetTriggerInput) (*GetTriggerOutput, error) {
	if input == nil {
		input = new(GetTriggerInput)
	}

	var output = new(GetTriggerOutput)
	httpResponse, err := c.sendRequest(input, http.MethodGet)
	if err != nil {
		return nil, err
	}
	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// UpdateTrigger updates trigger
func (c *Client) UpdateTrigger(input *UpdateTriggerInput) (*UpdateTriggerOutput, error) {
	if input == nil {
		input = new(UpdateTriggerInput)
	}

	httpResponse, err := c.sendRequest(input, http.MethodPut)
	if err != nil {
		return nil, err
	}
	var output = new(UpdateTriggerOutput)
	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// DeleteTrigger deletes trigger
func (c *Client) DeleteTrigger(input *DeleteTriggerInput) (*DeleteTriggerOutput, error) {
	if input == nil {
		input = new(DeleteTriggerInput)
	}

	var output = new(DeleteTriggerOutput)
	httpResponse, err := c.sendRequest(input, http.MethodDelete)
	if err != nil {
		return nil, err
	}

	output.Header = httpResponse.Header()
	return output, nil
}

// ListTriggers returns list of triggers
func (c *Client) ListTriggers(input *ListTriggersInput) (*ListTriggersOutput, error) {
	if input == nil {
		input = new(ListTriggersInput)
	}

	var output = new(ListTriggersOutput)
	httpResponse, err := c.sendRequest(input, http.MethodGet)
	if err != nil {
		return nil, err
	}
	output.Header = httpResponse.Header()
	json.Unmarshal(httpResponse.Body(), output)
	return output, nil
}

// InvokeFunction : invoke function in fc
func (c *Client) InvokeFunction(input *InvokeFunctionInput) (*InvokeFunctionOutput, error) {
	if input == nil {
		input = new(InvokeFunctionInput)
	}

	var output = new(InvokeFunctionOutput)
	httpResponse, err := c.sendRequest(input, http.MethodPost)
	if err != nil {
		return nil, err
	}
	output.Header = httpResponse.Header()
	output.Payload = httpResponse.Body()

	return output, nil
}

func (c *Client) sendRequest(input ServiceInput, httpMethod string) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	var serviceError = new(ServiceError)
	path := "/" + c.Config.APIVersion + input.GetPath()

	headerParams := make(map[string]string)
	for k, v := range input.GetHeaders() {
		headerParams[k] = v
	}
	headerParams["Host"] = c.Config.host
	headerParams[HTTPHeaderAccountID] = c.Config.AccountID
	headerParams[HTTPHeaderUserAgent] = c.Config.UserAgent
	headerParams["Accept"] = "application/json"
	// Caution: should not declare this as byte[] whose zero value is an empty byte array
	// if input has no payload, the http body should not be populated at all.
	var rawBody interface{}
	if input.GetPayload() != nil {
		switch input.GetPayload().(type) {
		case *[]byte:
			headerParams["Content-Type"] = "application/octet-stream"
			b := input.GetPayload().(*[]byte)
			headerParams["Content-MD5"] = MD5(*b)
			rawBody = *b
		default:
			headerParams["Content-Type"] = "application/json"
			b, err := json.Marshal(input.GetPayload())
			if err != nil {
				// TODO: return client side error
				return nil, nil
			}
			headerParams["Content-MD5"] = MD5(b)
			rawBody = b
		}
	}
	headerParams["Date"] = time.Now().UTC().Format(http.TimeFormat)
	if c.Config.SecurityToken != "" {
		headerParams[HTTPHeaderSecurityToken] = c.Config.SecurityToken
	}
	headerParams["Authorization"] = GetAuthStr(c.Config.AccessKeyID, c.Config.AccessKeySecret, httpMethod, headerParams, path)
	resp, err := c.Connect.SendRequest(c.Config.Endpoint+path, httpMethod, rawBody, headerParams, input.GetQueryParams())
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		serviceError.RequestID = resp.Header().Get(HTTPHeaderRequestID)
		serviceError.HTTPStatus = resp.StatusCode()
		json.Unmarshal(resp.Body(), &serviceError)
		return nil, serviceError
	}
	return resp, nil
}
