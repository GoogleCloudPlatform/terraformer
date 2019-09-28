package fc

// CDNEventsTriggerConfig defines the cdn events trigger config
type CDNEventsTriggerConfig struct {
	EventName    *string 			`json:"eventName"`
	EventVersion *string 			`json:"eventVersion"`
	Notes        *string 			`json:"notes"`
	Filter       map[string][]string 	`json:"filter"`
}

// NewCDNEventsTriggerConfig creates an empty CDNTEventsTriggerConfig
func NewCDNEventsTriggerConfig() *CDNEventsTriggerConfig {
	return &CDNEventsTriggerConfig{}
}

func (ctc *CDNEventsTriggerConfig) WithEventName(eventName string) *CDNEventsTriggerConfig {
	ctc.EventName = &eventName
	return ctc
}

func (ctc *CDNEventsTriggerConfig) WithEventVersion(eventVersion string) *CDNEventsTriggerConfig {
	ctc.EventVersion = &eventVersion
	return ctc
}

func (ctc *CDNEventsTriggerConfig) WithNotes(notes string) *CDNEventsTriggerConfig {
	ctc.Notes = &notes
	return ctc
}

func (ctc *CDNEventsTriggerConfig) WithFilter(filter map[string][]string) *CDNEventsTriggerConfig {
	ctc.Filter = make(map[string][]string, len(filter))
	for k,v := range filter {
		ctc.Filter[k] = v
	}
	return ctc
}
