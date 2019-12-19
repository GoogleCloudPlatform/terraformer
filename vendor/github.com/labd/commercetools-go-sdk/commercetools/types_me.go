// Automatically generated, do not edit

package commercetools

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

// MyCustomerDraft is a standalone struct
type MyCustomerDraft struct {
	VatID                  string        `json:"vatId,omitempty"`
	Title                  string        `json:"title,omitempty"`
	Password               string        `json:"password"`
	MiddleName             string        `json:"middleName,omitempty"`
	Locale                 string        `json:"locale,omitempty"`
	LastName               string        `json:"lastName,omitempty"`
	FirstName              string        `json:"firstName,omitempty"`
	Email                  string        `json:"email"`
	DefaultShippingAddress int           `json:"defaultShippingAddress,omitempty"`
	DefaultBillingAddress  int           `json:"defaultBillingAddress,omitempty"`
	DateOfBirth            Date          `json:"dateOfBirth,omitempty"`
	Custom                 *CustomFields `json:"custom,omitempty"`
	CompanyName            string        `json:"companyName,omitempty"`
	Addresses              []Address     `json:"addresses,omitempty"`
}

// MyLineItemDraft is a standalone struct
type MyLineItemDraft struct {
	VariantID           int                        `json:"variantId"`
	SupplyChannel       *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	ShippingDetails     *ItemShippingDetailsDraft  `json:"shippingDetails,omitempty"`
	Quantity            float64                    `json:"quantity"`
	ProductID           string                     `json:"productId"`
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	Custom              *CustomFieldsDraft         `json:"custom,omitempty"`
}

// MyOrderFromCartDraft is a standalone struct
type MyOrderFromCartDraft struct {
	Version int    `json:"version"`
	ID      string `json:"id"`
}
