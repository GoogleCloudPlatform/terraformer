// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// StateRoleEnum is an enum type
type StateRoleEnum string

// Enum values for StateRoleEnum
const (
	StateRoleEnumReviewIncludedInStatistics StateRoleEnum = "ReviewIncludedInStatistics"
	StateRoleEnumReturn                     StateRoleEnum = "Return"
)

// StateTypeEnum is an enum type
type StateTypeEnum string

// Enum values for StateTypeEnum
const (
	StateTypeEnumOrderState    StateTypeEnum = "OrderState"
	StateTypeEnumLineItemState StateTypeEnum = "LineItemState"
	StateTypeEnumProductState  StateTypeEnum = "ProductState"
	StateTypeEnumReviewState   StateTypeEnum = "ReviewState"
	StateTypeEnumPaymentState  StateTypeEnum = "PaymentState"
)

// StateUpdateAction uses action as discriminator attribute
type StateUpdateAction interface{}

func mapDiscriminatorStateUpdateAction(input interface{}) (StateUpdateAction, error) {
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
	case "addRoles":
		new := StateAddRolesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeInitial":
		new := StateChangeInitialAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeKey":
		new := StateChangeKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeType":
		new := StateChangeTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeRoles":
		new := StateRemoveRolesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := StateSetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setName":
		new := StateSetNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setRoles":
		new := StateSetRolesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTransitions":
		new := StateSetTransitionsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// State is of type LoggedResource
type State struct {
	Version        int              `json:"version"`
	LastModifiedAt time.Time        `json:"lastModifiedAt"`
	ID             string           `json:"id"`
	CreatedAt      time.Time        `json:"createdAt"`
	LastModifiedBy *LastModifiedBy  `json:"lastModifiedBy,omitempty"`
	CreatedBy      *CreatedBy       `json:"createdBy,omitempty"`
	Type           StateTypeEnum    `json:"type"`
	Transitions    []StateReference `json:"transitions,omitempty"`
	Roles          []StateRoleEnum  `json:"roles,omitempty"`
	Name           *LocalizedString `json:"name,omitempty"`
	Key            string           `json:"key"`
	Initial        bool             `json:"initial"`
	Description    *LocalizedString `json:"description,omitempty"`
	BuiltIn        bool             `json:"builtIn"`
}

// StateAddRolesAction implements the interface StateUpdateAction
type StateAddRolesAction struct {
	Roles []StateRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value
func (obj StateAddRolesAction) MarshalJSON() ([]byte, error) {
	type Alias StateAddRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addRoles", Alias: (*Alias)(&obj)})
}

// StateChangeInitialAction implements the interface StateUpdateAction
type StateChangeInitialAction struct {
	Initial bool `json:"initial"`
}

// MarshalJSON override to set the discriminator value
func (obj StateChangeInitialAction) MarshalJSON() ([]byte, error) {
	type Alias StateChangeInitialAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeInitial", Alias: (*Alias)(&obj)})
}

// StateChangeKeyAction implements the interface StateUpdateAction
type StateChangeKeyAction struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj StateChangeKeyAction) MarshalJSON() ([]byte, error) {
	type Alias StateChangeKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeKey", Alias: (*Alias)(&obj)})
}

// StateChangeTypeAction implements the interface StateUpdateAction
type StateChangeTypeAction struct {
	Type StateTypeEnum `json:"type"`
}

// MarshalJSON override to set the discriminator value
func (obj StateChangeTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StateChangeTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeType", Alias: (*Alias)(&obj)})
}

// StateDraft is a standalone struct
type StateDraft struct {
	Type        StateTypeEnum             `json:"type"`
	Transitions []StateResourceIdentifier `json:"transitions,omitempty"`
	Roles       []StateRoleEnum           `json:"roles,omitempty"`
	Name        *LocalizedString          `json:"name,omitempty"`
	Key         string                    `json:"key"`
	Initial     bool                      `json:"initial"`
	Description *LocalizedString          `json:"description,omitempty"`
}

// StatePagedQueryResponse is a standalone struct
type StatePagedQueryResponse struct {
	Total   int     `json:"total,omitempty"`
	Results []State `json:"results"`
	Offset  int     `json:"offset"`
	Count   int     `json:"count"`
}

// StateReference implements the interface Reference
type StateReference struct {
	ID  string `json:"id"`
	Obj *State `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StateReference) MarshalJSON() ([]byte, error) {
	type Alias StateReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "state", Alias: (*Alias)(&obj)})
}

// StateRemoveRolesAction implements the interface StateUpdateAction
type StateRemoveRolesAction struct {
	Roles []StateRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value
func (obj StateRemoveRolesAction) MarshalJSON() ([]byte, error) {
	type Alias StateRemoveRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeRoles", Alias: (*Alias)(&obj)})
}

// StateResourceIdentifier implements the interface ResourceIdentifier
type StateResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StateResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias StateResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "state", Alias: (*Alias)(&obj)})
}

// StateSetDescriptionAction implements the interface StateUpdateAction
type StateSetDescriptionAction struct {
	Description *LocalizedString `json:"description"`
}

// MarshalJSON override to set the discriminator value
func (obj StateSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// StateSetNameAction implements the interface StateUpdateAction
type StateSetNameAction struct {
	Name *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj StateSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

// StateSetRolesAction implements the interface StateUpdateAction
type StateSetRolesAction struct {
	Roles []StateRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value
func (obj StateSetRolesAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRoles", Alias: (*Alias)(&obj)})
}

// StateSetTransitionsAction implements the interface StateUpdateAction
type StateSetTransitionsAction struct {
	Transitions []StateResourceIdentifier `json:"transitions,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StateSetTransitionsAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetTransitionsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTransitions", Alias: (*Alias)(&obj)})
}

// StateUpdate is a standalone struct
type StateUpdate struct {
	Version int                 `json:"version"`
	Actions []StateUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StateUpdate) UnmarshalJSON(data []byte) error {
	type Alias StateUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorStateUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
