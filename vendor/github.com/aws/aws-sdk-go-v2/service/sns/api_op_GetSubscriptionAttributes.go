// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for GetSubscriptionAttributes.
type GetSubscriptionAttributesInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the subscription whose properties you want to get.
	//
	// SubscriptionArn is a required field
	SubscriptionArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetSubscriptionAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSubscriptionAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSubscriptionAttributesInput"}

	if s.SubscriptionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response for GetSubscriptionAttributes action.
type GetSubscriptionAttributesOutput struct {
	_ struct{} `type:"structure"`

	// A map of the subscription's attributes. Attributes in this map include the
	// following:
	//
	//    * ConfirmationWasAuthenticated – true if the subscription confirmation
	//    request was authenticated.
	//
	//    * DeliveryPolicy – The JSON serialization of the subscription's delivery
	//    policy.
	//
	//    * EffectiveDeliveryPolicy – The JSON serialization of the effective
	//    delivery policy that takes into account the topic delivery policy and
	//    account system defaults.
	//
	//    * FilterPolicy – The filter policy JSON that is assigned to the subscription.
	//
	//    * Owner – The AWS account ID of the subscription's owner.
	//
	//    * PendingConfirmation – true if the subscription hasn't been confirmed.
	//    To confirm a pending subscription, call the ConfirmSubscription action
	//    with a confirmation token.
	//
	//    * RawMessageDelivery – true if raw message delivery is enabled for the
	//    subscription. Raw messages are free of JSON formatting and can be sent
	//    to HTTP/S and Amazon SQS endpoints.
	//
	//    * RedrivePolicy – When specified, sends undeliverable messages to the
	//    specified Amazon SQS dead-letter queue. Messages that can't be delivered
	//    due to client errors (for example, when the subscribed endpoint is unreachable)
	//    or server errors (for example, when the service that powers the subscribed
	//    endpoint becomes unavailable) are held in the dead-letter queue for further
	//    analysis or reprocessing.
	//
	//    * SubscriptionArn – The subscription's ARN.
	//
	//    * TopicArn – The topic ARN that the subscription is associated with.
	Attributes map[string]string `type:"map"`
}

// String returns the string representation
func (s GetSubscriptionAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSubscriptionAttributes = "GetSubscriptionAttributes"

// GetSubscriptionAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Returns all of the properties of a subscription.
//
//    // Example sending a request using GetSubscriptionAttributesRequest.
//    req := client.GetSubscriptionAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/GetSubscriptionAttributes
func (c *Client) GetSubscriptionAttributesRequest(input *GetSubscriptionAttributesInput) GetSubscriptionAttributesRequest {
	op := &aws.Operation{
		Name:       opGetSubscriptionAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSubscriptionAttributesInput{}
	}

	req := c.newRequest(op, input, &GetSubscriptionAttributesOutput{})
	return GetSubscriptionAttributesRequest{Request: req, Input: input, Copy: c.GetSubscriptionAttributesRequest}
}

// GetSubscriptionAttributesRequest is the request type for the
// GetSubscriptionAttributes API operation.
type GetSubscriptionAttributesRequest struct {
	*aws.Request
	Input *GetSubscriptionAttributesInput
	Copy  func(*GetSubscriptionAttributesInput) GetSubscriptionAttributesRequest
}

// Send marshals and sends the GetSubscriptionAttributes API request.
func (r GetSubscriptionAttributesRequest) Send(ctx context.Context) (*GetSubscriptionAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSubscriptionAttributesResponse{
		GetSubscriptionAttributesOutput: r.Request.Data.(*GetSubscriptionAttributesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSubscriptionAttributesResponse is the response type for the
// GetSubscriptionAttributes API operation.
type GetSubscriptionAttributesResponse struct {
	*GetSubscriptionAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSubscriptionAttributes request.
func (r *GetSubscriptionAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
