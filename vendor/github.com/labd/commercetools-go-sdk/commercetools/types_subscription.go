// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// SubscriptionHealthStatus is an enum type
type SubscriptionHealthStatus string

// Enum values for SubscriptionHealthStatus
const (
	SubscriptionHealthStatusHealthy                           SubscriptionHealthStatus = "Healthy"
	SubscriptionHealthStatusConfigurationError                SubscriptionHealthStatus = "ConfigurationError"
	SubscriptionHealthStatusConfigurationErrorDeliveryStopped SubscriptionHealthStatus = "ConfigurationErrorDeliveryStopped"
	SubscriptionHealthStatusTemporaryError                    SubscriptionHealthStatus = "TemporaryError"
)

// DeliveryFormat uses type as discriminator attribute
type DeliveryFormat interface{}

func mapDiscriminatorDeliveryFormat(input interface{}) (DeliveryFormat, error) {
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
	case "CloudEvents":
		new := DeliveryCloudEventsFormat{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Platform":
		new := DeliveryPlatformFormat{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// Destination uses type as discriminator attribute
type Destination interface{}

func mapDiscriminatorDestination(input interface{}) (Destination, error) {
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
	case "EventGrid":
		new := AzureEventGridDestination{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "AzureServiceBus":
		new := AzureServiceBusDestination{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "GoogleCloudPubSub":
		new := GoogleCloudPubSubDestination{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "IronMQ":
		new := IronMqDestination{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "SNS":
		new := SnsDestination{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "SQS":
		new := SqsDestination{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// SubscriptionDelivery uses notificationType as discriminator attribute
type SubscriptionDelivery interface{}

func mapDiscriminatorSubscriptionDelivery(input interface{}) (SubscriptionDelivery, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["notificationType"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'notificationType'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "Message":
		new := MessageDelivery{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ResourceCreated":
		new := ResourceCreatedDelivery{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ResourceDeleted":
		new := ResourceDeletedDelivery{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ResourceUpdated":
		new := ResourceUpdatedDelivery{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// SubscriptionUpdateAction uses action as discriminator attribute
type SubscriptionUpdateAction interface{}

func mapDiscriminatorSubscriptionUpdateAction(input interface{}) (SubscriptionUpdateAction, error) {
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
		new := SubscriptionChangeDestinationAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		if new.Destination != nil {
			new.Destination, err = mapDiscriminatorDestination(new.Destination)
			if err != nil {
				return nil, err
			}
		}
		return new, nil
	case "setChanges":
		new := SubscriptionSetChangesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := SubscriptionSetKeyAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMessages":
		new := SubscriptionSetMessagesAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// AzureEventGridDestination implements the interface Destination
type AzureEventGridDestination struct {
	URI       string `json:"uri"`
	AccessKey string `json:"accessKey"`
}

// MarshalJSON override to set the discriminator value
func (obj AzureEventGridDestination) MarshalJSON() ([]byte, error) {
	type Alias AzureEventGridDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "EventGrid", Alias: (*Alias)(&obj)})
}

// AzureServiceBusDestination implements the interface Destination
type AzureServiceBusDestination struct {
	ConnectionString string `json:"connectionString"`
}

// MarshalJSON override to set the discriminator value
func (obj AzureServiceBusDestination) MarshalJSON() ([]byte, error) {
	type Alias AzureServiceBusDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "AzureServiceBus", Alias: (*Alias)(&obj)})
}

// ChangeSubscription is a standalone struct
type ChangeSubscription struct {
	ResourceTypeID string `json:"resourceTypeId"`
}

// DeliveryCloudEventsFormat implements the interface DeliveryFormat
type DeliveryCloudEventsFormat struct {
	CloudEventsVersion string `json:"cloudEventsVersion"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryCloudEventsFormat) MarshalJSON() ([]byte, error) {
	type Alias DeliveryCloudEventsFormat
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CloudEvents", Alias: (*Alias)(&obj)})
}

// DeliveryPlatformFormat implements the interface DeliveryFormat
type DeliveryPlatformFormat struct{}

// MarshalJSON override to set the discriminator value
func (obj DeliveryPlatformFormat) MarshalJSON() ([]byte, error) {
	type Alias DeliveryPlatformFormat
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "Platform", Alias: (*Alias)(&obj)})
}

// GoogleCloudPubSubDestination implements the interface Destination
type GoogleCloudPubSubDestination struct {
	Topic     string `json:"topic"`
	ProjectID string `json:"projectId"`
}

// MarshalJSON override to set the discriminator value
func (obj GoogleCloudPubSubDestination) MarshalJSON() ([]byte, error) {
	type Alias GoogleCloudPubSubDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "GoogleCloudPubSub", Alias: (*Alias)(&obj)})
}

// IronMqDestination implements the interface Destination
type IronMqDestination struct {
	URI string `json:"uri"`
}

// MarshalJSON override to set the discriminator value
func (obj IronMqDestination) MarshalJSON() ([]byte, error) {
	type Alias IronMqDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "IronMQ", Alias: (*Alias)(&obj)})
}

// MessageDelivery implements the interface SubscriptionDelivery
type MessageDelivery struct {
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ProjectKey                      string                   `json:"projectKey"`
	Version                         int                      `json:"version"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	PayloadNotIncluded              *PayloadNotIncluded      `json:"payloadNotIncluded"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	ID                              string                   `json:"id"`
	CreatedAt                       time.Time                `json:"createdAt"`
}

// MarshalJSON override to set the discriminator value
func (obj MessageDelivery) MarshalJSON() ([]byte, error) {
	type Alias MessageDelivery
	return json.Marshal(struct {
		NotificationType string `json:"notificationType"`
		*Alias
	}{NotificationType: "Message", Alias: (*Alias)(&obj)})
}

// MessageSubscription is a standalone struct
type MessageSubscription struct {
	Types          []string `json:"types,omitempty"`
	ResourceTypeID string   `json:"resourceTypeId"`
}

// PayloadNotIncluded is a standalone struct
type PayloadNotIncluded struct {
	Reason      string `json:"reason"`
	PayloadType string `json:"payloadType"`
}

// ResourceCreatedDelivery implements the interface SubscriptionDelivery
type ResourceCreatedDelivery struct {
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ProjectKey                      string                   `json:"projectKey"`
	Version                         int                      `json:"version"`
	ModifiedAt                      time.Time                `json:"modifiedAt"`
}

// MarshalJSON override to set the discriminator value
func (obj ResourceCreatedDelivery) MarshalJSON() ([]byte, error) {
	type Alias ResourceCreatedDelivery
	return json.Marshal(struct {
		NotificationType string `json:"notificationType"`
		*Alias
	}{NotificationType: "ResourceCreated", Alias: (*Alias)(&obj)})
}

// ResourceDeletedDelivery implements the interface SubscriptionDelivery
type ResourceDeletedDelivery struct {
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ProjectKey                      string                   `json:"projectKey"`
	Version                         int                      `json:"version"`
	ModifiedAt                      time.Time                `json:"modifiedAt"`
}

// MarshalJSON override to set the discriminator value
func (obj ResourceDeletedDelivery) MarshalJSON() ([]byte, error) {
	type Alias ResourceDeletedDelivery
	return json.Marshal(struct {
		NotificationType string `json:"notificationType"`
		*Alias
	}{NotificationType: "ResourceDeleted", Alias: (*Alias)(&obj)})
}

// ResourceUpdatedDelivery implements the interface SubscriptionDelivery
type ResourceUpdatedDelivery struct {
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	Resource                        Reference                `json:"resource"`
	ProjectKey                      string                   `json:"projectKey"`
	Version                         int                      `json:"version"`
	OldVersion                      int                      `json:"oldVersion"`
	ModifiedAt                      time.Time                `json:"modifiedAt"`
}

// MarshalJSON override to set the discriminator value
func (obj ResourceUpdatedDelivery) MarshalJSON() ([]byte, error) {
	type Alias ResourceUpdatedDelivery
	return json.Marshal(struct {
		NotificationType string `json:"notificationType"`
		*Alias
	}{NotificationType: "ResourceUpdated", Alias: (*Alias)(&obj)})
}

// SnsDestination implements the interface Destination
type SnsDestination struct {
	TopicArn     string `json:"topicArn"`
	AccessSecret string `json:"accessSecret"`
	AccessKey    string `json:"accessKey"`
}

// MarshalJSON override to set the discriminator value
func (obj SnsDestination) MarshalJSON() ([]byte, error) {
	type Alias SnsDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "SNS", Alias: (*Alias)(&obj)})
}

// SqsDestination implements the interface Destination
type SqsDestination struct {
	Region       string `json:"region"`
	QueueURL     string `json:"queueUrl"`
	AccessSecret string `json:"accessSecret"`
	AccessKey    string `json:"accessKey"`
}

// MarshalJSON override to set the discriminator value
func (obj SqsDestination) MarshalJSON() ([]byte, error) {
	type Alias SqsDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "SQS", Alias: (*Alias)(&obj)})
}

// Subscription is of type LoggedResource
type Subscription struct {
	Version        int                      `json:"version"`
	LastModifiedAt time.Time                `json:"lastModifiedAt"`
	ID             string                   `json:"id"`
	CreatedAt      time.Time                `json:"createdAt"`
	LastModifiedBy *LastModifiedBy          `json:"lastModifiedBy,omitempty"`
	CreatedBy      *CreatedBy               `json:"createdBy,omitempty"`
	Status         SubscriptionHealthStatus `json:"status"`
	Messages       []MessageSubscription    `json:"messages"`
	Key            string                   `json:"key,omitempty"`
	Format         DeliveryFormat           `json:"format"`
	Destination    Destination              `json:"destination"`
	Changes        []ChangeSubscription     `json:"changes"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Subscription) UnmarshalJSON(data []byte) error {
	type Alias Subscription
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorDestination(obj.Destination)
		if err != nil {
			return err
		}
	}
	if obj.Format != nil {
		var err error
		obj.Format, err = mapDiscriminatorDeliveryFormat(obj.Format)
		if err != nil {
			return err
		}
	}

	return nil
}

// SubscriptionChangeDestinationAction implements the interface SubscriptionUpdateAction
type SubscriptionChangeDestinationAction struct {
	Destination Destination `json:"destination"`
}

// MarshalJSON override to set the discriminator value
func (obj SubscriptionChangeDestinationAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionChangeDestinationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDestination", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SubscriptionChangeDestinationAction) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionChangeDestinationAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorDestination(obj.Destination)
		if err != nil {
			return err
		}
	}

	return nil
}

// SubscriptionDraft is a standalone struct
type SubscriptionDraft struct {
	Messages    []MessageSubscription `json:"messages,omitempty"`
	Key         string                `json:"key,omitempty"`
	Format      DeliveryFormat        `json:"format,omitempty"`
	Destination Destination           `json:"destination"`
	Changes     []ChangeSubscription  `json:"changes,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SubscriptionDraft) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorDestination(obj.Destination)
		if err != nil {
			return err
		}
	}
	if obj.Format != nil {
		var err error
		obj.Format, err = mapDiscriminatorDeliveryFormat(obj.Format)
		if err != nil {
			return err
		}
	}

	return nil
}

// SubscriptionPagedQueryResponse is a standalone struct
type SubscriptionPagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Results []Subscription `json:"results"`
	Offset  int            `json:"offset"`
	Count   int            `json:"count"`
}

// SubscriptionSetChangesAction implements the interface SubscriptionUpdateAction
type SubscriptionSetChangesAction struct {
	Changes []ChangeSubscription `json:"changes,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj SubscriptionSetChangesAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setChanges", Alias: (*Alias)(&obj)})
}

// SubscriptionSetKeyAction implements the interface SubscriptionUpdateAction
type SubscriptionSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj SubscriptionSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// SubscriptionSetMessagesAction implements the interface SubscriptionUpdateAction
type SubscriptionSetMessagesAction struct {
	Messages []MessageSubscription `json:"messages,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj SubscriptionSetMessagesAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetMessagesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMessages", Alias: (*Alias)(&obj)})
}

// SubscriptionUpdate is a standalone struct
type SubscriptionUpdate struct {
	Version int                        `json:"version"`
	Actions []SubscriptionUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SubscriptionUpdate) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorSubscriptionUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
