// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartCrawlerInput struct {
	_ struct{} `type:"structure"`

	// Name of the crawler to start.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartCrawlerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartCrawlerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartCrawlerInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartCrawlerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartCrawlerOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartCrawler = "StartCrawler"

// StartCrawlerRequest returns a request value for making API operation for
// AWS Glue.
//
// Starts a crawl using the specified crawler, regardless of what is scheduled.
// If the crawler is already running, returns a CrawlerRunningException (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-exceptions.html#aws-glue-api-exceptions-CrawlerRunningException).
//
//    // Example sending a request using StartCrawlerRequest.
//    req := client.StartCrawlerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/StartCrawler
func (c *Client) StartCrawlerRequest(input *StartCrawlerInput) StartCrawlerRequest {
	op := &aws.Operation{
		Name:       opStartCrawler,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartCrawlerInput{}
	}

	req := c.newRequest(op, input, &StartCrawlerOutput{})
	return StartCrawlerRequest{Request: req, Input: input, Copy: c.StartCrawlerRequest}
}

// StartCrawlerRequest is the request type for the
// StartCrawler API operation.
type StartCrawlerRequest struct {
	*aws.Request
	Input *StartCrawlerInput
	Copy  func(*StartCrawlerInput) StartCrawlerRequest
}

// Send marshals and sends the StartCrawler API request.
func (r StartCrawlerRequest) Send(ctx context.Context) (*StartCrawlerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartCrawlerResponse{
		StartCrawlerOutput: r.Request.Data.(*StartCrawlerOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartCrawlerResponse is the response type for the
// StartCrawler API operation.
type StartCrawlerResponse struct {
	*StartCrawlerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartCrawler request.
func (r *StartCrawlerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
