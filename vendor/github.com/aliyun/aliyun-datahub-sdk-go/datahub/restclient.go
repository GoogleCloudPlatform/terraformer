package datahub

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	HttpHeaderDate          = "Date"
	HttpHeaderUserAgent     = "User-Agent"
	HttpHeaderContentType   = "Content-Type"
	HttpHeaderContentLength = "Content-Length"
	HttpHeaderAuthorization = "Authorization"
	HttpHeaderSecurityToken = "x-datahub-security-token"
)

const (
	DatahubHeadersPrefix = "x-datahub-"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the level severity or above.
	dev := strings.ToLower(os.Getenv("GODATAHUB_DEV"))
	switch dev {
	case "true":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}
}

// DialContextFn was defined to make code more readable.
type DialContextFn func(ctx context.Context, network, address string) (net.Conn, error)

// TraceDialContext implements our own dialer in order to trace conn info.
func TraceDialContext(ctimeout time.Duration) DialContextFn {
	dialer := &net.Dialer{
		Timeout:   ctimeout,
		KeepAlive: ctimeout,
	}
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		conn, err := dialer.DialContext(ctx, network, addr)
		if err != nil {
			return nil, err
		}

		log.Debug("connect done, use", conn.LocalAddr().String())
		return conn, nil
	}
}

// DefaultHttpClient returns a default HTTP client with sensible values.
func DefaultHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext:           TraceDialContext(10 * time.Second),
			Proxy:                 http.ProxyFromEnvironment,
			MaxIdleConns:          100,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}
}

// DefaultUserAgent returns a default user agent
func DefaultUserAgent() string {
	return fmt.Sprintf("godatahub/%s golang/%s %s", DATAHUB_SDK_VERSION, runtime.Version(), runtime.GOOS)
}

// RestClient rest客户端
type RestClient struct {
	// Endpoint datahub服务的endpint
	Endpoint string
	// Useragent user agent
	Useragent string
	// HttpClient http client
	HttpClient *http.Client
	// Account
	Account Account
}

// NewRestClient create a new rest client
func NewRestClient(endpoint string, useragent string, httpclient *http.Client, account Account) *RestClient {
	if strings.HasSuffix(endpoint, "/") {
		endpoint = endpoint[0 : len(endpoint)-1]
	}
	return &RestClient{
		Endpoint:   endpoint,
		Useragent:  useragent,
		HttpClient: httpclient,
		Account:    account,
	}
}

// Get send HTTP Get method request
func (client *RestClient) Get(resource string, model RestModel) error {
	return client.request(http.MethodGet, resource, model)
}

// Post send HTTP Post method request
func (client *RestClient) Post(resource string, model RestModel) error {
	return client.request(http.MethodPost, resource, model)
}

// Put send HTTP Put method request
func (client *RestClient) Put(resource string, model RestModel) error {
	return client.request(http.MethodPut, resource, model)
}

// Delete send HTTP Delete method request
func (client *RestClient) Delete(resource string, model RestModel) error {
	return client.request(http.MethodDelete, resource, model)
}

func (client *RestClient) BuildSignature(header *http.Header, method, resource string) {
	builder := make([]string, 0, 5)
	builder = append(builder, method)
	builder = append(builder, header.Get(HttpHeaderContentType))
	builder = append(builder, header.Get(HttpHeaderDate))

	headersToSign := make(map[string][]string)
	for k, v := range *header {
		lower := strings.ToLower(k)
		if strings.HasPrefix(lower, DatahubHeadersPrefix) {
			headersToSign[lower] = v
		}
	}

	keys := make([]string, len(headersToSign))
	for k := range headersToSign {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		for _, v := range headersToSign[k] {
			builder = append(builder, fmt.Sprintf("%s:%s", k, v))
		}
	}

	builder = append(builder, resource)

	canonString := strings.Join(builder, "\n")

	log.Debug(fmt.Sprintf("canonString: %s, accesskey: %s", canonString, client.Account.GetAccountKey()))

	hash := hmac.New(sha1.New, []byte(client.Account.GetAccountKey()))
	hash.Write([]byte(canonString))
	crypto := hash.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(crypto)
	authorization := fmt.Sprintf("DATAHUB %s:%s", client.Account.GetAccountId(), signature)

	header.Add(HttpHeaderAuthorization, authorization)
}

func (client *RestClient) request(method, resource string, model RestModel) error {
	url := fmt.Sprintf("%s%s", client.Endpoint, resource)

	reqBody, err := model.RequestBodyEncode(method)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Add(DatahubHeadersPrefix+"client-version", DATAHUB_CLIENT_VERSION)
	req.Header.Add(HttpHeaderDate, time.Now().UTC().Format(http.TimeFormat))
	req.Header.Add(HttpHeaderUserAgent, client.Useragent)
	req.Header.Add(HttpHeaderContentType, "application/json")
	if client.Account.GetSecurityToken() != "" {
		req.Header.Add(HttpHeaderSecurityToken, client.Account.GetSecurityToken())
	}
	if len(reqBody) > 0 {
		req.Header.Add(HttpHeaderContentLength, fmt.Sprintf("%d", len(reqBody)))
	}
	client.BuildSignature(&req.Header, method, resource)

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respResult, err := NewCommonResponseResult(resp.StatusCode, &resp.Header, respBody)
	log.Debug(fmt.Sprintf("request id: %s\nrequest url: %s\nrequest headers: %v\nrequest body: %s\nresponse headers: %v\nresponse body: %s",
		respResult.RequestId, url, req.Header, string(reqBody), resp.Header, string(respBody)))
	if err != nil {
		return err
	}

	return model.ResponseBodyDecode(method, respBody)
}
