package fc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	TRIGGER_TYPE_OSS        = "oss"
	TRIGGER_TYPE_LOG        = "log"
	TRIGGER_TYPE_TIMER      = "timer"
	TRIGGER_TYPE_HTTP       = "http"
	TRIGGER_TYPE_TABLESTORE = "tablestore"
	TRIGGER_TYPE_CDN_EVENTS = "cdn_events"
	TRIGGER_TYPE_MNS_TOPIC  = "mns_topic"
)

// CreateTriggerInput defines trigger creation input
type CreateTriggerInput struct {
	ServiceName  *string
	FunctionName *string
	TriggerCreateObject
}

type TriggerCreateObject struct {
	TriggerName    *string     `json:"triggerName"`
	Description    *string     `json:"description"`
	SourceARN      *string     `json:"sourceArn"`
	TriggerType    *string     `json:"triggerType"`
	InvocationRole *string     `json:"invocationRole"`
	TriggerConfig  interface{} `json:"triggerConfig"`
	Qualifier      *string     `json:"qualifier"`

	err error `json:"-"`
}

func NewCreateTriggerInput(serviceName string, functionName string) *CreateTriggerInput {
	return &CreateTriggerInput{
		ServiceName:  &serviceName,
		FunctionName: &functionName,
	}
}

func (i *CreateTriggerInput) WithQualifier(qualifier string) *CreateTriggerInput {
	i.Qualifier = &qualifier
	return i
}

func (i *CreateTriggerInput) WithTriggerName(name string) *CreateTriggerInput {
	i.TriggerName = &name
	return i
}

func (i *CreateTriggerInput) WithDescription(desc string) *CreateTriggerInput {
	i.Description = &desc
	return i
}

func (i *CreateTriggerInput) WithSourceARN(arn string) *CreateTriggerInput {
	i.SourceARN = &arn
	return i
}

func (i *CreateTriggerInput) WithTriggerType(triggerType string) *CreateTriggerInput {
	i.TriggerType = &triggerType
	return i
}

func (i *CreateTriggerInput) WithInvocationRole(role string) *CreateTriggerInput {
	i.InvocationRole = &role
	return i
}

func (i *CreateTriggerInput) WithTriggerConfig(config interface{}) *CreateTriggerInput {
	i.TriggerConfig = &config
	return i
}

func (i *CreateTriggerInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *CreateTriggerInput) GetPath() string {
	return fmt.Sprintf(triggersPath, pathEscape(*i.ServiceName), pathEscape(*i.FunctionName))
}

func (i *CreateTriggerInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *CreateTriggerInput) GetPayload() interface{} {
	return i.TriggerCreateObject
}

func (i *CreateTriggerInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.FunctionName) {
		return fmt.Errorf("Function name is required but not provided")
	}
	if i.err != nil {
		return i.err
	}
	return nil
}

type CreateTriggerOutput struct {
	Header http.Header
	triggerMetadata
}

func (o CreateTriggerOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o CreateTriggerOutput) GetEtag() string {
	return GetEtag(o.Header)
}

func (o CreateTriggerOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

// MarshalJSON marshals trigger metadata and excludes RawTriggerConfig
func (o CreateTriggerOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(triggerMetadataDisplay{
		Header:           o.Header,
		TriggerName:      o.TriggerName,
		Description:      o.Description,
		SourceARN:        o.SourceARN,
		TriggerType:      o.TriggerType,
		InvocationRole:   o.InvocationRole,
		Qualifier:        o.Qualifier,
		TriggerConfig:    o.TriggerConfig,
		CreatedTime:      o.CreatedTime,
		LastModifiedTime: o.LastModifiedTime,
	})
}

type triggerMetadata struct {
	TriggerName      *string         `json:"triggerName"`
	Description      *string         `json:"description"`
	TriggerID        *string         `json:"triggerID"`
	SourceARN        *string         `json:"sourceArn"`
	TriggerType      *string         `json:"triggerType"`
	InvocationRole   *string         `json:"invocationRole"`
	Qualifier        *string         `json:"qualifier"`
	RawTriggerConfig json.RawMessage `json:"triggerConfig"`
	CreatedTime      *string         `json:"createdTime"`
	LastModifiedTime *string         `json:"lastModifiedTime"`

	TriggerConfig interface{} `json:"-"`
}

type triggerMetadataAlias triggerMetadata

// UnmarshalJSON unmarshals the data to trigger metadata and sets TriggerConfig field to an actual trigger config.
// User can use type switches/assertion to get the actual trigger config.
func (m *triggerMetadata) UnmarshalJSON(data []byte) error {
	// use triggerMetadataAlias instead of triggerMetadata to avoid recursive calls because Unmarshal calls UnmarshalJSON.
	tmp := triggerMetadataAlias{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	switch *tmp.TriggerType {
	case TRIGGER_TYPE_OSS:
		ossTriggerConfig := &OSSTriggerConfig{}
		if err := json.Unmarshal(tmp.RawTriggerConfig, ossTriggerConfig); err != nil {
			return err
		}
		tmp.TriggerConfig = ossTriggerConfig
	case TRIGGER_TYPE_LOG:
		logTriggerConfig := &LogTriggerConfig{}
		if err := json.Unmarshal(tmp.RawTriggerConfig, logTriggerConfig); err != nil {
			return err
		}
		tmp.TriggerConfig = logTriggerConfig
	case TRIGGER_TYPE_TIMER:
		timeTriggerConfig := &TimeTriggerConfig{}
		if err := json.Unmarshal(tmp.RawTriggerConfig, timeTriggerConfig); err != nil {
			return err
		}
		tmp.TriggerConfig = timeTriggerConfig
	case TRIGGER_TYPE_HTTP:
		httpTriggerConfig := &HTTPTriggerConfig{}
		if err := json.Unmarshal(tmp.RawTriggerConfig, httpTriggerConfig); err != nil {
			return err
		}
		tmp.TriggerConfig = httpTriggerConfig
	case TRIGGER_TYPE_TABLESTORE:
		tableStoreTriggerConfig := &TableStoreTriggerConfig{}
		if err := json.Unmarshal(tmp.RawTriggerConfig, tableStoreTriggerConfig); err != nil {
			return err
		}
		tmp.TriggerConfig = tableStoreTriggerConfig
	case TRIGGER_TYPE_CDN_EVENTS:
		cdnEventsTriggerConfig := &CDNEventsTriggerConfig{}
		if err := json.Unmarshal(tmp.RawTriggerConfig, cdnEventsTriggerConfig); err != nil {
			return err
		}
		tmp.TriggerConfig = cdnEventsTriggerConfig
	case TRIGGER_TYPE_MNS_TOPIC:
		mnsTriggerConfig := &MnsTopicTriggerConfig{}
		if err := json.Unmarshal(tmp.RawTriggerConfig, mnsTriggerConfig); err != nil {
			return err
		}
		tmp.TriggerConfig = mnsTriggerConfig
	default:
		return ErrUnknownTriggerType
	}
	*m = triggerMetadata(tmp)
	return nil
}

type triggerMetadataDisplay struct {
	Header           http.Header
	TriggerName      *string     `json:"triggerName"`
	Description      *string     `json:"description"`
	SourceARN        *string     `json:"sourceArn"`
	TriggerType      *string     `json:"triggerType"`
	InvocationRole   *string     `json:"invocationRole"`
	Qualifier        *string     `json:"qualifier"`
	TriggerConfig    interface{} `json:"triggerConfig"`
	CreatedTime      *string     `json:"createdTime"`
	LastModifiedTime *string     `json:"lastModifiedTime"`
}

type OSSTriggerConfig struct {
	Events []string          `json:"events"`
	Filter *OSSTriggerFilter `json:"filter"`
}

func NewOSSTriggerConfig() *OSSTriggerConfig {
	return &OSSTriggerConfig{}
}

func (c *OSSTriggerConfig) WithEvents(events []string) *OSSTriggerConfig {
	c.Events = events
	return c
}

func (c *OSSTriggerConfig) WithFilter(filter *OSSTriggerFilter) *OSSTriggerConfig {
	c.Filter = filter
	return c
}

func (c *OSSTriggerConfig) WithFilterKeyPrefix(prefix string) *OSSTriggerConfig {
	if c.Filter == nil {
		c.Filter = NewOSSTriggerFilter()
	}
	if c.Filter.Key == nil {
		c.Filter.Key = NewOSSTriggerKey()
	}
	c.Filter.Key.Prefix = &prefix
	return c
}

func (c *OSSTriggerConfig) WithFilterKeySuffix(suffix string) *OSSTriggerConfig {
	if c.Filter == nil {
		c.Filter = NewOSSTriggerFilter()
	}
	if c.Filter.Key == nil {
		c.Filter.Key = NewOSSTriggerKey()
	}
	c.Filter.Key.Suffix = &suffix
	return c
}

type OSSTriggerFilter struct {
	Key *OSSTriggerKey `json:"key"`
}

func NewOSSTriggerFilter() *OSSTriggerFilter {
	return &OSSTriggerFilter{}
}

func (f *OSSTriggerFilter) WithKey(key *OSSTriggerKey) *OSSTriggerFilter {
	f.Key = key
	return f
}

type OSSTriggerKey struct {
	Prefix *string `json:"prefix"`
	Suffix *string `json:"suffix"`
}

func NewOSSTriggerKey() *OSSTriggerKey {
	return &OSSTriggerKey{}
}

func (k *OSSTriggerKey) WithPrefix(prefix string) *OSSTriggerKey {
	k.Prefix = &prefix
	return k
}

func (k *OSSTriggerKey) WithSuffix(suffix string) *OSSTriggerKey {
	k.Suffix = &suffix
	return k
}

type GetTriggerInput struct {
	ServiceName  *string
	FunctionName *string
	TriggerName  *string
}

func NewGetTriggerInput(serviceName string, functionName string, triggerName string) *GetTriggerInput {
	return &GetTriggerInput{
		ServiceName:  &serviceName,
		FunctionName: &functionName,
		TriggerName:  &triggerName,
	}
}

func (i *GetTriggerInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *GetTriggerInput) GetPath() string {
	return fmt.Sprintf(singleTriggerPath, pathEscape(*i.ServiceName), pathEscape(*i.FunctionName), pathEscape(*i.TriggerName))
}

func (i *GetTriggerInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *GetTriggerInput) GetPayload() interface{} {
	return nil
}

func (i *GetTriggerInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.FunctionName) {
		return fmt.Errorf("Function name is required but not provided")
	}
	if IsBlank(i.TriggerName) {
		return fmt.Errorf("Trigger name is required but not provided")
	}
	return nil
}

// GetTriggerOutput define trigger response from fc
type GetTriggerOutput struct {
	Header http.Header
	triggerMetadata
}

func (o GetTriggerOutput) GetEtag() string {
	return GetEtag(o.Header)
}

func (o GetTriggerOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o GetTriggerOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

// MarshalJSON marshals trigger metadata and excludes RawTriggerConfig
func (o GetTriggerOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(triggerMetadataDisplay{
		Header:           o.Header,
		TriggerName:      o.TriggerName,
		Description:      o.Description,
		SourceARN:        o.SourceARN,
		TriggerType:      o.TriggerType,
		InvocationRole:   o.InvocationRole,
		Qualifier:        o.Qualifier,
		TriggerConfig:    o.TriggerConfig,
		CreatedTime:      o.CreatedTime,
		LastModifiedTime: o.LastModifiedTime,
	})
}

// TriggerUpdateObject defines update fields in Trigger
type TriggerUpdateObject struct {
	InvocationRole *string     `json:"invocationRole"`
	Description    *string     `json:"description"`
	TriggerConfig  interface{} `json:"triggerConfig"`
	Qualifier      *string     `json:"qualifier"`

	err error `json:"-"`
}

type UpdateTriggerInput struct {
	ServiceName  *string
	FunctionName *string
	TriggerName  *string
	TriggerUpdateObject
	IfMatch *string
}

func NewUpdateTriggerInput(serviceName string, functionName string, triggerName string) *UpdateTriggerInput {
	return &UpdateTriggerInput{
		ServiceName:  &serviceName,
		FunctionName: &functionName,
		TriggerName:  &triggerName,
	}
}

func (i *UpdateTriggerInput) WithDescription(desc string) *UpdateTriggerInput {
	i.Description = &desc
	return i
}

func (i *UpdateTriggerInput) WithInvocationRole(role string) *UpdateTriggerInput {
	i.InvocationRole = &role
	return i
}

func (i *UpdateTriggerInput) WithTriggerConfig(config interface{}) *UpdateTriggerInput {
	i.TriggerConfig = &config
	return i
}

func (s *UpdateTriggerInput) WithIfMatch(ifMatch string) *UpdateTriggerInput {
	s.IfMatch = &ifMatch
	return s
}

func (i *UpdateTriggerInput) WithQualifier(qualifier string) *UpdateTriggerInput {
	i.Qualifier = &qualifier
	return i
}

func (i *UpdateTriggerInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *UpdateTriggerInput) GetPath() string {
	return fmt.Sprintf(singleTriggerPath, pathEscape(*i.ServiceName), pathEscape(*i.FunctionName), pathEscape(*i.TriggerName))
}

func (i *UpdateTriggerInput) GetHeaders() Header {
	header := make(Header)
	if i.IfMatch != nil {
		header[ifMatch] = *i.IfMatch
	}
	return header
}

func (i *UpdateTriggerInput) GetPayload() interface{} {
	return i.TriggerUpdateObject
}

func (i *UpdateTriggerInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.FunctionName) {
		return fmt.Errorf("Function name is required but not provided")
	}
	if IsBlank(i.TriggerName) {
		return fmt.Errorf("Trigger name is required but not provided")
	}
	if i.err != nil {
		return i.err
	}
	return nil
}

type UpdateTriggerOutput struct {
	Header http.Header
	triggerMetadata
}

func (o UpdateTriggerOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o UpdateTriggerOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o UpdateTriggerOutput) GetEtag() string {
	return GetEtag(o.Header)
}

// MarshalJSON marshals trigger metadata and excludes RawTriggerConfig
func (o UpdateTriggerOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(triggerMetadataDisplay{
		Header:           o.Header,
		TriggerName:      o.TriggerName,
		Description:      o.Description,
		SourceARN:        o.SourceARN,
		TriggerType:      o.TriggerType,
		InvocationRole:   o.InvocationRole,
		Qualifier:        o.Qualifier,
		TriggerConfig:    o.TriggerConfig,
		CreatedTime:      o.CreatedTime,
		LastModifiedTime: o.LastModifiedTime,
	})
}

type ListTriggersInput struct {
	ServiceName  *string
	FunctionName *string
	Query
}

func NewListTriggersInput(serviceName string, functionName string) *ListTriggersInput {
	return &ListTriggersInput{
		ServiceName:  &serviceName,
		FunctionName: &functionName,
	}
}

func (i *ListTriggersInput) WithPrefix(prefix string) *ListTriggersInput {
	i.Prefix = &prefix
	return i
}

func (i *ListTriggersInput) WithStartKey(startKey string) *ListTriggersInput {
	i.StartKey = &startKey
	return i
}

func (i *ListTriggersInput) WithNextToken(nextToken string) *ListTriggersInput {
	i.NextToken = &nextToken
	return i
}

func (i *ListTriggersInput) WithLimit(limit int32) *ListTriggersInput {
	i.Limit = &limit
	return i
}

func (i *ListTriggersInput) GetQueryParams() url.Values {
	out := url.Values{}
	if i.Prefix != nil {
		out.Set("prefix", *i.Prefix)
	}

	if i.StartKey != nil {
		out.Set("startKey", *i.StartKey)
	}

	if i.NextToken != nil {
		out.Set("nextToken", *i.NextToken)
	}

	if i.Limit != nil {
		out.Set("limit", strconv.FormatInt(int64(*i.Limit), 10))
	}

	return out
}

func (i *ListTriggersInput) GetPath() string {
	return fmt.Sprintf(triggersPath, pathEscape(*i.ServiceName), pathEscape(*i.FunctionName))
}

func (i *ListTriggersInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *ListTriggersInput) GetPayload() interface{} {
	return nil
}

func (i *ListTriggersInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.FunctionName) {
		return fmt.Errorf("Function name is required but not provided")
	}
	return nil
}

// ListTriggersOutput defines the trigger response list
type ListTriggersOutput struct {
	Header    http.Header
	Triggers  []*triggerMetadata `json:"triggers"`
	NextToken *string            `json:"nextToken,omitempty"`
}

func (o ListTriggersOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o ListTriggersOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

type DeleteTriggerInput struct {
	ServiceName  *string
	FunctionName *string
	TriggerName  *string
	IfMatch      *string
}

func NewDeleteTriggerInput(serviceName string, functionName string, triggerName string) *DeleteTriggerInput {
	return &DeleteTriggerInput{
		ServiceName:  &serviceName,
		FunctionName: &functionName,
		TriggerName:  &triggerName,
	}
}

func (s *DeleteTriggerInput) WithIfMatch(ifMatch string) *DeleteTriggerInput {
	s.IfMatch = &ifMatch
	return s
}

func (i *DeleteTriggerInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *DeleteTriggerInput) GetPath() string {
	return fmt.Sprintf(singleTriggerPath, pathEscape(*i.ServiceName), pathEscape(*i.FunctionName), pathEscape(*i.TriggerName))
}

func (i *DeleteTriggerInput) GetHeaders() Header {
	header := make(Header)
	if i.IfMatch != nil {
		header[ifMatch] = *i.IfMatch
	}
	return header
}

func (i *DeleteTriggerInput) GetPayload() interface{} {
	return nil
}

func (i *DeleteTriggerInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.FunctionName) {
		return fmt.Errorf("Function name is required but not provided")
	}
	if IsBlank(i.TriggerName) {
		return fmt.Errorf("Trigger name is required but not provided")
	}
	return nil
}

type DeleteTriggerOutput struct {
	Header http.Header
}

func (o DeleteTriggerOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o DeleteTriggerOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}
