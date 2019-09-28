package fc

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	invocationTypeAsync = "Async"
	invocationTypeSync  = "Sync"
)

// Code defines the code location or includes the base64 encoded source
type Code struct {
	OSSBucketName *string `json:"ossBucketName"`
	OSSObjectName *string `json:"ossObjectName"`
	ZipFile       *string `json:"zipFile"`

	err error `json:"-"`
}

func NewCode() *Code {
	return &Code{}
}

func (c *Code) WithOSSBucketName(bucketName string) *Code {
	c.OSSBucketName = &bucketName
	return c
}

func (c *Code) WithOSSObjectName(objectName string) *Code {
	c.OSSObjectName = &objectName
	return c
}

func (c *Code) WithZipFile(zipFile []byte) *Code {
	encodedStr := base64.StdEncoding.EncodeToString(zipFile)
	c.ZipFile = &encodedStr
	return c
}

func (c *Code) WithDir(dir string) *Code {
	zipped := &bytes.Buffer{}
	err := ZipDir(dir, zipped)
	if err != nil {
		c.err = err
		return c
	}

	encoded := base64.StdEncoding.EncodeToString(zipped.Bytes())
	c.ZipFile = &encoded
	return c
}

func (c *Code) WithFiles(files ...string) *Code {
	zipFile, err := TmpZip(files)
	if err != nil {
		c.err = err
		return c
	}
	defer os.Remove(zipFile)
	data, err := ioutil.ReadFile(zipFile)
	if err != nil {
		c.err = err
		return c
	}
	encodedStr := base64.StdEncoding.EncodeToString(data)
	c.ZipFile = &encodedStr
	return c
}

// CreateFunctionInput defines function creation input
type CreateFunctionInput struct {
	ServiceName *string
	FunctionCreateObject
}

type FunctionCreateObject struct {
	FunctionName          *string           `json:"functionName"`
	Description           *string           `json:"description"`
	Runtime               *string           `json:"runtime"`
	Handler               *string           `json:"handler"`
	Timeout               *int32            `json:"timeout"`
	MemorySize            *int32            `json:"memorySize"`
	Code                  *Code             `json:"code"`
	EnvironmentVariables  map[string]string `json:"environmentVariables"`
	Initializer           *string           `json:"initializer"`
	InitializationTimeout *int32            `json:"initializationTimeout"`

	err error `json:"-"`
}

func NewCreateFunctionInput(serviceName string) *CreateFunctionInput {
	return &CreateFunctionInput{ServiceName: &serviceName}
}

func (i *CreateFunctionInput) WithFunctionName(functionName string) *CreateFunctionInput {
	i.FunctionName = &functionName
	return i
}

func (i *CreateFunctionInput) WithDescription(description string) *CreateFunctionInput {
	i.Description = &description
	return i
}

func (i *CreateFunctionInput) WithRuntime(runtime string) *CreateFunctionInput {
	i.Runtime = &runtime
	return i
}

func (i *CreateFunctionInput) WithHandler(handler string) *CreateFunctionInput {
	i.Handler = &handler
	return i
}

func (i *CreateFunctionInput) WithTimeout(timeout int32) *CreateFunctionInput {
	i.Timeout = &timeout
	return i
}

func (i *CreateFunctionInput) WithMemorySize(memory int32) *CreateFunctionInput {
	i.MemorySize = &memory
	return i
}

func (i *CreateFunctionInput) WithCode(code *Code) *CreateFunctionInput {
	if code != nil && code.err != nil {
		i.err = code.err
		return i
	}
	i.Code = code
	return i
}

func (i *CreateFunctionInput) WithEnvironmentVariables(env map[string]string) *CreateFunctionInput {
	i.EnvironmentVariables = env
	return i
}

func (i *CreateFunctionInput) WithInitializer(initializer string) *CreateFunctionInput {
	i.Initializer = &initializer
	return i
}

func (i *CreateFunctionInput) WithInitializationTimeout(initializationTimeout int32) *CreateFunctionInput {
	i.InitializationTimeout = &initializationTimeout
	return i
}

func (i *CreateFunctionInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *CreateFunctionInput) GetPath() string {
	return fmt.Sprintf(functionsPath, pathEscape(*i.ServiceName))
}

func (i *CreateFunctionInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *CreateFunctionInput) GetPayload() interface{} {
	return i.FunctionCreateObject
}

func (i *CreateFunctionInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if i.err != nil {
		return i.err
	}
	return nil
}

type CreateFunctionOutput struct {
	Header http.Header
	functionMetadata
}

func (o CreateFunctionOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o CreateFunctionOutput) GetEtag() string {
	return GetEtag(o.Header)
}

func (o CreateFunctionOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

// FunctionUpdateObject defines update fields in Function
type FunctionUpdateObject struct {
	Description           *string           `json:"description"`
	Runtime               *string           `json:"runtime"`
	Handler               *string           `json:"handler"`
	Timeout               *int32            `json:"timeout"`
	MemorySize            *int32            `json:"memorySize"`
	Code                  *Code             `json:"code"`
	EnvironmentVariables  map[string]string `json:"environmentVariables"`
	Initializer           *string           `json:"initializer"`
	InitializationTimeout *int32            `json:"initializationTimeout"`

	err error `json:"-"`
}

type UpdateFunctionInput struct {
	ServiceName  *string
	FunctionName *string
	FunctionUpdateObject
	IfMatch *string
}

func NewUpdateFunctionInput(serviceName string, functionName string) *UpdateFunctionInput {
	return &UpdateFunctionInput{
		ServiceName:  &serviceName,
		FunctionName: &functionName,
	}
}

func (i *UpdateFunctionInput) WithDescription(description string) *UpdateFunctionInput {
	i.Description = &description
	return i
}

func (i *UpdateFunctionInput) WithRuntime(runtime string) *UpdateFunctionInput {
	i.Runtime = &runtime
	return i
}

func (i *UpdateFunctionInput) WithHandler(handler string) *UpdateFunctionInput {
	i.Handler = &handler
	return i
}

func (i *UpdateFunctionInput) WithTimeout(timeout int32) *UpdateFunctionInput {
	i.Timeout = &timeout
	return i
}

func (i *UpdateFunctionInput) WithMemorySize(memory int32) *UpdateFunctionInput {
	i.MemorySize = &memory
	return i
}

func (i *UpdateFunctionInput) WithCode(code *Code) *UpdateFunctionInput {
	if code != nil && code.err != nil {
		i.err = code.err
		return i
	}
	i.Code = code
	return i
}

func (i *UpdateFunctionInput) WithEnvironmentVariables(env map[string]string) *UpdateFunctionInput {
	i.EnvironmentVariables = env
	return i
}

func (i *UpdateFunctionInput) WithInitializer(initializer string) *UpdateFunctionInput {
	i.Initializer = &initializer
	return i
}

func (i *UpdateFunctionInput) WithInitializationTimeout(initializationTimeout int32) *UpdateFunctionInput {
	i.InitializationTimeout = &initializationTimeout
	return i
}

func (s *UpdateFunctionInput) WithIfMatch(ifMatch string) *UpdateFunctionInput {
	s.IfMatch = &ifMatch
	return s
}

func (i *UpdateFunctionInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *UpdateFunctionInput) GetPath() string {
	return fmt.Sprintf(singleFunctionPath, pathEscape(*i.ServiceName), pathEscape(*i.FunctionName))
}

func (i *UpdateFunctionInput) GetHeaders() Header {
	header := make(Header)
	if i.IfMatch != nil {
		header[ifMatch] = *i.IfMatch
	}
	return header
}

func (i *UpdateFunctionInput) GetPayload() interface{} {
	return i.FunctionUpdateObject
}

func (i *UpdateFunctionInput) Validate() error {
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

type UpdateFunctionOutput struct {
	Header http.Header
	functionMetadata
}

func (o UpdateFunctionOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o UpdateFunctionOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o UpdateFunctionOutput) GetEtag() string {
	return GetEtag(o.Header)
}

type GetFunctionInput struct {
	ServiceName  *string
	FunctionName *string
	Qualifier    *string
}

func NewGetFunctionInput(serviceName string, functionName string) *GetFunctionInput {
	return &GetFunctionInput{
		ServiceName:  &serviceName,
		FunctionName: &functionName,
	}
}

func (i *GetFunctionInput) WithQualifier(qualifier string) *GetFunctionInput {
	i.Qualifier = &qualifier
	return i
}

func (i *GetFunctionInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *GetFunctionInput) GetPath() string {
	if !IsBlank(i.Qualifier) {
		return fmt.Sprintf(singleFunctionWithQualifierPath,
			pathEscape(*i.ServiceName), pathEscape(*i.Qualifier), pathEscape(*i.FunctionName))
	}
	return fmt.Sprintf(singleFunctionPath, pathEscape(*i.ServiceName), pathEscape(*i.FunctionName))
}

func (i *GetFunctionInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *GetFunctionInput) GetPayload() interface{} {
	return nil
}

func (i *GetFunctionInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.FunctionName) {
		return fmt.Errorf("Function name is required but not provided")
	}
	return nil
}

// GetFunctionOutput define function response from fc
type GetFunctionOutput struct {
	Header http.Header
	functionMetadata
}

func (o GetFunctionOutput) GetEtag() string {
	return GetEtag(o.Header)
}

func (o GetFunctionOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o GetFunctionOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

// functionMetadata define the function metadata
type functionMetadata struct {
	FunctionID            *string           `json:"functionId"`
	FunctionName          *string           `json:"functionName"`
	Description           *string           `json:"description"`
	Runtime               *string           `json:"runtime"`
	Handler               *string           `json:"handler"`
	Timeout               *int32            `json:"timeout"`
	MemorySize            *int32            `json:"memorySize"`
	CodeSize              *int64            `json:"codeSize"`
	CodeChecksum          *string           `json:"codeChecksum"`
	EnvironmentVariables  map[string]string `json:"environmentVariables"`
	CreatedTime           *string           `json:"createdTime"`
	LastModifiedTime      *string           `json:"lastModifiedTime"`
	Initializer           *string           `json:"initializer"`
	InitializationTimeout *int32            `json:"initializationTimeout"`
}

// GetFunctionCodeInput ...
type GetFunctionCodeInput struct {
	*GetFunctionInput
}

// NewGetFunctionCodeInput ...
func NewGetFunctionCodeInput(serviceName string, functionName string) *GetFunctionCodeInput {
	return &GetFunctionCodeInput{
		&GetFunctionInput{
			ServiceName:  &serviceName,
			FunctionName: &functionName,
		},
	}
}

func (i *GetFunctionCodeInput) WithQualifier(qualifier string) *GetFunctionCodeInput {
	i.Qualifier = &qualifier
	return i
}

// GetPath ...
func (i *GetFunctionCodeInput) GetPath() string {
	if !IsBlank(i.Qualifier) {
		return fmt.Sprintf(functionCodeWithQualifierPath,
			pathEscape(*i.ServiceName), pathEscape(*i.Qualifier), pathEscape(*i.FunctionName))
	}
	return fmt.Sprintf(functionCodePath, pathEscape(*i.ServiceName), pathEscape(*i.FunctionName))
}

type functionCodeMetadata struct {
	URL string `json:"url"`
}

// GetFunctionCodeOutput define function response from fc
type GetFunctionCodeOutput struct {
	Header http.Header
	functionCodeMetadata
}

// GetRequestID ...
func (o GetFunctionCodeOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

// String ...
func (o GetFunctionCodeOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

// ListFunctionsOutput defines the function response list
type ListFunctionsOutput struct {
	Header    http.Header
	Functions []*functionMetadata `json:"functions"`
	NextToken *string             `json:"nextToken,omitempty"`
}

func (o ListFunctionsOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o ListFunctionsOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

type ListFunctionsInput struct {
	ServiceName *string
	Qualifier   *string
	Query
}

func NewListFunctionsInput(serviceName string) *ListFunctionsInput {
	return &ListFunctionsInput{ServiceName: &serviceName}
}

func (i *ListFunctionsInput) WithQualifier(qualifier string) *ListFunctionsInput {
	i.Qualifier = &qualifier
	return i
}

func (i *ListFunctionsInput) WithPrefix(prefix string) *ListFunctionsInput {
	i.Prefix = &prefix
	return i
}

func (i *ListFunctionsInput) WithStartKey(startKey string) *ListFunctionsInput {
	i.StartKey = &startKey
	return i
}

func (i *ListFunctionsInput) WithNextToken(nextToken string) *ListFunctionsInput {
	i.NextToken = &nextToken
	return i
}

func (i *ListFunctionsInput) WithLimit(limit int32) *ListFunctionsInput {
	i.Limit = &limit
	return i
}

func (i *ListFunctionsInput) GetQueryParams() url.Values {
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

func (i *ListFunctionsInput) GetPath() string {
	if !IsBlank(i.Qualifier) {
		return fmt.Sprintf(functionsPathWithQualifierPath,
			pathEscape(*i.ServiceName), pathEscape(*i.Qualifier))
	}
	return fmt.Sprintf(functionsPath, pathEscape(*i.ServiceName))
}

func (i *ListFunctionsInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *ListFunctionsInput) GetPayload() interface{} {
	return nil
}

func (i *ListFunctionsInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	return nil
}

type DeleteFunctionInput struct {
	ServiceName  *string
	FunctionName *string
	IfMatch      *string
}

func NewDeleteFunctionInput(serviceName string, functionName string) *DeleteFunctionInput {
	return &DeleteFunctionInput{
		ServiceName:  &serviceName,
		FunctionName: &functionName,
	}
}

func (s *DeleteFunctionInput) WithIfMatch(ifMatch string) *DeleteFunctionInput {
	s.IfMatch = &ifMatch
	return s
}

func (i *DeleteFunctionInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *DeleteFunctionInput) GetPath() string {
	return fmt.Sprintf(singleFunctionPath, pathEscape(*i.ServiceName), pathEscape(*i.FunctionName))
}

func (i *DeleteFunctionInput) GetHeaders() Header {
	header := make(Header)
	if i.IfMatch != nil {
		header[ifMatch] = *i.IfMatch
	}
	return header
}

func (i *DeleteFunctionInput) GetPayload() interface{} {
	return nil
}

func (i *DeleteFunctionInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.FunctionName) {
		return fmt.Errorf("Function name is required but not provided")
	}
	return nil
}

type DeleteFunctionOutput struct {
	Header http.Header
}

func (o DeleteFunctionOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o DeleteFunctionOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

type InvokeFunctionInput struct {
	ServiceName  *string
	FunctionName *string
	Qualifier    *string
	Payload      *[]byte
	headers      Header
}

func (i *InvokeFunctionInput) WithQualifier(qualifier string) *InvokeFunctionInput {
	i.Qualifier = &qualifier
	return i
}

func NewInvokeFunctionInput(serviceName string, functionName string) *InvokeFunctionInput {
	return &InvokeFunctionInput{
		ServiceName:  &serviceName,
		FunctionName: &functionName,
		headers:      make(Header),
	}
}

func (i *InvokeFunctionInput) WithPayload(payload []byte) *InvokeFunctionInput {
	i.Payload = &payload
	return i
}

func (i *InvokeFunctionInput) WithInvocationType(invocationType string) *InvokeFunctionInput {
	i.headers[HTTPHeaderInvocationType] = invocationType
	return i
}

func (i *InvokeFunctionInput) WithLogType(logType string) *InvokeFunctionInput {
	i.headers[HTTPHeaderInvocationLogType] = logType
	return i
}

func (i *InvokeFunctionInput) WithHeader(key, value string) *InvokeFunctionInput {
	i.headers[key] = value
	return i
}

func (i *InvokeFunctionInput) WithAsyncInvocation() *InvokeFunctionInput {
	return i.WithInvocationType(invocationTypeAsync)
}

func (i *InvokeFunctionInput) WithSyncInvocation() *InvokeFunctionInput {
	return i.WithInvocationType(invocationTypeSync)
}

func (i *InvokeFunctionInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *InvokeFunctionInput) GetPath() string {
	if !IsBlank(i.Qualifier) {
		return fmt.Sprintf(invokeFunctionWithQualifierPath,
			pathEscape(*i.ServiceName), pathEscape(*i.Qualifier), pathEscape(*i.FunctionName))
	}
	return fmt.Sprintf(invokeFunctionPath, pathEscape(*i.ServiceName), pathEscape(*i.FunctionName))
}

func (i *InvokeFunctionInput) GetHeaders() Header {
	return i.headers
}

func (i *InvokeFunctionInput) GetPayload() interface{} {

	if i.Payload == nil || len(*i.Payload) <= 0 {
		// returning explicit untyped nil instead of i.Payload (interface nil)
		// see https://golang.org/doc/faq#nil_error
		return nil
	}
	return i.Payload
}

func (i *InvokeFunctionInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	if IsBlank(i.FunctionName) {
		return fmt.Errorf("Function name is required but not provided")
	}
	return nil
}

type InvokeFunctionOutput struct {
	Header  http.Header
	Payload []byte
}

func (o InvokeFunctionOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o InvokeFunctionOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

// GetErrorType returns error type occurred in function invocations
// will be empty string when no errors
func (o InvokeFunctionOutput) GetErrorType() string {
	return GetErrorType(o.Header)
}

// GetLogResult returns LogResults for the invocation
func (o InvokeFunctionOutput) GetLogResult() (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(o.Header.Get(HTTPHeaderInvocationLogResult))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
