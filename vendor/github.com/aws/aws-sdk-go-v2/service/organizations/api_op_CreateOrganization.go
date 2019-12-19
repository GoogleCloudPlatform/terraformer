// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateOrganizationInput struct {
	_ struct{} `type:"structure"`

	// Specifies the feature set supported by the new organization. Each feature
	// set supports different levels of functionality.
	//
	//    * CONSOLIDATED_BILLING: All member accounts have their bills consolidated
	//    to and paid by the master account. For more information, see Consolidated
	//    billing (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#feature-set-cb-only)
	//    in the AWS Organizations User Guide. The consolidated billing feature
	//    subset isn't available for organizations in the AWS GovCloud (US) Region.
	//
	//    * ALL: In addition to all the features supported by the consolidated billing
	//    feature set, the master account can also apply any policy type to any
	//    member account in the organization. For more information, see All features
	//    (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#feature-set-all)
	//    in the AWS Organizations User Guide.
	FeatureSet OrganizationFeatureSet `type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

type CreateOrganizationOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the newly created organization.
	Organization *Organization `type:"structure"`
}

// String returns the string representation
func (s CreateOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateOrganization = "CreateOrganization"

// CreateOrganizationRequest returns a request value for making API operation for
// AWS Organizations.
//
// Creates an AWS organization. The account whose user is calling the CreateOrganization
// operation automatically becomes the master account (https://docs.aws.amazon.com/IAM/latest/UserGuide/orgs_getting-started_concepts.html#account)
// of the new organization.
//
// This operation must be called using credentials from the account that is
// to become the new organization's master account. The principal must also
// have the relevant IAM permissions.
//
// By default (or if you set the FeatureSet parameter to ALL), the new organization
// is created with all features enabled and service control policies automatically
// enabled in the root. If you instead choose to create the organization supporting
// only the consolidated billing features by setting the FeatureSet parameter
// to CONSOLIDATED_BILLING", no policy types are enabled by default, and you
// can't use organization policies.
//
//    // Example sending a request using CreateOrganizationRequest.
//    req := client.CreateOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/CreateOrganization
func (c *Client) CreateOrganizationRequest(input *CreateOrganizationInput) CreateOrganizationRequest {
	op := &aws.Operation{
		Name:       opCreateOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateOrganizationInput{}
	}

	req := c.newRequest(op, input, &CreateOrganizationOutput{})
	return CreateOrganizationRequest{Request: req, Input: input, Copy: c.CreateOrganizationRequest}
}

// CreateOrganizationRequest is the request type for the
// CreateOrganization API operation.
type CreateOrganizationRequest struct {
	*aws.Request
	Input *CreateOrganizationInput
	Copy  func(*CreateOrganizationInput) CreateOrganizationRequest
}

// Send marshals and sends the CreateOrganization API request.
func (r CreateOrganizationRequest) Send(ctx context.Context) (*CreateOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateOrganizationResponse{
		CreateOrganizationOutput: r.Request.Data.(*CreateOrganizationOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateOrganizationResponse is the response type for the
// CreateOrganization API operation.
type CreateOrganizationResponse struct {
	*CreateOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateOrganization request.
func (r *CreateOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
