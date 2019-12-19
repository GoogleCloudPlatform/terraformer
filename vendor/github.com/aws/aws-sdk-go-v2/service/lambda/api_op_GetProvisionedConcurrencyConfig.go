// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetProvisionedConcurrencyConfigInput struct {
	_ struct{} `type:"structure"`

	// The name of the Lambda function.
	//
	// Name formats
	//
	//    * Function name - my-function.
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//    * Partial ARN - 123456789012:function:my-function.
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// The version number or alias name.
	//
	// Qualifier is a required field
	Qualifier *string `location:"querystring" locationName:"Qualifier" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetProvisionedConcurrencyConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetProvisionedConcurrencyConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetProvisionedConcurrencyConfigInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}

	if s.Qualifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("Qualifier"))
	}
	if s.Qualifier != nil && len(*s.Qualifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Qualifier", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetProvisionedConcurrencyConfigInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Qualifier != nil {
		v := *s.Qualifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Qualifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetProvisionedConcurrencyConfigOutput struct {
	_ struct{} `type:"structure"`

	// The amount of provisioned concurrency allocated.
	AllocatedProvisionedConcurrentExecutions *int64 `type:"integer"`

	// The amount of provisioned concurrency available.
	AvailableProvisionedConcurrentExecutions *int64 `type:"integer"`

	// The date and time that a user last updated the configuration, in ISO 8601
	// format (https://www.iso.org/iso-8601-date-and-time-format.html).
	LastModified *string `type:"string"`

	// The amount of provisioned concurrency requested.
	RequestedProvisionedConcurrentExecutions *int64 `min:"1" type:"integer"`

	// The status of the allocation process.
	Status ProvisionedConcurrencyStatusEnum `type:"string" enum:"true"`

	// For failed allocations, the reason that provisioned concurrency could not
	// be allocated.
	StatusReason *string `type:"string"`
}

// String returns the string representation
func (s GetProvisionedConcurrencyConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetProvisionedConcurrencyConfigOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AllocatedProvisionedConcurrentExecutions != nil {
		v := *s.AllocatedProvisionedConcurrentExecutions

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AllocatedProvisionedConcurrentExecutions", protocol.Int64Value(v), metadata)
	}
	if s.AvailableProvisionedConcurrentExecutions != nil {
		v := *s.AvailableProvisionedConcurrentExecutions

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AvailableProvisionedConcurrentExecutions", protocol.Int64Value(v), metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestedProvisionedConcurrentExecutions != nil {
		v := *s.RequestedProvisionedConcurrentExecutions

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestedProvisionedConcurrentExecutions", protocol.Int64Value(v), metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusReason != nil {
		v := *s.StatusReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StatusReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetProvisionedConcurrencyConfig = "GetProvisionedConcurrencyConfig"

// GetProvisionedConcurrencyConfigRequest returns a request value for making API operation for
// AWS Lambda.
//
// Retrieves the provisioned concurrency configuration for a function's alias
// or version.
//
//    // Example sending a request using GetProvisionedConcurrencyConfigRequest.
//    req := client.GetProvisionedConcurrencyConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetProvisionedConcurrencyConfig
func (c *Client) GetProvisionedConcurrencyConfigRequest(input *GetProvisionedConcurrencyConfigInput) GetProvisionedConcurrencyConfigRequest {
	op := &aws.Operation{
		Name:       opGetProvisionedConcurrencyConfig,
		HTTPMethod: "GET",
		HTTPPath:   "/2019-09-30/functions/{FunctionName}/provisioned-concurrency",
	}

	if input == nil {
		input = &GetProvisionedConcurrencyConfigInput{}
	}

	req := c.newRequest(op, input, &GetProvisionedConcurrencyConfigOutput{})
	return GetProvisionedConcurrencyConfigRequest{Request: req, Input: input, Copy: c.GetProvisionedConcurrencyConfigRequest}
}

// GetProvisionedConcurrencyConfigRequest is the request type for the
// GetProvisionedConcurrencyConfig API operation.
type GetProvisionedConcurrencyConfigRequest struct {
	*aws.Request
	Input *GetProvisionedConcurrencyConfigInput
	Copy  func(*GetProvisionedConcurrencyConfigInput) GetProvisionedConcurrencyConfigRequest
}

// Send marshals and sends the GetProvisionedConcurrencyConfig API request.
func (r GetProvisionedConcurrencyConfigRequest) Send(ctx context.Context) (*GetProvisionedConcurrencyConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetProvisionedConcurrencyConfigResponse{
		GetProvisionedConcurrencyConfigOutput: r.Request.Data.(*GetProvisionedConcurrencyConfigOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetProvisionedConcurrencyConfigResponse is the response type for the
// GetProvisionedConcurrencyConfig API operation.
type GetProvisionedConcurrencyConfigResponse struct {
	*GetProvisionedConcurrencyConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetProvisionedConcurrencyConfig request.
func (r *GetProvisionedConcurrencyConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
