package fc

// LogTriggerConfig ..
type LogTriggerConfig struct {
	SourceConfig      *SourceConfig          `json:"sourceConfig"`
	JobConfig         *JobConfig             `json:"jobConfig"`
	FunctionParameter map[string]interface{} `json:"functionParameter"`
	LogConfig         *JobLogConfig          `json:"logConfig"`
	Enable            *bool                  `json:"enable"`
}

func NewLogTriggerConfig() *LogTriggerConfig {
	return &LogTriggerConfig{}
}

func (ltc *LogTriggerConfig) WithSourceConfig(c *SourceConfig) *LogTriggerConfig {
	ltc.SourceConfig = c
	return ltc
}

func (ltc *LogTriggerConfig) WithJobConfig(c *JobConfig) *LogTriggerConfig {
	ltc.JobConfig = c
	return ltc
}

func (ltc *LogTriggerConfig) WithFunctionParameter(p map[string]interface{}) *LogTriggerConfig {
	ltc.FunctionParameter = p
	return ltc
}

func (ltc *LogTriggerConfig) WithLogConfig(c *JobLogConfig) *LogTriggerConfig {
	ltc.LogConfig = c
	return ltc
}

func (ltc *LogTriggerConfig) WithEnable(enable bool) *LogTriggerConfig {
	ltc.Enable = &enable
	return ltc
}

// SourceConfig ..
type SourceConfig struct {
	Logstore *string `json:"logstore"`
}

func NewSourceConfig() *SourceConfig {
	return &SourceConfig{}
}

func (sc *SourceConfig) WithLogstore(store string) *SourceConfig {
	sc.Logstore = &store
	return sc
}

// JobConfig maps to Log service's trigger config.
type JobConfig struct {
	MaxRetryTime    *int `json:"maxRetryTime"`
	TriggerInterval *int `json:"triggerInterval"`
}

func NewJobConfig() *JobConfig {
	return &JobConfig{}
}

func (jc *JobConfig) WithMaxRetryTime(retry int) *JobConfig {
	jc.MaxRetryTime = &retry
	return jc
}

func (jc *JobConfig) WithTriggerInterval(interval int) *JobConfig {
	jc.TriggerInterval = &interval
	return jc
}

// LogConfig ..
type JobLogConfig struct {
	Project  *string `json:"project"`
	Logstore *string `json:"logstore"`
}

func NewJobLogConfig() *JobLogConfig {
	return &JobLogConfig{}
}

func (jlc *JobLogConfig) WithProject(p string) *JobLogConfig {
	jlc.Project = &p
	return jlc
}

func (jlc *JobLogConfig) WithLogstore(s string) *JobLogConfig {
	jlc.Logstore = &s
	return jlc
}
