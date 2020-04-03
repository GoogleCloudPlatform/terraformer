// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to get the header information for the .csv file for
// the user import job.
type GetCSVHeaderInput struct {
	_ struct{} `type:"structure"`

	// The user pool ID for the user pool that the users are to be imported into.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCSVHeaderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCSVHeaderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCSVHeaderInput"}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server to the request to get the header
// information for the .csv file for the user import job.
type GetCSVHeaderOutput struct {
	_ struct{} `type:"structure"`

	// The header information for the .csv file for the user import job.
	CSVHeader []string `type:"list"`

	// The user pool ID for the user pool that the users are to be imported into.
	UserPoolId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetCSVHeaderOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCSVHeader = "GetCSVHeader"

// GetCSVHeaderRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Gets the header information for the .csv file to be used as input for the
// user import job.
//
//    // Example sending a request using GetCSVHeaderRequest.
//    req := client.GetCSVHeaderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetCSVHeader
func (c *Client) GetCSVHeaderRequest(input *GetCSVHeaderInput) GetCSVHeaderRequest {
	op := &aws.Operation{
		Name:       opGetCSVHeader,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCSVHeaderInput{}
	}

	req := c.newRequest(op, input, &GetCSVHeaderOutput{})
	return GetCSVHeaderRequest{Request: req, Input: input, Copy: c.GetCSVHeaderRequest}
}

// GetCSVHeaderRequest is the request type for the
// GetCSVHeader API operation.
type GetCSVHeaderRequest struct {
	*aws.Request
	Input *GetCSVHeaderInput
	Copy  func(*GetCSVHeaderInput) GetCSVHeaderRequest
}

// Send marshals and sends the GetCSVHeader API request.
func (r GetCSVHeaderRequest) Send(ctx context.Context) (*GetCSVHeaderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCSVHeaderResponse{
		GetCSVHeaderOutput: r.Request.Data.(*GetCSVHeaderOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCSVHeaderResponse is the response type for the
// GetCSVHeader API operation.
type GetCSVHeaderResponse struct {
	*GetCSVHeaderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCSVHeader request.
func (r *GetCSVHeaderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
