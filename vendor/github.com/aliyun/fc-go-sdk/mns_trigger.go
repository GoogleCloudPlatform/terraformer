package fc

// MnsTopicTriggerConfig ..
type MnsTopicTriggerConfig struct {
	FilterTag           *string `json:"filterTag"`
	NotifyContentFormat *string `json:"notifyContentFormat"`
	NotifyStrategy      *string `json:"notifyStrategy"`
}

// NewMnsTopicTriggerConfig ..
func NewMnsTopicTriggerConfig() *MnsTopicTriggerConfig {
	return &MnsTopicTriggerConfig{}
}

func (mtc *MnsTopicTriggerConfig) WithFilterTag(filterTag string) *MnsTopicTriggerConfig {
	mtc.FilterTag = &filterTag
	return mtc
}

func (mtc *MnsTopicTriggerConfig) WithNotifyContentFormat(notifyContentFormat string) *MnsTopicTriggerConfig {
	mtc.NotifyContentFormat = &notifyContentFormat
	return mtc
}

func (mtc *MnsTopicTriggerConfig) WithNotifyStrategy(notifyStrategy string) *MnsTopicTriggerConfig {
	mtc.NotifyStrategy = &notifyStrategy
	return mtc
}
