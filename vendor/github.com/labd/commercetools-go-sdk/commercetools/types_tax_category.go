// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// TaxCategoryUpdateAction uses action as discriminator attribute
type TaxCategoryUpdateAction interface{}

func mapDiscriminatorTaxCategoryUpdateAction(input interface{}) (TaxCategoryUpdateAction, error) {
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
	case "addTaxRate":
		new := TaxCategoryAddTaxRateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := TaxCategoryChangeNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeTaxRate":
		new := TaxCategoryRemoveTaxRateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "replaceTaxRate":
		new := TaxCategoryReplaceTaxRateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := TaxCategorySetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := TaxCategorySetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// SubRate is a standalone struct
type SubRate struct {
	Name   string   `json:"name"`
	Amount *float64 `json:"amount"`
}

// TaxCategory is of type LoggedResource
type TaxCategory struct {
	Version        int             `json:"version"`
	LastModifiedAt time.Time       `json:"lastModifiedAt"`
	ID             string          `json:"id"`
	CreatedAt      time.Time       `json:"createdAt"`
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	CreatedBy      *CreatedBy      `json:"createdBy,omitempty"`
	Rates          []TaxRate       `json:"rates"`
	Name           string          `json:"name"`
	Key            string          `json:"key,omitempty"`
	Description    string          `json:"description,omitempty"`
}

// TaxCategoryAddTaxRateAction implements the interface TaxCategoryUpdateAction
type TaxCategoryAddTaxRateAction struct {
	TaxRate *TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryAddTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryAddTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTaxRate", Alias: (*Alias)(&obj)})
}

// TaxCategoryChangeNameAction implements the interface TaxCategoryUpdateAction
type TaxCategoryChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// TaxCategoryDraft is a standalone struct
type TaxCategoryDraft struct {
	Rates       []TaxRateDraft `json:"rates"`
	Name        string         `json:"name"`
	Key         string         `json:"key,omitempty"`
	Description string         `json:"description,omitempty"`
}

// TaxCategoryPagedQueryResponse is a standalone struct
type TaxCategoryPagedQueryResponse struct {
	Total   int           `json:"total,omitempty"`
	Results []TaxCategory `json:"results"`
	Offset  int           `json:"offset"`
	Count   int           `json:"count"`
}

// TaxCategoryReference implements the interface Reference
type TaxCategoryReference struct {
	ID  string       `json:"id"`
	Obj *TaxCategory `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryReference) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "tax-category", Alias: (*Alias)(&obj)})
}

// TaxCategoryRemoveTaxRateAction implements the interface TaxCategoryUpdateAction
type TaxCategoryRemoveTaxRateAction struct {
	TaxRateID string `json:"taxRateId"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryRemoveTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryRemoveTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTaxRate", Alias: (*Alias)(&obj)})
}

// TaxCategoryReplaceTaxRateAction implements the interface TaxCategoryUpdateAction
type TaxCategoryReplaceTaxRateAction struct {
	TaxRateID string        `json:"taxRateId"`
	TaxRate   *TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryReplaceTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReplaceTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "replaceTaxRate", Alias: (*Alias)(&obj)})
}

// TaxCategoryResourceIdentifier implements the interface ResourceIdentifier
type TaxCategoryResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "tax-category", Alias: (*Alias)(&obj)})
}

// TaxCategorySetDescriptionAction implements the interface TaxCategoryUpdateAction
type TaxCategorySetDescriptionAction struct {
	Description string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategorySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// TaxCategorySetKeyAction implements the interface TaxCategoryUpdateAction
type TaxCategorySetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategorySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// TaxCategoryUpdate is a standalone struct
type TaxCategoryUpdate struct {
	Version int                       `json:"version"`
	Actions []TaxCategoryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TaxCategoryUpdate) UnmarshalJSON(data []byte) error {
	type Alias TaxCategoryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorTaxCategoryUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// TaxRate is a standalone struct
type TaxRate struct {
	SubRates        []SubRate   `json:"subRates,omitempty"`
	State           string      `json:"state,omitempty"`
	Name            string      `json:"name"`
	IncludedInPrice bool        `json:"includedInPrice"`
	ID              string      `json:"id,omitempty"`
	Country         CountryCode `json:"country"`
	Amount          *float64    `json:"amount"`
}

// TaxRateDraft is a standalone struct
type TaxRateDraft struct {
	SubRates        []SubRate   `json:"subRates,omitempty"`
	State           string      `json:"state,omitempty"`
	Name            string      `json:"name"`
	IncludedInPrice bool        `json:"includedInPrice"`
	Country         CountryCode `json:"country"`
	Amount          *float64    `json:"amount,omitempty"`
}
