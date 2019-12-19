// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// CustomerGroupUpdateAction uses action as discriminator attribute
type CustomerGroupUpdateAction interface{}

func mapDiscriminatorCustomerGroupUpdateAction(input interface{}) (CustomerGroupUpdateAction, error) {
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
	case "changeName":
		new := CustomerGroupChangeNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := CustomerGroupSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := CustomerGroupSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := CustomerGroupSetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// CustomerGroup is of type LoggedResource
type CustomerGroup struct {
	Version        int             `json:"version"`
	LastModifiedAt time.Time       `json:"lastModifiedAt"`
	ID             string          `json:"id"`
	CreatedAt      time.Time       `json:"createdAt"`
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	CreatedBy      *CreatedBy      `json:"createdBy,omitempty"`
	Name           string          `json:"name"`
	Key            string          `json:"key,omitempty"`
	Custom         *CustomFields   `json:"custom,omitempty"`
}

// CustomerGroupChangeNameAction implements the interface CustomerGroupUpdateAction
type CustomerGroupChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// CustomerGroupDraft is a standalone struct
type CustomerGroupDraft struct {
	Key       string        `json:"key,omitempty"`
	GroupName string        `json:"groupName"`
	Custom    *CustomFields `json:"custom,omitempty"`
}

// CustomerGroupPagedQueryResponse is a standalone struct
type CustomerGroupPagedQueryResponse struct {
	Total   int             `json:"total,omitempty"`
	Results []CustomerGroup `json:"results"`
	Offset  int             `json:"offset"`
	Count   int             `json:"count"`
}

// CustomerGroupReference implements the interface Reference
type CustomerGroupReference struct {
	ID  string         `json:"id"`
	Obj *CustomerGroup `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "customer-group", Alias: (*Alias)(&obj)})
}

// CustomerGroupResourceIdentifier implements the interface ResourceIdentifier
type CustomerGroupResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "customer-group", Alias: (*Alias)(&obj)})
}

// CustomerGroupSetCustomFieldAction implements the interface CustomerGroupUpdateAction
type CustomerGroupSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// CustomerGroupSetCustomTypeAction implements the interface CustomerGroupUpdateAction
type CustomerGroupSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// CustomerGroupSetKeyAction implements the interface CustomerGroupUpdateAction
type CustomerGroupSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// CustomerGroupUpdate is a standalone struct
type CustomerGroupUpdate struct {
	Version int                         `json:"version"`
	Actions []CustomerGroupUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerGroupUpdate) UnmarshalJSON(data []byte) error {
	type Alias CustomerGroupUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorCustomerGroupUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
