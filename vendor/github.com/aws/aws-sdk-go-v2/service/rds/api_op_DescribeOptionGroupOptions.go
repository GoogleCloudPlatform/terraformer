// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeOptionGroupOptionsInput struct {
	_ struct{} `type:"structure"`

	// A required parameter. Options available for the given engine name are described.
	//
	// EngineName is a required field
	EngineName *string `type:"string" required:"true"`

	// This parameter isn't currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// If specified, filters the results to include only options for the specified
	// major engine version.
	MajorEngineVersion *string `type:"string"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeOptionGroupOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeOptionGroupOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeOptionGroupOptionsInput"}

	if s.EngineName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EngineName"))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeOptionGroupOptionsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// List of available option group options.
	OptionGroupOptions []OptionGroupOption `locationNameList:"OptionGroupOption" type:"list"`
}

// String returns the string representation
func (s DescribeOptionGroupOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeOptionGroupOptions = "DescribeOptionGroupOptions"

// DescribeOptionGroupOptionsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Describes all available options.
//
//    // Example sending a request using DescribeOptionGroupOptionsRequest.
//    req := client.DescribeOptionGroupOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeOptionGroupOptions
func (c *Client) DescribeOptionGroupOptionsRequest(input *DescribeOptionGroupOptionsInput) DescribeOptionGroupOptionsRequest {
	op := &aws.Operation{
		Name:       opDescribeOptionGroupOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeOptionGroupOptionsInput{}
	}

	req := c.newRequest(op, input, &DescribeOptionGroupOptionsOutput{})
	return DescribeOptionGroupOptionsRequest{Request: req, Input: input, Copy: c.DescribeOptionGroupOptionsRequest}
}

// DescribeOptionGroupOptionsRequest is the request type for the
// DescribeOptionGroupOptions API operation.
type DescribeOptionGroupOptionsRequest struct {
	*aws.Request
	Input *DescribeOptionGroupOptionsInput
	Copy  func(*DescribeOptionGroupOptionsInput) DescribeOptionGroupOptionsRequest
}

// Send marshals and sends the DescribeOptionGroupOptions API request.
func (r DescribeOptionGroupOptionsRequest) Send(ctx context.Context) (*DescribeOptionGroupOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOptionGroupOptionsResponse{
		DescribeOptionGroupOptionsOutput: r.Request.Data.(*DescribeOptionGroupOptionsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeOptionGroupOptionsRequestPaginator returns a paginator for DescribeOptionGroupOptions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeOptionGroupOptionsRequest(input)
//   p := rds.NewDescribeOptionGroupOptionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeOptionGroupOptionsPaginator(req DescribeOptionGroupOptionsRequest) DescribeOptionGroupOptionsPaginator {
	return DescribeOptionGroupOptionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeOptionGroupOptionsInput
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

// DescribeOptionGroupOptionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeOptionGroupOptionsPaginator struct {
	aws.Pager
}

func (p *DescribeOptionGroupOptionsPaginator) CurrentPage() *DescribeOptionGroupOptionsOutput {
	return p.Pager.CurrentPage().(*DescribeOptionGroupOptionsOutput)
}

// DescribeOptionGroupOptionsResponse is the response type for the
// DescribeOptionGroupOptions API operation.
type DescribeOptionGroupOptionsResponse struct {
	*DescribeOptionGroupOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOptionGroupOptions request.
func (r *DescribeOptionGroupOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
