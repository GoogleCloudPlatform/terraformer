// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTargetHealthInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the target group.
	//
	// TargetGroupArn is a required field
	TargetGroupArn *string `type:"string" required:"true"`

	// The targets.
	Targets []TargetDescription `type:"list"`
}

// String returns the string representation
func (s DescribeTargetHealthInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTargetHealthInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTargetHealthInput"}

	if s.TargetGroupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetGroupArn"))
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTargetHealthOutput struct {
	_ struct{} `type:"structure"`

	// Information about the health of the targets.
	TargetHealthDescriptions []TargetHealthDescription `type:"list"`
}

// String returns the string representation
func (s DescribeTargetHealthOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTargetHealth = "DescribeTargetHealth"

// DescribeTargetHealthRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Describes the health of the specified targets or all of your targets.
//
//    // Example sending a request using DescribeTargetHealthRequest.
//    req := client.DescribeTargetHealthRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/DescribeTargetHealth
func (c *Client) DescribeTargetHealthRequest(input *DescribeTargetHealthInput) DescribeTargetHealthRequest {
	op := &aws.Operation{
		Name:       opDescribeTargetHealth,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTargetHealthInput{}
	}

	req := c.newRequest(op, input, &DescribeTargetHealthOutput{})
	return DescribeTargetHealthRequest{Request: req, Input: input, Copy: c.DescribeTargetHealthRequest}
}

// DescribeTargetHealthRequest is the request type for the
// DescribeTargetHealth API operation.
type DescribeTargetHealthRequest struct {
	*aws.Request
	Input *DescribeTargetHealthInput
	Copy  func(*DescribeTargetHealthInput) DescribeTargetHealthRequest
}

// Send marshals and sends the DescribeTargetHealth API request.
func (r DescribeTargetHealthRequest) Send(ctx context.Context) (*DescribeTargetHealthResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTargetHealthResponse{
		DescribeTargetHealthOutput: r.Request.Data.(*DescribeTargetHealthOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTargetHealthResponse is the response type for the
// DescribeTargetHealth API operation.
type DescribeTargetHealthResponse struct {
	*DescribeTargetHealthOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTargetHealth request.
func (r *DescribeTargetHealthResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
