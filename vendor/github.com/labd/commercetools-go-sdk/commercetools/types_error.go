// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ErrorObject uses code as discriminator attribute
type ErrorObject interface{}

func mapDiscriminatorErrorObject(input interface{}) (ErrorObject, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["code"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'code'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "access_denied":
		new := AccessDeniedError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ConcurrentModification":
		new := ConcurrentModificationError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DiscountCodeNonApplicable":
		new := DiscountCodeNonApplicableError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicateAttributeValue":
		new := DuplicateAttributeValueError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicateAttributeValues":
		new := DuplicateAttributeValuesError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicateField":
		new := DuplicateFieldError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicateFieldWithConflictingResource":
		new := DuplicateFieldWithConflictingResourceError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.ConflictingResource != nil {
			new.ConflictingResource, err = mapDiscriminatorReference(new.ConflictingResource)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "DuplicatePriceScope":
		new := DuplicatePriceScopeError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DuplicateVariantValues":
		new := DuplicateVariantValuesError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "EnumValueIsUsed":
		new := EnumValueIsUsedError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ExtensionBadResponse":
		new := ExtensionBadResponseError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ExtensionNoResponse":
		new := ExtensionNoResponseError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ExtensionUpdateActionsFailed":
		new := ExtensionUpdateActionsFailedError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "insufficient_scope":
		new := InsufficientScopeError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidCredentials":
		new := InvalidCredentialsError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidCurrentPassword":
		new := InvalidCurrentPasswordError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidField":
		new := InvalidFieldError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidInput":
		new := InvalidInputError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidItemShippingDetails":
		new := InvalidItemShippingDetailsError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidJsonInput":
		new := InvalidJSONInputError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidOperation":
		new := InvalidOperationError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InvalidSubject":
		new := InvalidSubjectError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "invalid_token":
		new := InvalidTokenError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "MatchingPriceNotFound":
		new := MatchingPriceNotFoundError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "MissingTaxRateForCountry":
		new := MissingTaxRateForCountryError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "NoMatchingProductDiscountFound":
		new := NoMatchingProductDiscountFoundError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OutOfStock":
		new := OutOfStockError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PriceChanged":
		new := PriceChangedError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReferenceExists":
		new := ReferenceExistsError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "RequiredField":
		new := RequiredFieldError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ResourceNotFound":
		new := ResourceNotFoundError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ShippingMethodDoesNotMatchCart":
		new := ShippingMethodDoesNotMatchCartError{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// AccessDeniedError implements the interface ErrorObject
type AccessDeniedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj AccessDeniedError) MarshalJSON() ([]byte, error) {
	type Alias AccessDeniedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "access_denied", Alias: (*Alias)(&obj)})
}

func (obj AccessDeniedError) Error() string {
	return obj.Message
}

// ConcurrentModificationError implements the interface ErrorObject
type ConcurrentModificationError struct {
	Message        string `json:"message"`
	CurrentVersion int    `json:"currentVersion,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ConcurrentModificationError) MarshalJSON() ([]byte, error) {
	type Alias ConcurrentModificationError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ConcurrentModification", Alias: (*Alias)(&obj)})
}

func (obj ConcurrentModificationError) Error() string {
	return obj.Message
}

// DiscountCodeNonApplicableError implements the interface ErrorObject
type DiscountCodeNonApplicableError struct {
	Message           string     `json:"message"`
	ValidityCheckTime *time.Time `json:"validityCheckTime,omitempty"`
	ValidUntil        *time.Time `json:"validUntil,omitempty"`
	ValidFrom         *time.Time `json:"validFrom,omitempty"`
	Reason            string     `json:"reason,omitempty"`
	DiscountCode      string     `json:"discountCode,omitempty"`
	DicountCodeID     string     `json:"dicountCodeId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeNonApplicableError) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeNonApplicableError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DiscountCodeNonApplicable", Alias: (*Alias)(&obj)})
}

func (obj DiscountCodeNonApplicableError) Error() string {
	return obj.Message
}

// DuplicateAttributeValueError implements the interface ErrorObject
type DuplicateAttributeValueError struct {
	Message   string     `json:"message"`
	Attribute *Attribute `json:"attribute"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateAttributeValueError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValueError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateAttributeValue", Alias: (*Alias)(&obj)})
}

func (obj DuplicateAttributeValueError) Error() string {
	return obj.Message
}

// DuplicateAttributeValuesError implements the interface ErrorObject
type DuplicateAttributeValuesError struct {
	Message    string      `json:"message"`
	Attributes []Attribute `json:"attributes"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateAttributeValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateAttributeValuesError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateAttributeValues", Alias: (*Alias)(&obj)})
}

func (obj DuplicateAttributeValuesError) Error() string {
	return obj.Message
}

// DuplicateFieldError implements the interface ErrorObject
type DuplicateFieldError struct {
	Message        string      `json:"message"`
	Field          string      `json:"field,omitempty"`
	DuplicateValue interface{} `json:"duplicateValue,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateFieldError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateField", Alias: (*Alias)(&obj)})
}

func (obj DuplicateFieldError) Error() string {
	return obj.Message
}

// DuplicateFieldWithConflictingResourceError implements the interface ErrorObject
type DuplicateFieldWithConflictingResourceError struct {
	Message             string      `json:"message"`
	Field               string      `json:"field"`
	DuplicateValue      interface{} `json:"duplicateValue"`
	ConflictingResource Reference   `json:"conflictingResource"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateFieldWithConflictingResourceError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateFieldWithConflictingResourceError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateFieldWithConflictingResource", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DuplicateFieldWithConflictingResourceError) UnmarshalJSON(data []byte) error {
	type Alias DuplicateFieldWithConflictingResourceError
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ConflictingResource != nil {
		var err error
		obj.ConflictingResource, err = mapDiscriminatorReference(obj.ConflictingResource)
		if err != nil {
			return err
		}
	}

	return nil
}

func (obj DuplicateFieldWithConflictingResourceError) Error() string {
	return obj.Message
}

// DuplicatePriceScopeError implements the interface ErrorObject
type DuplicatePriceScopeError struct {
	Message           string  `json:"message"`
	ConflictingPrices []Price `json:"conflictingPrices"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicatePriceScopeError) MarshalJSON() ([]byte, error) {
	type Alias DuplicatePriceScopeError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicatePriceScope", Alias: (*Alias)(&obj)})
}

func (obj DuplicatePriceScopeError) Error() string {
	return obj.Message
}

// DuplicateVariantValuesError implements the interface ErrorObject
type DuplicateVariantValuesError struct {
	Message       string         `json:"message"`
	VariantValues *VariantValues `json:"variantValues"`
}

// MarshalJSON override to set the discriminator value
func (obj DuplicateVariantValuesError) MarshalJSON() ([]byte, error) {
	type Alias DuplicateVariantValuesError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "DuplicateVariantValues", Alias: (*Alias)(&obj)})
}

func (obj DuplicateVariantValuesError) Error() string {
	return obj.Message
}

// EnumValueIsUsedError implements the interface ErrorObject
type EnumValueIsUsedError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj EnumValueIsUsedError) MarshalJSON() ([]byte, error) {
	type Alias EnumValueIsUsedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "EnumValueIsUsed", Alias: (*Alias)(&obj)})
}

func (obj EnumValueIsUsedError) Error() string {
	return obj.Message
}

// ErrorByExtension is a standalone struct
type ErrorByExtension struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id"`
}

// ErrorResponse is a standalone struct
type ErrorResponse struct {
	StatusCode       int           `json:"statusCode"`
	Message          string        `json:"message"`
	Errors           []ErrorObject `json:"errors,omitempty"`
	ErrorDescription string        `json:"error_description,omitempty"`
	ErrorMessage     string        `json:"error,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ErrorResponse) UnmarshalJSON(data []byte) error {
	type Alias ErrorResponse
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Errors {
		var err error
		obj.Errors[i], err = mapDiscriminatorErrorObject(obj.Errors[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (obj ErrorResponse) Error() string {
	return obj.Message
}

// ExtensionBadResponseError implements the interface ErrorObject
type ExtensionBadResponseError struct {
	Message            string            `json:"message"`
	LocalizedMessage   *LocalizedString  `json:"localizedMessage,omitempty"`
	ExtensionExtraInfo interface{}       `json:"extensionExtraInfo,omitempty"`
	ErrorByExtension   *ErrorByExtension `json:"errorByExtension"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionBadResponseError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionBadResponseError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ExtensionBadResponse", Alias: (*Alias)(&obj)})
}

func (obj ExtensionBadResponseError) Error() string {
	return obj.Message
}

// ExtensionNoResponseError implements the interface ErrorObject
type ExtensionNoResponseError struct {
	Message            string            `json:"message"`
	LocalizedMessage   *LocalizedString  `json:"localizedMessage,omitempty"`
	ExtensionExtraInfo interface{}       `json:"extensionExtraInfo,omitempty"`
	ErrorByExtension   *ErrorByExtension `json:"errorByExtension"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionNoResponseError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionNoResponseError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ExtensionNoResponse", Alias: (*Alias)(&obj)})
}

func (obj ExtensionNoResponseError) Error() string {
	return obj.Message
}

// ExtensionUpdateActionsFailedError implements the interface ErrorObject
type ExtensionUpdateActionsFailedError struct {
	Message            string            `json:"message"`
	LocalizedMessage   *LocalizedString  `json:"localizedMessage,omitempty"`
	ExtensionExtraInfo interface{}       `json:"extensionExtraInfo,omitempty"`
	ErrorByExtension   *ErrorByExtension `json:"errorByExtension"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionUpdateActionsFailedError) MarshalJSON() ([]byte, error) {
	type Alias ExtensionUpdateActionsFailedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ExtensionUpdateActionsFailed", Alias: (*Alias)(&obj)})
}

func (obj ExtensionUpdateActionsFailedError) Error() string {
	return obj.Message
}

// InsufficientScopeError implements the interface ErrorObject
type InsufficientScopeError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InsufficientScopeError) MarshalJSON() ([]byte, error) {
	type Alias InsufficientScopeError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "insufficient_scope", Alias: (*Alias)(&obj)})
}

func (obj InsufficientScopeError) Error() string {
	return obj.Message
}

// InvalidCredentialsError implements the interface ErrorObject
type InvalidCredentialsError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidCredentialsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCredentialsError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidCredentials", Alias: (*Alias)(&obj)})
}

func (obj InvalidCredentialsError) Error() string {
	return obj.Message
}

// InvalidCurrentPasswordError implements the interface ErrorObject
type InvalidCurrentPasswordError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidCurrentPasswordError) MarshalJSON() ([]byte, error) {
	type Alias InvalidCurrentPasswordError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidCurrentPassword", Alias: (*Alias)(&obj)})
}

func (obj InvalidCurrentPasswordError) Error() string {
	return obj.Message
}

// InvalidFieldError implements the interface ErrorObject
type InvalidFieldError struct {
	Message       string        `json:"message"`
	InvalidValue  interface{}   `json:"invalidValue"`
	Field         string        `json:"field"`
	AllowedValues []interface{} `json:"allowedValues,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidFieldError) MarshalJSON() ([]byte, error) {
	type Alias InvalidFieldError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidField", Alias: (*Alias)(&obj)})
}

func (obj InvalidFieldError) Error() string {
	return obj.Message
}

// InvalidInputError implements the interface ErrorObject
type InvalidInputError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidInputError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidInput", Alias: (*Alias)(&obj)})
}

func (obj InvalidInputError) Error() string {
	return obj.Message
}

// InvalidItemShippingDetailsError implements the interface ErrorObject
type InvalidItemShippingDetailsError struct {
	Message string `json:"message"`
	Subject string `json:"subject"`
	ItemID  string `json:"itemId"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidItemShippingDetailsError) MarshalJSON() ([]byte, error) {
	type Alias InvalidItemShippingDetailsError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidItemShippingDetails", Alias: (*Alias)(&obj)})
}

func (obj InvalidItemShippingDetailsError) Error() string {
	return obj.Message
}

// InvalidJSONInputError implements the interface ErrorObject
type InvalidJSONInputError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidJSONInputError) MarshalJSON() ([]byte, error) {
	type Alias InvalidJSONInputError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidJsonInput", Alias: (*Alias)(&obj)})
}

func (obj InvalidJSONInputError) Error() string {
	return obj.Message
}

// InvalidOperationError implements the interface ErrorObject
type InvalidOperationError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidOperationError) MarshalJSON() ([]byte, error) {
	type Alias InvalidOperationError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidOperation", Alias: (*Alias)(&obj)})
}

func (obj InvalidOperationError) Error() string {
	return obj.Message
}

// InvalidSubjectError implements the interface ErrorObject
type InvalidSubjectError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidSubjectError) MarshalJSON() ([]byte, error) {
	type Alias InvalidSubjectError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "InvalidSubject", Alias: (*Alias)(&obj)})
}

func (obj InvalidSubjectError) Error() string {
	return obj.Message
}

// InvalidTokenError implements the interface ErrorObject
type InvalidTokenError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj InvalidTokenError) MarshalJSON() ([]byte, error) {
	type Alias InvalidTokenError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "invalid_token", Alias: (*Alias)(&obj)})
}

func (obj InvalidTokenError) Error() string {
	return obj.Message
}

// MatchingPriceNotFoundError implements the interface ErrorObject
type MatchingPriceNotFoundError struct {
	Message       string                  `json:"message"`
	VariantID     int                     `json:"variantId"`
	ProductID     string                  `json:"productId"`
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	Currency      string                  `json:"currency,omitempty"`
	Country       string                  `json:"country,omitempty"`
	Channel       *ChannelReference       `json:"channel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MatchingPriceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias MatchingPriceNotFoundError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "MatchingPriceNotFound", Alias: (*Alias)(&obj)})
}

func (obj MatchingPriceNotFoundError) Error() string {
	return obj.Message
}

// MissingTaxRateForCountryError implements the interface ErrorObject
type MissingTaxRateForCountryError struct {
	Message       string `json:"message"`
	TaxCategoryID string `json:"taxCategoryId"`
	State         string `json:"state,omitempty"`
	Country       string `json:"country,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MissingTaxRateForCountryError) MarshalJSON() ([]byte, error) {
	type Alias MissingTaxRateForCountryError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "MissingTaxRateForCountry", Alias: (*Alias)(&obj)})
}

func (obj MissingTaxRateForCountryError) Error() string {
	return obj.Message
}

// NoMatchingProductDiscountFoundError implements the interface ErrorObject
type NoMatchingProductDiscountFoundError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj NoMatchingProductDiscountFoundError) MarshalJSON() ([]byte, error) {
	type Alias NoMatchingProductDiscountFoundError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "NoMatchingProductDiscountFound", Alias: (*Alias)(&obj)})
}

func (obj NoMatchingProductDiscountFoundError) Error() string {
	return obj.Message
}

// OutOfStockError implements the interface ErrorObject
type OutOfStockError struct {
	Message   string   `json:"message"`
	Skus      []string `json:"skus"`
	LineItems []string `json:"lineItems"`
}

// MarshalJSON override to set the discriminator value
func (obj OutOfStockError) MarshalJSON() ([]byte, error) {
	type Alias OutOfStockError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "OutOfStock", Alias: (*Alias)(&obj)})
}

func (obj OutOfStockError) Error() string {
	return obj.Message
}

// PriceChangedError implements the interface ErrorObject
type PriceChangedError struct {
	Message   string   `json:"message"`
	Shipping  bool     `json:"shipping"`
	LineItems []string `json:"lineItems"`
}

// MarshalJSON override to set the discriminator value
func (obj PriceChangedError) MarshalJSON() ([]byte, error) {
	type Alias PriceChangedError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "PriceChanged", Alias: (*Alias)(&obj)})
}

func (obj PriceChangedError) Error() string {
	return obj.Message
}

// ReferenceExistsError implements the interface ErrorObject
type ReferenceExistsError struct {
	Message      string          `json:"message"`
	ReferencedBy ReferenceTypeID `json:"referencedBy,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReferenceExistsError) MarshalJSON() ([]byte, error) {
	type Alias ReferenceExistsError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ReferenceExists", Alias: (*Alias)(&obj)})
}

func (obj ReferenceExistsError) Error() string {
	return obj.Message
}

// RequiredFieldError implements the interface ErrorObject
type RequiredFieldError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

// MarshalJSON override to set the discriminator value
func (obj RequiredFieldError) MarshalJSON() ([]byte, error) {
	type Alias RequiredFieldError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "RequiredField", Alias: (*Alias)(&obj)})
}

func (obj RequiredFieldError) Error() string {
	return obj.Message
}

// ResourceNotFoundError implements the interface ErrorObject
type ResourceNotFoundError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ResourceNotFoundError) MarshalJSON() ([]byte, error) {
	type Alias ResourceNotFoundError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ResourceNotFound", Alias: (*Alias)(&obj)})
}

func (obj ResourceNotFoundError) Error() string {
	return obj.Message
}

// ShippingMethodDoesNotMatchCartError implements the interface ErrorObject
type ShippingMethodDoesNotMatchCartError struct {
	Message string `json:"message"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodDoesNotMatchCartError) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodDoesNotMatchCartError
	return json.Marshal(struct {
		Code string `json:"code"`
		*Alias
	}{Code: "ShippingMethodDoesNotMatchCart", Alias: (*Alias)(&obj)})
}

func (obj ShippingMethodDoesNotMatchCartError) Error() string {
	return obj.Message
}

// VariantValues is a standalone struct
type VariantValues struct {
	SKU        string      `json:"sku,omitempty"`
	Prices     []Price     `json:"prices"`
	Attributes []Attribute `json:"attributes"`
}
