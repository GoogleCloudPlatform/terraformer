// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetPullRequestApprovalStatesInput struct {
	_ struct{} `type:"structure"`

	// The system-generated ID for the pull request.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`

	// The system-generated ID for the pull request revision.
	//
	// RevisionId is a required field
	RevisionId *string `locationName:"revisionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPullRequestApprovalStatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPullRequestApprovalStatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPullRequestApprovalStatesInput"}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}

	if s.RevisionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RevisionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetPullRequestApprovalStatesOutput struct {
	_ struct{} `type:"structure"`

	// Information about users who have approved the pull request.
	Approvals []Approval `locationName:"approvals" type:"list"`
}

// String returns the string representation
func (s GetPullRequestApprovalStatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPullRequestApprovalStates = "GetPullRequestApprovalStates"

// GetPullRequestApprovalStatesRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Gets information about the approval states for a specified pull request.
// Approval states only apply to pull requests that have one or more approval
// rules applied to them.
//
//    // Example sending a request using GetPullRequestApprovalStatesRequest.
//    req := client.GetPullRequestApprovalStatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetPullRequestApprovalStates
func (c *Client) GetPullRequestApprovalStatesRequest(input *GetPullRequestApprovalStatesInput) GetPullRequestApprovalStatesRequest {
	op := &aws.Operation{
		Name:       opGetPullRequestApprovalStates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPullRequestApprovalStatesInput{}
	}

	req := c.newRequest(op, input, &GetPullRequestApprovalStatesOutput{})
	return GetPullRequestApprovalStatesRequest{Request: req, Input: input, Copy: c.GetPullRequestApprovalStatesRequest}
}

// GetPullRequestApprovalStatesRequest is the request type for the
// GetPullRequestApprovalStates API operation.
type GetPullRequestApprovalStatesRequest struct {
	*aws.Request
	Input *GetPullRequestApprovalStatesInput
	Copy  func(*GetPullRequestApprovalStatesInput) GetPullRequestApprovalStatesRequest
}

// Send marshals and sends the GetPullRequestApprovalStates API request.
func (r GetPullRequestApprovalStatesRequest) Send(ctx context.Context) (*GetPullRequestApprovalStatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPullRequestApprovalStatesResponse{
		GetPullRequestApprovalStatesOutput: r.Request.Data.(*GetPullRequestApprovalStatesOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPullRequestApprovalStatesResponse is the response type for the
// GetPullRequestApprovalStates API operation.
type GetPullRequestApprovalStatesResponse struct {
	*GetPullRequestApprovalStatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPullRequestApprovalStates request.
func (r *GetPullRequestApprovalStatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
