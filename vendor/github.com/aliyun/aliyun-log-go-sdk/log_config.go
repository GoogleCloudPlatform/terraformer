package sls

import (
	"encoding/json"
	"errors"
)

// const InputTypes
const (
	InputTypeSyslog    = "syslog"
	InputTypeStreamlog = "streamlog"
	InputTypePlugin    = "plugin"
	InputTypeFile      = "file"
)

// const LogFileTypes
const (
	LogFileTypeApsaraLog    = "apsara_log"
	LogFileTypeRegexLog     = "common_reg_log"
	LogFileTypeJSONLog      = "json_log"
	LogFileTypeDelimiterLog = "delimiter_log"
)

// const OutputType
const (
	OutputTypeLogService = "LogService"
)

// const MergeType
const (
	MergeTypeTopic    = "topic"
	MergeTypeLogstore = "logstore"
)

const (
	TopicFormatNone         = "none"        // no topic
	TopicFormatMachineGroup = "group_topic" // machine group's topic
	// otherwise, file path regex.
	// eg, file path /var/log/nginx/access.log, TopicFormat: /var/log/([^/]+)/access\.log, so topic is 'nginx'
)

var NoConfigFieldError = errors.New("no this config field")
var InvalidTypeError = errors.New("invalid config type")

// IsValidInputType check if specific inputType is valid
func IsValidInputType(inputType string) bool {
	switch inputType {
	case InputTypeSyslog, InputTypeStreamlog, InputTypePlugin, InputTypeFile:
		return true
	}
	return false
}

// InputDetail defines log_config input
// @note : deprecated and no maintenance
type InputDetail struct {
	LogType       string   `json:"logType"`
	LogPath       string   `json:"logPath"`
	FilePattern   string   `json:"filePattern"`
	LocalStorage  bool     `json:"localStorage"`
	TimeKey       string   `json:"timeKey"`
	TimeFormat    string   `json:"timeFormat"`
	LogBeginRegex string   `json:"logBeginRegex"`
	Regex         string   `json:"regex"`
	Keys          []string `json:"key"`
	FilterKeys    []string `json:"filterKey"`
	FilterRegex   []string `json:"filterRegex"`
	TopicFormat   string   `json:"topicFormat"`
	Separator     string   `json:"separator"`
	AutoExtend    bool     `json:"autoExtend"`
}

func ConvertToInputDetail(detail InputDetailInterface) (*InputDetail, bool) {
	// ConvertToPluginLogConfigInputDetail need a plugin
	if mapVal, ok := detail.(map[string]interface{}); ok {
		if logType, ok := mapVal["logType"]; !ok || logType != "common_reg_log" {
			return nil, false
		}
	} else {
		return nil, false
	}
	buf, err := json.Marshal(detail)
	if err != nil {
		return nil, false
	}
	destDetail := &InputDetail{}
	err = json.Unmarshal(buf, destDetail)
	return destDetail, err == nil
}

type SensitiveKey struct {
	Key          string `json:"key"`
	Type         string `json:"type"`
	RegexBegin   string `json:"regex_begin"`
	RegexContent string `json:"regex_content"`
	All          bool   `json:"all"`
	ConstString  string `json:"const"`
}

// ApsaraLogConfigInputDetail apsara log config
type ApsaraLogConfigInputDetail struct {
	LocalFileConfigInputDetail
	LogBeginRegex string `json:"logBeginRegex"`
}

// InitApsaraLogConfigInputDetail ...
func InitApsaraLogConfigInputDetail(detail *ApsaraLogConfigInputDetail) {
	InitLocalFileConfigInputDetail(&detail.LocalFileConfigInputDetail)
	detail.LogBeginRegex = ".*"
	detail.LogType = LogFileTypeApsaraLog
}

func AddNecessaryApsaraLogInputConfigField(inputConfigDetail map[string]interface{}) {
	if _, ok := inputConfigDetail["logBeginRegex"]; !ok {
		inputConfigDetail["logBeginRegex"] = ".*"
	}
}

func ConvertToApsaraLogConfigInputDetail(detail InputDetailInterface) (*ApsaraLogConfigInputDetail, bool) {
	// ConvertToPluginLogConfigInputDetail need a plugin
	if mapVal, ok := detail.(map[string]interface{}); ok {
		if logType, ok := mapVal["logType"]; !ok || logType != "apsara_log" {
			return nil, false
		}
	} else {
		return nil, false
	}
	buf, err := json.Marshal(detail)
	if err != nil {
		return nil, false
	}
	destDetail := &ApsaraLogConfigInputDetail{}
	err = json.Unmarshal(buf, destDetail)
	return destDetail, err == nil
}

// RegexConfigInputDetail regex log config
type RegexConfigInputDetail struct {
	LocalFileConfigInputDetail
	Key           []string `json:"key"`
	LogBeginRegex string   `json:"logBeginRegex"`
	Regex         string   `json:"regex"`
}

// InitRegexConfigInputDetail ...
func InitRegexConfigInputDetail(detail *RegexConfigInputDetail) {
	InitLocalFileConfigInputDetail(&detail.LocalFileConfigInputDetail)
	detail.LogBeginRegex = ".*"
	detail.Regex = "(.*)"
	detail.LogType = LogFileTypeRegexLog
}

func AddNecessaryRegexLogInputConfigField(inputConfigDetail map[string]interface{}) {
	if _, ok := inputConfigDetail["logBeginRegex"]; !ok {
		inputConfigDetail["logBeginRegex"] = ".*"
	}

	if _, ok := inputConfigDetail["regex"]; !ok {
		inputConfigDetail["regex"] = "(.*)"
	}

	if _, ok := inputConfigDetail["key"]; !ok {
		inputConfigDetail["key"] = []string{"content"}
	}
}

func ConvertToRegexConfigInputDetail(detail InputDetailInterface) (*RegexConfigInputDetail, bool) {
	// ConvertToPluginLogConfigInputDetail need a plugin
	if mapVal, ok := detail.(map[string]interface{}); ok {
		if logType, ok := mapVal["logType"]; !ok || logType != "common_reg_log" {
			return nil, false
		}
	} else {
		return nil, false
	}
	buf, err := json.Marshal(detail)
	if err != nil {
		return nil, false
	}
	destDetail := &RegexConfigInputDetail{}
	err = json.Unmarshal(buf, destDetail)
	return destDetail, err == nil
}

// JSONConfigInputDetail pure json log config
type JSONConfigInputDetail struct {
	LocalFileConfigInputDetail
	TimeKey string `json:"timeKey"`
}

// InitJSONConfigInputDetail ...
func InitJSONConfigInputDetail(detail *JSONConfigInputDetail) {
	InitLocalFileConfigInputDetail(&detail.LocalFileConfigInputDetail)
	detail.LogType = LogFileTypeJSONLog
}

func AddNecessaryJSONLogInputConfigField(inputConfigDetail map[string]interface{}) {

}

func ConvertToJSONConfigInputDetail(detail InputDetailInterface) (*JSONConfigInputDetail, bool) {
	// ConvertToPluginLogConfigInputDetail need a plugin
	if mapVal, ok := detail.(map[string]interface{}); ok {
		if logType, ok := mapVal["logType"]; !ok || logType != "json_log" {
			return nil, false
		}
	} else {
		return nil, false
	}
	buf, err := json.Marshal(detail)
	if err != nil {
		return nil, false
	}
	destDetail := &JSONConfigInputDetail{}
	err = json.Unmarshal(buf, destDetail)
	return destDetail, err == nil
}

// DelimiterConfigInputDetail delimiter log config
type DelimiterConfigInputDetail struct {
	LocalFileConfigInputDetail
	Separator  string   `json:"separator"`
	Quote      string   `json:"quote"`
	Key        []string `json:"key"`
	TimeKey    string   `json:"timeKey,omitempty"`
	AutoExtend bool     `json:"autoExtend"`
}

// InitDelimiterConfigInputDetail ...
func InitDelimiterConfigInputDetail(detail *DelimiterConfigInputDetail) {
	InitLocalFileConfigInputDetail(&detail.LocalFileConfigInputDetail)
	detail.Quote = "\u0001"
	detail.AutoExtend = true
	detail.LogType = LogFileTypeDelimiterLog
}

func AddNecessaryDelimiterLogInputConfigField(inputConfigDetail map[string]interface{}) {
	if _, ok := inputConfigDetail["quote"]; !ok {
		inputConfigDetail["quote"] = "\u0001"
	}

	if _, ok := inputConfigDetail["autoExtend"]; !ok {
		inputConfigDetail["autoExtend"] = true
	}
}

func ConvertToDelimiterConfigInputDetail(detail InputDetailInterface) (*DelimiterConfigInputDetail, bool) {
	// ConvertToPluginLogConfigInputDetail need a plugin
	if mapVal, ok := detail.(map[string]interface{}); ok {
		if logType, ok := mapVal["logType"]; !ok || logType != "delimiter_log" {
			return nil, false
		}
	} else {
		return nil, false
	}
	buf, err := json.Marshal(detail)
	if err != nil {
		return nil, false
	}
	destDetail := &DelimiterConfigInputDetail{}
	err = json.Unmarshal(buf, destDetail)
	return destDetail, err == nil
}

// LocalFileConfigInputDetail all file input detail's basic config
type LocalFileConfigInputDetail struct {
	CommonConfigInputDetail
	LogType            string            `json:"logType"`
	LogPath            string            `json:"logPath"`
	FilePattern        string            `json:"filePattern"`
	TimeFormat         string            `json:"timeFormat"`
	TopicFormat        string            `json:"topicFormat,omitempty"`
	Preserve           bool              `json:"preserve"`
	PreserveDepth      int               `json:"preserveDepth"`
	FileEncoding       string            `json:"fileEncoding,omitempty"`
	DiscardUnmatch     bool              `json:"discardUnmatch"`
	MaxDepth           int               `json:"maxDepth"`
	TailExisted        bool              `json:"tailExisted"`
	DiscardNonUtf8     bool              `json:"discardNonUtf8"`
	DelaySkipBytes     int               `json:"delaySkipBytes"`
	IsDockerFile       bool              `json:"dockerFile"`
	DockerIncludeLabel map[string]string `json:"dockerIncludeLabel,omitempty"`
	DockerExcludeLabel map[string]string `json:"dockerExcludeLabel,omitempty"`
	DockerIncludeEnv   map[string]string `json:"dockerIncludeEnv,omitempty"`
	DockerExcludeEnv   map[string]string `json:"dockerExcludeEnv,omitempty"`
}

func GetFileConfigInputDetailType(detail InputDetailInterface) (string, bool) {
	// ConvertToPluginLogConfigInputDetail need a plugin
	if mapVal, ok := detail.(map[string]interface{}); ok {
		if logType, ok := mapVal["logType"]; ok {
			return logType.(string), true
		}
	}
	return "", false
}

// InitLocalFileConfigInputDetail ...
func InitLocalFileConfigInputDetail(detail *LocalFileConfigInputDetail) {
	InitCommonConfigInputDetail(&detail.CommonConfigInputDetail)
	detail.FileEncoding = "utf8"
	detail.MaxDepth = 100
	detail.TopicFormat = TopicFormatNone
	detail.Preserve = true
	detail.DiscardUnmatch = true
}

func AddNecessaryLocalFileInputConfigField(inputConfigDetail map[string]interface{}) {
	if _, ok := inputConfigDetail["fileEncoding"]; !ok {
		inputConfigDetail["fileEncoding"] = "utf8"
	}

	if _, ok := inputConfigDetail["maxDepth"]; !ok {
		inputConfigDetail["maxDepth"] = 100
	}

	if _, ok := inputConfigDetail["topicFormat"]; !ok {
		inputConfigDetail["topicFormat"] = TopicFormatNone
	}

	if _, ok := inputConfigDetail["preserve"]; !ok {
		inputConfigDetail["preserve"] = true
	}

	if _, ok := inputConfigDetail["discardUnmatch"]; !ok {
		inputConfigDetail["discardUnmatch"] = true
	}

	if _, ok := inputConfigDetail["timeFormat"]; !ok {
		inputConfigDetail["timeFormat"] = ""
	}
}

// PluginLogConfigInputDetail plugin log config, eg: docker stdout, binlog, mysql, http...
type PluginLogConfigInputDetail struct {
	CommonConfigInputDetail
	PluginDetail LogConfigPluginInput `json:"plugin"`
}

// InitPluginLogConfigInputDetail ...
func InitPluginLogConfigInputDetail(detail *PluginLogConfigInputDetail) {
	InitCommonConfigInputDetail(&detail.CommonConfigInputDetail)
}

func ConvertToPluginLogConfigInputDetail(detail InputDetailInterface) (*PluginLogConfigInputDetail, bool) {
	// ConvertToPluginLogConfigInputDetail need a plugin
	if mapVal, ok := detail.(map[string]interface{}); ok {
		if _, ok := mapVal["plugin"]; !ok {
			return nil, false
		}
		if _, ok := mapVal["logType"]; ok {
			return nil, false
		}
		buf, err := json.Marshal(detail)
		if err != nil {
			return nil, false
		}
		destDetail := &PluginLogConfigInputDetail{}
		err = json.Unmarshal(buf, destDetail)
		return destDetail, err == nil
	}
	return nil, false
}

// StreamLogConfigInputDetail syslog config
type StreamLogConfigInputDetail struct {
	CommonConfigInputDetail
	Tag string `json:"tag"`
}

// InitStreamLogConfigInputDetail ...
func InitStreamLogConfigInputDetail(detail *StreamLogConfigInputDetail) {
	InitCommonConfigInputDetail(&detail.CommonConfigInputDetail)
}

func ConvertToStreamLogConfigInputDetail(detail InputDetailInterface) (*StreamLogConfigInputDetail, bool) {
	// ConvertToStreamLogConfigInputDetail need a tag
	if mapVal, ok := detail.(map[string]interface{}); ok {
		if _, ok := mapVal["tag"]; !ok {
			return nil, false
		}
	} else {
		return nil, false
	}
	buf, err := json.Marshal(detail)
	if err != nil {
		return nil, false
	}
	destDetail := &StreamLogConfigInputDetail{}
	err = json.Unmarshal(buf, destDetail)
	return destDetail, err == nil
}

// CommonConfigInputDetail is all input detail's basic config
type CommonConfigInputDetail struct {
	LocalStorage    bool           `json:"localStorage"`
	FilterKeys      []string       `json:"filterKey,omitempty"`
	FilterRegex     []string       `json:"filterRegex,omitempty"`
	ShardHashKey    []string       `json:"shardHashKey,omitempty"`
	EnableTag       bool           `json:"enableTag"`
	EnableRawLog    bool           `json:"enableRawLog"`
	MaxSendRate     int            `json:"maxSendRate"`
	SendRateExpire  int            `json:"sendRateExpire"`
	SensitiveKeys   []SensitiveKey `json:"sensitive_keys,omitempty"`
	MergeType       string         `json:"mergeType,omitempty"`
	DelayAlarmBytes int            `json:"delayAlarmBytes,omitempty"`
	AdjustTimeZone  bool           `json:"adjustTimezone"`
	LogTimeZone     string         `json:"logTimezone,omitempty"`
	Priority        int            `json:"priority,omitempty"`
}

// InitCommonConfigInputDetail ...
func InitCommonConfigInputDetail(detail *CommonConfigInputDetail) {
	detail.LocalStorage = true
	detail.EnableTag = true
	detail.MaxSendRate = -1
	detail.MergeType = MergeTypeTopic
}

// AddNecessaryInputConfigField ...
func AddNecessaryInputConfigField(inputConfigDetail map[string]interface{}) {
	if _, ok := inputConfigDetail["localStorage"]; !ok {
		inputConfigDetail["localStorage"] = true
	}
	if _, ok := inputConfigDetail["enableTag"]; !ok {
		inputConfigDetail["enableTag"] = true
	}
	if _, ok := inputConfigDetail["maxSendRate"]; !ok {
		inputConfigDetail["maxSendRate"] = -1
	}
	if _, ok := inputConfigDetail["mergeType"]; !ok {
		inputConfigDetail["mergeType"] = MergeTypeTopic
	}

	if logTypeInterface, ok := inputConfigDetail["logType"]; ok {
		if logType, ok := logTypeInterface.(string); ok {
			AddNecessaryLocalFileInputConfigField(inputConfigDetail)
			switch logType {
			case LogFileTypeApsaraLog:
				AddNecessaryApsaraLogInputConfigField(inputConfigDetail)
			case LogFileTypeRegexLog:
				AddNecessaryRegexLogInputConfigField(inputConfigDetail)
			case LogFileTypeJSONLog:
				AddNecessaryJSONLogInputConfigField(inputConfigDetail)
			case LogFileTypeDelimiterLog:
				AddNecessaryDelimiterLogInputConfigField(inputConfigDetail)
			}
		}
	}
}

// UpdateInputConfigField ...
func UpdateInputConfigField(detail InputDetailInterface, key string, val interface{}) error {
	if mapVal, ok := detail.(map[string]interface{}); ok {
		if _, ok := mapVal[key]; !ok {
			return NoConfigFieldError
		}
		mapVal[key] = val
		return nil
	}
	return InvalidTypeError
}

// OutputDetail defines output
type OutputDetail struct {
	ProjectName  string `json:"projectName"`
	LogStoreName string `json:"logstoreName"`
}

// InputDetailInterface all input detail's interface
type InputDetailInterface interface {
}

// LogConfig defines log config
type LogConfig struct {
	Name         string               `json:"configName"`
	LogSample    string               `json:"logSample"`
	InputType    string               `json:"inputType"` // syslog plugin file
	InputDetail  InputDetailInterface `json:"inputDetail"`
	OutputType   string               `json:"outputType"`
	OutputDetail OutputDetail         `json:"outputDetail"`

	CreateTime     uint32 `json:"createTime,omitempty`
	LastModifyTime uint32 `json:"lastModifyTime,omitempty"`
}
