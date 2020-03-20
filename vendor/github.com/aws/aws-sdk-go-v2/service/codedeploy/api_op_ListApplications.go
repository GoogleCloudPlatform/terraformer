// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a ListApplications operation.
type ListApplicationsInput struct {
	_ struct{} `type:"structure"`

	// An identifier returned from the previous list applications call. It can be
	// used to return the next set of applications in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListApplicationsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a ListApplications operation.
type ListApplicationsOutput struct {
	_ struct{} `type:"structure"`

	// A list of application names.
	Applications []string `locationName:"applications" type:"list"`

	// If a large amount of information is returned, an identifier is also returned.
	// It can be used in a subsequent list applications call to return the next
	// set of applications in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListApplicationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListApplications = "ListApplications"

// ListApplicationsRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Lists the applications registered with the IAM user or AWS account.
//
//    // Example sending a request using ListApplicationsRequest.
//    req := client.ListApplicationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListApplications
func (c *Client) ListApplicationsRequest(input *ListApplicationsInput) ListApplicationsRequest {
	op := &aws.Operation{
		Name:       opListApplications,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListApplicationsInput{}
	}

	req := c.newRequest(op, input, &ListApplicationsOutput{})
	return ListApplicationsRequest{Request: req, Input: input, Copy: c.ListApplicationsRequest}
}

// ListApplicationsRequest is the request type for the
// ListApplications API operation.
type ListApplicationsRequest struct {
	*aws.Request
	Input *ListApplicationsInput
	Copy  func(*ListApplicationsInput) ListApplicationsRequest
}

// Send marshals and sends the ListApplications API request.
func (r ListApplicationsRequest) Send(ctx context.Context) (*ListApplicationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListApplicationsResponse{
		ListApplicationsOutput: r.Request.Data.(*ListApplicationsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListApplicationsRequestPaginator returns a paginator for ListApplications.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListApplicationsRequest(input)
//   p := codedeploy.NewListApplicationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListApplicationsPaginator(req ListApplicationsRequest) ListApplicationsPaginator {
	return ListApplicationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListApplicationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListApplicationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListApplicationsPaginator struct {
	aws.Pager
}

func (p *ListApplicationsPaginator) CurrentPage() *ListApplicationsOutput {
	return p.Pager.CurrentPage().(*ListApplicationsOutput)
}

// ListApplicationsResponse is the response type for the
// ListApplications API operation.
type ListApplicationsResponse struct {
	*ListApplicationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListApplications request.
func (r *ListApplicationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
