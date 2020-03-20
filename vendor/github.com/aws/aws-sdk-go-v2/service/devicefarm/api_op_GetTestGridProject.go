// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetTestGridProjectInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the Selenium testing project, from either CreateTestGridProject
	// or ListTestGridProjects.
	//
	// ProjectArn is a required field
	ProjectArn *string `locationName:"projectArn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTestGridProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTestGridProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTestGridProjectInput"}

	if s.ProjectArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectArn"))
	}
	if s.ProjectArn != nil && len(*s.ProjectArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectArn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTestGridProjectOutput struct {
	_ struct{} `type:"structure"`

	// A TestGridProject.
	TestGridProject *TestGridProject `locationName:"testGridProject" type:"structure"`
}

// String returns the string representation
func (s GetTestGridProjectOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTestGridProject = "GetTestGridProject"

// GetTestGridProjectRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Retrieves information about a Selenium testing project.
//
//    // Example sending a request using GetTestGridProjectRequest.
//    req := client.GetTestGridProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetTestGridProject
func (c *Client) GetTestGridProjectRequest(input *GetTestGridProjectInput) GetTestGridProjectRequest {
	op := &aws.Operation{
		Name:       opGetTestGridProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTestGridProjectInput{}
	}

	req := c.newRequest(op, input, &GetTestGridProjectOutput{})
	return GetTestGridProjectRequest{Request: req, Input: input, Copy: c.GetTestGridProjectRequest}
}

// GetTestGridProjectRequest is the request type for the
// GetTestGridProject API operation.
type GetTestGridProjectRequest struct {
	*aws.Request
	Input *GetTestGridProjectInput
	Copy  func(*GetTestGridProjectInput) GetTestGridProjectRequest
}

// Send marshals and sends the GetTestGridProject API request.
func (r GetTestGridProjectRequest) Send(ctx context.Context) (*GetTestGridProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTestGridProjectResponse{
		GetTestGridProjectOutput: r.Request.Data.(*GetTestGridProjectOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTestGridProjectResponse is the response type for the
// GetTestGridProject API operation.
type GetTestGridProjectResponse struct {
	*GetTestGridProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTestGridProject request.
func (r *GetTestGridProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
