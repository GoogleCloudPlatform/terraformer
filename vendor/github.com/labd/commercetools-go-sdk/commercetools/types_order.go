// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// OrderState is an enum type
type OrderState string

// Enum values for OrderState
const (
	OrderStateOpen      OrderState = "Open"
	OrderStateConfirmed OrderState = "Confirmed"
	OrderStateComplete  OrderState = "Complete"
	OrderStateCancelled OrderState = "Cancelled"
)

// PaymentState is an enum type
type PaymentState string

// Enum values for PaymentState
const (
	PaymentStateBalanceDue PaymentState = "BalanceDue"
	PaymentStateFailed     PaymentState = "Failed"
	PaymentStatePending    PaymentState = "Pending"
	PaymentStateCreditOwed PaymentState = "CreditOwed"
	PaymentStatePaid       PaymentState = "Paid"
)

// ReturnPaymentState is an enum type
type ReturnPaymentState string

// Enum values for ReturnPaymentState
const (
	ReturnPaymentStateNonRefundable ReturnPaymentState = "NonRefundable"
	ReturnPaymentStateInitial       ReturnPaymentState = "Initial"
	ReturnPaymentStateRefunded      ReturnPaymentState = "Refunded"
	ReturnPaymentStateNotRefunded   ReturnPaymentState = "NotRefunded"
)

// ReturnShipmentState is an enum type
type ReturnShipmentState string

// Enum values for ReturnShipmentState
const (
	ReturnShipmentStateAdvised     ReturnShipmentState = "Advised"
	ReturnShipmentStateReturned    ReturnShipmentState = "Returned"
	ReturnShipmentStateBackInStock ReturnShipmentState = "BackInStock"
	ReturnShipmentStateUnusable    ReturnShipmentState = "Unusable"
)

// ShipmentState is an enum type
type ShipmentState string

// Enum values for ShipmentState
const (
	ShipmentStateShipped   ShipmentState = "Shipped"
	ShipmentStateReady     ShipmentState = "Ready"
	ShipmentStatePending   ShipmentState = "Pending"
	ShipmentStateDelayed   ShipmentState = "Delayed"
	ShipmentStatePartial   ShipmentState = "Partial"
	ShipmentStateBackorder ShipmentState = "Backorder"
)

// OrderUpdateAction uses action as discriminator attribute
type OrderUpdateAction interface{}

func mapDiscriminatorOrderUpdateAction(input interface{}) (OrderUpdateAction, error) {
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
	case "addDelivery":
		new := OrderAddDeliveryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addItemShippingAddress":
		new := OrderAddItemShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addParcelToDelivery":
		new := OrderAddParcelToDeliveryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addPayment":
		new := OrderAddPaymentAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addReturnInfo":
		new := OrderAddReturnInfoAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeOrderState":
		new := OrderChangeOrderStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changePaymentState":
		new := OrderChangePaymentStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeShipmentState":
		new := OrderChangeShipmentStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "importCustomLineItemState":
		new := OrderImportCustomLineItemStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "importLineItemState":
		new := OrderImportLineItemStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeDelivery":
		new := OrderRemoveDeliveryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeItemShippingAddress":
		new := OrderRemoveItemShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeParcelFromDelivery":
		new := OrderRemoveParcelFromDeliveryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removePayment":
		new := OrderRemovePaymentAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setBillingAddress":
		new := OrderSetBillingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := OrderSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemCustomField":
		new := OrderSetCustomLineItemCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemCustomType":
		new := OrderSetCustomLineItemCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemShippingDetails":
		new := OrderSetCustomLineItemShippingDetailsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := OrderSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerEmail":
		new := OrderSetCustomerEmailAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerId":
		new := OrderSetCustomerIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryAddress":
		new := OrderSetDeliveryAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryItems":
		new := OrderSetDeliveryItemsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomField":
		new := OrderSetLineItemCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomType":
		new := OrderSetLineItemCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemShippingDetails":
		new := OrderSetLineItemShippingDetailsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocale":
		new := OrderSetLocaleAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setOrderNumber":
		new := OrderSetOrderNumberAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelItems":
		new := OrderSetParcelItemsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelMeasurements":
		new := OrderSetParcelMeasurementsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelTrackingData":
		new := OrderSetParcelTrackingDataAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setReturnPaymentState":
		new := OrderSetReturnPaymentStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setReturnShipmentState":
		new := OrderSetReturnShipmentStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddress":
		new := OrderSetShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionCustomLineItemState":
		new := OrderTransitionCustomLineItemStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionLineItemState":
		new := OrderTransitionLineItemStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionState":
		new := OrderTransitionStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "updateItemShippingAddress":
		new := OrderUpdateItemShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "updateSyncInfo":
		new := OrderUpdateSyncInfoAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// ReturnItem uses type as discriminator attribute
type ReturnItem interface{}

func mapDiscriminatorReturnItem(input interface{}) (ReturnItem, error) {
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
	case "CustomLineItemReturnItem":
		new := CustomLineItemReturnItem{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LineItemReturnItem":
		new := LineItemReturnItem{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// StagedOrderUpdateAction uses action as discriminator attribute
type StagedOrderUpdateAction interface{}

func mapDiscriminatorStagedOrderUpdateAction(input interface{}) (StagedOrderUpdateAction, error) {
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
	case "addCustomLineItem":
		new := StagedOrderAddCustomLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addDelivery":
		new := StagedOrderAddDeliveryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addDiscountCode":
		new := StagedOrderAddDiscountCodeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addItemShippingAddress":
		new := StagedOrderAddItemShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addLineItem":
		new := StagedOrderAddLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addParcelToDelivery":
		new := StagedOrderAddParcelToDeliveryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addPayment":
		new := StagedOrderAddPaymentAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addReturnInfo":
		new := StagedOrderAddReturnInfoAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addShoppingList":
		new := StagedOrderAddShoppingListAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeCustomLineItemMoney":
		new := StagedOrderChangeCustomLineItemMoneyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeCustomLineItemQuantity":
		new := StagedOrderChangeCustomLineItemQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemQuantity":
		new := StagedOrderChangeLineItemQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeOrderState":
		new := StagedOrderChangeOrderStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changePaymentState":
		new := StagedOrderChangePaymentStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeShipmentState":
		new := StagedOrderChangeShipmentStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTaxCalculationMode":
		new := StagedOrderChangeTaxCalculationModeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTaxMode":
		new := StagedOrderChangeTaxModeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTaxRoundingMode":
		new := StagedOrderChangeTaxRoundingModeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "importCustomLineItemState":
		new := StagedOrderImportCustomLineItemStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "importLineItemState":
		new := StagedOrderImportLineItemStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeCustomLineItem":
		new := StagedOrderRemoveCustomLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeDelivery":
		new := StagedOrderRemoveDeliveryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeDiscountCode":
		new := StagedOrderRemoveDiscountCodeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeItemShippingAddress":
		new := StagedOrderRemoveItemShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeLineItem":
		new := StagedOrderRemoveLineItemAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeParcelFromDelivery":
		new := StagedOrderRemoveParcelFromDeliveryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removePayment":
		new := StagedOrderRemovePaymentAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setBillingAddress":
		new := StagedOrderSetBillingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCountry":
		new := StagedOrderSetCountryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := StagedOrderSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemCustomField":
		new := StagedOrderSetCustomLineItemCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemCustomType":
		new := StagedOrderSetCustomLineItemCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemShippingDetails":
		new := StagedOrderSetCustomLineItemShippingDetailsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemTaxAmount":
		new := StagedOrderSetCustomLineItemTaxAmountAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemTaxRate":
		new := StagedOrderSetCustomLineItemTaxRateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomShippingMethod":
		new := StagedOrderSetCustomShippingMethodAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := StagedOrderSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerEmail":
		new := StagedOrderSetCustomerEmailAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerGroup":
		new := StagedOrderSetCustomerGroupAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerId":
		new := StagedOrderSetCustomerIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryAddress":
		new := StagedOrderSetDeliveryAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryItems":
		new := StagedOrderSetDeliveryItemsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomField":
		new := StagedOrderSetLineItemCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomType":
		new := StagedOrderSetLineItemCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemPrice":
		new := StagedOrderSetLineItemPriceAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemShippingDetails":
		new := StagedOrderSetLineItemShippingDetailsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemTaxAmount":
		new := StagedOrderSetLineItemTaxAmountAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemTaxRate":
		new := StagedOrderSetLineItemTaxRateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemTotalPrice":
		new := StagedOrderSetLineItemTotalPriceAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocale":
		new := StagedOrderSetLocaleAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setOrderNumber":
		new := StagedOrderSetOrderNumberAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setOrderTotalTax":
		new := StagedOrderSetOrderTotalTaxAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelItems":
		new := StagedOrderSetParcelItemsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelMeasurements":
		new := StagedOrderSetParcelMeasurementsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelTrackingData":
		new := StagedOrderSetParcelTrackingDataAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setReturnPaymentState":
		new := StagedOrderSetReturnPaymentStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setReturnShipmentState":
		new := StagedOrderSetReturnShipmentStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddress":
		new := StagedOrderSetShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddressAndCustomShippingMethod":
		new := StagedOrderSetShippingAddressAndCustomShippingMethodAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddressAndShippingMethod":
		new := StagedOrderSetShippingAddressAndShippingMethodAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingMethod":
		new := StagedOrderSetShippingMethodAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingMethodTaxAmount":
		new := StagedOrderSetShippingMethodTaxAmountAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingMethodTaxRate":
		new := StagedOrderSetShippingMethodTaxRateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingRateInput":
		new := StagedOrderSetShippingRateInputAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.ShippingRateInput != nil {
			new.ShippingRateInput, err = mapDiscriminatorShippingRateInputDraft(new.ShippingRateInput)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "transitionCustomLineItemState":
		new := StagedOrderTransitionCustomLineItemStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionLineItemState":
		new := StagedOrderTransitionLineItemStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionState":
		new := StagedOrderTransitionStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "updateItemShippingAddress":
		new := StagedOrderUpdateItemShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "updateSyncInfo":
		new := StagedOrderUpdateSyncInfoAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// CustomLineItemReturnItem implements the interface ReturnItem
type CustomLineItemReturnItem struct {
	ShipmentState    ReturnShipmentState `json:"shipmentState"`
	Quantity         int                 `json:"quantity"`
	PaymentState     ReturnPaymentState  `json:"paymentState"`
	LastModifiedAt   time.Time           `json:"lastModifiedAt"`
	ID               string              `json:"id"`
	CreatedAt        time.Time           `json:"createdAt"`
	Comment          string              `json:"comment,omitempty"`
	CustomLineItemID string              `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomLineItemReturnItem) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemReturnItem
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomLineItemReturnItem", Alias: (*Alias)(&obj)})
}

// Delivery is a standalone struct
type Delivery struct {
	Parcels   []Parcel       `json:"parcels"`
	Items     []DeliveryItem `json:"items"`
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	Address   *Address       `json:"address,omitempty"`
}

// DeliveryItem is a standalone struct
type DeliveryItem struct {
	Quantity float64 `json:"quantity"`
	ID       string  `json:"id"`
}

// DiscountedLineItemPriceDraft is a standalone struct
type DiscountedLineItemPriceDraft struct {
	Value             *Money                      `json:"value"`
	IncludedDiscounts []DiscountedLineItemPortion `json:"includedDiscounts"`
}

// ItemState is a standalone struct
type ItemState struct {
	State    *StateReference `json:"state"`
	Quantity float64         `json:"quantity"`
}

// LineItemImportDraft is a standalone struct
type LineItemImportDraft struct {
	Variant             *ProductVariantImportDraft `json:"variant"`
	TaxRate             *TaxRate                   `json:"taxRate,omitempty"`
	SupplyChannel       *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	State               []ItemState                `json:"state,omitempty"`
	ShippingDetails     *ItemShippingDetailsDraft  `json:"shippingDetails,omitempty"`
	Quantity            float64                    `json:"quantity"`
	ProductID           string                     `json:"productId,omitempty"`
	Price               *PriceDraft                `json:"price"`
	Name                *LocalizedString           `json:"name"`
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	Custom              *CustomFieldsDraft         `json:"custom,omitempty"`
}

// LineItemReturnItem implements the interface ReturnItem
type LineItemReturnItem struct {
	ShipmentState  ReturnShipmentState `json:"shipmentState"`
	Quantity       int                 `json:"quantity"`
	PaymentState   ReturnPaymentState  `json:"paymentState"`
	LastModifiedAt time.Time           `json:"lastModifiedAt"`
	ID             string              `json:"id"`
	CreatedAt      time.Time           `json:"createdAt"`
	Comment        string              `json:"comment,omitempty"`
	LineItemID     string              `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj LineItemReturnItem) MarshalJSON() ([]byte, error) {
	type Alias LineItemReturnItem
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "LineItemReturnItem", Alias: (*Alias)(&obj)})
}

// Order is of type LoggedResource
type Order struct {
	Version                   int                     `json:"version"`
	LastModifiedAt            time.Time               `json:"lastModifiedAt"`
	ID                        string                  `json:"id"`
	CreatedAt                 time.Time               `json:"createdAt"`
	LastModifiedBy            *LastModifiedBy         `json:"lastModifiedBy,omitempty"`
	CreatedBy                 *CreatedBy              `json:"createdBy,omitempty"`
	TotalPrice                *Money                  `json:"totalPrice"`
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
	PaymentState              PaymentState            `json:"paymentState,omitempty"`
	PaymentInfo               *PaymentInfo            `json:"paymentInfo,omitempty"`
	Origin                    CartOrigin              `json:"origin"`
	OrderState                OrderState              `json:"orderState"`
	OrderNumber               string                  `json:"orderNumber,omitempty"`
	Locale                    string                  `json:"locale,omitempty"`
	LineItems                 []LineItem              `json:"lineItems"`
	LastMessageSequenceNumber int                     `json:"lastMessageSequenceNumber"`
	ItemShippingAddresses     []Address               `json:"itemShippingAddresses,omitempty"`
	InventoryMode             InventoryMode           `json:"inventoryMode,omitempty"`
	DiscountCodes             []DiscountCodeInfo      `json:"discountCodes,omitempty"`
	CustomerID                string                  `json:"customerId,omitempty"`
	CustomerGroup             *CustomerGroupReference `json:"customerGroup,omitempty"`
	CustomerEmail             string                  `json:"customerEmail,omitempty"`
	CustomLineItems           []CustomLineItem        `json:"customLineItems"`
	Custom                    *CustomFields           `json:"custom,omitempty"`
	Country                   string                  `json:"country,omitempty"`
	CompletedAt               *time.Time              `json:"completedAt,omitempty"`
	Cart                      *CartReference          `json:"cart,omitempty"`
	BillingAddress            *Address                `json:"billingAddress,omitempty"`
	AnonymousID               string                  `json:"anonymousId,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Order) UnmarshalJSON(data []byte) error {
	type Alias Order
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

	return nil
}

// OrderAddDeliveryAction implements the interface OrderUpdateAction
type OrderAddDeliveryAction struct {
	Parcels []ParcelDraft  `json:"parcels,omitempty"`
	Items   []DeliveryItem `json:"items,omitempty"`
	Address *Address       `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderAddDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDelivery", Alias: (*Alias)(&obj)})
}

// OrderAddItemShippingAddressAction implements the interface OrderUpdateAction
type OrderAddItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addItemShippingAddress", Alias: (*Alias)(&obj)})
}

// OrderAddParcelToDeliveryAction implements the interface OrderUpdateAction
type OrderAddParcelToDeliveryAction struct {
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	Items        []DeliveryItem      `json:"items,omitempty"`
	DeliveryID   string              `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderAddParcelToDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddParcelToDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addParcelToDelivery", Alias: (*Alias)(&obj)})
}

// OrderAddPaymentAction implements the interface OrderUpdateAction
type OrderAddPaymentAction struct {
	Payment *PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderAddPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPayment", Alias: (*Alias)(&obj)})
}

// OrderAddReturnInfoAction implements the interface OrderUpdateAction
type OrderAddReturnInfoAction struct {
	ReturnTrackingID string            `json:"returnTrackingId,omitempty"`
	ReturnDate       *time.Time        `json:"returnDate,omitempty"`
	Items            []ReturnItemDraft `json:"items"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderAddReturnInfoAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddReturnInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addReturnInfo", Alias: (*Alias)(&obj)})
}

// OrderChangeOrderStateAction implements the interface OrderUpdateAction
type OrderChangeOrderStateAction struct {
	OrderState OrderState `json:"orderState"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderChangeOrderStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderChangeOrderStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeOrderState", Alias: (*Alias)(&obj)})
}

// OrderChangePaymentStateAction implements the interface OrderUpdateAction
type OrderChangePaymentStateAction struct {
	PaymentState PaymentState `json:"paymentState,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderChangePaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderChangePaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePaymentState", Alias: (*Alias)(&obj)})
}

// OrderChangeShipmentStateAction implements the interface OrderUpdateAction
type OrderChangeShipmentStateAction struct {
	ShipmentState ShipmentState `json:"shipmentState,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderChangeShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderChangeShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeShipmentState", Alias: (*Alias)(&obj)})
}

// OrderFromCartDraft is a standalone struct
type OrderFromCartDraft struct {
	Version      int          `json:"version"`
	PaymentState PaymentState `json:"paymentState,omitempty"`
	OrderNumber  string       `json:"orderNumber,omitempty"`
	ID           string       `json:"id"`
}

// OrderImportCustomLineItemStateAction implements the interface OrderUpdateAction
type OrderImportCustomLineItemStateAction struct {
	State            []ItemState `json:"state"`
	CustomLineItemID string      `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderImportCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderImportCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importCustomLineItemState", Alias: (*Alias)(&obj)})
}

// OrderImportDraft is a standalone struct
type OrderImportDraft struct {
	TotalPrice            *Money                           `json:"totalPrice"`
	TaxedPrice            *TaxedPrice                      `json:"taxedPrice,omitempty"`
	TaxRoundingMode       RoundingMode                     `json:"taxRoundingMode,omitempty"`
	ShippingInfo          *ShippingInfoImportDraft         `json:"shippingInfo,omitempty"`
	ShippingAddress       *Address                         `json:"shippingAddress,omitempty"`
	ShipmentState         ShipmentState                    `json:"shipmentState,omitempty"`
	PaymentState          PaymentState                     `json:"paymentState,omitempty"`
	OrderState            OrderState                       `json:"orderState,omitempty"`
	OrderNumber           string                           `json:"orderNumber,omitempty"`
	LineItems             []LineItemImportDraft            `json:"lineItems,omitempty"`
	ItemShippingAddresses []Address                        `json:"itemShippingAddresses,omitempty"`
	InventoryMode         InventoryMode                    `json:"inventoryMode,omitempty"`
	CustomerID            string                           `json:"customerId,omitempty"`
	CustomerGroup         *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	CustomerEmail         string                           `json:"customerEmail,omitempty"`
	CustomLineItems       []CustomLineItemDraft            `json:"customLineItems,omitempty"`
	Custom                *CustomFieldsDraft               `json:"custom,omitempty"`
	Country               string                           `json:"country,omitempty"`
	CompletedAt           *time.Time                       `json:"completedAt,omitempty"`
	BillingAddress        *Address                         `json:"billingAddress,omitempty"`
}

// OrderImportLineItemStateAction implements the interface OrderUpdateAction
type OrderImportLineItemStateAction struct {
	State      []ItemState `json:"state"`
	LineItemID string      `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderImportLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderImportLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importLineItemState", Alias: (*Alias)(&obj)})
}

// OrderPagedQueryResponse is a standalone struct
type OrderPagedQueryResponse struct {
	Total   int     `json:"total,omitempty"`
	Results []Order `json:"results"`
	Offset  int     `json:"offset"`
	Count   int     `json:"count"`
}

// OrderReference implements the interface Reference
type OrderReference struct {
	ID  string `json:"id"`
	Obj *Order `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderReference) MarshalJSON() ([]byte, error) {
	type Alias OrderReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "order", Alias: (*Alias)(&obj)})
}

// OrderRemoveDeliveryAction implements the interface OrderUpdateAction
type OrderRemoveDeliveryAction struct {
	DeliveryID string `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderRemoveDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemoveDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDelivery", Alias: (*Alias)(&obj)})
}

// OrderRemoveItemShippingAddressAction implements the interface OrderUpdateAction
type OrderRemoveItemShippingAddressAction struct {
	AddressKey string `json:"addressKey"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderRemoveItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemoveItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeItemShippingAddress", Alias: (*Alias)(&obj)})
}

// OrderRemoveParcelFromDeliveryAction implements the interface OrderUpdateAction
type OrderRemoveParcelFromDeliveryAction struct {
	ParcelID string `json:"parcelId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderRemoveParcelFromDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemoveParcelFromDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeParcelFromDelivery", Alias: (*Alias)(&obj)})
}

// OrderRemovePaymentAction implements the interface OrderUpdateAction
type OrderRemovePaymentAction struct {
	Payment *PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderRemovePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemovePaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePayment", Alias: (*Alias)(&obj)})
}

// OrderResourceIdentifier implements the interface ResourceIdentifier
type OrderResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias OrderResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "order", Alias: (*Alias)(&obj)})
}

// OrderSetBillingAddressAction implements the interface OrderUpdateAction
type OrderSetBillingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

// OrderSetCustomFieldAction implements the interface OrderUpdateAction
type OrderSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// OrderSetCustomLineItemCustomFieldAction implements the interface OrderUpdateAction
type OrderSetCustomLineItemCustomFieldAction struct {
	Value            interface{} `json:"value,omitempty"`
	Name             string      `json:"name"`
	CustomLineItemID string      `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomField", Alias: (*Alias)(&obj)})
}

// OrderSetCustomLineItemCustomTypeAction implements the interface OrderUpdateAction
type OrderSetCustomLineItemCustomTypeAction struct {
	Type             *TypeResourceIdentifier `json:"type,omitempty"`
	Fields           *FieldContainer         `json:"fields,omitempty"`
	CustomLineItemID string                  `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomType", Alias: (*Alias)(&obj)})
}

// OrderSetCustomLineItemShippingDetailsAction implements the interface OrderUpdateAction
type OrderSetCustomLineItemShippingDetailsAction struct {
	ShippingDetails  *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	CustomLineItemID string                    `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

// OrderSetCustomTypeAction implements the interface OrderUpdateAction
type OrderSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// OrderSetCustomerEmailAction implements the interface OrderUpdateAction
type OrderSetCustomerEmailAction struct {
	Email string `json:"email,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomerEmailAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomerEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerEmail", Alias: (*Alias)(&obj)})
}

// OrderSetCustomerIDAction implements the interface OrderUpdateAction
type OrderSetCustomerIDAction struct {
	CustomerID string `json:"customerId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomerIDAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomerIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerId", Alias: (*Alias)(&obj)})
}

// OrderSetDeliveryAddressAction implements the interface OrderUpdateAction
type OrderSetDeliveryAddressAction struct {
	DeliveryID string   `json:"deliveryId"`
	Address    *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetDeliveryAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddress", Alias: (*Alias)(&obj)})
}

// OrderSetDeliveryItemsAction implements the interface OrderUpdateAction
type OrderSetDeliveryItemsAction struct {
	Items      []DeliveryItem `json:"items"`
	DeliveryID string         `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetDeliveryItemsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryItems", Alias: (*Alias)(&obj)})
}

// OrderSetLineItemCustomFieldAction implements the interface OrderUpdateAction
type OrderSetLineItemCustomFieldAction struct {
	Value      interface{} `json:"value,omitempty"`
	Name       string      `json:"name"`
	LineItemID string      `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

// OrderSetLineItemCustomTypeAction implements the interface OrderUpdateAction
type OrderSetLineItemCustomTypeAction struct {
	Type       *TypeResourceIdentifier `json:"type,omitempty"`
	LineItemID string                  `json:"lineItemId"`
	Fields     *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

// OrderSetLineItemShippingDetailsAction implements the interface OrderUpdateAction
type OrderSetLineItemShippingDetailsAction struct {
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	LineItemID      string                    `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

// OrderSetLocaleAction implements the interface OrderUpdateAction
type OrderSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

// OrderSetOrderNumberAction implements the interface OrderUpdateAction
type OrderSetOrderNumberAction struct {
	OrderNumber string `json:"orderNumber,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetOrderNumberAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetOrderNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setOrderNumber", Alias: (*Alias)(&obj)})
}

// OrderSetParcelItemsAction implements the interface OrderUpdateAction
type OrderSetParcelItemsAction struct {
	ParcelID string         `json:"parcelId"`
	Items    []DeliveryItem `json:"items"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetParcelItemsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelItems", Alias: (*Alias)(&obj)})
}

// OrderSetParcelMeasurementsAction implements the interface OrderUpdateAction
type OrderSetParcelMeasurementsAction struct {
	ParcelID     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetParcelMeasurementsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelMeasurementsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelMeasurements", Alias: (*Alias)(&obj)})
}

// OrderSetParcelTrackingDataAction implements the interface OrderUpdateAction
type OrderSetParcelTrackingDataAction struct {
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	ParcelID     string        `json:"parcelId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetParcelTrackingDataAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelTrackingDataAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelTrackingData", Alias: (*Alias)(&obj)})
}

// OrderSetReturnPaymentStateAction implements the interface OrderUpdateAction
type OrderSetReturnPaymentStateAction struct {
	ReturnItemID string             `json:"returnItemId"`
	PaymentState ReturnPaymentState `json:"paymentState"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetReturnPaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetReturnPaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnPaymentState", Alias: (*Alias)(&obj)})
}

// OrderSetReturnShipmentStateAction implements the interface OrderUpdateAction
type OrderSetReturnShipmentStateAction struct {
	ShipmentState ReturnShipmentState `json:"shipmentState"`
	ReturnItemID  string              `json:"returnItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetReturnShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetReturnShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnShipmentState", Alias: (*Alias)(&obj)})
}

// OrderSetShippingAddressAction implements the interface OrderUpdateAction
type OrderSetShippingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

// OrderTransitionCustomLineItemStateAction implements the interface OrderUpdateAction
type OrderTransitionCustomLineItemStateAction struct {
	ToState              *StateResourceIdentifier `json:"toState"`
	Quantity             int                      `json:"quantity"`
	FromState            *StateResourceIdentifier `json:"fromState"`
	CustomLineItemID     string                   `json:"customLineItemId"`
	ActualTransitionDate *time.Time               `json:"actualTransitionDate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderTransitionCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderTransitionCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionCustomLineItemState", Alias: (*Alias)(&obj)})
}

// OrderTransitionLineItemStateAction implements the interface OrderUpdateAction
type OrderTransitionLineItemStateAction struct {
	ToState              *StateResourceIdentifier `json:"toState"`
	Quantity             int                      `json:"quantity"`
	LineItemID           string                   `json:"lineItemId"`
	FromState            *StateResourceIdentifier `json:"fromState"`
	ActualTransitionDate *time.Time               `json:"actualTransitionDate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderTransitionLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderTransitionLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionLineItemState", Alias: (*Alias)(&obj)})
}

// OrderTransitionStateAction implements the interface OrderUpdateAction
type OrderTransitionStateAction struct {
	State *StateResourceIdentifier `json:"state"`
	Force bool                     `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

// OrderUpdate is a standalone struct
type OrderUpdate struct {
	Version int                 `json:"version"`
	Actions []OrderUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderUpdate) UnmarshalJSON(data []byte) error {
	type Alias OrderUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorOrderUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// OrderUpdateItemShippingAddressAction implements the interface OrderUpdateAction
type OrderUpdateItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

// OrderUpdateSyncInfoAction implements the interface OrderUpdateAction
type OrderUpdateSyncInfoAction struct {
	SyncedAt   *time.Time                 `json:"syncedAt,omitempty"`
	ExternalID string                     `json:"externalId,omitempty"`
	Channel    *ChannelResourceIdentifier `json:"channel"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderUpdateSyncInfoAction) MarshalJSON() ([]byte, error) {
	type Alias OrderUpdateSyncInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateSyncInfo", Alias: (*Alias)(&obj)})
}

// Parcel is a standalone struct
type Parcel struct {
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	Items        []DeliveryItem      `json:"items,omitempty"`
	ID           string              `json:"id"`
	CreatedAt    time.Time           `json:"createdAt"`
}

// ParcelDraft is a standalone struct
type ParcelDraft struct {
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	Items        []DeliveryItem      `json:"items,omitempty"`
}

// ParcelMeasurements is a standalone struct
type ParcelMeasurements struct {
	WidthInMillimeter  float64 `json:"widthInMillimeter,omitempty"`
	WeightInGram       float64 `json:"weightInGram,omitempty"`
	LengthInMillimeter float64 `json:"lengthInMillimeter,omitempty"`
	HeightInMillimeter float64 `json:"heightInMillimeter,omitempty"`
}

// PaymentInfo is a standalone struct
type PaymentInfo struct {
	Payments []PaymentReference `json:"payments"`
}

// ProductVariantImportDraft is a standalone struct
type ProductVariantImportDraft struct {
	SKU        string       `json:"sku,omitempty"`
	Prices     []PriceDraft `json:"prices,omitempty"`
	Images     []Image      `json:"images,omitempty"`
	ID         int          `json:"id,omitempty"`
	Attributes []Attribute  `json:"attributes,omitempty"`
}

// ReturnInfo is a standalone struct
type ReturnInfo struct {
	ReturnTrackingID string       `json:"returnTrackingId,omitempty"`
	ReturnDate       *time.Time   `json:"returnDate,omitempty"`
	Items            []ReturnItem `json:"items"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReturnInfo) UnmarshalJSON(data []byte) error {
	type Alias ReturnInfo
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Items {
		var err error
		obj.Items[i], err = mapDiscriminatorReturnItem(obj.Items[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// ReturnItemDraft is a standalone struct
type ReturnItemDraft struct {
	ShipmentState    ReturnShipmentState `json:"shipmentState"`
	Quantity         int                 `json:"quantity"`
	LineItemID       string              `json:"lineItemId,omitempty"`
	CustomLineItemID string              `json:"customLineItemId,omitempty"`
	Comment          string              `json:"comment,omitempty"`
}

// ShippingInfoImportDraft is a standalone struct
type ShippingInfoImportDraft struct {
	TaxRate             *TaxRate                          `json:"taxRate,omitempty"`
	TaxCategory         *TaxCategoryResourceIdentifier    `json:"taxCategory,omitempty"`
	ShippingRate        *ShippingRateDraft                `json:"shippingRate"`
	ShippingMethodState ShippingMethodState               `json:"shippingMethodState,omitempty"`
	ShippingMethodName  string                            `json:"shippingMethodName"`
	ShippingMethod      *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	Price               *Money                            `json:"price"`
	DiscountedPrice     *DiscountedLineItemPriceDraft     `json:"discountedPrice,omitempty"`
	Deliveries          []Delivery                        `json:"deliveries,omitempty"`
}

// SyncInfo is a standalone struct
type SyncInfo struct {
	SyncedAt   time.Time         `json:"syncedAt"`
	ExternalID string            `json:"externalId,omitempty"`
	Channel    *ChannelReference `json:"channel"`
}

// TaxedItemPriceDraft is a standalone struct
type TaxedItemPriceDraft struct {
	TotalNet   *Money `json:"totalNet"`
	TotalGross *Money `json:"totalGross"`
}

// TrackingData is a standalone struct
type TrackingData struct {
	TrackingID          string `json:"trackingId,omitempty"`
	ProviderTransaction string `json:"providerTransaction,omitempty"`
	Provider            string `json:"provider,omitempty"`
	IsReturn            bool   `json:"isReturn"`
	Carrier             string `json:"carrier,omitempty"`
}
