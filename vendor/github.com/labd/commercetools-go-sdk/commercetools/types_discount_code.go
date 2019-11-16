// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// DiscountCodeUpdateAction uses action as discriminator attribute
type DiscountCodeUpdateAction interface{}

func mapDiscriminatorDiscountCodeUpdateAction(input interface{}) (DiscountCodeUpdateAction, error) {
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
	case "changeCartDiscounts":
		new := DiscountCodeChangeCartDiscountsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeGroups":
		new := DiscountCodeChangeGroupsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeIsActive":
		new := DiscountCodeChangeIsActiveAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCartPredicate":
		new := DiscountCodeSetCartPredicateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := DiscountCodeSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := DiscountCodeSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := DiscountCodeSetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMaxApplications":
		new := DiscountCodeSetMaxApplicationsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMaxApplicationsPerCustomer":
		new := DiscountCodeSetMaxApplicationsPerCustomerAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setName":
		new := DiscountCodeSetNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidFrom":
		new := DiscountCodeSetValidFromAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidFromAndUntil":
		new := DiscountCodeSetValidFromAndUntilAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidUntil":
		new := DiscountCodeSetValidUntilAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// DiscountCode is of type LoggedResource
type DiscountCode struct {
	Version                    int                     `json:"version"`
	LastModifiedAt             time.Time               `json:"lastModifiedAt"`
	ID                         string                  `json:"id"`
	CreatedAt                  time.Time               `json:"createdAt"`
	LastModifiedBy             *LastModifiedBy         `json:"lastModifiedBy,omitempty"`
	CreatedBy                  *CreatedBy              `json:"createdBy,omitempty"`
	ValidUntil                 *time.Time              `json:"validUntil,omitempty"`
	ValidFrom                  *time.Time              `json:"validFrom,omitempty"`
	References                 []Reference             `json:"references"`
	Name                       *LocalizedString        `json:"name,omitempty"`
	MaxApplicationsPerCustomer int                     `json:"maxApplicationsPerCustomer,omitempty"`
	MaxApplications            int                     `json:"maxApplications,omitempty"`
	IsActive                   bool                    `json:"isActive"`
	Groups                     []string                `json:"groups"`
	Description                *LocalizedString        `json:"description,omitempty"`
	Custom                     *CustomFields           `json:"custom,omitempty"`
	Code                       string                  `json:"code"`
	CartPredicate              string                  `json:"cartPredicate,omitempty"`
	CartDiscounts              []CartDiscountReference `json:"cartDiscounts"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountCode) UnmarshalJSON(data []byte) error {
	type Alias DiscountCode
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.References {
		var err error
		obj.References[i], err = mapDiscriminatorReference(obj.References[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// DiscountCodeChangeCartDiscountsAction implements the interface DiscountCodeUpdateAction
type DiscountCodeChangeCartDiscountsAction struct {
	CartDiscounts []CartDiscountResourceIdentifier `json:"cartDiscounts"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeChangeCartDiscountsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeCartDiscountsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCartDiscounts", Alias: (*Alias)(&obj)})
}

// DiscountCodeChangeGroupsAction implements the interface DiscountCodeUpdateAction
type DiscountCodeChangeGroupsAction struct {
	Groups []string `json:"groups"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeChangeGroupsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeGroupsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeGroups", Alias: (*Alias)(&obj)})
}

// DiscountCodeChangeIsActiveAction implements the interface DiscountCodeUpdateAction
type DiscountCodeChangeIsActiveAction struct {
	IsActive bool `json:"isActive"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

// DiscountCodeDraft is a standalone struct
type DiscountCodeDraft struct {
	ValidUntil                 *time.Time                       `json:"validUntil,omitempty"`
	ValidFrom                  *time.Time                       `json:"validFrom,omitempty"`
	Name                       *LocalizedString                 `json:"name,omitempty"`
	MaxApplicationsPerCustomer int                              `json:"maxApplicationsPerCustomer,omitempty"`
	MaxApplications            int                              `json:"maxApplications,omitempty"`
	IsActive                   bool                             `json:"isActive"`
	Groups                     []string                         `json:"groups,omitempty"`
	Description                *LocalizedString                 `json:"description,omitempty"`
	Custom                     *CustomFieldsDraft               `json:"custom,omitempty"`
	Code                       string                           `json:"code"`
	CartPredicate              string                           `json:"cartPredicate,omitempty"`
	CartDiscounts              []CartDiscountResourceIdentifier `json:"cartDiscounts"`
}

// DiscountCodePagedQueryResponse is a standalone struct
type DiscountCodePagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Results []DiscountCode `json:"results"`
	Offset  int            `json:"offset"`
	Count   int            `json:"count"`
}

// DiscountCodeReference implements the interface Reference
type DiscountCodeReference struct {
	ID  string        `json:"id"`
	Obj *DiscountCode `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeReference) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "discount-code", Alias: (*Alias)(&obj)})
}

// DiscountCodeResourceIdentifier implements the interface ResourceIdentifier
type DiscountCodeResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "discount-code", Alias: (*Alias)(&obj)})
}

// DiscountCodeSetCartPredicateAction implements the interface DiscountCodeUpdateAction
type DiscountCodeSetCartPredicateAction struct {
	CartPredicate string `json:"cartPredicate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetCartPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCartPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCartPredicate", Alias: (*Alias)(&obj)})
}

// DiscountCodeSetCustomFieldAction implements the interface DiscountCodeUpdateAction
type DiscountCodeSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// DiscountCodeSetCustomTypeAction implements the interface DiscountCodeUpdateAction
type DiscountCodeSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// DiscountCodeSetDescriptionAction implements the interface DiscountCodeUpdateAction
type DiscountCodeSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// DiscountCodeSetMaxApplicationsAction implements the interface DiscountCodeUpdateAction
type DiscountCodeSetMaxApplicationsAction struct {
	MaxApplications int `json:"maxApplications,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetMaxApplicationsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetMaxApplicationsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMaxApplications", Alias: (*Alias)(&obj)})
}

// DiscountCodeSetMaxApplicationsPerCustomerAction implements the interface DiscountCodeUpdateAction
type DiscountCodeSetMaxApplicationsPerCustomerAction struct {
	MaxApplicationsPerCustomer int `json:"maxApplicationsPerCustomer,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetMaxApplicationsPerCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetMaxApplicationsPerCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMaxApplicationsPerCustomer", Alias: (*Alias)(&obj)})
}

// DiscountCodeSetNameAction implements the interface DiscountCodeUpdateAction
type DiscountCodeSetNameAction struct {
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

// DiscountCodeSetValidFromAction implements the interface DiscountCodeUpdateAction
type DiscountCodeSetValidFromAction struct {
	ValidFrom *time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

// DiscountCodeSetValidFromAndUntilAction implements the interface DiscountCodeUpdateAction
type DiscountCodeSetValidFromAndUntilAction struct {
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	ValidFrom  *time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

// DiscountCodeSetValidUntilAction implements the interface DiscountCodeUpdateAction
type DiscountCodeSetValidUntilAction struct {
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}

// DiscountCodeUpdate is a standalone struct
type DiscountCodeUpdate struct {
	Version int                        `json:"version"`
	Actions []DiscountCodeUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountCodeUpdate) UnmarshalJSON(data []byte) error {
	type Alias DiscountCodeUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorDiscountCodeUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
