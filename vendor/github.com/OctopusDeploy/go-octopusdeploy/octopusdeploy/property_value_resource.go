package octopusdeploy

import (
	"encoding/json"
)

type SensitiveValue struct {
	HasValue bool   `json:"HasValue"`
	NewValue string `json:"NewValue"`
}

type PropertyValue string

type PropertyValueResource struct {
	*SensitiveValue
	*PropertyValue
}

func (d PropertyValueResource) MarshalJSON() ([]byte, error) {
	if d.SensitiveValue != nil {
		return json.Marshal(d.SensitiveValue)
	}

	if d.PropertyValue != nil {
		return json.Marshal(d.PropertyValue)
	}

	return json.Marshal(``)
}

func (d *PropertyValueResource) UnmarshalJSON(data []byte) error {
	// try unmarshal into a sensitive property, if that fails, it's just a normal property
	var spv SensitiveValue
	errUnmarshalSensitivePropertyValue := json.Unmarshal(data, &spv)

	if errUnmarshalSensitivePropertyValue != nil {
		var p PropertyValue
		errUnmarshalString := json.Unmarshal(data, &p)

		if errUnmarshalString != nil {
			return errUnmarshalString
		}

		d.PropertyValue = &p
		d.SensitiveValue = nil
		return nil
	}

	d.SensitiveValue = &spv
	return nil
}
