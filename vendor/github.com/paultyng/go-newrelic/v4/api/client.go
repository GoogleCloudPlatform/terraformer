package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tomnomnom/linkheader"

	resty "github.com/go-resty/resty/v2"
)

// Client represents the client state for the API.
type Client struct {
	RestyClient *resty.Client
}

// InfraClient represents the client state for the Infrastructure API
type InfraClient struct {
	Client
}

// NewInfraClient returns a new InfraClient for the specified apiKey.
func NewInfraClient(config Config) InfraClient {
	if config.BaseURL == "" {
		config.BaseURL = "https://infra-api.newrelic.com/v2"
	}

	return InfraClient{New(config)}
}

// ErrorResponse represents an error response from New Relic.
type ErrorResponse struct {
	Detail *ErrorDetail `json:"error,omitempty"`
}

func (e *ErrorResponse) Error() string {
	if e != nil && e.Detail != nil {
		return e.Detail.Title
	}
	return "Unknown error"
}

// ErrorDetail represents the details of an ErrorResponse from New Relic.
type ErrorDetail struct {
	Title string `json:"title,omitempty"`
}

// Config contains all the configuration data for the API Client.
type Config struct {
	// APIKey is the Admin API Key for your New Relic account.
	// This parameter is required.
	APIKey string

	// BaseURL is the base API URL for the client.
	// `Client` defaults to `https://api.newrelic.com/v2`.
	// Use `https://api.eu.newrelic.com/v2` for EU-based accounts.
	// `InfraClient` defaults to `https://infra-api.newrelic.com/v2`.
	// Use `https://intra-api.eu.newrelic.com/v2` for EU-based accounts.
	BaseURL string

	// ProxyURL sets the Resty client's proxy URL (optional).
	ProxyURL string

	// Debug sets the Resty client's debug mode.
	// Defaults to `false`.
	Debug bool

	// TLSConfig is passed to the Resty client's SetTLSClientConfig method (optional).
	// Used to set a custom root certificate or disable security.
	TLSConfig *tls.Config

	// UserAgent is passed to the Resty client's SetHeaders to allow overriding
	// the default user-agent header (go-newrelic/$version)
	UserAgent string

	// HttpTransport is passed to the Resty client's SetTransport method (optional).
	HTTPTransport http.RoundTripper
}

// New returns a new Client for the specified apiKey.
func New(config Config) Client {
	r := resty.New()

	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = "https://api.newrelic.com/v2"
	}

	proxyURL := config.ProxyURL
	if proxyURL != "" {
		r.SetProxy(proxyURL)
	}

	userAgent := config.UserAgent
	if userAgent == "" {
		userAgent = fmt.Sprintf("go-newrelic/%s (https://github.com/paultyng/go-newrelic)", Version)
	}

	r.SetHeaders(map[string]string{
		"X-Api-Key":  config.APIKey,
		"User-Agent": userAgent,
	})
	r.SetHostURL(baseURL)

	if config.TLSConfig != nil {
		r.SetTLSClientConfig(config.TLSConfig)
	}
	if config.Debug {
		r.SetDebug(true)
	}
	if config.HTTPTransport != nil {
		r.SetTransport(config.HTTPTransport)
	}

	c := Client{
		RestyClient: r,
	}

	return c
}

// Do exectes an API request with the specified parameters.
func (c *Client) Do(method string, path string, body interface{}, response interface{}) (string, error) {
	r := c.RestyClient.R().
		SetError(&ErrorResponse{}).
		SetHeader("Content-Type", "application/json")

	if body != nil {
		r = r.SetBody(body)
	}

	if response != nil {
		r = r.SetResult(response)
	}

	apiResponse, err := r.Execute(method, path)

	if err != nil {
		return "", err
	}

	nextPath := ""
	header := apiResponse.Header().Get("Link")
	if header != "" {
		links := linkheader.Parse(header)

		for _, link := range links.FilterByRel("next") {
			nextPath = link.URL
			break
		}
	}

	apiResponseBody := apiResponse.Body()
	if nextPath == "" && apiResponseBody != nil && len(apiResponseBody) > 0 {
		linksResponse := struct {
			Links struct {
				Next string `json:"next"`
			} `json:"links"`
		}{}

		err = json.Unmarshal(apiResponseBody, &linksResponse)
		if err != nil {
			return "", err
		}

		if linksResponse.Links.Next != "" {
			nextPath = linksResponse.Links.Next
		}
	}

	statusCode := apiResponse.StatusCode()
	statusClass := statusCode / 100 % 10

	if statusClass == 2 {
		return nextPath, nil
	}

	if statusCode == 404 {
		return "", ErrNotFound
	}

	rawError := apiResponse.Error()

	if rawError != nil {
		apiError := rawError.(*ErrorResponse)

		if apiError.Detail != nil {
			return "", apiError
		}
	}

	return "", fmt.Errorf("Unexpected status %v returned from API", apiResponse.StatusCode())
}
