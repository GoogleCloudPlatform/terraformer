// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the GetStackPolicy action.
type GetStackPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name or unique stack ID that is associated with the stack whose policy
	// you want to get.
	//
	// StackName is a required field
	StackName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetStackPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetStackPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetStackPolicyInput"}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for the GetStackPolicy action.
type GetStackPolicyOutput struct {
	_ struct{} `type:"structure"`

	// Structure containing the stack policy body. (For more information, go to
	// Prevent Updates to Stack Resources (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html)
	// in the AWS CloudFormation User Guide.)
	StackPolicyBody *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetStackPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetStackPolicy = "GetStackPolicy"

// GetStackPolicyRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns the stack policy for a specified stack. If a stack doesn't have a
// policy, a null value is returned.
//
//    // Example sending a request using GetStackPolicyRequest.
//    req := client.GetStackPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/GetStackPolicy
func (c *Client) GetStackPolicyRequest(input *GetStackPolicyInput) GetStackPolicyRequest {
	op := &aws.Operation{
		Name:       opGetStackPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetStackPolicyInput{}
	}

	req := c.newRequest(op, input, &GetStackPolicyOutput{})
	return GetStackPolicyRequest{Request: req, Input: input, Copy: c.GetStackPolicyRequest}
}

// GetStackPolicyRequest is the request type for the
// GetStackPolicy API operation.
type GetStackPolicyRequest struct {
	*aws.Request
	Input *GetStackPolicyInput
	Copy  func(*GetStackPolicyInput) GetStackPolicyRequest
}

// Send marshals and sends the GetStackPolicy API request.
func (r GetStackPolicyRequest) Send(ctx context.Context) (*GetStackPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetStackPolicyResponse{
		GetStackPolicyOutput: r.Request.Data.(*GetStackPolicyOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetStackPolicyResponse is the response type for the
// GetStackPolicy API operation.
type GetStackPolicyResponse struct {
	*GetStackPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetStackPolicy request.
func (r *GetStackPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
