package api

import (
	"fmt"
	"net/url"
	"strconv"
)

// ApplicationsFilters represents a set of filters to be
// used when querying New Relic applications.
type ApplicationsFilters struct {
	Name     *string
	Host     *string
	IDs      []int
	Language *string
}

// QueryApplications queries for New Relic applications
// with filters to narrow down the result set.  This can
// result in less paging required, for instance if you know
// the name of the application you are searching for ahead of time.
func (c *Client) QueryApplications(filters ApplicationsFilters) ([]Application, error) {
	applications := []Application{}

	reqURL, err := url.Parse("/applications.json")
	if err != nil {
		return nil, err
	}

	qs := reqURL.Query()
	if filters.Name != nil {
		qs.Set("filter[name]", *filters.Name)
	}
	if filters.Host != nil {
		qs.Set("filter[host]", *filters.Host)
	}
	for _, id := range filters.IDs {
		qs.Add("filter[ids]", strconv.Itoa(id))
	}
	if filters.Language != nil {
		qs.Set("filter[language]", *filters.Language)
	}
	reqURL.RawQuery = qs.Encode()

	nextPath := reqURL.String()

	for nextPath != "" {
		resp := struct {
			Applications []Application `json:"applications,omitempty"`
		}{}

		nextPath, err = c.Do("GET", nextPath, nil, &resp)
		if err != nil {
			return nil, err
		}

		applications = append(applications, resp.Applications...)
	}

	return applications, nil
}

// ListApplications lists all the applications you have access to.
func (c *Client) ListApplications() ([]Application, error) {
	return c.QueryApplications(ApplicationsFilters{})
}

// DeleteApplication deletes a non-reporting application from your account.
func (c *Client) DeleteApplication(id int) error {
	u := &url.URL{Path: fmt.Sprintf("/applications/%v.json", id)}
	_, err := c.Do("DELETE", u.String(), nil, nil)
	return err
}
