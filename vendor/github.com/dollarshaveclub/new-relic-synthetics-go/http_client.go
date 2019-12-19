package synthetics

import (
	"log"
	"net/http"
	"time"
)

// RetryableHTTPClient is the interface to an HTTP client that
// supports retries.
type RetryableHTTPClient interface {
	Do(func() (*http.Request, error)) (*http.Response, error)
}

// HTTPClient is the interface to an HTTP client.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type httpClientWithRetries struct {
	client  HTTPClient
	retries uint
}

func newHTTPClientWithRetries(client HTTPClient, retries uint) *httpClientWithRetries {
	return &httpClientWithRetries{client: client, retries: retries}
}

// Do performs a request with retries.
func (h *httpClientWithRetries) Do(reqFunc func() (*http.Request, error)) (*http.Response, error) {
	var response *http.Response
	for i := uint(0); i < h.retries; i++ {
		req, err := reqFunc()
		if err != nil {
			return nil, err
		}

		response, err := h.client.Do(req)
		if err != nil {
			return nil, err
		}

		if response.StatusCode == http.StatusTooManyRequests {
			if i != h.retries-1 {
				log.Printf("rate limited by synthetics, sleeping then retrying")
				time.Sleep((1 << i) * time.Second)
			}
			continue
		}

		return response, nil
	}

	return response, nil
}
