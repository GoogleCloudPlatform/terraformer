// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the GetTopicRule operation.
type GetTopicRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of the rule.
	//
	// RuleName is a required field
	RuleName *string `location:"uri" locationName:"ruleName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTopicRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTopicRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTopicRuleInput"}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}
	if s.RuleName != nil && len(*s.RuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTopicRuleInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RuleName != nil {
		v := *s.RuleName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ruleName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output from the GetTopicRule operation.
type GetTopicRuleOutput struct {
	_ struct{} `type:"structure"`

	// The rule.
	Rule *TopicRule `locationName:"rule" type:"structure"`

	// The rule ARN.
	RuleArn *string `locationName:"ruleArn" type:"string"`
}

// String returns the string representation
func (s GetTopicRuleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTopicRuleOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Rule != nil {
		v := s.Rule

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "rule", v, metadata)
	}
	if s.RuleArn != nil {
		v := *s.RuleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ruleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetTopicRule = "GetTopicRule"

// GetTopicRuleRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about the rule.
//
//    // Example sending a request using GetTopicRuleRequest.
//    req := client.GetTopicRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetTopicRuleRequest(input *GetTopicRuleInput) GetTopicRuleRequest {
	op := &aws.Operation{
		Name:       opGetTopicRule,
		HTTPMethod: "GET",
		HTTPPath:   "/rules/{ruleName}",
	}

	if input == nil {
		input = &GetTopicRuleInput{}
	}

	req := c.newRequest(op, input, &GetTopicRuleOutput{})
	return GetTopicRuleRequest{Request: req, Input: input, Copy: c.GetTopicRuleRequest}
}

// GetTopicRuleRequest is the request type for the
// GetTopicRule API operation.
type GetTopicRuleRequest struct {
	*aws.Request
	Input *GetTopicRuleInput
	Copy  func(*GetTopicRuleInput) GetTopicRuleRequest
}

// Send marshals and sends the GetTopicRule API request.
func (r GetTopicRuleRequest) Send(ctx context.Context) (*GetTopicRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTopicRuleResponse{
		GetTopicRuleOutput: r.Request.Data.(*GetTopicRuleOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTopicRuleResponse is the response type for the
// GetTopicRule API operation.
type GetTopicRuleResponse struct {
	*GetTopicRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTopicRule request.
func (r *GetTopicRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
