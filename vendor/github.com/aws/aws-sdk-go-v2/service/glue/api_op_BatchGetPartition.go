// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchGetPartitionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the partitions in question reside. If none
	// is supplied, the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database where the partitions reside.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// A list of partition values identifying the partitions to retrieve.
	//
	// PartitionsToGet is a required field
	PartitionsToGet []PartitionValueList `type:"list" required:"true"`

	// The name of the partitions' table.
	//
	// TableName is a required field
	TableName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s BatchGetPartitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetPartitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetPartitionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.PartitionsToGet == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartitionsToGet"))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 1))
	}
	if s.PartitionsToGet != nil {
		for i, v := range s.PartitionsToGet {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PartitionsToGet", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchGetPartitionOutput struct {
	_ struct{} `type:"structure"`

	// A list of the requested partitions.
	Partitions []Partition `type:"list"`

	// A list of the partition values in the request for which partitions were not
	// returned.
	UnprocessedKeys []PartitionValueList `type:"list"`
}

// String returns the string representation
func (s BatchGetPartitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetPartition = "BatchGetPartition"

// BatchGetPartitionRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves partitions in a batch request.
//
//    // Example sending a request using BatchGetPartitionRequest.
//    req := client.BatchGetPartitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchGetPartition
func (c *Client) BatchGetPartitionRequest(input *BatchGetPartitionInput) BatchGetPartitionRequest {
	op := &aws.Operation{
		Name:       opBatchGetPartition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetPartitionInput{}
	}

	req := c.newRequest(op, input, &BatchGetPartitionOutput{})
	return BatchGetPartitionRequest{Request: req, Input: input, Copy: c.BatchGetPartitionRequest}
}

// BatchGetPartitionRequest is the request type for the
// BatchGetPartition API operation.
type BatchGetPartitionRequest struct {
	*aws.Request
	Input *BatchGetPartitionInput
	Copy  func(*BatchGetPartitionInput) BatchGetPartitionRequest
}

// Send marshals and sends the BatchGetPartition API request.
func (r BatchGetPartitionRequest) Send(ctx context.Context) (*BatchGetPartitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetPartitionResponse{
		BatchGetPartitionOutput: r.Request.Data.(*BatchGetPartitionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetPartitionResponse is the response type for the
// BatchGetPartition API operation.
type BatchGetPartitionResponse struct {
	*BatchGetPartitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetPartition request.
func (r *BatchGetPartitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
