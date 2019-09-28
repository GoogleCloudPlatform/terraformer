package fc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	customDomainsPath = "/custom-domains"
	singleDomainPath  = customDomainsPath + "/%s"
)

type customDomainMetadata struct {
	DomainName       *string      `json:"domainName"`
	AccountID        *string      `json:"accountId"`
	Protocol         *string      `json:"protocol"`
	APIVersion       *string      `json:"apiVersion"`
	RouteConfig      *RouteConfig `json:"routeConfig"`
	CreatedTime      string       `json:"createdTime"`
	LastModifiedTime string       `json:"lastModifiedTime"`
}

// CreateCustomDomainInput defines input to create custom domain
type CreateCustomDomainInput struct {
	DomainName  *string      `json:"domainName"`
	Protocol    *string      `json:"protocol"`
	RouteConfig *RouteConfig `json:"routeConfig"`
}

// NewCreateCustomDomainInput...
func NewCreateCustomDomainInput() *CreateCustomDomainInput {
	return &CreateCustomDomainInput{}
}

func (c *CreateCustomDomainInput) WithDomainName(domainName string) *CreateCustomDomainInput {
	c.DomainName = &domainName
	return c
}

func (c *CreateCustomDomainInput) WithProtocol(protocol string) *CreateCustomDomainInput {
	c.Protocol = &protocol
	return c
}

func (c *CreateCustomDomainInput) WithRouteConfig(routeConfig *RouteConfig) *CreateCustomDomainInput {
	c.RouteConfig = routeConfig
	return c
}

func (c *CreateCustomDomainInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (c *CreateCustomDomainInput) GetPath() string {
	return customDomainsPath
}

func (c *CreateCustomDomainInput) GetHeaders() Header {
	return make(Header, 0)
}

func (c *CreateCustomDomainInput) GetPayload() interface{} {
	return c
}

func (c *CreateCustomDomainInput) Validate() error {
	return nil
}

// RouteConfig represents route config for a single domain
type RouteConfig struct {
	Routes []PathConfig `json:"routes" `
}

func NewRouteConfig() *RouteConfig {
	return &RouteConfig{}
}

func (r *RouteConfig) WithRoutes(pathConfig []PathConfig) *RouteConfig {
	r.Routes = pathConfig
	return r
}

// PathConfig represents path-function mapping
type PathConfig struct {
	Path         *string `json:"path" `
	ServiceName  *string `json:"serviceName" `
	FunctionName *string `json:"functionName" `
	Qualifier    *string `json:"qualifier" `
}

func NewPathConfig() *PathConfig {
	return &PathConfig{}
}

func (p *PathConfig) WithPath(path string) *PathConfig {
	p.Path = &path
	return p
}

func (p *PathConfig) WithServiceName(serviceName string) *PathConfig {
	p.ServiceName = &serviceName
	return p
}

func (p *PathConfig) WithFunctionName(functionName string) *PathConfig {
	p.FunctionName = &functionName
	return p
}

func (p *PathConfig) WithQualifier(qualifier string) *PathConfig {
	p.Qualifier = &qualifier
	return p
}

// CreateCustomDomainOutput define create custom domain response
type CreateCustomDomainOutput struct {
	Header http.Header
	customDomainMetadata
}

func (o CreateCustomDomainOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o CreateCustomDomainOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

// UpdateCustomDomainInput defines input to update custom domain
type UpdateCustomDomainObject struct {
	Protocol    *string      `json:"protocol"`
	RouteConfig *RouteConfig `json:"routeConfig"`
}

type UpdateCustomDomainInput struct {
	DomainName *string
	UpdateCustomDomainObject
}

func NewUpdateCustomDomainInput(domainName string) *UpdateCustomDomainInput {
	return &UpdateCustomDomainInput{DomainName: &domainName}
}

func (c *UpdateCustomDomainInput) WithProtocol(protocol string) *UpdateCustomDomainInput {
	c.Protocol = &protocol
	return c
}

func (c *UpdateCustomDomainInput) WithRouteConfig(routeConfig *RouteConfig) *UpdateCustomDomainInput {
	c.RouteConfig = routeConfig
	return c
}

func (c *UpdateCustomDomainInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (c *UpdateCustomDomainInput) GetPath() string {
	return fmt.Sprintf(singleDomainPath, pathEscape(*c.DomainName))
}

func (c *UpdateCustomDomainInput) GetHeaders() Header {
	header := make(Header)
	return header
}

func (c *UpdateCustomDomainInput) GetPayload() interface{} {
	return c.UpdateCustomDomainObject
}

func (c *UpdateCustomDomainInput) Validate() error {
	if IsBlank(c.DomainName) {
		return fmt.Errorf("Domain name is required but not provided")
	}
	return nil
}

// UpdateCustomDomainOutput define update custom domain response
type UpdateCustomDomainOutput struct {
	Header http.Header
	customDomainMetadata
}

func (o UpdateCustomDomainOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o UpdateCustomDomainOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

type GetCustomDomainInput struct {
	DomainName *string
}

func NewGetCustomDomainInput(domainName string) *GetCustomDomainInput {
	return &GetCustomDomainInput{DomainName: &domainName}
}

func (i *GetCustomDomainInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *GetCustomDomainInput) GetPath() string {
	return fmt.Sprintf(singleDomainPath, pathEscape(*i.DomainName))
}

func (i *GetCustomDomainInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *GetCustomDomainInput) GetPayload() interface{} {
	return nil
}

func (i *GetCustomDomainInput) Validate() error {
	if IsBlank(i.DomainName) {
		return fmt.Errorf("Domain name is required but not provided")
	}
	return nil
}

// GetCustomDomainOutput define get custom domain response
type GetCustomDomainOutput struct {
	Header http.Header
	customDomainMetadata
}

func (o GetCustomDomainOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o GetCustomDomainOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

// ListCustomDomains defines listCustomDomainsMetadata result
type ListCustomDomainsOutput struct {
	Header        http.Header
	CustomDomains []*customDomainMetadata `json:"customDomains"`
	NextToken     *string                 `json:"nextToken,omitempty"`
}

type ListCustomDomainsInput struct {
	Query
}

func NewListCustomDomainsInput() *ListCustomDomainsInput {
	return &ListCustomDomainsInput{}
}

func (i *ListCustomDomainsInput) WithPrefix(prefix string) *ListCustomDomainsInput {
	i.Prefix = &prefix
	return i
}

func (i *ListCustomDomainsInput) WithStartKey(startKey string) *ListCustomDomainsInput {
	i.StartKey = &startKey
	return i
}

func (i *ListCustomDomainsInput) WithNextToken(nextToken string) *ListCustomDomainsInput {
	i.NextToken = &nextToken
	return i
}

func (i *ListCustomDomainsInput) WithLimit(limit int32) *ListCustomDomainsInput {
	i.Limit = &limit
	return i
}

func (i *ListCustomDomainsInput) GetQueryParams() url.Values {
	out := url.Values{}
	if i.Prefix != nil {
		out.Set("prefix", *i.Prefix)
	}

	if i.StartKey != nil {
		out.Set("startKey", *i.StartKey)
	}

	if i.NextToken != nil {
		out.Set("nextToken", *i.NextToken)
	}

	if i.Limit != nil {
		out.Set("limit", strconv.FormatInt(int64(*i.Limit), 10))
	}

	return out
}

func (i *ListCustomDomainsInput) GetPath() string {
	return customDomainsPath
}

func (i *ListCustomDomainsInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *ListCustomDomainsInput) GetPayload() interface{} {
	return nil
}

func (i *ListCustomDomainsInput) Validate() error {
	return nil
}

func (o ListCustomDomainsOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o ListCustomDomainsOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

type DeleteCustomDomainInput struct {
	DomainName *string
}

func NewDeleteCustomDomainInput(domainName string) *DeleteCustomDomainInput {
	return &DeleteCustomDomainInput{DomainName: &domainName}
}

func (i *DeleteCustomDomainInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *DeleteCustomDomainInput) GetPath() string {
	return fmt.Sprintf(singleDomainPath, pathEscape(*i.DomainName))
}

func (i *DeleteCustomDomainInput) GetHeaders() Header {
	header := make(Header)
	return header
}

func (i *DeleteCustomDomainInput) GetPayload() interface{} {
	return nil
}

func (i *DeleteCustomDomainInput) Validate() error {
	if IsBlank(i.DomainName) {
		return fmt.Errorf("Domain name is required but not provided")
	}
	return nil
}

type DeleteCustomDomainOutput struct {
	Header http.Header
}

func (o DeleteCustomDomainOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o DeleteCustomDomainOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}