package fc

// TimeTriggerConfig defines the time trigger config
type TimeTriggerConfig struct {
	Payload        *string `json:"payload"`
	CronExpression *string `json:"cronExpression"`
	Enable         *bool   `json:"enable"`
}

// NewTimeTriggerConfig creates an empty TimeTriggerConfig
func NewTimeTriggerConfig() *TimeTriggerConfig {
	return &TimeTriggerConfig{}
}

func (ttc *TimeTriggerConfig) WithPayload(payload string) *TimeTriggerConfig {
	ttc.Payload = &payload
	return ttc
}

func (ttc *TimeTriggerConfig) WithCronExpression(cronExpression string) *TimeTriggerConfig {
	ttc.CronExpression = &cronExpression
	return ttc
}

func (ttc *TimeTriggerConfig) WithEnable(enable bool) *TimeTriggerConfig {
	ttc.Enable = &enable
	return ttc
}
