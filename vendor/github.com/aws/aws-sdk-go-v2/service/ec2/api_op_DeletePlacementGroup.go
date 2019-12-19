// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type DeletePlacementGroupInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The name of the placement group.
	//
	// GroupName is a required field
	GroupName *string `locationName:"groupName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePlacementGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePlacementGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePlacementGroupInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeletePlacementGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePlacementGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeletePlacementGroup = "DeletePlacementGroup"

// DeletePlacementGroupRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified placement group. You must terminate all instances in
// the placement group before you can delete the placement group. For more information,
// see Placement Groups (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DeletePlacementGroupRequest.
//    req := client.DeletePlacementGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeletePlacementGroup
func (c *Client) DeletePlacementGroupRequest(input *DeletePlacementGroupInput) DeletePlacementGroupRequest {
	op := &aws.Operation{
		Name:       opDeletePlacementGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePlacementGroupInput{}
	}

	req := c.newRequest(op, input, &DeletePlacementGroupOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeletePlacementGroupRequest{Request: req, Input: input, Copy: c.DeletePlacementGroupRequest}
}

// DeletePlacementGroupRequest is the request type for the
// DeletePlacementGroup API operation.
type DeletePlacementGroupRequest struct {
	*aws.Request
	Input *DeletePlacementGroupInput
	Copy  func(*DeletePlacementGroupInput) DeletePlacementGroupRequest
}

// Send marshals and sends the DeletePlacementGroup API request.
func (r DeletePlacementGroupRequest) Send(ctx context.Context) (*DeletePlacementGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePlacementGroupResponse{
		DeletePlacementGroupOutput: r.Request.Data.(*DeletePlacementGroupOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePlacementGroupResponse is the response type for the
// DeletePlacementGroup API operation.
type DeletePlacementGroupResponse struct {
	*DeletePlacementGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePlacementGroup request.
func (r *DeletePlacementGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
