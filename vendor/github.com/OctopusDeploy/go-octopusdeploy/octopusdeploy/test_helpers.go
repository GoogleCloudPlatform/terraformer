package octopusdeploy

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

type roundTripFunc func(r *http.Request) (*http.Response, error)

func (s roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return s(r)
}

// func getFakeOctopusClient(httpClient http.Client) *Client {
// 	return NewClient(&httpClient, "http://octopusserver", "FakeAPIKey")
// }

func getFakeOctopusClient(t *testing.T, apiPath string, statusCode int, responseBody string) *Client {
	httpClient := http.Client{}
	httpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, apiPath, r.URL.Path)
		return &http.Response{
			StatusCode: statusCode,
			Body:       ioutil.NopCloser(strings.NewReader(responseBody)),
		}, nil
	})
	return NewClient(&httpClient, "http://octopusserver", "FakeAPIKey")
}
