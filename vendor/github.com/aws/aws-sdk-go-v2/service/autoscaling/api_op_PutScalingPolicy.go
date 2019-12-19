// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether the ScalingAdjustment parameter is an absolute number or
	// a percentage of the current capacity. The valid values are ChangeInCapacity,
	// ExactCapacity, and PercentChangeInCapacity.
	//
	// Valid only if the policy type is StepScaling or SimpleScaling. For more information,
	// see Scaling Adjustment Types (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment)
	// in the Amazon EC2 Auto Scaling User Guide.
	AdjustmentType *string `min:"1" type:"string"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// The amount of time, in seconds, after a scaling activity completes before
	// any further dynamic scaling activities can start. If this parameter is not
	// specified, the default cooldown period for the group applies.
	//
	// Valid only if the policy type is SimpleScaling. For more information, see
	// Scaling Cooldowns (https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	Cooldown *int64 `type:"integer"`

	// The estimated time, in seconds, until a newly launched instance can contribute
	// to the CloudWatch metrics. The default is to use the value specified for
	// the default cooldown period for the group.
	//
	// Valid only if the policy type is StepScaling or TargetTrackingScaling.
	EstimatedInstanceWarmup *int64 `type:"integer"`

	// The aggregation type for the CloudWatch metrics. The valid values are Minimum,
	// Maximum, and Average. If the aggregation type is null, the value is treated
	// as Average.
	//
	// Valid only if the policy type is StepScaling.
	MetricAggregationType *string `min:"1" type:"string"`

	// The minimum number of instances to scale. If the value of AdjustmentType
	// is PercentChangeInCapacity, the scaling policy changes the DesiredCapacity
	// of the Auto Scaling group by at least this many instances. Otherwise, the
	// error is ValidationError.
	//
	// This property replaces the MinAdjustmentStep property. For example, suppose
	// that you create a step scaling policy to scale out an Auto Scaling group
	// by 25 percent and you specify a MinAdjustmentMagnitude of 2. If the group
	// has 4 instances and the scaling policy is performed, 25 percent of 4 is 1.
	// However, because you specified a MinAdjustmentMagnitude of 2, Amazon EC2
	// Auto Scaling scales out the group by 2 instances.
	//
	// Valid only if the policy type is SimpleScaling or StepScaling.
	MinAdjustmentMagnitude *int64 `type:"integer"`

	// Available for backward compatibility. Use MinAdjustmentMagnitude instead.
	MinAdjustmentStep *int64 `deprecated:"true" type:"integer"`

	// The name of the policy.
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`

	// The policy type. The valid values are SimpleScaling, StepScaling, and TargetTrackingScaling.
	// If the policy type is null, the value is treated as SimpleScaling.
	PolicyType *string `min:"1" type:"string"`

	// The amount by which a simple scaling policy scales the Auto Scaling group
	// in response to an alarm breach. The adjustment is based on the value that
	// you specified in the AdjustmentType parameter (either an absolute number
	// or a percentage). A positive value adds to the current capacity and a negative
	// value subtracts from the current capacity. For exact capacity, you must specify
	// a positive value.
	//
	// Conditional: If you specify SimpleScaling for the policy type, you must specify
	// this parameter. (Not used with any other policy type.)
	ScalingAdjustment *int64 `type:"integer"`

	// A set of adjustments that enable you to scale based on the size of the alarm
	// breach.
	//
	// Conditional: If you specify StepScaling for the policy type, you must specify
	// this parameter. (Not used with any other policy type.)
	StepAdjustments []StepAdjustment `type:"list"`

	// A target tracking scaling policy. Includes support for predefined or customized
	// metrics.
	//
	// For more information, see TargetTrackingConfiguration (https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_TargetTrackingConfiguration.html)
	// in the Amazon EC2 Auto Scaling API Reference.
	//
	// Conditional: If you specify TargetTrackingScaling for the policy type, you
	// must specify this parameter. (Not used with any other policy type.)
	TargetTrackingConfiguration *TargetTrackingConfiguration `type:"structure"`
}

// String returns the string representation
func (s PutScalingPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutScalingPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutScalingPolicyInput"}
	if s.AdjustmentType != nil && len(*s.AdjustmentType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AdjustmentType", 1))
	}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}
	if s.MetricAggregationType != nil && len(*s.MetricAggregationType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MetricAggregationType", 1))
	}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}
	if s.PolicyType != nil && len(*s.PolicyType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyType", 1))
	}
	if s.StepAdjustments != nil {
		for i, v := range s.StepAdjustments {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "StepAdjustments", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.TargetTrackingConfiguration != nil {
		if err := s.TargetTrackingConfiguration.Validate(); err != nil {
			invalidParams.AddNested("TargetTrackingConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of PutScalingPolicy.
type PutScalingPolicyOutput struct {
	_ struct{} `type:"structure"`

	// The CloudWatch alarms created for the target tracking scaling policy.
	Alarms []Alarm `type:"list"`

	// The Amazon Resource Name (ARN) of the policy.
	PolicyARN *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutScalingPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutScalingPolicy = "PutScalingPolicy"

// PutScalingPolicyRequest returns a request value for making API operation for
// Auto Scaling.
//
// Creates or updates a scaling policy for an Auto Scaling group. To update
// an existing scaling policy, use the existing policy name and set the parameters
// to change. Any existing parameter not changed in an update to an existing
// policy is not changed in this update request.
//
// For more information about using scaling policies to scale your Auto Scaling
// group automatically, see Dynamic Scaling (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scale-based-on-demand.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using PutScalingPolicyRequest.
//    req := client.PutScalingPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/PutScalingPolicy
func (c *Client) PutScalingPolicyRequest(input *PutScalingPolicyInput) PutScalingPolicyRequest {
	op := &aws.Operation{
		Name:       opPutScalingPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutScalingPolicyInput{}
	}

	req := c.newRequest(op, input, &PutScalingPolicyOutput{})
	return PutScalingPolicyRequest{Request: req, Input: input, Copy: c.PutScalingPolicyRequest}
}

// PutScalingPolicyRequest is the request type for the
// PutScalingPolicy API operation.
type PutScalingPolicyRequest struct {
	*aws.Request
	Input *PutScalingPolicyInput
	Copy  func(*PutScalingPolicyInput) PutScalingPolicyRequest
}

// Send marshals and sends the PutScalingPolicy API request.
func (r PutScalingPolicyRequest) Send(ctx context.Context) (*PutScalingPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutScalingPolicyResponse{
		PutScalingPolicyOutput: r.Request.Data.(*PutScalingPolicyOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutScalingPolicyResponse is the response type for the
// PutScalingPolicy API operation.
type PutScalingPolicyResponse struct {
	*PutScalingPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutScalingPolicy request.
func (r *PutScalingPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
