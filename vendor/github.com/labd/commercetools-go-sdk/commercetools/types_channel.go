// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ChannelRoleEnum is an enum type
type ChannelRoleEnum string

// Enum values for ChannelRoleEnum
const (
	ChannelRoleEnumInventorySupply     ChannelRoleEnum = "InventorySupply"
	ChannelRoleEnumProductDistribution ChannelRoleEnum = "ProductDistribution"
	ChannelRoleEnumOrderExport         ChannelRoleEnum = "OrderExport"
	ChannelRoleEnumOrderImport         ChannelRoleEnum = "OrderImport"
	ChannelRoleEnumPrimary             ChannelRoleEnum = "Primary"
)

// ChannelUpdateAction uses action as discriminator attribute
type ChannelUpdateAction interface{}

func mapDiscriminatorChannelUpdateAction(input interface{}) (ChannelUpdateAction, error) {
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
	case "addRoles":
		new := ChannelAddRolesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeDescription":
		new := ChannelChangeDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeKey":
		new := ChannelChangeKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := ChannelChangeNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeRoles":
		new := ChannelRemoveRolesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAddress":
		new := ChannelSetAddressAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := ChannelSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := ChannelSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setGeoLocation":
		new := ChannelSetGeoLocationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setRoles":
		new := ChannelSetRolesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Channel is of type LoggedResource
type Channel struct {
	Version                int                     `json:"version"`
	LastModifiedAt         time.Time               `json:"lastModifiedAt"`
	ID                     string                  `json:"id"`
	CreatedAt              time.Time               `json:"createdAt"`
	LastModifiedBy         *LastModifiedBy         `json:"lastModifiedBy,omitempty"`
	CreatedBy              *CreatedBy              `json:"createdBy,omitempty"`
	Roles                  []ChannelRoleEnum       `json:"roles"`
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
	Name                   *LocalizedString        `json:"name,omitempty"`
	Key                    string                  `json:"key"`
	GeoLocation            *GeoJSONPoint           `json:"geoLocation,omitempty"`
	Description            *LocalizedString        `json:"description,omitempty"`
	Custom                 *CustomFields           `json:"custom,omitempty"`
	Address                *Address                `json:"address,omitempty"`
}

// ChannelAddRolesAction implements the interface ChannelUpdateAction
type ChannelAddRolesAction struct {
	Roles []ChannelRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelAddRolesAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelAddRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addRoles", Alias: (*Alias)(&obj)})
}

// ChannelChangeDescriptionAction implements the interface ChannelUpdateAction
type ChannelChangeDescriptionAction struct {
	Description *LocalizedString `json:"description"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelChangeDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelChangeDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDescription", Alias: (*Alias)(&obj)})
}

// ChannelChangeKeyAction implements the interface ChannelUpdateAction
type ChannelChangeKeyAction struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelChangeKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelChangeKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeKey", Alias: (*Alias)(&obj)})
}

// ChannelChangeNameAction implements the interface ChannelUpdateAction
type ChannelChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ChannelDraft is a standalone struct
type ChannelDraft struct {
	Roles       []ChannelRoleEnum  `json:"roles,omitempty"`
	Name        *LocalizedString   `json:"name,omitempty"`
	Key         string             `json:"key"`
	GeoLocation *GeoJSONPoint      `json:"geoLocation,omitempty"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	Address     *Address           `json:"address,omitempty"`
}

// ChannelPagedQueryResponse is a standalone struct
type ChannelPagedQueryResponse struct {
	Total   int       `json:"total,omitempty"`
	Results []Channel `json:"results"`
	Offset  int       `json:"offset"`
	Count   int       `json:"count"`
}

// ChannelReference implements the interface Reference
type ChannelReference struct {
	ID  string   `json:"id"`
	Obj *Channel `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelReference) MarshalJSON() ([]byte, error) {
	type Alias ChannelReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "channel", Alias: (*Alias)(&obj)})
}

// ChannelRemoveRolesAction implements the interface ChannelUpdateAction
type ChannelRemoveRolesAction struct {
	Roles []ChannelRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelRemoveRolesAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelRemoveRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeRoles", Alias: (*Alias)(&obj)})
}

// ChannelResourceIdentifier implements the interface ResourceIdentifier
type ChannelResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ChannelResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "channel", Alias: (*Alias)(&obj)})
}

// ChannelSetAddressAction implements the interface ChannelUpdateAction
type ChannelSetAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelSetAddressAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddress", Alias: (*Alias)(&obj)})
}

// ChannelSetCustomFieldAction implements the interface ChannelUpdateAction
type ChannelSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// ChannelSetCustomTypeAction implements the interface ChannelUpdateAction
type ChannelSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// ChannelSetGeoLocationAction implements the interface ChannelUpdateAction
type ChannelSetGeoLocationAction struct {
	GeoLocation *GeoJSONPoint `json:"geoLocation,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelSetGeoLocationAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetGeoLocationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setGeoLocation", Alias: (*Alias)(&obj)})
}

// ChannelSetRolesAction implements the interface ChannelUpdateAction
type ChannelSetRolesAction struct {
	Roles []ChannelRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelSetRolesAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRoles", Alias: (*Alias)(&obj)})
}

// ChannelUpdate is a standalone struct
type ChannelUpdate struct {
	Version int                   `json:"version"`
	Actions []ChannelUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ChannelUpdate) UnmarshalJSON(data []byte) error {
	type Alias ChannelUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorChannelUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
