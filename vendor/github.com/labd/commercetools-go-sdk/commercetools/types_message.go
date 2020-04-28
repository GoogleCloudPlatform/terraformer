// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// Message uses type as discriminator attribute
type Message interface{}

func mapDiscriminatorMessage(input interface{}) (Message, error) {
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
	case "CategoryCreated":
		new := CategoryCreatedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CategorySlugChanged":
		new := CategorySlugChangedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomLineItemStateTransition":
		new := CustomLineItemStateTransitionMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressAdded":
		new := CustomerAddressAddedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressChanged":
		new := CustomerAddressChangedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressRemoved":
		new := CustomerAddressRemovedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerCompanyNameSet":
		new := CustomerCompanyNameSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerCreated":
		new := CustomerCreatedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerDateOfBirthSet":
		new := CustomerDateOfBirthSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerEmailChanged":
		new := CustomerEmailChangedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerEmailVerified":
		new := CustomerEmailVerifiedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerGroupSet":
		new := CustomerGroupSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryAdded":
		new := DeliveryAddedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryAddressSet":
		new := DeliveryAddressSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryItemsUpdated":
		new := DeliveryItemsUpdatedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryRemoved":
		new := DeliveryRemovedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InventoryEntryDeleted":
		new := InventoryEntryDeletedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LineItemStateTransition":
		new := LineItemStateTransitionMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderBillingAddressSet":
		new := OrderBillingAddressSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCreated":
		new := OrderCreatedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomLineItemDiscountSet":
		new := OrderCustomLineItemDiscountSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerEmailSet":
		new := OrderCustomerEmailSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerGroupSet":
		new := OrderCustomerGroupSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerSet":
		new := OrderCustomerSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDeleted":
		new := OrderDeletedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeAdded":
		new := OrderDiscountCodeAddedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeRemoved":
		new := OrderDiscountCodeRemovedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeStateSet":
		new := OrderDiscountCodeStateSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderEditApplied":
		new := OrderEditAppliedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderImported":
		new := OrderImportedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderLineItemAdded":
		new := OrderLineItemAddedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderLineItemDiscountSet":
		new := OrderLineItemDiscountSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderPaymentStateChanged":
		new := OrderPaymentStateChangedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReturnInfoAdded":
		new := OrderReturnInfoAddedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderReturnShipmentStateChanged":
		new := OrderReturnShipmentStateChangedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShipmentStateChanged":
		new := OrderShipmentStateChangedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingAddressSet":
		new := OrderShippingAddressSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingInfoSet":
		new := OrderShippingInfoSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingRateInputSet":
		new := OrderShippingRateInputSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.OldShippingRateInput != nil {
			new.OldShippingRateInput, err = mapDiscriminatorShippingRateInput(new.OldShippingRateInput)
			if err != nil {
				return nil, err
			}
		}
		if new.ShippingRateInput != nil {
			new.ShippingRateInput, err = mapDiscriminatorShippingRateInput(new.ShippingRateInput)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "OrderStateChanged":
		new := OrderStateChangedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderStateTransition":
		new := OrderStateTransitionMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderStoreSet":
		new := OrderStoreSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelAddedToDelivery":
		new := ParcelAddedToDeliveryMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelItemsUpdated":
		new := ParcelItemsUpdatedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelMeasurementsUpdated":
		new := ParcelMeasurementsUpdatedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelRemovedFromDelivery":
		new := ParcelRemovedFromDeliveryMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelTrackingDataUpdated":
		new := ParcelTrackingDataUpdatedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentCreated":
		new := PaymentCreatedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentInteractionAdded":
		new := PaymentInteractionAddedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentStatusInterfaceCodeSet":
		new := PaymentStatusInterfaceCodeSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentStatusStateTransition":
		new := PaymentStatusStateTransitionMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentTransactionAdded":
		new := PaymentTransactionAddedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentTransactionStateChanged":
		new := PaymentTransactionStateChangedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductAddedToCategory":
		new := ProductAddedToCategoryMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductCreated":
		new := ProductCreatedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductDeleted":
		new := ProductDeletedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductImageAdded":
		new := ProductImageAddedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPriceDiscountsSet":
		new := ProductPriceDiscountsSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPriceExternalDiscountSet":
		new := ProductPriceExternalDiscountSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPublished":
		new := ProductPublishedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductRemovedFromCategory":
		new := ProductRemovedFromCategoryMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductRevertedStagedChanges":
		new := ProductRevertedStagedChangesMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductSlugChanged":
		new := ProductSlugChangedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductStateTransition":
		new := ProductStateTransitionMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductUnpublished":
		new := ProductUnpublishedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductVariantDeleted":
		new := ProductVariantDeletedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReviewCreated":
		new := ReviewCreatedMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReviewRatingSet":
		new := ReviewRatingSetMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.Target != nil {
			new.Target, err = mapDiscriminatorReference(new.Target)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "ReviewStateTransition":
		new := ReviewStateTransitionMessage{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.Target != nil {
			new.Target, err = mapDiscriminatorReference(new.Target)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	}
	return nil, nil
}

// MessagePayload uses type as discriminator attribute
type MessagePayload interface{}

func mapDiscriminatorMessagePayload(input interface{}) (MessagePayload, error) {
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
	case "CategoryCreated":
		new := CategoryCreatedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CategorySlugChanged":
		new := CategorySlugChangedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomLineItemStateTransition":
		new := CustomLineItemStateTransitionMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressAdded":
		new := CustomerAddressAddedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressChanged":
		new := CustomerAddressChangedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerAddressRemoved":
		new := CustomerAddressRemovedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerCompanyNameSet":
		new := CustomerCompanyNameSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerCreated":
		new := CustomerCreatedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerDateOfBirthSet":
		new := CustomerDateOfBirthSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerEmailChanged":
		new := CustomerEmailChangedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerEmailVerified":
		new := CustomerEmailVerifiedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "CustomerGroupSet":
		new := CustomerGroupSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryAdded":
		new := DeliveryAddedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryAddressSet":
		new := DeliveryAddressSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryItemsUpdated":
		new := DeliveryItemsUpdatedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DeliveryRemoved":
		new := DeliveryRemovedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "InventoryEntryDeleted":
		new := InventoryEntryDeletedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LineItemStateTransition":
		new := LineItemStateTransitionMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderBillingAddressSet":
		new := OrderBillingAddressSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCreated":
		new := OrderCreatedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomLineItemDiscountSet":
		new := OrderCustomLineItemDiscountSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerEmailSet":
		new := OrderCustomerEmailSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerGroupSet":
		new := OrderCustomerGroupSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderCustomerSet":
		new := OrderCustomerSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDeleted":
		new := OrderDeletedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeAdded":
		new := OrderDiscountCodeAddedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeRemoved":
		new := OrderDiscountCodeRemovedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderDiscountCodeStateSet":
		new := OrderDiscountCodeStateSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderEditApplied":
		new := OrderEditAppliedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderImported":
		new := OrderImportedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderLineItemAdded":
		new := OrderLineItemAddedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderLineItemDiscountSet":
		new := OrderLineItemDiscountSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderPaymentStateChanged":
		new := OrderPaymentStateChangedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReturnInfoAdded":
		new := OrderReturnInfoAddedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderReturnShipmentStateChanged":
		new := OrderReturnShipmentStateChangedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShipmentStateChanged":
		new := OrderShipmentStateChangedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingAddressSet":
		new := OrderShippingAddressSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingInfoSet":
		new := OrderShippingInfoSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderShippingRateInputSet":
		new := OrderShippingRateInputSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.OldShippingRateInput != nil {
			new.OldShippingRateInput, err = mapDiscriminatorShippingRateInput(new.OldShippingRateInput)
			if err != nil {
				return nil, err
			}
		}
		if new.ShippingRateInput != nil {
			new.ShippingRateInput, err = mapDiscriminatorShippingRateInput(new.ShippingRateInput)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "OrderStateChanged":
		new := OrderStateChangedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderStateTransition":
		new := OrderStateTransitionMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "OrderStoreSet":
		new := OrderStoreSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelAddedToDelivery":
		new := ParcelAddedToDeliveryMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelItemsUpdated":
		new := ParcelItemsUpdatedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelMeasurementsUpdated":
		new := ParcelMeasurementsUpdatedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelRemovedFromDelivery":
		new := ParcelRemovedFromDeliveryMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ParcelTrackingDataUpdated":
		new := ParcelTrackingDataUpdatedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentCreated":
		new := PaymentCreatedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentInteractionAdded":
		new := PaymentInteractionAddedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentStatusInterfaceCodeSet":
		new := PaymentStatusInterfaceCodeSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentStatusStateTransition":
		new := PaymentStatusStateTransitionMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentTransactionAdded":
		new := PaymentTransactionAddedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PaymentTransactionStateChanged":
		new := PaymentTransactionStateChangedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductAddedToCategory":
		new := ProductAddedToCategoryMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductCreated":
		new := ProductCreatedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductDeleted":
		new := ProductDeletedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductImageAdded":
		new := ProductImageAddedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPriceDiscountsSet":
		new := ProductPriceDiscountsSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPriceExternalDiscountSet":
		new := ProductPriceExternalDiscountSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductPublished":
		new := ProductPublishedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductRemovedFromCategory":
		new := ProductRemovedFromCategoryMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductRevertedStagedChanges":
		new := ProductRevertedStagedChangesMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductSlugChanged":
		new := ProductSlugChangedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductStateTransition":
		new := ProductStateTransitionMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductUnpublished":
		new := ProductUnpublishedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ProductVariantDeleted":
		new := ProductVariantDeletedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReviewCreated":
		new := ReviewCreatedMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReviewRatingSet":
		new := ReviewRatingSetMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.Target != nil {
			new.Target, err = mapDiscriminatorReference(new.Target)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "ReviewStateTransition":
		new := ReviewStateTransitionMessagePayload{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.Target != nil {
			new.Target, err = mapDiscriminatorReference(new.Target)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	}
	return nil, nil
}

// CategoryCreatedMessage implements the interface Message
type CategoryCreatedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Category                        *Category                `json:"category"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias CategoryCreatedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CategoryCreated", Alias: (*Alias)(&obj)})
}

// CategoryCreatedMessagePayload implements the interface MessagePayload
type CategoryCreatedMessagePayload struct {
	Category *Category `json:"category"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CategoryCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CategoryCreated", Alias: (*Alias)(&obj)})
}

// CategorySlugChangedMessage implements the interface Message
type CategorySlugChangedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Slug                            *LocalizedString         `json:"slug"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySlugChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias CategorySlugChangedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CategorySlugChanged", Alias: (*Alias)(&obj)})
}

// CategorySlugChangedMessagePayload implements the interface MessagePayload
type CategorySlugChangedMessagePayload struct {
	Slug *LocalizedString `json:"slug"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySlugChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CategorySlugChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CategorySlugChanged", Alias: (*Alias)(&obj)})
}

// CustomLineItemStateTransitionMessage implements the interface Message
type CustomLineItemStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	TransitionDate                  time.Time                `json:"transitionDate"`
	ToState                         *StateReference          `json:"toState"`
	Quantity                        int                      `json:"quantity"`
	FromState                       *StateReference          `json:"fromState"`
	CustomLineItemID                string                   `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomLineItemStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemStateTransitionMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomLineItemStateTransition", Alias: (*Alias)(&obj)})
}

// CustomLineItemStateTransitionMessagePayload implements the interface MessagePayload
type CustomLineItemStateTransitionMessagePayload struct {
	TransitionDate   time.Time       `json:"transitionDate"`
	ToState          *StateReference `json:"toState"`
	Quantity         int             `json:"quantity"`
	FromState        *StateReference `json:"fromState"`
	CustomLineItemID string          `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomLineItemStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomLineItemStateTransition", Alias: (*Alias)(&obj)})
}

// CustomerAddressAddedMessage implements the interface Message
type CustomerAddressAddedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Address                         *Address                 `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressAddedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressAdded", Alias: (*Alias)(&obj)})
}

// CustomerAddressAddedMessagePayload implements the interface MessagePayload
type CustomerAddressAddedMessagePayload struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressAdded", Alias: (*Alias)(&obj)})
}

// CustomerAddressChangedMessage implements the interface Message
type CustomerAddressChangedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Address                         *Address                 `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressChangedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressChanged", Alias: (*Alias)(&obj)})
}

// CustomerAddressChangedMessagePayload implements the interface MessagePayload
type CustomerAddressChangedMessagePayload struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressChanged", Alias: (*Alias)(&obj)})
}

// CustomerAddressRemovedMessage implements the interface Message
type CustomerAddressRemovedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Address                         *Address                 `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressRemovedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressRemoved", Alias: (*Alias)(&obj)})
}

// CustomerAddressRemovedMessagePayload implements the interface MessagePayload
type CustomerAddressRemovedMessagePayload struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddressRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddressRemovedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerAddressRemoved", Alias: (*Alias)(&obj)})
}

// CustomerCompanyNameSetMessage implements the interface Message
type CustomerCompanyNameSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	CompanyName                     string                   `json:"companyName"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerCompanyNameSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerCompanyNameSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerCompanyNameSet", Alias: (*Alias)(&obj)})
}

// CustomerCompanyNameSetMessagePayload implements the interface MessagePayload
type CustomerCompanyNameSetMessagePayload struct {
	CompanyName string `json:"companyName"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerCompanyNameSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerCompanyNameSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerCompanyNameSet", Alias: (*Alias)(&obj)})
}

// CustomerCreatedMessage implements the interface Message
type CustomerCreatedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Customer                        *Customer                `json:"customer"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerCreatedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerCreated", Alias: (*Alias)(&obj)})
}

// CustomerCreatedMessagePayload implements the interface MessagePayload
type CustomerCreatedMessagePayload struct {
	Customer *Customer `json:"customer"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerCreated", Alias: (*Alias)(&obj)})
}

// CustomerDateOfBirthSetMessage implements the interface Message
type CustomerDateOfBirthSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	DateOfBirth                     Date                     `json:"dateOfBirth"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerDateOfBirthSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerDateOfBirthSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerDateOfBirthSet", Alias: (*Alias)(&obj)})
}

// CustomerDateOfBirthSetMessagePayload implements the interface MessagePayload
type CustomerDateOfBirthSetMessagePayload struct {
	DateOfBirth Date `json:"dateOfBirth"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerDateOfBirthSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerDateOfBirthSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerDateOfBirthSet", Alias: (*Alias)(&obj)})
}

// CustomerEmailChangedMessage implements the interface Message
type CustomerEmailChangedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Email                           string                   `json:"email"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerEmailChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailChangedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerEmailChanged", Alias: (*Alias)(&obj)})
}

// CustomerEmailChangedMessagePayload implements the interface MessagePayload
type CustomerEmailChangedMessagePayload struct {
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerEmailChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerEmailChanged", Alias: (*Alias)(&obj)})
}

// CustomerEmailVerifiedMessage implements the interface Message
type CustomerEmailVerifiedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerEmailVerifiedMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailVerifiedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerEmailVerified", Alias: (*Alias)(&obj)})
}

// CustomerEmailVerifiedMessagePayload implements the interface MessagePayload
type CustomerEmailVerifiedMessagePayload struct{}

// MarshalJSON override to set the discriminator value
func (obj CustomerEmailVerifiedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerEmailVerifiedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerEmailVerified", Alias: (*Alias)(&obj)})
}

// CustomerGroupSetMessage implements the interface Message
type CustomerGroupSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	CustomerGroup                   *CustomerGroupReference  `json:"customerGroup"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetMessage) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerGroupSet", Alias: (*Alias)(&obj)})
}

// CustomerGroupSetMessagePayload implements the interface MessagePayload
type CustomerGroupSetMessagePayload struct {
	CustomerGroup *CustomerGroupReference `json:"customerGroup"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CustomerGroupSet", Alias: (*Alias)(&obj)})
}

// DeliveryAddedMessage implements the interface Message
type DeliveryAddedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Delivery                        *Delivery                `json:"delivery"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryAdded", Alias: (*Alias)(&obj)})
}

// DeliveryAddedMessagePayload implements the interface MessagePayload
type DeliveryAddedMessagePayload struct {
	Delivery *Delivery `json:"delivery"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryAdded", Alias: (*Alias)(&obj)})
}

// DeliveryAddressSetMessage implements the interface Message
type DeliveryAddressSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	OldAddress                      *Address                 `json:"oldAddress,omitempty"`
	DeliveryID                      string                   `json:"deliveryId"`
	Address                         *Address                 `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryAddressSetMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddressSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryAddressSet", Alias: (*Alias)(&obj)})
}

// DeliveryAddressSetMessagePayload implements the interface MessagePayload
type DeliveryAddressSetMessagePayload struct {
	OldAddress *Address `json:"oldAddress,omitempty"`
	DeliveryID string   `json:"deliveryId"`
	Address    *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryAddressSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryAddressSet", Alias: (*Alias)(&obj)})
}

// DeliveryItemsUpdatedMessage implements the interface Message
type DeliveryItemsUpdatedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	OldItems                        []DeliveryItem           `json:"oldItems"`
	Items                           []DeliveryItem           `json:"items"`
	DeliveryID                      string                   `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryItemsUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryItemsUpdatedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryItemsUpdated", Alias: (*Alias)(&obj)})
}

// DeliveryItemsUpdatedMessagePayload implements the interface MessagePayload
type DeliveryItemsUpdatedMessagePayload struct {
	OldItems   []DeliveryItem `json:"oldItems"`
	Items      []DeliveryItem `json:"items"`
	DeliveryID string         `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryItemsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryItemsUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryItemsUpdated", Alias: (*Alias)(&obj)})
}

// DeliveryRemovedMessage implements the interface Message
type DeliveryRemovedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Delivery                        *Delivery                `json:"delivery"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias DeliveryRemovedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryRemoved", Alias: (*Alias)(&obj)})
}

// DeliveryRemovedMessagePayload implements the interface MessagePayload
type DeliveryRemovedMessagePayload struct {
	Delivery *Delivery `json:"delivery"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias DeliveryRemovedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "DeliveryRemoved", Alias: (*Alias)(&obj)})
}

// InventoryEntryDeletedMessage implements the interface Message
type InventoryEntryDeletedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	SupplyChannel                   *ChannelReference        `json:"supplyChannel"`
	SKU                             string                   `json:"sku"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryDeletedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "InventoryEntryDeleted", Alias: (*Alias)(&obj)})
}

// InventoryEntryDeletedMessagePayload implements the interface MessagePayload
type InventoryEntryDeletedMessagePayload struct {
	SupplyChannel *ChannelReference `json:"supplyChannel"`
	SKU           string            `json:"sku"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "InventoryEntryDeleted", Alias: (*Alias)(&obj)})
}

// LineItemStateTransitionMessage implements the interface Message
type LineItemStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	TransitionDate                  time.Time                `json:"transitionDate"`
	ToState                         *StateReference          `json:"toState"`
	Quantity                        int                      `json:"quantity"`
	LineItemID                      string                   `json:"lineItemId"`
	FromState                       *StateReference          `json:"fromState"`
}

// MarshalJSON override to set the discriminator value
func (obj LineItemStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias LineItemStateTransitionMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "LineItemStateTransition", Alias: (*Alias)(&obj)})
}

// LineItemStateTransitionMessagePayload implements the interface MessagePayload
type LineItemStateTransitionMessagePayload struct {
	TransitionDate time.Time       `json:"transitionDate"`
	ToState        *StateReference `json:"toState"`
	Quantity       int             `json:"quantity"`
	LineItemID     string          `json:"lineItemId"`
	FromState      *StateReference `json:"fromState"`
}

// MarshalJSON override to set the discriminator value
func (obj LineItemStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias LineItemStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "LineItemStateTransition", Alias: (*Alias)(&obj)})
}

// MessageConfiguration is a standalone struct
type MessageConfiguration struct {
	Enabled                 bool    `json:"enabled"`
	DeleteDaysAfterCreation float64 `json:"deleteDaysAfterCreation,omitempty"`
}

// MessageConfigurationDraft is a standalone struct
type MessageConfigurationDraft struct {
	Enabled                 bool    `json:"enabled"`
	DeleteDaysAfterCreation float64 `json:"deleteDaysAfterCreation"`
}

// MessagePagedQueryResponse is a standalone struct
type MessagePagedQueryResponse struct {
	Total   int       `json:"total,omitempty"`
	Results []Message `json:"results"`
	Offset  int       `json:"offset"`
	Limit   int       `json:"limit"`
	Count   int       `json:"count"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MessagePagedQueryResponse) UnmarshalJSON(data []byte) error {
	type Alias MessagePagedQueryResponse
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Results {
		var err error
		obj.Results[i], err = mapDiscriminatorMessage(obj.Results[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// OrderBillingAddressSetMessage implements the interface Message
type OrderBillingAddressSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	OldAddress                      *Address                 `json:"oldAddress,omitempty"`
	Address                         *Address                 `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderBillingAddressSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderBillingAddressSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderBillingAddressSet", Alias: (*Alias)(&obj)})
}

// OrderBillingAddressSetMessagePayload implements the interface MessagePayload
type OrderBillingAddressSetMessagePayload struct {
	OldAddress *Address `json:"oldAddress,omitempty"`
	Address    *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderBillingAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderBillingAddressSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderBillingAddressSet", Alias: (*Alias)(&obj)})
}

// OrderCreatedMessage implements the interface Message
type OrderCreatedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Order                           *Order                   `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCreatedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCreated", Alias: (*Alias)(&obj)})
}

// OrderCreatedMessagePayload implements the interface MessagePayload
type OrderCreatedMessagePayload struct {
	Order *Order `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCreated", Alias: (*Alias)(&obj)})
}

// OrderCustomLineItemDiscountSetMessage implements the interface Message
type OrderCustomLineItemDiscountSetMessage struct {
	Version                         int                                  `json:"version"`
	SequenceNumber                  int                                  `json:"sequenceNumber"`
	ResourceVersion                 int                                  `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers             `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                            `json:"resource"`
	LastModifiedBy                  *LastModifiedBy                      `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                            `json:"lastModifiedAt"`
	ID                              string                               `json:"id"`
	CreatedBy                       *CreatedBy                           `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                            `json:"createdAt"`
	TaxedPrice                      *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	DiscountedPricePerQuantity      []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	CustomLineItemID                string                               `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomLineItemDiscountSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemDiscountSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

// OrderCustomLineItemDiscountSetMessagePayload implements the interface MessagePayload
type OrderCustomLineItemDiscountSetMessagePayload struct {
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	CustomLineItemID           string                               `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomLineItemDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomLineItemDiscountSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

// OrderCustomerEmailSetMessage implements the interface Message
type OrderCustomerEmailSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	OldEmail                        string                   `json:"oldEmail,omitempty"`
	Email                           string                   `json:"email,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerEmailSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerEmailSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomerEmailSet", Alias: (*Alias)(&obj)})
}

// OrderCustomerEmailSetMessagePayload implements the interface MessagePayload
type OrderCustomerEmailSetMessagePayload struct {
	OldEmail string `json:"oldEmail,omitempty"`
	Email    string `json:"email,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerEmailSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerEmailSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomerEmailSet", Alias: (*Alias)(&obj)})
}

// OrderCustomerGroupSetMessage implements the interface Message
type OrderCustomerGroupSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	OldCustomerGroup                *CustomerGroupReference  `json:"oldCustomerGroup,omitempty"`
	CustomerGroup                   *CustomerGroupReference  `json:"customerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerGroupSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerGroupSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomerGroupSet", Alias: (*Alias)(&obj)})
}

// OrderCustomerGroupSetMessagePayload implements the interface MessagePayload
type OrderCustomerGroupSetMessagePayload struct {
	OldCustomerGroup *CustomerGroupReference `json:"oldCustomerGroup,omitempty"`
	CustomerGroup    *CustomerGroupReference `json:"customerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerGroupSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerGroupSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomerGroupSet", Alias: (*Alias)(&obj)})
}

// OrderCustomerSetMessage implements the interface Message
type OrderCustomerSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	OldCustomerGroup                *CustomerGroupReference  `json:"oldCustomerGroup,omitempty"`
	OldCustomer                     *CustomerReference       `json:"oldCustomer,omitempty"`
	CustomerGroup                   *CustomerGroupReference  `json:"customerGroup,omitempty"`
	Customer                        *CustomerReference       `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomerSet", Alias: (*Alias)(&obj)})
}

// OrderCustomerSetMessagePayload implements the interface MessagePayload
type OrderCustomerSetMessagePayload struct {
	OldCustomerGroup *CustomerGroupReference `json:"oldCustomerGroup,omitempty"`
	OldCustomer      *CustomerReference      `json:"oldCustomer,omitempty"`
	CustomerGroup    *CustomerGroupReference `json:"customerGroup,omitempty"`
	Customer         *CustomerReference      `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderCustomerSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderCustomerSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderCustomerSet", Alias: (*Alias)(&obj)})
}

// OrderDeletedMessage implements the interface Message
type OrderDeletedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Order                           *Order                   `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDeletedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDeleted", Alias: (*Alias)(&obj)})
}

// OrderDeletedMessagePayload implements the interface MessagePayload
type OrderDeletedMessagePayload struct {
	Order *Order `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDeleted", Alias: (*Alias)(&obj)})
}

// OrderDiscountCodeAddedMessage implements the interface Message
type OrderDiscountCodeAddedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	DiscountCode                    *DiscountCodeReference   `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeAddedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeAdded", Alias: (*Alias)(&obj)})
}

// OrderDiscountCodeAddedMessagePayload implements the interface MessagePayload
type OrderDiscountCodeAddedMessagePayload struct {
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeAdded", Alias: (*Alias)(&obj)})
}

// OrderDiscountCodeRemovedMessage implements the interface Message
type OrderDiscountCodeRemovedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	DiscountCode                    *DiscountCodeReference   `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeRemovedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeRemovedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeRemoved", Alias: (*Alias)(&obj)})
}

// OrderDiscountCodeRemovedMessagePayload implements the interface MessagePayload
type OrderDiscountCodeRemovedMessagePayload struct {
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeRemovedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeRemovedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeRemoved", Alias: (*Alias)(&obj)})
}

// OrderDiscountCodeStateSetMessage implements the interface Message
type OrderDiscountCodeStateSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	State                           DiscountCodeState        `json:"state"`
	OldState                        DiscountCodeState        `json:"oldState,omitempty"`
	DiscountCode                    *DiscountCodeReference   `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeStateSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeStateSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeStateSet", Alias: (*Alias)(&obj)})
}

// OrderDiscountCodeStateSetMessagePayload implements the interface MessagePayload
type OrderDiscountCodeStateSetMessagePayload struct {
	State        DiscountCodeState      `json:"state"`
	OldState     DiscountCodeState      `json:"oldState,omitempty"`
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderDiscountCodeStateSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderDiscountCodeStateSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderDiscountCodeStateSet", Alias: (*Alias)(&obj)})
}

// OrderEditAppliedMessage implements the interface Message
type OrderEditAppliedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Result                          *OrderEditApplied        `json:"result"`
	Edit                            *OrderEditReference      `json:"edit"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditAppliedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderEditAppliedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderEditApplied", Alias: (*Alias)(&obj)})
}

// OrderEditAppliedMessagePayload implements the interface MessagePayload
type OrderEditAppliedMessagePayload struct {
	Result *OrderEditApplied   `json:"result"`
	Edit   *OrderEditReference `json:"edit"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditAppliedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderEditAppliedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderEditApplied", Alias: (*Alias)(&obj)})
}

// OrderImportedMessage implements the interface Message
type OrderImportedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Order                           *Order                   `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderImportedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderImportedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderImported", Alias: (*Alias)(&obj)})
}

// OrderImportedMessagePayload implements the interface MessagePayload
type OrderImportedMessagePayload struct {
	Order *Order `json:"order"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderImportedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderImportedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderImported", Alias: (*Alias)(&obj)})
}

// OrderLineItemAddedMessage implements the interface Message
type OrderLineItemAddedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LineItem                        *LineItem                `json:"lineItem"`
	AddedQuantity                   int                      `json:"addedQuantity"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderLineItemAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemAddedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderLineItemAdded", Alias: (*Alias)(&obj)})
}

// OrderLineItemAddedMessagePayload implements the interface MessagePayload
type OrderLineItemAddedMessagePayload struct {
	LineItem      *LineItem `json:"lineItem"`
	AddedQuantity int       `json:"addedQuantity"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderLineItemAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderLineItemAdded", Alias: (*Alias)(&obj)})
}

// OrderLineItemDiscountSetMessage implements the interface Message
type OrderLineItemDiscountSetMessage struct {
	Version                         int                                  `json:"version"`
	SequenceNumber                  int                                  `json:"sequenceNumber"`
	ResourceVersion                 int                                  `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers             `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                            `json:"resource"`
	LastModifiedBy                  *LastModifiedBy                      `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                            `json:"lastModifiedAt"`
	ID                              string                               `json:"id"`
	CreatedBy                       *CreatedBy                           `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                            `json:"createdAt"`
	TotalPrice                      *Money                               `json:"totalPrice"`
	TaxedPrice                      *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	LineItemID                      string                               `json:"lineItemId"`
	DiscountedPricePerQuantity      []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderLineItemDiscountSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemDiscountSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

// OrderLineItemDiscountSetMessagePayload implements the interface MessagePayload
type OrderLineItemDiscountSetMessagePayload struct {
	TotalPrice                 *Money                               `json:"totalPrice"`
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	LineItemID                 string                               `json:"lineItemId"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderLineItemDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderLineItemDiscountSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderLineItemDiscountSet", Alias: (*Alias)(&obj)})
}

// OrderPaymentStateChangedMessage implements the interface Message
type OrderPaymentStateChangedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	PaymentState                    PaymentState             `json:"paymentState"`
	OldPaymentState                 PaymentState             `json:"oldPaymentState,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderPaymentStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderPaymentStateChangedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderPaymentStateChanged", Alias: (*Alias)(&obj)})
}

// OrderPaymentStateChangedMessagePayload implements the interface MessagePayload
type OrderPaymentStateChangedMessagePayload struct {
	PaymentState    PaymentState `json:"paymentState"`
	OldPaymentState PaymentState `json:"oldPaymentState,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderPaymentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderPaymentStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderPaymentStateChanged", Alias: (*Alias)(&obj)})
}

// OrderReturnInfoAddedMessage implements the interface Message
type OrderReturnInfoAddedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	ReturnInfo                      *ReturnInfo              `json:"returnInfo"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderReturnInfoAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnInfoAddedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReturnInfoAdded", Alias: (*Alias)(&obj)})
}

// OrderReturnInfoAddedMessagePayload implements the interface MessagePayload
type OrderReturnInfoAddedMessagePayload struct {
	ReturnInfo *ReturnInfo `json:"returnInfo"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderReturnInfoAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnInfoAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReturnInfoAdded", Alias: (*Alias)(&obj)})
}

// OrderReturnShipmentStateChangedMessage implements the interface Message
type OrderReturnShipmentStateChangedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	ReturnShipmentState             ReturnShipmentState      `json:"returnShipmentState"`
	ReturnItemID                    string                   `json:"returnItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderReturnShipmentStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnShipmentStateChangedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderReturnShipmentStateChanged", Alias: (*Alias)(&obj)})
}

// OrderReturnShipmentStateChangedMessagePayload implements the interface MessagePayload
type OrderReturnShipmentStateChangedMessagePayload struct {
	ReturnShipmentState ReturnShipmentState `json:"returnShipmentState"`
	ReturnItemID        string              `json:"returnItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderReturnShipmentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderReturnShipmentStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderReturnShipmentStateChanged", Alias: (*Alias)(&obj)})
}

// OrderShipmentStateChangedMessage implements the interface Message
type OrderShipmentStateChangedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	ShipmentState                   ShipmentState            `json:"shipmentState"`
	OldShipmentState                ShipmentState            `json:"oldShipmentState,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShipmentStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShipmentStateChangedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShipmentStateChanged", Alias: (*Alias)(&obj)})
}

// OrderShipmentStateChangedMessagePayload implements the interface MessagePayload
type OrderShipmentStateChangedMessagePayload struct {
	ShipmentState    ShipmentState `json:"shipmentState"`
	OldShipmentState ShipmentState `json:"oldShipmentState,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShipmentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShipmentStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShipmentStateChanged", Alias: (*Alias)(&obj)})
}

// OrderShippingAddressSetMessage implements the interface Message
type OrderShippingAddressSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	OldAddress                      *Address                 `json:"oldAddress,omitempty"`
	Address                         *Address                 `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingAddressSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingAddressSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingAddressSet", Alias: (*Alias)(&obj)})
}

// OrderShippingAddressSetMessagePayload implements the interface MessagePayload
type OrderShippingAddressSetMessagePayload struct {
	OldAddress *Address `json:"oldAddress,omitempty"`
	Address    *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingAddressSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingAddressSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingAddressSet", Alias: (*Alias)(&obj)})
}

// OrderShippingInfoSetMessage implements the interface Message
type OrderShippingInfoSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	ShippingInfo                    *ShippingInfo            `json:"shippingInfo,omitempty"`
	OldShippingInfo                 *ShippingInfo            `json:"oldShippingInfo,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingInfoSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingInfoSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingInfoSet", Alias: (*Alias)(&obj)})
}

// OrderShippingInfoSetMessagePayload implements the interface MessagePayload
type OrderShippingInfoSetMessagePayload struct {
	ShippingInfo    *ShippingInfo `json:"shippingInfo,omitempty"`
	OldShippingInfo *ShippingInfo `json:"oldShippingInfo,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingInfoSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingInfoSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingInfoSet", Alias: (*Alias)(&obj)})
}

// OrderShippingRateInputSetMessage implements the interface Message
type OrderShippingRateInputSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	ShippingRateInput               ShippingRateInput        `json:"shippingRateInput,omitempty"`
	OldShippingRateInput            ShippingRateInput        `json:"oldShippingRateInput,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingRateInputSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingRateInputSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingRateInputSet", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderShippingRateInputSetMessage) UnmarshalJSON(data []byte) error {
	type Alias OrderShippingRateInputSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.OldShippingRateInput != nil {
		var err error
		obj.OldShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.OldShippingRateInput)
		if err != nil {
			return err
		}
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

// OrderShippingRateInputSetMessagePayload implements the interface MessagePayload
type OrderShippingRateInputSetMessagePayload struct {
	ShippingRateInput    ShippingRateInput `json:"shippingRateInput,omitempty"`
	OldShippingRateInput ShippingRateInput `json:"oldShippingRateInput,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShippingRateInputSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShippingRateInputSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShippingRateInputSet", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderShippingRateInputSetMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias OrderShippingRateInputSetMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.OldShippingRateInput != nil {
		var err error
		obj.OldShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.OldShippingRateInput)
		if err != nil {
			return err
		}
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

// OrderStateChangedMessage implements the interface Message
type OrderStateChangedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	OrderState                      OrderState               `json:"orderState"`
	OldOrderState                   OrderState               `json:"oldOrderState"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderStateChangedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderStateChanged", Alias: (*Alias)(&obj)})
}

// OrderStateChangedMessagePayload implements the interface MessagePayload
type OrderStateChangedMessagePayload struct {
	OrderState    OrderState `json:"orderState"`
	OldOrderState OrderState `json:"oldOrderState"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderStateChanged", Alias: (*Alias)(&obj)})
}

// OrderStateTransitionMessage implements the interface Message
type OrderStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	State                           *StateReference          `json:"state"`
	Force                           bool                     `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderStateTransitionMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderStateTransition", Alias: (*Alias)(&obj)})
}

// OrderStateTransitionMessagePayload implements the interface MessagePayload
type OrderStateTransitionMessagePayload struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderStateTransition", Alias: (*Alias)(&obj)})
}

// OrderStoreSetMessage implements the interface Message
type OrderStoreSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Store                           *StoreKeyReference       `json:"store"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderStoreSetMessage) MarshalJSON() ([]byte, error) {
	type Alias OrderStoreSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderStoreSet", Alias: (*Alias)(&obj)})
}

// OrderStoreSetMessagePayload implements the interface MessagePayload
type OrderStoreSetMessagePayload struct {
	Store *StoreKeyReference `json:"store"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderStoreSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderStoreSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderStoreSet", Alias: (*Alias)(&obj)})
}

// ParcelAddedToDeliveryMessage implements the interface Message
type ParcelAddedToDeliveryMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Parcel                          *Parcel                  `json:"parcel"`
	Delivery                        *Delivery                `json:"delivery"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelAddedToDeliveryMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelAddedToDeliveryMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelAddedToDelivery", Alias: (*Alias)(&obj)})
}

// ParcelAddedToDeliveryMessagePayload implements the interface MessagePayload
type ParcelAddedToDeliveryMessagePayload struct {
	Parcel   *Parcel   `json:"parcel"`
	Delivery *Delivery `json:"delivery"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelAddedToDeliveryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelAddedToDeliveryMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelAddedToDelivery", Alias: (*Alias)(&obj)})
}

// ParcelItemsUpdatedMessage implements the interface Message
type ParcelItemsUpdatedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	ParcelID                        string                   `json:"parcelId"`
	OldItems                        []DeliveryItem           `json:"oldItems"`
	Items                           []DeliveryItem           `json:"items"`
	DeliveryID                      string                   `json:"deliveryId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelItemsUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelItemsUpdatedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelItemsUpdated", Alias: (*Alias)(&obj)})
}

// ParcelItemsUpdatedMessagePayload implements the interface MessagePayload
type ParcelItemsUpdatedMessagePayload struct {
	ParcelID   string         `json:"parcelId"`
	OldItems   []DeliveryItem `json:"oldItems"`
	Items      []DeliveryItem `json:"items"`
	DeliveryID string         `json:"deliveryId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelItemsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelItemsUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelItemsUpdated", Alias: (*Alias)(&obj)})
}

// ParcelMeasurementsUpdatedMessage implements the interface Message
type ParcelMeasurementsUpdatedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	ParcelID                        string                   `json:"parcelId"`
	Measurements                    *ParcelMeasurements      `json:"measurements,omitempty"`
	DeliveryID                      string                   `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelMeasurementsUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelMeasurementsUpdatedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelMeasurementsUpdated", Alias: (*Alias)(&obj)})
}

// ParcelMeasurementsUpdatedMessagePayload implements the interface MessagePayload
type ParcelMeasurementsUpdatedMessagePayload struct {
	ParcelID     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	DeliveryID   string              `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelMeasurementsUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelMeasurementsUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelMeasurementsUpdated", Alias: (*Alias)(&obj)})
}

// ParcelRemovedFromDeliveryMessage implements the interface Message
type ParcelRemovedFromDeliveryMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Parcel                          *Parcel                  `json:"parcel"`
	DeliveryID                      string                   `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelRemovedFromDeliveryMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelRemovedFromDeliveryMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelRemovedFromDelivery", Alias: (*Alias)(&obj)})
}

// ParcelRemovedFromDeliveryMessagePayload implements the interface MessagePayload
type ParcelRemovedFromDeliveryMessagePayload struct {
	Parcel     *Parcel `json:"parcel"`
	DeliveryID string  `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelRemovedFromDeliveryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelRemovedFromDeliveryMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelRemovedFromDelivery", Alias: (*Alias)(&obj)})
}

// ParcelTrackingDataUpdatedMessage implements the interface Message
type ParcelTrackingDataUpdatedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	TrackingData                    *TrackingData            `json:"trackingData,omitempty"`
	ParcelID                        string                   `json:"parcelId"`
	DeliveryID                      string                   `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelTrackingDataUpdatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ParcelTrackingDataUpdatedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelTrackingDataUpdated", Alias: (*Alias)(&obj)})
}

// ParcelTrackingDataUpdatedMessagePayload implements the interface MessagePayload
type ParcelTrackingDataUpdatedMessagePayload struct {
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	ParcelID     string        `json:"parcelId"`
	DeliveryID   string        `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj ParcelTrackingDataUpdatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ParcelTrackingDataUpdatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ParcelTrackingDataUpdated", Alias: (*Alias)(&obj)})
}

// PaymentCreatedMessage implements the interface Message
type PaymentCreatedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Payment                         *Payment                 `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentCreatedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentCreated", Alias: (*Alias)(&obj)})
}

// PaymentCreatedMessagePayload implements the interface MessagePayload
type PaymentCreatedMessagePayload struct {
	Payment *Payment `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentCreated", Alias: (*Alias)(&obj)})
}

// PaymentInteractionAddedMessage implements the interface Message
type PaymentInteractionAddedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Interaction                     *CustomFields            `json:"interaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentInteractionAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentInteractionAddedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentInteractionAdded", Alias: (*Alias)(&obj)})
}

// PaymentInteractionAddedMessagePayload implements the interface MessagePayload
type PaymentInteractionAddedMessagePayload struct {
	Interaction *CustomFields `json:"interaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentInteractionAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentInteractionAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentInteractionAdded", Alias: (*Alias)(&obj)})
}

// PaymentStatusInterfaceCodeSetMessage implements the interface Message
type PaymentStatusInterfaceCodeSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	PaymentID                       string                   `json:"paymentId"`
	InterfaceCode                   string                   `json:"interfaceCode"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentStatusInterfaceCodeSetMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusInterfaceCodeSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentStatusInterfaceCodeSet", Alias: (*Alias)(&obj)})
}

// PaymentStatusInterfaceCodeSetMessagePayload implements the interface MessagePayload
type PaymentStatusInterfaceCodeSetMessagePayload struct {
	PaymentID     string `json:"paymentId"`
	InterfaceCode string `json:"interfaceCode"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentStatusInterfaceCodeSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusInterfaceCodeSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentStatusInterfaceCodeSet", Alias: (*Alias)(&obj)})
}

// PaymentStatusStateTransitionMessage implements the interface Message
type PaymentStatusStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	State                           *StateReference          `json:"state"`
	Force                           bool                     `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentStatusStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusStateTransitionMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentStatusStateTransition", Alias: (*Alias)(&obj)})
}

// PaymentStatusStateTransitionMessagePayload implements the interface MessagePayload
type PaymentStatusStateTransitionMessagePayload struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentStatusStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentStatusStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentStatusStateTransition", Alias: (*Alias)(&obj)})
}

// PaymentTransactionAddedMessage implements the interface Message
type PaymentTransactionAddedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Transaction                     *Transaction             `json:"transaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransactionAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionAddedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentTransactionAdded", Alias: (*Alias)(&obj)})
}

// PaymentTransactionAddedMessagePayload implements the interface MessagePayload
type PaymentTransactionAddedMessagePayload struct {
	Transaction *Transaction `json:"transaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransactionAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentTransactionAdded", Alias: (*Alias)(&obj)})
}

// PaymentTransactionStateChangedMessage implements the interface Message
type PaymentTransactionStateChangedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	TransactionID                   string                   `json:"transactionId"`
	State                           TransactionState         `json:"state"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransactionStateChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionStateChangedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentTransactionStateChanged", Alias: (*Alias)(&obj)})
}

// PaymentTransactionStateChangedMessagePayload implements the interface MessagePayload
type PaymentTransactionStateChangedMessagePayload struct {
	TransactionID string           `json:"transactionId"`
	State         TransactionState `json:"state"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransactionStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransactionStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PaymentTransactionStateChanged", Alias: (*Alias)(&obj)})
}

// ProductAddedToCategoryMessage implements the interface Message
type ProductAddedToCategoryMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Staged                          bool                     `json:"staged"`
	Category                        *CategoryReference       `json:"category"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddedToCategoryMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductAddedToCategoryMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductAddedToCategory", Alias: (*Alias)(&obj)})
}

// ProductAddedToCategoryMessagePayload implements the interface MessagePayload
type ProductAddedToCategoryMessagePayload struct {
	Staged   bool               `json:"staged"`
	Category *CategoryReference `json:"category"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddedToCategoryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductAddedToCategoryMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductAddedToCategory", Alias: (*Alias)(&obj)})
}

// ProductCreatedMessage implements the interface Message
type ProductCreatedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	ProductProjection               *ProductProjection       `json:"productProjection"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductCreatedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductCreated", Alias: (*Alias)(&obj)})
}

// ProductCreatedMessagePayload implements the interface MessagePayload
type ProductCreatedMessagePayload struct {
	ProductProjection *ProductProjection `json:"productProjection"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductCreated", Alias: (*Alias)(&obj)})
}

// ProductDeletedMessage implements the interface Message
type ProductDeletedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	RemovedImageUrls                []string                 `json:"removedImageUrls"`
	CurrentProjection               *ProductProjection       `json:"currentProjection"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductDeletedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductDeleted", Alias: (*Alias)(&obj)})
}

// ProductDeletedMessagePayload implements the interface MessagePayload
type ProductDeletedMessagePayload struct {
	RemovedImageUrls  []string           `json:"removedImageUrls"`
	CurrentProjection *ProductProjection `json:"currentProjection"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductDeleted", Alias: (*Alias)(&obj)})
}

// ProductImageAddedMessage implements the interface Message
type ProductImageAddedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	VariantID                       int                      `json:"variantId"`
	Staged                          bool                     `json:"staged"`
	Image                           *Image                   `json:"image"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductImageAddedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductImageAddedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductImageAdded", Alias: (*Alias)(&obj)})
}

// ProductImageAddedMessagePayload implements the interface MessagePayload
type ProductImageAddedMessagePayload struct {
	VariantID int    `json:"variantId"`
	Staged    bool   `json:"staged"`
	Image     *Image `json:"image"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductImageAddedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductImageAddedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductImageAdded", Alias: (*Alias)(&obj)})
}

// ProductPriceDiscountsSetMessage implements the interface Message
type ProductPriceDiscountsSetMessage struct {
	Version                         int                                    `json:"version"`
	SequenceNumber                  int                                    `json:"sequenceNumber"`
	ResourceVersion                 int                                    `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers               `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                              `json:"resource"`
	LastModifiedBy                  *LastModifiedBy                        `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                              `json:"lastModifiedAt"`
	ID                              string                                 `json:"id"`
	CreatedBy                       *CreatedBy                             `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                              `json:"createdAt"`
	UpdatedPrices                   []ProductPriceDiscountsSetUpdatedPrice `json:"updatedPrices"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPriceDiscountsSetMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductPriceDiscountsSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductPriceDiscountsSet", Alias: (*Alias)(&obj)})
}

// ProductPriceDiscountsSetMessagePayload implements the interface MessagePayload
type ProductPriceDiscountsSetMessagePayload struct {
	UpdatedPrices []ProductPriceDiscountsSetUpdatedPrice `json:"updatedPrices"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPriceDiscountsSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductPriceDiscountsSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductPriceDiscountsSet", Alias: (*Alias)(&obj)})
}

// ProductPriceDiscountsSetUpdatedPrice is a standalone struct
type ProductPriceDiscountsSetUpdatedPrice struct {
	VariantKey string           `json:"variantKey,omitempty"`
	VariantID  int              `json:"variantId"`
	Staged     bool             `json:"staged"`
	SKU        string           `json:"sku,omitempty"`
	PriceID    string           `json:"priceId"`
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
}

// ProductPriceExternalDiscountSetMessage implements the interface Message
type ProductPriceExternalDiscountSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	VariantKey                      string                   `json:"variantKey,omitempty"`
	VariantID                       int                      `json:"variantId"`
	Staged                          bool                     `json:"staged"`
	SKU                             string                   `json:"sku,omitempty"`
	PriceID                         string                   `json:"priceId"`
	Discounted                      *DiscountedPrice         `json:"discounted,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPriceExternalDiscountSetMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductPriceExternalDiscountSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductPriceExternalDiscountSet", Alias: (*Alias)(&obj)})
}

// ProductPriceExternalDiscountSetMessagePayload implements the interface MessagePayload
type ProductPriceExternalDiscountSetMessagePayload struct {
	VariantKey string           `json:"variantKey,omitempty"`
	VariantID  int              `json:"variantId"`
	Staged     bool             `json:"staged"`
	SKU        string           `json:"sku,omitempty"`
	PriceID    string           `json:"priceId"`
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPriceExternalDiscountSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductPriceExternalDiscountSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductPriceExternalDiscountSet", Alias: (*Alias)(&obj)})
}

// ProductPublishedMessage implements the interface Message
type ProductPublishedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Scope                           ProductPublishScope      `json:"scope"`
	RemovedImageUrls                []interface{}            `json:"removedImageUrls"`
	ProductProjection               *ProductProjection       `json:"productProjection"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPublishedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductPublishedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductPublished", Alias: (*Alias)(&obj)})
}

// ProductPublishedMessagePayload implements the interface MessagePayload
type ProductPublishedMessagePayload struct {
	Scope             ProductPublishScope `json:"scope"`
	RemovedImageUrls  []interface{}       `json:"removedImageUrls"`
	ProductProjection *ProductProjection  `json:"productProjection"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPublishedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductPublishedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductPublished", Alias: (*Alias)(&obj)})
}

// ProductRemovedFromCategoryMessage implements the interface Message
type ProductRemovedFromCategoryMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Staged                          bool                     `json:"staged"`
	Category                        *CategoryReference       `json:"category"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemovedFromCategoryMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductRemovedFromCategoryMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductRemovedFromCategory", Alias: (*Alias)(&obj)})
}

// ProductRemovedFromCategoryMessagePayload implements the interface MessagePayload
type ProductRemovedFromCategoryMessagePayload struct {
	Staged   bool               `json:"staged"`
	Category *CategoryReference `json:"category"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemovedFromCategoryMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductRemovedFromCategoryMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductRemovedFromCategory", Alias: (*Alias)(&obj)})
}

// ProductRevertedStagedChangesMessage implements the interface Message
type ProductRevertedStagedChangesMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	RemovedImageUrls                []string                 `json:"removedImageUrls"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRevertedStagedChangesMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertedStagedChangesMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductRevertedStagedChanges", Alias: (*Alias)(&obj)})
}

// ProductRevertedStagedChangesMessagePayload implements the interface MessagePayload
type ProductRevertedStagedChangesMessagePayload struct {
	RemovedImageUrls []string `json:"removedImageUrls"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRevertedStagedChangesMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertedStagedChangesMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductRevertedStagedChanges", Alias: (*Alias)(&obj)})
}

// ProductSlugChangedMessage implements the interface Message
type ProductSlugChangedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Slug                            *LocalizedString         `json:"slug"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSlugChangedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductSlugChangedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductSlugChanged", Alias: (*Alias)(&obj)})
}

// ProductSlugChangedMessagePayload implements the interface MessagePayload
type ProductSlugChangedMessagePayload struct {
	Slug *LocalizedString `json:"slug"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSlugChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductSlugChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductSlugChanged", Alias: (*Alias)(&obj)})
}

// ProductStateTransitionMessage implements the interface Message
type ProductStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	State                           *StateReference          `json:"state"`
	Force                           bool                     `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductStateTransitionMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductStateTransition", Alias: (*Alias)(&obj)})
}

// ProductStateTransitionMessagePayload implements the interface MessagePayload
type ProductStateTransitionMessagePayload struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductStateTransition", Alias: (*Alias)(&obj)})
}

// ProductUnpublishedMessage implements the interface Message
type ProductUnpublishedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductUnpublishedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductUnpublishedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductUnpublished", Alias: (*Alias)(&obj)})
}

// ProductUnpublishedMessagePayload implements the interface MessagePayload
type ProductUnpublishedMessagePayload struct{}

// MarshalJSON override to set the discriminator value
func (obj ProductUnpublishedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductUnpublishedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductUnpublished", Alias: (*Alias)(&obj)})
}

// ProductVariantDeletedMessage implements the interface Message
type ProductVariantDeletedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Variant                         *ProductVariant          `json:"variant"`
	RemovedImageUrls                []string                 `json:"removedImageUrls"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductVariantDeletedMessage) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantDeletedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductVariantDeleted", Alias: (*Alias)(&obj)})
}

// ProductVariantDeletedMessagePayload implements the interface MessagePayload
type ProductVariantDeletedMessagePayload struct {
	Variant          *ProductVariant `json:"variant"`
	RemovedImageUrls []string        `json:"removedImageUrls"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductVariantDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductVariantDeleted", Alias: (*Alias)(&obj)})
}

// ReviewCreatedMessage implements the interface Message
type ReviewCreatedMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Review                          *Review                  `json:"review"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewCreatedMessage) MarshalJSON() ([]byte, error) {
	type Alias ReviewCreatedMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewCreated", Alias: (*Alias)(&obj)})
}

// ReviewCreatedMessagePayload implements the interface MessagePayload
type ReviewCreatedMessagePayload struct {
	Review *Review `json:"review"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewCreatedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewCreatedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewCreated", Alias: (*Alias)(&obj)})
}

// ReviewRatingSetMessage implements the interface Message
type ReviewRatingSetMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Target                          Reference                `json:"target,omitempty"`
	OldRating                       float64                  `json:"oldRating,omitempty"`
	NewRating                       float64                  `json:"newRating,omitempty"`
	IncludedInStatistics            bool                     `json:"includedInStatistics"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewRatingSetMessage) MarshalJSON() ([]byte, error) {
	type Alias ReviewRatingSetMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewRatingSet", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewRatingSetMessage) UnmarshalJSON(data []byte) error {
	type Alias ReviewRatingSetMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorReference(obj.Target)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReviewRatingSetMessagePayload implements the interface MessagePayload
type ReviewRatingSetMessagePayload struct {
	Target               Reference `json:"target,omitempty"`
	OldRating            float64   `json:"oldRating,omitempty"`
	NewRating            float64   `json:"newRating,omitempty"`
	IncludedInStatistics bool      `json:"includedInStatistics"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewRatingSetMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewRatingSetMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewRatingSet", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewRatingSetMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias ReviewRatingSetMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorReference(obj.Target)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReviewStateTransitionMessage implements the interface Message
type ReviewStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LastModifiedBy                  *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedBy                       *CreatedBy               `json:"createdBy,omitempty"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Target                          Reference                `json:"target"`
	OldState                        *StateReference          `json:"oldState"`
	OldIncludedInStatistics         bool                     `json:"oldIncludedInStatistics"`
	NewState                        *StateReference          `json:"newState"`
	NewIncludedInStatistics         bool                     `json:"newIncludedInStatistics"`
	Force                           bool                     `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewStateTransitionMessage) MarshalJSON() ([]byte, error) {
	type Alias ReviewStateTransitionMessage
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewStateTransition", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewStateTransitionMessage) UnmarshalJSON(data []byte) error {
	type Alias ReviewStateTransitionMessage
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorReference(obj.Target)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReviewStateTransitionMessagePayload implements the interface MessagePayload
type ReviewStateTransitionMessagePayload struct {
	Target                  Reference       `json:"target"`
	OldState                *StateReference `json:"oldState"`
	OldIncludedInStatistics bool            `json:"oldIncludedInStatistics"`
	NewState                *StateReference `json:"newState"`
	NewIncludedInStatistics bool            `json:"newIncludedInStatistics"`
	Force                   bool            `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewStateTransitionMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ReviewStateTransitionMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ReviewStateTransition", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewStateTransitionMessagePayload) UnmarshalJSON(data []byte) error {
	type Alias ReviewStateTransitionMessagePayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorReference(obj.Target)
		if err != nil {
			return err
		}
	}

	return nil
}

// UserProvidedIdentifiers is a standalone struct
type UserProvidedIdentifiers struct {
	Slug           *LocalizedString `json:"slug,omitempty"`
	SKU            string           `json:"sku,omitempty"`
	OrderNumber    string           `json:"orderNumber,omitempty"`
	Key            string           `json:"key,omitempty"`
	ExternalID     string           `json:"externalId,omitempty"`
	CustomerNumber string           `json:"customerNumber,omitempty"`
}
