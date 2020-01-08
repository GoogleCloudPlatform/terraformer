package fastly

import "fmt"

// Stats represent metrics of a Fastly service
type Stats struct {
	Requests                  uint64      `mapstructure:"requests"`                 // Number of requests processed.
	Hits                      uint64      `mapstructure:"hits"`                     // Number of cache hits.
	HitsTime                  float64     `mapstructure:"hits_time"`                // Total amount of time spent processing cache hits (in seconds).
	Miss                      uint64      `mapstructure:"miss"`                     // Number of cache misses.
	MissTime                  float64     `mapstructure:"miss_time"`                // Amount of time spent processing cache misses (in seconds).
	Pass                      uint64      `mapstructure:"pass"`                     // Number of requests that passed through the CDN without being cached.
	PassTime                  float64     `mapstructure:"pass_time"`                // Amount of time spent processing cache passes (in seconds).
	Synth                     uint64      `mapstructure:"synth"`                    // Number of requests that returned synth response.
	Errors                    uint64      `mapstructure:"errors"`                   // Number of cache errors.
	Restarts                  uint64      `mapstructure:"restarts"`                 // Number of restarts performed.
	HitRatio                  float64     `mapstructure:"hit_ratio"`                // Ratio of cache hits to cache misses (between 0 and 1).
	Bandwidth                 uint64      `mapstructure:"bandwidth"`                // Total bytes delivered (body_size + header_size).
	RequestBodyBytes          uint64      `mapstructure:"req_body_bytes"`           // Total body bytes received.
	RequestHeaderBytes        uint64      `mapstructure:"req_header_bytes"`         // Total header bytes received.
	ResponseBodyBytes         uint64      `mapstructure:"resp_body_bytes"`          // Total body bytes delivered.
	ResponseHeaderBytes       uint64      `mapstructure:"resp_header_bytes"`        // Total header bytes delivered.
	BERequestBodyBytes        uint64      `mapstructure:"bereq_body_bytes"`         // Total body bytes sent to origin.
	BERequestHeaderbytes      uint64      `mapstructure:"bereq_header_bytes"`       // Total header bytes sent to origin.
	Uncachable                uint64      `mapstructure:"uncachable"`               // Number of requests that were designated uncachable.
	Pipe                      uint64      `mapstructure:"pipe"`                     // Optional. Pipe operations performed (legacy feature).
	TLS                       uint64      `mapstructure:"tls"`                      // Number of requests that were received over TLS.
	TLSv10                    uint64      `mapstructure:"tls_v10"`                  // Number of requests received over TLS 1.0.
	TLSv11                    uint64      `mapstructure:"tls_v11"`                  // Number of requests received over TLS 1.`.
	TLSv12                    uint64      `mapstructure:"tls_v12"`                  // Number of requests received over TLS 1.2.
	TLSv13                    uint64      `mapstructure:"tls_v13"`                  // Number of requests received over TLS 1.3.
	Shield                    uint64      `mapstructure:"shield"`                   // Number of requests from shield to origin.
	ShieldResponseBodyBytes   uint64      `mapstructure:"shield_resp_body_bytes"`   // Total body bytes delivered via a shield.
	ShieldResponseHeaderBytes uint64      `mapstructure:"shield_resp_header_bytes"` // Total header bytes delivered via a shield.
	IPv6                      uint64      `mapstructure:"ipv6"`                     // Number of requests that were received over IPv6.
	OTFP                      uint64      `mapstructure:"otfp"`                     // Number of responses that came from the Fastly On-the-Fly Packager for On Demand Streaming service for video-on-demand.
	Video                     uint64      `mapstructure:"video"`                    // Number of responses with the video segment or video manifest MIME type (i.e., application/x-mpegurl, application/vnd.apple.mpegurl, application/f4m, application/dash+xml, application/vnd.ms-sstr+xml, ideo/mp2t, audio/aac, video/f4f, video/x-flv, video/mp4, audio/mp4).
	PCI                       uint64      `mapstructure:"pci"`                      // Number of responses with the PCI flag turned on.
	Log                       uint64      `mapstructure:"log"`                      // Number of log lines sent.
	HTTP2                     uint64      `mapstructure:"http2"`                    // Number of requests received over HTTP2.
	WAFLogged                 uint64      `mapstructure:"waf_logged"`               // Number of requests that triggered a WAF rule and were logged.
	WAFBlocked                uint64      `mapstructure:"waf_blocked"`              // Number of requests that triggered a WAF rule and were blocked.
	WAFPassed                 uint64      `mapstructure:"waf_passed"`               // Number of requests that triggered a WAF rule and were passed.
	AttackRequestBodyBytes    uint64      `mapstructure:"attack_req_body_bytes"`    // Total body bytes received from requests that triggered a WAF rule.
	AttachRequestHeaderBytes  uint64      `mapstructure:"attack_req_header_bytes"`  // Total header bytes received from requests that triggered a WAF rule.
	AttackResponseSynthBytes  uint64      `mapstructure:"attack_resp_synth_bytes"`  // Total bytes delivered for requests that triggered a WAF rule and returned a synthetic response.
	ImageOptimizer            uint64      `mapstructure:"imgopto"`                  // Number of responses that came from the Fastly Image Optimizer service.
	Status200                 uint64      `mapstructure:"status_200"`               // Number of responses sent with status code 200 (Success).
	Status204                 uint64      `mapstructure:"status_204"`               // Number of responses sent with status code 204 (No Content).
	Status301                 uint64      `mapstructure:"status_301"`               // Number of responses sent with status code 301 (Moved Permanently).
	Status302                 uint64      `mapstructure:"status_302"`               // Number of responses sent with status code 302 (Found).
	Status304                 uint64      `mapstructure:"status_304"`               // Number of responses sent with status code 304 (Not Modified).
	Status400                 uint64      `mapstructure:"status_400"`               // Number of responses sent with status code 400 (Bad Request).
	Status401                 uint64      `mapstructure:"status_401"`               // Number of responses sent with status code 401 (Unauthorized).
	Status403                 uint64      `mapstructure:"status_403"`               // Number of responses sent with status code 403 (Forbidden).
	Status404                 uint64      `mapstructure:"status_404"`               // Number of responses sent with status code 404 (Not Found).
	Status416                 uint64      `mapstructure:"status_416"`               // Number of responses sent with status code 416 (Range Not Satisfiable).
	Status500                 uint64      `mapstructure:"status_500"`               // Number of responses sent with status code 500 (Internal Server Error).
	Status501                 uint64      `mapstructure:"status_501"`               // Number of responses sent with status code 501 (Not Implemented).
	Status502                 uint64      `mapstructure:"status_502"`               // Number of responses sent with status code 502 (Bad Gateway).
	Status503                 uint64      `mapstructure:"status_503"`               // Number of responses sent with status code 503 (Service Unavailable).
	Status504                 uint64      `mapstructure:"status_504"`               // Number of responses sent with status code 504 (Gateway Timeout).
	Status505                 uint64      `mapstructure:"status_505"`               // Number of responses sent with status code 505 (HTTP Version Not Supported).
	Status1xx                 uint64      `mapstructure:"status_1xx"`               // Number of "Informational" category status codes delivered.
	Status2xx                 uint64      `mapstructure:"status_2xx"`               // Number of "Success" status codes delivered.
	Status3xx                 uint64      `mapstructure:"status_3xx"`               // Number of "Redirection" codes delivered.
	Status4xx                 uint64      `mapstructure:"status_4xx"`               // Number of "Client Error" codes delivered.
	Status5xx                 uint64      `mapstructure:"status_5xx"`               // Number of "Server Error" codes delivered.
	ObjectSize1k              uint64      `mapstructure:"object_size_1k"`           // Number of objects served that were under 1KB in size.
	ObjectSize10k             uint64      `mapstructure:"object_size_10k"`          // Number of objects served that were between 1KB and 10KB in size.
	ObjectSize100k            uint64      `mapstructure:"object_size_100k"`         // Number of objects served that were between 10KB and 100KB in size.
	ObjectSize1m              uint64      `mapstructure:"object_size_1m"`           // Number of objects served that were between 100KB and 1MB in size.
	ObjectSize10m             uint64      `mapstructure:"object_size_10m"`          // Number of objects served that were between 1MB and 10MB in size.
	ObjectSize100m            uint64      `mapstructure:"object_size_100m"`         // Number of objects served that were between 10MB and 100MB in size.
	ObjectSize1g              uint64      `mapstructure:"object_size_1g"`           // Number of objects served that were between 100MB and 1GB in size.
	MissHistogram             map[int]int `mapstructure:"miss_histogram"`           // Number of requests to origin in time buckets of 10s of milliseconds
	BilledHeaderBytes         uint64      `mapstructure:"billed_header_bytes"`
	BilledBodyBytes           uint64      `mapstructure:"billed_body_bytes"`
}

// GetStatsInput is an input to the GetStats function.
// Stats can be filtered by a Service ID, an individual stats field,
// time range (From and To), sampling rate (By) and/or Fastly region (Region)
// Allowed values for the fields are described at https://docs.fastly.com/api/stats
type GetStatsInput struct {
	Service string
	Field   string
	From    string
	To      string
	By      string
	Region  string
}

// StatsResponse is a response from the service stats API endpoint
type StatsResponse struct {
	Status  string            `mapstructure:"status"`
	Meta    map[string]string `mapstructure:"meta"`
	Message string            `mapstructure:"msg"`
	Data    []*Stats          `mapstructure:"data"`
}

// GetStats returns stats data based on GetStatsInput
func (c *Client) GetStats(i *GetStatsInput) (*StatsResponse, error) {

	p := "/stats"

	if i.Service != "" {
		p = fmt.Sprintf("%s/service/%s", p, i.Service)
	}

	if i.Field != "" {
		p = fmt.Sprintf("%s/field/%s", p, i.Field)
	}

	r, err := c.Get(p, &RequestOptions{
		Params: map[string]string{
			"from":   i.From,
			"to":     i.To,
			"by":     i.By,
			"region": i.Region,
		},
	})
	if err != nil {
		return nil, err
	}

	var sr *StatsResponse
	if err := decodeJSON(&sr, r.Body); err != nil {
		return nil, err
	}

	return sr, nil
}

// UsageStatsResponse is a response from the account usage API endpoint
type UsageStatsResponse struct {
	Status  string            `mapstructure:"status"`
	Meta    map[string]string `mapstructure:"meta"`
	Message string            `mapstructure:"msg"`
	Data    map[string]*Usage `mapstructure:"data"`
}

// Usage represents usage data of a single service or region
type Usage struct {
	Requests  uint64 `mapstructure:"requests"`
	Bandwidth uint64 `mapstructure:"bandwidth"`
}

// RegionsUsage is a list of aggregated usage data by Fastly's region
type RegionsUsage map[string]*Usage

// UsageStatsResponse is a response from the account usage API endpoint
type UsageResponse struct {
	Status  string            `mapstructure:"status"`
	Meta    map[string]string `mapstructure:"meta"`
	Message string            `mapstructure:"msg"`
	Data    *RegionsUsage     `mapstructure:"data"`
}

// GetUsageInput is used as an input to the GetUsage function
// Value for the input are described at https://docs.fastly.com/api/stats
type GetUsageInput struct {
	From   string
	To     string
	By     string
	Region string
}

// GetUsage returns usage information aggregated across all Fastly services and grouped by region.
func (c *Client) GetUsage(i *GetUsageInput) (*UsageResponse, error) {
	r, err := c.Get("/stats/usage", &RequestOptions{
		Params: map[string]string{
			"from":   i.From,
			"to":     i.To,
			"by":     i.By,
			"region": i.Region,
		},
	})
	if err != nil {
		return nil, err
	}
	var sr *UsageResponse
	if err := decodeJSON(&sr, r.Body); err != nil {
		return nil, err
	}

	return sr, nil
}

// UsageStatsResponse is a response from the account usage API endpoint
type UsageByServiceResponse struct {
	Status  string                  `mapstructure:"status"`
	Meta    map[string]string       `mapstructure:"meta"`
	Message string                  `mapstructure:"msg"`
	Data    *ServicesByRegionsUsage `mapstructure:"data"`
}

// ServicesUsage is a list of usage data by a service
type ServicesUsage map[string]*Usage

// ServicesByRegionsUsage is a list of ServicesUsage by Fastly's region
type ServicesByRegionsUsage map[string]*ServicesUsage

// GetUsageByService returns usage information aggregated by service and
// grouped by service and region.
func (c *Client) GetUsageByService(i *GetUsageInput) (*UsageByServiceResponse, error) {
	r, err := c.Get("/stats/usage_by_service", &RequestOptions{
		Params: map[string]string{
			"from":   i.From,
			"to":     i.To,
			"by":     i.By,
			"region": i.Region,
		},
	})
	if err != nil {
		return nil, err
	}
	var sr *UsageByServiceResponse
	if err := decodeJSON(&sr, r.Body); err != nil {
		return nil, err
	}

	return sr, nil
}

// RegionsResponse is a response from Fastly regions API endpoint
type RegionsResponse struct {
	Status  string            `mapstructure:"status"`
	Meta    map[string]string `mapstructure:"meta"`
	Message string            `mapstructure:"msg"`
	Data    []string          `mapstructure:"data"`
}

// GetRegions returns a list of Fastly regions
func (c *Client) GetRegions() (*RegionsResponse, error) {
	r, err := c.Get("stats/regions", nil)
	if err != nil {
		return nil, err
	}

	var rr *RegionsResponse
	if err := decodeJSON(&rr, r.Body); err != nil {
		return nil, err
	}

	return rr, nil
}
