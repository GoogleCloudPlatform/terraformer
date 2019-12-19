// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartMLLabelingSetGenerationTaskRunInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Simple Storage Service (Amazon S3) path where you generate the
	// labeling set.
	//
	// OutputS3Path is a required field
	OutputS3Path *string `type:"string" required:"true"`

	// The unique identifier of the machine learning transform.
	//
	// TransformId is a required field
	TransformId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartMLLabelingSetGenerationTaskRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartMLLabelingSetGenerationTaskRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartMLLabelingSetGenerationTaskRunInput"}

	if s.OutputS3Path == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputS3Path"))
	}

	if s.TransformId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransformId"))
	}
	if s.TransformId != nil && len(*s.TransformId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TransformId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartMLLabelingSetGenerationTaskRunOutput struct {
	_ struct{} `type:"structure"`

	// The unique run identifier that is associated with this task run.
	TaskRunId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartMLLabelingSetGenerationTaskRunOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartMLLabelingSetGenerationTaskRun = "StartMLLabelingSetGenerationTaskRun"

// StartMLLabelingSetGenerationTaskRunRequest returns a request value for making API operation for
// AWS Glue.
//
// Starts the active learning workflow for your machine learning transform to
// improve the transform's quality by generating label sets and adding labels.
//
// When the StartMLLabelingSetGenerationTaskRun finishes, AWS Glue will have
// generated a "labeling set" or a set of questions for humans to answer.
//
// In the case of the FindMatches transform, these questions are of the form,
// “What is the correct way to group these rows together into groups composed
// entirely of matching records?”
//
// After the labeling process is finished, you can upload your labels with a
// call to StartImportLabelsTaskRun. After StartImportLabelsTaskRun finishes,
// all future runs of the machine learning transform will use the new and improved
// labels and perform a higher-quality transformation.
//
//    // Example sending a request using StartMLLabelingSetGenerationTaskRunRequest.
//    req := client.StartMLLabelingSetGenerationTaskRunRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/StartMLLabelingSetGenerationTaskRun
func (c *Client) StartMLLabelingSetGenerationTaskRunRequest(input *StartMLLabelingSetGenerationTaskRunInput) StartMLLabelingSetGenerationTaskRunRequest {
	op := &aws.Operation{
		Name:       opStartMLLabelingSetGenerationTaskRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartMLLabelingSetGenerationTaskRunInput{}
	}

	req := c.newRequest(op, input, &StartMLLabelingSetGenerationTaskRunOutput{})
	return StartMLLabelingSetGenerationTaskRunRequest{Request: req, Input: input, Copy: c.StartMLLabelingSetGenerationTaskRunRequest}
}

// StartMLLabelingSetGenerationTaskRunRequest is the request type for the
// StartMLLabelingSetGenerationTaskRun API operation.
type StartMLLabelingSetGenerationTaskRunRequest struct {
	*aws.Request
	Input *StartMLLabelingSetGenerationTaskRunInput
	Copy  func(*StartMLLabelingSetGenerationTaskRunInput) StartMLLabelingSetGenerationTaskRunRequest
}

// Send marshals and sends the StartMLLabelingSetGenerationTaskRun API request.
func (r StartMLLabelingSetGenerationTaskRunRequest) Send(ctx context.Context) (*StartMLLabelingSetGenerationTaskRunResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartMLLabelingSetGenerationTaskRunResponse{
		StartMLLabelingSetGenerationTaskRunOutput: r.Request.Data.(*StartMLLabelingSetGenerationTaskRunOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartMLLabelingSetGenerationTaskRunResponse is the response type for the
// StartMLLabelingSetGenerationTaskRun API operation.
type StartMLLabelingSetGenerationTaskRunResponse struct {
	*StartMLLabelingSetGenerationTaskRunOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartMLLabelingSetGenerationTaskRun request.
func (r *StartMLLabelingSetGenerationTaskRunResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
