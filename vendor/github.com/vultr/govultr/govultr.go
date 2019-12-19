package govultr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	version     = "0.1.7"
	defaultBase = "https://api.vultr.com"
	userAgent   = "govultr/" + version
	rateLimit   = 600 * time.Millisecond
	retryLimit  = 3
)

// whiteListURI is an array of endpoints that should not have the API Key passed to them
var whiteListURI = [12]string{"/v1/regions/availability",
	"/v1/app/list",
	"/v1/os/list",
	"/v1/plans/list",
	"/v1/plans/list_baremetal",
	"/v1/plans/list_vc2",
	"/v1/plans/list_vc2z",
	"/v1/plans/list_vdc2",
	"/v1/regions/list",
	"/v1/regions/availability_vc2",
	"/v1/regions/availability_vdc2",
	"/v1/regions/availability_baremetal",
}

// APIKey contains a users API Key for interacting with the API
type APIKey struct {
	// API Key
	key string
}

// Client manages interaction with the Vultr V1 API
type Client struct {
	// Http Client used to interact with the Vultr V1 API
	client *retryablehttp.Client

	// BASE URL for APIs
	BaseURL *url.URL

	// User Agent for the client
	UserAgent string

	// API Key
	APIKey APIKey

	// Services used to interact with the API
	Account         AccountService
	API             APIService
	Application     ApplicationService
	Backup          BackupService
	BareMetalServer BareMetalServerService
	BlockStorage    BlockStorageService
	DNSDomain       DNSDomainService
	DNSRecord       DNSRecordService
	FirewallGroup   FirewallGroupService
	FirewallRule    FireWallRuleService
	ISO             ISOService
	Network         NetworkService
	OS              OSService
	Plan            PlanService
	Region          RegionService
	ReservedIP      ReservedIPService
	Server          ServerService
	Snapshot        SnapshotService
	SSHKey          SSHKeyService
	StartupScript   StartupScriptService
	User            UserService

	// Optional function called after every successful request made to the Vultr API
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// NewClient returns a Vultr API Client
func NewClient(httpClient *http.Client, key string) *Client {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBase)

	client := &Client{
		client:    retryablehttp.NewClient(),
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}

	client.client.HTTPClient = httpClient
	client.client.Logger = nil
	client.client.ErrorHandler = client.vultrErrorHandler
	client.SetRetryLimit(retryLimit)
	client.SetRateLimit(rateLimit)

	client.Account = &AccountServiceHandler{client}
	client.API = &APIServiceHandler{client}
	client.Application = &ApplicationServiceHandler{client}
	client.Backup = &BackupServiceHandler{client}
	client.BareMetalServer = &BareMetalServerServiceHandler{client}
	client.BlockStorage = &BlockStorageServiceHandler{client}
	client.DNSDomain = &DNSDomainServiceHandler{client}
	client.DNSRecord = &DNSRecordsServiceHandler{client}
	client.FirewallGroup = &FireWallGroupServiceHandler{client}
	client.FirewallRule = &FireWallRuleServiceHandler{client}
	client.ISO = &ISOServiceHandler{client}
	client.Network = &NetworkServiceHandler{client}
	client.OS = &OSServiceHandler{client}
	client.Plan = &PlanServiceHandler{client}
	client.Region = &RegionServiceHandler{client}
	client.Server = &ServerServiceHandler{client}
	client.ReservedIP = &ReservedIPServiceHandler{client}
	client.Snapshot = &SnapshotServiceHandler{client}
	client.SSHKey = &SSHKeyServiceHandler{client}
	client.StartupScript = &StartupScriptServiceHandler{client}
	client.User = &UserServiceHandler{client}

	apiKey := APIKey{key: key}
	client.APIKey = apiKey

	return client
}

// NewRequest creates an API Request
func (c *Client) NewRequest(ctx context.Context, method, uri string, body url.Values) (*http.Request, error) {

	path, err := url.Parse(uri)
	resolvedURL := c.BaseURL.ResolveReference(path)

	if err != nil {
		return nil, err
	}

	var reqBody io.Reader

	if body != nil {
		reqBody = strings.NewReader(body.Encode())
	} else {
		reqBody = nil
	}

	req, err := http.NewRequest(method, resolvedURL.String(), reqBody)

	if err != nil {
		return nil, err
	}

	req.Header.Add("API-key", c.APIKey.key)
	for _, v := range whiteListURI {
		if v == uri {
			req.Header.Del("API-key")
			break
		}
	}

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Accept", "application/json")

	if req.Method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

// DoWithContext sends an API Request and returns back the response. The API response is checked  to see if it was
// a successful call. A successful call is then checked to see if we need to unmarshal since some resources
// have their own implements of unmarshal.
func (c *Client) DoWithContext(ctx context.Context, r *http.Request, data interface{}) error {

	rreq, err := retryablehttp.FromRequest(r)

	if err != nil {
		return err
	}

	rreq = rreq.WithContext(ctx)

	res, err := c.client.Do(rreq)

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(r, res)
	}

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK {
		if data != nil {
			if string(body) == "[]" {
				data = nil
			} else {
				if err := json.Unmarshal(body, data); err != nil {
					return err
				}
			}
		}
		return nil
	}

	return errors.New(string(body))
}

// SetBaseURL Overrides the default BaseUrl
func (c *Client) SetBaseURL(baseURL string) error {
	updatedURL, err := url.Parse(baseURL)

	if err != nil {
		return err
	}

	c.BaseURL = updatedURL
	return nil
}

// SetRateLimit Overrides the default rateLimit. For performance, exponential
// backoff is used with the minimum wait being 2/3rds the time provided.
func (c *Client) SetRateLimit(time time.Duration) {
	c.client.RetryWaitMin = time / 3 * 2
	c.client.RetryWaitMax = time
}

// SetUserAgent Overrides the default UserAgent
func (c *Client) SetUserAgent(ua string) {
	c.UserAgent = ua
}

// OnRequestCompleted sets the API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

// SetRetryLimit overrides the default RetryLimit
func (c *Client) SetRetryLimit(n int) {
	c.client.RetryMax = n
}

func (c *Client) vultrErrorHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	if resp == nil {
		return nil, fmt.Errorf("gave up after %d attempts, last error unavailable (resp == nil)", c.client.RetryMax+1)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gave up after %d attempts, last error unavailable (error reading response body: %v)", c.client.RetryMax+1, err)
	}
	return nil, fmt.Errorf("gave up after %d attempts, last error: %#v", c.client.RetryMax+1, strings.TrimSpace(string(buf)))
}
