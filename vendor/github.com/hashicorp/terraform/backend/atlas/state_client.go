package atlas

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/go-rootcerts"
	"github.com/hashicorp/terraform/state/remote"
	"github.com/hashicorp/terraform/terraform"
)

const (
	// defaultAtlasServer is used when no address is given
	defaultAtlasServer = "https://atlas.hashicorp.com/"
	atlasTokenHeader   = "X-Atlas-Token"
)

// AtlasClient implements the Client interface for an Atlas compatible server.
type stateClient struct {
	Server      string
	ServerURL   *url.URL
	User        string
	Name        string
	AccessToken string
	RunId       string
	HTTPClient  *retryablehttp.Client

	conflictHandlingAttempted bool
}

func (c *stateClient) Get() (*remote.Payload, error) {
	// Make the HTTP request
	req, err := retryablehttp.NewRequest("GET", c.url().String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to make HTTP request: %v", err)
	}

	req.Header.Set(atlasTokenHeader, c.AccessToken)

	// Request the url
	client, err := c.http()
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle the common status codes
	switch resp.StatusCode {
	case http.StatusOK:
		// Handled after
	case http.StatusNoContent:
		return nil, nil
	case http.StatusNotFound:
		return nil, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("HTTP remote state endpoint requires auth")
	case http.StatusForbidden:
		return nil, fmt.Errorf("HTTP remote state endpoint invalid auth")
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("HTTP remote state internal server error")
	default:
		return nil, fmt.Errorf(
			"Unexpected HTTP response code: %d\n\nBody: %s",
			resp.StatusCode, c.readBody(resp.Body))
	}

	// Read in the body
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return nil, fmt.Errorf("Failed to read remote state: %v", err)
	}

	// Create the payload
	payload := &remote.Payload{
		Data: buf.Bytes(),
	}

	if len(payload.Data) == 0 {
		return nil, nil
	}

	// Check for the MD5
	if raw := resp.Header.Get("Content-MD5"); raw != "" {
		md5, err := base64.StdEncoding.DecodeString(raw)
		if err != nil {
			return nil, fmt.Errorf("Failed to decode Content-MD5 '%s': %v", raw, err)
		}

		payload.MD5 = md5
	} else {
		// Generate the MD5
		hash := md5.Sum(payload.Data)
		payload.MD5 = hash[:]
	}

	return payload, nil
}

func (c *stateClient) Put(state []byte) error {
	// Get the target URL
	base := c.url()

	// Generate the MD5
	hash := md5.Sum(state)
	b64 := base64.StdEncoding.EncodeToString(hash[:])

	// Make the HTTP client and request
	req, err := retryablehttp.NewRequest("PUT", base.String(), bytes.NewReader(state))
	if err != nil {
		return fmt.Errorf("Failed to make HTTP request: %v", err)
	}

	// Prepare the request
	req.Header.Set(atlasTokenHeader, c.AccessToken)
	req.Header.Set("Content-MD5", b64)
	req.Header.Set("Content-Type", "application/json")
	req.ContentLength = int64(len(state))

	// Make the request
	client, err := c.http()
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to upload state: %v", err)
	}
	defer resp.Body.Close()

	// Handle the error codes
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusConflict:
		return c.handleConflict(c.readBody(resp.Body), state)
	default:
		return fmt.Errorf(
			"HTTP error: %d\n\nBody: %s",
			resp.StatusCode, c.readBody(resp.Body))
	}
}

func (c *stateClient) Delete() error {
	// Make the HTTP request
	req, err := retryablehttp.NewRequest("DELETE", c.url().String(), nil)
	if err != nil {
		return fmt.Errorf("Failed to make HTTP request: %v", err)
	}
	req.Header.Set(atlasTokenHeader, c.AccessToken)

	// Make the request
	client, err := c.http()
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to delete state: %v", err)
	}
	defer resp.Body.Close()

	// Handle the error codes
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusNoContent:
		return nil
	case http.StatusNotFound:
		return nil
	default:
		return fmt.Errorf(
			"HTTP error: %d\n\nBody: %s",
			resp.StatusCode, c.readBody(resp.Body))
	}
}

func (c *stateClient) readBody(b io.Reader) string {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, b); err != nil {
		return fmt.Sprintf("Error reading body: %s", err)
	}

	result := buf.String()
	if result == "" {
		result = "<empty>"
	}

	return result
}

func (c *stateClient) url() *url.URL {
	values := url.Values{}

	values.Add("atlas_run_id", c.RunId)

	return &url.URL{
		Scheme:   c.ServerURL.Scheme,
		Host:     c.ServerURL.Host,
		Path:     path.Join("api/v1/terraform/state", c.User, c.Name),
		RawQuery: values.Encode(),
	}
}

func (c *stateClient) http() (*retryablehttp.Client, error) {
	if c.HTTPClient != nil {
		return c.HTTPClient, nil
	}
	tlsConfig := &tls.Config{}
	err := rootcerts.ConfigureTLS(tlsConfig, &rootcerts.Config{
		CAFile: os.Getenv("ATLAS_CAFILE"),
		CAPath: os.Getenv("ATLAS_CAPATH"),
	})
	if err != nil {
		return nil, err
	}
	rc := retryablehttp.NewClient()

	rc.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if err != nil {
			// don't bother retrying if the certs don't match
			if err, ok := err.(*url.Error); ok {
				if _, ok := err.Err.(x509.UnknownAuthorityError); ok {
					return false, nil
				}
			}
		}
		return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	}

	t := cleanhttp.DefaultTransport()
	t.TLSClientConfig = tlsConfig
	rc.HTTPClient.Transport = t

	c.HTTPClient = rc
	return rc, nil
}

// Atlas returns an HTTP 409 - Conflict if the pushed state reports the same
// Serial number but the checksum of the raw content differs. This can
// sometimes happen when Terraform changes state representation internally
// between versions in a way that's semantically neutral but affects the JSON
// output and therefore the checksum.
//
// Here we detect and handle this situation by ticking the serial and retrying
// iff for the previous state and the proposed state:
//
//   * the serials match
//   * the parsed states are Equal (semantically equivalent)
//
// In other words, in this situation Terraform can override Atlas's detected
// conflict by asserting that the state it is pushing is indeed correct.
func (c *stateClient) handleConflict(msg string, state []byte) error {
	log.Printf("[DEBUG] Handling Atlas conflict response: %s", msg)

	if c.conflictHandlingAttempted {
		log.Printf("[DEBUG] Already attempted conflict resolution; returning conflict.")
	} else {
		c.conflictHandlingAttempted = true
		log.Printf("[DEBUG] Atlas reported conflict, checking for equivalent states.")

		payload, err := c.Get()
		if err != nil {
			return conflictHandlingError(err)
		}

		currentState, err := terraform.ReadState(bytes.NewReader(payload.Data))
		if err != nil {
			return conflictHandlingError(err)
		}

		proposedState, err := terraform.ReadState(bytes.NewReader(state))
		if err != nil {
			return conflictHandlingError(err)
		}

		if statesAreEquivalent(currentState, proposedState) {
			log.Printf("[DEBUG] States are equivalent, incrementing serial and retrying.")
			proposedState.Serial++
			var buf bytes.Buffer
			if err := terraform.WriteState(proposedState, &buf); err != nil {
				return conflictHandlingError(err)

			}
			return c.Put(buf.Bytes())
		} else {
			log.Printf("[DEBUG] States are not equivalent, returning conflict.")
		}
	}

	return fmt.Errorf(
		"Atlas detected a remote state conflict.\n\nMessage: %s", msg)
}

func conflictHandlingError(err error) error {
	return fmt.Errorf(
		"Error while handling a conflict response from Atlas: %s", err)
}

func statesAreEquivalent(current, proposed *terraform.State) bool {
	return current.Serial == proposed.Serial && current.Equal(proposed)
}
