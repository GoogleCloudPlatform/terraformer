// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ReviewUpdateAction uses action as discriminator attribute
type ReviewUpdateAction interface{}

func mapDiscriminatorReviewUpdateAction(input interface{}) (ReviewUpdateAction, error) {
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
	case "setAuthorName":
		new := ReviewSetAuthorNameAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := ReviewSetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := ReviewSetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomer":
		new := ReviewSetCustomerAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := ReviewSetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocale":
		new := ReviewSetLocaleAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setRating":
		new := ReviewSetRatingAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTarget":
		new := ReviewSetTargetAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.Target != nil {
			new.Target, err = mapDiscriminatorResourceIdentifier(new.Target)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "setText":
		new := ReviewSetTextAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTitle":
		new := ReviewSetTitleAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionState":
		new := ReviewTransitionStateAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Review is of type LoggedResource
type Review struct {
	Version              int                `json:"version"`
	LastModifiedAt       time.Time          `json:"lastModifiedAt"`
	ID                   string             `json:"id"`
	CreatedAt            time.Time          `json:"createdAt"`
	LastModifiedBy       *LastModifiedBy    `json:"lastModifiedBy,omitempty"`
	CreatedBy            *CreatedBy         `json:"createdBy,omitempty"`
	UniquenessValue      string             `json:"uniquenessValue,omitempty"`
	Title                string             `json:"title,omitempty"`
	Text                 string             `json:"text,omitempty"`
	Target               Reference          `json:"target,omitempty"`
	State                *StateReference    `json:"state,omitempty"`
	Rating               float64            `json:"rating,omitempty"`
	Locale               string             `json:"locale,omitempty"`
	Key                  string             `json:"key,omitempty"`
	IncludedInStatistics bool               `json:"includedInStatistics"`
	Customer             *CustomerReference `json:"customer,omitempty"`
	Custom               *CustomFields      `json:"custom,omitempty"`
	AuthorName           string             `json:"authorName,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Review) UnmarshalJSON(data []byte) error {
	type Alias Review
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

// ReviewDraft is a standalone struct
type ReviewDraft struct {
	UniquenessValue string                      `json:"uniquenessValue,omitempty"`
	Title           string                      `json:"title,omitempty"`
	Text            string                      `json:"text,omitempty"`
	Target          ResourceIdentifier          `json:"target,omitempty"`
	State           *StateResourceIdentifier    `json:"state,omitempty"`
	Rating          float64                     `json:"rating,omitempty"`
	Locale          string                      `json:"locale,omitempty"`
	Key             string                      `json:"key,omitempty"`
	Customer        *CustomerResourceIdentifier `json:"customer,omitempty"`
	Custom          *CustomFieldsDraft          `json:"custom,omitempty"`
	AuthorName      string                      `json:"authorName,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewDraft) UnmarshalJSON(data []byte) error {
	type Alias ReviewDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorResourceIdentifier(obj.Target)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReviewPagedQueryResponse is a standalone struct
type ReviewPagedQueryResponse struct {
	Total   int      `json:"total,omitempty"`
	Results []Review `json:"results"`
	Offset  int      `json:"offset"`
	Count   int      `json:"count"`
}

// ReviewRatingStatistics is a standalone struct
type ReviewRatingStatistics struct {
	RatingsDistribution interface{} `json:"ratingsDistribution"`
	LowestRating        float64     `json:"lowestRating"`
	HighestRating       float64     `json:"highestRating"`
	Count               int         `json:"count"`
	AverageRating       float64     `json:"averageRating"`
}

// ReviewReference implements the interface Reference
type ReviewReference struct {
	ID  string  `json:"id"`
	Obj *Review `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewReference) MarshalJSON() ([]byte, error) {
	type Alias ReviewReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "review", Alias: (*Alias)(&obj)})
}

// ReviewResourceIdentifier implements the interface ResourceIdentifier
type ReviewResourceIdentifier struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ReviewResourceIdentifier
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "review", Alias: (*Alias)(&obj)})
}

// ReviewSetAuthorNameAction implements the interface ReviewUpdateAction
type ReviewSetAuthorNameAction struct {
	AuthorName string `json:"authorName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetAuthorNameAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetAuthorNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthorName", Alias: (*Alias)(&obj)})
}

// ReviewSetCustomFieldAction implements the interface ReviewUpdateAction
type ReviewSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// ReviewSetCustomTypeAction implements the interface ReviewUpdateAction
type ReviewSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// ReviewSetCustomerAction implements the interface ReviewUpdateAction
type ReviewSetCustomerAction struct {
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

// ReviewSetKeyAction implements the interface ReviewUpdateAction
type ReviewSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ReviewSetLocaleAction implements the interface ReviewUpdateAction
type ReviewSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

// ReviewSetRatingAction implements the interface ReviewUpdateAction
type ReviewSetRatingAction struct {
	Rating float64 `json:"rating,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetRatingAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetRatingAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRating", Alias: (*Alias)(&obj)})
}

// ReviewSetTargetAction implements the interface ReviewUpdateAction
type ReviewSetTargetAction struct {
	Target ResourceIdentifier `json:"target"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetTargetAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTargetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTarget", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewSetTargetAction) UnmarshalJSON(data []byte) error {
	type Alias ReviewSetTargetAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorResourceIdentifier(obj.Target)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReviewSetTextAction implements the interface ReviewUpdateAction
type ReviewSetTextAction struct {
	Text string `json:"text,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetTextAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTextAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setText", Alias: (*Alias)(&obj)})
}

// ReviewSetTitleAction implements the interface ReviewUpdateAction
type ReviewSetTitleAction struct {
	Title string `json:"title,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

// ReviewTransitionStateAction implements the interface ReviewUpdateAction
type ReviewTransitionStateAction struct {
	State *StateResourceIdentifier `json:"state"`
	Force bool                     `json:"force"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

// ReviewUpdate is a standalone struct
type ReviewUpdate struct {
	Version int                  `json:"version"`
	Actions []ReviewUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewUpdate) UnmarshalJSON(data []byte) error {
	type Alias ReviewUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorReviewUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
