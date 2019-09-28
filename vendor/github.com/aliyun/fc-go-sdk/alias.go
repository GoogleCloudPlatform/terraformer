package fc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type aliasMetadata struct {
	AliasName               *string           `json:"aliasName"`
	VersionID               *string            `json:"versionId"`
	Description             *string           `json:"description"`
	AdditionalVersionWeight map[string]float64 `json:"additionalVersionWeight"`
}

type AliasCreateObject struct {
	AliasName               *string           `json:"aliasName"`
	VersionID               *string            `json:"versionId"`
	Description             *string           `json:"description"`
	AdditionalVersionWeight map[string]float64 `json:"additionalVersionWeight"`
}

type CreateAliasInput struct {
	ServiceName             *string           `json:"serviceName"`
	AliasCreateObject
}

func NewCreateAliasInput(serviceName string) *CreateAliasInput {
	return &CreateAliasInput{ServiceName: &serviceName}
}

func (i *CreateAliasInput) WithAliasName(aliasName string) *CreateAliasInput {
	i.AliasName = &aliasName
	return i
}

func (i *CreateAliasInput) WithDescription(description string) *CreateAliasInput {
	i.Description = &description
	return i
}

func (i *CreateAliasInput) WithVersionID(versionID string) *CreateAliasInput {
	i.VersionID = &versionID
	return i
}

func (i *CreateAliasInput) WithAdditionalVersionWeight(additionalVersionWeight map[string]float64) *CreateAliasInput {
	i.AdditionalVersionWeight = additionalVersionWeight
	return i
}

func (i *CreateAliasInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *CreateAliasInput) GetPath() string {
	return fmt.Sprintf(aliasesPath, pathEscape(*i.ServiceName))
}

func (i *CreateAliasInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *CreateAliasInput) GetPayload() interface{} {
	return i.AliasCreateObject
}

func (i *CreateAliasInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	return nil
}

type CreateAliasOutput struct {
	Header http.Header
	aliasMetadata
}

func (o CreateAliasOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o CreateAliasOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o CreateAliasOutput) GetEtag() string {
	return GetEtag(o.Header)
}

type AliasUpdateObject struct {
	VersionID               *string            `json:"versionId"`
	Description             *string           `json:"description"`
	AdditionalVersionWeight map[string]float64 `json:"additionalVersionWeight"`
}

type UpdateAliasInput struct {
	ServiceName *string
	AliasName   *string
	AliasUpdateObject
	IfMatch     *string
}

func NewUpdateAliasInput(serviceName, aliasName string) *UpdateAliasInput {
	return &UpdateAliasInput{ServiceName: &serviceName, AliasName: &aliasName}
}

func (s *UpdateAliasInput) WithDescription(description string) *UpdateAliasInput {
	s.Description = &description
	return s
}

func (s *UpdateAliasInput) WithVersionID(versionID string) *UpdateAliasInput {
	s.VersionID = &versionID
	return s
}

func (s *UpdateAliasInput) WithAdditionalVersionWeight(additionalVersionWeight map[string]float64) *UpdateAliasInput {
	s.AdditionalVersionWeight = additionalVersionWeight
	return s
}

func (s *UpdateAliasInput) WithIfMatch(ifMatch string) *UpdateAliasInput {
	s.IfMatch = &ifMatch
	return s
}

func (i *UpdateAliasInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *UpdateAliasInput) GetPath() string {
	return fmt.Sprintf(singleAliasPath, pathEscape(*i.ServiceName), pathEscape(*i.AliasName))
}

func (i *UpdateAliasInput) GetHeaders() Header {
	header := make(Header)
	if i.IfMatch != nil {
		header[ifMatch] = *i.IfMatch
	}
	return header
}

func (i *UpdateAliasInput) GetPayload() interface{} {
	return i.AliasUpdateObject
}

func (i *UpdateAliasInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.AliasName) {
		return fmt.Errorf("Alias name is required but not provided")
	}
	return nil
}

type UpdateAliasOutput struct {
	Header http.Header
	aliasMetadata
}

func (o UpdateAliasOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o UpdateAliasOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o UpdateAliasOutput) GetEtag() string {
	return GetEtag(o.Header)
}

type ListAliasesInput struct {
	ServiceName *string
	Query
}

func NewListAliasesInput(serviceName string) *ListAliasesInput {
	return &ListAliasesInput{ServiceName: &serviceName}
}

func (i *ListAliasesInput) WithPrefix(prefix string) *ListAliasesInput {
	i.Prefix = &prefix
	return i
}

func (i *ListAliasesInput) WithStartKey(startKey string) *ListAliasesInput {
	i.StartKey = &startKey
	return i
}

func (i *ListAliasesInput) WithNextToken(nextToken string) *ListAliasesInput {
	i.NextToken = &nextToken
	return i
}

func (i *ListAliasesInput) WithLimit(limit int32) *ListAliasesInput {
	i.Limit = &limit
	return i
}

func (i *ListAliasesInput) GetQueryParams() url.Values {
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

func (i *ListAliasesInput) GetPath() string {
	return fmt.Sprintf(aliasesPath, pathEscape(*i.ServiceName))
}

func (i *ListAliasesInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *ListAliasesInput) GetPayload() interface{} {
	return nil
}

func (i *ListAliasesInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	return nil
}

type ListAliasesOutput struct {
	Header      http.Header
	Aliases     []*aliasMetadata `json:"aliases"`
	NextToken   *string          `json:"nextToken,omitempty"`
}

func (o ListAliasesOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o ListAliasesOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

type GetAliasInput struct {
	ServiceName *string
	AliasName   *string
}

func NewGetAliasInput(serviceName, aliasName string) *GetAliasInput {
	return &GetAliasInput{ServiceName: &serviceName, AliasName: &aliasName}
}

func (i *GetAliasInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *GetAliasInput) GetPath() string {
	return fmt.Sprintf(singleAliasPath, pathEscape(*i.ServiceName), pathEscape(*i.AliasName))
}

func (i *GetAliasInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *GetAliasInput) GetPayload() interface{} {
	return nil
}

func (i *GetAliasInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.AliasName) {
		return fmt.Errorf("Alias name is required but not provided")
	}
	return nil
}

type GetAliasOutput struct {
	Header http.Header
	aliasMetadata
}

func (o GetAliasOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o GetAliasOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o GetAliasOutput) GetEtag() string {
	return GetEtag(o.Header)
}

type DeleteAliasInput struct {
	ServiceName *string
	AliasName   *string
	IfMatch     *string
}

func NewDeleteAliasInput(serviceName, aliasName string) *DeleteAliasInput {
	return &DeleteAliasInput{ServiceName: &serviceName, AliasName: &aliasName}
}

func (s *DeleteAliasInput) WithIfMatch(ifMatch string) *DeleteAliasInput {
	s.IfMatch = &ifMatch
	return s
}

func (i *DeleteAliasInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *DeleteAliasInput) GetPath() string {
	return fmt.Sprintf(singleAliasPath, pathEscape(*i.ServiceName), pathEscape(*i.AliasName))
}

func (i *DeleteAliasInput) GetHeaders() Header {
	header := make(Header)
	if i.IfMatch != nil {
		header[ifMatch] = *i.IfMatch
	}
	return header
}

func (i *DeleteAliasInput) GetPayload() interface{} {
	return nil
}

func (i *DeleteAliasInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.AliasName) {
		return fmt.Errorf("Alias name is required but not provided")
	}
	return nil
}

type DeleteAliasOutput struct {
	Header http.Header
}

func (o DeleteAliasOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o DeleteAliasOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}
