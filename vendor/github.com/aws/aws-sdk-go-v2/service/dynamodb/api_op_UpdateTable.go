// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of an UpdateTable operation.
type UpdateTableInput struct {
	_ struct{} `type:"structure"`

	// An array of attributes that describe the key schema for the table and indexes.
	// If you are adding a new global secondary index to the table, AttributeDefinitions
	// must include the key element(s) of the new index.
	AttributeDefinitions []AttributeDefinition `type:"list"`

	// Controls how you are charged for read and write throughput and how you manage
	// capacity. When switching from pay-per-request to provisioned capacity, initial
	// provisioned capacity values must be set. The initial provisioned capacity
	// values are estimated based on the consumed read and write capacity of your
	// table and global secondary indexes over the past 30 minutes.
	//
	//    * PROVISIONED - We recommend using PROVISIONED for predictable workloads.
	//    PROVISIONED sets the billing mode to Provisioned Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).
	//
	//    * PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable
	//    workloads. PAY_PER_REQUEST sets the billing mode to On-Demand Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).
	BillingMode BillingMode `type:"string" enum:"true"`

	// An array of one or more global secondary indexes for the table. For each
	// index in the array, you can request one action:
	//
	//    * Create - add a new global secondary index to the table.
	//
	//    * Update - modify the provisioned throughput settings of an existing global
	//    secondary index.
	//
	//    * Delete - remove a global secondary index from the table.
	//
	// You can create or delete only one global secondary index per UpdateTable
	// operation.
	//
	// For more information, see Managing Global Secondary Indexes (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.OnlineOps.html)
	// in the Amazon DynamoDB Developer Guide.
	GlobalSecondaryIndexUpdates []GlobalSecondaryIndexUpdate `type:"list"`

	// The new provisioned throughput settings for the specified table or index.
	ProvisionedThroughput *ProvisionedThroughput `type:"structure"`

	// A list of replica update actions (create, delete, or update) for the table.
	//
	// This property only applies to Version 2019.11.21 (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html)
	// of global tables.
	ReplicaUpdates []ReplicationGroupUpdate `min:"1" type:"list"`

	// The new server-side encryption settings for the specified table.
	SSESpecification *SSESpecification `type:"structure"`

	// Represents the DynamoDB Streams configuration for the table.
	//
	// You receive a ResourceInUseException if you try to enable a stream on a table
	// that already has a stream, or if you try to disable a stream on a table that
	// doesn't have a stream.
	StreamSpecification *StreamSpecification `type:"structure"`

	// The name of the table to be updated.
	//
	// TableName is a required field
	TableName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTableInput"}
	if s.ReplicaUpdates != nil && len(s.ReplicaUpdates) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ReplicaUpdates", 1))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 3))
	}
	if s.AttributeDefinitions != nil {
		for i, v := range s.AttributeDefinitions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AttributeDefinitions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.GlobalSecondaryIndexUpdates != nil {
		for i, v := range s.GlobalSecondaryIndexUpdates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GlobalSecondaryIndexUpdates", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.ProvisionedThroughput != nil {
		if err := s.ProvisionedThroughput.Validate(); err != nil {
			invalidParams.AddNested("ProvisionedThroughput", err.(aws.ErrInvalidParams))
		}
	}
	if s.ReplicaUpdates != nil {
		for i, v := range s.ReplicaUpdates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ReplicaUpdates", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.StreamSpecification != nil {
		if err := s.StreamSpecification.Validate(); err != nil {
			invalidParams.AddNested("StreamSpecification", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of an UpdateTable operation.
type UpdateTableOutput struct {
	_ struct{} `type:"structure"`

	// Represents the properties of the table.
	TableDescription *TableDescription `type:"structure"`
}

// String returns the string representation
func (s UpdateTableOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTable = "UpdateTable"

// UpdateTableRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Modifies the provisioned throughput settings, global secondary indexes, or
// DynamoDB Streams settings for a given table.
//
// You can only perform one of the following operations at once:
//
//    * Modify the provisioned throughput settings of the table.
//
//    * Enable or disable DynamoDB Streams on the table.
//
//    * Remove a global secondary index from the table.
//
//    * Create a new global secondary index on the table. After the index begins
//    backfilling, you can use UpdateTable to perform other operations.
//
// UpdateTable is an asynchronous operation; while it is executing, the table
// status changes from ACTIVE to UPDATING. While it is UPDATING, you cannot
// issue another UpdateTable request. When the table returns to the ACTIVE state,
// the UpdateTable operation is complete.
//
//    // Example sending a request using UpdateTableRequest.
//    req := client.UpdateTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/UpdateTable
func (c *Client) UpdateTableRequest(input *UpdateTableInput) UpdateTableRequest {
	op := &aws.Operation{
		Name:       opUpdateTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTableInput{}
	}

	req := c.newRequest(op, input, &UpdateTableOutput{})

	if req.Config.EnableEndpointDiscovery {
		de := discovererDescribeEndpoints{
			Client:        c,
			Required:      false,
			EndpointCache: c.endpointCache,
			Params: map[string]*string{
				"op": &req.Operation.Name,
			},
		}

		for k, v := range de.Params {
			if v == nil {
				delete(de.Params, k)
			}
		}

		req.Handlers.Build.PushFrontNamed(aws.NamedHandler{
			Name: "crr.endpointdiscovery",
			Fn:   de.Handler,
		})
	}
	return UpdateTableRequest{Request: req, Input: input, Copy: c.UpdateTableRequest}
}

// UpdateTableRequest is the request type for the
// UpdateTable API operation.
type UpdateTableRequest struct {
	*aws.Request
	Input *UpdateTableInput
	Copy  func(*UpdateTableInput) UpdateTableRequest
}

// Send marshals and sends the UpdateTable API request.
func (r UpdateTableRequest) Send(ctx context.Context) (*UpdateTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTableResponse{
		UpdateTableOutput: r.Request.Data.(*UpdateTableOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTableResponse is the response type for the
// UpdateTable API operation.
type UpdateTableResponse struct {
	*UpdateTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTable request.
func (r *UpdateTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
