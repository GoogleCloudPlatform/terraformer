// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// TransactionState is an enum type
type TransactionState string

// Enum values for TransactionState
const (
	TransactionStateInitial TransactionState = "Initial"
	TransactionStatePending TransactionState = "Pending"
	TransactionStateSuccess TransactionState = "Success"
	TransactionStateFailure TransactionState = "Failure"
)

// TransactionType is an enum type
type TransactionType string

// Enum values for TransactionType
const (
	TransactionTypeAuthorization       TransactionType = "Authorization"
	TransactionTypeCancelAuthorization TransactionType = "CancelAuthorization"
	TransactionTypeCharge              TransactionType = "Charge"
	TransactionTypeRefund              TransactionType = "Refund"
	TransactionTypeChargeback          TransactionType = "Chargeback"
)

// PaymentUpdateAction uses action as discriminator attribute
type PaymentUpdateAction interface{}

func mapDiscriminatorPaymentUpdateAction(input interface{}) (PaymentUpdateAction, error) {
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
	case "addInterfaceInteraction":
		new := PaymentAddInterfaceInteractionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addTransaction":
		new := PaymentAddTransactionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAmountPlanned":
		new := PaymentChangeAmountPlannedAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTransactionInteractionId":
		new := PaymentChangeTransactionInteractionIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTransactionState":
		new := PaymentChangeTransactionStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTransactionTimestamp":
		new := PaymentChangeTransactionTimestampAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAmountPaid":
		new := PaymentSetAmountPaidAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAmountRefunded":
		new := PaymentSetAmountRefundedAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAnonymousId":
		new := PaymentSetAnonymousIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAuthorization":
		new := PaymentSetAuthorizationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := PaymentSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := PaymentSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomer":
		new := PaymentSetCustomerAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setExternalId":
		new := PaymentSetExternalIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setInterfaceId":
		new := PaymentSetInterfaceIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := PaymentSetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoInterface":
		new := PaymentSetMethodInfoInterfaceAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoMethod":
		new := PaymentSetMethodInfoMethodAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoName":
		new := PaymentSetMethodInfoNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setStatusInterfaceCode":
		new := PaymentSetStatusInterfaceCodeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setStatusInterfaceText":
		new := PaymentSetStatusInterfaceTextAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionState":
		new := PaymentTransitionStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Payment is of type LoggedResource
type Payment struct {
	Version               int                `json:"version"`
	LastModifiedAt        time.Time          `json:"lastModifiedAt"`
	ID                    string             `json:"id"`
	CreatedAt             time.Time          `json:"createdAt"`
	LastModifiedBy        *LastModifiedBy    `json:"lastModifiedBy,omitempty"`
	CreatedBy             *CreatedBy         `json:"createdBy,omitempty"`
	Transactions          []Transaction      `json:"transactions"`
	PaymentStatus         *PaymentStatus     `json:"paymentStatus"`
	PaymentMethodInfo     *PaymentMethodInfo `json:"paymentMethodInfo"`
	Key                   string             `json:"key,omitempty"`
	InterfaceInteractions []CustomFields     `json:"interfaceInteractions"`
	InterfaceID           string             `json:"interfaceId,omitempty"`
	ExternalID            string             `json:"externalId,omitempty"`
	Customer              *CustomerReference `json:"customer,omitempty"`
	Custom                *CustomFields      `json:"custom,omitempty"`
	AuthorizedUntil       string             `json:"authorizedUntil,omitempty"`
	AnonymousID           string             `json:"anonymousId,omitempty"`
	AmountRefunded        TypedMoney         `json:"amountRefunded,omitempty"`
	AmountPlanned         TypedMoney         `json:"amountPlanned"`
	AmountPaid            TypedMoney         `json:"amountPaid,omitempty"`
	AmountAuthorized      TypedMoney         `json:"amountAuthorized,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Payment) UnmarshalJSON(data []byte) error {
	type Alias Payment
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.AmountAuthorized != nil {
		var err error
		obj.AmountAuthorized, err = mapDiscriminatorTypedMoney(obj.AmountAuthorized)
		if err != nil {
			return err
		}
	}
	if obj.AmountPaid != nil {
		var err error
		obj.AmountPaid, err = mapDiscriminatorTypedMoney(obj.AmountPaid)
		if err != nil {
			return err
		}
	}
	if obj.AmountPlanned != nil {
		var err error
		obj.AmountPlanned, err = mapDiscriminatorTypedMoney(obj.AmountPlanned)
		if err != nil {
			return err
		}
	}
	if obj.AmountRefunded != nil {
		var err error
		obj.AmountRefunded, err = mapDiscriminatorTypedMoney(obj.AmountRefunded)
		if err != nil {
			return err
		}
	}

	return nil
}

// PaymentAddInterfaceInteractionAction implements the interface PaymentUpdateAction
type PaymentAddInterfaceInteractionAction struct {
	Type   *TypeResourceIdentifier `json:"type"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentAddInterfaceInteractionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddInterfaceInteractionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addInterfaceInteraction", Alias: (*Alias)(&obj)})
}

// PaymentAddTransactionAction implements the interface PaymentUpdateAction
type PaymentAddTransactionAction struct {
	Transaction *TransactionDraft `json:"transaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentAddTransactionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddTransactionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTransaction", Alias: (*Alias)(&obj)})
}

// PaymentChangeAmountPlannedAction implements the interface PaymentUpdateAction
type PaymentChangeAmountPlannedAction struct {
	Amount *Money `json:"amount"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeAmountPlannedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeAmountPlannedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAmountPlanned", Alias: (*Alias)(&obj)})
}

// PaymentChangeTransactionInteractionIDAction implements the interface PaymentUpdateAction
type PaymentChangeTransactionInteractionIDAction struct {
	TransactionID string `json:"transactionId"`
	InteractionID string `json:"interactionId"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeTransactionInteractionIDAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionInteractionIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionInteractionId", Alias: (*Alias)(&obj)})
}

// PaymentChangeTransactionStateAction implements the interface PaymentUpdateAction
type PaymentChangeTransactionStateAction struct {
	TransactionID string           `json:"transactionId"`
	State         TransactionState `json:"state"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeTransactionStateAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionState", Alias: (*Alias)(&obj)})
}

// PaymentChangeTransactionTimestampAction implements the interface PaymentUpdateAction
type PaymentChangeTransactionTimestampAction struct {
	TransactionID string    `json:"transactionId"`
	Timestamp     time.Time `json:"timestamp"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeTransactionTimestampAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionTimestampAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionTimestamp", Alias: (*Alias)(&obj)})
}

// PaymentDraft is a standalone struct
type PaymentDraft struct {
	Transactions          []TransactionDraft          `json:"transactions,omitempty"`
	PaymentStatus         *PaymentStatusDraft         `json:"paymentStatus,omitempty"`
	PaymentMethodInfo     *PaymentMethodInfo          `json:"paymentMethodInfo,omitempty"`
	Key                   string                      `json:"key,omitempty"`
	InterfaceInteractions []CustomFieldsDraft         `json:"interfaceInteractions,omitempty"`
	InterfaceID           string                      `json:"interfaceId,omitempty"`
	ExternalID            string                      `json:"externalId,omitempty"`
	Customer              *CustomerResourceIdentifier `json:"customer,omitempty"`
	Custom                *CustomFieldsDraft          `json:"custom,omitempty"`
	AuthorizedUntil       string                      `json:"authorizedUntil,omitempty"`
	AnonymousID           string                      `json:"anonymousId,omitempty"`
	AmountRefunded        *Money                      `json:"amountRefunded,omitempty"`
	AmountPlanned         *Money                      `json:"amountPlanned"`
	AmountPaid            *Money                      `json:"amountPaid,omitempty"`
	AmountAuthorized      *Money                      `json:"amountAuthorized,omitempty"`
}

// PaymentMethodInfo is a standalone struct
type PaymentMethodInfo struct {
	PaymentInterface string           `json:"paymentInterface,omitempty"`
	Name             *LocalizedString `json:"name,omitempty"`
	Method           string           `json:"method,omitempty"`
}

// PaymentPagedQueryResponse is a standalone struct
type PaymentPagedQueryResponse struct {
	Total   int       `json:"total,omitempty"`
	Results []Payment `json:"results"`
	Offset  int       `json:"offset"`
	Count   int       `json:"count"`
}

// PaymentReference implements the interface Reference
type PaymentReference struct {
	ID  string   `json:"id"`
	Obj *Payment `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "payment", Alias: (*Alias)(&obj)})
}

// PaymentResourceIdentifier implements the interface ResourceIdentifier
type PaymentResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias PaymentResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "payment", Alias: (*Alias)(&obj)})
}

// PaymentSetAmountPaidAction implements the interface PaymentUpdateAction
type PaymentSetAmountPaidAction struct {
	Amount *Money `json:"amount,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAmountPaidAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountPaidAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountPaid", Alias: (*Alias)(&obj)})
}

// PaymentSetAmountRefundedAction implements the interface PaymentUpdateAction
type PaymentSetAmountRefundedAction struct {
	Amount *Money `json:"amount,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAmountRefundedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountRefundedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountRefunded", Alias: (*Alias)(&obj)})
}

// PaymentSetAnonymousIDAction implements the interface PaymentUpdateAction
type PaymentSetAnonymousIDAction struct {
	AnonymousID string `json:"anonymousId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAnonymousIDAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAnonymousIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

// PaymentSetAuthorizationAction implements the interface PaymentUpdateAction
type PaymentSetAuthorizationAction struct {
	Until  *time.Time `json:"until,omitempty"`
	Amount *Money     `json:"amount,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAuthorizationAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAuthorizationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthorization", Alias: (*Alias)(&obj)})
}

// PaymentSetCustomFieldAction implements the interface PaymentUpdateAction
type PaymentSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// PaymentSetCustomTypeAction implements the interface PaymentUpdateAction
type PaymentSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// PaymentSetCustomerAction implements the interface PaymentUpdateAction
type PaymentSetCustomerAction struct {
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

// PaymentSetExternalIDAction implements the interface PaymentUpdateAction
type PaymentSetExternalIDAction struct {
	ExternalID string `json:"externalId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetExternalIDAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetExternalIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

// PaymentSetInterfaceIDAction implements the interface PaymentUpdateAction
type PaymentSetInterfaceIDAction struct {
	InterfaceID string `json:"interfaceId"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetInterfaceIDAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetInterfaceIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setInterfaceId", Alias: (*Alias)(&obj)})
}

// PaymentSetKeyAction implements the interface PaymentUpdateAction
type PaymentSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// PaymentSetMethodInfoInterfaceAction implements the interface PaymentUpdateAction
type PaymentSetMethodInfoInterfaceAction struct {
	Interface string `json:"interface"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetMethodInfoInterfaceAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoInterfaceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoInterface", Alias: (*Alias)(&obj)})
}

// PaymentSetMethodInfoMethodAction implements the interface PaymentUpdateAction
type PaymentSetMethodInfoMethodAction struct {
	Method string `json:"method,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetMethodInfoMethodAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoMethod", Alias: (*Alias)(&obj)})
}

// PaymentSetMethodInfoNameAction implements the interface PaymentUpdateAction
type PaymentSetMethodInfoNameAction struct {
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetMethodInfoNameAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoName", Alias: (*Alias)(&obj)})
}

// PaymentSetStatusInterfaceCodeAction implements the interface PaymentUpdateAction
type PaymentSetStatusInterfaceCodeAction struct {
	InterfaceCode string `json:"interfaceCode,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetStatusInterfaceCodeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetStatusInterfaceCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatusInterfaceCode", Alias: (*Alias)(&obj)})
}

// PaymentSetStatusInterfaceTextAction implements the interface PaymentUpdateAction
type PaymentSetStatusInterfaceTextAction struct {
	InterfaceText string `json:"interfaceText"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetStatusInterfaceTextAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetStatusInterfaceTextAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatusInterfaceText", Alias: (*Alias)(&obj)})
}

// PaymentStatus is a standalone struct
type PaymentStatus struct {
	State         *StateReference `json:"state,omitempty"`
	InterfaceText string          `json:"interfaceText,omitempty"`
	InterfaceCode string          `json:"interfaceCode,omitempty"`
}

// PaymentStatusDraft is a standalone struct
type PaymentStatusDraft struct {
	State         *StateResourceIdentifier `json:"state,omitempty"`
	InterfaceText string                   `json:"interfaceText,omitempty"`
	InterfaceCode string                   `json:"interfaceCode,omitempty"`
}

// PaymentTransitionStateAction implements the interface PaymentUpdateAction
type PaymentTransitionStateAction struct {
	State *StateResourceIdentifier `json:"state"`
	Force bool                     `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

// PaymentUpdate is a standalone struct
type PaymentUpdate struct {
	Version int                   `json:"version"`
	Actions []PaymentUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentUpdate) UnmarshalJSON(data []byte) error {
	type Alias PaymentUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorPaymentUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// Transaction is a standalone struct
type Transaction struct {
	Type          TransactionType  `json:"type"`
	Timestamp     *time.Time       `json:"timestamp,omitempty"`
	State         TransactionState `json:"state,omitempty"`
	InteractionID string           `json:"interactionId,omitempty"`
	ID            string           `json:"id"`
	Amount        TypedMoney       `json:"amount"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Amount != nil {
		var err error
		obj.Amount, err = mapDiscriminatorTypedMoney(obj.Amount)
		if err != nil {
			return err
		}
	}

	return nil
}

// TransactionDraft is a standalone struct
type TransactionDraft struct {
	Type          TransactionType  `json:"type"`
	Timestamp     *time.Time       `json:"timestamp,omitempty"`
	State         TransactionState `json:"state,omitempty"`
	InteractionID string           `json:"interactionId,omitempty"`
	Amount        *Money           `json:"amount"`
}
