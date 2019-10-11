package v5

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
)

// net/http RoundTripper interface, a.k.a. Transport
// https://godoc.org/net/http#RoundTripper
type RoundTripWithRetryBackoff struct {
	// Configuration fields for backoff.ExponentialBackOff
	InitialIntervalSeconds int64
	RandomizationFactor    float64
	Multiplier             float64
	MaxIntervalSeconds     int64
	// After MaxElapsedTime the ExponentialBackOff stops.
	// It never stops if MaxElapsedTime == 0.
	MaxElapsedTimeSeconds int64
}

func (r RoundTripWithRetryBackoff) RoundTrip(req *http.Request) (*http.Response, error) {
	var lastResponse *http.Response
	var lastError error

	retryableRoundTrip := func() error {
		lastResponse = nil
		lastError = nil

		// Fresh copy of the body for each retry.
		if req.Body != nil {
			originalBody, _ := req.GetBody()
			if originalBody != nil {
				req.Body = originalBody
			}
		}

		lastResponse, lastError = http.DefaultTransport.RoundTrip(req)
		// Detect Heroku API rate limiting
		// https://devcenter.heroku.com/articles/platform-api-reference#client-error-responses
		if lastResponse != nil && lastResponse.StatusCode == 429 {
			return fmt.Errorf("Heroku API rate limited: 429 Too Many Requests")
		}
		return nil
	}

	rateLimitRetryConfig := &backoff.ExponentialBackOff{
		Clock:               backoff.SystemClock,
		InitialInterval:     time.Duration(int64WithDefault(r.InitialIntervalSeconds, int64(30))) * time.Second,
		RandomizationFactor: float64WithDefault(r.RandomizationFactor, float64(0.25)),
		Multiplier:          float64WithDefault(r.Multiplier, float64(2)),
		MaxInterval:         time.Duration(int64WithDefault(r.MaxIntervalSeconds, int64(900))) * time.Second,
		MaxElapsedTime:      time.Duration(int64WithDefault(r.MaxElapsedTimeSeconds, int64(0))) * time.Second,
	}
	rateLimitRetryConfig.Reset()

	err := backoff.RetryNotify(retryableRoundTrip, rateLimitRetryConfig, notifyLog)
	// Propagate the rate limit error when retries eventually fail.
	if err != nil {
		if lastResponse != nil {
			lastResponse.Body.Close()
		}
		return nil, err
	}
	// Propagate all other response errors.
	if lastError != nil {
		if lastResponse != nil {
			lastResponse.Body.Close()
		}
		return nil, lastError
	}

	return lastResponse, nil
}

func int64WithDefault(v int64, defaultV int64) int64 {
	if v == int64(0) {
		return defaultV
	} else {
		return v
	}
}

func float64WithDefault(v float64, defaultV float64) float64 {
	if v == float64(0) {
		return defaultV
	} else {
		return v
	}
}

func notifyLog(err error, waitDuration time.Duration) {
	log.Printf("Will retry Heroku API request in %s, because %s", waitDuration, err)
}
