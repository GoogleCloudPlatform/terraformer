// Copyright (c) 2015 Jeevanandam M (jeeva@myjeeva.com), All rights reserved.
// resty source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package resty

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const (
	// GET HTTP method
	GET = "GET"

	// POST HTTP method
	POST = "POST"

	// PUT HTTP method
	PUT = "PUT"

	// DELETE HTTP method
	DELETE = "DELETE"

	// PATCH HTTP method
	PATCH = "PATCH"

	// HEAD HTTP method
	HEAD = "HEAD"

	// OPTIONS HTTP method
	OPTIONS = "OPTIONS"
)

var (
	hdrUserAgentKey     = http.CanonicalHeaderKey("User-Agent")
	hdrAcceptKey        = http.CanonicalHeaderKey("Accept")
	hdrContentTypeKey   = http.CanonicalHeaderKey("Content-Type")
	hdrContentLengthKey = http.CanonicalHeaderKey("Content-Length")
	hdrAuthorizationKey = http.CanonicalHeaderKey("Authorization")

	plainTextType   = "text/plain; charset=utf-8"
	jsonContentType = "application/json; charset=utf-8"
	formContentType = "application/x-www-form-urlencoded"

	plainTextCheck = regexp.MustCompile("(?i:text/plain)")
	jsonCheck      = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck       = regexp.MustCompile("(?i:[application|text]/xml)")

	hdrUserAgentValue = "go-resty v%s - https://github.com/go-resty/resty"
)

// Client type is used for HTTP/RESTful global values
// for all request raised from the client
type Client struct {
	HostURL    string
	QueryParam url.Values
	FormData   url.Values
	Header     http.Header
	UserInfo   *User
	Token      string
	Cookies    []*http.Cookie
	Error      reflect.Type
	Debug      bool
	Log        *log.Logger

	httpClient       *http.Client
	transport        *http.Transport
	setContentLength bool
	isHTTPMode       bool
	outputDirectory  string
	beforeRequest    []func(*Client, *Request) error
	afterResponse    []func(*Client, *Response) error
}

// User type is to hold an username and password information
type User struct {
	Username, Password string
}

// SetHostURL method is to set Host URL in the client instance. It will be used with request
// raised from this client with relative URL
//		// Setting HTTP address
//		resty.SetHostURL("http://myjeeva.com")
//
//		// Setting HTTPS address
//		resty.SetHostURL("https://myjeeva.com")
//
func (c *Client) SetHostURL(url string) *Client {
	c.HostURL = strings.TrimRight(url, "/")
	return c
}

// SetHeader method sets a single header field and its value in the client instance.
// These headers will be applied to all requests raised from this client instance.
// Also it can be overridden at request level header options, see `resty.R().SetHeader`
// or `resty.R().SetHeaders`.
//
// For Example: To set `Content-Type` and `Accept` as `application/json`
//
// 		resty.
// 			SetHeader("Content-Type", "application/json").
// 			SetHeader("Accept", "application/json")
//
func (c *Client) SetHeader(header, value string) *Client {
	c.Header.Set(header, value)
	return c
}

// SetHeaders method sets multiple headers field and its values at one go in the client instance.
// These headers will be applied to all requests raised from this client instance. Also it can be
// overridden at request level headers options, see `resty.R().SetHeaders` or `resty.R().SetHeader`.
//
// For Example: To set `Content-Type` and `Accept` as `application/json`
//
// 		resty.SetHeaders(map[string]string{
//				"Content-Type": "application/json",
//				"Accept": "application/json",
//			})
//
func (c *Client) SetHeaders(headers map[string]string) *Client {
	for h, v := range headers {
		c.Header.Set(h, v)
	}

	return c
}

// SetCookie method sets a single cookie in the client instance.
// These cookies will be added to all the request raised from this client instance.
// 		resty.SetCookie(&http.Cookie{
// 					Name:"go-resty",
//					Value:"This is cookie value",
//					Path: "/",
// 					Domain: "sample.com",
// 					MaxAge: 36000,
// 					HttpOnly: true,
//					Secure: false,
// 				})
//
func (c *Client) SetCookie(hc *http.Cookie) *Client {
	c.Cookies = append(c.Cookies, hc)
	return c
}

// SetCookies method sets an array of cookies in the client instance.
// These cookies will be added to all the request raised from this client instance.
// 		cookies := make([]*http.Cookie, 0)
//
//		cookies = append(cookies, &http.Cookie{
// 					Name:"go-resty-1",
//					Value:"This is cookie 1 value",
//					Path: "/",
// 					Domain: "sample.com",
// 					MaxAge: 36000,
// 					HttpOnly: true,
//					Secure: false,
// 				})
//
//		cookies = append(cookies, &http.Cookie{
// 					Name:"go-resty-2",
//					Value:"This is cookie 2 value",
//					Path: "/",
// 					Domain: "sample.com",
// 					MaxAge: 36000,
// 					HttpOnly: true,
//					Secure: false,
// 				})
//
//		// Setting a cookies into resty
// 		resty.SetCookies(cookies)
//
func (c *Client) SetCookies(cs []*http.Cookie) *Client {
	c.Cookies = append(c.Cookies, cs...)
	return c
}

// SetQueryParam method sets single paramater and its value in the client instance.
// It will be formed as query string for the request. For example: `search=kitchen%20papers&size=large`
// in the URL after `?` mark. These query params will be added to all the request raised from
// this client instance. Also it can be overridden at request level Query Param options,
// see `resty.R().SetQueryParam` or `resty.R().SetQueryParams`.
// 		resty.
//			SetQueryParam("search", "kitchen papers").
//			SetQueryParam("size", "large")
//
func (c *Client) SetQueryParam(param, value string) *Client {
	c.QueryParam.Add(param, value)
	return c
}

// SetQueryParams method sets multiple paramaters and its values at one go in the client instance.
// It will be formed as query string for the request. For example: `search=kitchen%20papers&size=large`
// in the URL after `?` mark. These query params will be added to all the request raised from this
// client instance. Also it can be overridden at request level Query Param options,
// see `resty.R().SetQueryParams` or `resty.R().SetQueryParam`.
// 		resty.SetQueryParams(map[string]string{
//				"search": "kitchen papers",
//				"size": "large",
//			})
//
func (c *Client) SetQueryParams(params map[string]string) *Client {
	for p, v := range params {
		c.QueryParam.Add(p, v)
	}

	return c
}

// SetFormData method sets Form parameters and its values in the client instance.
// It's applicable only HTTP method `POST` and `PUT` and requets content type would be set as
// `application/x-www-form-urlencoded`. These form data will be added to all the request raised from
// this client instance. Also it can be overridden at request level form data, see `resty.R().SetFormData`.
// 		resty.SetFormData(map[string]string{
//				"access_token": "BC594900-518B-4F7E-AC75-BD37F019E08F",
//				"user_id": "3455454545",
//			})
//
func (c *Client) SetFormData(data map[string]string) *Client {
	for k, v := range data {
		c.FormData.Add(k, v)
	}

	return c
}

// SetBasicAuth method sets the basic authentication header in the HTTP request. For example -
//		Authorization: Basic <base64-encoded-value>
//
// For example: To set the header for username "go-resty" and password "welcome"
// 		resty.SetBasicAuth("go-resty", "welcome")
//
// This basic auth information gets added to all the request rasied from this client instance.
// Also it can be overriden or set one at the request level is supported, see `resty.R().SetBasicAuth`.
//
func (c *Client) SetBasicAuth(username, password string) *Client {
	c.UserInfo = &User{Username: username, Password: password}
	return c
}

// SetAuthToken method sets bearer auth token header in the HTTP request. For exmaple -
// 		Authorization: Bearer <auth-token-value-comes-here>
//
// For example: To set auth token BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F
//
// 		resty.SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F")
//
// This bearer auth token gets added to all the request rasied from this client instance.
// Also it can be overriden or set one at the request level is supported, see `resty.R().SetAuthToken`.
//
func (c *Client) SetAuthToken(token string) *Client {
	c.Token = token
	return c
}

// R method creates a request instance, its used for Get, Post, Put, Delete, Patch, Head and Options.
func (c *Client) R() *Request {
	r := &Request{
		URL:        "",
		Method:     "",
		QueryParam: url.Values{},
		FormData:   url.Values{},
		Header:     http.Header{},
		Body:       nil,
		Result:     nil,
		Error:      nil,
		RawRequest: nil,
		client:     c,
		bodyBuf:    nil,
	}

	return r
}

// OnBeforeRequest method sets request middleware into the before request chain.
// Its gets applied after default `go-resty` request middlewares and before request
// been sent from `go-resty` to host server.
// 		resty.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
//				// Now you have access to Client and Request instance
//				// manipulate it as per your need
//
//				return nil 	// if its success otherwise return error
//			})
//
func (c *Client) OnBeforeRequest(m func(*Client, *Request) error) *Client {
	c.beforeRequest[len(c.beforeRequest)-1] = m
	c.beforeRequest = append(c.beforeRequest, requestLogger)

	return c
}

// OnAfterResponse method sets response middleware into the after response chain.
// Once we receive response from host server, default `go-resty` response middleware
// gets applied and then user assigened response middlewares applied.
// 		resty.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
//				// Now you have access to Client and Response instance
//				// manipulate it as per your need
//
//				return nil 	// if its success otherwise return error
//			})
//
func (c *Client) OnAfterResponse(m func(*Client, *Response) error) *Client {
	c.afterResponse = append(c.afterResponse, m)
	return c
}

// SetDebug method enables the debug mode on `go-resty` client. Client logs details of every request and response.
// For `Request` it logs information such as HTTP verb, Relative URL path, Host, Headers, Body if it has one.
// For `Response` it logs information such as Status, Response Time, Headers, Body if it has one.
//
func (c *Client) SetDebug(d bool) *Client {
	c.Debug = d
	return c
}

// SetLogger method sets given writer for logging go-resty request and response details.
// Default is os.Stderr
// 		file, _ := os.OpenFile("/Users/jeeva/go-resty.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//
//		resty.SetLogger(file)
//
func (c *Client) SetLogger(w io.Writer) *Client {
	c.Log = getLogger(w)
	return c
}

// SetContentLength method enables the HTTP header `Content-Length` value for every request.
// By default go-resty won't set `Content-Length`.
// 		resty.SetContentLength(true)
//
// Also you have an option to enable for particular request. See `resty.R().SetContentLength`
//
func (c *Client) SetContentLength(l bool) *Client {
	c.setContentLength = l
	return c
}

// SetError method is to register the global or client common `Error` object into go-resty.
// It is used for automatic unmarshalling if response status code is greater than 399 and
// content type either JSON or XML. Can be pointer or non-pointer.
// 		resty.SetError(&Error{})
//		// OR
//		resty.SetError(Error{})
//
func (c *Client) SetError(err interface{}) *Client {
	c.Error = getType(err)
	return c
}

// SetRedirectPolicy method sets the client redirect poilicy. go-resty provides ready to use
// redirect policies. Wanna create one for yourself refer `redirect.go`.
//
//		resty.SetRedirectPolicy(FlexibleRedirectPolicy(20))
//
// 		// Need multiple redirect policies together
//		resty.SetRedirectPolicy(FlexibleRedirectPolicy(20), DomainCheckRedirectPolicy("host1.com", "host2.net"))
//
func (c *Client) SetRedirectPolicy(policies ...interface{}) *Client {
	for _, p := range policies {
		if _, ok := p.(RedirectPolicy); !ok {
			c.Log.Printf("ERORR: %v does not implement resty.RedirectPolicy (missing Apply method)",
				runtime.FuncForPC(reflect.ValueOf(p).Pointer()).Name())
		}
	}

	c.httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for _, p := range policies {
			err := p.(RedirectPolicy).Apply(req, via)
			if err != nil {
				return err
			}
		}
		return nil // looks good, go ahead
	}

	return c
}

// SetHTTPMode method sets go-resty mode into HTTP
func (c *Client) SetHTTPMode() *Client {
	return c.SetMode("http")
}

// SetRESTMode method sets go-resty mode into RESTful
func (c *Client) SetRESTMode() *Client {
	return c.SetMode("rest")
}

// SetMode method sets go-resty client mode to given value such as 'http' & 'rest'.
// 	RESTful:
//		- No Redirect
//		- Automatic response unmarshal if it is JSON or XML
//	HTML:
//		- Up to 10 Redirects
//		- No automatic unmarshall. Response will be treated as `response.String()`
//
// If you want more redirects, use FlexibleRedirectPolicy
//		resty.SetRedirectPolicy(FlexibleRedirectPolicy(20))
//
func (c *Client) SetMode(mode string) *Client {
	if mode == "http" {
		c.isHTTPMode = true
		c.SetRedirectPolicy(FlexibleRedirectPolicy(10))
		c.afterResponse = []func(*Client, *Response) error{
			responseLogger,
			saveResponseIntoFile,
		}
	} else { // RESTful
		c.isHTTPMode = false
		c.SetRedirectPolicy(NoRedirectPolicy())
		c.afterResponse = []func(*Client, *Response) error{
			responseLogger,
			parseResponseBody,
			saveResponseIntoFile,
		}
	}

	return c
}

// Mode method returns the current client mode. Typically its a "http" or "rest".
// Default is "rest"
func (c *Client) Mode() string {
	if c.isHTTPMode {
		return "http"
	}

	return "rest"
}

// SetTLSClientConfig method sets TLSClientConfig for underling client Transport.
//
// For example:
// 		// One can set custom root-certificate. Refer: http://golang.org/pkg/crypto/tls/#example_Dial
//		resty.SetTLSClientConfig(&tls.Config{ RootCAs: roots })
//
// 		// or One can disable security check (https)
//		resty.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
// Note: This method overwrites existing `TLSClientConfig`.
//
func (c *Client) SetTLSClientConfig(config *tls.Config) *Client {
	c.transport.TLSClientConfig = config
	return c
}

// SetTimeout method sets timeout for request raised from client
//		resty.SetTimeout(time.Duration(1 * time.Minute))
//
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	c.transport.Dial = func(network, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(network, addr, timeout)
		if err != nil {
			c.Log.Printf("ERROR [%v]", err)
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(timeout))
		return conn, nil
	}

	return c
}

// SetProxy method sets the Proxy URL and Port for resty client.
//		resty.SetProxy("http://proxyserver:8888")
//
// Alternative: Without this `SetProxy` method, you can also set Proxy via environment variable.
// By default `Go` uses setting from `HTTP_PROXY`.
//
func (c *Client) SetProxy(proxyURL string) *Client {
	if pURL, err := url.Parse(proxyURL); err == nil {
		c.transport.Proxy = http.ProxyURL(pURL)
	} else {
		c.Log.Printf("ERROR [%v]", err)
	}

	return c
}

// RemoveProxy method removes the proxy configuration from resty client
//		resty.RemoveProxy()
//
func (c *Client) RemoveProxy() *Client {
	c.transport.Proxy = nil
	return c
}

// SetCertificates method helps to set client certificates into resty conveniently.
func (c *Client) SetCertificates(certs ...tls.Certificate) *Client {
	config := c.getTLSConfig()
	config.Certificates = append(config.Certificates, certs...)

	return c
}

// SetRootCertificate method helps to add one or more root certificates into resty client
// 		resty.SetRootCertificate("/path/to/root/pemFile.pem")
//
func (c *Client) SetRootCertificate(pemFilePath string) *Client {
	rootPemData, err := ioutil.ReadFile(pemFilePath)
	if err != nil {
		c.Log.Printf("ERROR [%v]", err)
		return c
	}

	config := c.getTLSConfig()
	if config.RootCAs == nil {
		config.RootCAs = x509.NewCertPool()
	}

	config.RootCAs.AppendCertsFromPEM(rootPemData)

	return c
}

// SetOutputDirectory method sets output directory for saving HTTP response into file.
// If the output directory not exists then resty creates one. This setting is optional one,
// if you're planning using absoule path in `Request.SetOutput` and can used together.
// 		resty.SetOutputDirectory("/save/http/response/here")
//
func (c *Client) SetOutputDirectory(dirPath string) *Client {
	err := createDirectory(dirPath)
	if err != nil {
		c.Log.Printf("ERROR [%v]", err)
	}

	c.outputDirectory = dirPath

	return c
}

// executes the given `Request` object and returns response
func (c *Client) execute(req *Request) (*Response, error) {
	// Apply Request middleware
	var err error
	for _, f := range c.beforeRequest {
		err = f(c, req)
		if err != nil {
			return nil, err
		}
	}

	req.Time = time.Now()
	c.httpClient.Transport = c.transport
	resp, err := c.httpClient.Do(req.RawRequest)
	if err != nil {
		return nil, err
	}

	response := &Response{
		Request:     req,
		ReceivedAt:  time.Now(),
		RawResponse: resp,
	}

	if !req.isSaveResponse {
		defer resp.Body.Close()
		response.Body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		response.size = int64(len(response.Body))
	}

	// Apply Response middleware
	for _, f := range c.afterResponse {
		err = f(c, response)
		if err != nil {
			break
		}
	}

	return response, err
}

// enables a log prefix
func (c *Client) enableLogPrefix() {
	c.Log.SetFlags(log.LstdFlags)
	c.Log.SetPrefix("RESTY ")
}

// disables a log prefix
func (c *Client) disableLogPrefix() {
	c.Log.SetFlags(0)
	c.Log.SetPrefix("")
}

// getting TLS client config if not exists then create one
func (c *Client) getTLSConfig() *tls.Config {
	if c.transport.TLSClientConfig == nil {
		c.transport.TLSClientConfig = &tls.Config{}
	}

	return c.transport.TLSClientConfig
}

//
// Request
//

// Request type is used to compose and send individual request from client
// go-resty is provide option override client level settings such as
//		Auth Token, Basic Auth credentials, Header, Query Param, Form Data, Error object
// and also you can add more options for that particular request
//
type Request struct {
	URL        string
	Method     string
	QueryParam url.Values
	FormData   url.Values
	Header     http.Header
	UserInfo   *User
	Token      string
	Body       interface{}
	Result     interface{}
	Error      interface{}
	Time       time.Time
	RawRequest *http.Request

	client           *Client
	bodyBuf          *bytes.Buffer
	isMultiPart      bool
	isFormData       bool
	setContentLength bool
	isSaveResponse   bool
	outputFile       string
}

// SetHeader method is to set a single header field and its value in the current request.
// For Example: To set `Content-Type` and `Accept` as `application/json`.
// 		resty.R().
//			SetHeader("Content-Type", "application/json").
//			SetHeader("Accept", "application/json")
//
// Also you can override header value, which was set at client instance level.
//
func (r *Request) SetHeader(header, value string) *Request {
	r.Header.Set(header, value)
	return r
}

// SetHeaders method sets multiple headers field and its values at one go in the current request.
// For Example: To set `Content-Type` and `Accept` as `application/json`
//
// 		resty.R().
//			SetHeaders(map[string]string{
//				"Content-Type": "application/json",
//				"Accept": "application/json",
//			})
// Also you can override header value, which was set at client instance level.
//
func (r *Request) SetHeaders(headers map[string]string) *Request {
	for h, v := range headers {
		r.Header.Set(h, v)
	}

	return r
}

// SetQueryParam method sets single paramater and its value in the current request.
// It will be formed as query string for the request.
// For example: `search=kitchen%20papers&size=large` in the URL after `?` mark.
// 		resty.R().
//			SetQueryParam("search", "kitchen papers").
//			SetQueryParam("size", "large")
// Also you can override query params value, which was set at client instance level
//
func (r *Request) SetQueryParam(param, value string) *Request {
	r.QueryParam.Add(param, value)
	return r
}

// SetQueryParams method sets multiple paramaters and its values at one go in the current request.
// It will be formed as query string for the request.
// For example: `search=kitchen%20papers&size=large` in the URL after `?` mark.
// 		resty.R().
//			SetQueryParams(map[string]string{
//				"search": "kitchen papers",
//				"size": "large",
//			})
// Also you can override query params value, which was set at client instance level
//
func (r *Request) SetQueryParams(params map[string]string) *Request {
	for p, v := range params {
		r.QueryParam.Add(p, v)
	}

	return r
}

// SetQueryString method provides ability to use string as an input to set URL query string for the request.
//
// Using String as an input
// 		resty.R().
//			SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more")
//
func (r *Request) SetQueryString(query string) *Request {
	values, err := url.ParseQuery(strings.TrimSpace(query))
	if err == nil {
		for k := range values {
			r.QueryParam.Add(k, values.Get(k))
		}
	} else {
		r.client.Log.Printf("ERROR [%v]", err)
	}
	return r
}

// SetFormData method sets Form parameters and its values in the current request.
// It's applicable only HTTP method `POST` and `PUT` and requets content type would be set as
// `application/x-www-form-urlencoded`.
// 		resty.R().
// 			SetFormData(map[string]string{
//				"access_token": "BC594900-518B-4F7E-AC75-BD37F019E08F",
//				"user_id": "3455454545",
//			})
// Also you can override form data value, which was set at client instance level
//
func (r *Request) SetFormData(data map[string]string) *Request {
	for k, v := range data {
		r.FormData.Add(k, v)
	}

	return r
}

// SetBody method sets the request body for the request. It supports various realtime need easy.
// We can say its quite handy or powerful. Supported request body data types is `string`, `[]byte`,
// `struct` and `map`. Body value can be pointer or non-pointer. Automatic marshalling
// for JSON and XML content type, if it is `struct` or `map`.
//
// For Example:
//
// Struct as a body input, based on content type, it will be marshalled.
//		resty.R().
//			SetBody(User{
//				Username: "jeeva@myjeeva.com",
//				Password: "welcome2resty",
//			})
//
// Map as a body input, based on content type, it will be marshalled.
//		resty.R().
//			SetBody(map[string]interface{}{
//				"username": "jeeva@myjeeva.com",
//				"password": "welcome2resty",
//				"address": &Address{
//					Address1: "1111 This is my street",
//					Address2: "Apt 201",
//					City: "My City",
//					State: "My State",
//					ZipCode: 00000,
//				},
//			})
//
// String as a body input. Suitable for any need as a string input.
//		resty.R().
//			SetBody(`{
//				"username": "jeeva@getrightcare.com",
//				"password": "admin"
//			}`)
//
// []byte as a body input. Suitable for raw request such as file upload, serialize & deserialize, etc.
// 		resty.R().
//			SetBody([]byte("This is my raw request, sent as-is"))
//
func (r *Request) SetBody(body interface{}) *Request {
	r.Body = body
	return r
}

// SetResult method is to register the response `Result` object for automatic unmarshalling in the RESTful mode
// if response status code is between 200 and 299 and content type either JSON or XML.
//
// Note: Result object can be pointer or non-pointer.
//		resty.R().SetResult(&AuthToken{})
//		// OR
//		resty.R().SetResult(AuthToken{})
//
// Accessing a result value
//		response.Result().(*AuthToken)
//
func (r *Request) SetResult(res interface{}) *Request {
	r.Result = getPointer(res)
	return r
}

// SetError method is to register the request `Error` object for automatic unmarshalling in the RESTful mode
// if response status code is greater than 399 and content type either JSON or XML.
//
// Note: Error object can be pointer or non-pointer.
// 		resty.R().SetError(&AuthError{})
//		// OR
//		resty.R().SetError(AuthError{})
//
// Accessing a error value
//		response.Error().(*AuthError)
//
func (r *Request) SetError(err interface{}) *Request {
	r.Error = getPointer(err)
	return r
}

// SetFile method is to set single file field name and its path for multipart upload.
//	resty.R().
//		SetFile("my_file", "/Users/jeeva/Gas Bill - Sep.pdf")
//
func (r *Request) SetFile(param, filePath string) *Request {
	r.FormData.Set("@"+param, filePath)
	r.isMultiPart = true

	return r
}

// SetFiles method is to set multiple file field name and its path for multipart upload.
//	resty.R().
//		SetFiles(map[string]string{
//				"my_file1": "/Users/jeeva/Gas Bill - Sep.pdf",
//				"my_file2": "/Users/jeeva/Electricity Bill - Sep.pdf",
//				"my_file3": "/Users/jeeva/Water Bill - Sep.pdf",
//			})
//
func (r *Request) SetFiles(files map[string]string) *Request {
	for f, fp := range files {
		r.FormData.Set("@"+f, fp)
	}
	r.isMultiPart = true

	return r
}

// SetContentLength method sets the HTTP header `Content-Length` value for current request.
// By default go-resty won't set `Content-Length`. Also you have an option to enable for every
// request. See `resty.SetContentLength`
// 		resty.R().SetContentLength(true)
//
func (r *Request) SetContentLength(l bool) *Request {
	r.setContentLength = true

	return r
}

// SetBasicAuth method sets the basic authentication header in the current HTTP request.
// For Header example:
//		Authorization: Basic <base64-encoded-value>
//
// To set the header for username "go-resty" and password "welcome"
// 		resty.R().SetBasicAuth("go-resty", "welcome")
//
// This method overrides the credentials set by method `resty.SetBasicAuth`.
//
func (r *Request) SetBasicAuth(username, password string) *Request {
	r.UserInfo = &User{Username: username, Password: password}
	return r
}

// SetAuthToken method sets bearer auth token header in the current HTTP request. For Header exmaple:
// 		Authorization: Bearer <auth-token-value-comes-here>
//
// For example: To set auth token BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F
//
// 		resty.R().SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F")
//
// This method overrides the Auth token set by method `resty.SetAuthToken`.
//
func (r *Request) SetAuthToken(token string) *Request {
	r.Token = token
	return r
}

// SetOutput method sets the output file for current HTTP request. Current HTTP response will be
// saved into given file. It is similar to `curl -o` flag. Absoulte path or relative path can be used.
// If is it relative path then output file goes under the output directory, as mentioned
// in the `Client.SetOutputDirectory`.
// 		resty.R().
// 			SetOutput("/Users/jeeva/Downloads/ReplyWithHeader-v5.1-beta.zip").
// 			Get("http://bit.ly/1LouEKr")
//
// Note: In this scenario `Response.Body` might be nil.
func (r *Request) SetOutput(file string) *Request {
	r.outputFile = file
	r.isSaveResponse = true
	return r
}

//
// HTTP verb method starts here
//

// Get method does GET HTTP request. It's defined in section 4.3.1 of RFC7231.
func (r *Request) Get(url string) (*Response, error) {
	return r.Execute(GET, url)
}

// Head method does HEAD HTTP request. It's defined in section 4.3.2 of RFC7231.
func (r *Request) Head(url string) (*Response, error) {
	return r.Execute(HEAD, url)
}

// Post method does POST HTTP request. It's defined in section 4.3.3 of RFC7231.
func (r *Request) Post(url string) (*Response, error) {
	return r.Execute(POST, url)
}

// Put method does PUT HTTP request. It's defined in section 4.3.4 of RFC7231.
func (r *Request) Put(url string) (*Response, error) {
	return r.Execute(PUT, url)
}

// Delete method does DELETE HTTP request. It's defined in section 4.3.5 of RFC7231.
func (r *Request) Delete(url string) (*Response, error) {
	return r.Execute(DELETE, url)
}

// Options method does OPTIONS HTTP request. It's defined in section 4.3.7 of RFC7231.
func (r *Request) Options(url string) (*Response, error) {
	return r.Execute(OPTIONS, url)
}

// Patch method does PATCH HTTP request. It's defined in section 2 of RFC5789.
func (r *Request) Patch(url string) (*Response, error) {
	return r.Execute(PATCH, url)
}

// Execute method performs the HTTP request with given HTTP method and URL
// for current `Request`.
// 		resp, err := resty.R().Execute(resty.GET, "http://httpbin.org/get")
//
func (r *Request) Execute(method, url string) (*Response, error) {
	if r.isMultiPart && !(method == POST || method == PUT) {
		return nil, fmt.Errorf("Multipart content is not allowed in HTTP verb [%v]", method)
	}

	r.Method = method
	r.URL = url

	return r.client.execute(r)
}

//
// Response
//

// Response is an object represents executed request and its values.
type Response struct {
	Body        []byte
	ReceivedAt  time.Time
	Request     *Request
	RawResponse *http.Response

	size int64
}

// Status method returns the HTTP status string for the executed request.
//	For example: 200 OK
func (r *Response) Status() string {
	return r.RawResponse.Status
}

// StatusCode method returns the HTTP status code for the executed request.
//	For example: 200
func (r *Response) StatusCode() int {
	return r.RawResponse.StatusCode
}

// Result method returns the response value as an object if it has one
func (r *Response) Result() interface{} {
	return r.Request.Result
}

// Error method returns the error object if it has one
func (r *Response) Error() interface{} {
	return r.Request.Error
}

// Header method returns the response headers
func (r *Response) Header() http.Header {
	return r.RawResponse.Header
}

// Cookies method to access all the response cookies
func (r *Response) Cookies() []*http.Cookie {
	return r.RawResponse.Cookies()
}

// String method returns the body of the server response as String.
func (r *Response) String() string {
	if r.Body == nil {
		return ""
	}

	return strings.TrimSpace(string(r.Body))
}

// Time method returns the time of HTTP response time that from request we sent and received a request.
// See `response.ReceivedAt` to know when client recevied response and see `response.Request.Time` to know
// when client sent a request.
func (r *Response) Time() time.Duration {
	return r.ReceivedAt.Sub(r.Request.Time)
}

// Size method returns the HTTP response size in bytes. Ya, you can relay on HTTP `Content-Length` header,
// however it won't be good for chucked transfer/compressed response. Since Resty calculates response size
// at the client end. You will get actual size of the http response.
func (r *Response) Size() int64 {
	return r.size
}

//
// Helper methods
//

// IsStringEmpty method tells whether given string is empty or not
func IsStringEmpty(str string) bool {
	return (len(strings.TrimSpace(str)) == 0)
}

// DetectContentType method is used to figure out `Request.Body` content type for request header
func DetectContentType(body interface{}) string {
	contentType := plainTextType
	switch getBaseKind(body) {
	case reflect.Struct, reflect.Map:
		contentType = jsonContentType
	case reflect.String:
		contentType = plainTextType
	default:
		contentType = http.DetectContentType(body.([]byte))
	}

	return contentType
}

// IsJSONType method is to check JSON content type or not
func IsJSONType(ct string) bool {
	return jsonCheck.MatchString(ct)
}

// IsXMLType method is to check XML content type or not
func IsXMLType(ct string) bool {
	return xmlCheck.MatchString(ct)
}

// Unmarshal content into object from JSON or XML
func Unmarshal(ct string, b []byte, d interface{}) (err error) {
	if IsJSONType(ct) {
		err = json.Unmarshal(b, d)
	} else if IsXMLType(ct) {
		err = xml.Unmarshal(b, d)
	}

	return
}

func getLogger(w io.Writer) *log.Logger {
	return log.New(w, "RESTY ", log.LstdFlags)
}

func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

func getRequestBodyString(r *Request) (body string) {
	body = "***** NO CONTENT *****"
	if isPayloadSupported(r.Method) {
		// multipart/form-data OR form data
		if r.isMultiPart || r.isFormData {
			body = string(r.bodyBuf.Bytes())
			return
		}

		// request body data
		if r.Body != nil {
			contentType := r.Header.Get(hdrContentTypeKey)
			var prtBodyBytes []byte
			var err error
			kind := reflect.ValueOf(r.Body).Kind()
			if IsJSONType(contentType) && (kind == reflect.Struct || kind == reflect.Map) {
				prtBodyBytes, err = json.MarshalIndent(&r.Body, "", "   ")
			} else if IsXMLType(contentType) && (kind == reflect.Struct) {
				prtBodyBytes, err = xml.MarshalIndent(&r.Body, "", "   ")
			} else if b, ok := r.Body.(string); ok {
				if IsJSONType(contentType) {
					bodyBytes := []byte(b)
					var out bytes.Buffer
					if err = json.Indent(&out, bodyBytes, "", "   "); err == nil {
						prtBodyBytes = out.Bytes()
					}
				} else {
					body = b
					return
				}
			} else if b, ok := r.Body.([]byte); ok {
				body = base64.StdEncoding.EncodeToString(b)
			}

			if prtBodyBytes != nil {
				body = string(prtBodyBytes)
			}
		}

	}

	return
}

func getResponseBodyString(res *Response) string {
	bodyStr := "***** NO CONTENT *****"
	if res.Body != nil {
		ct := res.Header().Get(hdrContentTypeKey)
		if IsJSONType(ct) {
			var out bytes.Buffer
			if err := json.Indent(&out, res.Body, "", "   "); err == nil {
				bodyStr = string(out.Bytes())
			}
		} else {
			str := res.String()
			if !IsStringEmpty(str) {
				bodyStr = str
			}
		}
	}

	return bodyStr
}

func getPointer(v interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		return v
	}
	return reflect.New(vv.Type()).Interface()
}

func isPayloadSupported(m string) bool {
	return (m == POST || m == PUT || m == DELETE || m == PATCH)
}

func getBaseKind(v interface{}) reflect.Kind {
	return getType(v).Kind()
}

func getType(v interface{}) reflect.Type {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		return vv.Elem().Type()
	}
	return vv.Type()
}

func createDirectory(dir string) (err error) {
	if _, err = os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(dir, 0755); err != nil {
				return
			}
		}
	}
	return
}
