// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribePolicyInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the policy that you want details about. You
	// can get the ID from the ListPolicies or ListPoliciesForTarget operations.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a policy ID string
	// requires "p-" followed by from 8 to 128 lowercase or uppercase letters, digits,
	// or the underscore character (_).
	//
	// PolicyId is a required field
	PolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePolicyInput"}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribePolicyOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the specified policy.
	Policy *Policy `type:"structure"`
}

// String returns the string representation
func (s DescribePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePolicy = "DescribePolicy"

// DescribePolicyRequest returns a request value for making API operation for
// AWS Organizations.
//
// Retrieves information about a policy.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using DescribePolicyRequest.
//    req := client.DescribePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DescribePolicy
func (c *Client) DescribePolicyRequest(input *DescribePolicyInput) DescribePolicyRequest {
	op := &aws.Operation{
		Name:       opDescribePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePolicyInput{}
	}

	req := c.newRequest(op, input, &DescribePolicyOutput{})
	return DescribePolicyRequest{Request: req, Input: input, Copy: c.DescribePolicyRequest}
}

// DescribePolicyRequest is the request type for the
// DescribePolicy API operation.
type DescribePolicyRequest struct {
	*aws.Request
	Input *DescribePolicyInput
	Copy  func(*DescribePolicyInput) DescribePolicyRequest
}

// Send marshals and sends the DescribePolicy API request.
func (r DescribePolicyRequest) Send(ctx context.Context) (*DescribePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePolicyResponse{
		DescribePolicyOutput: r.Request.Data.(*DescribePolicyOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePolicyResponse is the response type for the
// DescribePolicy API operation.
type DescribePolicyResponse struct {
	*DescribePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePolicy request.
func (r *DescribePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
