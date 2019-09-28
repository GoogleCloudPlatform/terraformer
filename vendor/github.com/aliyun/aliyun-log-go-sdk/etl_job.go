package sls

import "encoding/json"

type ETLJob struct {
	JobName        string          `json:"etlJobName"`
	SourceConfig   *SourceConfig   `json:"sourceConfig"`
	TriggerConfig  *TriggerConfig  `json:"triggerConfig"`
	FunctionConfig *FunctionConfig `json:"functionConfig"`

	// TODO: change this to map[string]interface{} once log service fixes the format
	FunctionParameter interface{}   `json:"functionParameter"`
	LogConfig         *JobLogConfig `json:"logConfig"`
	Enable            bool          `json:"enable"`

	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
}

type etlJobAlias ETLJob

// This can be removed once log service returns function parameter in json type.
// "functionParameter":{"a":1} instead of "functionParameter":"{\"a\":1}"
func (job *ETLJob) UnmarshalJSON(data []byte) error {
	output := etlJobAlias{}
	if err := json.Unmarshal(data, &output); err != nil {
		return err
	}
	param := map[string]interface{}{}
	paramStr, ok := output.FunctionParameter.(string)
	if ok {
		if err := json.Unmarshal([]byte(paramStr), &param); err != nil {
			return err
		}
		output.FunctionParameter = param
	}
	*job = ETLJob(output)
	return nil
}

type SourceConfig struct {
	LogstoreName string `json:"logstoreName"`
}

type TriggerConfig struct {
	MaxRetryTime    int    `json:"maxRetryTime"`
	TriggerInterval int    `json:"triggerInterval"`
	RoleARN         string `json:"roleArn"`
}

type FunctionConfig struct {
	FunctionProvider string `json:"functionProvider"`
	Endpoint         string `json:"endpoint"`
	AccountID        string `json:"accountId"`
	RegionName       string `json:"regionName"`
	ServiceName      string `json:"serviceName"`
	FunctionName     string `json:"functionName"`
}

type JobLogConfig struct {
	Endpoint     string `json:"endpoint"`
	ProjectName  string `json:"projectName"`
	LogstoreName string `json:"logstoreName"`
}
