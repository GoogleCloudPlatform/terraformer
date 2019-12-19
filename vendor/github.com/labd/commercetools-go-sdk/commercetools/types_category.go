// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// CategoryUpdateAction uses action as discriminator attribute
type CategoryUpdateAction interface{}

func mapDiscriminatorCategoryUpdateAction(input interface{}) (CategoryUpdateAction, error) {
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
	case "addAsset":
		new := CategoryAddAssetAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAssetName":
		new := CategoryChangeAssetNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAssetOrder":
		new := CategoryChangeAssetOrderAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := CategoryChangeNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeOrderHint":
		new := CategoryChangeOrderHintAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeParent":
		new := CategoryChangeParentAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeSlug":
		new := CategoryChangeSlugAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeAsset":
		new := CategoryRemoveAssetAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetCustomField":
		new := CategorySetAssetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetCustomType":
		new := CategorySetAssetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetDescription":
		new := CategorySetAssetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetKey":
		new := CategorySetAssetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetSources":
		new := CategorySetAssetSourcesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetTags":
		new := CategorySetAssetTagsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := CategorySetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := CategorySetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := CategorySetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setExternalId":
		new := CategorySetExternalIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := CategorySetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMetaDescription":
		new := CategorySetMetaDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMetaKeywords":
		new := CategorySetMetaKeywordsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMetaTitle":
		new := CategorySetMetaTitleAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Category is of type LoggedResource
type Category struct {
	Version         int                 `json:"version"`
	LastModifiedAt  time.Time           `json:"lastModifiedAt"`
	ID              string              `json:"id"`
	CreatedAt       time.Time           `json:"createdAt"`
	LastModifiedBy  *LastModifiedBy     `json:"lastModifiedBy,omitempty"`
	CreatedBy       *CreatedBy          `json:"createdBy,omitempty"`
	Slug            *LocalizedString    `json:"slug"`
	Parent          *CategoryReference  `json:"parent,omitempty"`
	OrderHint       string              `json:"orderHint"`
	Name            *LocalizedString    `json:"name"`
	MetaTitle       *LocalizedString    `json:"metaTitle,omitempty"`
	MetaKeywords    *LocalizedString    `json:"metaKeywords,omitempty"`
	MetaDescription *LocalizedString    `json:"metaDescription,omitempty"`
	Key             string              `json:"key,omitempty"`
	ExternalID      string              `json:"externalId,omitempty"`
	Description     *LocalizedString    `json:"description,omitempty"`
	Custom          *CustomFields       `json:"custom,omitempty"`
	Assets          []Asset             `json:"assets,omitempty"`
	Ancestors       []CategoryReference `json:"ancestors"`
}

// CategoryAddAssetAction implements the interface CategoryUpdateAction
type CategoryAddAssetAction struct {
	Position float64     `json:"position,omitempty"`
	Asset    *AssetDraft `json:"asset"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryAddAssetAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryAddAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAsset", Alias: (*Alias)(&obj)})
}

// CategoryChangeAssetNameAction implements the interface CategoryUpdateAction
type CategoryChangeAssetNameAction struct {
	Name     *LocalizedString `json:"name"`
	AssetKey string           `json:"assetKey,omitempty"`
	AssetID  string           `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryChangeAssetNameAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeAssetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetName", Alias: (*Alias)(&obj)})
}

// CategoryChangeAssetOrderAction implements the interface CategoryUpdateAction
type CategoryChangeAssetOrderAction struct {
	AssetOrder []string `json:"assetOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryChangeAssetOrderAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeAssetOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetOrder", Alias: (*Alias)(&obj)})
}

// CategoryChangeNameAction implements the interface CategoryUpdateAction
type CategoryChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// CategoryChangeOrderHintAction implements the interface CategoryUpdateAction
type CategoryChangeOrderHintAction struct {
	OrderHint string `json:"orderHint"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryChangeOrderHintAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeOrderHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeOrderHint", Alias: (*Alias)(&obj)})
}

// CategoryChangeParentAction implements the interface CategoryUpdateAction
type CategoryChangeParentAction struct {
	Parent *CategoryResourceIdentifier `json:"parent"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryChangeParentAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeParentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeParent", Alias: (*Alias)(&obj)})
}

// CategoryChangeSlugAction implements the interface CategoryUpdateAction
type CategoryChangeSlugAction struct {
	Slug *LocalizedString `json:"slug"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryChangeSlugAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSlug", Alias: (*Alias)(&obj)})
}

// CategoryDraft is a standalone struct
type CategoryDraft struct {
	Slug            *LocalizedString            `json:"slug"`
	Parent          *CategoryResourceIdentifier `json:"parent,omitempty"`
	OrderHint       string                      `json:"orderHint,omitempty"`
	Name            *LocalizedString            `json:"name"`
	MetaTitle       *LocalizedString            `json:"metaTitle,omitempty"`
	MetaKeywords    *LocalizedString            `json:"metaKeywords,omitempty"`
	MetaDescription *LocalizedString            `json:"metaDescription,omitempty"`
	Key             string                      `json:"key,omitempty"`
	ExternalID      string                      `json:"externalId,omitempty"`
	Description     *LocalizedString            `json:"description,omitempty"`
	Custom          *CustomFieldsDraft          `json:"custom,omitempty"`
	Assets          []AssetDraft                `json:"assets,omitempty"`
}

// CategoryPagedQueryResponse is a standalone struct
type CategoryPagedQueryResponse struct {
	Total   int        `json:"total,omitempty"`
	Results []Category `json:"results"`
	Offset  int        `json:"offset"`
	Count   int        `json:"count"`
}

// CategoryReference implements the interface Reference
type CategoryReference struct {
	ID  string    `json:"id"`
	Obj *Category `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryReference) MarshalJSON() ([]byte, error) {
	type Alias CategoryReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "category", Alias: (*Alias)(&obj)})
}

// CategoryRemoveAssetAction implements the interface CategoryUpdateAction
type CategoryRemoveAssetAction struct {
	AssetKey string `json:"assetKey,omitempty"`
	AssetID  string `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryRemoveAssetAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryRemoveAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAsset", Alias: (*Alias)(&obj)})
}

// CategoryResourceIdentifier implements the interface ResourceIdentifier
type CategoryResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias CategoryResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "category", Alias: (*Alias)(&obj)})
}

// CategorySetAssetCustomFieldAction implements the interface CategoryUpdateAction
type CategorySetAssetCustomFieldAction struct {
	Value    interface{} `json:"value,omitempty"`
	Name     string      `json:"name"`
	AssetKey string      `json:"assetKey,omitempty"`
	AssetID  string      `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomField", Alias: (*Alias)(&obj)})
}

// CategorySetAssetCustomTypeAction implements the interface CategoryUpdateAction
type CategorySetAssetCustomTypeAction struct {
	Type     *TypeResourceIdentifier `json:"type,omitempty"`
	Fields   interface{}             `json:"fields,omitempty"`
	AssetKey string                  `json:"assetKey,omitempty"`
	AssetID  string                  `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomType", Alias: (*Alias)(&obj)})
}

// CategorySetAssetDescriptionAction implements the interface CategoryUpdateAction
type CategorySetAssetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
	AssetKey    string           `json:"assetKey,omitempty"`
	AssetID     string           `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetDescription", Alias: (*Alias)(&obj)})
}

// CategorySetAssetKeyAction implements the interface CategoryUpdateAction
type CategorySetAssetKeyAction struct {
	AssetKey string `json:"assetKey,omitempty"`
	AssetID  string `json:"assetId"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetKey", Alias: (*Alias)(&obj)})
}

// CategorySetAssetSourcesAction implements the interface CategoryUpdateAction
type CategorySetAssetSourcesAction struct {
	Sources  []AssetSource `json:"sources"`
	AssetKey string        `json:"assetKey,omitempty"`
	AssetID  string        `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetSourcesAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetSourcesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetSources", Alias: (*Alias)(&obj)})
}

// CategorySetAssetTagsAction implements the interface CategoryUpdateAction
type CategorySetAssetTagsAction struct {
	Tags     []string `json:"tags,omitempty"`
	AssetKey string   `json:"assetKey,omitempty"`
	AssetID  string   `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetTagsAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetTagsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetTags", Alias: (*Alias)(&obj)})
}

// CategorySetCustomFieldAction implements the interface CategoryUpdateAction
type CategorySetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// CategorySetCustomTypeAction implements the interface CategoryUpdateAction
type CategorySetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// CategorySetDescriptionAction implements the interface CategoryUpdateAction
type CategorySetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// CategorySetExternalIDAction implements the interface CategoryUpdateAction
type CategorySetExternalIDAction struct {
	ExternalID string `json:"externalId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetExternalIDAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetExternalIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

// CategorySetKeyAction implements the interface CategoryUpdateAction
type CategorySetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// CategorySetMetaDescriptionAction implements the interface CategoryUpdateAction
type CategorySetMetaDescriptionAction struct {
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetMetaDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaDescription", Alias: (*Alias)(&obj)})
}

// CategorySetMetaKeywordsAction implements the interface CategoryUpdateAction
type CategorySetMetaKeywordsAction struct {
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetMetaKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaKeywords", Alias: (*Alias)(&obj)})
}

// CategorySetMetaTitleAction implements the interface CategoryUpdateAction
type CategorySetMetaTitleAction struct {
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetMetaTitleAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaTitle", Alias: (*Alias)(&obj)})
}

// CategoryUpdate is a standalone struct
type CategoryUpdate struct {
	Version int                    `json:"version"`
	Actions []CategoryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CategoryUpdate) UnmarshalJSON(data []byte) error {
	type Alias CategoryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorCategoryUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
