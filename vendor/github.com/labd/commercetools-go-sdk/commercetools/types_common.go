// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// MoneyType is an enum type
type MoneyType string

// Enum values for MoneyType
const (
	MoneyTypeCentPrecision MoneyType = "centPrecision"
	MoneyTypeHighPrecision MoneyType = "highPrecision"
)

// ReferenceTypeID is an enum type
type ReferenceTypeID string

// Enum values for ReferenceTypeID
const (
	ReferenceTypeIDCart             ReferenceTypeID = "cart"
	ReferenceTypeIDCartDiscount     ReferenceTypeID = "cart-discount"
	ReferenceTypeIDCategory         ReferenceTypeID = "category"
	ReferenceTypeIDChannel          ReferenceTypeID = "channel"
	ReferenceTypeIDCustomer         ReferenceTypeID = "customer"
	ReferenceTypeIDCustomerGroup    ReferenceTypeID = "customer-group"
	ReferenceTypeIDDiscountCode     ReferenceTypeID = "discount-code"
	ReferenceTypeIDKeyValueDocument ReferenceTypeID = "key-value-document"
	ReferenceTypeIDPayment          ReferenceTypeID = "payment"
	ReferenceTypeIDProduct          ReferenceTypeID = "product"
	ReferenceTypeIDProductType      ReferenceTypeID = "product-type"
	ReferenceTypeIDProductDiscount  ReferenceTypeID = "product-discount"
	ReferenceTypeIDOrder            ReferenceTypeID = "order"
	ReferenceTypeIDReview           ReferenceTypeID = "review"
	ReferenceTypeIDShoppingList     ReferenceTypeID = "shopping-list"
	ReferenceTypeIDShippingMethod   ReferenceTypeID = "shipping-method"
	ReferenceTypeIDState            ReferenceTypeID = "state"
	ReferenceTypeIDStore            ReferenceTypeID = "store"
	ReferenceTypeIDTaxCategory      ReferenceTypeID = "tax-category"
	ReferenceTypeIDType             ReferenceTypeID = "type"
	ReferenceTypeIDZone             ReferenceTypeID = "zone"
	ReferenceTypeIDInventoryEntry   ReferenceTypeID = "inventory-entry"
	ReferenceTypeIDOrderEdit        ReferenceTypeID = "order-edit"
)

// CountryCode is of type string
type CountryCode string

// CurrencyCode is of type string
type CurrencyCode string

// Locale is of type string
type Locale string

// LocalizedString is a map
type LocalizedString map[string]string

// GeoJSON uses type as discriminator attribute
type GeoJSON interface{}

func mapDiscriminatorGeoJSON(input interface{}) (GeoJSON, error) {
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
	case "Point":
		new := GeoJSONPoint{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// KeyReference uses typeId as discriminator attribute
type KeyReference interface{}

func mapDiscriminatorKeyReference(input interface{}) (KeyReference, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "store":
		new := StoreKeyReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Reference uses typeId as discriminator attribute
type Reference interface{}

func mapDiscriminatorReference(input interface{}) (Reference, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "cart-discount":
		new := CartDiscountReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "cart":
		new := CartReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "category":
		new := CategoryReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "channel":
		new := ChannelReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "key-value-document":
		new := CustomObjectReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "customer-group":
		new := CustomerGroupReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "customer":
		new := CustomerReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "discount-code":
		new := DiscountCodeReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "inventory-entry":
		new := InventoryEntryReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "order-edit":
		new := OrderEditReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "order":
		new := OrderReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "payment":
		new := PaymentReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product-discount":
		new := ProductDiscountReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product":
		new := ProductReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product-type":
		new := ProductTypeReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "review":
		new := ReviewReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "shipping-method":
		new := ShippingMethodReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "shopping-list":
		new := ShoppingListReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "state":
		new := StateReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "store":
		new := StoreReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "tax-category":
		new := TaxCategoryReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "type":
		new := TypeReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "zone":
		new := ZoneReference{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// ResourceIdentifier uses typeId as discriminator attribute
type ResourceIdentifier interface{}

func mapDiscriminatorResourceIdentifier(input interface{}) (ResourceIdentifier, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["typeId"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'typeId'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "cart-discount":
		new := CartDiscountResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "cart":
		new := CartResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "category":
		new := CategoryResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "channel":
		new := ChannelResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "customer-group":
		new := CustomerGroupResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "customer":
		new := CustomerResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "discount-code":
		new := DiscountCodeResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "inventory-entry":
		new := InventoryEntryResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "order-edit":
		new := OrderEditResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "order":
		new := OrderResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "payment":
		new := PaymentResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product-discount":
		new := ProductDiscountResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product":
		new := ProductResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "product-type":
		new := ProductTypeResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "review":
		new := ReviewResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "shipping-method":
		new := ShippingMethodResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "shopping-list":
		new := ShoppingListResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "state":
		new := StateResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "store":
		new := StoreResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "tax-category":
		new := TaxCategoryResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "type":
		new := TypeResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "zone":
		new := ZoneResourceIdentifier{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// TypedMoney uses type as discriminator attribute
type TypedMoney interface{}

func mapDiscriminatorTypedMoney(input interface{}) (TypedMoney, error) {
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
	case "centPrecision":
		new := CentPrecisionMoney{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "highPrecision":
		new := HighPrecisionMoney{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Address is a standalone struct
type Address struct {
	Title                 string      `json:"title,omitempty"`
	StreetNumber          string      `json:"streetNumber,omitempty"`
	StreetName            string      `json:"streetName,omitempty"`
	State                 string      `json:"state,omitempty"`
	Salutation            string      `json:"salutation,omitempty"`
	Region                string      `json:"region,omitempty"`
	PostalCode            string      `json:"postalCode,omitempty"`
	Phone                 string      `json:"phone,omitempty"`
	POBox                 string      `json:"pOBox,omitempty"`
	Mobile                string      `json:"mobile,omitempty"`
	LastName              string      `json:"lastName,omitempty"`
	Key                   string      `json:"key,omitempty"`
	ID                    string      `json:"id,omitempty"`
	FirstName             string      `json:"firstName,omitempty"`
	Fax                   string      `json:"fax,omitempty"`
	ExternalID            string      `json:"externalId,omitempty"`
	Email                 string      `json:"email,omitempty"`
	Department            string      `json:"department,omitempty"`
	Country               CountryCode `json:"country"`
	Company               string      `json:"company,omitempty"`
	City                  string      `json:"city,omitempty"`
	Building              string      `json:"building,omitempty"`
	Apartment             string      `json:"apartment,omitempty"`
	AdditionalStreetInfo  string      `json:"additionalStreetInfo,omitempty"`
	AdditionalAddressInfo string      `json:"additionalAddressInfo,omitempty"`
}

// Asset is a standalone struct
type Asset struct {
	Tags        []string         `json:"tags,omitempty"`
	Sources     []AssetSource    `json:"sources"`
	Name        *LocalizedString `json:"name"`
	Key         string           `json:"key,omitempty"`
	ID          string           `json:"id"`
	Description *LocalizedString `json:"description,omitempty"`
	Custom      *CustomFields    `json:"custom,omitempty"`
}

// AssetDimensions is a standalone struct
type AssetDimensions struct {
	W float64 `json:"w"`
	H float64 `json:"h"`
}

// AssetDraft is a standalone struct
type AssetDraft struct {
	Tags        []string           `json:"tags,omitempty"`
	Sources     []AssetSource      `json:"sources"`
	Name        *LocalizedString   `json:"name"`
	Key         string             `json:"key,omitempty"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
}

// AssetSource is a standalone struct
type AssetSource struct {
	URI         string           `json:"uri"`
	Key         string           `json:"key,omitempty"`
	Dimensions  *AssetDimensions `json:"dimensions,omitempty"`
	ContentType string           `json:"contentType,omitempty"`
}

// BaseResource is a standalone struct
type BaseResource struct {
	Version        int       `json:"version"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
}

// CentPrecisionMoney implements the interface TypedMoney
type CentPrecisionMoney struct {
	CurrencyCode   CurrencyCode `json:"currencyCode"`
	CentAmount     int          `json:"centAmount"`
	FractionDigits float64      `json:"fractionDigits"`
}

// MarshalJSON override to set the discriminator value
func (obj CentPrecisionMoney) MarshalJSON() ([]byte, error) {
	type Alias CentPrecisionMoney
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "centPrecision", Alias: (*Alias)(&obj)})
}

// ClientLogging is a standalone struct
type ClientLogging struct {
	ExternalUserID string             `json:"externalUserId,omitempty"`
	Customer       *CustomerReference `json:"customer,omitempty"`
	ClientID       string             `json:"clientId,omitempty"`
	AnonymousID    string             `json:"anonymousId,omitempty"`
}

// CreatedBy is of type ClientLogging
type CreatedBy struct {
	ExternalUserID string             `json:"externalUserId,omitempty"`
	Customer       *CustomerReference `json:"customer,omitempty"`
	ClientID       string             `json:"clientId,omitempty"`
	AnonymousID    string             `json:"anonymousId,omitempty"`
}

// DiscountedPrice is a standalone struct
type DiscountedPrice struct {
	Value    *Money                    `json:"value"`
	Discount *ProductDiscountReference `json:"discount"`
}

// GeoJSONPoint implements the interface GeoJSON
type GeoJSONPoint struct {
	Coordinates []float64 `json:"coordinates"`
}

// MarshalJSON override to set the discriminator value
func (obj GeoJSONPoint) MarshalJSON() ([]byte, error) {
	type Alias GeoJSONPoint
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "Point", Alias: (*Alias)(&obj)})
}

// HighPrecisionMoney implements the interface TypedMoney
type HighPrecisionMoney struct {
	CurrencyCode   CurrencyCode `json:"currencyCode"`
	CentAmount     int          `json:"centAmount"`
	FractionDigits float64      `json:"fractionDigits"`
	PreciseAmount  int          `json:"preciseAmount"`
}

// MarshalJSON override to set the discriminator value
func (obj HighPrecisionMoney) MarshalJSON() ([]byte, error) {
	type Alias HighPrecisionMoney
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "highPrecision", Alias: (*Alias)(&obj)})
}

// Image is a standalone struct
type Image struct {
	URL        string           `json:"url"`
	Label      string           `json:"label,omitempty"`
	Dimensions *ImageDimensions `json:"dimensions"`
}

// ImageDimensions is a standalone struct
type ImageDimensions struct {
	W float64 `json:"w"`
	H float64 `json:"h"`
}

// LastModifiedBy is of type ClientLogging
type LastModifiedBy struct {
	ExternalUserID string             `json:"externalUserId,omitempty"`
	Customer       *CustomerReference `json:"customer,omitempty"`
	ClientID       string             `json:"clientId,omitempty"`
	AnonymousID    string             `json:"anonymousId,omitempty"`
}

// LoggedResource is of type BaseResource
type LoggedResource struct {
	Version        int             `json:"version"`
	LastModifiedAt time.Time       `json:"lastModifiedAt"`
	ID             string          `json:"id"`
	CreatedAt      time.Time       `json:"createdAt"`
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	CreatedBy      *CreatedBy      `json:"createdBy,omitempty"`
}

// Money is a standalone struct
type Money struct {
	CurrencyCode CurrencyCode `json:"currencyCode"`
	CentAmount   int          `json:"centAmount"`
}

// PagedQueryResponse is a standalone struct
type PagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Results []BaseResource `json:"results"`
	Offset  int            `json:"offset"`
	Meta    interface{}    `json:"meta,omitempty"`
	Facets  *FacetResults  `json:"facets,omitempty"`
	Count   int            `json:"count"`
}

// Price is a standalone struct
type Price struct {
	Value         *Money                  `json:"value"`
	ValidUntil    *time.Time              `json:"validUntil,omitempty"`
	ValidFrom     *time.Time              `json:"validFrom,omitempty"`
	Tiers         []PriceTier             `json:"tiers,omitempty"`
	ID            string                  `json:"id,omitempty"`
	Discounted    *DiscountedPrice        `json:"discounted,omitempty"`
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	Custom        *CustomFields           `json:"custom,omitempty"`
	Country       CountryCode             `json:"country,omitempty"`
	Channel       *ChannelReference       `json:"channel,omitempty"`
}

// PriceDraft is a standalone struct
type PriceDraft struct {
	Value         *Money                           `json:"value"`
	ValidUntil    *time.Time                       `json:"validUntil,omitempty"`
	ValidFrom     *time.Time                       `json:"validFrom,omitempty"`
	Tiers         []PriceTier                      `json:"tiers,omitempty"`
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	Custom        *CustomFieldsDraft               `json:"custom,omitempty"`
	Country       CountryCode                      `json:"country,omitempty"`
	Channel       *ChannelResourceIdentifier       `json:"channel,omitempty"`
}

// PriceTier is a standalone struct
type PriceTier struct {
	Value           *Money `json:"value"`
	MinimumQuantity int    `json:"minimumQuantity"`
}

// ScopedPrice is a standalone struct
type ScopedPrice struct {
	Value         TypedMoney              `json:"value"`
	ValidUntil    *time.Time              `json:"validUntil,omitempty"`
	ValidFrom     *time.Time              `json:"validFrom,omitempty"`
	ID            string                  `json:"id"`
	Discounted    *DiscountedPrice        `json:"discounted,omitempty"`
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	Custom        *CustomFields           `json:"custom,omitempty"`
	CurrentValue  TypedMoney              `json:"currentValue"`
	Country       CountryCode             `json:"country,omitempty"`
	Channel       *ChannelReference       `json:"channel,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ScopedPrice) UnmarshalJSON(data []byte) error {
	type Alias ScopedPrice
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.CurrentValue != nil {
		var err error
		obj.CurrentValue, err = mapDiscriminatorTypedMoney(obj.CurrentValue)
		if err != nil {
			return err
		}
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorTypedMoney(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// Update is a standalone struct
type Update struct {
	Version int            `json:"version"`
	Actions []UpdateAction `json:"actions"`
}

// UpdateAction is a standalone struct
type UpdateAction struct {
	Action string `json:"action"`
}
