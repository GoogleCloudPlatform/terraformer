package cloudflare

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// AccessGroup defines a group for allowing or disallowing access to
// one or more Access applications.
type AccessGroup struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Name      string     `json:"name"`

	// The include group works like an OR logical operator. The user must
	// satisfy one of the rules.
	Include []interface{} `json:"include"`

	// The exclude group works like a NOT logical operator. The user must
	// not satisfy all of the rules in exclude.
	Exclude []interface{} `json:"exclude"`

	// The require group works like a AND logical operator. The user must
	// satisfy all of the rules in require.
	Require []interface{} `json:"require"`
}

// AccessGroupEmail is used for managing access based on the email.
// For example, restrict access to users with the email addresses
// `test@example.com` or `someone@example.com`.
type AccessGroupEmail struct {
	Email struct {
		Email string `json:"email"`
	} `json:"email"`
}

// AccessGroupEmailDomain is used for managing access based on an email
// domain domain such as `example.com` instead of individual addresses.
type AccessGroupEmailDomain struct {
	EmailDomain struct {
		Domain string `json:"domain"`
	} `json:"email_domain"`
}

// AccessGroupIP is used for managing access based in the IP. It
// accepts individual IPs or CIDRs.
type AccessGroupIP struct {
	IP struct {
		IP string `json:"ip"`
	} `json:"ip"`
}

// AccessGroupEveryone is used for managing access to everyone.
type AccessGroupEveryone struct {
	Everyone struct{} `json:"everyone"`
}

// AccessGroupAccessGroup is used for managing access based on an
// access group.
type AccessGroupAccessGroup struct {
	Group struct {
		ID string `json:"id"`
	} `json:"group"`
}

// AccessGroupListResponse represents the response from the list
// access group endpoint.
type AccessGroupListResponse struct {
	Result []AccessGroup `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// AccessGroupDetailResponse is the API response, containing a single
// access group.
type AccessGroupDetailResponse struct {
	Success  bool        `json:"success"`
	Errors   []string    `json:"errors"`
	Messages []string    `json:"messages"`
	Result   AccessGroup `json:"result"`
}

// AccessGroups returns all access groups for an access application.
//
// API reference: https://api.cloudflare.com/#access-groups-list-access-groups
func (api *API) AccessGroups(accountID string, pageOpts PaginationOptions) ([]AccessGroup, ResultInfo, error) {
	v := url.Values{}
	if pageOpts.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(pageOpts.PerPage))
	}
	if pageOpts.Page > 0 {
		v.Set("page", strconv.Itoa(pageOpts.Page))
	}

	uri := fmt.Sprintf(
		"/accounts/%s/access/groups",
		accountID,
	)

	if len(v) > 0 {
		uri = uri + "?" + v.Encode()
	}

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []AccessGroup{}, ResultInfo{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessGroupListResponse AccessGroupListResponse
	err = json.Unmarshal(res, &accessGroupListResponse)
	if err != nil {
		return []AccessGroup{}, ResultInfo{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessGroupListResponse.Result, accessGroupListResponse.ResultInfo, nil
}

// AccessGroup returns a single group based on the group ID.
//
// API reference: https://api.cloudflare.com/#access-groups-access-group-details
func (api *API) AccessGroup(accountID, groupID string) (AccessGroup, error) {
	uri := fmt.Sprintf(
		"/accounts/%s/access/groups/%s",
		accountID,
		groupID,
	)

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return AccessGroup{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessGroupDetailResponse AccessGroupDetailResponse
	err = json.Unmarshal(res, &accessGroupDetailResponse)
	if err != nil {
		return AccessGroup{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessGroupDetailResponse.Result, nil
}

// CreateAccessGroup creates a new access group.
//
// API reference: https://api.cloudflare.com/#access-groups-create-access-group
func (api *API) CreateAccessGroup(accountID string, accessGroup AccessGroup) (AccessGroup, error) {
	uri := fmt.Sprintf(
		"/accounts/%s/access/groups",
		accountID,
	)

	res, err := api.makeRequest("POST", uri, accessGroup)
	if err != nil {
		return AccessGroup{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessGroupDetailResponse AccessGroupDetailResponse
	err = json.Unmarshal(res, &accessGroupDetailResponse)
	if err != nil {
		return AccessGroup{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessGroupDetailResponse.Result, nil
}

// UpdateAccessGroup updates an existing access group.
//
// API reference: https://api.cloudflare.com/#access-groups-update-access-group
func (api *API) UpdateAccessGroup(accountID string, accessGroup AccessGroup) (AccessGroup, error) {
	if accessGroup.ID == "" {
		return AccessGroup{}, errors.Errorf("access group ID cannot be empty")
	}
	uri := fmt.Sprintf(
		"/accounts/%s/access/groups/%s",
		accountID,
		accessGroup.ID,
	)

	res, err := api.makeRequest("PUT", uri, accessGroup)
	if err != nil {
		return AccessGroup{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessGroupDetailResponse AccessGroupDetailResponse
	err = json.Unmarshal(res, &accessGroupDetailResponse)
	if err != nil {
		return AccessGroup{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessGroupDetailResponse.Result, nil
}

// DeleteAccessGroup deletes an access group.
//
// API reference: https://api.cloudflare.com/#access-groups-delete-access-group
func (api *API) DeleteAccessGroup(accountID, groupID string) error {
	uri := fmt.Sprintf(
		"/accounts/%s/access/groups/%s",
		accountID,
		groupID,
	)

	_, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}

	return nil
}
