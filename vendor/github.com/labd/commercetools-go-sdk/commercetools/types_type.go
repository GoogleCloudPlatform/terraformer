// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ResourceTypeID is an enum type
type ResourceTypeID string

// Enum values for ResourceTypeID
const (
	ResourceTypeIDAsset                       ResourceTypeID = "asset"
	ResourceTypeIDCategory                    ResourceTypeID = "category"
	ResourceTypeIDChannel                     ResourceTypeID = "channel"
	ResourceTypeIDCustomer                    ResourceTypeID = "customer"
	ResourceTypeIDOrder                       ResourceTypeID = "order"
	ResourceTypeIDOrderEdit                   ResourceTypeID = "order-edit"
	ResourceTypeIDInventoryEntry              ResourceTypeID = "inventory-entry"
	ResourceTypeIDLineItem                    ResourceTypeID = "line-item"
	ResourceTypeIDCustomLineItem              ResourceTypeID = "custom-line-item"
	ResourceTypeIDProductPrice                ResourceTypeID = "product-price"
	ResourceTypeIDPayment                     ResourceTypeID = "payment"
	ResourceTypeIDPaymentInterfaceInteraction ResourceTypeID = "payment-interface-interaction"
	ResourceTypeIDReview                      ResourceTypeID = "review"
	ResourceTypeIDShoppingList                ResourceTypeID = "shopping-list"
	ResourceTypeIDShoppingListTextLineItem    ResourceTypeID = "shopping-list-text-line-item"
	ResourceTypeIDDiscountCode                ResourceTypeID = "discount-code"
	ResourceTypeIDCartDiscount                ResourceTypeID = "cart-discount"
	ResourceTypeIDCustomerGroup               ResourceTypeID = "customer-group"
)

// TypeTextInputHint is an enum type
type TypeTextInputHint string

// Enum values for TypeTextInputHint
const (
	TypeTextInputHintSingleLine TypeTextInputHint = "SingleLine"
	TypeTextInputHintMultiLine  TypeTextInputHint = "MultiLine"
)

// FieldContainer is a map
type FieldContainer map[string]interface{}

// FieldType uses name as discriminator attribute
type FieldType interface{}

func mapDiscriminatorFieldType(input interface{}) (FieldType, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["name"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'name'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "Boolean":
		new := CustomFieldBooleanType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DateTime":
		new := CustomFieldDateTimeType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Date":
		new := CustomFieldDateType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Enum":
		new := CustomFieldEnumType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LocalizedEnum":
		new := CustomFieldLocalizedEnumType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LocalizedString":
		new := CustomFieldLocalizedStringType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Money":
		new := CustomFieldMoneyType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Number":
		new := CustomFieldNumberType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Reference":
		new := CustomFieldReferenceType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Set":
		new := CustomFieldSetType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.ElementType != nil {
			new.ElementType, err = mapDiscriminatorFieldType(new.ElementType)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "String":
		new := CustomFieldStringType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Time":
		new := CustomFieldTimeType{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// TypeUpdateAction uses action as discriminator attribute
type TypeUpdateAction interface{}

func mapDiscriminatorTypeUpdateAction(input interface{}) (TypeUpdateAction, error) {
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
	case "addEnumValue":
		new := TypeAddEnumValueAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addFieldDefinition":
		new := TypeAddFieldDefinitionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addLocalizedEnumValue":
		new := TypeAddLocalizedEnumValueAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeEnumValueLabel":
		new := TypeChangeEnumValueLabelAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeEnumValueOrder":
		new := TypeChangeEnumValueOrderAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeFieldDefinitionLabel":
		new := TypeChangeFieldDefinitionLabelAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeFieldDefinitionOrder":
		new := TypeChangeFieldDefinitionOrderAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeInputHint":
		new := TypeChangeInputHintAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeKey":
		new := TypeChangeKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLabel":
		new := TypeChangeLabelAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLocalizedEnumValueLabel":
		new := TypeChangeLocalizedEnumValueLabelAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLocalizedEnumValueOrder":
		new := TypeChangeLocalizedEnumValueOrderAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := TypeChangeNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeFieldDefinition":
		new := TypeRemoveFieldDefinitionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := TypeSetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// CustomFieldBooleanType implements the interface FieldType
type CustomFieldBooleanType struct{}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldBooleanType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldBooleanType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Boolean", Alias: (*Alias)(&obj)})
}

// CustomFieldDateTimeType implements the interface FieldType
type CustomFieldDateTimeType struct{}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldDateTimeType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldDateTimeType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "DateTime", Alias: (*Alias)(&obj)})
}

// CustomFieldDateType implements the interface FieldType
type CustomFieldDateType struct{}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldDateType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldDateType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Date", Alias: (*Alias)(&obj)})
}

// CustomFieldEnumType implements the interface FieldType
type CustomFieldEnumType struct {
	Values []CustomFieldEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldEnumType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldEnumType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Enum", Alias: (*Alias)(&obj)})
}

// CustomFieldEnumValue is a standalone struct
type CustomFieldEnumValue struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

// CustomFieldLocalizedEnumType implements the interface FieldType
type CustomFieldLocalizedEnumType struct {
	Values []CustomFieldLocalizedEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldLocalizedEnumType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldLocalizedEnumType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "LocalizedEnum", Alias: (*Alias)(&obj)})
}

// CustomFieldLocalizedEnumValue is a standalone struct
type CustomFieldLocalizedEnumValue struct {
	Label *LocalizedString `json:"label"`
	Key   string           `json:"key"`
}

// CustomFieldLocalizedStringType implements the interface FieldType
type CustomFieldLocalizedStringType struct{}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldLocalizedStringType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldLocalizedStringType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "LocalizedString", Alias: (*Alias)(&obj)})
}

// CustomFieldMoneyType implements the interface FieldType
type CustomFieldMoneyType struct{}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldMoneyType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldMoneyType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Money", Alias: (*Alias)(&obj)})
}

// CustomFieldNumberType implements the interface FieldType
type CustomFieldNumberType struct{}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldNumberType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldNumberType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Number", Alias: (*Alias)(&obj)})
}

// CustomFieldReferenceType implements the interface FieldType
type CustomFieldReferenceType struct {
	ReferenceTypeID ReferenceTypeID `json:"referenceTypeId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldReferenceType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldReferenceType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Reference", Alias: (*Alias)(&obj)})
}

// CustomFieldSetType implements the interface FieldType
type CustomFieldSetType struct {
	ElementType FieldType `json:"elementType"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldSetType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldSetType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Set", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomFieldSetType) UnmarshalJSON(data []byte) error {
	type Alias CustomFieldSetType
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ElementType != nil {
		var err error
		obj.ElementType, err = mapDiscriminatorFieldType(obj.ElementType)
		if err != nil {
			return err
		}
	}

	return nil
}

// CustomFieldStringType implements the interface FieldType
type CustomFieldStringType struct{}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldStringType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldStringType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "String", Alias: (*Alias)(&obj)})
}

// CustomFieldTimeType implements the interface FieldType
type CustomFieldTimeType struct{}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldTimeType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldTimeType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Time", Alias: (*Alias)(&obj)})
}

// CustomFields is a standalone struct
type CustomFields struct {
	Type   *TypeReference  `json:"type"`
	Fields *FieldContainer `json:"fields"`
}

// CustomFieldsDraft is a standalone struct
type CustomFieldsDraft struct {
	Type   *TypeResourceIdentifier `json:"type"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// FieldDefinition is a standalone struct
type FieldDefinition struct {
	Type      FieldType         `json:"type"`
	Required  bool              `json:"required"`
	Name      string            `json:"name"`
	Label     *LocalizedString  `json:"label"`
	InputHint TypeTextInputHint `json:"inputHint,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *FieldDefinition) UnmarshalJSON(data []byte) error {
	type Alias FieldDefinition
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Type != nil {
		var err error
		obj.Type, err = mapDiscriminatorFieldType(obj.Type)
		if err != nil {
			return err
		}
	}

	return nil
}

// Type is of type LoggedResource
type Type struct {
	Version          int               `json:"version"`
	LastModifiedAt   time.Time         `json:"lastModifiedAt"`
	ID               string            `json:"id"`
	CreatedAt        time.Time         `json:"createdAt"`
	LastModifiedBy   *LastModifiedBy   `json:"lastModifiedBy,omitempty"`
	CreatedBy        *CreatedBy        `json:"createdBy,omitempty"`
	ResourceTypeIds  []ResourceTypeID  `json:"resourceTypeIds"`
	Name             *LocalizedString  `json:"name"`
	Key              string            `json:"key"`
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions"`
	Description      *LocalizedString  `json:"description,omitempty"`
}

// TypeAddEnumValueAction implements the interface TypeUpdateAction
type TypeAddEnumValueAction struct {
	Value     *CustomFieldEnumValue `json:"value"`
	FieldName string                `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeAddEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addEnumValue", Alias: (*Alias)(&obj)})
}

// TypeAddFieldDefinitionAction implements the interface TypeUpdateAction
type TypeAddFieldDefinitionAction struct {
	FieldDefinition *FieldDefinition `json:"fieldDefinition"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeAddFieldDefinitionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddFieldDefinitionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addFieldDefinition", Alias: (*Alias)(&obj)})
}

// TypeAddLocalizedEnumValueAction implements the interface TypeUpdateAction
type TypeAddLocalizedEnumValueAction struct {
	Value     *CustomFieldLocalizedEnumValue `json:"value"`
	FieldName string                         `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeAddLocalizedEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddLocalizedEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLocalizedEnumValue", Alias: (*Alias)(&obj)})
}

// TypeChangeEnumValueLabelAction implements the interface TypeUpdateAction
type TypeChangeEnumValueLabelAction struct {
	Value     *CustomFieldEnumValue `json:"value"`
	FieldName string                `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeEnumValueLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeEnumValueLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEnumValueLabel", Alias: (*Alias)(&obj)})
}

// TypeChangeEnumValueOrderAction implements the interface TypeUpdateAction
type TypeChangeEnumValueOrderAction struct {
	Keys      []string `json:"keys"`
	FieldName string   `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEnumValueOrder", Alias: (*Alias)(&obj)})
}

// TypeChangeFieldDefinitionLabelAction implements the interface TypeUpdateAction
type TypeChangeFieldDefinitionLabelAction struct {
	Label     *LocalizedString `json:"label"`
	FieldName string           `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeFieldDefinitionLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeFieldDefinitionLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeFieldDefinitionLabel", Alias: (*Alias)(&obj)})
}

// TypeChangeFieldDefinitionOrderAction implements the interface TypeUpdateAction
type TypeChangeFieldDefinitionOrderAction struct {
	FieldNames []string `json:"fieldNames"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeFieldDefinitionOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeFieldDefinitionOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeFieldDefinitionOrder", Alias: (*Alias)(&obj)})
}

// TypeChangeInputHintAction implements the interface TypeUpdateAction
type TypeChangeInputHintAction struct {
	InputHint TypeTextInputHint `json:"inputHint"`
	FieldName string            `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeInputHintAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeInputHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeInputHint", Alias: (*Alias)(&obj)})
}

// TypeChangeKeyAction implements the interface TypeUpdateAction
type TypeChangeKeyAction struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeKeyAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeKey", Alias: (*Alias)(&obj)})
}

// TypeChangeLabelAction implements the interface TypeUpdateAction
type TypeChangeLabelAction struct {
	Label     *LocalizedString `json:"label"`
	FieldName string           `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLabel", Alias: (*Alias)(&obj)})
}

// TypeChangeLocalizedEnumValueLabelAction implements the interface TypeUpdateAction
type TypeChangeLocalizedEnumValueLabelAction struct {
	Value     *CustomFieldLocalizedEnumValue `json:"value"`
	FieldName string                         `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeLocalizedEnumValueLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLocalizedEnumValueLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLocalizedEnumValueLabel", Alias: (*Alias)(&obj)})
}

// TypeChangeLocalizedEnumValueOrderAction implements the interface TypeUpdateAction
type TypeChangeLocalizedEnumValueOrderAction struct {
	Keys      []string `json:"keys"`
	FieldName string   `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeLocalizedEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLocalizedEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLocalizedEnumValueOrder", Alias: (*Alias)(&obj)})
}

// TypeChangeNameAction implements the interface TypeUpdateAction
type TypeChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// TypeDraft is a standalone struct
type TypeDraft struct {
	ResourceTypeIds  []ResourceTypeID  `json:"resourceTypeIds"`
	Name             *LocalizedString  `json:"name"`
	Key              string            `json:"key"`
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions,omitempty"`
	Description      *LocalizedString  `json:"description,omitempty"`
}

// TypePagedQueryResponse is a standalone struct
type TypePagedQueryResponse struct {
	Total   int    `json:"total,omitempty"`
	Results []Type `json:"results"`
	Offset  int    `json:"offset"`
	Count   int    `json:"count"`
}

// TypeReference implements the interface Reference
type TypeReference struct {
	ID  string `json:"id"`
	Obj *Type  `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeReference) MarshalJSON() ([]byte, error) {
	type Alias TypeReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "type", Alias: (*Alias)(&obj)})
}

// TypeRemoveFieldDefinitionAction implements the interface TypeUpdateAction
type TypeRemoveFieldDefinitionAction struct {
	FieldName string `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeRemoveFieldDefinitionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeRemoveFieldDefinitionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeFieldDefinition", Alias: (*Alias)(&obj)})
}

// TypeResourceIdentifier implements the interface ResourceIdentifier
type TypeResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias TypeResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "type", Alias: (*Alias)(&obj)})
}

// TypeSetDescriptionAction implements the interface TypeUpdateAction
type TypeSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// TypeUpdate is a standalone struct
type TypeUpdate struct {
	Version int                `json:"version"`
	Actions []TypeUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TypeUpdate) UnmarshalJSON(data []byte) error {
	type Alias TypeUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorTypeUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
