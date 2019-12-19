package fastly

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RequestOptions is the list of options to pass to the request.
type RequestOptions struct {
	// Params is a map of key-value pairs that will be added to the Request.
	Params map[string]string

	// Headers is a map of key-value pairs that will be added to the Request.
	Headers map[string]string

	// Body is an io.Reader object that will be streamed or uploaded with the
	// Request. BodyLength is the final size of the Body.
	Body       io.Reader
	BodyLength int64
}

// RawRequest accepts a verb, URL, and RequestOptions struct and returns the
// constructed http.Request and any errors that occurred
func (c *Client) RawRequest(verb, p string, ro *RequestOptions) (*http.Request, error) {
	// Ensure we have request options.
	if ro == nil {
		ro = new(RequestOptions)
	}

	// Append the path to the URL.
	u := strings.TrimRight(c.url.String(), "/") + "/" + strings.TrimLeft(p, "/")

	// Create the request object.
	request, err := http.NewRequest(verb, u, ro.Body)
	if err != nil {
		return nil, err
	}

	var params = make(url.Values)
	for k, v := range ro.Params {
		params.Add(k, v)
	}
	request.URL.RawQuery = params.Encode()

	// Set the API key.
	if len(c.apiKey) > 0 {
		request.Header.Set(APIKeyHeader, c.apiKey)
	}

	// Set the User-Agent.
	request.Header.Set("User-Agent", UserAgent)

	// Add any custom headers.
	for k, v := range ro.Headers {
		request.Header.Add(k, v)
	}

	// Add Content-Length if we have it.
	if ro.BodyLength > 0 {
		request.ContentLength = ro.BodyLength
	}

	return request, nil
}

// SimpleGet combines the RawRequest and Request methods,
// but doesn't add any parameters or change any encoding in the URL
// passed to it. It's mostly for calling the URLs given to us
// directly from Fastly without mangling them.
func (c *Client) SimpleGet(target string) (*http.Response, error) {
	// We parse the URL and then convert it right back to a string
	// later; this just acts as a check that Fastly isn't sending
	// us nonsense.
	url, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(c.apiKey) > 0 {
		request.Header.Set(APIKeyHeader, c.apiKey)
	}
	request.Header.Set("User-Agent", UserAgent)

	resp, err := checkResp(c.HTTPClient.Do(request))
	if err != nil {
		return resp, err
	}
	return resp, nil
}
