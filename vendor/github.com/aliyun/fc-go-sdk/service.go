package fc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	servicesPath       = "/services"
	singleServicePath  = servicesPath + "/%s"
	functionsPath      = singleServicePath + "/functions"
	singleFunctionPath = functionsPath + "/%s"
	functionCodePath   = singleFunctionPath + "/code"
	triggersPath       = singleFunctionPath + "/triggers"
	singleTriggerPath  = triggersPath + "/%s"
	invokeFunctionPath = singleFunctionPath + "/invocations"
	versionsPath       = singleServicePath + "/versions"
	singleVersionPath  = versionsPath + "/%s"
	aliasesPath        = singleServicePath + "/aliases"
	singleAliasPath    = aliasesPath + "/%s"

	singleServiceWithQualifierPath  = servicesPath + "/%s.%s"
	functionsPathWithQualifierPath  = singleServiceWithQualifierPath + "/functions"
	singleFunctionWithQualifierPath = functionsPathWithQualifierPath + "/%s"
	functionCodeWithQualifierPath   = singleFunctionWithQualifierPath + "/code"
	invokeFunctionWithQualifierPath = singleFunctionWithQualifierPath + "/invocations"

	printIndent = "  "

	ifMatch = "If-Match"
)

type ServiceInput interface {
	GetQueryParams() url.Values
	GetPath() string
	GetHeaders() Header
	GetPayload() interface{}
	Validate() error
}

// LogConfig defines the log config for service
type LogConfig struct {
	Project  *string `json:"project"`
	Logstore *string `json:"logstore"`
}

func NewLogConfig() *LogConfig {
	return &LogConfig{}
}

func (l *LogConfig) WithProject(project string) *LogConfig {
	l.Project = &project
	return l
}

func (l *LogConfig) WithLogstore(logstore string) *LogConfig {
	l.Logstore = &logstore
	return l
}

// VPCConfig defines the VPC config for service
type VPCConfig struct {
	VPCID           *string  `json:"vpcId"`
	VSwitchIDs      []string `json:"vSwitchIds"`
	SecurityGroupID *string  `json:"securityGroupId"`
}

func NewVPCConfig() *VPCConfig {
	return &VPCConfig{}
}

func (l *VPCConfig) WithVPCID(vpcID string) *VPCConfig {
	l.VPCID = &vpcID
	return l
}

func (l *VPCConfig) WithVSwitchIDs(vSwitchIDs []string) *VPCConfig {
	l.VSwitchIDs = vSwitchIDs
	return l
}

func (l *VPCConfig) WithSecurityGroupID(securityGroupID string) *VPCConfig {
	l.SecurityGroupID = &securityGroupID
	return l
}

// NASMountConfig defines the nas binding info for service
type NASMountConfig struct {
	ServerAddr string `json:"serverAddr"`
	MountDir   string `json:"mountDir"`
}

func NewNASMountConfig(serverAddr, mountDir string) NASMountConfig {
	return NASMountConfig{
		ServerAddr: serverAddr,
		MountDir:   mountDir,
	}
}

// NASConfig defines the NAS config info
// UserID/GroupID is the uid/gid of the user access the NFS file system
type NASConfig struct {
	UserID      *int32           `json:"userId"`
	GroupID     *int32           `json:"groupId"`
	MountPoints []NASMountConfig `json:"mountPoints"`
}

func NewNASConfig() *NASConfig {
	return &NASConfig{}
}

func (n *NASConfig) WithUserID(userID int32) *NASConfig {
	n.UserID = &userID
	return n
}

func (n *NASConfig) WithGroupID(groupID int32) *NASConfig {
	n.GroupID = &groupID
	return n
}

func (n *NASConfig) WithMountPoints(mountPoints []NASMountConfig) *NASConfig {
	n.MountPoints = mountPoints
	return n
}

// CreateServiceInput defines input to create service
type CreateServiceInput struct {
	ServiceName    *string    `json:"serviceName"`
	Description    *string    `json:"description"`
	Role           *string    `json:"role"`
	LogConfig      *LogConfig `json:"logConfig"`
	VPCConfig      *VPCConfig `json:"vpcConfig"`
	InternetAccess *bool      `json:"internetAccess"`
	NASConfig      *NASConfig `json:"nasConfig"`
}

func NewCreateServiceInput() *CreateServiceInput {
	return &CreateServiceInput{}
}

func (s *CreateServiceInput) WithServiceName(serviceName string) *CreateServiceInput {
	s.ServiceName = &serviceName
	return s
}

func (s *CreateServiceInput) WithDescription(description string) *CreateServiceInput {
	s.Description = &description
	return s
}

func (s *CreateServiceInput) WithRole(role string) *CreateServiceInput {
	s.Role = &role
	return s
}

func (s *CreateServiceInput) WithLogConfig(logConfig *LogConfig) *CreateServiceInput {
	s.LogConfig = logConfig
	return s
}

func (s *CreateServiceInput) WithVPCConfig(vpcConfig *VPCConfig) *CreateServiceInput {
	s.VPCConfig = vpcConfig
	return s
}

func (s *CreateServiceInput) WithNASConfig(nasConfig *NASConfig) *CreateServiceInput {
	s.NASConfig = nasConfig
	return s
}

func (s *CreateServiceInput) WithInternetAccess(access bool) *CreateServiceInput {
	s.InternetAccess = &access
	return s
}

func (i *CreateServiceInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *CreateServiceInput) GetPath() string {
	return servicesPath
}

func (i *CreateServiceInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *CreateServiceInput) GetPayload() interface{} {
	return i
}

func (i *CreateServiceInput) Validate() error {
	return nil
}

// CreateServiceOutput define get service response
type CreateServiceOutput struct {
	Header http.Header
	serviceMetadata
}

func (o CreateServiceOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o CreateServiceOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o CreateServiceOutput) GetEtag() string {
	return GetEtag(o.Header)
}

// ServiceUpdateObject defines the service update fields
type ServiceUpdateObject struct {
	Description    *string    `json:"description"`
	Role           *string    `json:"role"`
	LogConfig      *LogConfig `json:"logConfig"`
	VPCConfig      *VPCConfig `json:"vpcConfig"`
	InternetAccess *bool      `json:"internetAccess"`
	NASConfig      *NASConfig `json:"nasConfig"`
}

type UpdateServiceInput struct {
	ServiceName *string
	ServiceUpdateObject
	IfMatch *string
}

func NewUpdateServiceInput(serviceName string) *UpdateServiceInput {
	return &UpdateServiceInput{ServiceName: &serviceName}
}

func (s *UpdateServiceInput) WithDescription(description string) *UpdateServiceInput {
	s.Description = &description
	return s
}

func (s *UpdateServiceInput) WithRole(role string) *UpdateServiceInput {
	s.Role = &role
	return s
}

func (s *UpdateServiceInput) WithLogConfig(logConfig *LogConfig) *UpdateServiceInput {
	s.LogConfig = logConfig
	return s
}

func (s *UpdateServiceInput) WithVPCConfig(vpcConfig *VPCConfig) *UpdateServiceInput {
	s.VPCConfig = vpcConfig
	return s
}

func (s *UpdateServiceInput) WithNASConfig(nasConfig *NASConfig) *UpdateServiceInput {
	s.NASConfig = nasConfig
	return s
}

func (s *UpdateServiceInput) WithInternetAccess(access bool) *UpdateServiceInput {
	s.InternetAccess = &access
	return s
}

func (s *UpdateServiceInput) WithIfMatch(ifMatch string) *UpdateServiceInput {
	s.IfMatch = &ifMatch
	return s
}

func (i *UpdateServiceInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *UpdateServiceInput) GetPath() string {
	return fmt.Sprintf(singleServicePath, pathEscape(*i.ServiceName))
}

func (i *UpdateServiceInput) GetHeaders() Header {
	header := make(Header)
	if i.IfMatch != nil {
		header[ifMatch] = *i.IfMatch
	}
	return header
}

func (i *UpdateServiceInput) GetPayload() interface{} {
	return i.ServiceUpdateObject
}

func (i *UpdateServiceInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	return nil
}

// UpdateServiceOutput define get service response
type UpdateServiceOutput struct {
	Header http.Header
	serviceMetadata
}

func (o UpdateServiceOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o UpdateServiceOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o UpdateServiceOutput) GetEtag() string {
	return GetEtag(o.Header)
}

// GetServiceOutput define get service response
type GetServiceOutput struct {
	Header http.Header
	serviceMetadata
}

func (o GetServiceOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o GetServiceOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

func (o GetServiceOutput) GetEtag() string {
	return GetEtag(o.Header)
}

// serviceMetadata defines the detail service object
type serviceMetadata struct {
	ServiceName      *string    `json:"serviceName"`
	Description      *string    `json:"description"`
	Role             *string    `json:"role"`
	LogConfig        *LogConfig `json:"logConfig"`
	VPCConfig        *VPCConfig `json:"vpcConfig"`
	InternetAccess   *bool      `json:"internetAccess"`
	ServiceID        *string    `json:"serviceId"`
	CreatedTime      *string    `json:"createdTime"`
	LastModifiedTime *string    `json:"lastModifiedTime"`
	NASConfig        *NASConfig `json:"nasConfig"`
}

// ListServicesOutput defines listServiceMetadata result
type ListServicesOutput struct {
	Header    http.Header
	Services  []*serviceMetadata `json:"services"`
	NextToken *string            `json:"nextToken,omitempty"`
}

func (o ListServicesOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}
func (o ListServicesOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}

type ListServicesInput struct {
	Query
}

func NewListServicesInput() *ListServicesInput {
	return &ListServicesInput{}
}

func (i *ListServicesInput) WithPrefix(prefix string) *ListServicesInput {
	i.Prefix = &prefix
	return i
}

func (i *ListServicesInput) WithStartKey(startKey string) *ListServicesInput {
	i.StartKey = &startKey
	return i
}

func (i *ListServicesInput) WithNextToken(nextToken string) *ListServicesInput {
	i.NextToken = &nextToken
	return i
}

func (i *ListServicesInput) WithLimit(limit int32) *ListServicesInput {
	i.Limit = &limit
	return i
}

func (i *ListServicesInput) GetQueryParams() url.Values {
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

func (i *ListServicesInput) GetPath() string {
	return servicesPath
}

func (i *ListServicesInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *ListServicesInput) GetPayload() interface{} {
	return nil
}

func (i *ListServicesInput) Validate() error {
	return nil
}

type GetServiceInput struct {
	ServiceName *string
	Qualifier   *string
}

func NewGetServiceInput(serviceName string) *GetServiceInput {
	return &GetServiceInput{ServiceName: &serviceName}
}

func (i *GetServiceInput) WithQualifier(qualifier string) *GetServiceInput {
	i.Qualifier = &qualifier
	return i
}

func (i *GetServiceInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *GetServiceInput) GetPath() string {
	if !IsBlank(i.Qualifier) {
		return fmt.Sprintf(singleServiceWithQualifierPath, pathEscape(*i.ServiceName), pathEscape(*i.Qualifier))
	}
	return fmt.Sprintf(singleServicePath, pathEscape(*i.ServiceName))
}

func (i *GetServiceInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *GetServiceInput) GetPayload() interface{} {
	return nil
}

func (i *GetServiceInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	return nil
}

type DeleteServiceInput struct {
	ServiceName *string
	IfMatch     *string
}

func NewDeleteServiceInput(serviceName string) *DeleteServiceInput {
	return &DeleteServiceInput{ServiceName: &serviceName}
}

func (s *DeleteServiceInput) WithIfMatch(ifMatch string) *DeleteServiceInput {
	s.IfMatch = &ifMatch
	return s
}

func (i *DeleteServiceInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (i *DeleteServiceInput) GetPath() string {
	return fmt.Sprintf(singleServicePath, pathEscape(*i.ServiceName))
}

func (i *DeleteServiceInput) GetHeaders() Header {
	header := make(Header)
	if i.IfMatch != nil {
		header[ifMatch] = *i.IfMatch
	}
	return header
}

func (i *DeleteServiceInput) GetPayload() interface{} {
	return nil
}

func (i *DeleteServiceInput) Validate() error {
	if IsBlank(i.ServiceName) {
		return fmt.Errorf("Service name is required but not provided")
	}
	return nil
}

type DeleteServiceOutput struct {
	Header http.Header
}

func (o DeleteServiceOutput) String() string {
	b, err := json.MarshalIndent(o, "", printIndent)
	if err != nil {
		return ""
	}
	return string(b)
}

func (o DeleteServiceOutput) GetRequestID() string {
	return GetRequestID(o.Header)
}
