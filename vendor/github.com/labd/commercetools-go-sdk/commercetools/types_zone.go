// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ZoneUpdateAction uses action as discriminator attribute
type ZoneUpdateAction interface{}

func mapDiscriminatorZoneUpdateAction(input interface{}) (ZoneUpdateAction, error) {
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
	case "addLocation":
		new := ZoneAddLocationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := ZoneChangeNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeLocation":
		new := ZoneRemoveLocationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := ZoneSetDescriptionAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := ZoneSetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Location is a standalone struct
type Location struct {
	State   string      `json:"state,omitempty"`
	Country CountryCode `json:"country"`
}

// Zone is of type BaseResource
type Zone struct {
	Version        int        `json:"version"`
	LastModifiedAt time.Time  `json:"lastModifiedAt"`
	ID             string     `json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	Name           string     `json:"name"`
	Locations      []Location `json:"locations"`
	Key            string     `json:"key,omitempty"`
	Description    string     `json:"description,omitempty"`
}

// ZoneAddLocationAction implements the interface ZoneUpdateAction
type ZoneAddLocationAction struct {
	Location *Location `json:"location"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneAddLocationAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneAddLocationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLocation", Alias: (*Alias)(&obj)})
}

// ZoneChangeNameAction implements the interface ZoneUpdateAction
type ZoneChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ZoneDraft is a standalone struct
type ZoneDraft struct {
	Name        string     `json:"name"`
	Locations   []Location `json:"locations"`
	Key         string     `json:"key,omitempty"`
	Description string     `json:"description,omitempty"`
}

// ZonePagedQueryResponse is a standalone struct
type ZonePagedQueryResponse struct {
	Total   int    `json:"total,omitempty"`
	Results []Zone `json:"results"`
	Offset  int    `json:"offset"`
	Count   int    `json:"count"`
}

// ZoneReference implements the interface Reference
type ZoneReference struct {
	ID  string `json:"id"`
	Obj *Zone  `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneReference) MarshalJSON() ([]byte, error) {
	type Alias ZoneReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "zone", Alias: (*Alias)(&obj)})
}

// ZoneRemoveLocationAction implements the interface ZoneUpdateAction
type ZoneRemoveLocationAction struct {
	Location *Location `json:"location"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneRemoveLocationAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneRemoveLocationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLocation", Alias: (*Alias)(&obj)})
}

// ZoneResourceIdentifier implements the interface ResourceIdentifier
type ZoneResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ZoneResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "zone", Alias: (*Alias)(&obj)})
}

// ZoneSetDescriptionAction implements the interface ZoneUpdateAction
type ZoneSetDescriptionAction struct {
	Description string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// ZoneSetKeyAction implements the interface ZoneUpdateAction
type ZoneSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ZoneUpdate is a standalone struct
type ZoneUpdate struct {
	Version int                `json:"version"`
	Actions []ZoneUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ZoneUpdate) UnmarshalJSON(data []byte) error {
	type Alias ZoneUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorZoneUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
