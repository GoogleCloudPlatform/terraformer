// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// OrderEditResult uses type as discriminator attribute
type OrderEditResult interface{}

func mapDiscriminatorOrderEditResult(input interface{}) (OrderEditResult, error) {
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
	case "Applied":
		new := OrderEditApplied{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "NotProcessed":
		new := OrderEditNotProcessed{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "PreviewFailure":
		new := OrderEditPreviewFailure{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		for i := range new.Errors {
			new.Errors[i], err = mapDiscriminatorErrorObject(new.Errors[i])
		}
		return new, nil
	case "PreviewSuccess":
		new := OrderEditPreviewSuccess{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		for i := range new.MessagePayloads {
			new.MessagePayloads[i], err = mapDiscriminatorMessagePayload(new.MessagePayloads[i])
		}
		return new, nil
	}
	return nil, nil
}

// OrderEditUpdateAction uses action as discriminator attribute
type OrderEditUpdateAction interface{}

func mapDiscriminatorOrderEditUpdateAction(input interface{}) (OrderEditUpdateAction, error) {
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
	case "addStagedAction":
		new := OrderEditAddStagedActionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.StagedAction != nil {
			new.StagedAction, err = mapDiscriminatorStagedOrderUpdateAction(new.StagedAction)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "setComment":
		new := OrderEditSetCommentAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := OrderEditSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := OrderEditSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := OrderEditSetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setStagedActions":
		new := OrderEditSetStagedActionsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		for i := range new.StagedActions {
			new.StagedActions[i], err = mapDiscriminatorStagedOrderUpdateAction(new.StagedActions[i])
		}
		return new, nil
	}
	return nil, nil
}

// OrderEdit is of type LoggedResource
type OrderEdit struct {
	Version        int                       `json:"version"`
	ID             string                    `json:"id"`
	LastModifiedBy *LastModifiedBy           `json:"lastModifiedBy,omitempty"`
	CreatedBy      *CreatedBy                `json:"createdBy,omitempty"`
	StagedActions  []StagedOrderUpdateAction `json:"stagedActions"`
	Result         OrderEditResult           `json:"result"`
	Resource       *OrderReference           `json:"resource"`
	LastModifiedAt time.Time                 `json:"lastModifiedAt"`
	Key            string                    `json:"key,omitempty"`
	Custom         *CustomFields             `json:"custom,omitempty"`
	CreatedAt      time.Time                 `json:"createdAt"`
	Comment        string                    `json:"comment,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEdit) UnmarshalJSON(data []byte) error {
	type Alias OrderEdit
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Result != nil {
		var err error
		obj.Result, err = mapDiscriminatorOrderEditResult(obj.Result)
		if err != nil {
			return err
		}
	}
	for i := range obj.StagedActions {
		var err error
		obj.StagedActions[i], err = mapDiscriminatorStagedOrderUpdateAction(obj.StagedActions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// OrderEditAddStagedActionAction implements the interface OrderEditUpdateAction
type OrderEditAddStagedActionAction struct {
	StagedAction StagedOrderUpdateAction `json:"stagedAction"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditAddStagedActionAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditAddStagedActionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addStagedAction", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditAddStagedActionAction) UnmarshalJSON(data []byte) error {
	type Alias OrderEditAddStagedActionAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.StagedAction != nil {
		var err error
		obj.StagedAction, err = mapDiscriminatorStagedOrderUpdateAction(obj.StagedAction)
		if err != nil {
			return err
		}
	}

	return nil
}

// OrderEditApplied implements the interface OrderEditResult
type OrderEditApplied struct {
	ExcerptBeforeEdit *OrderExcerpt `json:"excerptBeforeEdit"`
	ExcerptAfterEdit  *OrderExcerpt `json:"excerptAfterEdit"`
	AppliedAt         time.Time     `json:"appliedAt"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditApplied) MarshalJSON() ([]byte, error) {
	type Alias OrderEditApplied
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "Applied", Alias: (*Alias)(&obj)})
}

// OrderEditApply is a standalone struct
type OrderEditApply struct {
	ResourceVersion int `json:"resourceVersion"`
	EditVersion     int `json:"editVersion"`
}

// OrderEditDraft is a standalone struct
type OrderEditDraft struct {
	StagedActions []StagedOrderUpdateAction `json:"stagedActions,omitempty"`
	Resource      *OrderReference           `json:"resource"`
	Key           string                    `json:"key,omitempty"`
	DryRun        bool                      `json:"dryRun"`
	Custom        *CustomFieldsDraft        `json:"custom,omitempty"`
	Comment       string                    `json:"comment,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditDraft) UnmarshalJSON(data []byte) error {
	type Alias OrderEditDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.StagedActions {
		var err error
		obj.StagedActions[i], err = mapDiscriminatorStagedOrderUpdateAction(obj.StagedActions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// OrderEditNotProcessed implements the interface OrderEditResult
type OrderEditNotProcessed struct{}

// MarshalJSON override to set the discriminator value
func (obj OrderEditNotProcessed) MarshalJSON() ([]byte, error) {
	type Alias OrderEditNotProcessed
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "NotProcessed", Alias: (*Alias)(&obj)})
}

// OrderEditPagedQueryResponse is a standalone struct
type OrderEditPagedQueryResponse struct {
	Total   int         `json:"total,omitempty"`
	Results []OrderEdit `json:"results"`
	Offset  int         `json:"offset"`
	Count   int         `json:"count"`
}

// OrderEditPreviewFailure implements the interface OrderEditResult
type OrderEditPreviewFailure struct {
	Errors []ErrorObject `json:"errors"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditPreviewFailure) MarshalJSON() ([]byte, error) {
	type Alias OrderEditPreviewFailure
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PreviewFailure", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditPreviewFailure) UnmarshalJSON(data []byte) error {
	type Alias OrderEditPreviewFailure
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

// OrderEditPreviewSuccess implements the interface OrderEditResult
type OrderEditPreviewSuccess struct {
	Preview         *StagedOrder     `json:"preview"`
	MessagePayloads []MessagePayload `json:"messagePayloads"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditPreviewSuccess) MarshalJSON() ([]byte, error) {
	type Alias OrderEditPreviewSuccess
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "PreviewSuccess", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditPreviewSuccess) UnmarshalJSON(data []byte) error {
	type Alias OrderEditPreviewSuccess
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.MessagePayloads {
		var err error
		obj.MessagePayloads[i], err = mapDiscriminatorMessagePayload(obj.MessagePayloads[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// OrderEditReference implements the interface Reference
type OrderEditReference struct {
	ID  string     `json:"id"`
	Obj *OrderEdit `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditReference) MarshalJSON() ([]byte, error) {
	type Alias OrderEditReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "order-edit", Alias: (*Alias)(&obj)})
}

// OrderEditResourceIdentifier implements the interface ResourceIdentifier
type OrderEditResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias OrderEditResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "order-edit", Alias: (*Alias)(&obj)})
}

// OrderEditSetCommentAction implements the interface OrderEditUpdateAction
type OrderEditSetCommentAction struct {
	Comment string `json:"comment,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditSetCommentAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetCommentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setComment", Alias: (*Alias)(&obj)})
}

// OrderEditSetCustomFieldAction implements the interface OrderEditUpdateAction
type OrderEditSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// OrderEditSetCustomTypeAction implements the interface OrderEditUpdateAction
type OrderEditSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields interface{}             `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// OrderEditSetKeyAction implements the interface OrderEditUpdateAction
type OrderEditSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// OrderEditSetStagedActionsAction implements the interface OrderEditUpdateAction
type OrderEditSetStagedActionsAction struct {
	StagedActions []StagedOrderUpdateAction `json:"stagedActions"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderEditSetStagedActionsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetStagedActionsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStagedActions", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditSetStagedActionsAction) UnmarshalJSON(data []byte) error {
	type Alias OrderEditSetStagedActionsAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.StagedActions {
		var err error
		obj.StagedActions[i], err = mapDiscriminatorStagedOrderUpdateAction(obj.StagedActions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// OrderEditUpdate is a standalone struct
type OrderEditUpdate struct {
	Version int                     `json:"version"`
	DryRun  bool                    `json:"dryRun"`
	Actions []OrderEditUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditUpdate) UnmarshalJSON(data []byte) error {
	type Alias OrderEditUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorOrderEditUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// OrderExcerpt is a standalone struct
type OrderExcerpt struct {
	Version    int         `json:"version"`
	TotalPrice *Money      `json:"totalPrice"`
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitempty"`
}

// StagedOrder is of type Order
type StagedOrder struct {
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

// StagedOrderAddCustomLineItemAction implements the interface StagedOrderUpdateAction
type StagedOrderAddCustomLineItemAction struct {
	TaxCategory     *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	Slug            string                         `json:"slug"`
	Quantity        float64                        `json:"quantity,omitempty"`
	Name            *LocalizedString               `json:"name"`
	Money           *Money                         `json:"money"`
	ExternalTaxRate *ExternalTaxRateDraft          `json:"externalTaxRate,omitempty"`
	Custom          *CustomFieldsDraft             `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderAddCustomLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddCustomLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addCustomLineItem", Alias: (*Alias)(&obj)})
}

// StagedOrderAddDeliveryAction implements the interface StagedOrderUpdateAction
type StagedOrderAddDeliveryAction struct {
	Parcels []ParcelDraft  `json:"parcels,omitempty"`
	Items   []DeliveryItem `json:"items,omitempty"`
	Address *Address       `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderAddDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDelivery", Alias: (*Alias)(&obj)})
}

// StagedOrderAddDiscountCodeAction implements the interface StagedOrderUpdateAction
type StagedOrderAddDiscountCodeAction struct {
	Code string `json:"code"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderAddDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDiscountCode", Alias: (*Alias)(&obj)})
}

// StagedOrderAddItemShippingAddressAction implements the interface StagedOrderUpdateAction
type StagedOrderAddItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addItemShippingAddress", Alias: (*Alias)(&obj)})
}

// StagedOrderAddLineItemAction implements the interface StagedOrderUpdateAction
type StagedOrderAddLineItemAction struct {
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
func (obj StagedOrderAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

// StagedOrderAddParcelToDeliveryAction implements the interface StagedOrderUpdateAction
type StagedOrderAddParcelToDeliveryAction struct {
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	Items        []DeliveryItem      `json:"items,omitempty"`
	DeliveryID   string              `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderAddParcelToDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddParcelToDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addParcelToDelivery", Alias: (*Alias)(&obj)})
}

// StagedOrderAddPaymentAction implements the interface StagedOrderUpdateAction
type StagedOrderAddPaymentAction struct {
	Payment *PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderAddPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPayment", Alias: (*Alias)(&obj)})
}

// StagedOrderAddReturnInfoAction implements the interface StagedOrderUpdateAction
type StagedOrderAddReturnInfoAction struct {
	ReturnTrackingID string            `json:"returnTrackingId,omitempty"`
	ReturnDate       *time.Time        `json:"returnDate,omitempty"`
	Items            []ReturnItemDraft `json:"items"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderAddReturnInfoAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddReturnInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addReturnInfo", Alias: (*Alias)(&obj)})
}

// StagedOrderAddShoppingListAction implements the interface StagedOrderUpdateAction
type StagedOrderAddShoppingListAction struct {
	SupplyChannel       *ChannelResourceIdentifier      `json:"supplyChannel,omitempty"`
	ShoppingList        *ShoppingListResourceIdentifier `json:"shoppingList"`
	DistributionChannel *ChannelResourceIdentifier      `json:"distributionChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderAddShoppingListAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddShoppingListAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShoppingList", Alias: (*Alias)(&obj)})
}

// StagedOrderChangeCustomLineItemMoneyAction implements the interface StagedOrderUpdateAction
type StagedOrderChangeCustomLineItemMoneyAction struct {
	Money            *Money `json:"money"`
	CustomLineItemID string `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderChangeCustomLineItemMoneyAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeCustomLineItemMoneyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomLineItemMoney", Alias: (*Alias)(&obj)})
}

// StagedOrderChangeCustomLineItemQuantityAction implements the interface StagedOrderUpdateAction
type StagedOrderChangeCustomLineItemQuantityAction struct {
	Quantity         float64 `json:"quantity"`
	CustomLineItemID string  `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderChangeCustomLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeCustomLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomLineItemQuantity", Alias: (*Alias)(&obj)})
}

// StagedOrderChangeLineItemQuantityAction implements the interface StagedOrderUpdateAction
type StagedOrderChangeLineItemQuantityAction struct {
	Quantity           float64                     `json:"quantity"`
	LineItemID         string                      `json:"lineItemId"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalPrice      *Money                      `json:"externalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

// StagedOrderChangeOrderStateAction implements the interface StagedOrderUpdateAction
type StagedOrderChangeOrderStateAction struct {
	OrderState OrderState `json:"orderState"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderChangeOrderStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeOrderStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeOrderState", Alias: (*Alias)(&obj)})
}

// StagedOrderChangePaymentStateAction implements the interface StagedOrderUpdateAction
type StagedOrderChangePaymentStateAction struct {
	PaymentState PaymentState `json:"paymentState,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderChangePaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangePaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePaymentState", Alias: (*Alias)(&obj)})
}

// StagedOrderChangeShipmentStateAction implements the interface StagedOrderUpdateAction
type StagedOrderChangeShipmentStateAction struct {
	ShipmentState ShipmentState `json:"shipmentState,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderChangeShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeShipmentState", Alias: (*Alias)(&obj)})
}

// StagedOrderChangeTaxCalculationModeAction implements the interface StagedOrderUpdateAction
type StagedOrderChangeTaxCalculationModeAction struct {
	TaxCalculationMode TaxCalculationMode `json:"taxCalculationMode"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderChangeTaxCalculationModeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeTaxCalculationModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxCalculationMode", Alias: (*Alias)(&obj)})
}

// StagedOrderChangeTaxModeAction implements the interface StagedOrderUpdateAction
type StagedOrderChangeTaxModeAction struct {
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderChangeTaxModeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeTaxModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxMode", Alias: (*Alias)(&obj)})
}

// StagedOrderChangeTaxRoundingModeAction implements the interface StagedOrderUpdateAction
type StagedOrderChangeTaxRoundingModeAction struct {
	TaxRoundingMode RoundingMode `json:"taxRoundingMode"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderChangeTaxRoundingModeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeTaxRoundingModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxRoundingMode", Alias: (*Alias)(&obj)})
}

// StagedOrderImportCustomLineItemStateAction implements the interface StagedOrderUpdateAction
type StagedOrderImportCustomLineItemStateAction struct {
	State            []ItemState `json:"state"`
	CustomLineItemID string      `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderImportCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderImportCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importCustomLineItemState", Alias: (*Alias)(&obj)})
}

// StagedOrderImportLineItemStateAction implements the interface StagedOrderUpdateAction
type StagedOrderImportLineItemStateAction struct {
	State      []ItemState `json:"state"`
	LineItemID string      `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderImportLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderImportLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importLineItemState", Alias: (*Alias)(&obj)})
}

// StagedOrderRemoveCustomLineItemAction implements the interface StagedOrderUpdateAction
type StagedOrderRemoveCustomLineItemAction struct {
	CustomLineItemID string `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderRemoveCustomLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveCustomLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeCustomLineItem", Alias: (*Alias)(&obj)})
}

// StagedOrderRemoveDeliveryAction implements the interface StagedOrderUpdateAction
type StagedOrderRemoveDeliveryAction struct {
	DeliveryID string `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderRemoveDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDelivery", Alias: (*Alias)(&obj)})
}

// StagedOrderRemoveDiscountCodeAction implements the interface StagedOrderUpdateAction
type StagedOrderRemoveDiscountCodeAction struct {
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderRemoveDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDiscountCode", Alias: (*Alias)(&obj)})
}

// StagedOrderRemoveItemShippingAddressAction implements the interface StagedOrderUpdateAction
type StagedOrderRemoveItemShippingAddressAction struct {
	AddressKey string `json:"addressKey"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderRemoveItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeItemShippingAddress", Alias: (*Alias)(&obj)})
}

// StagedOrderRemoveLineItemAction implements the interface StagedOrderUpdateAction
type StagedOrderRemoveLineItemAction struct {
	ShippingDetailsToRemove *ItemShippingDetailsDraft   `json:"shippingDetailsToRemove,omitempty"`
	Quantity                float64                     `json:"quantity,omitempty"`
	LineItemID              string                      `json:"lineItemId"`
	ExternalTotalPrice      *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalPrice           *Money                      `json:"externalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

// StagedOrderRemoveParcelFromDeliveryAction implements the interface StagedOrderUpdateAction
type StagedOrderRemoveParcelFromDeliveryAction struct {
	ParcelID string `json:"parcelId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderRemoveParcelFromDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveParcelFromDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeParcelFromDelivery", Alias: (*Alias)(&obj)})
}

// StagedOrderRemovePaymentAction implements the interface StagedOrderUpdateAction
type StagedOrderRemovePaymentAction struct {
	Payment *PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderRemovePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemovePaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePayment", Alias: (*Alias)(&obj)})
}

// StagedOrderSetBillingAddressAction implements the interface StagedOrderUpdateAction
type StagedOrderSetBillingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCountryAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCountryAction struct {
	Country string `json:"country,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCountryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCountryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCountry", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCustomFieldAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCustomLineItemCustomFieldAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCustomLineItemCustomFieldAction struct {
	Value            interface{} `json:"value,omitempty"`
	Name             string      `json:"name"`
	CustomLineItemID string      `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCustomLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomField", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCustomLineItemCustomTypeAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCustomLineItemCustomTypeAction struct {
	Type             *TypeResourceIdentifier `json:"type,omitempty"`
	Fields           *FieldContainer         `json:"fields,omitempty"`
	CustomLineItemID string                  `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCustomLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomType", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCustomLineItemShippingDetailsAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCustomLineItemShippingDetailsAction struct {
	ShippingDetails  *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	CustomLineItemID string                    `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCustomLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCustomLineItemTaxAmountAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCustomLineItemTaxAmountAction struct {
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
	CustomLineItemID  string                  `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCustomLineItemTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemTaxAmount", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCustomLineItemTaxRateAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCustomLineItemTaxRateAction struct {
	ExternalTaxRate  *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	CustomLineItemID string                `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCustomLineItemTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemTaxRate", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCustomShippingMethodAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCustomShippingMethodAction struct {
	TaxCategory        *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	ShippingRate       *ShippingRateDraft             `json:"shippingRate"`
	ShippingMethodName string                         `json:"shippingMethodName"`
	ExternalTaxRate    *ExternalTaxRateDraft          `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCustomShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomShippingMethod", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCustomTypeAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCustomerEmailAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCustomerEmailAction struct {
	Email string `json:"email,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCustomerEmailAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomerEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerEmail", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCustomerGroupAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCustomerGroupAction struct {
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCustomerGroupAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomerGroupAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerGroup", Alias: (*Alias)(&obj)})
}

// StagedOrderSetCustomerIDAction implements the interface StagedOrderUpdateAction
type StagedOrderSetCustomerIDAction struct {
	CustomerID string `json:"customerId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetCustomerIDAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomerIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerId", Alias: (*Alias)(&obj)})
}

// StagedOrderSetDeliveryAddressAction implements the interface StagedOrderUpdateAction
type StagedOrderSetDeliveryAddressAction struct {
	DeliveryID string   `json:"deliveryId"`
	Address    *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetDeliveryAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetDeliveryAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddress", Alias: (*Alias)(&obj)})
}

// StagedOrderSetDeliveryItemsAction implements the interface StagedOrderUpdateAction
type StagedOrderSetDeliveryItemsAction struct {
	Items      []DeliveryItem `json:"items"`
	DeliveryID string         `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetDeliveryItemsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetDeliveryItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryItems", Alias: (*Alias)(&obj)})
}

// StagedOrderSetLineItemCustomFieldAction implements the interface StagedOrderUpdateAction
type StagedOrderSetLineItemCustomFieldAction struct {
	Value      interface{} `json:"value,omitempty"`
	Name       string      `json:"name"`
	LineItemID string      `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

// StagedOrderSetLineItemCustomTypeAction implements the interface StagedOrderUpdateAction
type StagedOrderSetLineItemCustomTypeAction struct {
	Type       *TypeResourceIdentifier `json:"type,omitempty"`
	LineItemID string                  `json:"lineItemId"`
	Fields     *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

// StagedOrderSetLineItemPriceAction implements the interface StagedOrderUpdateAction
type StagedOrderSetLineItemPriceAction struct {
	LineItemID    string `json:"lineItemId"`
	ExternalPrice *Money `json:"externalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetLineItemPriceAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemPrice", Alias: (*Alias)(&obj)})
}

// StagedOrderSetLineItemShippingDetailsAction implements the interface StagedOrderUpdateAction
type StagedOrderSetLineItemShippingDetailsAction struct {
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	LineItemID      string                    `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

// StagedOrderSetLineItemTaxAmountAction implements the interface StagedOrderUpdateAction
type StagedOrderSetLineItemTaxAmountAction struct {
	LineItemID        string                  `json:"lineItemId"`
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetLineItemTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTaxAmount", Alias: (*Alias)(&obj)})
}

// StagedOrderSetLineItemTaxRateAction implements the interface StagedOrderUpdateAction
type StagedOrderSetLineItemTaxRateAction struct {
	LineItemID      string                `json:"lineItemId"`
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetLineItemTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTaxRate", Alias: (*Alias)(&obj)})
}

// StagedOrderSetLineItemTotalPriceAction implements the interface StagedOrderUpdateAction
type StagedOrderSetLineItemTotalPriceAction struct {
	LineItemID         string                      `json:"lineItemId"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetLineItemTotalPriceAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemTotalPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTotalPrice", Alias: (*Alias)(&obj)})
}

// StagedOrderSetLocaleAction implements the interface StagedOrderUpdateAction
type StagedOrderSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

// StagedOrderSetOrderNumberAction implements the interface StagedOrderUpdateAction
type StagedOrderSetOrderNumberAction struct {
	OrderNumber string `json:"orderNumber,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetOrderNumberAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetOrderNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setOrderNumber", Alias: (*Alias)(&obj)})
}

// StagedOrderSetOrderTotalTaxAction implements the interface StagedOrderUpdateAction
type StagedOrderSetOrderTotalTaxAction struct {
	ExternalTotalGross  *Money       `json:"externalTotalGross"`
	ExternalTaxPortions []TaxPortion `json:"externalTaxPortions,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetOrderTotalTaxAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetOrderTotalTaxAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setOrderTotalTax", Alias: (*Alias)(&obj)})
}

// StagedOrderSetParcelItemsAction implements the interface StagedOrderUpdateAction
type StagedOrderSetParcelItemsAction struct {
	ParcelID string         `json:"parcelId"`
	Items    []DeliveryItem `json:"items"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetParcelItemsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetParcelItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelItems", Alias: (*Alias)(&obj)})
}

// StagedOrderSetParcelMeasurementsAction implements the interface StagedOrderUpdateAction
type StagedOrderSetParcelMeasurementsAction struct {
	ParcelID     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetParcelMeasurementsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetParcelMeasurementsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelMeasurements", Alias: (*Alias)(&obj)})
}

// StagedOrderSetParcelTrackingDataAction implements the interface StagedOrderUpdateAction
type StagedOrderSetParcelTrackingDataAction struct {
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	ParcelID     string        `json:"parcelId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetParcelTrackingDataAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetParcelTrackingDataAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelTrackingData", Alias: (*Alias)(&obj)})
}

// StagedOrderSetReturnPaymentStateAction implements the interface StagedOrderUpdateAction
type StagedOrderSetReturnPaymentStateAction struct {
	ReturnItemID string             `json:"returnItemId"`
	PaymentState ReturnPaymentState `json:"paymentState"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetReturnPaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetReturnPaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnPaymentState", Alias: (*Alias)(&obj)})
}

// StagedOrderSetReturnShipmentStateAction implements the interface StagedOrderUpdateAction
type StagedOrderSetReturnShipmentStateAction struct {
	ShipmentState ReturnShipmentState `json:"shipmentState"`
	ReturnItemID  string              `json:"returnItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetReturnShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetReturnShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnShipmentState", Alias: (*Alias)(&obj)})
}

// StagedOrderSetShippingAddressAction implements the interface StagedOrderUpdateAction
type StagedOrderSetShippingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

// StagedOrderSetShippingAddressAndCustomShippingMethodAction implements the interface StagedOrderUpdateAction
type StagedOrderSetShippingAddressAndCustomShippingMethodAction struct {
	TaxCategory        *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	ShippingRate       *ShippingRateDraft             `json:"shippingRate"`
	ShippingMethodName string                         `json:"shippingMethodName"`
	ExternalTaxRate    *ExternalTaxRateDraft          `json:"externalTaxRate,omitempty"`
	Address            *Address                       `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetShippingAddressAndCustomShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingAddressAndCustomShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressAndCustomShippingMethod", Alias: (*Alias)(&obj)})
}

// StagedOrderSetShippingAddressAndShippingMethodAction implements the interface StagedOrderUpdateAction
type StagedOrderSetShippingAddressAndShippingMethodAction struct {
	ShippingMethod  *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft             `json:"externalTaxRate,omitempty"`
	Address         *Address                          `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetShippingAddressAndShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingAddressAndShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressAndShippingMethod", Alias: (*Alias)(&obj)})
}

// StagedOrderSetShippingMethodAction implements the interface StagedOrderUpdateAction
type StagedOrderSetShippingMethodAction struct {
	ShippingMethod  *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft             `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethod", Alias: (*Alias)(&obj)})
}

// StagedOrderSetShippingMethodTaxAmountAction implements the interface StagedOrderUpdateAction
type StagedOrderSetShippingMethodTaxAmountAction struct {
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetShippingMethodTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingMethodTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethodTaxAmount", Alias: (*Alias)(&obj)})
}

// StagedOrderSetShippingMethodTaxRateAction implements the interface StagedOrderUpdateAction
type StagedOrderSetShippingMethodTaxRateAction struct {
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetShippingMethodTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingMethodTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethodTaxRate", Alias: (*Alias)(&obj)})
}

// StagedOrderSetShippingRateInputAction implements the interface StagedOrderUpdateAction
type StagedOrderSetShippingRateInputAction struct {
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderSetShippingRateInputAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingRateInputAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingRateInput", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedOrderSetShippingRateInputAction) UnmarshalJSON(data []byte) error {
	type Alias StagedOrderSetShippingRateInputAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInputDraft(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}

	return nil
}

// StagedOrderTransitionCustomLineItemStateAction implements the interface StagedOrderUpdateAction
type StagedOrderTransitionCustomLineItemStateAction struct {
	ToState              *StateResourceIdentifier `json:"toState"`
	Quantity             int                      `json:"quantity"`
	FromState            *StateResourceIdentifier `json:"fromState"`
	CustomLineItemID     string                   `json:"customLineItemId"`
	ActualTransitionDate *time.Time               `json:"actualTransitionDate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderTransitionCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderTransitionCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionCustomLineItemState", Alias: (*Alias)(&obj)})
}

// StagedOrderTransitionLineItemStateAction implements the interface StagedOrderUpdateAction
type StagedOrderTransitionLineItemStateAction struct {
	ToState              *StateResourceIdentifier `json:"toState"`
	Quantity             int                      `json:"quantity"`
	LineItemID           string                   `json:"lineItemId"`
	FromState            *StateResourceIdentifier `json:"fromState"`
	ActualTransitionDate *time.Time               `json:"actualTransitionDate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderTransitionLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderTransitionLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionLineItemState", Alias: (*Alias)(&obj)})
}

// StagedOrderTransitionStateAction implements the interface StagedOrderUpdateAction
type StagedOrderTransitionStateAction struct {
	State *StateResourceIdentifier `json:"state"`
	Force bool                     `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

// StagedOrderUpdateItemShippingAddressAction implements the interface StagedOrderUpdateAction
type StagedOrderUpdateItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

// StagedOrderUpdateSyncInfoAction implements the interface StagedOrderUpdateAction
type StagedOrderUpdateSyncInfoAction struct {
	SyncedAt   *time.Time                 `json:"syncedAt,omitempty"`
	ExternalID string                     `json:"externalId,omitempty"`
	Channel    *ChannelResourceIdentifier `json:"channel"`
}

// MarshalJSON override to set the discriminator value
func (obj StagedOrderUpdateSyncInfoAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderUpdateSyncInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateSyncInfo", Alias: (*Alias)(&obj)})
}
