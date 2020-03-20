package octopusdeploy

type ActionTemplateParameter struct {

	// default value
	DefaultValue PropertyValueResource `json:"DefaultValue,omitempty"`

	// display settings
	DisplaySettings map[string]string `json:"DisplaySettings,omitempty"`

	// help text
	HelpText string `json:"HelpText,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// label
	Label string `json:"Label,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn string `json:"LastModifiedOn,omitempty"` // datetime

	// links
	Links Links `json:"Links,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}
