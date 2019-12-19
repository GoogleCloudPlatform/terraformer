// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteSecurityConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the security configuration.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSecurityConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSecurityConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSecurityConfigurationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteSecurityConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSecurityConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSecurityConfiguration = "DeleteSecurityConfiguration"

// DeleteSecurityConfigurationRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Deletes a security configuration.
//
//    // Example sending a request using DeleteSecurityConfigurationRequest.
//    req := client.DeleteSecurityConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/DeleteSecurityConfiguration
func (c *Client) DeleteSecurityConfigurationRequest(input *DeleteSecurityConfigurationInput) DeleteSecurityConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteSecurityConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSecurityConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeleteSecurityConfigurationOutput{})
	return DeleteSecurityConfigurationRequest{Request: req, Input: input, Copy: c.DeleteSecurityConfigurationRequest}
}

// DeleteSecurityConfigurationRequest is the request type for the
// DeleteSecurityConfiguration API operation.
type DeleteSecurityConfigurationRequest struct {
	*aws.Request
	Input *DeleteSecurityConfigurationInput
	Copy  func(*DeleteSecurityConfigurationInput) DeleteSecurityConfigurationRequest
}

// Send marshals and sends the DeleteSecurityConfiguration API request.
func (r DeleteSecurityConfigurationRequest) Send(ctx context.Context) (*DeleteSecurityConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSecurityConfigurationResponse{
		DeleteSecurityConfigurationOutput: r.Request.Data.(*DeleteSecurityConfigurationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSecurityConfigurationResponse is the response type for the
// DeleteSecurityConfiguration API operation.
type DeleteSecurityConfigurationResponse struct {
	*DeleteSecurityConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSecurityConfiguration request.
func (r *DeleteSecurityConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
