// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// AnonymousCartSignInMode is an enum type
type AnonymousCartSignInMode string

// Enum values for AnonymousCartSignInMode
const (
	AnonymousCartSignInModeMergeWithExistingCustomerCart AnonymousCartSignInMode = "MergeWithExistingCustomerCart"
	AnonymousCartSignInModeUseAsNewActiveCustomerCart    AnonymousCartSignInMode = "UseAsNewActiveCustomerCart"
)

// CustomerUpdateAction uses action as discriminator attribute
type CustomerUpdateAction interface{}

func mapDiscriminatorCustomerUpdateAction(input interface{}) (CustomerUpdateAction, error) {
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
		new := CustomerAddAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addBillingAddressId":
		new := CustomerAddBillingAddressIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addShippingAddressId":
		new := CustomerAddShippingAddressIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAddress":
		new := CustomerChangeAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeEmail":
		new := CustomerChangeEmailAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeAddress":
		new := CustomerRemoveAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeBillingAddressId":
		new := CustomerRemoveBillingAddressIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeShippingAddressId":
		new := CustomerRemoveShippingAddressIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCompanyName":
		new := CustomerSetCompanyNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := CustomerSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := CustomerSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerGroup":
		new := CustomerSetCustomerGroupAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerNumber":
		new := CustomerSetCustomerNumberAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDateOfBirth":
		new := CustomerSetDateOfBirthAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDefaultBillingAddress":
		new := CustomerSetDefaultBillingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDefaultShippingAddress":
		new := CustomerSetDefaultShippingAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setExternalId":
		new := CustomerSetExternalIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setFirstName":
		new := CustomerSetFirstNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := CustomerSetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLastName":
		new := CustomerSetLastNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocale":
		new := CustomerSetLocaleAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMiddleName":
		new := CustomerSetMiddleNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSalutation":
		new := CustomerSetSalutationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTitle":
		new := CustomerSetTitleAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setVatId":
		new := CustomerSetVatIDAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Customer is of type LoggedResource
type Customer struct {
	Version                  int                     `json:"version"`
	LastModifiedAt           time.Time               `json:"lastModifiedAt"`
	ID                       string                  `json:"id"`
	CreatedAt                time.Time               `json:"createdAt"`
	LastModifiedBy           *LastModifiedBy         `json:"lastModifiedBy,omitempty"`
	CreatedBy                *CreatedBy              `json:"createdBy,omitempty"`
	VatID                    string                  `json:"vatId,omitempty"`
	Title                    string                  `json:"title,omitempty"`
	ShippingAddressIds       []string                `json:"shippingAddressIds,omitempty"`
	Salutation               string                  `json:"salutation,omitempty"`
	Password                 string                  `json:"password"`
	MiddleName               string                  `json:"middleName,omitempty"`
	Locale                   string                  `json:"locale,omitempty"`
	LastName                 string                  `json:"lastName,omitempty"`
	Key                      string                  `json:"key,omitempty"`
	IsEmailVerified          bool                    `json:"isEmailVerified"`
	FirstName                string                  `json:"firstName,omitempty"`
	ExternalID               string                  `json:"externalId,omitempty"`
	Email                    string                  `json:"email"`
	DefaultShippingAddressID string                  `json:"defaultShippingAddressId,omitempty"`
	DefaultBillingAddressID  string                  `json:"defaultBillingAddressId,omitempty"`
	DateOfBirth              Date                    `json:"dateOfBirth,omitempty"`
	CustomerNumber           string                  `json:"customerNumber,omitempty"`
	CustomerGroup            *CustomerGroupReference `json:"customerGroup,omitempty"`
	Custom                   *CustomFields           `json:"custom,omitempty"`
	CompanyName              string                  `json:"companyName,omitempty"`
	BillingAddressIds        []string                `json:"billingAddressIds,omitempty"`
	Addresses                []Address               `json:"addresses"`
}

// CustomerAddAddressAction implements the interface CustomerUpdateAction
type CustomerAddAddressAction struct {
	Address *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAddress", Alias: (*Alias)(&obj)})
}

// CustomerAddBillingAddressIDAction implements the interface CustomerUpdateAction
type CustomerAddBillingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddBillingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddBillingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addBillingAddressId", Alias: (*Alias)(&obj)})
}

// CustomerAddShippingAddressIDAction implements the interface CustomerUpdateAction
type CustomerAddShippingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerAddShippingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerAddShippingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingAddressId", Alias: (*Alias)(&obj)})
}

// CustomerChangeAddressAction implements the interface CustomerUpdateAction
type CustomerChangeAddressAction struct {
	AddressID string   `json:"addressId"`
	Address   *Address `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerChangeAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerChangeAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAddress", Alias: (*Alias)(&obj)})
}

// CustomerChangeEmailAction implements the interface CustomerUpdateAction
type CustomerChangeEmailAction struct {
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerChangeEmailAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerChangeEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEmail", Alias: (*Alias)(&obj)})
}

// CustomerChangePassword is a standalone struct
type CustomerChangePassword struct {
	Version         int    `json:"version"`
	NewPassword     string `json:"newPassword"`
	ID              string `json:"id"`
	CurrentPassword string `json:"currentPassword"`
}

// CustomerCreateEmailToken is a standalone struct
type CustomerCreateEmailToken struct {
	Version    int    `json:"version,omitempty"`
	TTLMinutes int    `json:"ttlMinutes"`
	ID         string `json:"id"`
}

// CustomerCreatePasswordResetToken is a standalone struct
type CustomerCreatePasswordResetToken struct {
	TTLMinutes int    `json:"ttlMinutes,omitempty"`
	Email      string `json:"email"`
}

// CustomerDraft is a standalone struct
type CustomerDraft struct {
	VatID                  string                           `json:"vatId,omitempty"`
	Title                  string                           `json:"title,omitempty"`
	ShippingAddresses      []int                            `json:"shippingAddresses,omitempty"`
	Salutation             string                           `json:"salutation,omitempty"`
	Password               string                           `json:"password"`
	MiddleName             string                           `json:"middleName,omitempty"`
	Locale                 string                           `json:"locale,omitempty"`
	LastName               string                           `json:"lastName,omitempty"`
	Key                    string                           `json:"key,omitempty"`
	IsEmailVerified        bool                             `json:"isEmailVerified"`
	FirstName              string                           `json:"firstName,omitempty"`
	ExternalID             string                           `json:"externalId,omitempty"`
	Email                  string                           `json:"email"`
	DefaultShippingAddress int                              `json:"defaultShippingAddress,omitempty"`
	DefaultBillingAddress  int                              `json:"defaultBillingAddress,omitempty"`
	DateOfBirth            Date                             `json:"dateOfBirth,omitempty"`
	CustomerNumber         string                           `json:"customerNumber,omitempty"`
	CustomerGroup          *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	Custom                 *CustomFieldsDraft               `json:"custom,omitempty"`
	CompanyName            string                           `json:"companyName,omitempty"`
	BillingAddresses       []int                            `json:"billingAddresses,omitempty"`
	AnonymousID            string                           `json:"anonymousId,omitempty"`
	AnonymousCartID        string                           `json:"anonymousCartId,omitempty"`
	Addresses              []Address                        `json:"addresses,omitempty"`
}

// CustomerEmailVerify is a standalone struct
type CustomerEmailVerify struct {
	Version    int    `json:"version,omitempty"`
	TokenValue string `json:"tokenValue"`
}

// CustomerPagedQueryResponse is a standalone struct
type CustomerPagedQueryResponse struct {
	Total   int        `json:"total,omitempty"`
	Results []Customer `json:"results"`
	Offset  int        `json:"offset"`
	Count   int        `json:"count"`
}

// CustomerReference implements the interface Reference
type CustomerReference struct {
	ID  string    `json:"id"`
	Obj *Customer `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "customer", Alias: (*Alias)(&obj)})
}

// CustomerRemoveAddressAction implements the interface CustomerUpdateAction
type CustomerRemoveAddressAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerRemoveAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAddress", Alias: (*Alias)(&obj)})
}

// CustomerRemoveBillingAddressIDAction implements the interface CustomerUpdateAction
type CustomerRemoveBillingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerRemoveBillingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveBillingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeBillingAddressId", Alias: (*Alias)(&obj)})
}

// CustomerRemoveShippingAddressIDAction implements the interface CustomerUpdateAction
type CustomerRemoveShippingAddressIDAction struct {
	AddressID string `json:"addressId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerRemoveShippingAddressIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerRemoveShippingAddressIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingAddressId", Alias: (*Alias)(&obj)})
}

// CustomerResetPassword is a standalone struct
type CustomerResetPassword struct {
	Version     int    `json:"version,omitempty"`
	TokenValue  string `json:"tokenValue"`
	NewPassword string `json:"newPassword"`
}

// CustomerResourceIdentifier implements the interface ResourceIdentifier
type CustomerResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias CustomerResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "customer", Alias: (*Alias)(&obj)})
}

// CustomerSetCompanyNameAction implements the interface CustomerUpdateAction
type CustomerSetCompanyNameAction struct {
	CompanyName string `json:"companyName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetCompanyNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCompanyNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCompanyName", Alias: (*Alias)(&obj)})
}

// CustomerSetCustomFieldAction implements the interface CustomerUpdateAction
type CustomerSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// CustomerSetCustomTypeAction implements the interface CustomerUpdateAction
type CustomerSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// CustomerSetCustomerGroupAction implements the interface CustomerUpdateAction
type CustomerSetCustomerGroupAction struct {
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetCustomerGroupAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomerGroupAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerGroup", Alias: (*Alias)(&obj)})
}

// CustomerSetCustomerNumberAction implements the interface CustomerUpdateAction
type CustomerSetCustomerNumberAction struct {
	CustomerNumber string `json:"customerNumber,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetCustomerNumberAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetCustomerNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerNumber", Alias: (*Alias)(&obj)})
}

// CustomerSetDateOfBirthAction implements the interface CustomerUpdateAction
type CustomerSetDateOfBirthAction struct {
	DateOfBirth Date `json:"dateOfBirth,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetDateOfBirthAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDateOfBirthAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDateOfBirth", Alias: (*Alias)(&obj)})
}

// CustomerSetDefaultBillingAddressAction implements the interface CustomerUpdateAction
type CustomerSetDefaultBillingAddressAction struct {
	AddressID string `json:"addressId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetDefaultBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDefaultBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultBillingAddress", Alias: (*Alias)(&obj)})
}

// CustomerSetDefaultShippingAddressAction implements the interface CustomerUpdateAction
type CustomerSetDefaultShippingAddressAction struct {
	AddressID string `json:"addressId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetDefaultShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetDefaultShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefaultShippingAddress", Alias: (*Alias)(&obj)})
}

// CustomerSetExternalIDAction implements the interface CustomerUpdateAction
type CustomerSetExternalIDAction struct {
	ExternalID string `json:"externalId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetExternalIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetExternalIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

// CustomerSetFirstNameAction implements the interface CustomerUpdateAction
type CustomerSetFirstNameAction struct {
	FirstName string `json:"firstName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetFirstNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetFirstNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setFirstName", Alias: (*Alias)(&obj)})
}

// CustomerSetKeyAction implements the interface CustomerUpdateAction
type CustomerSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// CustomerSetLastNameAction implements the interface CustomerUpdateAction
type CustomerSetLastNameAction struct {
	LastName string `json:"lastName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetLastNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetLastNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLastName", Alias: (*Alias)(&obj)})
}

// CustomerSetLocaleAction implements the interface CustomerUpdateAction
type CustomerSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

// CustomerSetMiddleNameAction implements the interface CustomerUpdateAction
type CustomerSetMiddleNameAction struct {
	MiddleName string `json:"middleName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetMiddleNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetMiddleNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMiddleName", Alias: (*Alias)(&obj)})
}

// CustomerSetSalutationAction implements the interface CustomerUpdateAction
type CustomerSetSalutationAction struct {
	Salutation string `json:"salutation,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetSalutationAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetSalutationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSalutation", Alias: (*Alias)(&obj)})
}

// CustomerSetTitleAction implements the interface CustomerUpdateAction
type CustomerSetTitleAction struct {
	Title string `json:"title,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

// CustomerSetVatIDAction implements the interface CustomerUpdateAction
type CustomerSetVatIDAction struct {
	VatID string `json:"vatId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerSetVatIDAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerSetVatIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setVatId", Alias: (*Alias)(&obj)})
}

// CustomerSignInResult is a standalone struct
type CustomerSignInResult struct {
	Customer *Customer   `json:"customer"`
	Cart     interface{} `json:"cart,omitempty"`
}

// CustomerSignin is a standalone struct
type CustomerSignin struct {
	Password                string                  `json:"password"`
	Email                   string                  `json:"email"`
	AnonymousID             string                  `json:"anonymousId,omitempty"`
	AnonymousCartSignInMode AnonymousCartSignInMode `json:"anonymousCartSignInMode,omitempty"`
	AnonymousCartID         string                  `json:"anonymousCartId,omitempty"`
}

// CustomerToken is a standalone struct
type CustomerToken struct {
	Value          string     `json:"value"`
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
	ID             string     `json:"id"`
	ExpiresAt      time.Time  `json:"expiresAt"`
	CustomerID     string     `json:"customerId"`
	CreatedAt      time.Time  `json:"createdAt"`
}

// CustomerUpdate is a standalone struct
type CustomerUpdate struct {
	Version int                    `json:"version"`
	Actions []CustomerUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerUpdate) UnmarshalJSON(data []byte) error {
	type Alias CustomerUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorCustomerUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
