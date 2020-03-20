// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to get information about a collection of BasePathMapping resources.
type GetBasePathMappingsInput struct {
	_ struct{} `type:"structure"`

	// [Required] The domain name of a BasePathMapping resource.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domain_name" type:"string" required:"true"`

	// The maximum number of returned results per page. The default value is 25
	// and the maximum value is 500.
	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	// The current pagination position in the paged result set.
	Position *string `location:"querystring" locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetBasePathMappingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBasePathMappingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBasePathMappingsInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBasePathMappingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domain_name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents a collection of BasePathMapping resources.
//
// Use Custom Domain Names (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type GetBasePathMappingsOutput struct {
	_ struct{} `type:"structure"`

	// The current page of elements from this collection.
	Items []BasePathMapping `locationName:"item" type:"list"`

	Position *string `locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetBasePathMappingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBasePathMappingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Items != nil {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "item", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetBasePathMappings = "GetBasePathMappings"

// GetBasePathMappingsRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Represents a collection of BasePathMapping resources.
//
//    // Example sending a request using GetBasePathMappingsRequest.
//    req := client.GetBasePathMappingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetBasePathMappingsRequest(input *GetBasePathMappingsInput) GetBasePathMappingsRequest {
	op := &aws.Operation{
		Name:       opGetBasePathMappings,
		HTTPMethod: "GET",
		HTTPPath:   "/domainnames/{domain_name}/basepathmappings",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"position"},
			OutputTokens:    []string{"position"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetBasePathMappingsInput{}
	}

	req := c.newRequest(op, input, &GetBasePathMappingsOutput{})
	return GetBasePathMappingsRequest{Request: req, Input: input, Copy: c.GetBasePathMappingsRequest}
}

// GetBasePathMappingsRequest is the request type for the
// GetBasePathMappings API operation.
type GetBasePathMappingsRequest struct {
	*aws.Request
	Input *GetBasePathMappingsInput
	Copy  func(*GetBasePathMappingsInput) GetBasePathMappingsRequest
}

// Send marshals and sends the GetBasePathMappings API request.
func (r GetBasePathMappingsRequest) Send(ctx context.Context) (*GetBasePathMappingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBasePathMappingsResponse{
		GetBasePathMappingsOutput: r.Request.Data.(*GetBasePathMappingsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetBasePathMappingsRequestPaginator returns a paginator for GetBasePathMappings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetBasePathMappingsRequest(input)
//   p := apigateway.NewGetBasePathMappingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetBasePathMappingsPaginator(req GetBasePathMappingsRequest) GetBasePathMappingsPaginator {
	return GetBasePathMappingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetBasePathMappingsInput
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

// GetBasePathMappingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetBasePathMappingsPaginator struct {
	aws.Pager
}

func (p *GetBasePathMappingsPaginator) CurrentPage() *GetBasePathMappingsOutput {
	return p.Pager.CurrentPage().(*GetBasePathMappingsOutput)
}

// GetBasePathMappingsResponse is the response type for the
// GetBasePathMappings API operation.
type GetBasePathMappingsResponse struct {
	*GetBasePathMappingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBasePathMappings request.
func (r *GetBasePathMappingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
