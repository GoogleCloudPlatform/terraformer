// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// MyCartUpdateAction uses action as discriminator attribute
type MyCartUpdateAction interface{}

func mapDiscriminatorMyCartUpdateAction(input interface{}) (MyCartUpdateAction, error) {
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
	case "addDiscountCode":
		new := MyCartAddDiscountCodeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addItemShippingAddress":
		new := MyCartAddItemShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addLineItem":
		new := MyCartAddLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addPayment":
		new := MyCartAddPaymentAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "applyDeltaToLineItemShippingDetailsTargets":
		new := MyCartApplyDeltaToLineItemShippingDetailsTargetsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemQuantity":
		new := MyCartChangeLineItemQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTaxMode":
		new := MyCartChangeTaxModeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "recalculate":
		new := MyCartRecalculateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeDiscountCode":
		new := MyCartRemoveDiscountCodeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeItemShippingAddress":
		new := MyCartRemoveItemShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeLineItem":
		new := MyCartRemoveLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removePayment":
		new := MyCartRemovePaymentAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setBillingAddress":
		new := MyCartSetBillingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCountry":
		new := MyCartSetCountryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := MyCartSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomShippingMethod":
		new := MyCartSetCustomShippingMethodAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := MyCartSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeleteDaysAfterLastModification":
		new := MyCartSetDeleteDaysAfterLastModificationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomField":
		new := MyCartSetLineItemCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomType":
		new := MyCartSetLineItemCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemShippingDetails":
		new := MyCartSetLineItemShippingDetailsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocale":
		new := MyCartSetLocaleAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddress":
		new := MyCartSetShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingMethod":
		new := MyCartSetShippingMethodAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "updateItemShippingAddress":
		new := MyCartUpdateItemShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// MyCustomerUpdateAction uses action as discriminator attribute
type MyCustomerUpdateAction interface{}

func mapDiscriminatorMyCustomerUpdateAction(input interface{}) (MyCustomerUpdateAction, error) {
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
	case "addAddress":
		new := MyCustomerAddAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addBillingAddressId":
		new := MyCustomerAddBillingAddressIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addShippingAddressId":
		new := MyCustomerAddShippingAddressIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAddress":
		new := MyCustomerChangeAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeEmail":
		new := MyCustomerChangeEmailAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeAddress":
		new := MyCustomerRemoveAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeBillingAddressId":
		new := MyCustomerRemoveBillingAddressIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeShippingAddressId":
		new := MyCustomerRemoveShippingAddressIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCompanyName":
		new := MyCustomerSetCompanyNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := MyCustomerSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := MyCustomerSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDateOfBirth":
		new := MyCustomerSetDateOfBirthAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDefaultBillingAddress":
		new := MyCustomerSetDefaultBillingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDefaultShippingAddress":
		new := MyCustomerSetDefaultShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setFirstName":
		new := MyCustomerSetFirstNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLastName":
		new := MyCustomerSetLastNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocale":
		new := MyCustomerSetLocaleAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMiddleName":
		new := MyCustomerSetMiddleNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSalutation":
		new := MyCustomerSetSalutationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTitle":
		new := MyCustomerSetTitleAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setVatId":
		new := MyCustomerSetVatIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// MyPaymentUpdateAction uses action as discriminator attribute
type MyPaymentUpdateAction interface{}

func mapDiscriminatorMyPaymentUpdateAction(input interface{}) (MyPaymentUpdateAction, error) {
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
	case "addTransaction":
		new := MyPaymentAddTransactionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAmountPlanned":
		new := MyPaymentChangeAmountPlannedAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := MyPaymentSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoInterface":
		new := MyPaymentSetMethodInfoInterfaceAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoMethod":
		new := MyPaymentSetMethodInfoMethodAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoName":
		new := MyPaymentSetMethodInfoNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// MyShoppingListUpdateAction uses action as discriminator attribute
type MyShoppingListUpdateAction interface{}

func mapDiscriminatorMyShoppingListUpdateAction(input interface{}) (MyShoppingListUpdateAction, error) {
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
		new := MyShoppingListAddLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addTextLineItem":
		new := MyShoppingListAddTextLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemQuantity":
		new := MyShoppingListChangeLineItemQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemsOrder":
		new := MyShoppingListChangeLineItemsOrderAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := MyShoppingListChangeNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemName":
		new := MyShoppingListChangeTextLineItemNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemQuantity":
		new := MyShoppingListChangeTextLineItemQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTextLineItemsOrder":
		new := MyShoppingListChangeTextLineItemsOrderAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeLineItem":
		new := MyShoppingListRemoveLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeTextLineItem":
		new := MyShoppingListRemoveTextLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := MyShoppingListSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := MyShoppingListSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeleteDaysAfterLastModification":
		new := MyShoppingListSetDeleteDaysAfterLastModificationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := MyShoppingListSetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomField":
		new := MyShoppingListSetLineItemCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomType":
		new := MyShoppingListSetLineItemCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemCustomField":
		new := MyShoppingListSetTextLineItemCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemCustomType":
		new := MyShoppingListSetTextLineItemCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTextLineItemDescription":
		new := MyShoppingListSetTextLineItemDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// MyCart is of type BaseResource
type MyCart struct {
	Version                         int                     `json:"version"`
	TotalPrice                      TypedMoney              `json:"totalPrice"`
	TaxedPrice                      *TaxedPrice             `json:"taxedPrice,omitempty"`
	TaxRoundingMode                 RoundingMode            `json:"taxRoundingMode"`
	TaxMode                         TaxMode                 `json:"taxMode"`
	TaxCalculationMode              TaxCalculationMode      `json:"taxCalculationMode"`
	Store                           *StoreKeyReference      `json:"store,omitempty"`
	ShippingRateInput               ShippingRateInput       `json:"shippingRateInput,omitempty"`
	ShippingInfo                    *ShippingInfo           `json:"shippingInfo,omitempty"`
	ShippingAddress                 *Address                `json:"shippingAddress,omitempty"`
	RefusedGifts                    []CartDiscountReference `json:"refusedGifts"`
	PaymentInfo                     *PaymentInfo            `json:"paymentInfo,omitempty"`
	Origin                          CartOrigin              `json:"origin"`
	Locale                          string                  `json:"locale,omitempty"`
	LineItems                       []LineItem              `json:"lineItems"`
	LastModifiedBy                  *LastModifiedBy         `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time               `json:"lastModifiedAt"`
	ItemShippingAddresses           []Address               `json:"itemShippingAddresses,omitempty"`
	InventoryMode                   InventoryMode           `json:"inventoryMode,omitempty"`
	ID                              string                  `json:"id"`
	DiscountCodes                   []DiscountCodeInfo      `json:"discountCodes,omitempty"`
	DeleteDaysAfterLastModification int                     `json:"deleteDaysAfterLastModification,omitempty"`
	CustomerID                      string                  `json:"customerId,omitempty"`
	CustomerGroup                   *CustomerGroupReference `json:"customerGroup,omitempty"`
	CustomerEmail                   string                  `json:"customerEmail,omitempty"`
	CustomLineItems                 []CustomLineItem        `json:"customLineItems"`
	Custom                          *CustomFields           `json:"custom,omitempty"`
	CreatedBy                       *CreatedBy              `json:"createdBy,omitempty"`
	CreatedAt                       time.Time               `json:"createdAt"`
	Country                         CountryCode             `json:"country,omitempty"`
	CartState                       CartState               `json:"cartState"`
	BillingAddress                  *Address                `json:"billingAddress,omitempty"`
	AnonymousID                     string                  `json:"anonymousId,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyCart) UnmarshalJSON(data []byte) error {
	type Alias MyCart
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}
	if obj.TotalPrice != nil {
		var err error
		obj.TotalPrice, err = mapDiscriminatorTypedMoney(obj.TotalPrice)
		if err != nil {
			return err
		}
	}

	return nil
}

// MyCartAddDiscountCodeAction implements the interface MyCartUpdateAction
type MyCartAddDiscountCodeAction struct {
	Code string `json:"code"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartAddDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartAddDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDiscountCode", Alias: (*Alias)(&obj)})
}

// MyCartAddItemShippingAddressAction implements the interface MyCartUpdateAction
type MyCartAddItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartAddItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addItemShippingAddress", Alias: (*Alias)(&obj)})
}

// MyCartAddLineItemAction implements the interface MyCartUpdateAction
type MyCartAddLineItemAction struct {
	VariantID           int                         `json:"variantId,omitempty"`
	SupplyChannel       *ChannelResourceIdentifier  `json:"supplyChannel,omitempty"`
	SKU                 string                      `json:"sku,omitempty"`
	ShippingDetails     *ItemShippingDetailsDraft   `json:"shippingDetails,omitempty"`
	Quantity            float64                     `json:"quantity,omitempty"`
	ProductID           string                      `json:"productId,omitempty"`
	ExternalTotalPrice  *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalTaxRate     *ExternalTaxRateDraft       `json:"externalTaxRate,omitempty"`
	ExternalPrice       *Money                      `json:"externalPrice,omitempty"`
	DistributionChannel *ChannelResourceIdentifier  `json:"distributionChannel,omitempty"`
	Custom              *CustomFieldsDraft          `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

// MyCartAddPaymentAction implements the interface MyCartUpdateAction
type MyCartAddPaymentAction struct {
	Payment *PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartAddPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartAddPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPayment", Alias: (*Alias)(&obj)})
}

// MyCartApplyDeltaToLineItemShippingDetailsTargetsAction implements the interface MyCartUpdateAction
type MyCartApplyDeltaToLineItemShippingDetailsTargetsAction struct {
	TargetsDelta []ItemShippingTarget `json:"targetsDelta"`
	LineItemID   string               `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartApplyDeltaToLineItemShippingDetailsTargetsAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartApplyDeltaToLineItemShippingDetailsTargetsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "applyDeltaToLineItemShippingDetailsTargets", Alias: (*Alias)(&obj)})
}

// MyCartChangeLineItemQuantityAction implements the interface MyCartUpdateAction
type MyCartChangeLineItemQuantityAction struct {
	Quantity           float64                     `json:"quantity"`
	LineItemID         string                      `json:"lineItemId"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalPrice      *Money                      `json:"externalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

// MyCartChangeTaxModeAction implements the interface MyCartUpdateAction
type MyCartChangeTaxModeAction struct {
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartChangeTaxModeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartChangeTaxModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxMode", Alias: (*Alias)(&obj)})
}

// MyCartDraft is a standalone struct
type MyCartDraft struct {
	TaxMode                         TaxMode                           `json:"taxMode,omitempty"`
	ShippingMethod                  *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	ShippingAddress                 *Address                          `json:"shippingAddress,omitempty"`
	Locale                          string                            `json:"locale,omitempty"`
	LineItems                       []MyLineItemDraft                 `json:"lineItems,omitempty"`
	ItemShippingAddresses           []Address                         `json:"itemShippingAddresses,omitempty"`
	InventoryMode                   InventoryMode                     `json:"inventoryMode,omitempty"`
	DeleteDaysAfterLastModification int                               `json:"deleteDaysAfterLastModification,omitempty"`
	CustomerEmail                   string                            `json:"customerEmail,omitempty"`
	Custom                          *CustomFieldsDraft                `json:"custom,omitempty"`
	Currency                        CurrencyCode                      `json:"currency"`
	Country                         string                            `json:"country,omitempty"`
	BillingAddress                  *Address                          `json:"billingAddress,omitempty"`
}

// MyCartRecalculateAction implements the interface MyCartUpdateAction
type MyCartRecalculateAction struct {
	UpdateProductData bool `json:"updateProductData"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartRecalculateAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRecalculateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "recalculate", Alias: (*Alias)(&obj)})
}

// MyCartRemoveDiscountCodeAction implements the interface MyCartUpdateAction
type MyCartRemoveDiscountCodeAction struct {
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartRemoveDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRemoveDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDiscountCode", Alias: (*Alias)(&obj)})
}

// MyCartRemoveItemShippingAddressAction implements the interface MyCartUpdateAction
type MyCartRemoveItemShippingAddressAction struct {
	AddressKey string `json:"addressKey"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartRemoveItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRemoveItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeItemShippingAddress", Alias: (*Alias)(&obj)})
}

// MyCartRemoveLineItemAction implements the interface MyCartUpdateAction
type MyCartRemoveLineItemAction struct {
	ShippingDetailsToRemove *ItemShippingDetailsDraft   `json:"shippingDetailsToRemove,omitempty"`
	Quantity                float64                     `json:"quantity,omitempty"`
	LineItemID              string                      `json:"lineItemId"`
	ExternalTotalPrice      *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalPrice           *Money                      `json:"externalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

// MyCartRemovePaymentAction implements the interface MyCartUpdateAction
type MyCartRemovePaymentAction struct {
	Payment *PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartRemovePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartRemovePaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePayment", Alias: (*Alias)(&obj)})
}

// MyCartSetBillingAddressAction implements the interface MyCartUpdateAction
type MyCartSetBillingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

// MyCartSetCountryAction implements the interface MyCartUpdateAction
type MyCartSetCountryAction struct {
	Country CountryCode `json:"country,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetCountryAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetCountryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCountry", Alias: (*Alias)(&obj)})
}

// MyCartSetCustomFieldAction implements the interface MyCartUpdateAction
type MyCartSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// MyCartSetCustomShippingMethodAction implements the interface MyCartUpdateAction
type MyCartSetCustomShippingMethodAction struct {
	TaxCategory        *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	ShippingRate       *ShippingRateDraft             `json:"shippingRate"`
	ShippingMethodName string                         `json:"shippingMethodName"`
	ExternalTaxRate    *ExternalTaxRateDraft          `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetCustomShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetCustomShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomShippingMethod", Alias: (*Alias)(&obj)})
}

// MyCartSetCustomTypeAction implements the interface MyCartUpdateAction
type MyCartSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// MyCartSetDeleteDaysAfterLastModificationAction implements the interface MyCartUpdateAction
type MyCartSetDeleteDaysAfterLastModificationAction struct {
	DeleteDaysAfterLastModification int `json:"deleteDaysAfterLastModification,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

// MyCartSetLineItemCustomFieldAction implements the interface MyCartUpdateAction
type MyCartSetLineItemCustomFieldAction struct {
	Value      interface{} `json:"value,omitempty"`
	Name       string      `json:"name"`
	LineItemID string      `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

// MyCartSetLineItemCustomTypeAction implements the interface MyCartUpdateAction
type MyCartSetLineItemCustomTypeAction struct {
	Type       *TypeResourceIdentifier `json:"type,omitempty"`
	LineItemID string                  `json:"lineItemId"`
	Fields     *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

// MyCartSetLineItemShippingDetailsAction implements the interface MyCartUpdateAction
type MyCartSetLineItemShippingDetailsAction struct {
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	LineItemID      string                    `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

// MyCartSetLocaleAction implements the interface MyCartUpdateAction
type MyCartSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

// MyCartSetShippingAddressAction implements the interface MyCartUpdateAction
type MyCartSetShippingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

// MyCartSetShippingMethodAction implements the interface MyCartUpdateAction
type MyCartSetShippingMethodAction struct {
	ShippingMethod  *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft             `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartSetShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartSetShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethod", Alias: (*Alias)(&obj)})
}

// MyCartUpdateItemShippingAddressAction implements the interface MyCartUpdateAction
type MyCartUpdateItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCartUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCartUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

// MyCustomer is of type BaseResource
type MyCustomer struct {
	Version                  int                     `json:"version"`
	VatID                    string                  `json:"vatId,omitempty"`
	Title                    string                  `json:"title,omitempty"`
	Stores                   []StoreKeyReference     `json:"stores,omitempty"`
	ShippingAddressIds       []string                `json:"shippingAddressIds,omitempty"`
	Salutation               string                  `json:"salutation,omitempty"`
	Password                 string                  `json:"password"`
	MiddleName               string                  `json:"middleName,omitempty"`
	Locale                   string                  `json:"locale,omitempty"`
	LastName                 string                  `json:"lastName,omitempty"`
	LastModifiedBy           *LastModifiedBy         `json:"lastModifiedBy,omitempty"`
	LastModifiedAt           time.Time               `json:"lastModifiedAt"`
	Key                      string                  `json:"key,omitempty"`
	IsEmailVerified          bool                    `json:"isEmailVerified"`
	ID                       string                  `json:"id"`
	FirstName                string                  `json:"firstName,omitempty"`
	ExternalID               string                  `json:"externalId,omitempty"`
	Email                    string                  `json:"email"`
	DefaultShippingAddressID string                  `json:"defaultShippingAddressId,omitempty"`
	DefaultBillingAddressID  string                  `json:"defaultBillingAddressId,omitempty"`
	DateOfBirth              Date                    `json:"dateOfBirth,omitempty"`
	CustomerNumber           string                  `json:"customerNumber,omitempty"`
	CustomerGroup            *CustomerGroupReference `json:"customerGroup,omitempty"`
	Custom                   *CustomFields           `json:"custom,omitempty"`
	CreatedBy                *CreatedBy              `json:"createdBy,omitempty"`
	CreatedAt                time.Time               `json:"createdAt"`
	CompanyName              string                  `json:"companyName,omitempty"`
	BillingAddressIds        []string                `json:"billingAddressIds,omitempty"`
	Addresses                []Address               `json:"addresses"`
}

// MyCustomerAddAddressAction implements the interface MyCustomerUpdateAction
type MyCustomerAddAddressAction struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerAddAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerAddAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAddress", Alias: (*Alias)(&obj)})
}

// MyCustomerAddBillingAddressIDAction implements the interface MyCustomerUpdateAction
type MyCustomerAddBillingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerAddBillingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerAddBillingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addBillingAddressId", Alias: (*Alias)(&obj)})
}

// MyCustomerAddShippingAddressIDAction implements the interface MyCustomerUpdateAction
type MyCustomerAddShippingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerAddShippingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerAddShippingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingAddressId", Alias: (*Alias)(&obj)})
}

// MyCustomerChangeAddressAction implements the interface MyCustomerUpdateAction
type MyCustomerChangeAddressAction struct {
	AddressID string   `json:"addressId"`
	Address   *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerChangeAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerChangeAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAddress", Alias: (*Alias)(&obj)})
}

// MyCustomerChangeEmailAction implements the interface MyCustomerUpdateAction
type MyCustomerChangeEmailAction struct {
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerChangeEmailAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerChangeEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEmail", Alias: (*Alias)(&obj)})
}

// MyCustomerDraft is a standalone struct
type MyCustomerDraft struct {
	VatID                  string                    `json:"vatId,omitempty"`
	Title                  string                    `json:"title,omitempty"`
	Stores                 []StoreResourceIdentifier `json:"stores,omitempty"`
	Password               string                    `json:"password"`
	MiddleName             string                    `json:"middleName,omitempty"`
	Locale                 string                    `json:"locale,omitempty"`
	LastName               string                    `json:"lastName,omitempty"`
	FirstName              string                    `json:"firstName,omitempty"`
	Email                  string                    `json:"email"`
	DefaultShippingAddress int                       `json:"defaultShippingAddress,omitempty"`
	DefaultBillingAddress  int                       `json:"defaultBillingAddress,omitempty"`
	DateOfBirth            Date                      `json:"dateOfBirth,omitempty"`
	Custom                 *CustomFields             `json:"custom,omitempty"`
	CompanyName            string                    `json:"companyName,omitempty"`
	Addresses              []Address                 `json:"addresses,omitempty"`
}

// MyCustomerRemoveAddressAction implements the interface MyCustomerUpdateAction
type MyCustomerRemoveAddressAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerRemoveAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerRemoveAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAddress", Alias: (*Alias)(&obj)})
}

// MyCustomerRemoveBillingAddressIDAction implements the interface MyCustomerUpdateAction
type MyCustomerRemoveBillingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerRemoveBillingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerRemoveBillingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeBillingAddressId", Alias: (*Alias)(&obj)})
}

// MyCustomerRemoveShippingAddressIDAction implements the interface MyCustomerUpdateAction
type MyCustomerRemoveShippingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerRemoveShippingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerRemoveShippingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingAddressId", Alias: (*Alias)(&obj)})
}

// MyCustomerSetCompanyNameAction implements the interface MyCustomerUpdateAction
type MyCustomerSetCompanyNameAction struct {
	CompanyName string `json:"companyName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetCompanyNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetCompanyNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCompanyName", Alias: (*Alias)(&obj)})
}

// MyCustomerSetCustomFieldAction implements the interface MyCustomerUpdateAction
type MyCustomerSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// MyCustomerSetCustomTypeAction implements the interface MyCustomerUpdateAction
type MyCustomerSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// MyCustomerSetDateOfBirthAction implements the interface MyCustomerUpdateAction
type MyCustomerSetDateOfBirthAction struct {
	DateOfBirth Date `json:"dateOfBirth,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetDateOfBirthAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetDateOfBirthAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDateOfBirth", Alias: (*Alias)(&obj)})
}

// MyCustomerSetDefaultBillingAddressAction implements the interface MyCustomerUpdateAction
type MyCustomerSetDefaultBillingAddressAction struct {
	AddressID string `json:"addressId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetDefaultBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetDefaultBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultBillingAddress", Alias: (*Alias)(&obj)})
}

// MyCustomerSetDefaultShippingAddressAction implements the interface MyCustomerUpdateAction
type MyCustomerSetDefaultShippingAddressAction struct {
	AddressID string `json:"addressId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetDefaultShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetDefaultShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultShippingAddress", Alias: (*Alias)(&obj)})
}

// MyCustomerSetFirstNameAction implements the interface MyCustomerUpdateAction
type MyCustomerSetFirstNameAction struct {
	FirstName string `json:"firstName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetFirstNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetFirstNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setFirstName", Alias: (*Alias)(&obj)})
}

// MyCustomerSetLastNameAction implements the interface MyCustomerUpdateAction
type MyCustomerSetLastNameAction struct {
	LastName string `json:"lastName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetLastNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetLastNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLastName", Alias: (*Alias)(&obj)})
}

// MyCustomerSetLocaleAction implements the interface MyCustomerUpdateAction
type MyCustomerSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

// MyCustomerSetMiddleNameAction implements the interface MyCustomerUpdateAction
type MyCustomerSetMiddleNameAction struct {
	MiddleName string `json:"middleName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetMiddleNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetMiddleNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMiddleName", Alias: (*Alias)(&obj)})
}

// MyCustomerSetSalutationAction implements the interface MyCustomerUpdateAction
type MyCustomerSetSalutationAction struct {
	Salutation string `json:"salutation,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetSalutationAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetSalutationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSalutation", Alias: (*Alias)(&obj)})
}

// MyCustomerSetTitleAction implements the interface MyCustomerUpdateAction
type MyCustomerSetTitleAction struct {
	Title string `json:"title,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

// MyCustomerSetVatIDAction implements the interface MyCustomerUpdateAction
type MyCustomerSetVatIDAction struct {
	VatID string `json:"vatId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyCustomerSetVatIDAction) MarshalJSON() ([]byte, error) {
	type Alias MyCustomerSetVatIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setVatId", Alias: (*Alias)(&obj)})
}

// MyLineItemDraft is a standalone struct
type MyLineItemDraft struct {
	VariantID           int                        `json:"variantId"`
	SupplyChannel       *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	SKU                 string                     `json:"sku,omitempty"`
	ShippingDetails     *ItemShippingDetailsDraft  `json:"shippingDetails,omitempty"`
	Quantity            float64                    `json:"quantity"`
	ProductID           string                     `json:"productId"`
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	Custom              *CustomFieldsDraft         `json:"custom,omitempty"`
}

// MyOrder is of type BaseResource
type MyOrder struct {
	Version                   int                     `json:"version"`
	TotalPrice                TypedMoney              `json:"totalPrice"`
	TaxedPrice                *TaxedPrice             `json:"taxedPrice,omitempty"`
	TaxRoundingMode           RoundingMode            `json:"taxRoundingMode,omitempty"`
	TaxMode                   TaxMode                 `json:"taxMode,omitempty"`
	TaxCalculationMode        TaxCalculationMode      `json:"taxCalculationMode,omitempty"`
	SyncInfo                  []SyncInfo              `json:"syncInfo"`
	Store                     *StoreKeyReference      `json:"store,omitempty"`
	State                     *StateReference         `json:"state,omitempty"`
	ShippingRateInput         ShippingRateInput       `json:"shippingRateInput,omitempty"`
	ShippingInfo              *ShippingInfo           `json:"shippingInfo,omitempty"`
	ShippingAddress           *Address                `json:"shippingAddress,omitempty"`
	ShipmentState             ShipmentState           `json:"shipmentState,omitempty"`
	ReturnInfo                []ReturnInfo            `json:"returnInfo,omitempty"`
	RefusedGifts              []CartDiscountReference `json:"refusedGifts"`
	PaymentState              PaymentState            `json:"paymentState,omitempty"`
	PaymentInfo               *PaymentInfo            `json:"paymentInfo,omitempty"`
	Origin                    CartOrigin              `json:"origin"`
	OrderState                OrderState              `json:"orderState"`
	OrderNumber               string                  `json:"orderNumber,omitempty"`
	Locale                    string                  `json:"locale,omitempty"`
	LineItems                 []LineItem              `json:"lineItems"`
	LastModifiedBy            *LastModifiedBy         `json:"lastModifiedBy,omitempty"`
	LastModifiedAt            time.Time               `json:"lastModifiedAt"`
	LastMessageSequenceNumber int                     `json:"lastMessageSequenceNumber"`
	ItemShippingAddresses     []Address               `json:"itemShippingAddresses,omitempty"`
	InventoryMode             InventoryMode           `json:"inventoryMode,omitempty"`
	ID                        string                  `json:"id"`
	DiscountCodes             []DiscountCodeInfo      `json:"discountCodes,omitempty"`
	CustomerID                string                  `json:"customerId,omitempty"`
	CustomerGroup             *CustomerGroupReference `json:"customerGroup,omitempty"`
	CustomerEmail             string                  `json:"customerEmail,omitempty"`
	CustomLineItems           []CustomLineItem        `json:"customLineItems"`
	Custom                    *CustomFields           `json:"custom,omitempty"`
	CreatedBy                 *CreatedBy              `json:"createdBy,omitempty"`
	CreatedAt                 time.Time               `json:"createdAt"`
	Country                   string                  `json:"country,omitempty"`
	CompletedAt               *time.Time              `json:"completedAt,omitempty"`
	Cart                      *CartReference          `json:"cart,omitempty"`
	BillingAddress            *Address                `json:"billingAddress,omitempty"`
	AnonymousID               string                  `json:"anonymousId,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyOrder) UnmarshalJSON(data []byte) error {
	type Alias MyOrder
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}
	if obj.TotalPrice != nil {
		var err error
		obj.TotalPrice, err = mapDiscriminatorTypedMoney(obj.TotalPrice)
		if err != nil {
			return err
		}
	}

	return nil
}

// MyOrderFromCartDraft is a standalone struct
type MyOrderFromCartDraft struct {
	Version int    `json:"version"`
	ID      string `json:"id"`
}

// MyPayment is a standalone struct
type MyPayment struct {
	Version           int                `json:"version"`
	Transactions      []Transaction      `json:"transactions"`
	PaymentMethodInfo *PaymentMethodInfo `json:"paymentMethodInfo"`
	ID                string             `json:"id"`
	Customer          *CustomerReference `json:"customer,omitempty"`
	Custom            *CustomFields      `json:"custom,omitempty"`
	AnonymousID       string             `json:"anonymousId,omitempty"`
	AmountPlanned     TypedMoney         `json:"amountPlanned"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyPayment) UnmarshalJSON(data []byte) error {
	type Alias MyPayment
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.AmountPlanned != nil {
		var err error
		obj.AmountPlanned, err = mapDiscriminatorTypedMoney(obj.AmountPlanned)
		if err != nil {
			return err
		}
	}

	return nil
}

// MyPaymentAddTransactionAction implements the interface MyPaymentUpdateAction
type MyPaymentAddTransactionAction struct {
	Transaction *TransactionDraft `json:"transaction"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentAddTransactionAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentAddTransactionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTransaction", Alias: (*Alias)(&obj)})
}

// MyPaymentChangeAmountPlannedAction implements the interface MyPaymentUpdateAction
type MyPaymentChangeAmountPlannedAction struct {
	Amount *Money `json:"amount"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentChangeAmountPlannedAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentChangeAmountPlannedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAmountPlanned", Alias: (*Alias)(&obj)})
}

// MyPaymentDraft is a standalone struct
type MyPaymentDraft struct {
	Transaction       *MyTransactionDraft `json:"transaction,omitempty"`
	PaymentMethodInfo *PaymentMethodInfo  `json:"paymentMethodInfo,omitempty"`
	Custom            *CustomFieldsDraft  `json:"custom,omitempty"`
	AmountPlanned     *Money              `json:"amountPlanned"`
}

// MyPaymentPagedQueryResponse is a standalone struct
type MyPaymentPagedQueryResponse struct {
	Total   int         `json:"total,omitempty"`
	Results []MyPayment `json:"results"`
	Offset  int         `json:"offset"`
	Limit   int         `json:"limit"`
	Count   int         `json:"count"`
}

// MyPaymentSetCustomFieldAction implements the interface MyPaymentUpdateAction
type MyPaymentSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// MyPaymentSetMethodInfoInterfaceAction implements the interface MyPaymentUpdateAction
type MyPaymentSetMethodInfoInterfaceAction struct {
	Interface string `json:"interface"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentSetMethodInfoInterfaceAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetMethodInfoInterfaceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoInterface", Alias: (*Alias)(&obj)})
}

// MyPaymentSetMethodInfoMethodAction implements the interface MyPaymentUpdateAction
type MyPaymentSetMethodInfoMethodAction struct {
	Method string `json:"method,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentSetMethodInfoMethodAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetMethodInfoMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoMethod", Alias: (*Alias)(&obj)})
}

// MyPaymentSetMethodInfoNameAction implements the interface MyPaymentUpdateAction
type MyPaymentSetMethodInfoNameAction struct {
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyPaymentSetMethodInfoNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyPaymentSetMethodInfoNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoName", Alias: (*Alias)(&obj)})
}

// MyPaymentUpdate is a standalone struct
type MyPaymentUpdate struct {
	Version int                     `json:"version"`
	Actions []MyPaymentUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyPaymentUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyPaymentUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorMyPaymentUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MyShoppingListAddLineItemAction implements the interface MyShoppingListUpdateAction
type MyShoppingListAddLineItemAction struct {
	VariantID int                `json:"variantId,omitempty"`
	SKU       string             `json:"sku,omitempty"`
	Quantity  int                `json:"quantity,omitempty"`
	ProductID string             `json:"productId,omitempty"`
	Custom    *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt   *time.Time         `json:"addedAt,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

// MyShoppingListAddTextLineItemAction implements the interface MyShoppingListUpdateAction
type MyShoppingListAddTextLineItemAction struct {
	Quantity    int                `json:"quantity,omitempty"`
	Name        *LocalizedString   `json:"name"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	AddedAt     *time.Time         `json:"addedAt,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListAddTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListAddTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTextLineItem", Alias: (*Alias)(&obj)})
}

// MyShoppingListChangeLineItemQuantityAction implements the interface MyShoppingListUpdateAction
type MyShoppingListChangeLineItemQuantityAction struct {
	Quantity   int    `json:"quantity"`
	LineItemID string `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

// MyShoppingListChangeLineItemsOrderAction implements the interface MyShoppingListUpdateAction
type MyShoppingListChangeLineItemsOrderAction struct {
	LineItemOrder []string `json:"lineItemOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemsOrder", Alias: (*Alias)(&obj)})
}

// MyShoppingListChangeNameAction implements the interface MyShoppingListUpdateAction
type MyShoppingListChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// MyShoppingListChangeTextLineItemNameAction implements the interface MyShoppingListUpdateAction
type MyShoppingListChangeTextLineItemNameAction struct {
	TextLineItemID string           `json:"textLineItemId"`
	Name           *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeTextLineItemNameAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeTextLineItemNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemName", Alias: (*Alias)(&obj)})
}

// MyShoppingListChangeTextLineItemQuantityAction implements the interface MyShoppingListUpdateAction
type MyShoppingListChangeTextLineItemQuantityAction struct {
	TextLineItemID string `json:"textLineItemId"`
	Quantity       int    `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeTextLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeTextLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemQuantity", Alias: (*Alias)(&obj)})
}

// MyShoppingListChangeTextLineItemsOrderAction implements the interface MyShoppingListUpdateAction
type MyShoppingListChangeTextLineItemsOrderAction struct {
	TextLineItemOrder []string `json:"textLineItemOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListChangeTextLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListChangeTextLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTextLineItemsOrder", Alias: (*Alias)(&obj)})
}

// MyShoppingListDraft is a standalone struct
type MyShoppingListDraft struct {
	TextLineItems                   []TextLineItemDraft         `json:"textLineItems,omitempty"`
	Name                            *LocalizedString            `json:"name"`
	LineItems                       []ShoppingListLineItemDraft `json:"lineItems,omitempty"`
	Description                     *LocalizedString            `json:"description,omitempty"`
	DeleteDaysAfterLastModification int                         `json:"deleteDaysAfterLastModification,omitempty"`
	Custom                          *CustomFieldsDraft          `json:"custom,omitempty"`
}

// MyShoppingListRemoveLineItemAction implements the interface MyShoppingListUpdateAction
type MyShoppingListRemoveLineItemAction struct {
	Quantity   int    `json:"quantity,omitempty"`
	LineItemID string `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

// MyShoppingListRemoveTextLineItemAction implements the interface MyShoppingListUpdateAction
type MyShoppingListRemoveTextLineItemAction struct {
	TextLineItemID string `json:"textLineItemId"`
	Quantity       int    `json:"quantity,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListRemoveTextLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListRemoveTextLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTextLineItem", Alias: (*Alias)(&obj)})
}

// MyShoppingListSetCustomFieldAction implements the interface MyShoppingListUpdateAction
type MyShoppingListSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// MyShoppingListSetCustomTypeAction implements the interface MyShoppingListUpdateAction
type MyShoppingListSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// MyShoppingListSetDeleteDaysAfterLastModificationAction implements the interface MyShoppingListUpdateAction
type MyShoppingListSetDeleteDaysAfterLastModificationAction struct {
	DeleteDaysAfterLastModification int `json:"deleteDaysAfterLastModification,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

// MyShoppingListSetDescriptionAction implements the interface MyShoppingListUpdateAction
type MyShoppingListSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// MyShoppingListSetLineItemCustomFieldAction implements the interface MyShoppingListUpdateAction
type MyShoppingListSetLineItemCustomFieldAction struct {
	Value      interface{} `json:"value,omitempty"`
	Name       string      `json:"name"`
	LineItemID string      `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

// MyShoppingListSetLineItemCustomTypeAction implements the interface MyShoppingListUpdateAction
type MyShoppingListSetLineItemCustomTypeAction struct {
	Type       *TypeResourceIdentifier `json:"type,omitempty"`
	LineItemID string                  `json:"lineItemId"`
	Fields     *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

// MyShoppingListSetTextLineItemCustomFieldAction implements the interface MyShoppingListUpdateAction
type MyShoppingListSetTextLineItemCustomFieldAction struct {
	Value          interface{} `json:"value,omitempty"`
	TextLineItemID string      `json:"textLineItemId"`
	Name           string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetTextLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetTextLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomField", Alias: (*Alias)(&obj)})
}

// MyShoppingListSetTextLineItemCustomTypeAction implements the interface MyShoppingListUpdateAction
type MyShoppingListSetTextLineItemCustomTypeAction struct {
	Type           *TypeResourceIdentifier `json:"type,omitempty"`
	TextLineItemID string                  `json:"textLineItemId"`
	Fields         *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetTextLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetTextLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemCustomType", Alias: (*Alias)(&obj)})
}

// MyShoppingListSetTextLineItemDescriptionAction implements the interface MyShoppingListUpdateAction
type MyShoppingListSetTextLineItemDescriptionAction struct {
	TextLineItemID string           `json:"textLineItemId"`
	Description    *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj MyShoppingListSetTextLineItemDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias MyShoppingListSetTextLineItemDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTextLineItemDescription", Alias: (*Alias)(&obj)})
}

// MyShoppingListUpdate is a standalone struct
type MyShoppingListUpdate struct {
	Version int                          `json:"version"`
	Actions []MyShoppingListUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MyShoppingListUpdate) UnmarshalJSON(data []byte) error {
	type Alias MyShoppingListUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorMyShoppingListUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MyTransactionDraft is a standalone struct
type MyTransactionDraft struct {
	Type          TransactionType `json:"type"`
	Timestamp     *time.Time      `json:"timestamp,omitempty"`
	InteractionID string          `json:"interactionId,omitempty"`
	Amount        *Money          `json:"amount"`
}
