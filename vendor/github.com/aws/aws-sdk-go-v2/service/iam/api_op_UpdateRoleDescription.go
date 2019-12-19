// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateRoleDescriptionInput struct {
	_ struct{} `type:"structure"`

	// The new description that you want to apply to the specified role.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// The name of the role that you want to modify.
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateRoleDescriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRoleDescriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRoleDescriptionInput"}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateRoleDescriptionOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the modified role.
	Role *Role `type:"structure"`
}

// String returns the string representation
func (s UpdateRoleDescriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateRoleDescription = "UpdateRoleDescription"

// UpdateRoleDescriptionRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Use UpdateRole instead.
//
// Modifies only the description of a role. This operation performs the same
// function as the Description parameter in the UpdateRole operation.
//
//    // Example sending a request using UpdateRoleDescriptionRequest.
//    req := client.UpdateRoleDescriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateRoleDescription
func (c *Client) UpdateRoleDescriptionRequest(input *UpdateRoleDescriptionInput) UpdateRoleDescriptionRequest {
	op := &aws.Operation{
		Name:       opUpdateRoleDescription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRoleDescriptionInput{}
	}

	req := c.newRequest(op, input, &UpdateRoleDescriptionOutput{})
	return UpdateRoleDescriptionRequest{Request: req, Input: input, Copy: c.UpdateRoleDescriptionRequest}
}

// UpdateRoleDescriptionRequest is the request type for the
// UpdateRoleDescription API operation.
type UpdateRoleDescriptionRequest struct {
	*aws.Request
	Input *UpdateRoleDescriptionInput
	Copy  func(*UpdateRoleDescriptionInput) UpdateRoleDescriptionRequest
}

// Send marshals and sends the UpdateRoleDescription API request.
func (r UpdateRoleDescriptionRequest) Send(ctx context.Context) (*UpdateRoleDescriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRoleDescriptionResponse{
		UpdateRoleDescriptionOutput: r.Request.Data.(*UpdateRoleDescriptionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRoleDescriptionResponse is the response type for the
// UpdateRoleDescription API operation.
type UpdateRoleDescriptionResponse struct {
	*UpdateRoleDescriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRoleDescription request.
func (r *UpdateRoleDescriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
