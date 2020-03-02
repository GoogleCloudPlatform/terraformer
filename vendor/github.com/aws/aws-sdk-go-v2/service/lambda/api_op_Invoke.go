// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type InvokeInput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// Up to 3583 bytes of base64-encoded data about the invoking client to pass
	// to the function in the context object.
	ClientContext *string `location:"header" locationName:"X-Amz-Client-Context" type:"string"`

	// The name of the Lambda function, version, or alias.
	//
	// Name formats
	//
	//    * Function name - my-function (name-only), my-function:v1 (with alias).
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//    * Partial ARN - 123456789012:function:my-function.
	//
	// You can append a version number or alias to any of the formats. The length
	// constraint applies only to the full ARN. If you specify only the function
	// name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// Choose from the following options.
	//
	//    * RequestResponse (default) - Invoke the function synchronously. Keep
	//    the connection open until the function returns a response or times out.
	//    The API response includes the function response and additional data.
	//
	//    * Event - Invoke the function asynchronously. Send events that fail multiple
	//    times to the function's dead-letter queue (if it's configured). The API
	//    response only includes a status code.
	//
	//    * DryRun - Validate parameter values and verify that the user or role
	//    has permission to invoke the function.
	InvocationType InvocationType `location:"header" locationName:"X-Amz-Invocation-Type" type:"string" enum:"true"`

	// Set to Tail to include the execution log in the response.
	LogType LogType `location:"header" locationName:"X-Amz-Log-Type" type:"string" enum:"true"`

	// The JSON that you want to provide to your Lambda function as input.
	Payload []byte `type:"blob" sensitive:"true"`

	// Specify a version or alias to invoke a published version of the function.
	Qualifier *string `location:"querystring" locationName:"Qualifier" min:"1" type:"string"`
}

// String returns the string representation
func (s InvokeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InvokeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InvokeInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
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
func (s InvokeInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ClientContext != nil {
		v := *s.ClientContext

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amz-Client-Context", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.InvocationType) > 0 {
		v := s.InvocationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amz-Invocation-Type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.LogType) > 0 {
		v := s.LogType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amz-Log-Type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Payload != nil {
		v := s.Payload

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Payload", protocol.BytesStream(v), metadata)
	}
	if s.Qualifier != nil {
		v := *s.Qualifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Qualifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type InvokeOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The version of the function that executed. When you invoke a function with
	// an alias, this indicates which version the alias resolved to.
	ExecutedVersion *string `location:"header" locationName:"X-Amz-Executed-Version" min:"1" type:"string"`

	// If present, indicates that an error occurred during function execution. Details
	// about the error are included in the response payload.
	FunctionError *string `location:"header" locationName:"X-Amz-Function-Error" type:"string"`

	// The last 4 KB of the execution log, which is base64 encoded.
	LogResult *string `location:"header" locationName:"X-Amz-Log-Result" type:"string"`

	// The response from the function, or an error object.
	Payload []byte `type:"blob" sensitive:"true"`

	// The HTTP status code is in the 200 range for a successful request. For the
	// RequestResponse invocation type, this status code is 200. For the Event invocation
	// type, this status code is 202. For the DryRun invocation type, the status
	// code is 204.
	StatusCode *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s InvokeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s InvokeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExecutedVersion != nil {
		v := *s.ExecutedVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amz-Executed-Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionError != nil {
		v := *s.FunctionError

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amz-Function-Error", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LogResult != nil {
		v := *s.LogResult

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amz-Log-Result", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Payload != nil {
		v := s.Payload

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Payload", protocol.BytesStream(v), metadata)
	}
	// ignoring invalid encode state, StatusCode. StatusCode
	return nil
}

const opInvoke = "Invoke"

// InvokeRequest returns a request value for making API operation for
// AWS Lambda.
//
// Invokes a Lambda function. You can invoke a function synchronously (and wait
// for the response), or asynchronously. To invoke a function asynchronously,
// set InvocationType to Event.
//
// For synchronous invocation (https://docs.aws.amazon.com/lambda/latest/dg/invocation-sync.html),
// details about the function response, including errors, are included in the
// response body and headers. For either invocation type, you can find more
// information in the execution log (https://docs.aws.amazon.com/lambda/latest/dg/monitoring-functions.html)
// and trace (https://docs.aws.amazon.com/lambda/latest/dg/lambda-x-ray.html).
//
// When an error occurs, your function may be invoked multiple times. Retry
// behavior varies by error type, client, event source, and invocation type.
// For example, if you invoke a function asynchronously and it returns an error,
// Lambda executes the function up to two more times. For more information,
// see Retry Behavior (https://docs.aws.amazon.com/lambda/latest/dg/retries-on-errors.html).
//
// For asynchronous invocation (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html),
// Lambda adds events to a queue before sending them to your function. If your
// function does not have enough capacity to keep up with the queue, events
// may be lost. Occasionally, your function may receive the same event multiple
// times, even if no error occurs. To retain events that were not processed,
// configure your function with a dead-letter queue (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq).
//
// The status code in the API response doesn't reflect function errors. Error
// codes are reserved for errors that prevent your function from executing,
// such as permissions errors, limit errors (https://docs.aws.amazon.com/lambda/latest/dg/limits.html),
// or issues with your function's code and configuration. For example, Lambda
// returns TooManyRequestsException if executing the function would cause you
// to exceed a concurrency limit at either the account level (ConcurrentInvocationLimitExceeded)
// or function level (ReservedFunctionConcurrentInvocationLimitExceeded).
//
// For functions with a long timeout, your client might be disconnected during
// synchronous invocation while it waits for a response. Configure your HTTP
// client, SDK, firewall, proxy, or operating system to allow for long connections
// with timeout or keep-alive settings.
//
// This operation requires permission for the lambda:InvokeFunction action.
//
//    // Example sending a request using InvokeRequest.
//    req := client.InvokeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/Invoke
func (c *Client) InvokeRequest(input *InvokeInput) InvokeRequest {
	op := &aws.Operation{
		Name:       opInvoke,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/invocations",
	}

	if input == nil {
		input = &InvokeInput{}
	}

	req := c.newRequest(op, input, &InvokeOutput{})
	return InvokeRequest{Request: req, Input: input, Copy: c.InvokeRequest}
}

// InvokeRequest is the request type for the
// Invoke API operation.
type InvokeRequest struct {
	*aws.Request
	Input *InvokeInput
	Copy  func(*InvokeInput) InvokeRequest
}

// Send marshals and sends the Invoke API request.
func (r InvokeRequest) Send(ctx context.Context) (*InvokeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &InvokeResponse{
		InvokeOutput: r.Request.Data.(*InvokeOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// InvokeResponse is the response type for the
// Invoke API operation.
type InvokeResponse struct {
	*InvokeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// Invoke request.
func (r *InvokeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
