// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type RetireGrantInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier of the grant to retire. The grant ID is returned in the
	// response to a CreateGrant operation.
	//
	//    * Grant ID Example - 0123456789012345678901234567890123456789012345678901234567890123
	GrantId *string `min:"1" type:"string"`

	// Token that identifies the grant to be retired.
	GrantToken *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the CMK associated with the grant.
	//
	// For example: arn:aws:kms:us-east-2:444455556666:key/1234abcd-12ab-34cd-56ef-1234567890ab
	KeyId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s RetireGrantInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RetireGrantInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RetireGrantInput"}
	if s.GrantId != nil && len(*s.GrantId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GrantId", 1))
	}
	if s.GrantToken != nil && len(*s.GrantToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GrantToken", 1))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RetireGrantOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RetireGrantOutput) String() string {
	return awsutil.Prettify(s)
}

const opRetireGrant = "RetireGrant"

// RetireGrantRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Retires a grant. To clean up, you can retire a grant when you're done using
// it. You should revoke a grant when you intend to actively deny operations
// that depend on it. The following are permitted to call this API:
//
//    * The AWS account (root user) under which the grant was created
//
//    * The RetiringPrincipal, if present in the grant
//
//    * The GranteePrincipal, if RetireGrant is an operation specified in the
//    grant
//
// You must identify the grant to retire by its grant token or by a combination
// of the grant ID and the Amazon Resource Name (ARN) of the customer master
// key (CMK). A grant token is a unique variable-length base64-encoded string.
// A grant ID is a 64 character unique identifier of a grant. The CreateGrant
// operation returns both.
//
//    // Example sending a request using RetireGrantRequest.
//    req := client.RetireGrantRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/RetireGrant
func (c *Client) RetireGrantRequest(input *RetireGrantInput) RetireGrantRequest {
	op := &aws.Operation{
		Name:       opRetireGrant,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RetireGrantInput{}
	}

	req := c.newRequest(op, input, &RetireGrantOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RetireGrantRequest{Request: req, Input: input, Copy: c.RetireGrantRequest}
}

// RetireGrantRequest is the request type for the
// RetireGrant API operation.
type RetireGrantRequest struct {
	*aws.Request
	Input *RetireGrantInput
	Copy  func(*RetireGrantInput) RetireGrantRequest
}

// Send marshals and sends the RetireGrant API request.
func (r RetireGrantRequest) Send(ctx context.Context) (*RetireGrantResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RetireGrantResponse{
		RetireGrantOutput: r.Request.Data.(*RetireGrantOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RetireGrantResponse is the response type for the
// RetireGrant API operation.
type RetireGrantResponse struct {
	*RetireGrantOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RetireGrant request.
func (r *RetireGrantResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
