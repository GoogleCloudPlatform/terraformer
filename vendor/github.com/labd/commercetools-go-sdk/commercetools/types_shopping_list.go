// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ShoppingListUpdateAction uses action as discriminator attribute
type ShoppingListUpdateAction interface{}

func mapDiscriminatorShoppingListUpdateAction(input interface{}) (ShoppingListUpdateAction, error) {
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
	case "addLineItem":
		new := ShoppingListAddLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addTextLineItem":
		new := ShoppingListAddTextLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemQuantity":
		new := ShoppingListChangeLineItemQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemsOrder":
		new := ShoppingListChangeLineItemsOrderAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := ShoppingListChangeNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemName":
		new := ShoppingListChangeTextLineItemNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemQuantity":
		new := ShoppingListChangeTextLineItemQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemsOrder":
		new := ShoppingListChangeTextLineItemsOrderAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeLineItem":
		new := ShoppingListRemoveLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeTextLineItem":
		new := ShoppingListRemoveTextLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAnonymousId":
		new := ShoppingListSetAnonymousIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := ShoppingListSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := ShoppingListSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomer":
		new := ShoppingListSetCustomerAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeleteDaysAfterLastModification":
		new := ShoppingListSetDeleteDaysAfterLastModificationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := ShoppingListSetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := ShoppingListSetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomField":
		new := ShoppingListSetLineItemCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomType":
		new := ShoppingListSetLineItemCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSlug":
		new := ShoppingListSetSlugAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemCustomField":
		new := ShoppingListSetTextLineItemCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemCustomType":
		new := ShoppingListSetTextLineItemCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemDescription":
		new := ShoppingListSetTextLineItemDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// ShoppingList is of type LoggedResource
type ShoppingList struct {
	Version                         int                    `json:"version"`
	LastModifiedAt                  time.Time              `json:"lastModifiedAt"`
	ID                              string                 `json:"id"`
	CreatedAt                       time.Time              `json:"createdAt"`
	LastModifiedBy                  *LastModifiedBy        `json:"lastModifiedBy,omitempty"`
	CreatedBy                       *CreatedBy             `json:"createdBy,omitempty"`
	TextLineItems                   []TextLineItem         `json:"textLineItems,omitempty"`
	Slug                            *LocalizedString       `json:"slug,omitempty"`
	Name                            *LocalizedString       `json:"name"`
	LineItems                       []ShoppingListLineItem `json:"lineItems,omitempty"`
	Key                             string                 `json:"key,omitempty"`
	Description                     *LocalizedString       `json:"description,omitempty"`
	DeleteDaysAfterLastModification int                    `json:"deleteDaysAfterLastModification,omitempty"`
	Customer                        *CustomerReference     `json:"customer,omitempty"`
	Custom                          *CustomFields          `json:"custom,omitempty"`
	AnonymousID                     string                 `json:"anonymousId,omitempty"`
}

// ShoppingListAddLineItemAction implements the interface ShoppingListUpdateAction
type ShoppingListAddLineItemAction struct {
	VariantID int                `json:"variantId,omitempty"`
	SKU       string             `json:"sku,omitempty"`
	Quantity  int                `json:"quantity,omitempty"`
	ProductID string             `json:"productId,omitempty"`
	Custom    *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt   *time.Time         `json:"addedAt,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

// ShoppingListAddTextLineItemAction implements the interface ShoppingListUpdateAction
type ShoppingListAddTextLineItemAction struct {
	Quantity    int                `json:"quantity,omitempty"`
	Name        *LocalizedString   `json:"name"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt     *time.Time         `json:"addedAt,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListAddTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListAddTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTextLineItem", Alias: (*Alias)(&obj)})
}

// ShoppingListChangeLineItemQuantityAction implements the interface ShoppingListUpdateAction
type ShoppingListChangeLineItemQuantityAction struct {
	Quantity   int    `json:"quantity"`
	LineItemID string `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

// ShoppingListChangeLineItemsOrderAction implements the interface ShoppingListUpdateAction
type ShoppingListChangeLineItemsOrderAction struct {
	LineItemOrder []string `json:"lineItemOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemsOrder", Alias: (*Alias)(&obj)})
}

// ShoppingListChangeNameAction implements the interface ShoppingListUpdateAction
type ShoppingListChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ShoppingListChangeTextLineItemNameAction implements the interface ShoppingListUpdateAction
type ShoppingListChangeTextLineItemNameAction struct {
	TextLineItemID string           `json:"textLineItemId"`
	Name           *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeTextLineItemNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemName", Alias: (*Alias)(&obj)})
}

// ShoppingListChangeTextLineItemQuantityAction implements the interface ShoppingListUpdateAction
type ShoppingListChangeTextLineItemQuantityAction struct {
	TextLineItemID string `json:"textLineItemId"`
	Quantity       int    `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeTextLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemQuantity", Alias: (*Alias)(&obj)})
}

// ShoppingListChangeTextLineItemsOrderAction implements the interface ShoppingListUpdateAction
type ShoppingListChangeTextLineItemsOrderAction struct {
	TextLineItemOrder []string `json:"textLineItemOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListChangeTextLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListChangeTextLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemsOrder", Alias: (*Alias)(&obj)})
}

// ShoppingListDraft is a standalone struct
type ShoppingListDraft struct {
	TextLineItems                   []TextLineItemDraft         `json:"textLineItems,omitempty"`
	Slug                            *LocalizedString            `json:"slug,omitempty"`
	Name                            *LocalizedString            `json:"name"`
	LineItems                       []ShoppingListLineItemDraft `json:"lineItems,omitempty"`
	Key                             string                      `json:"key,omitempty"`
	Description                     *LocalizedString            `json:"description,omitempty"`
	DeleteDaysAfterLastModification int                         `json:"deleteDaysAfterLastModification,omitempty"`
	Customer                        *CustomerResourceIdentifier `json:"customer,omitempty"`
	Custom                          *CustomFieldsDraft          `json:"custom,omitempty"`
	AnonymousID                     string                      `json:"anonymousId,omitempty"`
}

// ShoppingListLineItem is a standalone struct
type ShoppingListLineItem struct {
	VariantID     int                   `json:"variantId,omitempty"`
	Variant       *ProductVariant       `json:"variant,omitempty"`
	Quantity      float64               `json:"quantity"`
	ProductType   *ProductTypeReference `json:"productType"`
	ProductSlug   *LocalizedString      `json:"productSlug,omitempty"`
	ProductID     string                `json:"productId"`
	Name          *LocalizedString      `json:"name"`
	ID            string                `json:"id"`
	DeactivatedAt *time.Time            `json:"deactivatedAt,omitempty"`
	Custom        *CustomFields         `json:"custom,omitempty"`
	AddedAt       time.Time             `json:"addedAt"`
}

// ShoppingListLineItemDraft is a standalone struct
type ShoppingListLineItemDraft struct {
	VariantID int                `json:"variantId,omitempty"`
	SKU       string             `json:"sku,omitempty"`
	Quantity  float64            `json:"quantity,omitempty"`
	ProductID string             `json:"productId,omitempty"`
	Custom    *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt   *time.Time         `json:"addedAt,omitempty"`
}

// ShoppingListPagedQueryResponse is a standalone struct
type ShoppingListPagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Results []ShoppingList `json:"results"`
	Offset  int            `json:"offset"`
	Count   int            `json:"count"`
}

// ShoppingListReference implements the interface Reference
type ShoppingListReference struct {
	ID  string        `json:"id"`
	Obj *ShoppingList `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListReference) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "shopping-list", Alias: (*Alias)(&obj)})
}

// ShoppingListRemoveLineItemAction implements the interface ShoppingListUpdateAction
type ShoppingListRemoveLineItemAction struct {
	Quantity   int    `json:"quantity,omitempty"`
	LineItemID string `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

// ShoppingListRemoveTextLineItemAction implements the interface ShoppingListUpdateAction
type ShoppingListRemoveTextLineItemAction struct {
	TextLineItemID string `json:"textLineItemId"`
	Quantity       int    `json:"quantity,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListRemoveTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListRemoveTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTextLineItem", Alias: (*Alias)(&obj)})
}

// ShoppingListResourceIdentifier implements the interface ResourceIdentifier
type ShoppingListResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "shopping-list", Alias: (*Alias)(&obj)})
}

// ShoppingListSetAnonymousIDAction implements the interface ShoppingListUpdateAction
type ShoppingListSetAnonymousIDAction struct {
	AnonymousID string `json:"anonymousId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetAnonymousIDAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetAnonymousIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

// ShoppingListSetCustomFieldAction implements the interface ShoppingListUpdateAction
type ShoppingListSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// ShoppingListSetCustomTypeAction implements the interface ShoppingListUpdateAction
type ShoppingListSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// ShoppingListSetCustomerAction implements the interface ShoppingListUpdateAction
type ShoppingListSetCustomerAction struct {
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

// ShoppingListSetDeleteDaysAfterLastModificationAction implements the interface ShoppingListUpdateAction
type ShoppingListSetDeleteDaysAfterLastModificationAction struct {
	DeleteDaysAfterLastModification int `json:"deleteDaysAfterLastModification,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

// ShoppingListSetDescriptionAction implements the interface ShoppingListUpdateAction
type ShoppingListSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// ShoppingListSetKeyAction implements the interface ShoppingListUpdateAction
type ShoppingListSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ShoppingListSetLineItemCustomFieldAction implements the interface ShoppingListUpdateAction
type ShoppingListSetLineItemCustomFieldAction struct {
	Value      interface{} `json:"value,omitempty"`
	Name       string      `json:"name"`
	LineItemID string      `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

// ShoppingListSetLineItemCustomTypeAction implements the interface ShoppingListUpdateAction
type ShoppingListSetLineItemCustomTypeAction struct {
	Type       *TypeResourceIdentifier `json:"type,omitempty"`
	LineItemID string                  `json:"lineItemId"`
	Fields     *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

// ShoppingListSetSlugAction implements the interface ShoppingListUpdateAction
type ShoppingListSetSlugAction struct {
	Slug *LocalizedString `json:"slug,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetSlugAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSlug", Alias: (*Alias)(&obj)})
}

// ShoppingListSetTextLineItemCustomFieldAction implements the interface ShoppingListUpdateAction
type ShoppingListSetTextLineItemCustomFieldAction struct {
	Value          interface{} `json:"value,omitempty"`
	TextLineItemID string      `json:"textLineItemId"`
	Name           string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetTextLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomField", Alias: (*Alias)(&obj)})
}

// ShoppingListSetTextLineItemCustomTypeAction implements the interface ShoppingListUpdateAction
type ShoppingListSetTextLineItemCustomTypeAction struct {
	Type           *TypeResourceIdentifier `json:"type,omitempty"`
	TextLineItemID string                  `json:"textLineItemId"`
	Fields         *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetTextLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomType", Alias: (*Alias)(&obj)})
}

// ShoppingListSetTextLineItemDescriptionAction implements the interface ShoppingListUpdateAction
type ShoppingListSetTextLineItemDescriptionAction struct {
	TextLineItemID string           `json:"textLineItemId"`
	Description    *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShoppingListSetTextLineItemDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShoppingListSetTextLineItemDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemDescription", Alias: (*Alias)(&obj)})
}

// ShoppingListUpdate is a standalone struct
type ShoppingListUpdate struct {
	Version int                        `json:"version"`
	Actions []ShoppingListUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShoppingListUpdate) UnmarshalJSON(data []byte) error {
	type Alias ShoppingListUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorShoppingListUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// TextLineItem is a standalone struct
type TextLineItem struct {
	Quantity    float64          `json:"quantity"`
	Name        *LocalizedString `json:"name"`
	ID          string           `json:"id"`
	Description *LocalizedString `json:"description,omitempty"`
	Custom      *CustomFields    `json:"custom,omitempty"`
	AddedAt     time.Time        `json:"addedAt"`
}

// TextLineItemDraft is a standalone struct
type TextLineItemDraft struct {
	Quantity    float64            `json:"quantity,omitempty"`
	Name        *LocalizedString   `json:"name"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt     *time.Time         `json:"addedAt,omitempty"`
}
