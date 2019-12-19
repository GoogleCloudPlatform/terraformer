// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ExtensionAction is an enum type
type ExtensionAction string

// Enum values for ExtensionAction
const (
	ExtensionActionCreate ExtensionAction = "Create"
	ExtensionActionUpdate ExtensionAction = "Update"
)

// ExtensionResourceTypeID is an enum type
type ExtensionResourceTypeID string

// Enum values for ExtensionResourceTypeID
const (
	ExtensionResourceTypeIDCart     ExtensionResourceTypeID = "cart"
	ExtensionResourceTypeIDOrder    ExtensionResourceTypeID = "order"
	ExtensionResourceTypeIDPayment  ExtensionResourceTypeID = "payment"
	ExtensionResourceTypeIDCustomer ExtensionResourceTypeID = "customer"
)

// ExtensionDestination uses type as discriminator attribute
type ExtensionDestination interface{}

func mapDiscriminatorExtensionDestination(input interface{}) (ExtensionDestination, error) {
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
	case "AWSLambda":
		new := ExtensionAWSLambdaDestination{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "HTTP":
		new := ExtensionHTTPDestination{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.Authentication != nil {
			new.Authentication, err = mapDiscriminatorExtensionHTTPDestinationAuthentication(new.Authentication)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	}
	return nil, nil
}

// ExtensionHTTPDestinationAuthentication uses type as discriminator attribute
type ExtensionHTTPDestinationAuthentication interface{}

func mapDiscriminatorExtensionHTTPDestinationAuthentication(input interface{}) (ExtensionHTTPDestinationAuthentication, error) {
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
	case "AuthorizationHeader":
		new := ExtensionAuthorizationHeaderAuthentication{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "AzureFunctions":
		new := ExtensionAzureFunctionsAuthentication{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// ExtensionUpdateAction uses action as discriminator attribute
type ExtensionUpdateAction interface{}

func mapDiscriminatorExtensionUpdateAction(input interface{}) (ExtensionUpdateAction, error) {
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
	case "changeDestination":
		new := ExtensionChangeDestinationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.Destination != nil {
			new.Destination, err = mapDiscriminatorExtensionDestination(new.Destination)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "changeTriggers":
		new := ExtensionChangeTriggersAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := ExtensionSetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTimeoutInMs":
		new := ExtensionSetTimeoutInMsAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Extension is of type LoggedResource
type Extension struct {
	Version        int                  `json:"version"`
	LastModifiedAt time.Time            `json:"lastModifiedAt"`
	ID             string               `json:"id"`
	CreatedAt      time.Time            `json:"createdAt"`
	LastModifiedBy *LastModifiedBy      `json:"lastModifiedBy,omitempty"`
	CreatedBy      *CreatedBy           `json:"createdBy,omitempty"`
	Triggers       []ExtensionTrigger   `json:"triggers"`
	TimeoutInMs    int                  `json:"timeoutInMs,omitempty"`
	Key            string               `json:"key,omitempty"`
	Destination    ExtensionDestination `json:"destination"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Extension) UnmarshalJSON(data []byte) error {
	type Alias Extension
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorExtensionDestination(obj.Destination)
		if err != nil {
			return err
		}
	}

	return nil
}

// ExtensionAWSLambdaDestination implements the interface ExtensionDestination
type ExtensionAWSLambdaDestination struct {
	Arn          string `json:"arn"`
	AccessSecret string `json:"accessSecret"`
	AccessKey    string `json:"accessKey"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionAWSLambdaDestination) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAWSLambdaDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "AWSLambda", Alias: (*Alias)(&obj)})
}

// ExtensionAuthorizationHeaderAuthentication implements the interface ExtensionHTTPDestinationAuthentication
type ExtensionAuthorizationHeaderAuthentication struct {
	HeaderValue string `json:"headerValue"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionAuthorizationHeaderAuthentication) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAuthorizationHeaderAuthentication
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "AuthorizationHeader", Alias: (*Alias)(&obj)})
}

// ExtensionAzureFunctionsAuthentication implements the interface ExtensionHTTPDestinationAuthentication
type ExtensionAzureFunctionsAuthentication struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionAzureFunctionsAuthentication) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAzureFunctionsAuthentication
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "AzureFunctions", Alias: (*Alias)(&obj)})
}

// ExtensionChangeDestinationAction implements the interface ExtensionUpdateAction
type ExtensionChangeDestinationAction struct {
	Destination ExtensionDestination `json:"destination"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionChangeDestinationAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeDestinationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDestination", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionChangeDestinationAction) UnmarshalJSON(data []byte) error {
	type Alias ExtensionChangeDestinationAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorExtensionDestination(obj.Destination)
		if err != nil {
			return err
		}
	}

	return nil
}

// ExtensionChangeTriggersAction implements the interface ExtensionUpdateAction
type ExtensionChangeTriggersAction struct {
	Triggers []ExtensionTrigger `json:"triggers"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionChangeTriggersAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeTriggersAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTriggers", Alias: (*Alias)(&obj)})
}

// ExtensionDraft is a standalone struct
type ExtensionDraft struct {
	Triggers    []ExtensionTrigger   `json:"triggers"`
	TimeoutInMs int                  `json:"timeoutInMs,omitempty"`
	Key         string               `json:"key,omitempty"`
	Destination ExtensionDestination `json:"destination"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionDraft) UnmarshalJSON(data []byte) error {
	type Alias ExtensionDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorExtensionDestination(obj.Destination)
		if err != nil {
			return err
		}
	}

	return nil
}

// ExtensionHTTPDestination implements the interface ExtensionDestination
type ExtensionHTTPDestination struct {
	URL            string                                 `json:"url"`
	Authentication ExtensionHTTPDestinationAuthentication `json:"authentication,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionHTTPDestination) MarshalJSON() ([]byte, error) {
	type Alias ExtensionHTTPDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "HTTP", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionHTTPDestination) UnmarshalJSON(data []byte) error {
	type Alias ExtensionHTTPDestination
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Authentication != nil {
		var err error
		obj.Authentication, err = mapDiscriminatorExtensionHTTPDestinationAuthentication(obj.Authentication)
		if err != nil {
			return err
		}
	}

	return nil
}

// ExtensionInput is a standalone struct
type ExtensionInput struct {
	Resource Reference       `json:"resource"`
	Action   ExtensionAction `json:"action"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionInput) UnmarshalJSON(data []byte) error {
	type Alias ExtensionInput
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

// ExtensionPagedQueryResponse is a standalone struct
type ExtensionPagedQueryResponse struct {
	Total   int         `json:"total,omitempty"`
	Results []Extension `json:"results"`
	Offset  int         `json:"offset"`
	Count   int         `json:"count"`
}

// ExtensionSetKeyAction implements the interface ExtensionUpdateAction
type ExtensionSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ExtensionSetTimeoutInMsAction implements the interface ExtensionUpdateAction
type ExtensionSetTimeoutInMsAction struct {
	TimeoutInMs int `json:"timeoutInMs,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionSetTimeoutInMsAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionSetTimeoutInMsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTimeoutInMs", Alias: (*Alias)(&obj)})
}

// ExtensionTrigger is a standalone struct
type ExtensionTrigger struct {
	ResourceTypeID ExtensionResourceTypeID `json:"resourceTypeId"`
	Actions        []ExtensionAction       `json:"actions"`
}

// ExtensionUpdate is a standalone struct
type ExtensionUpdate struct {
	Version int                     `json:"version"`
	Actions []ExtensionUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionUpdate) UnmarshalJSON(data []byte) error {
	type Alias ExtensionUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorExtensionUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
