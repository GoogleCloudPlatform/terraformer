package fc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

var (
	ListDirectionBackward = "BACKWARD"
	ListDirectionForward = "FORWARD"
)

type versionMetadata struct {
	VersionID        *string     `json:"versionId"`
	Description      *string    `json:"description"`
	CreatedTime      *string    `json:"createdTime"`
	LastModifiedTime *string    `json:"lastModifiedTime"`
}

type ServiceVersionPublishObject struct {
	Description *string
}

type PublishServiceVersionInput struct {
	ServiceName *string
	ServiceVersionPublishObject
	IfMatch     *string
}

func NewPublishServiceVersionInput(serviceName string) *PublishServiceVersionInput {
	return &PublishServiceVersionInput{ServiceName: &serviceName}
}

func (i *PublishServiceVersionInput) WithDescription(description string) *PublishServiceVersionInput {
	i.Description = &description
	return i
}

func (i *PublishServiceVersionInput) WithIfMatch(ifMatch string) *PublishServiceVersionInput {
	i.IfMatch = &ifMatch
	return i
}

func (i *PublishServiceVersionInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *PublishServiceVersionInput) GetPath() string {
	return fmt.Sprintf(versionsPath, pathEscape(*i.ServiceName))
}

func (i *PublishServiceVersionInput) GetHeaders() Header {
	header := make(Header)
	if i.IfMatch != nil {
		header[ifMatch] = *i.IfMatch
	}
	return header
}

func (i *PublishServiceVersionInput) GetPayload() interface{} {
	return i.ServiceVersionPublishObject
}

func (i *PublishServiceVersionInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	return nil
}

type PublishServiceVersionOutput struct {
	Header http.Header
	versionMetadata
}

func (o PublishServiceVersionOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o PublishServiceVersionOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o PublishServiceVersionOutput) GetEtag() string {
	return GetEtag(o.Header)
}

type ListServiceVersionsInput struct {
	ServiceName *string
	StartKey    *string
	NextToken   *string
	Limit       *int32
	Direction   *string
}

func NewListServiceVersionsInput(serviceName string) *ListServiceVersionsInput {
	return &ListServiceVersionsInput{ServiceName: &serviceName}
}

func (i *ListServiceVersionsInput) WithStartKey(startKey string) *ListServiceVersionsInput {
	i.StartKey = &startKey
	return i
}

func (i *ListServiceVersionsInput) WithNextToken(nextToken string) *ListServiceVersionsInput {
	i.NextToken = &nextToken
	return i
}

func (i *ListServiceVersionsInput) WithLimit(limit int32) *ListServiceVersionsInput {
	i.Limit = &limit
	return i
}

func (i *ListServiceVersionsInput) WithBackwardDirection() *ListServiceVersionsInput {
	i.Direction = &ListDirectionBackward
	return i
}

func (i *ListServiceVersionsInput) WithForwardDirection() *ListServiceVersionsInput {
	i.Direction = &ListDirectionForward
	return i
}

func (i *ListServiceVersionsInput) GetQueryParams() url.Values {
	out := url.Values{}

	if i.StartKey != nil {
		out.Set("startKey", *i.StartKey)
	}

	if i.NextToken != nil {
		out.Set("nextToken", *i.NextToken)
	}

	if i.Limit != nil {
		out.Set("limit", strconv.FormatInt(int64(*i.Limit), 10))
	}

	if i.Direction != nil {
		out.Set("direction", *i.Direction)
	}

	return out
}

func (i *ListServiceVersionsInput) GetPath() string {
	return fmt.Sprintf(versionsPath, pathEscape(*i.ServiceName))
}

func (i *ListServiceVersionsInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *ListServiceVersionsInput) GetPayload() interface{} {
	return nil
}

func (i *ListServiceVersionsInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	return nil
}

type ListServiceVersionsOutput struct {
	Header    http.Header
	Versions  []*versionMetadata `json:"versions"`
	NextToken *string            `json:"nextToken,omitempty"`
	Direction *string            `json:"direction"`
}

func (o ListServiceVersionsOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o ListServiceVersionsOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

type DeleteServiceVersionInput struct {
	ServiceName *string
	VersionID   *string
}

func NewDeleteServiceVersionInput(serviceName string, versionID string) *DeleteServiceVersionInput {
	return &DeleteServiceVersionInput{ServiceName: &serviceName, VersionID: &versionID}
}

func (i *DeleteServiceVersionInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *DeleteServiceVersionInput) GetPath() string {
	return fmt.Sprintf(singleVersionPath, pathEscape(*i.ServiceName), pathEscape(*i.VersionID))
}

func (i *DeleteServiceVersionInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *DeleteServiceVersionInput) GetPayload() interface{} {
	return nil
}

func (i *DeleteServiceVersionInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.VersionID) {
		return fmt.Errorf("Version ID is required but not provided")
	}
	return nil
}

type DeleteServiceVersionOutput struct {
	Header http.Header
}

func (o DeleteServiceVersionOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o DeleteServiceVersionOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

