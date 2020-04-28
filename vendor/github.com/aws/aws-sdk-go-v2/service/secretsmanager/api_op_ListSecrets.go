// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSecretsInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Limits the number of results that you want to include in the response.
	// If you don't include this parameter, it defaults to a value that's specific
	// to the operation. If additional items exist beyond the maximum you specify,
	// the NextToken response element is present and has a value (isn't null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Secrets Manager might return
	// fewer results than the maximum even when there are more results available.
	// You should check NextToken after every operation to ensure that you receive
	// all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// (Optional) Use this parameter in a request if you receive a NextToken response
	// in a previous request that indicates that there's more output available.
	// In a subsequent call, set it to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListSecretsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSecretsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSecretsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListSecretsOutput struct {
	_ struct{} `type:"structure"`

	// If present in the response, this value indicates that there's more output
	// available than what's included in the current response. This can occur even
	// when the response includes no values at all, such as when you ask for a filtered
	// view of a very long list. Use this value in the NextToken request parameter
	// in a subsequent call to the operation to continue processing and get the
	// next part of the output. You should repeat this until the NextToken response
	// element comes back empty (as null).
	NextToken *string `min:"1" type:"string"`

	// A list of the secrets in the account.
	SecretList []SecretListEntry `type:"list"`
}

// String returns the string representation
func (s ListSecretsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSecrets = "ListSecrets"

// ListSecretsRequest returns a request value for making API operation for
// AWS Secrets Manager.
//
// Lists all of the secrets that are stored by Secrets Manager in the AWS account.
// To list the versions currently stored for a specific secret, use ListSecretVersionIds.
// The encrypted fields SecretString and SecretBinary are not included in the
// output. To get that information, call the GetSecretValue operation.
//
// Always check the NextToken response parameter when calling any of the List*
// operations. These operations can occasionally return an empty or shorter
// than expected list of results even when there are more results available.
// When this happens, the NextToken response parameter contains a value to pass
// to the next call to the same API to request the next part of the list.
//
// Minimum permissions
//
// To run this command, you must have the following permissions:
//
//    * secretsmanager:ListSecrets
//
// Related operations
//
//    * To list the versions attached to a secret, use ListSecretVersionIds.
//
//    // Example sending a request using ListSecretsRequest.
//    req := client.ListSecretsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/ListSecrets
func (c *Client) ListSecretsRequest(input *ListSecretsInput) ListSecretsRequest {
	op := &aws.Operation{
		Name:       opListSecrets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListSecretsInput{}
	}

	req := c.newRequest(op, input, &ListSecretsOutput{})
	return ListSecretsRequest{Request: req, Input: input, Copy: c.ListSecretsRequest}
}

// ListSecretsRequest is the request type for the
// ListSecrets API operation.
type ListSecretsRequest struct {
	*aws.Request
	Input *ListSecretsInput
	Copy  func(*ListSecretsInput) ListSecretsRequest
}

// Send marshals and sends the ListSecrets API request.
func (r ListSecretsRequest) Send(ctx context.Context) (*ListSecretsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSecretsResponse{
		ListSecretsOutput: r.Request.Data.(*ListSecretsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSecretsRequestPaginator returns a paginator for ListSecrets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSecretsRequest(input)
//   p := secretsmanager.NewListSecretsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSecretsPaginator(req ListSecretsRequest) ListSecretsPaginator {
	return ListSecretsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSecretsInput
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

// ListSecretsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSecretsPaginator struct {
	aws.Pager
}

func (p *ListSecretsPaginator) CurrentPage() *ListSecretsOutput {
	return p.Pager.CurrentPage().(*ListSecretsOutput)
}

// ListSecretsResponse is the response type for the
// ListSecrets API operation.
type ListSecretsResponse struct {
	*ListSecretsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSecrets request.
func (r *ListSecretsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
