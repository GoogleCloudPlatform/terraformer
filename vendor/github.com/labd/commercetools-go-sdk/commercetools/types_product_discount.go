// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ProductDiscountUpdateAction uses action as discriminator attribute
type ProductDiscountUpdateAction interface{}

func mapDiscriminatorProductDiscountUpdateAction(input interface{}) (ProductDiscountUpdateAction, error) {
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
	case "changeIsActive":
		new := ProductDiscountChangeIsActiveAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := ProductDiscountChangeNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changePredicate":
		new := ProductDiscountChangePredicateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeSortOrder":
		new := ProductDiscountChangeSortOrderAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeValue":
		new := ProductDiscountChangeValueAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.Value != nil {
			new.Value, err = mapDiscriminatorProductDiscountValue(new.Value)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "setDescription":
		new := ProductDiscountSetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := ProductDiscountSetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidFrom":
		new := ProductDiscountSetValidFromAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidFromAndUntil":
		new := ProductDiscountSetValidFromAndUntilAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidUntil":
		new := ProductDiscountSetValidUntilAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// ProductDiscountValue uses type as discriminator attribute
type ProductDiscountValue interface{}

func mapDiscriminatorProductDiscountValue(input interface{}) (ProductDiscountValue, error) {
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
	case "absolute":
		new := ProductDiscountValueAbsolute{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "external":
		new := ProductDiscountValueExternal{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "relative":
		new := ProductDiscountValueRelative{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// ProductDiscount is of type LoggedResource
type ProductDiscount struct {
	Version        int                  `json:"version"`
	LastModifiedAt time.Time            `json:"lastModifiedAt"`
	ID             string               `json:"id"`
	CreatedAt      time.Time            `json:"createdAt"`
	LastModifiedBy *LastModifiedBy      `json:"lastModifiedBy,omitempty"`
	CreatedBy      *CreatedBy           `json:"createdBy,omitempty"`
	Value          ProductDiscountValue `json:"value"`
	ValidUntil     *time.Time           `json:"validUntil,omitempty"`
	ValidFrom      *time.Time           `json:"validFrom,omitempty"`
	SortOrder      string               `json:"sortOrder"`
	References     []Reference          `json:"references"`
	Predicate      string               `json:"predicate"`
	Name           *LocalizedString     `json:"name"`
	Key            string               `json:"key,omitempty"`
	IsActive       bool                 `json:"isActive"`
	Description    *LocalizedString     `json:"description,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscount) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscount
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
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorProductDiscountValue(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// ProductDiscountChangeIsActiveAction implements the interface ProductDiscountUpdateAction
type ProductDiscountChangeIsActiveAction struct {
	IsActive bool `json:"isActive"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

// ProductDiscountChangeNameAction implements the interface ProductDiscountUpdateAction
type ProductDiscountChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ProductDiscountChangePredicateAction implements the interface ProductDiscountUpdateAction
type ProductDiscountChangePredicateAction struct {
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountChangePredicateAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangePredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePredicate", Alias: (*Alias)(&obj)})
}

// ProductDiscountChangeSortOrderAction implements the interface ProductDiscountUpdateAction
type ProductDiscountChangeSortOrderAction struct {
	SortOrder string `json:"sortOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountChangeSortOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeSortOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSortOrder", Alias: (*Alias)(&obj)})
}

// ProductDiscountChangeValueAction implements the interface ProductDiscountUpdateAction
type ProductDiscountChangeValueAction struct {
	Value ProductDiscountValue `json:"value"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountChangeValueAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeValue", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountChangeValueAction) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountChangeValueAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorProductDiscountValue(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// ProductDiscountDraft is a standalone struct
type ProductDiscountDraft struct {
	Value       ProductDiscountValue `json:"value"`
	ValidUntil  *time.Time           `json:"validUntil,omitempty"`
	ValidFrom   *time.Time           `json:"validFrom,omitempty"`
	SortOrder   string               `json:"sortOrder"`
	Predicate   string               `json:"predicate"`
	Name        *LocalizedString     `json:"name"`
	Key         string               `json:"key,omitempty"`
	IsActive    bool                 `json:"isActive"`
	Description *LocalizedString     `json:"description,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountDraft) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorProductDiscountValue(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// ProductDiscountMatchQuery is a standalone struct
type ProductDiscountMatchQuery struct {
	VariantID float64 `json:"variantId"`
	Staged    bool    `json:"staged"`
	ProductID string  `json:"productId"`
	Price     *Price  `json:"price"`
}

// ProductDiscountPagedQueryResponse is a standalone struct
type ProductDiscountPagedQueryResponse struct {
	Total   int               `json:"total,omitempty"`
	Results []ProductDiscount `json:"results"`
	Offset  int               `json:"offset"`
	Count   int               `json:"count"`
}

// ProductDiscountReference implements the interface Reference
type ProductDiscountReference struct {
	ID  string           `json:"id"`
	Obj *ProductDiscount `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "product-discount", Alias: (*Alias)(&obj)})
}

// ProductDiscountResourceIdentifier implements the interface ResourceIdentifier
type ProductDiscountResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "product-discount", Alias: (*Alias)(&obj)})
}

// ProductDiscountSetDescriptionAction implements the interface ProductDiscountUpdateAction
type ProductDiscountSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// ProductDiscountSetKeyAction implements the interface ProductDiscountUpdateAction
type ProductDiscountSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ProductDiscountSetValidFromAction implements the interface ProductDiscountUpdateAction
type ProductDiscountSetValidFromAction struct {
	ValidFrom *time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

// ProductDiscountSetValidFromAndUntilAction implements the interface ProductDiscountUpdateAction
type ProductDiscountSetValidFromAndUntilAction struct {
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	ValidFrom  *time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

// ProductDiscountSetValidUntilAction implements the interface ProductDiscountUpdateAction
type ProductDiscountSetValidUntilAction struct {
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}

// ProductDiscountUpdate is a standalone struct
type ProductDiscountUpdate struct {
	Version int                           `json:"version"`
	Actions []ProductDiscountUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorProductDiscountUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// ProductDiscountValueAbsolute implements the interface ProductDiscountValue
type ProductDiscountValueAbsolute struct {
	Money []Money `json:"money"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountValueAbsolute) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueAbsolute
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "absolute", Alias: (*Alias)(&obj)})
}

// ProductDiscountValueExternal implements the interface ProductDiscountValue
type ProductDiscountValueExternal struct{}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountValueExternal) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueExternal
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "external", Alias: (*Alias)(&obj)})
}

// ProductDiscountValueRelative implements the interface ProductDiscountValue
type ProductDiscountValueRelative struct {
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountValueRelative) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueRelative
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "relative", Alias: (*Alias)(&obj)})
}
