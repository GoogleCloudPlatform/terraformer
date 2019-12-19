// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// FacetTypes is an enum type
type FacetTypes string

// Enum values for FacetTypes
const (
	FacetTypesTerms  FacetTypes = "terms"
	FacetTypesRange  FacetTypes = "range"
	FacetTypesFilter FacetTypes = "filter"
)

// TermFacetResultType is an enum type
type TermFacetResultType string

// Enum values for TermFacetResultType
const (
	TermFacetResultTypeText     TermFacetResultType = "text"
	TermFacetResultTypeDate     TermFacetResultType = "date"
	TermFacetResultTypeTime     TermFacetResultType = "time"
	TermFacetResultTypeDatetime TermFacetResultType = "datetime"
	TermFacetResultTypeBoolean  TermFacetResultType = "boolean"
	TermFacetResultTypeNumber   TermFacetResultType = "number"
)

// CategoryOrderHints is a map
type CategoryOrderHints map[string]string

// FacetResult uses type as discriminator attribute
type FacetResult interface{}

func mapDiscriminatorFacetResult(input interface{}) (FacetResult, error) {
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
	case "filter":
		new := FilteredFacetResult{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "range":
		new := RangeFacetResult{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "terms":
		new := TermFacetResult{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// ProductUpdateAction uses action as discriminator attribute
type ProductUpdateAction interface{}

func mapDiscriminatorProductUpdateAction(input interface{}) (ProductUpdateAction, error) {
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
		new := ProductAddAssetAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addExternalImage":
		new := ProductAddExternalImageAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addPrice":
		new := ProductAddPriceAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addToCategory":
		new := ProductAddToCategoryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addVariant":
		new := ProductAddVariantAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAssetName":
		new := ProductChangeAssetNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAssetOrder":
		new := ProductChangeAssetOrderAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeMasterVariant":
		new := ProductChangeMasterVariantAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := ProductChangeNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changePrice":
		new := ProductChangePriceAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeSlug":
		new := ProductChangeSlugAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "legacySetSku":
		new := ProductLegacySetSkuAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "moveImageToPosition":
		new := ProductMoveImageToPositionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "publish":
		new := ProductPublishAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeAsset":
		new := ProductRemoveAssetAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeFromCategory":
		new := ProductRemoveFromCategoryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeImage":
		new := ProductRemoveImageAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removePrice":
		new := ProductRemovePriceAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeVariant":
		new := ProductRemoveVariantAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "revertStagedChanges":
		new := ProductRevertStagedChangesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "revertStagedVariantChanges":
		new := ProductRevertStagedVariantChangesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetCustomField":
		new := ProductSetAssetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetCustomType":
		new := ProductSetAssetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetDescription":
		new := ProductSetAssetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetKey":
		new := ProductSetAssetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetSources":
		new := ProductSetAssetSourcesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetTags":
		new := ProductSetAssetTagsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAttribute":
		new := ProductSetAttributeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAttributeInAllVariants":
		new := ProductSetAttributeInAllVariantsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCategoryOrderHint":
		new := ProductSetCategoryOrderHintAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := ProductSetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDiscountedPrice":
		new := ProductSetDiscountedPriceAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setImageLabel":
		new := ProductSetImageLabelAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := ProductSetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMetaDescription":
		new := ProductSetMetaDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMetaKeywords":
		new := ProductSetMetaKeywordsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMetaTitle":
		new := ProductSetMetaTitleAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setPrices":
		new := ProductSetPricesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setProductPriceCustomField":
		new := ProductSetProductPriceCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setProductPriceCustomType":
		new := ProductSetProductPriceCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setProductVariantKey":
		new := ProductSetProductVariantKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSearchKeywords":
		new := ProductSetSearchKeywordsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSku":
		new := ProductSetSkuAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTaxCategory":
		new := ProductSetTaxCategoryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionState":
		new := ProductTransitionStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "unpublish":
		new := ProductUnpublishAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// SuggestTokenizer uses type as discriminator attribute
type SuggestTokenizer interface{}

func mapDiscriminatorSuggestTokenizer(input interface{}) (SuggestTokenizer, error) {
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
	case "custom":
		new := CustomTokenizer{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "whitespace":
		new := WhitespaceTokenizer{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Attribute is a standalone struct
type Attribute struct {
	Value interface{} `json:"value"`
	Name  string      `json:"name"`
}

// AttributeValue is a standalone struct
type AttributeValue struct{}

// CustomTokenizer implements the interface SuggestTokenizer
type CustomTokenizer struct {
	Inputs []string `json:"inputs"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomTokenizer) MarshalJSON() ([]byte, error) {
	type Alias CustomTokenizer
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "custom", Alias: (*Alias)(&obj)})
}

// FacetResultRange is a standalone struct
type FacetResultRange struct {
	Total        int     `json:"total"`
	ToStr        string  `json:"toStr"`
	To           float64 `json:"to"`
	ProductCount int     `json:"productCount,omitempty"`
	Min          float64 `json:"min"`
	Mean         float64 `json:"mean"`
	Max          float64 `json:"max"`
	FromStr      string  `json:"fromStr"`
	From         float64 `json:"from"`
	Count        int     `json:"count"`
}

// FacetResultTerm is a standalone struct
type FacetResultTerm struct {
	Term         interface{} `json:"term"`
	ProductCount int         `json:"productCount,omitempty"`
	Count        int         `json:"count"`
}

// FacetResults is a standalone struct
type FacetResults struct {
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *FacetResults) UnmarshalJSON(data []byte) error {
	type Alias FacetResults
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	return nil
}

// FilteredFacetResult implements the interface FacetResult
type FilteredFacetResult struct {
	ProductCount int `json:"productCount,omitempty"`
	Count        int `json:"count"`
}

// MarshalJSON override to set the discriminator value
func (obj FilteredFacetResult) MarshalJSON() ([]byte, error) {
	type Alias FilteredFacetResult
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "filter", Alias: (*Alias)(&obj)})
}

// Product is of type LoggedResource
type Product struct {
	Version                int                     `json:"version"`
	LastModifiedAt         time.Time               `json:"lastModifiedAt"`
	ID                     string                  `json:"id"`
	CreatedAt              time.Time               `json:"createdAt"`
	LastModifiedBy         *LastModifiedBy         `json:"lastModifiedBy,omitempty"`
	CreatedBy              *CreatedBy              `json:"createdBy,omitempty"`
	TaxCategory            *TaxCategoryReference   `json:"taxCategory,omitempty"`
	State                  *StateReference         `json:"state,omitempty"`
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
	ProductType            *ProductTypeReference   `json:"productType"`
	MasterData             *ProductCatalogData     `json:"masterData"`
	Key                    string                  `json:"key,omitempty"`
}

// ProductAddAssetAction implements the interface ProductUpdateAction
type ProductAddAssetAction struct {
	VariantID int         `json:"variantId,omitempty"`
	Staged    bool        `json:"staged"`
	SKU       string      `json:"sku,omitempty"`
	Position  float64     `json:"position,omitempty"`
	Asset     *AssetDraft `json:"asset"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddAssetAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAsset", Alias: (*Alias)(&obj)})
}

// ProductAddExternalImageAction implements the interface ProductUpdateAction
type ProductAddExternalImageAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged"`
	SKU       string `json:"sku,omitempty"`
	Image     *Image `json:"image"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddExternalImageAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddExternalImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addExternalImage", Alias: (*Alias)(&obj)})
}

// ProductAddPriceAction implements the interface ProductUpdateAction
type ProductAddPriceAction struct {
	VariantID int         `json:"variantId,omitempty"`
	Staged    bool        `json:"staged"`
	SKU       string      `json:"sku,omitempty"`
	Price     *PriceDraft `json:"price"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddPriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPrice", Alias: (*Alias)(&obj)})
}

// ProductAddToCategoryAction implements the interface ProductUpdateAction
type ProductAddToCategoryAction struct {
	Staged    bool                        `json:"staged"`
	OrderHint string                      `json:"orderHint,omitempty"`
	Category  *CategoryResourceIdentifier `json:"category"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddToCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddToCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addToCategory", Alias: (*Alias)(&obj)})
}

// ProductAddVariantAction implements the interface ProductUpdateAction
type ProductAddVariantAction struct {
	Staged     bool         `json:"staged"`
	SKU        string       `json:"sku,omitempty"`
	Prices     []PriceDraft `json:"prices,omitempty"`
	Key        string       `json:"key,omitempty"`
	Images     []Image      `json:"images,omitempty"`
	Attributes []Attribute  `json:"attributes,omitempty"`
	Assets     []Asset      `json:"assets,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addVariant", Alias: (*Alias)(&obj)})
}

// ProductCatalogData is a standalone struct
type ProductCatalogData struct {
	Staged           *ProductData `json:"staged"`
	Published        bool         `json:"published"`
	HasStagedChanges bool         `json:"hasStagedChanges"`
	Current          *ProductData `json:"current"`
}

// ProductChangeAssetNameAction implements the interface ProductUpdateAction
type ProductChangeAssetNameAction struct {
	VariantID int              `json:"variantId,omitempty"`
	Staged    bool             `json:"staged"`
	SKU       string           `json:"sku,omitempty"`
	Name      *LocalizedString `json:"name"`
	AssetKey  string           `json:"assetKey,omitempty"`
	AssetID   string           `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangeAssetNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeAssetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetName", Alias: (*Alias)(&obj)})
}

// ProductChangeAssetOrderAction implements the interface ProductUpdateAction
type ProductChangeAssetOrderAction struct {
	VariantID  int      `json:"variantId,omitempty"`
	Staged     bool     `json:"staged"`
	SKU        string   `json:"sku,omitempty"`
	AssetOrder []string `json:"assetOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangeAssetOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeAssetOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetOrder", Alias: (*Alias)(&obj)})
}

// ProductChangeMasterVariantAction implements the interface ProductUpdateAction
type ProductChangeMasterVariantAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged"`
	SKU       string `json:"sku,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangeMasterVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeMasterVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMasterVariant", Alias: (*Alias)(&obj)})
}

// ProductChangeNameAction implements the interface ProductUpdateAction
type ProductChangeNameAction struct {
	Staged bool             `json:"staged"`
	Name   *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ProductChangePriceAction implements the interface ProductUpdateAction
type ProductChangePriceAction struct {
	Staged  bool        `json:"staged"`
	PriceID string      `json:"priceId"`
	Price   *PriceDraft `json:"price"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangePriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangePriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePrice", Alias: (*Alias)(&obj)})
}

// ProductChangeSlugAction implements the interface ProductUpdateAction
type ProductChangeSlugAction struct {
	Staged bool             `json:"staged"`
	Slug   *LocalizedString `json:"slug"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangeSlugAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSlug", Alias: (*Alias)(&obj)})
}

// ProductData is a standalone struct
type ProductData struct {
	Variants           []ProductVariant    `json:"variants"`
	Slug               *LocalizedString    `json:"slug"`
	SearchKeywords     *SearchKeywords     `json:"searchKeywords"`
	Name               *LocalizedString    `json:"name"`
	MetaTitle          *LocalizedString    `json:"metaTitle,omitempty"`
	MetaKeywords       *LocalizedString    `json:"metaKeywords,omitempty"`
	MetaDescription    *LocalizedString    `json:"metaDescription,omitempty"`
	MasterVariant      *ProductVariant     `json:"masterVariant"`
	Description        *LocalizedString    `json:"description,omitempty"`
	CategoryOrderHints *CategoryOrderHints `json:"categoryOrderHints,omitempty"`
	Categories         []CategoryReference `json:"categories"`
}

// ProductDraft is a standalone struct
type ProductDraft struct {
	Variants           []ProductVariantDraft          `json:"variants,omitempty"`
	TaxCategory        *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	State              *StateResourceIdentifier       `json:"state,omitempty"`
	Slug               *LocalizedString               `json:"slug"`
	SearchKeywords     *SearchKeywords                `json:"searchKeywords,omitempty"`
	Publish            bool                           `json:"publish"`
	ProductType        *ProductTypeResourceIdentifier `json:"productType"`
	Name               *LocalizedString               `json:"name"`
	MetaTitle          *LocalizedString               `json:"metaTitle,omitempty"`
	MetaKeywords       *LocalizedString               `json:"metaKeywords,omitempty"`
	MetaDescription    *LocalizedString               `json:"metaDescription,omitempty"`
	MasterVariant      *ProductVariantDraft           `json:"masterVariant,omitempty"`
	Key                string                         `json:"key,omitempty"`
	Description        *LocalizedString               `json:"description,omitempty"`
	CategoryOrderHints *CategoryOrderHints            `json:"categoryOrderHints,omitempty"`
	Categories         []CategoryResourceIdentifier   `json:"categories,omitempty"`
}

// ProductLegacySetSkuAction implements the interface ProductUpdateAction
type ProductLegacySetSkuAction struct {
	VariantID int    `json:"variantId"`
	SKU       string `json:"sku,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductLegacySetSkuAction) MarshalJSON() ([]byte, error) {
	type Alias ProductLegacySetSkuAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "legacySetSku", Alias: (*Alias)(&obj)})
}

// ProductMoveImageToPositionAction implements the interface ProductUpdateAction
type ProductMoveImageToPositionAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged"`
	SKU       string `json:"sku,omitempty"`
	Position  int    `json:"position"`
	ImageURL  string `json:"imageUrl"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductMoveImageToPositionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductMoveImageToPositionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "moveImageToPosition", Alias: (*Alias)(&obj)})
}

// ProductPagedQueryResponse is a standalone struct
type ProductPagedQueryResponse struct {
	Total   int       `json:"total,omitempty"`
	Results []Product `json:"results"`
	Offset  int       `json:"offset"`
	Count   int       `json:"count"`
}

// ProductProjection is of type BaseResource
type ProductProjection struct {
	Version                int                     `json:"version"`
	LastModifiedAt         time.Time               `json:"lastModifiedAt"`
	ID                     string                  `json:"id"`
	CreatedAt              time.Time               `json:"createdAt"`
	Variants               []ProductVariant        `json:"variants"`
	TaxCategory            *TaxCategoryReference   `json:"taxCategory,omitempty"`
	State                  *StateReference         `json:"state,omitempty"`
	Slug                   *LocalizedString        `json:"slug"`
	SearchKeywords         *SearchKeywords         `json:"searchKeywords,omitempty"`
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
	Published              bool                    `json:"published"`
	ProductType            *ProductTypeReference   `json:"productType"`
	Name                   *LocalizedString        `json:"name"`
	MetaTitle              *LocalizedString        `json:"metaTitle,omitempty"`
	MetaKeywords           *LocalizedString        `json:"metaKeywords,omitempty"`
	MetaDescription        *LocalizedString        `json:"metaDescription,omitempty"`
	MasterVariant          *ProductVariant         `json:"masterVariant"`
	Key                    string                  `json:"key,omitempty"`
	HasStagedChanges       bool                    `json:"hasStagedChanges"`
	Description            *LocalizedString        `json:"description,omitempty"`
	CategoryOrderHints     *CategoryOrderHints     `json:"categoryOrderHints,omitempty"`
	Categories             []CategoryReference     `json:"categories"`
}

// ProductProjectionPagedQueryResponse is a standalone struct
type ProductProjectionPagedQueryResponse struct {
	Total   int                 `json:"total,omitempty"`
	Results []ProductProjection `json:"results"`
	Offset  int                 `json:"offset"`
	Count   int                 `json:"count"`
}

// ProductProjectionPagedSearchResponse is a standalone struct
type ProductProjectionPagedSearchResponse struct {
	Total   int                 `json:"total,omitempty"`
	Results []ProductProjection `json:"results"`
	Offset  int                 `json:"offset"`
	Facets  *FacetResults       `json:"facets"`
	Count   int                 `json:"count"`
}

// ProductPublishAction implements the interface ProductUpdateAction
type ProductPublishAction struct {
	Scope ProductPublishScope `json:"scope,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductPublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "publish", Alias: (*Alias)(&obj)})
}

// ProductReference implements the interface Reference
type ProductReference struct {
	ID  string   `json:"id"`
	Obj *Product `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductReference) MarshalJSON() ([]byte, error) {
	type Alias ProductReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "product", Alias: (*Alias)(&obj)})
}

// ProductRemoveAssetAction implements the interface ProductUpdateAction
type ProductRemoveAssetAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged"`
	SKU       string `json:"sku,omitempty"`
	AssetKey  string `json:"assetKey,omitempty"`
	AssetID   string `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemoveAssetAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAsset", Alias: (*Alias)(&obj)})
}

// ProductRemoveFromCategoryAction implements the interface ProductUpdateAction
type ProductRemoveFromCategoryAction struct {
	Staged   bool                        `json:"staged"`
	Category *CategoryResourceIdentifier `json:"category"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemoveFromCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveFromCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeFromCategory", Alias: (*Alias)(&obj)})
}

// ProductRemoveImageAction implements the interface ProductUpdateAction
type ProductRemoveImageAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged"`
	SKU       string `json:"sku,omitempty"`
	ImageURL  string `json:"imageUrl"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemoveImageAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeImage", Alias: (*Alias)(&obj)})
}

// ProductRemovePriceAction implements the interface ProductUpdateAction
type ProductRemovePriceAction struct {
	Staged  bool   `json:"staged"`
	PriceID string `json:"priceId"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemovePriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemovePriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePrice", Alias: (*Alias)(&obj)})
}

// ProductRemoveVariantAction implements the interface ProductUpdateAction
type ProductRemoveVariantAction struct {
	Staged bool   `json:"staged"`
	SKU    string `json:"sku,omitempty"`
	ID     int    `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemoveVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeVariant", Alias: (*Alias)(&obj)})
}

// ProductResourceIdentifier implements the interface ResourceIdentifier
type ProductResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ProductResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "product", Alias: (*Alias)(&obj)})
}

// ProductRevertStagedChangesAction implements the interface ProductUpdateAction
type ProductRevertStagedChangesAction struct{}

// MarshalJSON override to set the discriminator value
func (obj ProductRevertStagedChangesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertStagedChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "revertStagedChanges", Alias: (*Alias)(&obj)})
}

// ProductRevertStagedVariantChangesAction implements the interface ProductUpdateAction
type ProductRevertStagedVariantChangesAction struct {
	VariantID int `json:"variantId"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRevertStagedVariantChangesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertStagedVariantChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "revertStagedVariantChanges", Alias: (*Alias)(&obj)})
}

// ProductSetAssetCustomFieldAction implements the interface ProductUpdateAction
type ProductSetAssetCustomFieldAction struct {
	VariantID int         `json:"variantId,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	Staged    bool        `json:"staged"`
	SKU       string      `json:"sku,omitempty"`
	Name      string      `json:"name"`
	AssetKey  string      `json:"assetKey,omitempty"`
	AssetID   string      `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomField", Alias: (*Alias)(&obj)})
}

// ProductSetAssetCustomTypeAction implements the interface ProductUpdateAction
type ProductSetAssetCustomTypeAction struct {
	VariantID int                     `json:"variantId,omitempty"`
	Type      *TypeResourceIdentifier `json:"type,omitempty"`
	Staged    bool                    `json:"staged"`
	SKU       string                  `json:"sku,omitempty"`
	Fields    interface{}             `json:"fields,omitempty"`
	AssetKey  string                  `json:"assetKey,omitempty"`
	AssetID   string                  `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomType", Alias: (*Alias)(&obj)})
}

// ProductSetAssetDescriptionAction implements the interface ProductUpdateAction
type ProductSetAssetDescriptionAction struct {
	VariantID   int              `json:"variantId,omitempty"`
	Staged      bool             `json:"staged"`
	SKU         string           `json:"sku,omitempty"`
	Description *LocalizedString `json:"description,omitempty"`
	AssetKey    string           `json:"assetKey,omitempty"`
	AssetID     string           `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetDescription", Alias: (*Alias)(&obj)})
}

// ProductSetAssetKeyAction implements the interface ProductUpdateAction
type ProductSetAssetKeyAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged"`
	SKU       string `json:"sku,omitempty"`
	AssetKey  string `json:"assetKey,omitempty"`
	AssetID   string `json:"assetId"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetKey", Alias: (*Alias)(&obj)})
}

// ProductSetAssetSourcesAction implements the interface ProductUpdateAction
type ProductSetAssetSourcesAction struct {
	VariantID int           `json:"variantId,omitempty"`
	Staged    bool          `json:"staged"`
	Sources   []AssetSource `json:"sources"`
	SKU       string        `json:"sku,omitempty"`
	AssetKey  string        `json:"assetKey,omitempty"`
	AssetID   string        `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetSourcesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetSourcesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetSources", Alias: (*Alias)(&obj)})
}

// ProductSetAssetTagsAction implements the interface ProductUpdateAction
type ProductSetAssetTagsAction struct {
	VariantID int      `json:"variantId,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	Staged    bool     `json:"staged"`
	SKU       string   `json:"sku,omitempty"`
	AssetKey  string   `json:"assetKey,omitempty"`
	AssetID   string   `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetTagsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetTagsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetTags", Alias: (*Alias)(&obj)})
}

// ProductSetAttributeAction implements the interface ProductUpdateAction
type ProductSetAttributeAction struct {
	VariantID int         `json:"variantId,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	Staged    bool        `json:"staged"`
	SKU       string      `json:"sku,omitempty"`
	Name      string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAttributeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAttributeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttribute", Alias: (*Alias)(&obj)})
}

// ProductSetAttributeInAllVariantsAction implements the interface ProductUpdateAction
type ProductSetAttributeInAllVariantsAction struct {
	Value  interface{} `json:"value,omitempty"`
	Staged bool        `json:"staged"`
	Name   string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAttributeInAllVariantsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAttributeInAllVariantsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttributeInAllVariants", Alias: (*Alias)(&obj)})
}

// ProductSetCategoryOrderHintAction implements the interface ProductUpdateAction
type ProductSetCategoryOrderHintAction struct {
	Staged     bool   `json:"staged"`
	OrderHint  string `json:"orderHint,omitempty"`
	CategoryID string `json:"categoryId"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetCategoryOrderHintAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetCategoryOrderHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCategoryOrderHint", Alias: (*Alias)(&obj)})
}

// ProductSetDescriptionAction implements the interface ProductUpdateAction
type ProductSetDescriptionAction struct {
	Staged      bool             `json:"staged"`
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// ProductSetDiscountedPriceAction implements the interface ProductUpdateAction
type ProductSetDiscountedPriceAction struct {
	Staged     bool             `json:"staged"`
	PriceID    string           `json:"priceId"`
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetDiscountedPriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetDiscountedPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDiscountedPrice", Alias: (*Alias)(&obj)})
}

// ProductSetImageLabelAction implements the interface ProductUpdateAction
type ProductSetImageLabelAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged"`
	SKU       string `json:"sku,omitempty"`
	Label     string `json:"label,omitempty"`
	ImageURL  string `json:"imageUrl"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetImageLabelAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetImageLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setImageLabel", Alias: (*Alias)(&obj)})
}

// ProductSetKeyAction implements the interface ProductUpdateAction
type ProductSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ProductSetMetaDescriptionAction implements the interface ProductUpdateAction
type ProductSetMetaDescriptionAction struct {
	Staged          bool             `json:"staged"`
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetMetaDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaDescription", Alias: (*Alias)(&obj)})
}

// ProductSetMetaKeywordsAction implements the interface ProductUpdateAction
type ProductSetMetaKeywordsAction struct {
	Staged       bool             `json:"staged"`
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetMetaKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaKeywords", Alias: (*Alias)(&obj)})
}

// ProductSetMetaTitleAction implements the interface ProductUpdateAction
type ProductSetMetaTitleAction struct {
	Staged    bool             `json:"staged"`
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetMetaTitleAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaTitle", Alias: (*Alias)(&obj)})
}

// ProductSetPricesAction implements the interface ProductUpdateAction
type ProductSetPricesAction struct {
	VariantID int          `json:"variantId,omitempty"`
	Staged    bool         `json:"staged"`
	SKU       string       `json:"sku,omitempty"`
	Prices    []PriceDraft `json:"prices"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetPricesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetPricesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPrices", Alias: (*Alias)(&obj)})
}

// ProductSetProductPriceCustomFieldAction implements the interface ProductUpdateAction
type ProductSetProductPriceCustomFieldAction struct {
	Value   interface{} `json:"value,omitempty"`
	Staged  bool        `json:"staged"`
	PriceID string      `json:"priceId"`
	Name    string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetProductPriceCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductPriceCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductPriceCustomField", Alias: (*Alias)(&obj)})
}

// ProductSetProductPriceCustomTypeAction implements the interface ProductUpdateAction
type ProductSetProductPriceCustomTypeAction struct {
	Type    *TypeResourceIdentifier `json:"type,omitempty"`
	Staged  bool                    `json:"staged"`
	PriceID string                  `json:"priceId"`
	Fields  *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetProductPriceCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductPriceCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductPriceCustomType", Alias: (*Alias)(&obj)})
}

// ProductSetProductVariantKeyAction implements the interface ProductUpdateAction
type ProductSetProductVariantKeyAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged"`
	SKU       string `json:"sku,omitempty"`
	Key       string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetProductVariantKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductVariantKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductVariantKey", Alias: (*Alias)(&obj)})
}

// ProductSetSearchKeywordsAction implements the interface ProductUpdateAction
type ProductSetSearchKeywordsAction struct {
	Staged         bool            `json:"staged"`
	SearchKeywords *SearchKeywords `json:"searchKeywords"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetSearchKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetSearchKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSearchKeywords", Alias: (*Alias)(&obj)})
}

// ProductSetSkuAction implements the interface ProductUpdateAction
type ProductSetSkuAction struct {
	VariantID int    `json:"variantId"`
	Staged    bool   `json:"staged"`
	SKU       string `json:"sku,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetSkuAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetSkuAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSku", Alias: (*Alias)(&obj)})
}

// ProductSetTaxCategoryAction implements the interface ProductUpdateAction
type ProductSetTaxCategoryAction struct {
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetTaxCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetTaxCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTaxCategory", Alias: (*Alias)(&obj)})
}

// ProductTransitionStateAction implements the interface ProductUpdateAction
type ProductTransitionStateAction struct {
	State *StateResourceIdentifier `json:"state,omitempty"`
	Force bool                     `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

// ProductUnpublishAction implements the interface ProductUpdateAction
type ProductUnpublishAction struct{}

// MarshalJSON override to set the discriminator value
func (obj ProductUnpublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductUnpublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "unpublish", Alias: (*Alias)(&obj)})
}

// ProductUpdate is a standalone struct
type ProductUpdate struct {
	Version int                   `json:"version"`
	Actions []ProductUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProductUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorProductUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// ProductVariant is a standalone struct
type ProductVariant struct {
	SKU                   string                      `json:"sku,omitempty"`
	ScopedPriceDiscounted bool                        `json:"scopedPriceDiscounted"`
	ScopedPrice           *ScopedPrice                `json:"scopedPrice,omitempty"`
	Prices                []Price                     `json:"prices,omitempty"`
	Price                 *Price                      `json:"price,omitempty"`
	Key                   string                      `json:"key,omitempty"`
	IsMatchingVariant     bool                        `json:"isMatchingVariant"`
	Images                []Image                     `json:"images,omitempty"`
	ID                    int                         `json:"id"`
	Availability          *ProductVariantAvailability `json:"availability,omitempty"`
	Attributes            []Attribute                 `json:"attributes,omitempty"`
	Assets                []Asset                     `json:"assets,omitempty"`
}

// ProductVariantAvailability is a standalone struct
type ProductVariantAvailability struct {
	RestockableInDays int                                   `json:"restockableInDays,omitempty"`
	IsOnStock         bool                                  `json:"isOnStock"`
	Channels          *ProductVariantChannelAvailabilityMap `json:"channels,omitempty"`
	AvailableQuantity int                                   `json:"availableQuantity,omitempty"`
}

// ProductVariantChannelAvailability is a standalone struct
type ProductVariantChannelAvailability struct {
	RestockableInDays int  `json:"restockableInDays,omitempty"`
	IsOnStock         bool `json:"isOnStock"`
	AvailableQuantity int  `json:"availableQuantity,omitempty"`
}

// ProductVariantChannelAvailabilityMap is a standalone struct
type ProductVariantChannelAvailabilityMap struct {
}

// ProductVariantDraft is a standalone struct
type ProductVariantDraft struct {
	SKU        string       `json:"sku,omitempty"`
	Prices     []PriceDraft `json:"prices,omitempty"`
	Key        string       `json:"key,omitempty"`
	Images     []Image      `json:"images,omitempty"`
	Attributes []Attribute  `json:"attributes,omitempty"`
	Assets     []AssetDraft `json:"assets,omitempty"`
}

// RangeFacetResult implements the interface FacetResult
type RangeFacetResult struct {
	Ranges []FacetResultRange `json:"ranges"`
}

// MarshalJSON override to set the discriminator value
func (obj RangeFacetResult) MarshalJSON() ([]byte, error) {
	type Alias RangeFacetResult
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "range", Alias: (*Alias)(&obj)})
}

// SearchKeyword is a standalone struct
type SearchKeyword struct {
	Text             string           `json:"text"`
	SuggestTokenizer SuggestTokenizer `json:"suggestTokenizer,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SearchKeyword) UnmarshalJSON(data []byte) error {
	type Alias SearchKeyword
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.SuggestTokenizer != nil {
		var err error
		obj.SuggestTokenizer, err = mapDiscriminatorSuggestTokenizer(obj.SuggestTokenizer)
		if err != nil {
			return err
		}
	}

	return nil
}

// SearchKeywords is a standalone struct
type SearchKeywords struct {
}

// Suggestion is a standalone struct
type Suggestion struct {
	Text string `json:"text"`
}

// SuggestionResult is a standalone struct
type SuggestionResult struct {
}

// TermFacetResult implements the interface FacetResult
type TermFacetResult struct {
	Total    int                 `json:"total"`
	Terms    []FacetResultTerm   `json:"terms"`
	Other    int                 `json:"other"`
	Missing  int                 `json:"missing"`
	DataType TermFacetResultType `json:"dataType"`
}

// MarshalJSON override to set the discriminator value
func (obj TermFacetResult) MarshalJSON() ([]byte, error) {
	type Alias TermFacetResult
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "terms", Alias: (*Alias)(&obj)})
}

// WhitespaceTokenizer implements the interface SuggestTokenizer
type WhitespaceTokenizer struct{}

// MarshalJSON override to set the discriminator value
func (obj WhitespaceTokenizer) MarshalJSON() ([]byte, error) {
	type Alias WhitespaceTokenizer
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "whitespace", Alias: (*Alias)(&obj)})
}
