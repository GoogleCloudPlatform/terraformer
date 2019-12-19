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
	"strings"
	"time"

	"github.com/google/jsonapi"
)

// WAFConfigurationSet represents information about a configuration_set.
type WAFConfigurationSet struct {
	ID string `jsonapi:"primary,configuration_set"`
}

// WAF is the information about a firewall object.
type WAF struct {
	ID                string     `jsonapi:"primary,waf"`
	Version           int        `jsonapi:"attr,version"`
	PrefetchCondition string     `jsonapi:"attr,prefetch_condition"`
	Response          string     `jsonapi:"attr,response"`
	LastPush          *time.Time `jsonapi:"attr,last_push,iso8601"`

	ConfigurationSet *WAFConfigurationSet `jsonapi:"relation,configuration_set"`
}

// wafType is used for reflection because JSONAPI wants to know what it's
// decoding into.
var wafType = reflect.TypeOf(new(WAF))

// ListWAFsInput is used as input to the ListWAFs function.
type ListWAFsInput struct {
	// Service is the ID of the service (required).
	Service string

	// Version is the specific configuration version (required).
	Version int
}

// ListWAFs returns the list of wafs for the configuration version.
func (c *Client) ListWAFs(i *ListWAFsInput) ([]*WAF, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/wafs", i.Service, i.Version)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	data, err := jsonapi.UnmarshalManyPayload(resp.Body, wafType)
	if err != nil {
		return nil, err
	}

	wafs := make([]*WAF, len(data))
	for i := range data {
		typed, ok := data[i].(*WAF)
		if !ok {
			return nil, fmt.Errorf("got back a non-WAF response")
		}
		wafs[i] = typed
	}
	return wafs, nil
}

// CreateWAFInput is used as input to the CreateWAF function.
type CreateWAFInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	ID                string `jsonapi:"primary,waf"`
	PrefetchCondition string `jsonapi:"attr,prefetch_condition,omitempty"`
	Response          string `jsonapi:"attr,response,omitempty"`
}

// CreateWAF creates a new Fastly WAF.
func (c *Client) CreateWAF(i *CreateWAFInput) (*WAF, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	path := fmt.Sprintf("/service/%s/version/%d/wafs", i.Service, i.Version)
	resp, err := c.PostJSONAPI(path, i, nil)
	if err != nil {
		return nil, err
	}

	var waf WAF
	if err := jsonapi.UnmarshalPayload(resp.Body, &waf); err != nil {
		return nil, err
	}
	return &waf, nil
}

// GetWAFInput is used as input to the GetWAF function.
type GetWAFInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// ID is the id of the WAF to get.
	ID string
}

func (c *Client) GetWAF(i *GetWAFInput) (*WAF, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.ID == "" {
		return nil, ErrMissingWAFID
	}

	path := fmt.Sprintf("/service/%s/version/%d/wafs/%s", i.Service, i.Version, i.ID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var waf WAF
	if err := jsonapi.UnmarshalPayload(resp.Body, &waf); err != nil {
		return nil, err
	}
	return &waf, nil
}

// UpdateWAFInput is used as input to the UpdateWAF function.
type UpdateWAFInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	ID                string `jsonapi:"primary,waf"`
	PrefetchCondition string `jsonapi:"attr,prefetch_condition,omitempty"`
	Response          string `jsonapi:"attr,response,omitempty"`
}

// UpdateWAF updates a specific WAF.
func (c *Client) UpdateWAF(i *UpdateWAFInput) (*WAF, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.Version == 0 {
		return nil, ErrMissingVersion
	}

	if i.ID == "" {
		return nil, ErrMissingWAFID
	}

	path := fmt.Sprintf("/service/%s/version/%d/wafs/%s", i.Service, i.Version, i.ID)
	resp, err := c.PatchJSONAPI(path, i, nil)
	if err != nil {
		return nil, err
	}

	var waf WAF
	if err := jsonapi.UnmarshalPayload(resp.Body, &waf); err != nil {
		return nil, err
	}
	return &waf, nil
}

// DeleteWAFInput is used as input to the DeleteWAFInput function.
type DeleteWAFInput struct {
	// Service is the ID of the service. Version is the specific configuration
	// version. Both fields are required.
	Service string
	Version int

	// ID is the id of the WAF to delete.
	ID string
}

func (c *Client) DeleteWAF(i *DeleteWAFInput) error {
	if i.Service == "" {
		return ErrMissingService
	}

	if i.Version == 0 {
		return ErrMissingVersion
	}

	if i.ID == "" {
		return ErrMissingWAFID
	}

	path := fmt.Sprintf("/service/%s/version/%d/wafs/%s", i.Service, i.Version, i.ID)
	_, err := c.Delete(path, nil)
	return err
}

// OWASP is the information about an OWASP object.
type OWASP struct {
	ID                               string     `jsonapi:"primary,owasp"`
	AllowedHTTPVersions              string     `jsonapi:"attr,allowed_http_versions"`
	AllowedMethods                   string     `jsonapi:"attr,allowed_methods"`
	AllowedRequestContentType        string     `jsonapi:"attr,allowed_request_content_type"`
	AllowedRequestContentTypeCharset string     `jsonapi:"attr,allowed_request_content_type_charset"`
	ArgLength                        int        `jsonapi:"attr,arg_length"`
	ArgNameLength                    int        `jsonapi:"attr,arg_name_length"`
	CombinedFileSizes                int        `jsonapi:"attr,combined_file_sizes"`
	CreatedAt                        *time.Time `jsonapi:"attr,created_at,iso8601"`
	CriticalAnomalyScore             int        `jsonapi:"attr,critical_anomaly_score"`
	CRSValidateUTF8Encoding          bool       `jsonapi:"attr,crs_validate_utf8_encoding"`
	ErrorAnomalyScore                int        `jsonapi:"attr,error_anomaly_score"`
	HighRiskCountryCodes             string     `jsonapi:"attr,high_risk_country_codes"`
	HTTPViolationScoreThreshold      int        `jsonapi:"attr,http_violation_score_threshold"`
	InboundAnomalyScoreThreshold     int        `jsonapi:"attr,inbound_anomaly_score_threshold"`
	LFIScoreThreshold                int        `jsonapi:"attr,lfi_score_threshold"`
	MaxFileSize                      int        `jsonapi:"attr,max_file_size"`
	MaxNumArgs                       int        `jsonapi:"attr,max_num_args"`
	NoticeAnomalyScore               int        `jsonapi:"attr,notice_anomaly_score"`
	ParanoiaLevel                    int        `jsonapi:"attr,paranoia_level"`
	PHPInjectionScoreThreshold       int        `jsonapi:"attr,php_injection_score_threshold"`
	RCEScoreThreshold                int        `jsonapi:"attr,rce_score_threshold"`
	RestrictedExtensions             string     `jsonapi:"attr,restricted_extensions"`
	RestrictedHeaders                string     `jsonapi:"attr,restricted_headers"`
	RFIScoreThreshold                int        `jsonapi:"attr,rfi_score_threshold"`
	SessionFixationScoreThreshold    int        `jsonapi:"attr,session_fixation_score_threshold"`
	SQLInjectionScoreThreshold       int        `jsonapi:"attr,sql_injection_score_threshold"`
	TotalArgLength                   int        `jsonapi:"attr,total_arg_length"`
	UpdatedAt                        *time.Time `jsonapi:"attr,updated_at,iso8601"`
	WarningAnomalyScore              int        `jsonapi:"attr,warning_anomaly_score"`
	XSSScoreThreshold                int        `jsonapi:"attr,xss_score_threshold"`
}

// GetOWASPInput is used as input to the GetOWASP function.
type GetOWASPInput struct {
	// Service is the ID of the service. ID is the ID of the firewall.
	// Both fields are required.
	Service string
	ID      string
}

// GetOWASP gets OWASP settings for a service firewall object.
func (c *Client) GetOWASP(i *GetOWASPInput) (*OWASP, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.ID == "" {
		return nil, ErrMissingWAFID
	}

	path := fmt.Sprintf("/service/%s/wafs/%s/owasp", i.Service, i.ID)

	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var owasp OWASP
	if err := jsonapi.UnmarshalPayload(resp.Body, &owasp); err != nil {
		return nil, err
	}
	return &owasp, nil
}

// CreateOWASPInput is used as input to the CreateOWASP function.
type CreateOWASPInput struct {
	// Service is the ID of the service. ID is the ID of the firewall.
	// Both fields are required.
	Service string
	ID      string `jsonapi:"primary,owasp"`
	Type    string `jsonapi:"attr,type"`
}

// CreateOWASP creates an OWASP settings object for a service firewall object.
func (c *Client) CreateOWASP(i *CreateOWASPInput) (*OWASP, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.ID == "" {
		return nil, ErrMissingWAFID
	}

	path := fmt.Sprintf("/service/%s/wafs/%s/owasp", i.Service, i.ID)
	resp, err := c.PostJSONAPI(path, i, nil)
	if err != nil {
		return nil, err
	}

	var owasp OWASP
	if err := jsonapi.UnmarshalPayload(resp.Body, &owasp); err != nil {
		return nil, err
	}
	return &owasp, nil
}

// UpdateOWASPInput is used as input to the CreateOWASP function.
type UpdateOWASPInput struct {
	// Service is the ID of the service. ID is the ID of the firewall.
	// Both fields are required.
	Service string
	ID      string
	OWASPID string `jsonapi:"primary,owasp,omitempty"`

	Type                             string     `jsonapi:"attr,type"`
	AllowedHTTPVersions              string     `jsonapi:"attr,allowed_http_versions,omitempty"`
	AllowedMethods                   string     `jsonapi:"attr,allowed_methods,omitempty"`
	AllowedRequestContentType        string     `jsonapi:"attr,allowed_request_content_type,omitempty"`
	AllowedRequestContentTypeCharset string     `jsonapi:"attr,allowed_request_content_type_charset,omitempty"`
	ArgLength                        int        `jsonapi:"attr,arg_length,omitempty"`
	ArgNameLength                    int        `jsonapi:"attr,arg_name_length,omitempty"`
	CombinedFileSizes                int        `jsonapi:"attr,combined_file_sizes,omitempty"`
	CreatedAt                        *time.Time `jsonapi:"attr,created_at,omitempty,iso8601"`
	CriticalAnomalyScore             int        `jsonapi:"attr,critical_anomaly_score,omitempty"`
	CRSValidateUTF8Encoding          bool       `jsonapi:"attr,crs_validate_utf8_encoding,omitempty"`
	ErrorAnomalyScore                int        `jsonapi:"attr,error_anomaly_score,omitempty"`
	HighRiskCountryCodes             string     `jsonapi:"attr,high_risk_country_codes,omitempty"`
	HTTPViolationScoreThreshold      int        `jsonapi:"attr,http_violation_score_threshold,omitempty"`
	InboundAnomalyScoreThreshold     int        `jsonapi:"attr,inbound_anomaly_score_threshold,omitempty"`
	LFIScoreThreshold                int        `jsonapi:"attr,lfi_score_threshold,omitempty"`
	MaxFileSize                      int        `jsonapi:"attr,max_file_size,omitempty"`
	MaxNumArgs                       int        `jsonapi:"attr,max_num_args,omitempty"`
	NoticeAnomalyScore               int        `jsonapi:"attr,notice_anomaly_score,omitempty"`
	ParanoiaLevel                    int        `jsonapi:"attr,paranoia_level,omitempty"`
	PHPInjectionScoreThreshold       int        `jsonapi:"attr,php_injection_score_threshold,omitempty"`
	RCEScoreThreshold                int        `jsonapi:"attr,rce_score_threshold,omitempty"`
	RestrictedExtensions             string     `jsonapi:"attr,restricted_extensions,omitempty"`
	RestrictedHeaders                string     `jsonapi:"attr,restricted_headers,omitempty"`
	RFIScoreThreshold                int        `jsonapi:"attr,rfi_score_threshold,omitempty"`
	SessionFixationScoreThreshold    int        `jsonapi:"attr,session_fixation_score_threshold,omitempty"`
	SQLInjectionScoreThreshold       int        `jsonapi:"attr,sql_injection_score_threshold,omitempty"`
	TotalArgLength                   int        `jsonapi:"attr,total_arg_length,omitempty"`
	UpdatedAt                        *time.Time `jsonapi:"attr,updated_at,omitempty,iso8601"`
	WarningAnomalyScore              int        `jsonapi:"attr,warning_anomaly_score,omitempty"`
	XSSScoreThreshold                int        `jsonapi:"attr,xss_score_threshold,omitempty"`
}

// CreateOWASP creates an OWASP settings object for a service firewall object.
func (c *Client) UpdateOWASP(i *UpdateOWASPInput) (*OWASP, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.ID == "" {
		return nil, ErrMissingWAFID
	}

	if i.OWASPID == "" {
		return nil, ErrMissingOWASPID
	}

	path := fmt.Sprintf("/service/%s/wafs/%s/owasp", i.Service, i.ID)
	resp, err := c.PatchJSONAPI(path, i, nil)
	if err != nil {
		return nil, err
	}

	var owasp OWASP
	if err := jsonapi.UnmarshalPayload(resp.Body, &owasp); err != nil {
		return nil, err
	}
	return &owasp, nil
}

// Rule is the information about a WAF rule.
type Rule struct {
	ID       string `jsonapi:"primary,rule"`
	RuleID   string `jsonapi:"attr,rule_id,omitempty"`
	Severity int    `jsonapi:"attr,severity,omitempty"`
	Message  string `jsonapi:"attr,message,omitempty"`
}

// rulesType is used for reflection because JSONAPI wants to know what it's
// decoding into.
var rulesType = reflect.TypeOf(new(Rule))

// GetRules returns the list of wafs for the configuration version.
func (c *Client) GetRules() ([]*Rule, error) {
	path := fmt.Sprintf("/wafs/rules")
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	data, err := jsonapi.UnmarshalManyPayload(resp.Body, rulesType)
	if err != nil {
		return nil, err
	}

	rules := make([]*Rule, len(data))
	for i := range data {
		typed, ok := data[i].(*Rule)
		if !ok {
			return nil, fmt.Errorf("got back a non-Rules response")
		}
		rules[i] = typed
	}

	return rules, nil
}

// GetRuleVCLInput is used as input to the GetRuleVCL function.
type GetRuleInput struct {
	// RuleID is the ID of the rule and is required.
	RuleID string
}

// GetRule gets a Rule using the Rule ID.
func (c *Client) GetRule(i *GetRuleInput) (*Rule, error) {
	if i.RuleID == "" {
		return nil, ErrMissingRuleID
	}

	path := fmt.Sprintf("/wafs/rules/%s", i.RuleID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var rule Rule
	if err := jsonapi.UnmarshalPayload(resp.Body, &rule); err != nil {
		return nil, err
	}
	return &rule, nil
}

// RuleVCL is the information about a Rule's VCL.
type RuleVCL struct {
	ID  string `jsonapi:"primary,rule_vcl"`
	VCL string `jsonapi:"attr,vcl,omitempty"`
}

// GetRuleVCL gets the VCL for a Rule.
func (c *Client) GetRuleVCL(i *GetRuleInput) (*RuleVCL, error) {
	if i.RuleID == "" {
		return nil, ErrMissingRuleID
	}

	path := fmt.Sprintf("/wafs/rules/%s/vcl", i.RuleID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var vcl RuleVCL
	if err := jsonapi.UnmarshalPayload(resp.Body, &vcl); err != nil {
		return nil, err
	}
	return &vcl, nil
}

// GetWAFRuleVCLInput is used as input to the GetWAFRuleVCL function.
type GetWAFRuleVCLInput struct {
	// ID is the ID of the firewall. RuleID is the ID of the rule.
	// Both are required.
	ID     string
	RuleID string
}

// GetWAFRuleVCL gets the VCL for a role associated with a firewall WAF.
func (c *Client) GetWAFRuleVCL(i *GetWAFRuleVCLInput) (*RuleVCL, error) {
	if i.ID == "" {
		return nil, ErrMissingWAFID
	}

	if i.RuleID == "" {
		return nil, ErrMissingRuleID
	}

	path := fmt.Sprintf("/wafs/%s/rules/%s/vcl", i.ID, i.RuleID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var vcl RuleVCL
	if err := jsonapi.UnmarshalPayload(resp.Body, &vcl); err != nil {
		return nil, err
	}
	return &vcl, nil
}

// Ruleset is the information about a firewall object's ruleset.
type Ruleset struct {
	ID       string     `jsonapi:"primary,ruleset"`
	VCL      string     `jsonapi:"attr,vcl,omitempty"`
	LastPush *time.Time `jsonapi:"attr,last_push,omitempty,iso8601"`
}

// GetWAFRuleRuleSetsInput is used as input to the GetWAFRuleRuleSets function.
type GetWAFRuleRuleSetsInput struct {
	// Service is the ID of the service. ID is the ID of the firewall.
	// Both fields are required.
	Service string
	ID      string
}

// GetWAFRuleRuleSets gets the VCL for rulesets associated with a firewall WAF.
func (c *Client) GetWAFRuleRuleSets(i *GetWAFRuleRuleSetsInput) (*Ruleset, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.ID == "" {
		return nil, ErrMissingWAFID
	}

	path := fmt.Sprintf("/service/%s/wafs/%s/ruleset", i.Service, i.ID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var ruleset Ruleset
	if err := jsonapi.UnmarshalPayload(resp.Body, &ruleset); err != nil {
		return nil, err
	}
	return &ruleset, nil
}

// UpdateWAFRuleRuleSetsInput is used as input to the UpdateWAFRuleSets function.
type UpdateWAFRuleRuleSetsInput struct {
	// Service is the ID of the service. ID is the ID of the firewall.
	// Both fields are required.
	Service string
	ID      string `jsonapi:"primary,ruleset"`
}

// UpdateWAFRuleSets updates the rulesets for a role associated with a firewall WAF.
func (c *Client) UpdateWAFRuleSets(i *UpdateWAFRuleRuleSetsInput) (*Ruleset, error) {
	if i.Service == "" {
		return nil, ErrMissingService
	}

	if i.ID == "" {
		return nil, ErrMissingWAFID
	}

	path := fmt.Sprintf("/service/%s/wafs/%s/ruleset", i.Service, i.ID)
	resp, err := c.PatchJSONAPI(path, i, nil)
	if err != nil {
		return nil, err
	}

	var ruleset Ruleset
	if err := jsonapi.UnmarshalPayload(resp.Body, &ruleset); err != nil {
		return nil, err
	}
	return &ruleset, nil
}

// GetWAFRuleStatusesInput specifies the parameters for the GetWAFRuleStatuses call
type GetWAFRuleStatusesInput struct {
	Service string
	WAF     string
	Filters GetWAFRuleStatusesFilters
}

// WAFRuleStatus stores the information about a rule received from Fastly
type WAFRuleStatus struct {
	ID     string `jsonapi:"primary,rule_status"` // This is the ID of the status, not the ID of the rule. Currently, it is of the format ${WAF_ID}-${rule_ID}, if you want to infer those based on this field.
	Status string `jsonapi:"attr,status"`

	Tag string `jsonapi:"attr,name,omitempty"` // This will only be set in a response for modifying rules based on tag.

	// HACK: These two fields are supposed to be sent in response
	// to requests for rule status data, but the entire "Relationships"
	// field is currently missing from Fastly responses, so they are
	// instead inferred from the status ID (see inferIDs method).
	// WAF  newTypeThatDoesntExistNow `jsonapi:"relation,waf"`
	// Rule newTypeThatDoesntExistNow `jsonapi:"relation,rule"`
}

// GetWAFRuleStatusesFilters provides a set of parameters for filtering the
// results of the call to get the rules associated with a WAF.
type GetWAFRuleStatusesFilters struct {
	Status     string
	Accuracy   int
	Maturity   int
	Message    string
	Revision   int
	RuleID     string
	TagID      int    // Filter by a single tag ID.
	TagName    string // Filter by single tag name.
	Version    string
	Tags       []int // Return all rules with any of the specified tag IDs.
	MaxResults int   // Max number of returned results per request.
	Page       int   // Which page of results to return.
}

// formatFilters converts user input into query parameters for filtering
// Fastly results for rules in a WAF.
func (i *GetWAFRuleStatusesInput) formatFilters() map[string]string {
	input := i.Filters
	result := map[string]string{}
	pairings := map[string]interface{}{
		"filter[status]":           input.Status,
		"filter[rule][accuracy]":   input.Accuracy,
		"filter[rule][maturity]":   input.Maturity,
		"filter[rule][message]":    input.Message,
		"filter[rule][revision]":   input.Revision,
		"filter[rule][rule_id]":    input.RuleID,
		"filter[rule][tags]":       input.TagID,
		"filter[rule][tags][name]": input.TagName,
		"filter[rule][version]":    input.Version,
		"include":                  input.Tags,
		"page[size]":               input.MaxResults,
		"page[number]":             input.Page, // starts at 1, not 0
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
		case "[]int":
			// convert ints to strings
			toStrings := []string{}
			values := value.([]int)
			for _, i := range values {
				toStrings = append(toStrings, strconv.Itoa(i))
			}
			// concat strings
			if len(values) > 0 {
				result[key] = strings.Join(toStrings, ",")
			}
		}
	}
	return result
}

// GetWAFRuleStatuses fetches the status of a subset of rules associated with a WAF.
func (c *Client) GetWAFRuleStatuses(i *GetWAFRuleStatusesInput) (GetWAFRuleStatusesResponse, error) {
	statusResponse := GetWAFRuleStatusesResponse{Rules: []*WAFRuleStatus{}}
	if i.Service == "" {
		return statusResponse, ErrMissingService
	}
	if i.WAF == "" {
		return statusResponse, ErrMissingWAFID
	}

	path := fmt.Sprintf("/service/%s/wafs/%s/rule_statuses", i.Service, i.WAF)
	filters := &RequestOptions{Params: i.formatFilters()}

	resp, err := c.Get(path, filters)
	if err != nil {
		return statusResponse, err
	}
	err = c.interpretWAFRuleStatusesPage(&statusResponse, resp)
	// NOTE: It's possible for statusResponse to be partially completed before an error
	// was encountered, so the presence of a statusResponse doesn't preclude the presence of
	// an error.
	return statusResponse, err
}

// interpretWAFRuleStatusesPage accepts a Fastly response for a set of WAF rule statuses
// and unmarshals the results. If there are more pages of results, it fetches the next
// page, adds that response to the array of results, and repeats until all results have
// been fetched.
func (c *Client) interpretWAFRuleStatusesPage(answer *GetWAFRuleStatusesResponse, received *http.Response) error {
	// before we pull the status info out of the response body, fetch
	// pagination info from it:
	pages, body, err := getPages(received.Body)
	if err != nil {
		return err
	}
	data, err := jsonapi.UnmarshalManyPayload(body, reflect.TypeOf(new(WAFRuleStatus)))
	if err != nil {
		return err
	}

	for i := range data {
		typed, ok := data[i].(*WAFRuleStatus)
		if !ok {
			return fmt.Errorf("got back response of unexpected type")
		}
		answer.Rules = append(answer.Rules, typed)
	}
	if pages.Next != "" {
		// NOTE: pages.Next URL includes filters already
		resp, err := c.SimpleGet(pages.Next)
		if err != nil {
			return err
		}
		c.interpretWAFRuleStatusesPage(answer, resp)
	}
	return nil
}

// linksResponse is used to pull the "Links" pagination fields from
// a call to Fastly; these are excluded from the results of the jsonapi
// call to `UnmarshalManyPayload()`, so we have to fetch them separately.
type linksResponse struct {
	Links paginationInfo `json:"links"`
}

// paginationInfo stores links to searches related to the current one, showing
// any information about additional results being stored on another page
type paginationInfo struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Next  string `json:"next,omitempty"`
}

// GetWAFRuleStatusesResponse is the data returned to the user from a GetWAFRuleStatus call
type GetWAFRuleStatusesResponse struct {
	Rules []*WAFRuleStatus
}

// getPages parses a response to get the pagination data without destroying
// the reader we receive as "resp.Body"; this essentially copies resp.Body
// and returns it so we can use it again.
func getPages(body io.Reader) (paginationInfo, io.Reader, error) {
	var buf bytes.Buffer
	tee := io.TeeReader(body, &buf)

	bodyBytes, err := ioutil.ReadAll(tee)
	if err != nil {
		return paginationInfo{}, nil, err
	}

	var pages linksResponse
	json.Unmarshal(bodyBytes, &pages)
	return pages.Links, bytes.NewReader(buf.Bytes()), nil
}

// GetWAFRuleStatusInput specifies the parameters for the GetWAFRuleStatus call.
type GetWAFRuleStatusInput struct {
	ID      int
	Service string
	WAF     string
}

// GetWAFRuleStatus fetches the status of a single rule associated with a WAF.
func (c *Client) GetWAFRuleStatus(i *GetWAFRuleStatusInput) (WAFRuleStatus, error) {
	if i.ID == 0 {
		return WAFRuleStatus{}, ErrMissingRuleID
	}
	if i.Service == "" {
		return WAFRuleStatus{}, ErrMissingService
	}
	if i.WAF == "" {
		return WAFRuleStatus{}, ErrMissingWAFID
	}

	path := fmt.Sprintf("/service/%s/wafs/%s/rules/%d/rule_status", i.Service, i.WAF, i.ID)
	resp, err := c.Get(path, nil)
	if err != nil {
		return WAFRuleStatus{}, err
	}

	var status WAFRuleStatus
	err = jsonapi.UnmarshalPayload(resp.Body, &status)
	return status, err
}

// UpdateWAFRuleStatusInput specifies the parameters for the UpdateWAFRuleStatus call.
type UpdateWAFRuleStatusInput struct {
	ID      string `jsonapi:"primary,rule_status"` // The ID of the rule status. Currently in the format ${WAF_ID}-${rule_ID}.
	RuleID  int
	Service string
	WAF     string
	Status  string `jsonapi:"attr,status"`
}

// validate makes sure the UpdateWAFRuleStatusInput instance has all
// fields we need to request a change.
func (i UpdateWAFRuleStatusInput) validate() error {
	if i.ID == "" {
		return ErrMissingID
	}
	if i.RuleID == 0 {
		return ErrMissingRuleID
	}
	if i.Service == "" {
		return ErrMissingService
	}
	if i.WAF == "" {
		return ErrMissingWAFID
	}
	if i.Status == "" {
		return ErrMissingStatus
	}
	return nil
}

// UpdateWAFRuleStatus changes the status of a single rule associated with a WAF.
func (c *Client) UpdateWAFRuleStatus(i *UpdateWAFRuleStatusInput) (WAFRuleStatus, error) {
	if err := i.validate(); err != nil {
		return WAFRuleStatus{}, err
	}

	path := fmt.Sprintf("/service/%s/wafs/%s/rules/%d/rule_status", i.Service, i.WAF, i.RuleID)

	var buf bytes.Buffer
	err := jsonapi.MarshalPayload(&buf, i)
	if err != nil {
		return WAFRuleStatus{}, err
	}

	options := &RequestOptions{
		Body: &buf,
		Headers: map[string]string{
			"Content-Type": jsonapi.MediaType,
			"Accept":       jsonapi.MediaType,
		},
	}

	resp, err := c.Patch(path, options)
	if err != nil {
		return WAFRuleStatus{}, err
	}

	var status WAFRuleStatus
	err = jsonapi.UnmarshalPayload(resp.Body, &status)
	return status, err
}

// UpdateWAFRuleTagStatusInput specifies the parameters for the UpdateWAFRuleStatus call.
type UpdateWAFRuleTagStatusInput struct {
	Service string
	WAF     string
	Status  string `json:"status"` // `jsonapi:"attr,status"`
	Tag     string `json:"name"`   // `jsonapi:"attr,name"`
	Force   bool   `json:"force"`  // `jsonapi:"attr,force"`
	// HACK: This won't work with the jsonapi struct tags, because the POST body expected by
	// Fastly doesn't conform to the jsonapi spec -- there's no ID field at the top level,
	// and there's no way for us to indicate the "type" of the entity without a primary key.
	// ID field is required: http://jsonapi.org/format/#document-resource-objects
}

// updateWAFRuleTagStatusBody is the top-level object sent to Fastly based on
// UpdateWAFRuleTagStatusInput from the user.
type updateWAFRuleTagStatusBody struct {
	Data updateWAFRuleTagStatusData `json:"data"`
}

type updateWAFRuleTagStatusData struct {
	Type       string                       `json:"type"`       // hard-coded because we can't use jsonapi
	Attributes *UpdateWAFRuleTagStatusInput `json:"attributes"` // supplied by user as input
}

// validate makes sure the UpdateWAFRuleStatusInput instance has all
// fields we need to request a change. Almost, but not quite, identical to
// UpdateWAFRuleStatusInput.validate()
func (i UpdateWAFRuleTagStatusInput) validate() error {
	if i.Tag == "" {
		return ErrMissingTag
	}
	if i.Service == "" {
		return ErrMissingService
	}
	if i.WAF == "" {
		return ErrMissingWAFID
	}
	if i.Status == "" {
		return ErrMissingStatus
	}
	return nil
}

// UpdateWAFRuleTagStatus changes the status of a single rule associated with a WAF.
// NOTE: This call currently appears to return *all* rules attached to the WAF, rather
// than just the ones that were modified by the call.
func (c *Client) UpdateWAFRuleTagStatus(input *UpdateWAFRuleTagStatusInput) (GetWAFRuleStatusesResponse, error) {
	if err := input.validate(); err != nil {
		return GetWAFRuleStatusesResponse{}, err
	}

	path := fmt.Sprintf("/service/%s/wafs/%s/rule_statuses", input.Service, input.WAF)

	body := updateWAFRuleTagStatusBody{
		Data: updateWAFRuleTagStatusData{
			Type:       "rule_status",
			Attributes: input,
		},
	}
	encoded, err := json.Marshal(body)
	if err != nil {
		return GetWAFRuleStatusesResponse{}, err
	}

	options := &RequestOptions{
		Body: bytes.NewReader(encoded),
		Headers: map[string]string{
			"Content-Type": jsonapi.MediaType,
			"Accept":       jsonapi.MediaType,
		},
	}

	resp, err := c.Post(path, options)
	if err != nil {
		return GetWAFRuleStatusesResponse{}, err
	}

	statusResponse := GetWAFRuleStatusesResponse{Rules: []*WAFRuleStatus{}}
	err = c.interpretWAFRuleStatusesPage(&statusResponse, resp)

	return statusResponse, err
}

// UpdateWAFConfigSetInput is used as input to the UpdateWAFConfigSet function.
type UpdateWAFConfigSetInput struct {
	WAFList     []ConfigSetWAFs
	ConfigSetID string
}

// ConfigSetWAFs used to store the ID of a WAF needed to update config set relationships
type ConfigSetWAFs struct {
	ID string `jsonapi:"primary,waf"`
}

// UpdateWAFConfigSetResponse stores the list of WAFs returned from the call to update its config set
type UpdateWAFConfigSetResponse struct {
	IDs []ConfigSetWAFs `jsonapi:"primary,waf"`
}

// UpdateWAFConfigSet updates a list of WAFs with the given configset
func (c *Client) UpdateWAFConfigSet(i *UpdateWAFConfigSetInput) (UpdateWAFConfigSetResponse, error) {
	if err := i.validate(); err != nil {
		return UpdateWAFConfigSetResponse{}, err
	}
	var wafs []interface{}
	for _, w := range i.WAFList {
		wafs = append(wafs, &w)
	}

	path := fmt.Sprintf("/wafs/configuration_sets/%s/relationships/wafs", i.ConfigSetID)
	resp, err := c.PatchJSONAPI(path, wafs, nil)
	if err != nil {
		return UpdateWAFConfigSetResponse{}, err
	}

	wafConfigSetResponse := UpdateWAFConfigSetResponse{}

	err = c.interpretWAFCongfigSetResponse(&wafConfigSetResponse, resp)
	if err != nil {
		return UpdateWAFConfigSetResponse{}, err
	}

	return wafConfigSetResponse, nil
}

// validate makes sure the UpdateWAFConfigSetInput instance has all
// fields we need to assign the config set to the WAF(s)
func (i UpdateWAFConfigSetInput) validate() error {
	if i.ConfigSetID == "" {
		return ErrMissingConfigSetID
	}

	if len(i.WAFList) == 0 {
		return ErrMissingWAFList
	}

	return nil
}

// interpretWAFCongfigSetResponse accepts a Fastly response containing a set of WAF ID's that
// where given to associate with the config set and unmarshals the results.
// If there are more pages of results, it fetches the next
// page, adds that response to the array of results, and repeats until all results have
// been fetched.
func (c *Client) interpretWAFCongfigSetResponse(answer *UpdateWAFConfigSetResponse, received *http.Response) error {
	pages, body, err := getPages(received.Body)
	if err != nil {
		return err
	}
	data, err := jsonapi.UnmarshalManyPayload(body, reflect.TypeOf([]ConfigSetWAFs{}))
	if err != nil {
		return err
	}

	for i := range data {
		typed, ok := data[i].(*ConfigSetWAFs)
		if !ok {
			return fmt.Errorf("got back response of unexpected type")
		}
		answer.IDs = append(answer.IDs, *typed)
	}

	if pages.Next != "" {
		resp, err := c.SimpleGet(pages.Next)
		if err != nil {
			return err
		}
		c.interpretWAFCongfigSetResponse(answer, resp)
	}

	return nil
}
