package fastly

import "fmt"

// RealtimeStats is a response from Fastly's real-time analytics endpoint
type RealtimeStatsResponse struct {
	Timestamp      uint64          `mapstructure:"Timestamp"`
	Data           []*RealtimeData `mapstructure:"Data"`
	Error          string          `mapstructure:"Error"`
	AggregateDelay uint32          `mapstructure:"AggregateDelay"`
}

// RealtimeData represents combined stats for all Fastly's POPs and aggregate of them.
// It also includes a timestamp of when the stats were recorded
type RealtimeData struct {
	Datacenter map[string]*Stats `mapstructure:"datacenter"`
	Aggregated *Stats            `mapstructure:"aggregated"`
	Recorded   uint64            `mapstructure:"recorded"`
}

// GetRealtimeStatsInput is an input parameter to GetRealtimeStats function
type GetRealtimeStatsInput struct {
	Service   string
	Timestamp uint64
	Limit     uint32
}

// GetRealtimeStats returns realtime stats for a service based on the GetRealtimeStatsInput
// parameter. The realtime stats work in a rolling fasion where first request will return
// a timestamp which should be passed to consequentive call and so on.
// More details at https://docs.fastly.com/api/analytics
func (c *RTSClient) GetRealtimeStats(i *GetRealtimeStatsInput) (*RealtimeStatsResponse, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	path := fmt.Sprintf("/v1/channel/%s/ts/%d", i.Service, i.Timestamp)

	if i.Limit != 0 {
		path = fmt.Sprintf("%s/limit/%d", path, i.Limit)
	}

	resp, err := c.client.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var s *RealtimeStatsResponse
	if err := decodeJSON(&s, resp.Body); err != nil {
		return nil, err
	}
	return s, nil
}
