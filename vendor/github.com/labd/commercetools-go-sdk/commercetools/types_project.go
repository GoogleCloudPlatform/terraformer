// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ProjectUpdateAction uses action as discriminator attribute
type ProjectUpdateAction interface{}

func mapDiscriminatorProjectUpdateAction(input interface{}) (ProjectUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "changeCountries":
		new := ProjectChangeCountriesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeCurrencies":
		new := ProjectChangeCurrenciesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLanguages":
		new := ProjectChangeLanguagesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeMessagesConfiguration":
		new := ProjectChangeMessagesConfigurationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeMessagesEnabled":
		new := ProjectChangeMessagesEnabledAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := ProjectChangeNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setExternalOAuth":
		new := ProjectSetExternalOAuthAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingRateInputType":
		new := ProjectSetShippingRateInputTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.ShippingRateInputType != nil {
			new.ShippingRateInputType, err = mapDiscriminatorShippingRateInputType(new.ShippingRateInputType)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	}
	return nil, nil
}

// ShippingRateInputType uses type as discriminator attribute
type ShippingRateInputType interface{}

func mapDiscriminatorShippingRateInputType(input interface{}) (ShippingRateInputType, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "CartClassification":
		new := CartClassificationType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CartScore":
		new := CartScoreType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CartValue":
		new := CartValueType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// CartClassificationType implements the interface ShippingRateInputType
type CartClassificationType struct {
	Values []CustomFieldLocalizedEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value
func (obj CartClassificationType) MarshalJSON() ([]byte, error) {
	type Alias CartClassificationType
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartClassification", Alias: (*Alias)(&obj)})
}

// CartScoreType implements the interface ShippingRateInputType
type CartScoreType struct{}

// MarshalJSON override to set the discriminator value
func (obj CartScoreType) MarshalJSON() ([]byte, error) {
	type Alias CartScoreType
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartScore", Alias: (*Alias)(&obj)})
}

// CartValueType implements the interface ShippingRateInputType
type CartValueType struct{}

// MarshalJSON override to set the discriminator value
func (obj CartValueType) MarshalJSON() ([]byte, error) {
	type Alias CartValueType
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartValue", Alias: (*Alias)(&obj)})
}

// ExternalOAuth is a standalone struct
type ExternalOAuth struct {
	URL                 string `json:"url"`
	AuthorizationHeader string `json:"authorizationHeader"`
}

// Project is a standalone struct
type Project struct {
	Version               int                   `json:"version"`
	TrialUntil            string                `json:"trialUntil,omitempty"`
	ShippingRateInputType ShippingRateInputType `json:"shippingRateInputType,omitempty"`
	Name                  string                `json:"name"`
	Messages              *MessageConfiguration `json:"messages"`
	Languages             []Locale              `json:"languages"`
	Key                   string                `json:"key"`
	ExternalOAuth         *ExternalOAuth        `json:"externalOAuth,omitempty"`
	Currencies            []CurrencyCode        `json:"currencies"`
	CreatedAt             time.Time             `json:"createdAt"`
	Countries             []CountryCode         `json:"countries"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Project) UnmarshalJSON(data []byte) error {
	type Alias Project
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInputType != nil {
		var err error
		obj.ShippingRateInputType, err = mapDiscriminatorShippingRateInputType(obj.ShippingRateInputType)
		if err != nil {
			return err
		}
	}

	return nil
}

// ProjectChangeCountriesAction implements the interface ProjectUpdateAction
type ProjectChangeCountriesAction struct {
	Countries []CountryCode `json:"countries"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeCountriesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeCountriesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCountries", Alias: (*Alias)(&obj)})
}

// ProjectChangeCurrenciesAction implements the interface ProjectUpdateAction
type ProjectChangeCurrenciesAction struct {
	Currencies []CurrencyCode `json:"currencies"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeCurrenciesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeCurrenciesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCurrencies", Alias: (*Alias)(&obj)})
}

// ProjectChangeLanguagesAction implements the interface ProjectUpdateAction
type ProjectChangeLanguagesAction struct {
	Languages []Locale `json:"languages"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeLanguagesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeLanguagesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLanguages", Alias: (*Alias)(&obj)})
}

// ProjectChangeMessagesConfigurationAction implements the interface ProjectUpdateAction
type ProjectChangeMessagesConfigurationAction struct {
	MessagesConfiguration *MessageConfigurationDraft `json:"messagesConfiguration"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeMessagesConfigurationAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeMessagesConfigurationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMessagesConfiguration", Alias: (*Alias)(&obj)})
}

// ProjectChangeMessagesEnabledAction implements the interface ProjectUpdateAction
type ProjectChangeMessagesEnabledAction struct {
	MessagesEnabled bool `json:"messagesEnabled"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeMessagesEnabledAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeMessagesEnabledAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMessagesEnabled", Alias: (*Alias)(&obj)})
}

// ProjectChangeNameAction implements the interface ProjectUpdateAction
type ProjectChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ProjectSetExternalOAuthAction implements the interface ProjectUpdateAction
type ProjectSetExternalOAuthAction struct {
	ExternalOAuth *ExternalOAuth `json:"externalOAuth,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectSetExternalOAuthAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectSetExternalOAuthAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalOAuth", Alias: (*Alias)(&obj)})
}

// ProjectSetShippingRateInputTypeAction implements the interface ProjectUpdateAction
type ProjectSetShippingRateInputTypeAction struct {
	ShippingRateInputType ShippingRateInputType `json:"shippingRateInputType,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectSetShippingRateInputTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectSetShippingRateInputTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingRateInputType", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProjectSetShippingRateInputTypeAction) UnmarshalJSON(data []byte) error {
	type Alias ProjectSetShippingRateInputTypeAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInputType != nil {
		var err error
		obj.ShippingRateInputType, err = mapDiscriminatorShippingRateInputType(obj.ShippingRateInputType)
		if err != nil {
			return err
		}
	}

	return nil
}

// ProjectUpdate is a standalone struct
type ProjectUpdate struct {
	Version int                   `json:"version"`
	Actions []ProjectUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProjectUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProjectUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorProjectUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
