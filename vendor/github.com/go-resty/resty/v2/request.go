// Copyright (c) 2015-2019 Jeevanandam M (jeeva@myjeeva.com), All rights reserved.
// resty source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package resty

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Request struct and methods
//_______________________________________________________________________

// Request struct is used to compose and fire individual request from
// resty client. Request provides an options to override client level
// settings and also an options for the request composition.
type Request struct {
	URL        string
	Method     string
	Token      string
	QueryParam url.Values
	FormData   url.Values
	Header     http.Header
	Time       time.Time
	Body       interface{}
	Result     interface{}
	Error      interface{}
	RawRequest *http.Request
	SRV        *SRVRecord
	UserInfo   *User
	Cookies    []*http.Cookie

	isMultiPart         bool
	isFormData          bool
	setContentLength    bool
	isSaveResponse      bool
	notParseResponse    bool
	jsonEscapeHTML      bool
	trace               bool
	outputFile          string
	fallbackContentType string
	ctx                 context.Context
	pathParams          map[string]string
	values              map[string]interface{}
	client              *Client
	bodyBuf             *bytes.Buffer
	clientTrace         *clientTrace
	multipartFiles      []*File
	multipartFields     []*MultipartField
}

// Context method returns the Context if its already set in request
// otherwise it creates new one using `context.Background()`.
func (r *Request) Context() context.Context {
	if r.ctx == nil {
		return context.Background()
	}
	return r.ctx
}

// SetContext method sets the context.Context for current Request. It allows
// to interrupt the request execution if ctx.Done() channel is closed.
// See https://blog.golang.org/context article and the "context" package
// documentation.
func (r *Request) SetContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

// SetHeader method is to set a single header field and its value in the current request.
//
// For Example: To set `Content-Type` and `Accept` as `application/json`.
// 		client.R().
//			SetHeader("Content-Type", "application/json").
//			SetHeader("Accept", "application/json")
//
// Also you can override header value, which was set at client instance level.
func (r *Request) SetHeader(header, value string) *Request {
	r.Header.Set(header, value)
	return r
}

// SetHeaders method sets multiple headers field and its values at one go in the current request.
//
// For Example: To set `Content-Type` and `Accept` as `application/json`
//
// 		client.R().
//			SetHeaders(map[string]string{
//				"Content-Type": "application/json",
//				"Accept": "application/json",
//			})
// Also you can override header value, which was set at client instance level.
func (r *Request) SetHeaders(headers map[string]string) *Request {
	for h, v := range headers {
		r.SetHeader(h, v)
	}
	return r
}

// SetQueryParam method sets single parameter and its value in the current request.
// It will be formed as query string for the request.
//
// For Example: `search=kitchen%20papers&size=large` in the URL after `?` mark.
// 		client.R().
//			SetQueryParam("search", "kitchen papers").
//			SetQueryParam("size", "large")
// Also you can override query params value, which was set at client instance level.
func (r *Request) SetQueryParam(param, value string) *Request {
	r.QueryParam.Set(param, value)
	return r
}

// SetQueryParams method sets multiple parameters and its values at one go in the current request.
// It will be formed as query string for the request.
//
// For Example: `search=kitchen%20papers&size=large` in the URL after `?` mark.
// 		client.R().
//			SetQueryParams(map[string]string{
//				"search": "kitchen papers",
//				"size": "large",
//			})
// Also you can override query params value, which was set at client instance level.
func (r *Request) SetQueryParams(params map[string]string) *Request {
	for p, v := range params {
		r.SetQueryParam(p, v)
	}
	return r
}

// SetQueryParamsFromValues method appends multiple parameters with multi-value
// (`url.Values`) at one go in the current request. It will be formed as
// query string for the request.
//
// For Example: `status=pending&status=approved&status=open` in the URL after `?` mark.
// 		client.R().
//			SetQueryParamsFromValues(url.Values{
//				"status": []string{"pending", "approved", "open"},
//			})
// Also you can override query params value, which was set at client instance level.
func (r *Request) SetQueryParamsFromValues(params url.Values) *Request {
	for p, v := range params {
		for _, pv := range v {
			r.QueryParam.Add(p, pv)
		}
	}
	return r
}

// SetQueryString method provides ability to use string as an input to set URL query string for the request.
//
// Using String as an input
// 		client.R().
//			SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more")
func (r *Request) SetQueryString(query string) *Request {
	params, err := url.ParseQuery(strings.TrimSpace(query))
	if err == nil {
		for p, v := range params {
			for _, pv := range v {
				r.QueryParam.Add(p, pv)
			}
		}
	} else {
		r.client.log.Errorf("%v", err)
	}
	return r
}

// SetFormData method sets Form parameters and their values in the current request.
// It's applicable only HTTP method `POST` and `PUT` and requests content type would be set as
// `application/x-www-form-urlencoded`.
// 		client.R().
// 			SetFormData(map[string]string{
//				"access_token": "BC594900-518B-4F7E-AC75-BD37F019E08F",
//				"user_id": "3455454545",
//			})
// Also you can override form data value, which was set at client instance level.
func (r *Request) SetFormData(data map[string]string) *Request {
	for k, v := range data {
		r.FormData.Set(k, v)
	}
	return r
}

// SetFormDataFromValues method appends multiple form parameters with multi-value
// (`url.Values`) at one go in the current request.
// 		client.R().
//			SetFormDataFromValues(url.Values{
//				"search_criteria": []string{"book", "glass", "pencil"},
//			})
// Also you can override form data value, which was set at client instance level.
func (r *Request) SetFormDataFromValues(data url.Values) *Request {
	for k, v := range data {
		for _, kv := range v {
			r.FormData.Add(k, kv)
		}
	}
	return r
}

// SetBody method sets the request body for the request. It supports various realtime needs as easy.
// We can say its quite handy or powerful. Supported request body data types is `string`,
// `[]byte`, `struct`, `map`, `slice` and `io.Reader`. Body value can be pointer or non-pointer.
// Automatic marshalling for JSON and XML content type, if it is `struct`, `map`, or `slice`.
//
// Note: `io.Reader` is processed as bufferless mode while sending request.
//
// For Example: Struct as a body input, based on content type, it will be marshalled.
//		client.R().
//			SetBody(User{
//				Username: "jeeva@myjeeva.com",
//				Password: "welcome2resty",
//			})
//
// Map as a body input, based on content type, it will be marshalled.
//		client.R().
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
//		client.R().
//			SetBody(`{
//				"username": "jeeva@getrightcare.com",
//				"password": "admin"
//			}`)
//
// []byte as a body input. Suitable for raw request such as file upload, serialize & deserialize, etc.
// 		client.R().
//			SetBody([]byte("This is my raw request, sent as-is"))
func (r *Request) SetBody(body interface{}) *Request {
	r.Body = body
	return r
}

// SetResult method is to register the response `Result` object for automatic unmarshalling for the request,
// if response status code is between 200 and 299 and content type either JSON or XML.
//
// Note: Result object can be pointer or non-pointer.
//		client.R().SetResult(&AuthToken{})
//		// OR
//		client.R().SetResult(AuthToken{})
//
// Accessing a result value from response instance.
//		response.Result().(*AuthToken)
func (r *Request) SetResult(res interface{}) *Request {
	r.Result = getPointer(res)
	return r
}

// SetError method is to register the request `Error` object for automatic unmarshalling for the request,
// if response status code is greater than 399 and content type either JSON or XML.
//
// Note: Error object can be pointer or non-pointer.
// 		client.R().SetError(&AuthError{})
//		// OR
//		client.R().SetError(AuthError{})
//
// Accessing a error value from response instance.
//		response.Error().(*AuthError)
func (r *Request) SetError(err interface{}) *Request {
	r.Error = getPointer(err)
	return r
}

// SetFile method is to set single file field name and its path for multipart upload.
//	client.R().
//		SetFile("my_file", "/Users/jeeva/Gas Bill - Sep.pdf")
func (r *Request) SetFile(param, filePath string) *Request {
	r.isMultiPart = true
	r.FormData.Set("@"+param, filePath)
	return r
}

// SetFiles method is to set multiple file field name and its path for multipart upload.
//	client.R().
//		SetFiles(map[string]string{
//				"my_file1": "/Users/jeeva/Gas Bill - Sep.pdf",
//				"my_file2": "/Users/jeeva/Electricity Bill - Sep.pdf",
//				"my_file3": "/Users/jeeva/Water Bill - Sep.pdf",
//			})
func (r *Request) SetFiles(files map[string]string) *Request {
	r.isMultiPart = true
	for f, fp := range files {
		r.FormData.Set("@"+f, fp)
	}
	return r
}

// SetFileReader method is to set single file using io.Reader for multipart upload.
//	client.R().
//		SetFileReader("profile_img", "my-profile-img.png", bytes.NewReader(profileImgBytes)).
//		SetFileReader("notes", "user-notes.txt", bytes.NewReader(notesBytes))
func (r *Request) SetFileReader(param, fileName string, reader io.Reader) *Request {
	r.isMultiPart = true
	r.multipartFiles = append(r.multipartFiles, &File{
		Name:      fileName,
		ParamName: param,
		Reader:    reader,
	})
	return r
}

// SetMultipartField method is to set custom data using io.Reader for multipart upload.
func (r *Request) SetMultipartField(param, fileName, contentType string, reader io.Reader) *Request {
	r.isMultiPart = true
	r.multipartFields = append(r.multipartFields, &MultipartField{
		Param:       param,
		FileName:    fileName,
		ContentType: contentType,
		Reader:      reader,
	})
	return r
}

// SetMultipartFields method is to set multiple data fields using io.Reader for multipart upload.
//
// For Example:
// 	client.R().SetMultipartFields(
// 		&resty.MultipartField{
//			Param:       "uploadManifest1",
//			FileName:    "upload-file-1.json",
//			ContentType: "application/json",
//			Reader:      strings.NewReader(`{"input": {"name": "Uploaded document 1", "_filename" : ["file1.txt"]}}`),
//		},
//		&resty.MultipartField{
//			Param:       "uploadManifest2",
//			FileName:    "upload-file-2.json",
//			ContentType: "application/json",
//			Reader:      strings.NewReader(`{"input": {"name": "Uploaded document 2", "_filename" : ["file2.txt"]}}`),
//		})
//
// If you have slice already, then simply call-
// 	client.R().SetMultipartFields(fields...)
func (r *Request) SetMultipartFields(fields ...*MultipartField) *Request {
	r.isMultiPart = true
	r.multipartFields = append(r.multipartFields, fields...)
	return r
}

// SetContentLength method sets the HTTP header `Content-Length` value for current request.
// By default Resty won't set `Content-Length`. Also you have an option to enable for every
// request.
//
// See `Client.SetContentLength`
// 		client.R().SetContentLength(true)
func (r *Request) SetContentLength(l bool) *Request {
	r.setContentLength = true
	return r
}

// SetBasicAuth method sets the basic authentication header in the current HTTP request.
//
// For Example:
//		Authorization: Basic <base64-encoded-value>
//
// To set the header for username "go-resty" and password "welcome"
// 		client.R().SetBasicAuth("go-resty", "welcome")
//
// This method overrides the credentials set by method `Client.SetBasicAuth`.
func (r *Request) SetBasicAuth(username, password string) *Request {
	r.UserInfo = &User{Username: username, Password: password}
	return r
}

// SetAuthToken method sets bearer auth token header in the current HTTP request. Header example:
// 		Authorization: Bearer <auth-token-value-comes-here>
//
// For Example: To set auth token BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F
//
// 		client.R().SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F")
//
// This method overrides the Auth token set by method `Client.SetAuthToken`.
func (r *Request) SetAuthToken(token string) *Request {
	r.Token = token
	return r
}

// SetOutput method sets the output file for current HTTP request. Current HTTP response will be
// saved into given file. It is similar to `curl -o` flag. Absolute path or relative path can be used.
// If is it relative path then output file goes under the output directory, as mentioned
// in the `Client.SetOutputDirectory`.
// 		client.R().
// 			SetOutput("/Users/jeeva/Downloads/ReplyWithHeader-v5.1-beta.zip").
// 			Get("http://bit.ly/1LouEKr")
//
// Note: In this scenario `Response.Body` might be nil.
func (r *Request) SetOutput(file string) *Request {
	r.outputFile = file
	r.isSaveResponse = true
	return r
}

// SetSRV method sets the details to query the service SRV record and execute the
// request.
// 		client.R().
//			SetSRV(SRVRecord{"web", "testservice.com"}).
//			Get("/get")
func (r *Request) SetSRV(srv *SRVRecord) *Request {
	r.SRV = srv
	return r
}

// SetDoNotParseResponse method instructs `Resty` not to parse the response body automatically.
// Resty exposes the raw response body as `io.ReadCloser`. Also do not forget to close the body,
// otherwise you might get into connection leaks, no connection reuse.
//
// Note: Response middlewares are not applicable, if you use this option. Basically you have
// taken over the control of response parsing from `Resty`.
func (r *Request) SetDoNotParseResponse(parse bool) *Request {
	r.notParseResponse = parse
	return r
}

// SetPathParams method sets multiple URL path key-value pairs at one go in the
// Resty current request instance.
// 		client.R().SetPathParams(map[string]string{
// 		   "userId": "sample@sample.com",
// 		   "subAccountId": "100002",
// 		})
//
// 		Result:
// 		   URL - /v1/users/{userId}/{subAccountId}/details
// 		   Composed URL - /v1/users/sample@sample.com/100002/details
// It replace the value of the key while composing request URL. Also you can
// override Path Params value, which was set at client instance level.
func (r *Request) SetPathParams(params map[string]string) *Request {
	for p, v := range params {
		r.pathParams[p] = v
	}
	return r
}

// ExpectContentType method allows to provide fallback `Content-Type` for automatic unmarshalling
// when `Content-Type` response header is unavailable.
func (r *Request) ExpectContentType(contentType string) *Request {
	r.fallbackContentType = contentType
	return r
}

// SetJSONEscapeHTML method is to enable/disable the HTML escape on JSON marshal.
//
// Note: This option only applicable to standard JSON Marshaller.
func (r *Request) SetJSONEscapeHTML(b bool) *Request {
	r.jsonEscapeHTML = b
	return r
}

// SetCookie method appends a single cookie in the current request instance.
// 		client.R().SetCookie(&http.Cookie{
// 					Name:"go-resty",
// 					Value:"This is cookie value",
// 				})
//
// Note: Method appends the Cookie value into existing Cookie if already existing.
//
// Since v2.1.0
func (r *Request) SetCookie(hc *http.Cookie) *Request {
	r.Cookies = append(r.Cookies, hc)
	return r
}

// SetCookies method sets an array of cookies in the current request instance.
// 		cookies := []*http.Cookie{
// 			&http.Cookie{
// 				Name:"go-resty-1",
// 				Value:"This is cookie 1 value",
// 			},
// 			&http.Cookie{
// 				Name:"go-resty-2",
// 				Value:"This is cookie 2 value",
// 			},
// 		}
//
//		// Setting a cookies into resty's current request
// 		client.R().SetCookies(cookies)
//
// Note: Method appends the Cookie value into existing Cookie if already existing.
//
// Since v2.1.0
func (r *Request) SetCookies(rs []*http.Cookie) *Request {
	r.Cookies = append(r.Cookies, rs...)
	return r
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// HTTP request tracing
//_______________________________________________________________________

// EnableTrace method enables trace for the current request
// using `httptrace.ClientTrace` and provides insights.
//
// 		client := resty.New()
//
// 		resp, err := client.R().EnableTrace().Get("https://httpbin.org/get")
// 		fmt.Println("Error:", err)
// 		fmt.Println("Trace Info:", resp.Request.TraceInfo())
//
// See `Client.EnableTrace` available too to get trace info for all requests.
//
// Since v2.0.0
func (r *Request) EnableTrace() *Request {
	r.trace = true
	return r
}

// TraceInfo method returns the trace info for the request.
//
// Since v2.0.0
func (r *Request) TraceInfo() TraceInfo {
	ct := r.clientTrace
	return TraceInfo{
		DNSLookup:     ct.dnsDone.Sub(ct.dnsStart),
		ConnTime:      ct.gotConn.Sub(ct.getConn),
		TLSHandshake:  ct.tlsHandshakeDone.Sub(ct.tlsHandshakeStart),
		ServerTime:    ct.gotFirstResponseByte.Sub(ct.wroteRequest),
		ResponseTime:  ct.endTime.Sub(ct.gotFirstResponseByte),
		TotalTime:     ct.endTime.Sub(ct.getConn),
		IsConnReused:  ct.gotConnInfo.Reused,
		IsConnWasIdle: ct.gotConnInfo.WasIdle,
		ConnIdleTime:  ct.gotConnInfo.IdleTime,
	}
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// HTTP verb method starts here
//_______________________________________________________________________

// Get method does GET HTTP request. It's defined in section 4.3.1 of RFC7231.
func (r *Request) Get(url string) (*Response, error) {
	return r.Execute(MethodGet, url)
}

// Head method does HEAD HTTP request. It's defined in section 4.3.2 of RFC7231.
func (r *Request) Head(url string) (*Response, error) {
	return r.Execute(MethodHead, url)
}

// Post method does POST HTTP request. It's defined in section 4.3.3 of RFC7231.
func (r *Request) Post(url string) (*Response, error) {
	return r.Execute(MethodPost, url)
}

// Put method does PUT HTTP request. It's defined in section 4.3.4 of RFC7231.
func (r *Request) Put(url string) (*Response, error) {
	return r.Execute(MethodPut, url)
}

// Delete method does DELETE HTTP request. It's defined in section 4.3.5 of RFC7231.
func (r *Request) Delete(url string) (*Response, error) {
	return r.Execute(MethodDelete, url)
}

// Options method does OPTIONS HTTP request. It's defined in section 4.3.7 of RFC7231.
func (r *Request) Options(url string) (*Response, error) {
	return r.Execute(MethodOptions, url)
}

// Patch method does PATCH HTTP request. It's defined in section 2 of RFC5789.
func (r *Request) Patch(url string) (*Response, error) {
	return r.Execute(MethodPatch, url)
}

// Execute method performs the HTTP request with given HTTP method and URL
// for current `Request`.
// 		resp, err := client.R().Execute(resty.GET, "http://httpbin.org/get")
func (r *Request) Execute(method, url string) (*Response, error) {
	var addrs []*net.SRV
	var err error

	if r.isMultiPart && !(method == MethodPost || method == MethodPut || method == MethodPatch) {
		return nil, fmt.Errorf("multipart content is not allowed in HTTP verb [%v]", method)
	}

	if r.SRV != nil {
		_, addrs, err = net.LookupSRV(r.SRV.Service, "tcp", r.SRV.Domain)
		if err != nil {
			return nil, err
		}
	}

	r.Method = method
	r.URL = r.selectAddr(addrs, url, 0)

	if r.client.RetryCount == 0 {
		return r.client.execute(r)
	}

	var resp *Response
	attempt := 0
	err = Backoff(
		func() (*Response, error) {
			attempt++

			r.URL = r.selectAddr(addrs, url, attempt)

			resp, err = r.client.execute(r)
			if err != nil {
				r.client.log.Errorf("%v, Attempt %v", err, attempt)
			}

			return resp, err
		},
		Retries(r.client.RetryCount),
		WaitTime(r.client.RetryWaitTime),
		MaxWaitTime(r.client.RetryMaxWaitTime),
		RetryConditions(r.client.RetryConditions),
	)

	return resp, err
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// SRVRecord struct
//_______________________________________________________________________

// SRVRecord struct holds the data to query the SRV record for the
// following service.
type SRVRecord struct {
	Service string
	Domain  string
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Request Unexported methods
//_______________________________________________________________________

func (r *Request) fmtBodyString() (body string) {
	body = "***** NO CONTENT *****"
	if isPayloadSupported(r.Method, r.client.AllowGetMethodPayload) {
		if _, ok := r.Body.(io.Reader); ok {
			body = "***** BODY IS io.Reader *****"
			return
		}

		// multipart or form-data
		if r.isMultiPart || r.isFormData {
			body = r.bodyBuf.String()
			return
		}

		// request body data
		if r.Body == nil {
			return
		}
		var prtBodyBytes []byte
		var err error

		contentType := r.Header.Get(hdrContentTypeKey)
		kind := kindOf(r.Body)
		if canJSONMarshal(contentType, kind) {
			prtBodyBytes, err = json.MarshalIndent(&r.Body, "", "   ")
		} else if IsXMLType(contentType) && (kind == reflect.Struct) {
			prtBodyBytes, err = xml.MarshalIndent(&r.Body, "", "   ")
		} else if b, ok := r.Body.(string); ok {
			if IsJSONType(contentType) {
				bodyBytes := []byte(b)
				out := acquireBuffer()
				defer releaseBuffer(out)
				if err = json.Indent(out, bodyBytes, "", "   "); err == nil {
					prtBodyBytes = out.Bytes()
				}
			} else {
				body = b
				return
			}
		} else if b, ok := r.Body.([]byte); ok {
			body = base64.StdEncoding.EncodeToString(b)
		}

		if prtBodyBytes != nil && err == nil {
			body = string(prtBodyBytes)
		}
	}

	return
}

func (r *Request) selectAddr(addrs []*net.SRV, path string, attempt int) string {
	if addrs == nil {
		return path
	}

	idx := attempt % len(addrs)
	domain := strings.TrimRight(addrs[idx].Target, ".")
	path = strings.TrimLeft(path, "/")

	return fmt.Sprintf("%s://%s:%d/%s", r.client.scheme, domain, addrs[idx].Port, path)
}

func (r *Request) initValuesMap() {
	if r.values == nil {
		r.values = make(map[string]interface{})
	}
}

var noescapeJSONMarshal = func(v interface{}) ([]byte, error) {
	buf := acquireBuffer()
	defer releaseBuffer(buf)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	return buf.Bytes(), err
}
