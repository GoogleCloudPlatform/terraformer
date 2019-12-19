// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

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

// CategoryCreatedMessage is of type Message
type CategoryCreatedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Category                        *Category                `json:"category"`
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

// CategorySlugChangedMessage is of type Message
type CategorySlugChangedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Slug                            *LocalizedString         `json:"slug"`
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

// CustomLineItemStateTransitionMessage is of type Message
type CustomLineItemStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	TransitionDate                  time.Time                `json:"transitionDate"`
	ToState                         *StateReference          `json:"toState"`
	Quantity                        int                      `json:"quantity"`
	FromState                       *StateReference          `json:"fromState"`
	CustomLineItemID                string                   `json:"customLineItemId"`
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

// CustomerAddressAddedMessage is of type Message
type CustomerAddressAddedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Address                         *Address                 `json:"address"`
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

// CustomerAddressChangedMessage is of type Message
type CustomerAddressChangedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Address                         *Address                 `json:"address"`
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

// CustomerAddressRemovedMessage is of type Message
type CustomerAddressRemovedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Address                         *Address                 `json:"address"`
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

// CustomerCompanyNameSetMessage is of type Message
type CustomerCompanyNameSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	CompanyName                     string                   `json:"companyName"`
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

// CustomerCreatedMessage is of type Message
type CustomerCreatedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Customer                        *Customer                `json:"customer"`
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

// CustomerDateOfBirthSetMessage is of type Message
type CustomerDateOfBirthSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	DateOfBirth                     Date                     `json:"dateOfBirth"`
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

// CustomerEmailChangedMessage is of type Message
type CustomerEmailChangedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Email                           string                   `json:"email"`
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

// CustomerEmailVerifiedMessage is of type Message
type CustomerEmailVerifiedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
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

// CustomerGroupSetMessage is of type Message
type CustomerGroupSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	CustomerGroup                   *CustomerGroupReference  `json:"customerGroup"`
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

// DeliveryAddedMessage is of type Message
type DeliveryAddedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Delivery                        *Delivery                `json:"delivery"`
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

// DeliveryAddressSetMessage is of type Message
type DeliveryAddressSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	OldAddress                      *Address                 `json:"oldAddress,omitempty"`
	DeliveryID                      string                   `json:"deliveryId"`
	Address                         *Address                 `json:"address,omitempty"`
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

// DeliveryItemsUpdatedMessage is of type Message
type DeliveryItemsUpdatedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	OldItems                        []DeliveryItem           `json:"oldItems"`
	Items                           []DeliveryItem           `json:"items"`
	DeliveryID                      string                   `json:"deliveryId"`
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

// DeliveryRemovedMessage is of type Message
type DeliveryRemovedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Delivery                        *Delivery                `json:"delivery"`
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

// InventoryEntryDeletedMessage is of type Message
type InventoryEntryDeletedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	SupplyChannel                   *ChannelReference        `json:"supplyChannel"`
	SKU                             string                   `json:"sku"`
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

// LineItemStateTransitionMessage is of type Message
type LineItemStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	TransitionDate                  time.Time                `json:"transitionDate"`
	ToState                         *StateReference          `json:"toState"`
	Quantity                        int                      `json:"quantity"`
	LineItemID                      string                   `json:"lineItemId"`
	FromState                       *StateReference          `json:"fromState"`
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

// Message is of type BaseResource
type Message struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}

	return nil
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
	Count   int       `json:"count"`
}

// OrderBillingAddressSetMessage is of type Message
type OrderBillingAddressSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	OldAddress                      *Address                 `json:"oldAddress,omitempty"`
	Address                         *Address                 `json:"address,omitempty"`
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

// OrderCreatedMessage is of type Message
type OrderCreatedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Order                           *Order                   `json:"order"`
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

// OrderCustomLineItemDiscountSetMessage is of type Message
type OrderCustomLineItemDiscountSetMessage struct {
	Version                         int                                  `json:"version"`
	LastModifiedAt                  time.Time                            `json:"lastModifiedAt"`
	ID                              string                               `json:"id"`
	CreatedAt                       time.Time                            `json:"createdAt"`
	Type                            string                               `json:"type"`
	SequenceNumber                  int                                  `json:"sequenceNumber"`
	ResourceVersion                 int                                  `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers             `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                            `json:"resource"`
	TaxedPrice                      *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	DiscountedPricePerQuantity      []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	CustomLineItemID                string                               `json:"customLineItemId"`
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

// OrderCustomerEmailSetMessage is of type Message
type OrderCustomerEmailSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	OldEmail                        string                   `json:"oldEmail,omitempty"`
	Email                           string                   `json:"email,omitempty"`
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

// OrderCustomerSetMessage is of type Message
type OrderCustomerSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	OldCustomerGroup                *CustomerGroupReference  `json:"oldCustomerGroup,omitempty"`
	OldCustomer                     *CustomerReference       `json:"oldCustomer,omitempty"`
	CustomerGroup                   *CustomerGroupReference  `json:"customerGroup,omitempty"`
	Customer                        *CustomerReference       `json:"customer,omitempty"`
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

// OrderDeletedMessage is of type Message
type OrderDeletedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Order                           *Order                   `json:"order"`
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

// OrderDiscountCodeAddedMessage is of type Message
type OrderDiscountCodeAddedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	DiscountCode                    *DiscountCodeReference   `json:"discountCode"`
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

// OrderDiscountCodeRemovedMessage is of type Message
type OrderDiscountCodeRemovedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	DiscountCode                    *DiscountCodeReference   `json:"discountCode"`
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

// OrderDiscountCodeStateSetMessage is of type Message
type OrderDiscountCodeStateSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	State                           DiscountCodeState        `json:"state"`
	OldState                        DiscountCodeState        `json:"oldState,omitempty"`
	DiscountCode                    *DiscountCodeReference   `json:"discountCode"`
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

// OrderEditAppliedMessage is of type Message
type OrderEditAppliedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Result                          *OrderEditApplied        `json:"result"`
	Edit                            *OrderEditReference      `json:"edit"`
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

// OrderImportedMessage is of type Message
type OrderImportedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Order                           *Order                   `json:"order"`
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

// OrderLineItemAddedMessage is of type Message
type OrderLineItemAddedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	LineItem                        *LineItem                `json:"lineItem"`
	AddedQuantity                   int                      `json:"addedQuantity"`
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

// OrderLineItemDiscountSetMessage is of type Message
type OrderLineItemDiscountSetMessage struct {
	Version                         int                                  `json:"version"`
	LastModifiedAt                  time.Time                            `json:"lastModifiedAt"`
	ID                              string                               `json:"id"`
	CreatedAt                       time.Time                            `json:"createdAt"`
	Type                            string                               `json:"type"`
	SequenceNumber                  int                                  `json:"sequenceNumber"`
	ResourceVersion                 int                                  `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers             `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                            `json:"resource"`
	TotalPrice                      *Money                               `json:"totalPrice"`
	TaxedPrice                      *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	LineItemID                      string                               `json:"lineItemId"`
	DiscountedPricePerQuantity      []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
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

// OrderPaymentStateChangedMessage is of type Message
type OrderPaymentStateChangedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	PaymentState                    PaymentState             `json:"paymentState"`
	OldPaymentState                 PaymentState             `json:"oldPaymentState"`
}

// OrderPaymentStateChangedMessagePayload implements the interface MessagePayload
type OrderPaymentStateChangedMessagePayload struct {
	PaymentState    PaymentState `json:"paymentState"`
	OldPaymentState PaymentState `json:"oldPaymentState"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderPaymentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderPaymentStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderPaymentStateChanged", Alias: (*Alias)(&obj)})
}

// OrderReturnInfoAddedMessage is of type Message
type OrderReturnInfoAddedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ReturnInfo                      *ReturnInfo              `json:"returnInfo"`
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

// OrderReturnShipmentStateChangedMessage is of type Message
type OrderReturnShipmentStateChangedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ReturnShipmentState             ReturnShipmentState      `json:"returnShipmentState"`
	ReturnItemID                    string                   `json:"returnItemId"`
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

// OrderShipmentStateChangedMessage is of type Message
type OrderShipmentStateChangedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ShipmentState                   ShipmentState            `json:"shipmentState"`
	OldShipmentState                ShipmentState            `json:"oldShipmentState"`
}

// OrderShipmentStateChangedMessagePayload implements the interface MessagePayload
type OrderShipmentStateChangedMessagePayload struct {
	ShipmentState    ShipmentState `json:"shipmentState"`
	OldShipmentState ShipmentState `json:"oldShipmentState"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderShipmentStateChangedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias OrderShipmentStateChangedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "OrderShipmentStateChanged", Alias: (*Alias)(&obj)})
}

// OrderShippingAddressSetMessage is of type Message
type OrderShippingAddressSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	OldAddress                      *Address                 `json:"oldAddress,omitempty"`
	Address                         *Address                 `json:"address,omitempty"`
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

// OrderShippingInfoSetMessage is of type Message
type OrderShippingInfoSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ShippingInfo                    *ShippingInfo            `json:"shippingInfo,omitempty"`
	OldShippingInfo                 *ShippingInfo            `json:"oldShippingInfo,omitempty"`
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

// OrderShippingRateInputSetMessage is of type Message
type OrderShippingRateInputSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ShippingRateInput               ShippingRateInput        `json:"shippingRateInput,omitempty"`
	OldShippingRateInput            ShippingRateInput        `json:"oldShippingRateInput,omitempty"`
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

// OrderStateChangedMessage is of type Message
type OrderStateChangedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	OrderState                      OrderState               `json:"orderState"`
	OldOrderState                   OrderState               `json:"oldOrderState"`
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

// OrderStateTransitionMessage is of type Message
type OrderStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	State                           *StateReference          `json:"state"`
	Force                           bool                     `json:"force"`
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

// ParcelAddedToDeliveryMessage is of type Message
type ParcelAddedToDeliveryMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Parcel                          *Parcel                  `json:"parcel"`
	Delivery                        *Delivery                `json:"delivery"`
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

// ParcelItemsUpdatedMessage is of type Message
type ParcelItemsUpdatedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ParcelID                        string                   `json:"parcelId"`
	OldItems                        []DeliveryItem           `json:"oldItems"`
	Items                           []DeliveryItem           `json:"items"`
	DeliveryID                      string                   `json:"deliveryId,omitempty"`
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

// ParcelMeasurementsUpdatedMessage is of type Message
type ParcelMeasurementsUpdatedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ParcelID                        string                   `json:"parcelId"`
	Measurements                    *ParcelMeasurements      `json:"measurements,omitempty"`
	DeliveryID                      string                   `json:"deliveryId"`
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

// ParcelRemovedFromDeliveryMessage is of type Message
type ParcelRemovedFromDeliveryMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Parcel                          *Parcel                  `json:"parcel"`
	DeliveryID                      string                   `json:"deliveryId"`
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

// ParcelTrackingDataUpdatedMessage is of type Message
type ParcelTrackingDataUpdatedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	TrackingData                    *TrackingData            `json:"trackingData,omitempty"`
	ParcelID                        string                   `json:"parcelId"`
	DeliveryID                      string                   `json:"deliveryId"`
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

// PaymentCreatedMessage is of type Message
type PaymentCreatedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Payment                         *Payment                 `json:"payment"`
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

// PaymentInteractionAddedMessage is of type Message
type PaymentInteractionAddedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Interaction                     *CustomFields            `json:"interaction"`
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

// PaymentStatusInterfaceCodeSetMessage is of type Message
type PaymentStatusInterfaceCodeSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	PaymentID                       string                   `json:"paymentId"`
	InterfaceCode                   string                   `json:"interfaceCode"`
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

// PaymentStatusStateTransitionMessage is of type Message
type PaymentStatusStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	State                           *StateReference          `json:"state"`
	Force                           bool                     `json:"force"`
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

// PaymentTransactionAddedMessage is of type Message
type PaymentTransactionAddedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Transaction                     *Transaction             `json:"transaction"`
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

// PaymentTransactionStateChangedMessage is of type Message
type PaymentTransactionStateChangedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	TransactionID                   string                   `json:"transactionId"`
	State                           TransactionState         `json:"state"`
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

// ProductCreatedMessage is of type Message
type ProductCreatedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ProductProjection               *ProductProjection       `json:"productProjection"`
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

// ProductDeletedMessage is of type Message
type ProductDeletedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	RemovedImageUrls                []string                 `json:"removedImageUrls"`
	CurrentProjection               *ProductProjection       `json:"currentProjection"`
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

// ProductImageAddedMessage is of type Message
type ProductImageAddedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	VariantID                       int                      `json:"variantId"`
	Staged                          bool                     `json:"staged"`
	Image                           *Image                   `json:"image"`
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

// ProductPriceDiscountsSetMessage is of type Message
type ProductPriceDiscountsSetMessage struct {
	Version                         int                                    `json:"version"`
	LastModifiedAt                  time.Time                              `json:"lastModifiedAt"`
	ID                              string                                 `json:"id"`
	CreatedAt                       time.Time                              `json:"createdAt"`
	Type                            string                                 `json:"type"`
	SequenceNumber                  int                                    `json:"sequenceNumber"`
	ResourceVersion                 int                                    `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers               `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                              `json:"resource"`
	UpdatedPrices                   []ProductPriceDiscountsSetUpdatedPrice `json:"updatedPrices"`
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

// ProductPriceExternalDiscountSetMessage is of type Message
type ProductPriceExternalDiscountSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	VariantKey                      string                   `json:"variantKey,omitempty"`
	VariantID                       int                      `json:"variantId"`
	Staged                          bool                     `json:"staged"`
	SKU                             string                   `json:"sku,omitempty"`
	PriceID                         string                   `json:"priceId"`
	Discounted                      *DiscountedPrice         `json:"discounted,omitempty"`
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

// ProductPublishedMessage is of type Message
type ProductPublishedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Scope                           ProductPublishScope      `json:"scope"`
	RemovedImageUrls                []interface{}            `json:"removedImageUrls"`
	ProductProjection               *ProductProjection       `json:"productProjection"`
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

// ProductRevertedStagedChangesMessage is of type Message
type ProductRevertedStagedChangesMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	RemovedImageUrls                []interface{}            `json:"removedImageUrls"`
}

// ProductRevertedStagedChangesMessagePayload implements the interface MessagePayload
type ProductRevertedStagedChangesMessagePayload struct {
	RemovedImageUrls []interface{} `json:"removedImageUrls"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRevertedStagedChangesMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertedStagedChangesMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductRevertedStagedChanges", Alias: (*Alias)(&obj)})
}

// ProductSlugChangedMessage is of type Message
type ProductSlugChangedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Slug                            *LocalizedString         `json:"slug"`
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

// ProductStateTransitionMessage is of type Message
type ProductStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	State                           *StateReference          `json:"state"`
	Force                           bool                     `json:"force"`
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

// ProductUnpublishedMessage is of type Message
type ProductUnpublishedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
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

// ProductVariantDeletedMessage is of type Message
type ProductVariantDeletedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Variant                         *ProductVariant          `json:"variant"`
	RemovedImageUrls                []interface{}            `json:"removedImageUrls"`
}

// ProductVariantDeletedMessagePayload implements the interface MessagePayload
type ProductVariantDeletedMessagePayload struct {
	Variant          *ProductVariant `json:"variant"`
	RemovedImageUrls []interface{}   `json:"removedImageUrls"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductVariantDeletedMessagePayload) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantDeletedMessagePayload
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "ProductVariantDeleted", Alias: (*Alias)(&obj)})
}

// ReviewCreatedMessage is of type Message
type ReviewCreatedMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Review                          *Review                  `json:"review"`
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

// ReviewRatingSetMessage is of type Message
type ReviewRatingSetMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Target                          Reference                `json:"target,omitempty"`
	OldRating                       float64                  `json:"oldRating,omitempty"`
	NewRating                       float64                  `json:"newRating,omitempty"`
	IncludedInStatistics            bool                     `json:"includedInStatistics"`
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

// ReviewStateTransitionMessage is of type Message
type ReviewStateTransitionMessage struct {
	Version                         int                      `json:"version"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
	Type                            string                   `json:"type"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	Target                          Reference                `json:"target"`
	OldState                        *StateReference          `json:"oldState"`
	OldIncludedInStatistics         bool                     `json:"oldIncludedInStatistics"`
	NewState                        *StateReference          `json:"newState"`
	NewIncludedInStatistics         bool                     `json:"newIncludedInStatistics"`
	Force                           bool                     `json:"force"`
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
