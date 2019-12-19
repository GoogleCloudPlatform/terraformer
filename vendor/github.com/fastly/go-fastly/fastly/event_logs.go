package fastly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/google/jsonapi"
)

// Events represents an event_logs item response from the Fastly API.
type Event struct {
	ID          string                 `jsonapi:"primary,event"`
	CustomerID  string                 `jsonapi:"attr,customer_id"`
	Description string                 `jsonapi:"attr,description"`
	EventType   string                 `jsonapi:"attr,event_type"`
	IP          string                 `jsonapi:"attr,ip"`
	Metadata    map[string]interface{} `jsonapi:"attr,metadata,omitempty"`
	ServiceID   string                 `jsonapi:"attr,service_id"`
	UserID      string                 `jsonapi:"attr,user_id"`
	CreatedAt   *time.Time             `jsonapi:"attr,created_at,iso8601"`
	Admin       bool                   `jsonapi:"attr,admin"`
}

// GetAPIEventsFilter is used as input to the GetAPIEvents function.
type GetAPIEventsFilterInput struct {
	// CustomerID to Limit the returned events to a specific customer.
	CustomerID string

	// ServiceID to Limit the returned events to a specific service.
	ServiceID string

	// EventType to Limit the returned events to a specific event type. See above for event codes.
	EventType string

	// UserID to Limit the returned events to a specific user.
	UserID string

	// Number is the Pagination page number.
	PageNumber int

	// Size is the Number of items to return on each paginated page.
	MaxResults int
}

// eventLinksResponse is used to pull the "Links" pagination fields from
// a call to Fastly; these are excluded from the results of the jsonapi
// call to `UnmarshalManyPayload()`, so we have to fetch them separately.
// type EventLinksResponse struct {
// 	Links EventsPaginationInfo `json:"links"`
// }

// EventsPaginationInfo stores links to searches related to the current one, showing
// any information about additional results being stored on another page
type EventsPaginationInfo struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Next  string `json:"next,omitempty"`
}

// GetAPIEventsResponse is the data returned to the user from a GetAPIEvents call
type GetAPIEventsResponse struct {
	Events []*Event
	Links  EventsPaginationInfo `json:"links"`
}

// GetAPIEvents lists all the events for a particular customer
func (c *Client) GetAPIEvents(i *GetAPIEventsFilterInput) (GetAPIEventsResponse, error) {
	eventsResponse := GetAPIEventsResponse{
		Events: []*Event{},
		Links:  EventsPaginationInfo{},
	}

	path := fmt.Sprintf("/events")

	filters := &RequestOptions{Params: i.formatEventFilters()}

	resp, err := c.Get(path, filters)

	if err != nil {
		return eventsResponse, err
	}

	err = c.interpretAPIEventsPage(&eventsResponse, i.PageNumber, resp)
	// NOTE: It's possible for eventsResponse to be partially completed before an error
	// was encountered, so the presence of a statusResponse doesn't preclude the presence of
	// an error.
	// }
	return eventsResponse, err
}

// GetAPIEventInput is used as input to the GetAPIEvent function.
type GetAPIEventInput struct {
	// EventID is the ID of the event and is required.
	EventID string
}

// GetAPIEvent gets a specific event
func (c *Client) GetAPIEvent(i *GetAPIEventInput) (*Event, error) {
	if i.EventID == "" {
		return nil, ErrMissingEventID
	}

	path := fmt.Sprintf("/events/%s", i.EventID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var event Event
	if err := jsonapi.UnmarshalPayload(resp.Body, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

// interpretAPIEventsPage accepts a Fastly response for a set of WAF rule statuses
// and unmarshals the results. If there are more pages of results, it fetches the next
// page, adds that response to the array of results, and repeats until all results have
// been fetched.
func (c *Client) interpretAPIEventsPage(answer *GetAPIEventsResponse, pageNum int, received *http.Response) error {
	// before we pull the status info out of the response body, fetch
	// pagination info from it:
	pages, body, err := getEventsPages(received.Body)
	if err != nil {
		return err
	}
	answer.Links = pages
	data, err := jsonapi.UnmarshalManyPayload(body, reflect.TypeOf(new(Event)))
	if err != nil {
		return err
	}

	for i := range data {
		typed, ok := data[i].(*Event)
		if !ok {
			return fmt.Errorf("got back response of unexpected type")
		}
		answer.Events = append(answer.Events, typed)
	}
	if pageNum == 0 {
		if pages.Next != "" {
			// NOTE: pages.Next URL includes filters already
			resp, err := c.SimpleGet(pages.Next)
			if err != nil {
				return err
			}
			c.interpretAPIEventsPage(answer, pageNum, resp)
		}
		return nil
	}
	return nil
}

// getEventsPages parses a response to get the pagination data without destroying
// the reader we receive as "resp.Body"; this essentially copies resp.Body
// and returns it so we can use it again.
func getEventsPages(body io.Reader) (EventsPaginationInfo, io.Reader, error) {
	var buf bytes.Buffer
	tee := io.TeeReader(body, &buf)

	bodyBytes, err := ioutil.ReadAll(tee)
	if err != nil {
		return EventsPaginationInfo{}, nil, err
	}

	var pages *GetAPIEventsResponse
	json.Unmarshal(bodyBytes, &pages)
	return pages.Links, bytes.NewReader(buf.Bytes()), nil
}

// formatEventFilters converts user input into query parameters for filtering
// Fastly events.
func (i *GetAPIEventsFilterInput) formatEventFilters() map[string]string {
	result := map[string]string{}
	pairings := map[string]interface{}{
		"filter[customer_id]": i.CustomerID,
		"filter[service_id]":  i.ServiceID,
		"filter[event_type]":  i.EventType,
		"filter[user_id]":     i.UserID,
		"page[size]":          i.MaxResults,
		"page[number]":        i.PageNumber, // starts at 1, not 0
	}
	// NOTE: This setup means we will not be able to send the zero value
	// of any of these filters. It doesn't appear we would need to at present.

	for key, value := range pairings {
		switch t := reflect.TypeOf(value).String(); t {
		case "string":
			if value != "" {
				result[key] = value.(string)
			}
		case "int":
			if value != 0 {
				result[key] = strconv.Itoa(value.(int))
			}
		}

	}
	return result
}
