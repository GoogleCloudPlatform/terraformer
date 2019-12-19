// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// InventoryEntryUpdateAction uses action as discriminator attribute
type InventoryEntryUpdateAction interface{}

func mapDiscriminatorInventoryEntryUpdateAction(input interface{}) (InventoryEntryUpdateAction, error) {
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
	case "addQuantity":
		new := InventoryEntryAddQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeQuantity":
		new := InventoryEntryChangeQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeQuantity":
		new := InventoryEntryRemoveQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := InventoryEntrySetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := InventoryEntrySetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setExpectedDelivery":
		new := InventoryEntrySetExpectedDeliveryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setRestockableInDays":
		new := InventoryEntrySetRestockableInDaysAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSupplyChannel":
		new := InventoryEntrySetSupplyChannelAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// InventoryEntry is of type LoggedResource
type InventoryEntry struct {
	Version           int                        `json:"version"`
	LastModifiedAt    time.Time                  `json:"lastModifiedAt"`
	ID                string                     `json:"id"`
	CreatedAt         time.Time                  `json:"createdAt"`
	LastModifiedBy    *LastModifiedBy            `json:"lastModifiedBy,omitempty"`
	CreatedBy         *CreatedBy                 `json:"createdBy,omitempty"`
	SupplyChannel     *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	SKU               string                     `json:"sku"`
	RestockableInDays int                        `json:"restockableInDays,omitempty"`
	QuantityOnStock   int                        `json:"quantityOnStock"`
	ExpectedDelivery  *time.Time                 `json:"expectedDelivery,omitempty"`
	Custom            *CustomFields              `json:"custom,omitempty"`
	AvailableQuantity int                        `json:"availableQuantity"`
}

// InventoryEntryAddQuantityAction implements the interface InventoryEntryUpdateAction
type InventoryEntryAddQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryAddQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryAddQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addQuantity", Alias: (*Alias)(&obj)})
}

// InventoryEntryChangeQuantityAction implements the interface InventoryEntryUpdateAction
type InventoryEntryChangeQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryChangeQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryChangeQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeQuantity", Alias: (*Alias)(&obj)})
}

// InventoryEntryDraft is a standalone struct
type InventoryEntryDraft struct {
	SupplyChannel     *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	SKU               string                     `json:"sku"`
	RestockableInDays int                        `json:"restockableInDays,omitempty"`
	QuantityOnStock   int                        `json:"quantityOnStock"`
	ExpectedDelivery  *time.Time                 `json:"expectedDelivery,omitempty"`
	Custom            *CustomFieldsDraft         `json:"custom,omitempty"`
}

// InventoryEntryReference implements the interface Reference
type InventoryEntryReference struct {
	ID  string          `json:"id"`
	Obj *InventoryEntry `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryReference) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "inventory-entry", Alias: (*Alias)(&obj)})
}

// InventoryEntryRemoveQuantityAction implements the interface InventoryEntryUpdateAction
type InventoryEntryRemoveQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryRemoveQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryRemoveQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeQuantity", Alias: (*Alias)(&obj)})
}

// InventoryEntryResourceIdentifier implements the interface ResourceIdentifier
type InventoryEntryResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "inventory-entry", Alias: (*Alias)(&obj)})
}

// InventoryEntrySetCustomFieldAction implements the interface InventoryEntryUpdateAction
type InventoryEntrySetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntrySetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// InventoryEntrySetCustomTypeAction implements the interface InventoryEntryUpdateAction
type InventoryEntrySetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntrySetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// InventoryEntrySetExpectedDeliveryAction implements the interface InventoryEntryUpdateAction
type InventoryEntrySetExpectedDeliveryAction struct {
	ExpectedDelivery *time.Time `json:"expectedDelivery,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntrySetExpectedDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetExpectedDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExpectedDelivery", Alias: (*Alias)(&obj)})
}

// InventoryEntrySetRestockableInDaysAction implements the interface InventoryEntryUpdateAction
type InventoryEntrySetRestockableInDaysAction struct {
	RestockableInDays int `json:"restockableInDays,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntrySetRestockableInDaysAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetRestockableInDaysAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRestockableInDays", Alias: (*Alias)(&obj)})
}

// InventoryEntrySetSupplyChannelAction implements the interface InventoryEntryUpdateAction
type InventoryEntrySetSupplyChannelAction struct {
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntrySetSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupplyChannel", Alias: (*Alias)(&obj)})
}

// InventoryEntryUpdate is a standalone struct
type InventoryEntryUpdate struct {
	Version int                          `json:"version"`
	Actions []InventoryEntryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InventoryEntryUpdate) UnmarshalJSON(data []byte) error {
	type Alias InventoryEntryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorInventoryEntryUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// InventoryPagedQueryResponse is a standalone struct
type InventoryPagedQueryResponse struct {
	Total   int              `json:"total,omitempty"`
	Results []InventoryEntry `json:"results"`
	Offset  int              `json:"offset"`
	Count   int              `json:"count"`
}
